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

package _go

import (
	"testing"

	"github.com/dolthub/go-mysql-server/sql"
)

func TestAlterTableAddPrimaryKey(t *testing.T) {
	RunScripts(t, []ScriptTest{
		{
			Name: "Add Primary Key",
			SetUpScript: []string{
				"CREATE TABLE test1 (a INT, b INT);",
				"CREATE TABLE test2 (a INT, b INT, c INT);",
				"CREATE TABLE pkTable1 (a INT PRIMARY KEY);",
			},
			Assertions: []ScriptTestAssertion{
				{
					Query:    "ALTER TABLE test1 ADD PRIMARY KEY (a);",
					Expected: []sql.Row{},
				},
				{
					// Test the pk by inserting a duplicate value
					Query:       "INSERT into test1 values (1, 2), (1, 3);",
					ExpectedErr: "duplicate primary key",
				},
				{
					Query:    "ALTER TABLE test2 ADD PRIMARY KEY (a, b);",
					Expected: []sql.Row{},
				},
				{
					// Test the pk by inserting a duplicate value
					Query:       "INSERT into test2 values (1, 2, 3), (1, 2, 4);",
					ExpectedErr: "duplicate primary key",
				},
				{
					Query:       "ALTER TABLE pkTable1 ADD PRIMARY KEY (a);",
					ExpectedErr: "Multiple primary keys defined",
				},
				{
					// TODO: This statement fails in analysis, because it can't find a table named
					//       doesNotExist – since IF EXISTS is specified, the analyzer should skip
					//       errors on resolving the table in this case.
					Skip:     true,
					Query:    "ALTER TABLE IF EXISTS doesNotExist ADD PRIMARY KEY (a, b);",
					Expected: []sql.Row{},
				},
			},
		},
	})
}
