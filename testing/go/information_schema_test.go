package _go

import (
	"testing"

	"github.com/dolthub/go-mysql-server/sql"
)

func TestInfoSchemaColumns(t *testing.T) {
	RunScripts(t, []ScriptTest{
		{
			Name: "information_schema.columns",
			SetUpScript: []string{
				"create table test_table (id int)",
				"create view test_view as select * from test_table",
			},
			Assertions: []ScriptTestAssertion{
				{
					Query: `SELECT DISTINCT table_schema FROM information_schema.columns ORDER BY table_schema;`,
					Expected: []sql.Row{
						{"information_schema"}, {"pg_catalog"}, {"public"},
					},
				},
				{
					Query: `SELECT table_catalog, table_schema, table_name, column_name FROM information_schema.columns WHERE table_schema='public' ORDER BY table_name;`,
					Expected: []sql.Row{
						{"postgres", "public", "test_table", "id"},
						{"postgres", "public", "test_view", ""},
					},
				},
				{
					Query:    `CREATE SCHEMA test_schema;`,
					Expected: []sql.Row{},
				},
				{
					Query:    `SET SEARCH_PATH TO test_schema;`,
					Expected: []sql.Row{},
				},
				{
					Query:    `CREATE TABLE test_table2 (id2 INT);`,
					Expected: []sql.Row{},
				},
				{
					Query: `SELECT DISTINCT table_schema FROM information_schema.columns order by table_schema;`,
					Expected: []sql.Row{
						{"information_schema"}, {"pg_catalog"}, {"public"}, {"test_schema"},
					},
				},
				{
					Query: `SELECT table_catalog, table_schema, table_name, column_name FROM information_schema.columns WHERE table_schema='test_schema';`,
					Expected: []sql.Row{
						{"postgres", "test_schema", "test_table2", "id2"},
					},
				},
				{
					Skip:     true, // TODO: need ENUM type for column_type column
					Query:    "SELECT * FROM information_schema.columns;",
					Expected: []sql.Row{},
				},
				{
					Skip:  true, // TODO: table name in column returns `table not found: columns` error
					Query: `SELECT columns.table_name from "information_schema"."columns" WHERE table_name='test_table';`,
					Expected: []sql.Row{
						{"test_table"},
					},
				},
			},
		},
	})
}

func TestInfoSchemaSchemata(t *testing.T) {
	RunScripts(t, []ScriptTest{
		{
			Name:     "information_schema.schemata",
			Database: "newdb",
			SetUpScript: []string{
				"create schema test_schema",
			},
			Assertions: []ScriptTestAssertion{
				{
					Query: `SELECT catalog_name, schema_name FROM information_schema.schemata order by schema_name;`,
					Expected: []sql.Row{
						{"newdb", "information_schema"},
						{"newdb", "pg_catalog"},
						{"newdb", "public"},
						{"newdb", "test_schema"},
					},
				},
			},
		},
	})
}

func TestInfoSchemaTables(t *testing.T) {
	RunScripts(t, []ScriptTest{
		{
			Name: "information_schema.tables",
			SetUpScript: []string{
				"create table test_table (id int)",
			},
			Assertions: []ScriptTestAssertion{
				{
					Query: `SELECT DISTINCT table_schema FROM information_schema.tables order by table_schema;`,
					Expected: []sql.Row{
						{"information_schema"}, {"pg_catalog"}, {"public"},
					},
				},
				{
					Query: `SELECT table_catalog, table_schema FROM information_schema.tables group by table_catalog, table_schema order by table_schema;`,
					Expected: []sql.Row{
						{"postgres", "information_schema"},
						{"postgres", "pg_catalog"},
						{"postgres", "public"},
					},
				},
				{
					Query: `SELECT table_catalog, table_schema, table_name FROM information_schema.tables WHERE table_schema='public';`,
					Expected: []sql.Row{
						{"postgres", "public", "test_table"},
					},
				},
				{
					Query:    `CREATE SCHEMA test_schema;`,
					Expected: []sql.Row{},
				},
				{
					Query:    `SET SEARCH_PATH TO test_schema;`,
					Expected: []sql.Row{},
				},
				{
					Query:    `CREATE TABLE test_table2 (id INT);`,
					Expected: []sql.Row{},
				},
				{
					Query: `SELECT DISTINCT table_schema FROM information_schema.tables order by table_schema;`,
					Expected: []sql.Row{
						{"information_schema"}, {"pg_catalog"}, {"public"}, {"test_schema"},
					},
				},
				{
					Query: `SELECT table_catalog, table_schema FROM information_schema.tables group by table_catalog, table_schema order by table_schema;`,
					Expected: []sql.Row{
						{"postgres", "information_schema"},
						{"postgres", "pg_catalog"},
						{"postgres", "public"},
						{"postgres", "test_schema"},
					},
				},
				{
					Query: `SELECT table_catalog, table_schema, table_name FROM information_schema.tables WHERE table_schema='test_schema';`,
					Expected: []sql.Row{
						{"postgres", "test_schema", "test_table2"},
					},
				},
				{
					Skip:     true, // TODO: need ENUM type for table_type column
					Query:    "SELECT * FROM information_schema.tables ORDER BY table_name;",
					Expected: []sql.Row{},
				},
				{
					Query:    `SELECT "table_schema", "table_name", obj_description(('"' || "table_schema" || '"."' || "table_name" || '"')::regclass, 'pg_class') AS table_comment FROM "information_schema"."tables" WHERE ("table_schema" = 'test_schema' AND "table_name" = 'test_table2')`,
					Expected: []sql.Row{{"test_schema", "test_table2", ""}},
				},
			},
		},
	})
}
