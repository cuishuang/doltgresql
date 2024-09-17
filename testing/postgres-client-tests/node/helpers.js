const args = process.argv.slice(2);
const user = args[0];
const port = args[1];
const version = args[2];

export const dbName = "doltgres";

export function getArgs() {
  return { user, port };
}

export function getDoltgresVersion() {
  return version;
}

export function getConfig() {
  const { user, port } = getArgs();
  return {
    host: "localhost",
    port: port,
    database: dbName,
    user: user,
  };
}

export function assertQueryResult(q, expected, data, matcher) {
  if (matcher) {
    return matcher(data, expected);
  }
  if (q.toLowerCase().includes("dolt_commit")) {
    if (data.rows.length !== 1) return false;
    const hash = data.rows[0].dolt_commit.slice(1, -1);
    // dolt_commit row returns 32 character hash enclosed in brackets.
    return hash.length === 32;
  }
  if (q.toLowerCase().includes("dolt_merge")) {
    if (data.rows.length !== 1) return false;
    const res = data.rows[0].dolt_merge.slice(1, -1).split(",");
    const [hash, fastForward, conflicts, message] = res;
    return (
      hash.length === 32 &&
      expected.fastForward === fastForward &&
      expected.conflicts === conflicts &&
      expected.message === message
    );
  }

  // Does partial matching of actual and expected results.
  const partialRes = {
    command: data.command,
    rowCount: data.rowCount,
    oid: data.oid,
    rows: data.rows,
    fields: data.fields,
  };
  return JSON.stringify(expected) === JSON.stringify(partialRes);
}
