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

// nodeCreateView handles *tree.CreateView nodes.
func nodeCreateView(node *tree.CreateView) (*vitess.DDL, error) {
	if node == nil {
		return nil, nil
	}
	//TODO: AFAICT, CREATE VIEW uses a substring of the original string, which won't necessarily work for us.
	// Although it takes a parsed definition, this definition isn't sent to integrators.
	return nil, fmt.Errorf("CREATE VIEW is not yet supported")
}
