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

package _go

import (
	"testing"

	"github.com/dolthub/go-mysql-server/sql"
)

// https://www.postgresql.org/docs/15/functions-math.html
func TestFunctionsMath(t *testing.T) {
	RunScripts(t, []ScriptTest{
		{
			Name: "cbrt",
			SetUpScript: []string{
				`CREATE TABLE test (pk INT primary key, v1 INT, v2 FLOAT4, v3 FLOAT8, v4 VARCHAR(255));`,
				`INSERT INTO test VALUES (1, -1, -2, -3, '-5'), (2, 7, 11, 13, '17'), (3, 19, -23, 29, '-31');`,
			},
			Assertions: []ScriptTestAssertion{
				{
					Query: `SELECT cbrt(v1), cbrt(v2), cbrt(v3) FROM test ORDER BY pk;`,
					Skip:  true, // Our values are slightly different
					Expected: []sql.Row{
						{-1.0, -1.259921049894873, -1.4422495703074083},
						{1.9129311827723892, 2.2239800905693157, 2.3513346877207573},
						{2.668401648721945, -2.8438669798515654, 3.0723168256858475},
					},
				},
				{
					Query: `SELECT round(cbrt(v1)::numeric, 10), round(cbrt(v2)::numeric, 10), round(cbrt(v3)::numeric, 10) FROM test ORDER BY pk;`,
					Cols:  []string{"round", "round", "round"},
					Expected: []sql.Row{
						{-1.0000000000, -1.2599210499, -1.4422495703},
						{1.9129311828, 2.2239800906, 2.3513346877},
						{2.6684016487, -2.8438669799, 3.0723168257},
					},
				},
				{
					Query:       `SELECT cbrt(v4) FROM test ORDER BY pk;`,
					ExpectedErr: "function cbrt(varchar(255)) does not exist",
				},
				{
					Query: `SELECT cbrt('64');`,
					Cols:  []string{"cbrt"},
					Expected: []sql.Row{
						{4.0},
					},
				},
				{
					Query: `SELECT round(cbrt('64'));`,
					Cols:  []string{"round"},
					Expected: []sql.Row{
						{4.0},
					},
				},
			},
		},
		{
			Name: "gcd",
			SetUpScript: []string{
				`CREATE TABLE test (pk INT primary key, v1 INT4, v2 INT8, v3 FLOAT8, v4 VARCHAR(255));`,
				`INSERT INTO test VALUES (1, -2, -4, -6, '-8'), (2, 10, 12, 14.14, '16.16'), (3, 18, -20, 22.22, '-24.24');`,
			},
			Assertions: []ScriptTestAssertion{
				{
					Query: `SELECT gcd(v1, 10), gcd(v2, 20) FROM test ORDER BY pk;`,
					Cols:  []string{"gcd", "gcd"},
					Expected: []sql.Row{
						{2, 4},
						{10, 4},
						{2, 20},
					},
				},
				{
					Query:       `SELECT gcd(v3, 10) FROM test ORDER BY pk;`,
					ExpectedErr: "function gcd(double precision, integer)",
				},
				{
					Query:       `SELECT gcd(v4, 10) FROM test ORDER BY pk;`,
					ExpectedErr: "function gcd(varchar(255), integer) does not exist",
				},
				{
					Query: `SELECT gcd(36, '48');`,
					Cols:  []string{"gcd"},
					Expected: []sql.Row{
						{12},
					},
				},
				{
					Query: `SELECT gcd('36', 48);`,
					Cols:  []string{"gcd"},
					Expected: []sql.Row{
						{12},
					},
				},
				{
					Query: `SELECT gcd(1, 0), gcd(0, 1), gcd(0, 0);`,
					Cols:  []string{"gcd", "gcd", "gcd"},
					Expected: []sql.Row{
						{1, 1, 0},
					},
				},
			},
		},
		{
			Name: "lcm",
			SetUpScript: []string{
				`CREATE TABLE test (pk INT primary key, v1 INT4, v2 INT8, v3 FLOAT8, v4 VARCHAR(255));`,
				`INSERT INTO test VALUES (1, -2, -4, -6, '-8'), (2, 10, 12, 14.14, '16.16'), (3, 18, -20, 22.22, '-24.24');`,
			},
			Assertions: []ScriptTestAssertion{
				{
					Query: `SELECT lcm(v1, 10), lcm(v2, 20) FROM test ORDER BY pk;`,
					Cols:  []string{"lcm", "lcm"},
					Expected: []sql.Row{
						{10, 20},
						{10, 60},
						{90, 20},
					},
				},
				{
					Query:       `SELECT lcm(v3, 10) FROM test ORDER BY pk;`,
					ExpectedErr: "function lcm(double precision, integer)",
				},
				{
					Query:       `SELECT lcm(v4, 10) FROM test ORDER BY pk;`,
					ExpectedErr: "function lcm(varchar(255), integer) does not exist",
				},
				{
					Query: `SELECT lcm(36, '48');`,
					Expected: []sql.Row{
						{144},
					},
				},
				{
					Query: `SELECT lcm('36', 48);`,
					Expected: []sql.Row{
						{144},
					},
				},
				{
					Query: `SELECT lcm(1, 0), lcm(0, 1), lcm(0, 0);`,
					Expected: []sql.Row{
						{0, 0, 0},
					},
				},
			},
		},
	})
}

func TestFunctionsOID(t *testing.T) {
	RunScripts(t, []ScriptTest{
		{
			Name: "to_regclass",
			SetUpScript: []string{
				`CREATE TABLE testing (pk INT primary key, v1 INT UNIQUE);`,
				`CREATE TABLE "Testing2" (pk INT primary key, v1 INT);`,
			},
			Assertions: []ScriptTestAssertion{
				{
					Query: `SELECT to_regclass('testing');`,
					Cols:  []string{"to_regclass"},
					Expected: []sql.Row{
						{"testing"},
					},
				},
				{
					Query: `SELECT to_regclass('Testing2');`,
					Expected: []sql.Row{
						{nil},
					},
				},
				{
					Query: `SELECT to_regclass('"Testing2"');`,
					Expected: []sql.Row{
						{"Testing2"},
					},
				},
				{
					Query: `SELECT to_regclass(('testing'::regclass)::text);`,
					Expected: []sql.Row{
						{"testing"},
					},
				},
				{
					Query: `SELECT to_regclass((('testing'::regclass)::oid)::text);`,
					Expected: []sql.Row{
						{nil},
					},
				},
			},
		},
		{
			Name: "to_regproc",
			Assertions: []ScriptTestAssertion{
				{
					Query: `SELECT to_regproc('acos');`,
					Cols:  []string{"to_regproc"},
					Expected: []sql.Row{
						{"acos"},
					},
				},
				{
					Query: `SELECT to_regproc('acos"');`,
					Expected: []sql.Row{
						{nil},
					},
				},
				{
					Query: `SELECT to_regproc(('acos'::regproc)::text);`,
					Expected: []sql.Row{
						{"acos"},
					},
				},
				{
					Query: `SELECT to_regproc((('acos'::regproc)::oid)::text);`,
					Expected: []sql.Row{
						{nil},
					},
				},
			},
		},
		{
			Name: "to_regtype",
			Assertions: []ScriptTestAssertion{
				{
					Query: `SELECT to_regtype('integer');`,
					Cols:  []string{"to_regtype"},
					Expected: []sql.Row{
						{"integer"},
					},
				},
				{
					Skip:  true, // TODO: to_regtype should work with array types
					Query: `SELECT to_regtype('integer[]');`,
					Expected: []sql.Row{
						{"integer[]"},
					},
				},
				{
					Query: `SELECT to_regtype('int4');`,
					Expected: []sql.Row{
						{"integer"},
					},
				},
				{
					Query: `SELECT to_regtype('varchar');`,
					Expected: []sql.Row{
						{"character varying"},
					},
				},
				{
					Query: `SELECT to_regtype('varchar(10)');`,
					Expected: []sql.Row{
						{"character varying"},
					},
				},
				{
					Query: `SELECT to_regtype('timestamp');`,
					Expected: []sql.Row{
						{"timestamp without time zone"},
					},
				},
				{
					Query: `SELECT to_regtype('timestamp without time zone');`,
					Expected: []sql.Row{
						{"timestamp without time zone"},
					},
				},
				{
					Query: `SELECT to_regtype('integer"');`,
					Expected: []sql.Row{
						{nil},
					},
				},
				{
					Query: `SELECT to_regtype(('integer'::regtype)::text);`,
					Expected: []sql.Row{
						{"integer"},
					},
				},
				{
					Query: `SELECT to_regtype((('integer'::regtype)::oid)::text);`,
					Expected: []sql.Row{
						{nil},
					},
				},
			},
		},
	})
}

func TestSystemInformationFunctions(t *testing.T) {
	RunScripts(t, []ScriptTest{
		{
			Name:     "current_database",
			Database: "test",
			Assertions: []ScriptTestAssertion{
				{
					Query: `SELECT current_database();`,
					Cols:  []string{"current_database"},
					Expected: []sql.Row{
						{"test"},
					},
				},
				{
					Query:       `SELECT current_database;`,
					ExpectedErr: `column "current_database" could not be found in any table in scope`,
				},
				// TODO: Implement table function for current_database
				{
					Query: `SELECT * FROM current_database();`,
					Skip:  true,
					Expected: []sql.Row{
						{"test"},
					},
				},
				{
					Query:       `SELECT * FROM current_database;`,
					ExpectedErr: "table not found: current_database",
				},
			},
		},
		{
			Name:     "current_catalog",
			Database: "test",
			Assertions: []ScriptTestAssertion{
				{
					Skip:  true, // TODO: current_catalog currently returns current_database column name
					Query: `SELECT current_catalog;`,
					Cols:  []string{"current_catalog"},
					Expected: []sql.Row{
						{"test"},
					},
				},
				{
					Query: `SELECT current_catalog;`,
					Expected: []sql.Row{
						{"test"},
					},
				},
				{
					Query:       `SELECT current_catalog();`,
					ExpectedErr: `ERROR: at or near "(": syntax error (SQLSTATE XX000)`,
				},
				// // TODO: Implement table function for current_catalog
				{
					Query: `SELECT * FROM current_catalog;`,
					Skip:  true,
					Expected: []sql.Row{
						{"test"},
					},
				},
				{
					Query:       `SELECT * FROM current_catalog();`,
					ExpectedErr: `ERROR: at or near "(": syntax error (SQLSTATE XX000)`,
				},
			},
		},
		{
			Name: "current_schema",
			Assertions: []ScriptTestAssertion{
				{
					Skip:  true, // TODO: current_schema currently returns column name in quotes
					Query: `SELECT current_schema();`,
					Cols:  []string{"\"current_schema\""},
					Expected: []sql.Row{
						{"postgres"},
					},
				},
				{
					Query: `SELECT current_schema();`,
					Expected: []sql.Row{
						{"postgres"},
					},
				},
				{
					Query:    "CREATE SCHEMA test_schema;",
					Expected: []sql.Row{},
				},
				{
					Query:    `SET SEARCH_PATH TO test_schema;`,
					Expected: []sql.Row{},
				},
				{
					Query: `SELECT current_schema();`,
					Expected: []sql.Row{
						{"test_schema"},
					},
				},
				{
					Query:    `SET SEARCH_PATH TO public;`,
					Expected: []sql.Row{},
				},
				{
					Query: `SELECT current_schema();`,
					Expected: []sql.Row{
						{"public"},
					},
				},
				{
					Query: `SELECT current_schema;`,
					Expected: []sql.Row{
						{"public"},
					},
				},
				// TODO: Implement table function for current_schema
				{
					Query: `SELECT * FROM current_schema();`,
					Skip:  true,
					Expected: []sql.Row{
						{"public"},
					},
				},
				{
					Query: `SELECT * FROM current_schema;`,
					Skip:  true,
					Expected: []sql.Row{
						{"public"},
					},
				},
			},
		},
		{
			Name: "current_schemas",
			Assertions: []ScriptTestAssertion{
				{ // TODO: Not sure why Postgres does not display "$user", which is postgres here
					Query: `SELECT current_schemas(true);`,
					Cols:  []string{"current_schemas"},
					Expected: []sql.Row{
						{"{pg_catalog,postgres,public}"},
					},
				},
				{ // TODO: Not sure why Postgres does not display "$user" here
					Query: `SELECT current_schemas(false);`,
					Expected: []sql.Row{
						{"{postgres,public}"},
					},
				},
				{
					Query:    "CREATE SCHEMA test_schema;",
					Expected: []sql.Row{},
				},
				{
					Query:    `SET SEARCH_PATH TO test_schema;`,
					Expected: []sql.Row{},
				},
				{
					Query: `SELECT current_schemas(true);`,
					Expected: []sql.Row{
						{"{pg_catalog,test_schema}"},
					},
				},
				{
					Query: `SELECT current_schemas(false);`,
					Expected: []sql.Row{
						{"{test_schema}"},
					},
				},
				{
					Query:    `SET SEARCH_PATH TO public, test_schema;`,
					Expected: []sql.Row{},
				},
				{
					Query: `SELECT current_schemas(true);`,
					Expected: []sql.Row{
						{"{pg_catalog,public,test_schema}"},
					},
				},
				{
					Query: `SELECT current_schemas(false);`,
					Expected: []sql.Row{
						{"{public,test_schema}"},
					},
				},
			},
		},
		{
			Name: "version",
			Assertions: []ScriptTestAssertion{
				{
					Query: `SELECT version();`,
					Cols:  []string{"version"},
					Expected: []sql.Row{
						{"PostgreSQL 15.5"},
					},
				},
			},
		},
		{
			Name: "col_description",
			Assertions: []ScriptTestAssertion{
				{
					Query: `SELECT col_description(100, 1);`,
					Cols:  []string{"col_description"},
					Expected: []sql.Row{
						{""},
					},
				},
				{
					Query:       `SELECT col_description('not_a_table'::regclass, 1);`,
					ExpectedErr: `relation "not_a_table" does not exist`,
				},
				{
					Query:    `CREATE TABLE test_table (id INT);`,
					Expected: []sql.Row{},
				},
				{
					Skip:     true, // TODO: Implement column comments
					Query:    `COMMENT ON COLUMN test_table.id IS 'This is col id';`,
					Expected: []sql.Row{},
				},
				{
					Skip:  true, // TODO: Implement column object comments
					Query: `SELECT col_description('test_table'::regclass, 1);`,
					Cols:  []string{"col_description"},
					Expected: []sql.Row{
						{"This is col id"},
					},
				},
			},
		},
		{
			Name: "obj_description",
			Assertions: []ScriptTestAssertion{
				{
					Query: `SELECT obj_description(100, 'pg_class');`,
					Cols:  []string{"obj_description"},
					Expected: []sql.Row{
						{""},
					},
				},
				{
					Query:       `SELECT obj_description('does-not-exist'::regproc, 'pg_class');`,
					ExpectedErr: `function "does-not-exist" does not exist`,
				},
				{
					Skip:  true, // TODO: Implement database object comments
					Query: `SELECT obj_description('sinh'::regproc, 'pg_proc');`,
					Cols:  []string{"col_description"},
					Expected: []sql.Row{
						{"hyperbolic sine"},
					},
				},
			},
		},
		{
			Name: "shobj_description",
			Assertions: []ScriptTestAssertion{
				{
					Query: `SELECT shobj_description(100, 'pg_class');`,
					Cols:  []string{"shobj_description"},
					Expected: []sql.Row{
						{""},
					},
				},
				{
					Query:       `SELECT shobj_description('does-not-exist'::regproc, 'pg_class');`,
					ExpectedErr: `function "does-not-exist" does not exist`,
				},
				{
					Skip:     true, // TODO: Implement tablespaces
					Query:    `CREATE TABLESPACE tblspc_2 LOCATION '/';`,
					Expected: []sql.Row{},
				},
				{
					Skip:     true, // TODO: Implement shared database object comments
					Query:    `COMMENT ON TABLESPACE tblspc_2 IS 'Store a few of the things';`,
					Expected: []sql.Row{},
				},
				{
					Skip: true, // TODO: Implement shared database object comments
					Query: `SELECT shobj_description(
                 (SELECT oid FROM pg_tablespace WHERE spcname = 'tblspc_2'),
                 'pg_tablespace');`,
					Cols: []string{"shobj_description"},
					Expected: []sql.Row{
						{"Store a few of the things"},
					},
				},
			},
		},
		{
			Name: "format_type",
			Assertions: []ScriptTestAssertion{
				// Without typemod
				{
					Query: `SELECT format_type('integer'::regtype, null);`,
					Cols:  []string{"format_type"},
					Expected: []sql.Row{
						{"integer"},
					},
				},
				{
					Query: `SELECT format_type('character varying'::regtype, null);`,
					Expected: []sql.Row{
						{"character varying"},
					},
				},
				{
					Query: `SELECT format_type('varchar'::regtype, null);`,
					Expected: []sql.Row{
						{"character varying"},
					},
				},
				{
					Query: `SELECT format_type('date'::regtype, null);`,
					Expected: []sql.Row{
						{"date"},
					},
				},
				{
					Query: `SELECT format_type('timestamptz'::regtype, null);`,
					Expected: []sql.Row{
						{"timestamp with time zone"},
					},
				},
				{
					Query: `SELECT format_type('bool'::regtype, null);`,
					Expected: []sql.Row{
						{"boolean"},
					},
				},
				{
					Query: `SELECT format_type(1007, null);`,
					Expected: []sql.Row{
						{"integer[]"},
					},
				},
				{
					Query: `SELECT format_type('"char"'::regtype, null);`,
					Expected: []sql.Row{
						{"\"char\""},
					},
				},
				{
					Skip:  true, // TODO: regtype should work with array types
					Query: `SELECT format_type('real[]'::regtype, null);`,
					Expected: []sql.Row{
						{"real[]"},
					},
				},
				// With typemod
				{
					Query: `SELECT format_type('character varying'::regtype, 100);`,
					Expected: []sql.Row{
						{"character varying(96)"},
					},
				},
				{
					Query: `SELECT format_type('text'::regtype, 0);`,
					Expected: []sql.Row{
						{"text(0)"},
					},
				},
				{
					Query: `SELECT format_type('text'::regtype, 4);`,
					Expected: []sql.Row{
						{"text(4)"},
					},
				},
				{
					Query: `SELECT format_type('text'::regtype, -1);`,
					Expected: []sql.Row{
						{"text"},
					},
				},
				{
					Query: `SELECT format_type('name'::regtype, 0);`,
					Expected: []sql.Row{
						{"name(0)"},
					},
				},
				{
					Query: `SELECT format_type('bpchar'::regtype, -1);`,
					Expected: []sql.Row{
						{"bpchar"},
					},
				},
				{
					Query: `SELECT format_type('bpchar'::regtype, 10);`,
					Expected: []sql.Row{
						{"character(6)"},
					},
				},
				{
					Query: `SELECT format_type('bpchar'::regtype, 10);`,
					Expected: []sql.Row{
						{"character(6)"},
					},
				},
				{
					Query: `SELECT format_type('character'::regtype, 4);`,
					Expected: []sql.Row{
						{"character"},
					},
				},
				{
					Query: `SELECT format_type('varchar'::regtype, 0);`,
					Expected: []sql.Row{
						{"character varying"},
					},
				},
				{
					Query: `SELECT format_type('"char"'::regtype, 0);`,
					Expected: []sql.Row{
						{"\"char\"(0)"},
					},
				},
				{
					Query: `SELECT format_type('numeric'::regtype, 12);`,
					Expected: []sql.Row{
						{"numeric(0,8)"},
					},
				},
				// OID does not exist
				{
					Query: `SELECT format_type(874938247, 20);`,
					Cols:  []string{"format_type"},
					Expected: []sql.Row{
						{"???"},
					},
				},
				{
					Query: `SELECT format_type(874938247, null);`,
					Cols:  []string{"format_type"},
					Expected: []sql.Row{
						{"???"},
					},
				},
			},
		},
		{
			Name: "pg_get_constraintdef",
			SetUpScript: []string{
				`CREATE TABLE testing (pk INT primary key, v1 INT UNIQUE);`,
				`CREATE TABLE testing2 (pk INT primary key, pktesting INT REFERENCES testing(pk), v1 TEXT);`,
				`CREATE TABLE testing3 (pk1 INT, pk2 INT, PRIMARY KEY (pk1, pk2));`,
				// TODO: Uncomment when check constraints supported
				// `ALTER TABLE testing2 ADD CONSTRAINT v1_check CHECK (v1 != '');`,
			},
			Assertions: []ScriptTestAssertion{
				{
					Query:    `SELECT pg_get_constraintdef(845743985);`,
					Cols:     []string{"pg_get_constraintdef"},
					Expected: []sql.Row{{""}},
				},
				{
					Query: `SELECT pg_get_constraintdef(oid) FROM pg_catalog.pg_constraint WHERE conrelid='testing'::regclass;`,
					Cols:  []string{"pg_get_constraintdef"},
					Expected: []sql.Row{
						{"PRIMARY KEY (pk)"},
						{"UNIQUE (v1)"},
					},
				},
				{
					Query: `SELECT pg_get_constraintdef(oid) FROM pg_catalog.pg_constraint WHERE conrelid='testing2'::regclass LIMIT 1;`,
					Expected: []sql.Row{
						{"PRIMARY KEY (pk)"},
					},
				},
				{
					Skip:  true, // TODO: Foreign keys don't work
					Query: `SELECT pg_get_constraintdef(oid) FROM pg_catalog.pg_constraint WHERE conrelid='testing2'::regclass;`,
					Expected: []sql.Row{
						{"PRIMARY KEY (pk)"},
						{"FOREIGN KEY (pktesting) REFERENCES testing(pk)"},
					},
				},
				{
					Query: `SELECT pg_get_constraintdef(oid) FROM pg_catalog.pg_constraint WHERE conrelid='testing3'::regclass;`,
					Expected: []sql.Row{
						{"PRIMARY KEY (pk1, pk2)"},
					},
				},
				{
					Query:       `SELECT pg_get_constraintdef(oid, true) FROM pg_catalog.pg_constraint WHERE conrelid='testing3'::regclass;`,
					ExpectedErr: "pretty printing is not yet supported",
				},
				{
					Query: `SELECT pg_get_constraintdef(oid, false) FROM pg_catalog.pg_constraint WHERE conrelid='testing3'::regclass;`,
					Expected: []sql.Row{
						{"PRIMARY KEY (pk1, pk2)"},
					},
				},
			},
		},
	})
}
