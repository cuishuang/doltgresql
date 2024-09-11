// Copyright 2023 Dolthub, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"sync/atomic"

	"github.com/dolthub/dolt/go/libraries/doltcore/sqlserver"
	"github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/expression"
	"github.com/dolthub/go-mysql-server/sql/plan"
	"github.com/dolthub/go-mysql-server/sql/transform"
	"github.com/dolthub/vitess/go/mysql"
	"github.com/dolthub/vitess/go/sqltypes"
	"github.com/dolthub/vitess/go/vt/sqlparser"
	"github.com/jackc/pgx/v5/pgproto3"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/lib/pq/oid"
	"github.com/sirupsen/logrus"

	"github.com/dolthub/doltgresql/core"
	"github.com/dolthub/doltgresql/core/dataloader"
	"github.com/dolthub/doltgresql/postgres/connection"
	"github.com/dolthub/doltgresql/postgres/messages"
	"github.com/dolthub/doltgresql/postgres/parser/parser"
	"github.com/dolthub/doltgresql/server/ast"
	pgexprs "github.com/dolthub/doltgresql/server/expression"
	"github.com/dolthub/doltgresql/server/node"
	pgtypes "github.com/dolthub/doltgresql/server/types"
)

// ConnectionHandler is responsible for the entire lifecycle of a user connection: receiving messages they send,
// executing queries, sending the correct messages in return, and terminating the connection when appropriate.
type ConnectionHandler struct {
	mysqlConn          *mysql.Conn
	preparedStatements map[string]PreparedStatementData
	portals            map[string]PortalData
	doltgresHandler    *Handler
	mysqlHandler       mysql.Handler
	pgTypeMap          *pgtype.Map
	waitForSync        bool
	// copyFromStdinState is set when this connection is in the COPY FROM STDIN mode, meaning it is waiting on
	// COPY DATA messages from the client to import data into tables.
	copyFromStdinState *copyFromStdinState
}

// copyFromStdinState tracks the metadata for an import of data into a table using a COPY FROM STDIN statement. When
// this statement is processed, the server accepts COPY DATA messages from the client with chunks of data to load
// into a table.
type copyFromStdinState struct {
	copyFromStdinNode *node.CopyFrom
	dataLoader        *dataloader.TabularDataLoader
}

// Set this env var to disable panic handling in the connection, which is useful when debugging a panic
const disablePanicHandlingEnvVar = "DOLT_PGSQL_PANIC"

// handlePanics determines whether panics should be handled in the connection handler. See |disablePanicHandlingEnvVar|.
var handlePanics = true

func init() {
	if _, ok := os.LookupEnv(disablePanicHandlingEnvVar); ok {
		handlePanics = false
	}
}

// NewConnectionHandler returns a new ConnectionHandler for the connection provided
func NewConnectionHandler(conn net.Conn, handler mysql.Handler) *ConnectionHandler {
	mysqlConn := &mysql.Conn{
		Conn:        conn,
		PrepareData: make(map[uint32]*mysql.PrepareData),
	}
	mysqlConn.ConnectionID = atomic.AddUint32(&connectionIDCounter, 1)

	// Postgres has a two-stage procedure for prepared queries. First the query is parsed via a |Parse| message, and
	// the result is stored in the |preparedStatements| map by the name provided. Then one or more |Bind| messages
	// provide parameters for the query, and the result is stored in |portals|. Finally, a call to |Execute| executes
	// the named portal.
	preparedStatements := make(map[string]PreparedStatementData)
	portals := make(map[string]PortalData)

	// TODO: how to expose these to create our own handler?
	server := sqlserver.GetRunningServer()
	h := &Handler{
		e:                 server.Engine,
		sm:                server.SessionManager(),
		readTimeout:       0,     // cfg.ConnReadTimeout,
		disableMultiStmts: false, // cfg.DisableClientMultiStatements,
		maxLoggedQueryLen: 0,     // cfg.MaxLoggedQueryLen, ???
		encodeLoggedQuery: false, // cfg.EncodeLoggedQuery,
		sel:               nil,   // TODO
		handler:           handler,
	}

	// TODO: should we use this backend???
	//pgproto3.NewBackend()
	return &ConnectionHandler{
		mysqlConn:          mysqlConn,
		preparedStatements: preparedStatements,
		portals:            portals,
		doltgresHandler:    h,
		mysqlHandler:       handler,
		pgTypeMap:          pgtype.NewMap(),
	}
}

// HandleConnection handles a connection's session, reading messages, executing queries, and sending responses.
// Expected to run in a goroutine per connection.
func (h *ConnectionHandler) HandleConnection() {
	var returnErr error
	if handlePanics {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Listener recovered panic: %v", r)

				var eomErr error
				if returnErr != nil {
					eomErr = returnErr
				} else if rErr, ok := r.(error); ok {
					eomErr = rErr
				} else {
					eomErr = fmt.Errorf("panic: %v", r)
				}

				// Sending eom can panic, which means we must recover again
				defer func() {
					if r := recover(); r != nil {
						fmt.Printf("Listener recovered panic: %v", r)
					}
				}()
				h.endOfMessages(eomErr)
			}

			if returnErr != nil {
				fmt.Println(returnErr.Error())
			}

			h.mysqlHandler.ConnectionClosed(h.mysqlConn)
			if err := h.Conn().Close(); err != nil {
				fmt.Printf("Failed to properly close connection:\n%v\n", err)
			}
		}()
	}
	h.mysqlHandler.NewConnection(h.mysqlConn)

	startupMessage, err := h.receiveStartupMessage()
	if err != nil {
		returnErr = err
		return
	}

	err = h.sendClientStartupMessages(startupMessage)
	if err != nil {
		returnErr = err
		return
	}

	err = h.chooseInitialDatabase(startupMessage)
	if err != nil {
		returnErr = err
		return
	}

	if err := connection.Send(h.Conn(), messages.ReadyForQuery{
		Indicator: messages.ReadyForQueryTransactionIndicator_Idle,
	}); err != nil {
		returnErr = err
		return
	}

	// Main session loop: read messages one at a time off the connection until we receive a |Terminate| message, in
	// which case we hang up, or the connection is closed by the client, which generates an io.EOF from the connection.
	for {
		stop, err := h.receiveMessage()
		if err != nil {
			returnErr = err
			break
		}

		if stop {
			break
		}
	}
}

// Conn returns the underlying net.Conn for this connection.
func (h *ConnectionHandler) Conn() net.Conn {
	return h.mysqlConn.Conn
}

// receiveMessage reads a single message off the connection and processes it, returning an error if no message could be
// received from the connection. Otherwise (a message is received successfully), the message is processed and any
// error is handled appropriately. The return value indicates whether the connection should be closed.
func (h *ConnectionHandler) receiveMessage() (bool, error) {
	var endOfMessages bool
	// For the time being, we handle panics in this function and treat them the same as errors so that they don't
	// forcibly close the connection. Contrast this with the panic handling logic in HandleConnection, where we treat any
	// panic as unrecoverable to the connection. As we fill out the implementation, we can revisit this decision and
	// rethink our posture over whether panics should terminate a connection.
	if handlePanics {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Listener recovered panic: %v", r)

				var eomErr error
				if rErr, ok := r.(error); ok {
					eomErr = rErr
				} else {
					eomErr = fmt.Errorf("panic: %v", r)
				}

				if !endOfMessages && h.waitForSync {
					if syncErr := connection.DiscardToSync(h.Conn()); syncErr != nil {
						fmt.Println(syncErr.Error())
					}
				}
				h.endOfMessages(eomErr)
			}
		}()
	}

	message, err := connection.Receive(h.Conn())
	if err != nil {
		return false, err
	}

	if ds, ok := message.(sql.DebugStringer); ok && logrus.IsLevelEnabled(logrus.DebugLevel) {
		logrus.Debugf("Received message: %s", ds.DebugString())
	} else {
		logrus.Debugf("Received message: %s", message.DefaultMessage().Name)
	}

	var stop bool
	stop, endOfMessages, err = h.handleMessage(message)
	if err != nil {
		if !endOfMessages && h.waitForSync {
			if syncErr := connection.DiscardToSync(h.Conn()); syncErr != nil {
				fmt.Println(syncErr.Error())
			}
		}
		h.endOfMessages(err)
	} else if endOfMessages {
		h.endOfMessages(nil)
	}

	return stop, nil
}

// receiveStartupMessage reads a startup message from the connection given and returns it. Some startup messages will
// result in the establishment of a new connection, which is also returned.
func (h *ConnectionHandler) receiveStartupMessage() (messages.StartupMessage, error) {
	var startupMessage messages.StartupMessage
	// The initial message may be one of a few different messages, so we'll check for those.
InitialMessageLoop:
	for {
		initialMessages, err := connection.ReceiveIntoAny(h.Conn(),
			messages.StartupMessage{},
			messages.SSLRequest{},
			messages.GSSENCRequest{})
		if err != nil {
			if err == io.EOF {
				return messages.StartupMessage{}, nil
			}
			return messages.StartupMessage{}, err
		}

		if len(initialMessages) != 1 {
			return messages.StartupMessage{}, fmt.Errorf("expected a single message upon starting connection, terminating connection")
		}

		initialMessage := initialMessages[0]
		switch initialMessage := initialMessage.(type) {
		case messages.StartupMessage:
			startupMessage = initialMessage
			break InitialMessageLoop
		case messages.SSLRequest:
			hasCertificate := len(certificate.Certificate) > 0
			if err := connection.Send(h.Conn(), messages.SSLResponse{
				SupportsSSL: hasCertificate,
			}); err != nil {
				return messages.StartupMessage{}, err
			}
			// If we have a certificate and the client has asked for SSL support, then we switch here.
			// This involves swapping out our underlying net connection for a new one.
			// We can't start in SSL mode, as the client does not attempt the handshake until after our response.
			if hasCertificate {
				conn := tls.Server(h.Conn(), &tls.Config{
					Certificates: []tls.Certificate{certificate},
				})
				h.mysqlConn.Conn = conn
			}
		case messages.GSSENCRequest:
			if err = connection.Send(h.Conn(), messages.GSSENCResponse{
				SupportsGSSAPI: false,
			}); err != nil {
				return messages.StartupMessage{}, err
			}
		default:
			return messages.StartupMessage{}, fmt.Errorf("unexpected initial message, terminating connection")
		}
	}

	return startupMessage, nil
}

// chooseInitialDatabase attempts to choose the initial database for the connection, if one is specified in the
// startup message provided
func (h *ConnectionHandler) chooseInitialDatabase(startupMessage messages.StartupMessage) error {
	if db, ok := startupMessage.Parameters["database"]; ok && len(db) > 0 {
		err := h.mysqlHandler.ComQuery(context.Background(), h.mysqlConn, fmt.Sprintf("USE `%s`;", db), func(res *sqltypes.Result, more bool) error {
			return nil
		})
		if err != nil {
			_ = connection.Send(h.Conn(), messages.ErrorResponse{
				Severity:     messages.ErrorResponseSeverity_Fatal,
				SqlStateCode: "3D000",
				Message:      fmt.Sprintf(`"database "%s" does not exist"`, db),
				Optional: messages.ErrorResponseOptionalFields{
					Routine: "InitPostgres",
				},
			})
			return err
		}
	} else {
		// If a database isn't specified, then we attempt to connect to a database with the same name as the user,
		// ignoring any error
		_ = h.mysqlHandler.ComQuery(context.Background(), h.mysqlConn, fmt.Sprintf("USE `%s`;", h.mysqlConn.User), func(res *sqltypes.Result, more bool) error {
			return nil
		})
	}
	return nil
}

// handleMessages processes the message provided and returns status flags indicating what the connection should do next.
// If the |stop| response parameter is true, it indicates that the connection should be closed by the caller. If the
// |endOfMessages| response parameter is true, it indicates that no more messages are expected for the current operation
// and a READY FOR QUERY message should be sent back to the client so it can send the next query.
func (h *ConnectionHandler) handleMessage(message connection.Message) (stop, endOfMessages bool, err error) {
	switch message := message.(type) {
	case messages.Terminate:
		return true, false, nil
	case messages.Sync:
		h.waitForSync = false
		return false, true, nil
	case messages.Query:
		endOfMessages, err = h.handleQuery(message)
		return false, endOfMessages, err
	case messages.Parse:
		return false, false, h.handleParse(message)
	case messages.Describe:
		return false, false, h.handleDescribe(message)
	case messages.Bind:
		return false, false, h.handleBind(message)
	case messages.Execute:
		return false, false, h.handleExecute(message)
	case messages.Close:
		if message.ClosingPreparedStatement {
			delete(h.preparedStatements, message.Target)
		} else {
			delete(h.portals, message.Target)
		}
		return false, false, connection.Send(h.Conn(), messages.CloseComplete{})
	case messages.CopyData:
		return h.handleCopyData(message)
	case messages.CopyDone:
		return h.handleCopyDone(message)
	case messages.CopyFail:
		return h.handleCopyFail(message)
	default:
		return false, true, fmt.Errorf(`Unhandled message "%s"`, message.DefaultMessage().Name)
	}
}

// handleQuery handles a query message, and returns a boolean flag, |endOfMessages| indicating if no other messages are
// expected as part of this query, in which case the server will send a READY FOR QUERY message back to the client so
// that it can send its next query.
func (h *ConnectionHandler) handleQuery(message messages.Query) (endOfMessages bool, err error) {
	handled, err := h.handledPSQLCommands(message.String)
	if handled || err != nil {
		return true, err
	}

	// TODO: Remove this once we support `SELECT * FROM function()` syntax
	// Github issue: https://github.com/dolthub/doltgresql/issues/464
	handled, err = h.handledWorkbenchCommands(message.String)
	if handled || err != nil {
		return true, err
	}

	query, err := h.convertQuery(message.String)
	if err != nil {
		return true, err
	}

	// A query message destroys the unnamed statement and the unnamed portal
	delete(h.preparedStatements, "")
	delete(h.portals, "")

	// Certain statement types get handled directly by the handler instead of being passed to the engine
	handled, endOfMessages, err = h.handleQueryOutsideEngine(query)
	if handled {
		return endOfMessages, err
	}

	return true, h.query(query)
}

// handleQueryOutsideEngine handles any queries that should be handled by the handler directly, rather than being
// passed to the engine. The response parameter |handled| is true if the query was handled, |endOfMessages| is true
// if no more messages are expected for this query and server should send the client a READY FOR QUERY message,
// and any error that occurred while handling the query.
func (h *ConnectionHandler) handleQueryOutsideEngine(query ConvertedQuery) (handled bool, endOfMessages bool, err error) {
	switch stmt := query.AST.(type) {
	case *sqlparser.Deallocate:
		// TODO: handle ALL keyword
		return true, true, h.deallocatePreparedStatement(stmt.Name, h.preparedStatements, query, h.Conn())
	case sqlparser.InjectedStatement:
		switch injectedStmt := stmt.Statement.(type) {
		case node.DiscardStatement:
			return true, true, h.discardAll(query, h.Conn())
		case *node.CopyFrom:
			// When copying data from STDIN, the data is sent to the server as CopyData messages
			// We send endOfMessages=false since the server will be in COPY DATA mode and won't
			// be ready for more queries util COPY DATA mode is completed.
			if injectedStmt.Stdin {
				return true, false, h.handleCopyFromStdinQuery(injectedStmt, h.Conn())
			}
		}
	}
	return false, true, nil
}

// handleParse handles a parse message, returning any error that occurs
func (h *ConnectionHandler) handleParse(message messages.Parse) error {
	h.waitForSync = true

	// TODO: "Named prepared statements must be explicitly closed before they can be redefined by another Parse message, but this is not required for the unnamed statement"
	query, err := h.convertQuery(message.Query)
	if err != nil {
		return err
	}

	if query.AST == nil {
		// special case: empty query
		h.preparedStatements[message.Name] = PreparedStatementData{
			Query: query,
		}
		return nil
	}

	parsedQuery, fields, err := h.doltgresHandler.ComPrepareParsed(context.Background(), h.mysqlConn, query.String, query.AST)
	if err != nil {
		return err
	}

	analyzedPlan, ok := parsedQuery.(sql.Node)
	if !ok {
		return fmt.Errorf("expected a sql.Node, got %T", parsedQuery)
	}

	// A valid Parse message must have ParameterObjectIDs if there are any binding variables.
	bindVarTypes := message.ParameterObjectIDs
	if len(bindVarTypes) == 0 {
		// NOTE: This is used for Prepared Statement Tests only.
		bindVarTypes, err = extractBindVarTypes(analyzedPlan)
		if err != nil {
			return err
		}
	}

	h.preparedStatements[message.Name] = PreparedStatementData{
		Query:        query,
		ReturnFields: fields,
		BindVarTypes: bindVarTypes,
	}

	return connection.Send(h.Conn(), messages.ParseComplete{})
}

// handleDescribe handles a Describe message, returning any error that occurs
func (h *ConnectionHandler) handleDescribe(message messages.Describe) error {
	var fields []pgproto3.FieldDescription
	var bindvarTypes []uint32
	var tag string

	h.waitForSync = true
	if message.IsPrepared {
		preparedStatementData, ok := h.preparedStatements[message.Target]
		if !ok {
			return fmt.Errorf("prepared statement %s does not exist", message.Target)
		}

		fields = preparedStatementData.ReturnFields
		bindvarTypes = preparedStatementData.BindVarTypes
		tag = preparedStatementData.Query.StatementTag
	} else {
		portalData, ok := h.portals[message.Target]
		if !ok {
			return fmt.Errorf("portal %s does not exist", message.Target)
		}

		fields = portalData.Fields
		tag = portalData.Query.StatementTag
	}

	return h.sendDescribeResponse(h.Conn(), fields, bindvarTypes, tag)
}

// handleBind handles a bind message, returning any error that occurs
func (h *ConnectionHandler) handleBind(message messages.Bind) error {
	h.waitForSync = true

	// TODO: a named portal object lasts till the end of the current transaction, unless explicitly destroyed
	//  we need to destroy the named portal as a side effect of the transaction ending
	logrus.Tracef("binding portal %q to prepared statement %s", message.DestinationPortal, message.SourcePreparedStatement)
	preparedData, ok := h.preparedStatements[message.SourcePreparedStatement]
	if !ok {
		return fmt.Errorf("prepared statement %s does not exist", message.SourcePreparedStatement)
	}

	if preparedData.Query.AST == nil {
		// special case: empty query
		h.portals[message.DestinationPortal] = PortalData{
			Query:        preparedData.Query,
			IsEmptyQuery: true,
		}
		return connection.Send(h.Conn(), messages.BindComplete{})
	}

	bindVars, err := h.convertBindParameters(preparedData.BindVarTypes, message.ParameterFormatCodes, message.ParameterValues)
	if err != nil {
		return err
	}

	analyzedPlan, fields, err := h.doltgresHandler.ComBind(context.Background(), h.mysqlConn, preparedData.Query.String, preparedData.Query.AST, bindVars)
	if err != nil {
		return err
	}

	boundPlan, ok := analyzedPlan.(sql.Node)
	if !ok {
		return fmt.Errorf("expected a sql.Node, got %T", analyzedPlan)
	}

	h.portals[message.DestinationPortal] = PortalData{
		Query:     preparedData.Query,
		Fields:    fields,
		BoundPlan: boundPlan,
	}
	return connection.Send(h.Conn(), messages.BindComplete{})
}

// handleExecute handles an execute message, returning any error that occurs
func (h *ConnectionHandler) handleExecute(message messages.Execute) error {
	h.waitForSync = true

	// TODO: implement the RowMax
	portalData, ok := h.portals[message.Portal]
	if !ok {
		return fmt.Errorf("portal %s does not exist", message.Portal)
	}

	logrus.Tracef("executing portal %s with contents %v", message.Portal, portalData)
	query := portalData.Query

	// we need the CommandComplete message defined here because it's altered by the callback below
	complete := messages.CommandComplete{
		Query: query.String,
		Tag:   query.StatementTag,
	}

	if portalData.IsEmptyQuery {
		return connection.Send(h.Conn(), messages.EmptyQueryResponse{})
	}

	// Certain statement types get handled directly by the handler instead of being passed to the engine
	handled, _, err := h.handleQueryOutsideEngine(query)
	if handled {
		return err
	}

	err = h.doltgresHandler.ComExecuteBound(context.Background(), h.mysqlConn, query.String, portalData.BoundPlan, spoolRowsCallback(h.Conn(), &complete, true))
	if err != nil {
		return err
	}

	return connection.Send(h.Conn(), complete)
}

// handleCopyData handles the COPY DATA message, by loading the data sent from the client. The |stop| response parameter
// is true if the connection handler should shut down the connection, |endOfMessages| is true if no more COPY DATA
// messages are expected, and the server should tell the client that it is ready for the next query, and |err| contains
// any error that occurred while processing the COPY DATA message.
func (h *ConnectionHandler) handleCopyData(message messages.CopyData) (stop bool, endOfMessages bool, err error) {
	if h.copyFromStdinState == nil {
		return false, true,
			fmt.Errorf("COPY DATA message received without a COPY FROM STDIN operation in progress")
	}

	// Grab a sql.Context
	ctxProvider, ok := h.mysqlHandler.(sql.ContextProvider)
	if !ok {
		return false, true, fmt.Errorf("%T does not implement server.ContextProvider", h.mysqlHandler)
	}
	sqlCtx, err := ctxProvider.NewContext(context.Background(), h.mysqlConn, "")
	if err != nil {
		return false, false, err
	}

	dataLoader := h.copyFromStdinState.dataLoader
	if dataLoader == nil {
		copyFromStdinNode := h.copyFromStdinState.copyFromStdinNode
		if copyFromStdinNode == nil {
			return false, false, fmt.Errorf("no COPY FROM STDIN node found")
		}

		// TODO: It would be better to get the table from the copyFromStdinNode – not by calling core.GetSqlTableFromContext
		table, err := core.GetSqlTableFromContext(sqlCtx, copyFromStdinNode.DatabaseName, copyFromStdinNode.TableName)
		if err != nil {
			return false, true, err
		}
		if table == nil {
			return false, true, fmt.Errorf(`relation "%s" does not exist`, copyFromStdinNode.TableName.String())
		}
		insertableTable, ok := table.(sql.InsertableTable)
		if !ok {
			return false, true, fmt.Errorf(`table "%s" is read-only`, copyFromStdinNode.TableName.String())
		}

		dataLoader, err = dataloader.NewTabularDataLoader(sqlCtx, insertableTable, copyFromStdinNode.Delimiter, copyFromStdinNode.Null)
		if err != nil {
			return false, false, err
		}

		h.copyFromStdinState.dataLoader = dataLoader
	}

	byteReader := bytes.NewReader(message.Data)
	reader := bufio.NewReader(byteReader)
	if err = dataLoader.LoadChunk(sqlCtx, reader); err != nil {
		return false, false, err
	}

	// We expect to see more CopyData messages until we see either a CopyDone or CopyFail message, so
	// return false for endOfMessages
	return false, false, nil
}

// handleCopyDone handles a COPY DONE message by finalizing the in-progress COPY DATA operation and committing the
// loaded table data. The |stop| response parameter is true if the connection handler should shut down the connection,
// |endOfMessages| is true if no more COPY DATA messages are expected, and the server should tell the client that it is
// ready for the next query, and |err| contains any error that occurred while processing the COPY DATA message.
func (h *ConnectionHandler) handleCopyDone(_ messages.CopyDone) (stop bool, endOfMessages bool, err error) {
	if h.copyFromStdinState == nil {
		return false, true,
			fmt.Errorf("COPY DONE message received without a COPY FROM STDIN operation in progress")
	}

	dataLoader := h.copyFromStdinState.dataLoader
	if dataLoader == nil {
		return false, true,
			fmt.Errorf("no data loader found for COPY FROM STDIN operation")
	}

	ctxProvider, ok := h.mysqlHandler.(sql.ContextProvider)
	if !ok {
		return false, true, fmt.Errorf("%T does not implement server.ContextProvider", h.mysqlHandler)
	}
	sqlCtx, err := ctxProvider.NewContext(context.Background(), h.mysqlConn, "")
	if err != nil {
		return false, false, err
	}

	loadDataResults, err := dataLoader.Finish(sqlCtx)
	if err != nil {
		return false, false, err
	}

	h.copyFromStdinState = nil
	// We send back endOfMessage=true, since the COPY DONE message ends the COPY DATA flow and the server is ready
	// to accept the next query now.
	return false, true, connection.Send(h.Conn(), messages.CommandComplete{
		Query: "",
		Rows:  loadDataResults.RowsLoaded,
		Tag:   "COPY",
	})
}

// handleCopyDone handles a COPY FAIL message by aborting the in-progress COPY DATA operation.  The |stop| response
// parameter is true if the connection handler should shut down the connection, |endOfMessages| is true if no more
// COPY DATA messages are expected, and the server should tell the client that it is ready for the next query, and
// |err| contains any error that occurred while processing the COPY DATA message.
func (h *ConnectionHandler) handleCopyFail(_ messages.CopyFail) (stop bool, endOfMessages bool, err error) {
	if h.copyFromStdinState == nil {
		return false, true,
			fmt.Errorf("COPY FAIL message received without a COPY FROM STDIN operation in progress")
	}

	dataLoader := h.copyFromStdinState.dataLoader
	if dataLoader == nil {
		return false, true,
			fmt.Errorf("no data loader found for COPY FROM STDIN operation")
	}

	h.copyFromStdinState = nil
	// We send back endOfMessage=true, since the COPY FAIL message ends the COPY DATA flow and the server is ready
	// to accept the next query now.
	return false, true, nil
}

func (h *ConnectionHandler) deallocatePreparedStatement(name string, preparedStatements map[string]PreparedStatementData, query ConvertedQuery, conn net.Conn) error {
	_, ok := preparedStatements[name]
	if !ok {
		return fmt.Errorf("prepared statement %s does not exist", name)
	}
	delete(preparedStatements, name)

	commandComplete := messages.CommandComplete{
		Query: query.String,
		Tag:   query.StatementTag,
	}

	return connection.Send(conn, commandComplete)
}

func extractBindVarTypes(queryPlan sql.Node) ([]uint32, error) {
	inspectNode := queryPlan
	switch queryPlan := queryPlan.(type) {
	case *plan.InsertInto:
		inspectNode = queryPlan.Source
	}

	types := make([]uint32, 0)
	var err error
	extractBindVars := func(expr sql.Expression) bool {
		if err != nil {
			return false
		}
		switch e := expr.(type) {
		case *expression.BindVar:
			var oid uint32
			if doltgresType, ok := e.Type().(pgtypes.DoltgresType); ok {
				oid = doltgresType.OID()
			} else {
				// TODO: error here?
				oid, err = messages.VitessTypeToObjectID(e.Type().Type())
				if err != nil {
					err = fmt.Errorf("could not determine OID for placeholder %s: %w", e.Name, err)
					return false
				}
			}
			types = append(types, oid)
		case *pgexprs.ExplicitCast:
			if bindVar, ok := e.Child().(*expression.BindVar); ok {
				var oid uint32
				if doltgresType, ok := bindVar.Type().(pgtypes.DoltgresType); ok {
					oid = doltgresType.OID()
				} else {
					oid, err = messages.VitessTypeToObjectID(e.Type().Type())
					if err != nil {
						err = fmt.Errorf("could not determine OID for placeholder %s: %w", bindVar.Name, err)
						return false
					}
				}
				types = append(types, oid)
				return false
			}
		// $1::text and similar get converted to a Convert expression wrapping the bindvar
		case *expression.Convert:
			if bindVar, ok := e.Child.(*expression.BindVar); ok {
				var oid uint32
				oid, err = messages.VitessTypeToObjectID(e.Type().Type())
				if err != nil {
					err = fmt.Errorf("could not determine OID for placeholder %s: %w", bindVar.Name, err)
					return false
				}
				types = append(types, oid)
				return false
			}
		}

		return true
	}

	transform.InspectExpressions(inspectNode, extractBindVars)
	return types, err
}

// convertBindParameters handles the conversion from bind parameters to variable values.
func (h *ConnectionHandler) convertBindParameters(types []uint32, formatCodes []int32, values []messages.BindParameterValue) (map[string]sqlparser.Expr, error) {
	bindings := make(map[string]sqlparser.Expr, len(values))
	for i := range values {
		typ := types[i]
		var bindVarString string
		// We'll rely on a library to decode each format, which will deal with text and binary representations for us
		if err := h.pgTypeMap.Scan(typ, int16(formatCodes[i]), values[i].Data, &bindVarString); err != nil {
			return nil, err
		}

		pgTyp, ok := OidToDoltgresType[typ]
		if !ok {
			return nil, fmt.Errorf("unhandled oid type: %v", typ)
		}
		v, err := pgTyp.IoInput(nil, bindVarString)
		if err != nil {
			return nil, err
		}
		bindings[fmt.Sprintf("v%d", i+1)] = sqlparser.InjectedExpr{Expression: pgexprs.NewUnsafeLiteral(v, pgTyp)}
	}
	return bindings, nil
}

// sendClientStartupMessages sends introductory messages to the client and returns any error
// TODO: implement users and authentication
func (h *ConnectionHandler) sendClientStartupMessages(startupMessage messages.StartupMessage) error {
	if user, ok := startupMessage.Parameters["user"]; ok && len(user) > 0 {
		var host string
		if h.Conn().RemoteAddr().Network() == "unix" {
			host = "localhost"
		} else {
			host, _, _ = net.SplitHostPort(h.Conn().RemoteAddr().String())
			if len(host) == 0 {
				host = "localhost"
			}
		}

		h.mysqlConn.User = user
		h.mysqlConn.UserData = sql.MysqlConnectionUser{
			User: user,
			Host: host,
		}
	} else {
		h.mysqlConn.User = "doltgres"
		h.mysqlConn.UserData = sql.MysqlConnectionUser{
			User: "doltgres",
			Host: "localhost",
		}
	}

	if err := connection.Send(h.Conn(), messages.AuthenticationOk{}); err != nil {
		return err
	}

	if err := connection.Send(h.Conn(), messages.ParameterStatus{
		Name:  "server_version",
		Value: "15.0",
	}); err != nil {
		return err
	}

	if err := connection.Send(h.Conn(), messages.ParameterStatus{
		Name:  "client_encoding",
		Value: "UTF8",
	}); err != nil {
		return err
	}

	if err := connection.Send(h.Conn(), messages.BackendKeyData{
		ProcessID: int32(processID),
		SecretKey: 0,
	}); err != nil {
		return err
	}

	return nil
}

// query runs the given query and sends a CommandComplete message to the client
func (h *ConnectionHandler) query(query ConvertedQuery) error {
	commandComplete := messages.CommandComplete{
		Query: query.String,
		Tag:   query.StatementTag,
	}

	err := h.doltgresHandler.ComQuery(context.Background(), h.mysqlConn, query.String, query.AST, spoolRowsCallback(h.Conn(), &commandComplete, false))
	if err != nil {
		if strings.HasPrefix(err.Error(), "syntax error at position") {
			return fmt.Errorf("This statement is not yet supported")
		}
		return err
	}

	if err := connection.Send(h.Conn(), commandComplete); err != nil {
		return err
	}

	return nil
}

// spoolRowsCallback returns a callback function that will send RowDescription message, then a DataRow message for
// each row in the result set.
func spoolRowsCallback(conn net.Conn, commandComplete *messages.CommandComplete, isExecute bool) func(res *Result) error {
	return func(res *Result) error {
		if messages.ReturnsRow(commandComplete.Tag) {
			// EXECUTE does not send RowDescription; instead it should be sent from DESCRIBE prior to it
			if !isExecute {
				if err := connection.Send(conn, messages.RowDescription{
					Fields: res.Fields,
				}); err != nil {
					return err
				}
			}

			for _, row := range res.Rows {
				if err := connection.Send(conn, messages.DataRow{
					Values: row.val,
				}); err != nil {
					return err
				}
			}
		}

		if commandComplete.IsIUD() {
			commandComplete.Rows = int32(res.RowsAffected)
		} else {
			commandComplete.Rows += int32(len(res.Rows))
		}
		return nil
	}
}

// sendDescribeResponse sends a response message for a Describe message
func (h *ConnectionHandler) sendDescribeResponse(conn net.Conn, fields []pgproto3.FieldDescription, types []uint32, tag string) (err error) {
	// The prepared statement variant of the describe command returns the OIDs of the parameters.
	if types != nil {
		if err := connection.Send(conn, messages.ParameterDescription{
			ObjectIDs: types,
		}); err != nil {
			return err
		}
	}

	if messages.ReturnsRow(tag) {
		// Both variants finish with a row description.
		return connection.Send(conn, messages.RowDescription{
			Fields: fields,
		})
	} else {
		return connection.Send(conn, messages.NoData{})
	}
}

// handledPSQLCommands handles the special PSQL commands, such as \l and \dt.
func (h *ConnectionHandler) handledPSQLCommands(statement string) (bool, error) {
	statement = strings.ToLower(statement)
	// Command: \l
	if statement == "select d.datname as \"name\",\n       pg_catalog.pg_get_userbyid(d.datdba) as \"owner\",\n       pg_catalog.pg_encoding_to_char(d.encoding) as \"encoding\",\n       d.datcollate as \"collate\",\n       d.datctype as \"ctype\",\n       d.daticulocale as \"icu locale\",\n       case d.datlocprovider when 'c' then 'libc' when 'i' then 'icu' end as \"locale provider\",\n       pg_catalog.array_to_string(d.datacl, e'\\n') as \"access privileges\"\nfrom pg_catalog.pg_database d\norder by 1;" {
		query, err := h.convertQuery(`select d.datname as "Name", 'postgres' as "Owner", 'UTF8' as "Encoding", 'en_US.UTF-8' as "Collate", 'en_US.UTF-8' as "Ctype", 'en-US' as "ICU Locale", case d.datlocprovider when 'c' then 'libc' when 'i' then 'icu' end as "locale provider", '' as "access privileges" from pg_catalog.pg_database d order by 1;`)
		if err != nil {
			return false, err
		}
		return true, h.query(query)
	}
	// Command: \l on psql 16
	if statement == "select\n  d.datname as \"name\",\n  pg_catalog.pg_get_userbyid(d.datdba) as \"owner\",\n  pg_catalog.pg_encoding_to_char(d.encoding) as \"encoding\",\n  case d.datlocprovider when 'c' then 'libc' when 'i' then 'icu' end as \"locale provider\",\n  d.datcollate as \"collate\",\n  d.datctype as \"ctype\",\n  d.daticulocale as \"icu locale\",\n  null as \"icu rules\",\n  pg_catalog.array_to_string(d.datacl, e'\\n') as \"access privileges\"\nfrom pg_catalog.pg_database d\norder by 1;" {
		query, err := h.convertQuery(`select d.datname as "Name", 'postgres' as "Owner", 'UTF8' as "Encoding", 'en_US.UTF-8' as "Collate", 'en_US.UTF-8' as "Ctype", 'en-US' as "ICU Locale", case d.datlocprovider when 'c' then 'libc' when 'i' then 'icu' end as "locale provider", '' as "access privileges" from pg_catalog.pg_database d order by 1;`)
		if err != nil {
			return false, err
		}
		return true, h.query(query)
	}
	// Command: \dt
	if statement == "select n.nspname as \"schema\",\n  c.relname as \"name\",\n  case c.relkind when 'r' then 'table' when 'v' then 'view' when 'm' then 'materialized view' when 'i' then 'index' when 's' then 'sequence' when 't' then 'toast table' when 'f' then 'foreign table' when 'p' then 'partitioned table' when 'i' then 'partitioned index' end as \"type\",\n  pg_catalog.pg_get_userbyid(c.relowner) as \"owner\"\nfrom pg_catalog.pg_class c\n     left join pg_catalog.pg_namespace n on n.oid = c.relnamespace\n     left join pg_catalog.pg_am am on am.oid = c.relam\nwhere c.relkind in ('r','p','')\n      and n.nspname <> 'pg_catalog'\n      and n.nspname !~ '^pg_toast'\n      and n.nspname <> 'information_schema'\n  and pg_catalog.pg_table_is_visible(c.oid)\norder by 1,2;" {
		return true, h.query(ConvertedQuery{
			String:       `SELECT table_schema AS 'Schema', TABLE_NAME AS 'Name', 'table' AS 'Type', 'postgres' AS 'Owner' FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA <> 'pg_catalog' AND CONVERT(TABLE_TYPE, CHAR) = 'BASE TABLE' ORDER BY 2;`,
			StatementTag: "SELECT",
		})
	}
	// Command: \d
	if statement == "select n.nspname as \"schema\",\n  c.relname as \"name\",\n  case c.relkind when 'r' then 'table' when 'v' then 'view' when 'm' then 'materialized view' when 'i' then 'index' when 's' then 'sequence' when 't' then 'toast table' when 'f' then 'foreign table' when 'p' then 'partitioned table' when 'i' then 'partitioned index' end as \"type\",\n  pg_catalog.pg_get_userbyid(c.relowner) as \"owner\"\nfrom pg_catalog.pg_class c\n     left join pg_catalog.pg_namespace n on n.oid = c.relnamespace\n     left join pg_catalog.pg_am am on am.oid = c.relam\nwhere c.relkind in ('r','p','v','m','s','f','')\n      and n.nspname <> 'pg_catalog'\n      and n.nspname !~ '^pg_toast'\n      and n.nspname <> 'information_schema'\n  and pg_catalog.pg_table_is_visible(c.oid)\norder by 1,2;" {
		return true, h.query(ConvertedQuery{
			String:       `SELECT table_schema AS 'Schema', TABLE_NAME AS 'Name', IF(TABLE_TYPE = 'VIEW', 'view', 'table') AS 'Type', 'postgres' AS 'Owner' FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA <> 'pg_catalog' AND CONVERT(TABLE_TYPE, CHAR) = 'BASE TABLE' OR CONVERT(TABLE_TYPE, CHAR) = 'VIEW' ORDER BY 2;`,
			StatementTag: "SELECT",
		})
	}
	// Alternate \d for psql 14
	if statement == "select n.nspname as \"schema\",\n  c.relname as \"name\",\n  case c.relkind when 'r' then 'table' when 'v' then 'view' when 'm' then 'materialized view' when 'i' then 'index' when 's' then 'sequence' when 's' then 'special' when 't' then 'toast table' when 'f' then 'foreign table' when 'p' then 'partitioned table' when 'i' then 'partitioned index' end as \"type\",\n  pg_catalog.pg_get_userbyid(c.relowner) as \"owner\"\nfrom pg_catalog.pg_class c\n     left join pg_catalog.pg_namespace n on n.oid = c.relnamespace\n     left join pg_catalog.pg_am am on am.oid = c.relam\nwhere c.relkind in ('r','p','v','m','s','f','')\n      and n.nspname <> 'pg_catalog'\n      and n.nspname !~ '^pg_toast'\n      and n.nspname <> 'information_schema'\n  and pg_catalog.pg_table_is_visible(c.oid)\norder by 1,2;" {
		return true, h.query(ConvertedQuery{
			String:       `SELECT table_schema AS 'Schema', TABLE_NAME AS 'Name', IF(TABLE_TYPE = 'VIEW', 'view', 'table') AS 'Type', 'postgres' AS 'Owner' FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA <> 'pg_catalog' AND CONVERT(TABLE_TYPE, CHAR) = 'BASE TABLE' OR CONVERT(TABLE_TYPE, CHAR) = 'VIEW' ORDER BY 2;`,
			StatementTag: "SELECT",
		})
	}
	// Command: \d table_name
	if strings.HasPrefix(statement, "select c.oid,\n  n.nspname,\n  c.relname\nfrom pg_catalog.pg_class c\n     left join pg_catalog.pg_namespace n on n.oid = c.relnamespace\nwhere c.relname operator(pg_catalog.~) '^(") && strings.HasSuffix(statement, ")$' collate pg_catalog.default\n  and pg_catalog.pg_table_is_visible(c.oid)\norder by 2, 3;") {
		// There are >at least< 15 separate statements sent for this command, which is far too much to validate and
		// implement, so we'll just return an error for now
		return true, fmt.Errorf("PSQL command not yet supported")
	}
	// Command: \dn
	if statement == "select n.nspname as \"name\",\n  pg_catalog.pg_get_userbyid(n.nspowner) as \"owner\"\nfrom pg_catalog.pg_namespace n\nwhere n.nspname !~ '^pg_' and n.nspname <> 'information_schema'\norder by 1;" {
		return true, h.query(ConvertedQuery{
			String:       `SELECT 'public' AS 'Name', 'pg_database_owner' AS 'Owner';`,
			StatementTag: "SELECT",
		})
	}
	// Command: \df
	if statement == "select n.nspname as \"schema\",\n  p.proname as \"name\",\n  pg_catalog.pg_get_function_result(p.oid) as \"result data type\",\n  pg_catalog.pg_get_function_arguments(p.oid) as \"argument data types\",\n case p.prokind\n  when 'a' then 'agg'\n  when 'w' then 'window'\n  when 'p' then 'proc'\n  else 'func'\n end as \"type\"\nfrom pg_catalog.pg_proc p\n     left join pg_catalog.pg_namespace n on n.oid = p.pronamespace\nwhere pg_catalog.pg_function_is_visible(p.oid)\n      and n.nspname <> 'pg_catalog'\n      and n.nspname <> 'information_schema'\norder by 1, 2, 4;" {
		return true, h.query(ConvertedQuery{
			String:       `SELECT '' AS 'Schema', '' AS 'Name', '' AS 'Result data type', '' AS 'Argument data types', '' AS 'Type' FROM dual LIMIT 0;`,
			StatementTag: "SELECT",
		})
	}
	// Command: \dv
	if statement == "select n.nspname as \"schema\",\n  c.relname as \"name\",\n  case c.relkind when 'r' then 'table' when 'v' then 'view' when 'm' then 'materialized view' when 'i' then 'index' when 's' then 'sequence' when 't' then 'toast table' when 'f' then 'foreign table' when 'p' then 'partitioned table' when 'i' then 'partitioned index' end as \"type\",\n  pg_catalog.pg_get_userbyid(c.relowner) as \"owner\"\nfrom pg_catalog.pg_class c\n     left join pg_catalog.pg_namespace n on n.oid = c.relnamespace\nwhere c.relkind in ('v','')\n      and n.nspname <> 'pg_catalog'\n      and n.nspname !~ '^pg_toast'\n      and n.nspname <> 'information_schema'\n  and pg_catalog.pg_table_is_visible(c.oid)\norder by 1,2;" {
		return true, h.query(ConvertedQuery{
			String:       `SELECT table_schema AS 'Schema', TABLE_NAME AS 'Name', 'view' AS 'Type', 'postgres' AS 'Owner' FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA <> 'pg_catalog' AND TABLE_TYPE = 'VIEW' ORDER BY 2;`,
			StatementTag: "SELECT",
		})
	}
	// Command: \du
	if statement == "select r.rolname, r.rolsuper, r.rolinherit,\n  r.rolcreaterole, r.rolcreatedb, r.rolcanlogin,\n  r.rolconnlimit, r.rolvaliduntil,\n  array(select b.rolname\n        from pg_catalog.pg_auth_members m\n        join pg_catalog.pg_roles b on (m.roleid = b.oid)\n        where m.member = r.oid) as memberof\n, r.rolreplication\n, r.rolbypassrls\nfrom pg_catalog.pg_roles r\nwhere r.rolname !~ '^pg_'\norder by 1;" {
		// We don't support users yet, so we'll just return nothing for now
		return true, h.query(ConvertedQuery{
			String:       `SELECT '' FROM dual LIMIT 0;`,
			StatementTag: "SELECT",
		})
	}
	return false, nil
}

// handledWorkbenchCommands handles commands used by some workbenches, such as dolt-workbench.
func (h *ConnectionHandler) handledWorkbenchCommands(statement string) (bool, error) {
	lower := strings.ToLower(statement)
	if lower == "select * from current_schema()" || lower == "select * from current_schema();" {
		return true, h.query(ConvertedQuery{
			String:       `SELECT search_path AS 'current_schema';`,
			StatementTag: "SELECT",
		})
	}
	if lower == "select * from current_database()" || lower == "select * from current_database();" {
		return true, h.query(ConvertedQuery{
			String:       `SELECT DATABASE() AS 'current_database';`,
			StatementTag: "SELECT",
		})
	}
	return false, nil
}

// endOfMessages should be called from HandleConnection or a function within HandleConnection. This represents the end
// of the message slice, which may occur naturally (all relevant response messages have been sent) or on error. Once
// endOfMessages has been called, no further messages should be sent, and the connection loop should wait for the next
// query. A nil error should be provided if this is being called naturally.
func (h *ConnectionHandler) endOfMessages(err error) {
	if err != nil {
		h.sendError(h.Conn(), err)
	}
	if sendErr := connection.Send(h.Conn(), messages.ReadyForQuery{
		Indicator: messages.ReadyForQueryTransactionIndicator_Idle,
	}); sendErr != nil {
		// We panic here for the same reason as above.
		panic(sendErr)
	}
}

// sendError sends the given error to the client. This should generally never be called directly.
func (h *ConnectionHandler) sendError(conn net.Conn, err error) {
	fmt.Println(err.Error())
	if sendErr := connection.Send(conn, messages.ErrorResponse{
		Severity:     messages.ErrorResponseSeverity_Error,
		SqlStateCode: "XX000", // internal_error for now
		Message:      err.Error(),
	}); sendErr != nil {
		// If we're unable to send anything to the connection, then there's something wrong with the connection and
		// we should terminate it. This will be caught in HandleConnection's defer block.
		panic(sendErr)
	}
}

// convertQuery takes the given Postgres query, and converts it as an ast.ConvertedQuery that will work with the handler.
func (h *ConnectionHandler) convertQuery(query string) (ConvertedQuery, error) {
	s, err := parser.Parse(query)
	if err != nil {
		return ConvertedQuery{}, err
	}
	if len(s) > 1 {
		return ConvertedQuery{}, fmt.Errorf("only a single statement at a time is currently supported")
	}
	if len(s) == 0 {
		return ConvertedQuery{String: query}, nil
	}
	vitessAST, err := ast.Convert(s[0])
	stmtTag := s[0].AST.StatementTag()
	if err != nil {
		return ConvertedQuery{}, err
	}
	if vitessAST == nil {
		return ConvertedQuery{
			String:       s[0].AST.String(),
			StatementTag: stmtTag,
		}, nil
	}
	return ConvertedQuery{
		String:       query,
		AST:          vitessAST,
		StatementTag: stmtTag,
	}, nil
}

// discardAll handles the DISCARD ALL command
func (h *ConnectionHandler) discardAll(query ConvertedQuery, conn net.Conn) error {
	err := h.mysqlHandler.ComResetConnection(h.mysqlConn)
	if err != nil {
		return err
	}

	commandComplete := messages.CommandComplete{
		Query: query.String,
		Tag:   query.StatementTag,
	}

	return connection.Send(conn, commandComplete)
}

// handleCopyFromStdinQuery handles the COPY FROM STDIN query at the Doltgres layer, without passing it to the engine.
// COPY FROM STDIN can't be handled directly by the GMS engine, since COPY FROM STDIN relies on multiple messages sent
// over the wire.
func (h *ConnectionHandler) handleCopyFromStdinQuery(copyFrom *node.CopyFrom, conn net.Conn) error {
	ctxProvider, ok := h.mysqlHandler.(sql.ContextProvider)
	if !ok {
		return fmt.Errorf("%T does not implement server.ContextProvider", h.mysqlHandler)
	}
	sqlCtx, err := ctxProvider.NewContext(context.Background(), h.mysqlConn, "")
	if err != nil {
		return err
	}

	if err := copyFrom.Validate(sqlCtx); err != nil {
		return err
	}

	h.copyFromStdinState = &copyFromStdinState{
		copyFromStdinNode: copyFrom,
	}

	return connection.Send(conn, messages.CopyInResponse{
		IsTextual: true,
	})
}

var OidToDoltgresType = map[uint32]pgtypes.DoltgresType{
	uint32(oid.T_bool):             pgtypes.Bool,
	uint32(oid.T_bytea):            pgtypes.Bytea,
	uint32(oid.T_char):             pgtypes.InternalChar,
	uint32(oid.T_name):             pgtypes.Name,
	uint32(oid.T_int8):             pgtypes.Int64,
	uint32(oid.T_int2):             pgtypes.Int16,
	uint32(oid.T_int2vector):       pgtypes.Unknown,
	uint32(oid.T_int4):             pgtypes.Int32,
	uint32(oid.T_regproc):          pgtypes.Regproc,
	uint32(oid.T_text):             pgtypes.Text,
	uint32(oid.T_oid):              pgtypes.Oid,
	uint32(oid.T_tid):              pgtypes.Unknown,
	uint32(oid.T_xid):              pgtypes.Xid,
	uint32(oid.T_cid):              pgtypes.Unknown,
	uint32(oid.T_oidvector):        pgtypes.Unknown,
	uint32(oid.T_pg_ddl_command):   pgtypes.Unknown,
	uint32(oid.T_pg_type):          pgtypes.Unknown,
	uint32(oid.T_pg_attribute):     pgtypes.Unknown,
	uint32(oid.T_pg_proc):          pgtypes.Unknown,
	uint32(oid.T_pg_class):         pgtypes.Unknown,
	uint32(oid.T_json):             pgtypes.Json,
	uint32(oid.T_xml):              pgtypes.Unknown,
	uint32(oid.T__xml):             pgtypes.Unknown,
	uint32(oid.T_pg_node_tree):     pgtypes.Unknown,
	uint32(oid.T__json):            pgtypes.JsonArray,
	uint32(oid.T_smgr):             pgtypes.Unknown,
	uint32(oid.T_index_am_handler): pgtypes.Unknown,
	uint32(oid.T_point):            pgtypes.Unknown,
	uint32(oid.T_lseg):             pgtypes.Unknown,
	uint32(oid.T_path):             pgtypes.Unknown,
	uint32(oid.T_box):              pgtypes.Unknown,
	uint32(oid.T_polygon):          pgtypes.Unknown,
	uint32(oid.T_line):             pgtypes.Unknown,
	uint32(oid.T__line):            pgtypes.Unknown,
	uint32(oid.T_cidr):             pgtypes.Unknown,
	uint32(oid.T__cidr):            pgtypes.Unknown,
	uint32(oid.T_float4):           pgtypes.Float32,
	uint32(oid.T_float8):           pgtypes.Float64,
	uint32(oid.T_abstime):          pgtypes.Unknown,
	uint32(oid.T_reltime):          pgtypes.Unknown,
	uint32(oid.T_tinterval):        pgtypes.Unknown,
	uint32(oid.T_unknown):          pgtypes.Unknown,
	uint32(oid.T_circle):           pgtypes.Unknown,
	uint32(oid.T__circle):          pgtypes.Unknown,
	uint32(oid.T_money):            pgtypes.Unknown,
	uint32(oid.T__money):           pgtypes.Unknown,
	uint32(oid.T_macaddr):          pgtypes.Unknown,
	uint32(oid.T_inet):             pgtypes.Unknown,
	uint32(oid.T__bool):            pgtypes.BoolArray,
	uint32(oid.T__bytea):           pgtypes.ByteaArray,
	uint32(oid.T__char):            pgtypes.InternalCharArray,
	uint32(oid.T__name):            pgtypes.NameArray,
	uint32(oid.T__int2):            pgtypes.Int16Array,
	uint32(oid.T__int2vector):      pgtypes.Unknown,
	uint32(oid.T__int4):            pgtypes.Int32Array,
	uint32(oid.T__regproc):         pgtypes.RegprocArray,
	uint32(oid.T__text):            pgtypes.TextArray,
	uint32(oid.T__tid):             pgtypes.Unknown,
	uint32(oid.T__xid):             pgtypes.XidArray,
	uint32(oid.T__cid):             pgtypes.Unknown,
	uint32(oid.T__oidvector):       pgtypes.Unknown,
	uint32(oid.T__bpchar):          pgtypes.BpCharArray,
	uint32(oid.T__varchar):         pgtypes.VarCharArray,
	uint32(oid.T__int8):            pgtypes.Int64Array,
	uint32(oid.T__point):           pgtypes.Unknown,
	uint32(oid.T__lseg):            pgtypes.Unknown,
	uint32(oid.T__path):            pgtypes.Unknown,
	uint32(oid.T__box):             pgtypes.Unknown,
	uint32(oid.T__float4):          pgtypes.Float32Array,
	uint32(oid.T__float8):          pgtypes.Float64Array,
	uint32(oid.T__abstime):         pgtypes.Unknown,
	uint32(oid.T__reltime):         pgtypes.Unknown,
	uint32(oid.T__tinterval):       pgtypes.Unknown,
	uint32(oid.T__polygon):         pgtypes.Unknown,
	uint32(oid.T__oid):             pgtypes.OidArray,
	uint32(oid.T_aclitem):          pgtypes.Unknown,
	uint32(oid.T__aclitem):         pgtypes.Unknown,
	uint32(oid.T__macaddr):         pgtypes.Unknown,
	uint32(oid.T__inet):            pgtypes.Unknown,
	uint32(oid.T_bpchar):           pgtypes.BpChar,
	uint32(oid.T_varchar):          pgtypes.VarChar,
	uint32(oid.T_date):             pgtypes.Date,
	uint32(oid.T_time):             pgtypes.Time,
	uint32(oid.T_timestamp):        pgtypes.Timestamp,
	uint32(oid.T__timestamp):       pgtypes.TimestampArray,
	uint32(oid.T__date):            pgtypes.DateArray,
	uint32(oid.T__time):            pgtypes.TimeArray,
	uint32(oid.T_timestamptz):      pgtypes.TimestampTZ,
	uint32(oid.T__timestamptz):     pgtypes.TimestampTZArray,
	uint32(oid.T_interval):         pgtypes.Interval,
	uint32(oid.T__interval):        pgtypes.IntervalArray,
	uint32(oid.T__numeric):         pgtypes.NumericArray,
	uint32(oid.T_pg_database):      pgtypes.Unknown,
	uint32(oid.T__cstring):         pgtypes.Unknown,
	uint32(oid.T_timetz):           pgtypes.TimeTZ,
	uint32(oid.T__timetz):          pgtypes.TimeTZArray,
	uint32(oid.T_bit):              pgtypes.Unknown,
	uint32(oid.T__bit):             pgtypes.Unknown,
	uint32(oid.T_varbit):           pgtypes.Unknown,
	uint32(oid.T__varbit):          pgtypes.Unknown,
	uint32(oid.T_numeric):          pgtypes.Numeric,
	uint32(oid.T_refcursor):        pgtypes.Unknown,
	uint32(oid.T__refcursor):       pgtypes.Unknown,
	uint32(oid.T_regprocedure):     pgtypes.Unknown,
	uint32(oid.T_regoper):          pgtypes.Unknown,
	uint32(oid.T_regoperator):      pgtypes.Unknown,
	uint32(oid.T_regclass):         pgtypes.Regclass,
	uint32(oid.T_regtype):          pgtypes.Regtype,
	uint32(oid.T__regprocedure):    pgtypes.Unknown,
	uint32(oid.T__regoper):         pgtypes.Unknown,
	uint32(oid.T__regoperator):     pgtypes.Unknown,
	uint32(oid.T__regclass):        pgtypes.RegclassArray,
	uint32(oid.T__regtype):         pgtypes.RegtypeArray,
	uint32(oid.T_record):           pgtypes.Unknown,
	uint32(oid.T_cstring):          pgtypes.Unknown,
	uint32(oid.T_any):              pgtypes.Unknown,
	uint32(oid.T_anyarray):         pgtypes.AnyArray,
	uint32(oid.T_void):             pgtypes.Unknown,
	uint32(oid.T_trigger):          pgtypes.Unknown,
	uint32(oid.T_language_handler): pgtypes.Unknown,
	uint32(oid.T_internal):         pgtypes.Unknown,
	uint32(oid.T_opaque):           pgtypes.Unknown,
	uint32(oid.T_anyelement):       pgtypes.AnyElement,
	uint32(oid.T__record):          pgtypes.Unknown,
	uint32(oid.T_anynonarray):      pgtypes.AnyNonArray,
	uint32(oid.T_pg_authid):        pgtypes.Unknown,
	uint32(oid.T_pg_auth_members):  pgtypes.Unknown,
	uint32(oid.T__txid_snapshot):   pgtypes.Unknown,
	uint32(oid.T_uuid):             pgtypes.Uuid,
	uint32(oid.T__uuid):            pgtypes.UuidArray,
	uint32(oid.T_txid_snapshot):    pgtypes.Unknown,
	uint32(oid.T_fdw_handler):      pgtypes.Unknown,
	uint32(oid.T_pg_lsn):           pgtypes.Unknown,
	uint32(oid.T__pg_lsn):          pgtypes.Unknown,
	uint32(oid.T_tsm_handler):      pgtypes.Unknown,
	uint32(oid.T_anyenum):          pgtypes.Unknown,
	uint32(oid.T_tsvector):         pgtypes.Unknown,
	uint32(oid.T_tsquery):          pgtypes.Unknown,
	uint32(oid.T_gtsvector):        pgtypes.Unknown,
	uint32(oid.T__tsvector):        pgtypes.Unknown,
	uint32(oid.T__gtsvector):       pgtypes.Unknown,
	uint32(oid.T__tsquery):         pgtypes.Unknown,
	uint32(oid.T_regconfig):        pgtypes.Unknown,
	uint32(oid.T__regconfig):       pgtypes.Unknown,
	uint32(oid.T_regdictionary):    pgtypes.Unknown,
	uint32(oid.T__regdictionary):   pgtypes.Unknown,
	uint32(oid.T_jsonb):            pgtypes.JsonB,
	uint32(oid.T__jsonb):           pgtypes.JsonBArray,
	uint32(oid.T_anyrange):         pgtypes.Unknown,
	uint32(oid.T_event_trigger):    pgtypes.Unknown,
	uint32(oid.T_int4range):        pgtypes.Unknown,
	uint32(oid.T__int4range):       pgtypes.Unknown,
	uint32(oid.T_numrange):         pgtypes.Unknown,
	uint32(oid.T__numrange):        pgtypes.Unknown,
	uint32(oid.T_tsrange):          pgtypes.Unknown,
	uint32(oid.T__tsrange):         pgtypes.Unknown,
	uint32(oid.T_tstzrange):        pgtypes.Unknown,
	uint32(oid.T__tstzrange):       pgtypes.Unknown,
	uint32(oid.T_daterange):        pgtypes.Unknown,
	uint32(oid.T__daterange):       pgtypes.Unknown,
	uint32(oid.T_int8range):        pgtypes.Unknown,
	uint32(oid.T__int8range):       pgtypes.Unknown,
	uint32(oid.T_pg_shseclabel):    pgtypes.Unknown,
	uint32(oid.T_regnamespace):     pgtypes.Unknown,
	uint32(oid.T__regnamespace):    pgtypes.Unknown,
	uint32(oid.T_regrole):          pgtypes.Unknown,
	uint32(oid.T__regrole):         pgtypes.Unknown,
}
