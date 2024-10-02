// Copyright 2024 Dolthub, Inc.
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

package pgcatalog

import (
	"io"

	"github.com/dolthub/go-mysql-server/sql"

	"github.com/dolthub/doltgresql/server/tables"
	pgtypes "github.com/dolthub/doltgresql/server/types"
)

// PgOperatorName is a constant to the pg_operator name.
const PgOperatorName = "pg_operator"

// InitPgOperator handles registration of the pg_operator handler.
func InitPgOperator() {
	tables.AddHandler(PgCatalogName, PgOperatorName, PgOperatorHandler{})
}

// PgOperatorHandler is the handler for the pg_operator table.
type PgOperatorHandler struct{}

var _ tables.Handler = PgOperatorHandler{}

// Name implements the interface tables.Handler.
func (p PgOperatorHandler) Name() string {
	return PgOperatorName
}

// RowIter implements the interface tables.Handler.
func (p PgOperatorHandler) RowIter(ctx *sql.Context) (sql.RowIter, error) {
	// TODO: Implement pg_operator row iter
	return emptyRowIter()
}

// Schema implements the interface tables.Handler.
func (p PgOperatorHandler) Schema() sql.PrimaryKeySchema {
	return sql.PrimaryKeySchema{
		Schema:     pgOperatorSchema,
		PkOrdinals: nil,
	}
}

// pgOperatorSchema is the schema for pg_operator.
var pgOperatorSchema = sql.Schema{
	{Name: "oid", Type: pgtypes.Oid, Default: nil, Nullable: false, Source: PgOperatorName},
	{Name: "oprname", Type: pgtypes.Name, Default: nil, Nullable: false, Source: PgOperatorName},
	{Name: "oprnamespace", Type: pgtypes.Oid, Default: nil, Nullable: false, Source: PgOperatorName},
	{Name: "oprowner", Type: pgtypes.Oid, Default: nil, Nullable: false, Source: PgOperatorName},
	{Name: "oprkind", Type: pgtypes.InternalChar, Default: nil, Nullable: false, Source: PgOperatorName},
	{Name: "oprcanmerge", Type: pgtypes.Bool, Default: nil, Nullable: false, Source: PgOperatorName},
	{Name: "oprcanhash", Type: pgtypes.Bool, Default: nil, Nullable: false, Source: PgOperatorName},
	{Name: "oprleft", Type: pgtypes.Oid, Default: nil, Nullable: false, Source: PgOperatorName},
	{Name: "oprright", Type: pgtypes.Oid, Default: nil, Nullable: false, Source: PgOperatorName},
	{Name: "oprresult", Type: pgtypes.Oid, Default: nil, Nullable: false, Source: PgOperatorName},
	{Name: "oprcom", Type: pgtypes.Oid, Default: nil, Nullable: false, Source: PgOperatorName},
	{Name: "oprnegate", Type: pgtypes.Oid, Default: nil, Nullable: false, Source: PgOperatorName},
	{Name: "oprcode", Type: pgtypes.Regproc, Default: nil, Nullable: false, Source: PgOperatorName},
	{Name: "oprrest", Type: pgtypes.Regproc, Default: nil, Nullable: false, Source: PgOperatorName},
	{Name: "oprjoin", Type: pgtypes.Regproc, Default: nil, Nullable: false, Source: PgOperatorName},
}

// pgOperatorRowIter is the sql.RowIter for the pg_operator table.
type pgOperatorRowIter struct {
}

var _ sql.RowIter = (*pgOperatorRowIter)(nil)

// Next implements the interface sql.RowIter.
func (iter *pgOperatorRowIter) Next(ctx *sql.Context) (sql.Row, error) {
	return nil, io.EOF
}

// Close implements the interface sql.RowIter.
func (iter *pgOperatorRowIter) Close(ctx *sql.Context) error {
	return nil
}
