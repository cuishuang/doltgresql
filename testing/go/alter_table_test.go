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

func TestAlterTableAddForeignKeyConstraint(t *testing.T) {
	RunScripts(t, []ScriptTest{
		{
			Name: "Add Foreign Key Constraint",
			SetUpScript: []string{
				"create table child (pk int primary key, c1 int);",
				"insert into child values (1,1), (2,2), (3,3);",
				"create index idx_child_c1 on child (pk, c1);",
				"create table parent (pk int primary key, c1 int, c2 int);",
				"insert into parent values (1, 1, 10);",
			},
			Assertions: []ScriptTestAssertion{
				{
					Query:    "ALTER TABLE parent ADD FOREIGN KEY (c1) REFERENCES child (pk) ON DELETE CASCADE;",
					Expected: []sql.Row{},
				},
				{
					// Test that the FK constraint is working
					Query:       "INSERT INTO parent VALUES (10, 10, 10);",
					ExpectedErr: "Foreign key violation",
				},
				{
					Query:       "ALTER TABLE parent ADD FOREIGN KEY (c2) REFERENCES child (pk);",
					ExpectedErr: "Foreign key violation",
				},
				{
					// Test an FK reference over multiple columns
					Query:       "ALTER TABLE parent ADD FOREIGN KEY (c1, c2) REFERENCES child (pk, c1);",
					ExpectedErr: "Foreign key violation",
				},
				{
					// Unsupported syntax: MATCH PARTIAL
					Query:       "ALTER TABLE parent ADD FOREIGN KEY (c1, c2) REFERENCES child (pk, c1) MATCH PARTIAL;",
					ExpectedErr: "MATCH PARTIAL is not yet supported",
				},
			},
		},
	})
}

func TestAlterTableAddPrimaryKey(t *testing.T) {
	RunScripts(t, []ScriptTest{
		{
			Name: "Add Primary Key",
			SetUpScript: []string{
				"CREATE TABLE test1 (a INT, b INT);",
				"CREATE TABLE test2 (a INT, b INT, c INT);",
				"CREATE TABLE pkTable1 (a INT PRIMARY KEY);",
				"CREATE TABLE duplicateRows (a INT, b INT);",
				"INSERT INTO duplicateRows VALUES (1, 2), (1, 2);",
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
					Query:       "ALTER TABLE duplicateRows ADD PRIMARY KEY (a);",
					ExpectedErr: "duplicate primary key",
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

func TestAlterTableAddColumn(t *testing.T) {
	RunScripts(t, []ScriptTest{
		{
			Name: "Add Column",
			SetUpScript: []string{
				"CREATE TABLE test1 (a INT, b INT);",
				"INSERT INTO test1 VALUES (1, 1);",
			},
			Assertions: []ScriptTestAssertion{
				{
					Query:    "ALTER TABLE test1 ADD COLUMN c INT NOT NULL DEFAULT 42;",
					Expected: []sql.Row{},
				},
				{
					Query:    "select * from test1;",
					Expected: []sql.Row{{1, 1, 42}},
				},
			},
		},
	})
}

func TestAlterTableDropColumn(t *testing.T) {
	RunScripts(t, []ScriptTest{
		{
			Name: "Drop Column",
			SetUpScript: []string{
				"CREATE TABLE test1 (a INT, b INT, c INT, d INT);",
				"INSERT INTO test1 VALUES (1, 2, 3, 4);",
			},
			Assertions: []ScriptTestAssertion{
				{
					Query:    "ALTER TABLE test1 DROP COLUMN c;",
					Expected: []sql.Row{},
				},
				{
					Query:    "select * from test1;",
					Expected: []sql.Row{{1, 2, 4}},
				},
				{
					Query:    "ALTER TABLE test1 DROP COLUMN d;",
					Expected: []sql.Row{},
				},
				{
					Query:    "select * from test1;",
					Expected: []sql.Row{{1, 2}},
				},
				{
					// TODO: Skipped until we support conditional execution on existence of column
					Skip:     true,
					Query:    "ALTER TABLE test1 DROP COLUMN IF EXISTS zzz;",
					Expected: []sql.Row{},
				},
				{
					// TODO: Even though we're setting IF EXISTS, this query still fails with an
					//       error about the table not existing.
					Skip:     true,
					Query:    "ALTER TABLE IF EXISTS doesNotExist DROP COLUMN z;",
					Expected: []sql.Row{},
				},
			},
		},
	})
}

func TestAlterTableRenameColumn(t *testing.T) {
	RunScripts(t, []ScriptTest{
		{
			Name: "Rename Column",
			SetUpScript: []string{
				"CREATE TABLE test1 (a INT, b INT, c INT, d INT);",
				"INSERT INTO test1 VALUES (1, 2, 3, 4);",
			},
			Assertions: []ScriptTestAssertion{
				{
					Query:    "ALTER TABLE test1 RENAME COLUMN c to jjj;",
					Expected: []sql.Row{},
				},
				{
					Query:    "select * from test1 where jjj=3;",
					Expected: []sql.Row{{1, 2, 3, 4}},
				},
			},
		},
	})
}
