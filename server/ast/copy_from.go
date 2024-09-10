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

package ast

import (
	"fmt"

	"github.com/dolthub/dolt/go/libraries/doltcore/doltdb"
	vitess "github.com/dolthub/vitess/go/vt/sqlparser"

	"github.com/dolthub/doltgresql/postgres/parser/sem/tree"
	pgnodes "github.com/dolthub/doltgresql/server/node"
)

// nodeCopyFrom handles *tree.CopyFrom nodes.
func nodeCopyFrom(node *tree.CopyFrom) (vitess.Statement, error) {
	if node == nil {
		return nil, nil
	}
	if len(node.Columns) != 0 {
		return nil, fmt.Errorf("COPY FROM does not yet support a column list")
	}
	if node.Stdin {
		return nil, fmt.Errorf("COPY FROM does not yet support reading from STDIN")
	}
	if node.Options.CopyFormat != tree.CopyFormatText {
		return nil, fmt.Errorf("COPY FROM does not yet support non-text formats")
	}
	if node.Options.Destination != nil {
		return nil, fmt.Errorf("COPY FROM does not work with destinations")
	}
	return vitess.InjectedStatement{
		Statement: pgnodes.NewCopyFrom(node.Table.Catalog(), doltdb.TableName{
			Name:   node.Table.Object(),
			Schema: node.Table.Schema(),
		}, node.File),
		Children: nil,
	}, nil
}
