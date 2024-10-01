import { doltCommitFields, doltDocsFields } from "../fields.js";

const readmeText = `# README
## My List

- Item 1
- Item 2
`;

const updatedReadmeText = `${readmeText}-Item 3`;

export const docsTests = [
  {
    q: "select * from dolt_docs",
    res: {
      command: "SELECT",
      rowCount: 0,
      oid: null,
      rows: [],
      fields: doltDocsFields,
    },
  },
  {
    q: "INSERT INTO dolt_docs VALUES ($1, $2);",
    p: ["README.md", readmeText],
    res: {
      command: "INSERT",
      rowCount: 1,
      oid: 0,
      rows: [],
      fields: [],
    },
  },
  {
    q: `select * from dolt_docs where doc_name=$1`,
    p: ["README.md"],
    res: {
      command: "SELECT",
      rowCount: 1,
      oid: null,
      rows: [{ doc_name: "README.md", doc_text: readmeText }],
      fields: doltDocsFields,
    },
  },
  {
    skip: true, // TODO: the ON CONFLICT clause provided is not yet supported
    q: "INSERT INTO schema.searches VALUES ($1, $2) ON CONFLICT (doc_name) DO UPDATE SET doc_text = $2",
    p: ["README.md", updatedReadmeText],
    res: {
      command: "INSERT",
      rowCount: null,
      oid: null,
      rows: [],
      fields: [],
    },
  },
  {
    skip: true,
    q: `select * from dolt_docs where doc_name=$1`,
    p: ["README.md"],
    res: {
      command: "SELECT",
      rowCount: 1,
      oid: null,
      rows: [{ doc_name: "README.md", doc_text: updatedReadmeText }],
      fields: [],
    },
  },
  {
    q: `SELECT DOLT_COMMIT('-A', '-m', $1::text)`,
    p: ["Add dolt_docs table"],
    res: {
      command: "SELECT",
      rowCount: 1,
      oid: null,
      rows: [{ dolt_commit: "" }],
      fields: doltCommitFields,
    },
  },
  {
    q: `SET SEARCH_PATH = 'testschema';`,
    res: {
      command: "SET",
      rowCount: null,
      oid: null,
      rows: [],
      fields: [],
    },
  },
  // dolt_docs should not be schema-specific
  {
    q: `select * from dolt_docs where doc_name=$1`,
    p: ["README.md"],
    res: {
      command: "SELECT",
      rowCount: 1,
      oid: null,
      rows: [{ doc_name: "README.md", doc_text: readmeText }],
      fields: doltDocsFields,
    },
  },
  {
    q: "DELETE FROM dolt_docs WHERE doc_name=$1",
    p: ["README.md"],
    res: {
      command: "DELETE",
      rowCount: 1,
      oid: null,
      rows: [],
      fields: [],
    },
  },
  {
    q: `select * from dolt_docs where doc_name=$1`,
    p: ["README.md"],
    res: {
      command: "SELECT",
      rowCount: 0,
      oid: null,
      rows: [],
      fields: doltDocsFields,
    },
  },
  {
    q: `SET SEARCH_PATH = 'public';`,
    res: {
      command: "SET",
      rowCount: null,
      oid: null,
      rows: [],
      fields: [],
    },
  },
  {
    q: `select * from dolt_docs where doc_name=$1`,
    p: ["README.md"],
    res: {
      command: "SELECT",
      rowCount: 0,
      oid: null,
      rows: [],
      fields: doltDocsFields,
    },
  },
  {
    q: `SELECT DOLT_COMMIT('-A', '-m', $1::text)`,
    p: ["Remove README"],
    res: {
      command: "SELECT",
      rowCount: 1,
      oid: null,
      rows: [{ dolt_commit: "" }],
      fields: doltCommitFields,
    },
  },
];
