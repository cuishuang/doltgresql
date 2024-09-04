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

package regression

import (
	"testing"

	"github.com/dolthub/go-mysql-server/sql"
)

func TestPsqlCrosstab(t *testing.T) {
	t.Skip()
	_ = RunTests(t, RegressionFileName_psql_crosstab)
}

func init() {
	RegisterRegressionFile(RegressionFile{
		RegressionFileName: RegressionFileName_psql_crosstab,
		DependsOn:          []RegressionFileName{RegressionFileName_test_setup},
		Statements: []RegressionFileStatement{
			{
				Statement: `CREATE TABLE ctv_data (v, h, c, i, d) AS
VALUES
   ('v1','h2','foo', 3, '2015-04-01'::date),
   ('v2','h1','bar', 3, '2015-01-02'),
   ('v1','h0','baz', NULL, '2015-07-12'),
   ('v0','h4','qux', 4, '2015-07-15'),
   ('v0','h4','dbl', -3, '2014-12-15'),
   ('v0',NULL,'qux', 5, '2014-07-15'),
   ('v1','h2','quux',7, '2015-04-04');`,
			},
			{
				Statement: `ANALYZE ctv_data;`,
			},
			{
				Statement: `SELECT v, EXTRACT(year FROM d), count(*)
 FROM ctv_data
 GROUP BY 1, 2
 ORDER BY 1, 2;`,
				Results: []sql.Row{{`v0`, 2014, 2}, {`v0`, 2015, 1}, {`v1`, 2015, 3}, {`v2`, 2015, 1}},
			},
			{
				Statement: ` \crosstabview
 v  | 2014 | 2015 
----+------+------
 v0 |    2 |    1
 v1 |      |    3
 v2 |      |    1
(3 rows)
SELECT v, to_char(d, 'Mon') AS "month name", EXTRACT(month FROM d) AS num,
 count(*) FROM ctv_data  GROUP BY 1,2,3 ORDER BY 1
 \crosstabview v "month name" 4 num
 v  | Jan | Apr | Jul | Dec 
----+-----+-----+-----+-----
 v0 |     |     |   2 |   1
 v1 |     |   2 |   1 |    
 v2 |   1 |     |     |    
(3 rows)
SELECT EXTRACT(year FROM d) AS year, to_char(d,'Mon') AS """month"" name",
  EXTRACT(month FROM d) AS month,
  format('sum=%s avg=%s', sum(i), avg(i)::numeric(2,1))
  FROM ctv_data
  GROUP BY EXTRACT(year FROM d), to_char(d,'Mon'), EXTRACT(month FROM d)
ORDER BY month
\crosstabview """month"" name" year format year
 "month" name |      2014       |      2015      
--------------+-----------------+----------------
 Jan          |                 | sum=3 avg=3.0
 Apr          |                 | sum=10 avg=5.0
 Jul          | sum=5 avg=5.0   | sum=4 avg=4.0
 Dec          | sum=-3 avg=-3.0 | 
(4 rows)
SELECT v, h, string_agg(c, E'\n') FROM ctv_data GROUP BY v, h ORDER BY 1,2,3
 \crosstabview 1 2 3
 v  | h4  |     | h0  |  h2  | h1  
----+-----+-----+-----+------+-----
 v0 | qux+| qux |     |      | 
    | dbl |     |     |      | 
 v1 |     |     | baz | foo +| 
    |     |     |     | quux | 
 v2 |     |     |     |      | bar
(3 rows)
SELECT v,h, string_agg(c, E'\n') AS c, row_number() OVER(ORDER BY h) AS r
FROM ctv_data GROUP BY v, h ORDER BY 1,3,2
 \crosstabview v h c r
 v  | h0  | h1  |  h2  | h4  |     
----+-----+-----+------+-----+-----
 v0 |     |     |      | qux+| qux
    |     |     |      | dbl | 
 v1 | baz |     | foo +|     | 
    |     |     | quux |     | 
 v2 |     | bar |      |     | 
(3 rows)
SELECT v, h, string_agg(c, E'\n') AS c, row_number() OVER(ORDER BY h DESC) AS r
FROM ctv_data GROUP BY v, h ORDER BY 1,3,2
 \crosstabview v h c r
 v  |     | h4  |  h2  | h1  | h0  
----+-----+-----+------+-----+-----
 v0 | qux | qux+|      |     | 
    |     | dbl |      |     | 
 v1 |     |     | foo +|     | baz
    |     |     | quux |     | 
 v2 |     |     |      | bar | 
(3 rows)
SELECT v,h, string_agg(c, E'\n') AS c, row_number() OVER(ORDER BY h NULLS LAST) AS r
FROM ctv_data GROUP BY v, h ORDER BY 1,3,2
 \crosstabview v h c r
 v  | h0  | h1  |  h2  | h4  |     
----+-----+-----+------+-----+-----
 v0 |     |     |      | qux+| qux
    |     |     |      | dbl | 
 v1 | baz |     | foo +|     | 
    |     |     | quux |     | 
 v2 |     | bar |      |     | 
(3 rows)
SELECT null,null \crosstabview
\crosstabview: query must return at least three columns
SELECT null,null,null \crosstabview
 ?column? |  
----------+--
          | 
(1 row)
\pset null '#null#'
SELECT v,h, string_agg(i::text, E'\n') AS i FROM ctv_data
GROUP BY v, h ORDER BY h,v
 \crosstabview v h i
 v  |   h0   | h1 | h2 | h4 | #null# 
----+--------+----+----+----+--------
 v1 | #null# |    | 3 +|    | 
    |        |    | 7  |    | 
 v2 |        | 3  |    |    | 
 v0 |        |    |    | 4 +| 5
    |        |    |    | -3 | 
(3 rows)
\pset null ''
SELECT v,h,string_agg(i::text, E'\n'), string_agg(c, E'\n')
FROM ctv_data GROUP BY v, h ORDER BY h,v
 \crosstabview 2 1 4
 h  |  v1  | v2  | v0  
----+------+-----+-----
 h0 | baz  |     | 
 h1 |      | bar | 
 h2 | foo +|     | 
    | quux |     | 
 h4 |      |     | qux+
    |      |     | dbl
    |      |     | qux
(5 rows)
SELECT v,h, string_agg(i::text, E'\n') AS i, string_agg(c, E'\n') AS c
FROM ctv_data GROUP BY v, h ORDER BY h,v
 \crosstabview 1 "h" 4
 v  | h0  | h1  |  h2  | h4  |     
----+-----+-----+------+-----+-----
 v1 | baz |     | foo +|     | 
    |     |     | quux |     | 
 v2 |     | bar |      |     | 
 v0 |     |     |      | qux+| qux
    |     |     |      | dbl | 
(3 rows)
SELECT 1 as "22", 2 as b, 3 as "Foo"
 \crosstabview "22" B "Foo"
 22 | 2 
----+---
  1 | 3
(1 row)
SELECT v,h,c,i FROM ctv_data
 \crosstabview v h j
\crosstabview: column name not found: "j"
SELECT 1 as "22", 2 as b, 3 as "Foo"
 \crosstabview 1 2 Foo
\crosstabview: column name not found: "foo"
SELECT 1 as "22", 2 as b, 3 as "Foo"
 \crosstabview 1 "B" "Foo"
\crosstabview: column name not found: "B"
SELECT v,h,i,c FROM ctv_data
 \crosstabview 2 1 5
\crosstabview: column number 5 is out of range 1..4
SELECT v,h,i,c FROM ctv_data
 \crosstabview 2 h 4
\crosstabview: vertical and horizontal headers must be different columns
SELECT a,a,1 FROM generate_series(1,3000) AS a
 \crosstabview
\crosstabview: maximum number of columns (1600) exceeded
SELECT 1 \crosstabview
\crosstabview: query must return at least three columns
DROP TABLE ctv_data;`,
			},
			{
				Statement: `CREATE TABLE ctv_data (x int, y int, v text);`,
			},
			{
				Statement: `INSERT INTO ctv_data SELECT 1, x, '*' || x FROM generate_series(1,10) x;`,
			},
			{
				Statement: `SELECT * FROM ctv_data \crosstabview
 x | 1  | 2  | 3  | 4  | 5  | 6  | 7  | 8  | 9  | 10  
---+----+----+----+----+----+----+----+----+----+-----
 1 | *1 | *2 | *3 | *4 | *5 | *6 | *7 | *8 | *9 | *10
(1 row)
INSERT INTO ctv_data VALUES (1, 10, '*'); -- duplicate data to cause error`,
			},
			{
				Statement: `SELECT * FROM ctv_data \crosstabview
\crosstabview: query result contains multiple data values for row "1", column "10"
DROP TABLE ctv_data;`,
			},
		},
	})
}
