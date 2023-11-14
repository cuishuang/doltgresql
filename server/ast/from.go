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

	vitess "github.com/dolthub/vitess/go/vt/sqlparser"

	"github.com/dolthub/doltgresql/postgres/parser/sem/tree"
)

// nodeFrom handles *tree.From nodes.
func nodeFrom(node tree.From) (vitess.TableExprs, error) {
	if len(node.Tables) == 0 {
		return nil, nil
	}
	asOfExpr, err := nodeExpr(node.AsOf.Expr)
	if err != nil {
		return nil, err
	}
	tableExprs, err := nodeTableExprs(node.Tables)
	if err != nil {
		return nil, err
	}
	if asOfExpr != nil {
		//TODO: determine if this will always be Time
		asOf := &vitess.AsOf{
			Time: asOfExpr,
		}
		for _, tableExpr := range tableExprs {
			if aliasedTableExpr, ok := tableExpr.(*vitess.AliasedTableExpr); ok {
				aliasedTableExpr.AsOf = asOf
			} else {
				return nil, fmt.Errorf("this particular usage of AS OF is not yet supported")
			}
		}
	}
	return tableExprs, err
}
