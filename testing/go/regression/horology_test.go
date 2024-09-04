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

func TestHorology(t *testing.T) {
	t.Skip()
	_ = RunTests(t, RegressionFileName_horology)
}

func init() {
	RegisterRegressionFile(RegressionFile{
		RegressionFileName: RegressionFileName_horology,
		DependsOn:          []RegressionFileName{RegressionFileName_test_setup, RegressionFileName_date, RegressionFileName_time, RegressionFileName_timetz, RegressionFileName_timestamp, RegressionFileName_timestamptz, RegressionFileName_interval},
		Statements: []RegressionFileStatement{
			{
				Statement: `SET DateStyle = 'Postgres, MDY';`,
			},
			{
				Statement: `SHOW TimeZone;  -- Many of these tests depend on the prevailing setting`,
				Results:   []sql.Row{{`PST8PDT`}},
			},
			{
				Statement: `SELECT timestamp with time zone '20011227 040506+08';`,
				Results:   []sql.Row{{`Wed Dec 26 12:05:06 2001 PST`}},
			},
			{
				Statement: `SELECT timestamp with time zone '20011227 040506-08';`,
				Results:   []sql.Row{{`Thu Dec 27 04:05:06 2001 PST`}},
			},
			{
				Statement: `SELECT timestamp with time zone '20011227 040506.789+08';`,
				Results:   []sql.Row{{`Wed Dec 26 12:05:06.789 2001 PST`}},
			},
			{
				Statement: `SELECT timestamp with time zone '20011227 040506.789-08';`,
				Results:   []sql.Row{{`Thu Dec 27 04:05:06.789 2001 PST`}},
			},
			{
				Statement: `SELECT timestamp with time zone '20011227T040506+08';`,
				Results:   []sql.Row{{`Wed Dec 26 12:05:06 2001 PST`}},
			},
			{
				Statement: `SELECT timestamp with time zone '20011227T040506-08';`,
				Results:   []sql.Row{{`Thu Dec 27 04:05:06 2001 PST`}},
			},
			{
				Statement: `SELECT timestamp with time zone '20011227T040506.789+08';`,
				Results:   []sql.Row{{`Wed Dec 26 12:05:06.789 2001 PST`}},
			},
			{
				Statement: `SELECT timestamp with time zone '20011227T040506.789-08';`,
				Results:   []sql.Row{{`Thu Dec 27 04:05:06.789 2001 PST`}},
			},
			{
				Statement: `SELECT timestamp with time zone '2001-12-27 04:05:06.789-08';`,
				Results:   []sql.Row{{`Thu Dec 27 04:05:06.789 2001 PST`}},
			},
			{
				Statement: `SELECT timestamp with time zone '2001.12.27 04:05:06.789-08';`,
				Results:   []sql.Row{{`Thu Dec 27 04:05:06.789 2001 PST`}},
			},
			{
				Statement: `SELECT timestamp with time zone '2001/12/27 04:05:06.789-08';`,
				Results:   []sql.Row{{`Thu Dec 27 04:05:06.789 2001 PST`}},
			},
			{
				Statement: `SELECT timestamp with time zone '12/27/2001 04:05:06.789-08';`,
				Results:   []sql.Row{{`Thu Dec 27 04:05:06.789 2001 PST`}},
			},
			{
				Statement:   `SELECT timestamp with time zone '27/12/2001 04:05:06.789-08';`,
				ErrorString: `date/time field value out of range: "27/12/2001 04:05:06.789-08"`,
			},
			{
				Statement: `set datestyle to dmy;`,
			},
			{
				Statement: `SELECT timestamp with time zone '27/12/2001 04:05:06.789-08';`,
				Results:   []sql.Row{{`Thu 27 Dec 04:05:06.789 2001 PST`}},
			},
			{
				Statement: `reset datestyle;`,
			},
			{
				Statement: `SELECT timestamp with time zone 'Y2001M12D27H04M05S06.789+08';`,
				Results:   []sql.Row{{`Wed Dec 26 12:05:06.789 2001 PST`}},
			},
			{
				Statement: `SELECT timestamp with time zone 'Y2001M12D27H04M05S06.789-08';`,
				Results:   []sql.Row{{`Thu Dec 27 04:05:06.789 2001 PST`}},
			},
			{
				Statement: `SELECT timestamp with time zone 'Y2001M12D27H04MM05S06.789+08';`,
				Results:   []sql.Row{{`Wed Dec 26 12:05:06.789 2001 PST`}},
			},
			{
				Statement: `SELECT timestamp with time zone 'Y2001M12D27H04MM05S06.789-08';`,
				Results:   []sql.Row{{`Thu Dec 27 04:05:06.789 2001 PST`}},
			},
			{
				Statement: `SELECT timestamp with time zone 'J2452271+08';`,
				Results:   []sql.Row{{`Wed Dec 26 08:00:00 2001 PST`}},
			},
			{
				Statement: `SELECT timestamp with time zone 'J2452271-08';`,
				Results:   []sql.Row{{`Thu Dec 27 00:00:00 2001 PST`}},
			},
			{
				Statement: `SELECT timestamp with time zone 'J2452271.5+08';`,
				Results:   []sql.Row{{`Wed Dec 26 20:00:00 2001 PST`}},
			},
			{
				Statement: `SELECT timestamp with time zone 'J2452271.5-08';`,
				Results:   []sql.Row{{`Thu Dec 27 12:00:00 2001 PST`}},
			},
			{
				Statement: `SELECT timestamp with time zone 'J2452271 04:05:06+08';`,
				Results:   []sql.Row{{`Wed Dec 26 12:05:06 2001 PST`}},
			},
			{
				Statement: `SELECT timestamp with time zone 'J2452271 04:05:06-08';`,
				Results:   []sql.Row{{`Thu Dec 27 04:05:06 2001 PST`}},
			},
			{
				Statement: `SELECT timestamp with time zone 'J2452271T040506+08';`,
				Results:   []sql.Row{{`Wed Dec 26 12:05:06 2001 PST`}},
			},
			{
				Statement: `SELECT timestamp with time zone 'J2452271T040506-08';`,
				Results:   []sql.Row{{`Thu Dec 27 04:05:06 2001 PST`}},
			},
			{
				Statement: `SELECT timestamp with time zone 'J2452271T040506.789+08';`,
				Results:   []sql.Row{{`Wed Dec 26 12:05:06.789 2001 PST`}},
			},
			{
				Statement: `SELECT timestamp with time zone 'J2452271T040506.789-08';`,
				Results:   []sql.Row{{`Thu Dec 27 04:05:06.789 2001 PST`}},
			},
			{
				Statement: `SELECT timestamp with time zone '12.27.2001 04:05:06.789+08';`,
				Results:   []sql.Row{{`Wed Dec 26 12:05:06.789 2001 PST`}},
			},
			{
				Statement: `SELECT timestamp with time zone '12.27.2001 04:05:06.789-08';`,
				Results:   []sql.Row{{`Thu Dec 27 04:05:06.789 2001 PST`}},
			},
			{
				Statement: `SET DateStyle = 'German';`,
			},
			{
				Statement: `SELECT timestamp with time zone '27.12.2001 04:05:06.789+08';`,
				Results:   []sql.Row{{`26.12.2001 12:05:06.789 PST`}},
			},
			{
				Statement: `SELECT timestamp with time zone '27.12.2001 04:05:06.789-08';`,
				Results:   []sql.Row{{`27.12.2001 04:05:06.789 PST`}},
			},
			{
				Statement: `SET DateStyle = 'ISO';`,
			},
			{
				Statement: `SELECT time without time zone '040506.789+08';`,
				Results:   []sql.Row{{`04:05:06.789`}},
			},
			{
				Statement: `SELECT time without time zone '040506.789-08';`,
				Results:   []sql.Row{{`04:05:06.789`}},
			},
			{
				Statement: `SELECT time without time zone 'T040506.789+08';`,
				Results:   []sql.Row{{`04:05:06.789`}},
			},
			{
				Statement: `SELECT time without time zone 'T040506.789-08';`,
				Results:   []sql.Row{{`04:05:06.789`}},
			},
			{
				Statement: `SELECT time with time zone '040506.789+08';`,
				Results:   []sql.Row{{`04:05:06.789+08`}},
			},
			{
				Statement: `SELECT time with time zone '040506.789-08';`,
				Results:   []sql.Row{{`04:05:06.789-08`}},
			},
			{
				Statement: `SELECT time with time zone 'T040506.789+08';`,
				Results:   []sql.Row{{`04:05:06.789+08`}},
			},
			{
				Statement: `SELECT time with time zone 'T040506.789-08';`,
				Results:   []sql.Row{{`04:05:06.789-08`}},
			},
			{
				Statement: `SELECT time with time zone 'T040506.789 +08';`,
				Results:   []sql.Row{{`04:05:06.789+08`}},
			},
			{
				Statement: `SELECT time with time zone 'T040506.789 -08';`,
				Results:   []sql.Row{{`04:05:06.789-08`}},
			},
			{
				Statement: `SET DateStyle = 'Postgres, MDY';`,
			},
			{
				Statement: `SELECT date 'J1520447' AS "Confucius' Birthday";`,
				Results:   []sql.Row{{`09-28-0551 BC`}},
			},
			{
				Statement: `SELECT date 'J0' AS "Julian Epoch";`,
				Results:   []sql.Row{{`11-24-4714 BC`}},
			},
			{
				Statement: `SELECT date '1981-02-03' + time '04:05:06' AS "Date + Time";`,
				Results:   []sql.Row{{`Tue Feb 03 04:05:06 1981`}},
			},
			{
				Statement: `SELECT date '1991-02-03' + time with time zone '04:05:06 PST' AS "Date + Time PST";`,
				Results:   []sql.Row{{`Sun Feb 03 04:05:06 1991 PST`}},
			},
			{
				Statement: `SELECT date '2001-02-03' + time with time zone '04:05:06 UTC' AS "Date + Time UTC";`,
				Results:   []sql.Row{{`Fri Feb 02 20:05:06 2001 PST`}},
			},
			{
				Statement: `SELECT date '1991-02-03' + interval '2 years' AS "Add Two Years";`,
				Results:   []sql.Row{{`Wed Feb 03 00:00:00 1993`}},
			},
			{
				Statement: `SELECT date '2001-12-13' - interval '2 years' AS "Subtract Two Years";`,
				Results:   []sql.Row{{`Mon Dec 13 00:00:00 1999`}},
			},
			{
				Statement: `SELECT date '1991-02-03' - time '04:05:06' AS "Subtract Time";`,
				Results:   []sql.Row{{`Sat Feb 02 19:54:54 1991`}},
			},
			{
				Statement:   `SELECT date '1991-02-03' - time with time zone '04:05:06 UTC' AS "Subtract Time UTC";`,
				ErrorString: `operator does not exist: date - time with time zone`,
			},
			{
				Statement: `SELECT timestamp without time zone '1996-03-01' - interval '1 second' AS "Feb 29";`,
				Results:   []sql.Row{{`Thu Feb 29 23:59:59 1996`}},
			},
			{
				Statement: `SELECT timestamp without time zone '1999-03-01' - interval '1 second' AS "Feb 28";`,
				Results:   []sql.Row{{`Sun Feb 28 23:59:59 1999`}},
			},
			{
				Statement: `SELECT timestamp without time zone '2000-03-01' - interval '1 second' AS "Feb 29";`,
				Results:   []sql.Row{{`Tue Feb 29 23:59:59 2000`}},
			},
			{
				Statement: `SELECT timestamp without time zone '1999-12-01' + interval '1 month - 1 second' AS "Dec 31";`,
				Results:   []sql.Row{{`Fri Dec 31 23:59:59 1999`}},
			},
			{
				Statement: `SELECT timestamp without time zone 'Jan 1, 4713 BC' + interval '106000000 days' AS "Feb 23, 285506";`,
				Results:   []sql.Row{{`Fri Feb 23 00:00:00 285506`}},
			},
			{
				Statement: `SELECT timestamp without time zone 'Jan 1, 4713 BC' + interval '107000000 days' AS "Jan 20, 288244";`,
				Results:   []sql.Row{{`Sat Jan 20 00:00:00 288244`}},
			},
			{
				Statement: `SELECT timestamp without time zone 'Jan 1, 4713 BC' + interval '109203489 days' AS "Dec 31, 294276";`,
				Results:   []sql.Row{{`Sun Dec 31 00:00:00 294276`}},
			},
			{
				Statement: `SELECT timestamp without time zone '12/31/294276' - timestamp without time zone '12/23/1999' AS "106751991 Days";`,
				Results:   []sql.Row{{`@ 106751991 days`}},
			},
			{
				Statement: `SELECT (timestamp without time zone 'today' = (timestamp without time zone 'yesterday' + interval '1 day')) as "True";`,
				Results:   []sql.Row{{true}},
			},
			{
				Statement: `SELECT (timestamp without time zone 'today' = (timestamp without time zone 'tomorrow' - interval '1 day')) as "True";`,
				Results:   []sql.Row{{true}},
			},
			{
				Statement: `SELECT (timestamp without time zone 'today 10:30' = (timestamp without time zone 'yesterday' + interval '1 day 10 hr 30 min')) as "True";`,
				Results:   []sql.Row{{true}},
			},
			{
				Statement: `SELECT (timestamp without time zone '10:30 today' = (timestamp without time zone 'yesterday' + interval '1 day 10 hr 30 min')) as "True";`,
				Results:   []sql.Row{{true}},
			},
			{
				Statement: `SELECT (timestamp without time zone 'tomorrow' = (timestamp without time zone 'yesterday' + interval '2 days')) as "True";`,
				Results:   []sql.Row{{true}},
			},
			{
				Statement: `SELECT (timestamp without time zone 'tomorrow 16:00:00' = (timestamp without time zone 'today' + interval '1 day 16 hours')) as "True";`,
				Results:   []sql.Row{{true}},
			},
			{
				Statement: `SELECT (timestamp without time zone '16:00:00 tomorrow' = (timestamp without time zone 'today' + interval '1 day 16 hours')) as "True";`,
				Results:   []sql.Row{{true}},
			},
			{
				Statement: `SELECT (timestamp without time zone 'yesterday 12:34:56' = (timestamp without time zone 'tomorrow' - interval '2 days - 12:34:56')) as "True";`,
				Results:   []sql.Row{{true}},
			},
			{
				Statement: `SELECT (timestamp without time zone '12:34:56 yesterday' = (timestamp without time zone 'tomorrow' - interval '2 days - 12:34:56')) as "True";`,
				Results:   []sql.Row{{true}},
			},
			{
				Statement: `SELECT (timestamp without time zone 'tomorrow' > 'now') as "True";`,
				Results:   []sql.Row{{true}},
			},
			{
				Statement: `SELECT date '1994-01-01' + time '11:00' AS "Jan_01_1994_11am";`,
				Results:   []sql.Row{{`Sat Jan 01 11:00:00 1994`}},
			},
			{
				Statement: `SELECT date '1994-01-01' + time '10:00' AS "Jan_01_1994_10am";`,
				Results:   []sql.Row{{`Sat Jan 01 10:00:00 1994`}},
			},
			{
				Statement: `SELECT date '1994-01-01' + timetz '11:00-5' AS "Jan_01_1994_8am";`,
				Results:   []sql.Row{{`Sat Jan 01 08:00:00 1994 PST`}},
			},
			{
				Statement: `SELECT timestamptz(date '1994-01-01', time with time zone '11:00-5') AS "Jan_01_1994_8am";`,
				Results:   []sql.Row{{`Sat Jan 01 08:00:00 1994 PST`}},
			},
			{
				Statement: `SELECT d1 + interval '1 year' AS one_year FROM TIMESTAMP_TBL;`,
				Results:   []sql.Row{{`-infinity`}, {`infinity`}, {`Fri Jan 01 00:00:00 1971`}, {`Tue Feb 10 17:32:01 1998`}, {`Tue Feb 10 17:32:01 1998`}, {`Tue Feb 10 17:32:02 1998`}, {`Tue Feb 10 17:32:01.4 1998`}, {`Tue Feb 10 17:32:01.5 1998`}, {`Tue Feb 10 17:32:01.6 1998`}, {`Fri Jan 02 00:00:00 1998`}, {`Fri Jan 02 03:04:05 1998`}, {`Tue Feb 10 17:32:01 1998`}, {`Tue Feb 10 17:32:01 1998`}, {`Tue Feb 10 17:32:01 1998`}, {`Tue Feb 10 17:32:01 1998`}, {`Wed Jun 10 17:32:01 1998`}, {`Sun Sep 22 18:19:20 2002`}, {`Thu Mar 15 08:14:01 2001`}, {`Thu Mar 15 13:14:02 2001`}, {`Thu Mar 15 12:14:03 2001`}, {`Thu Mar 15 03:14:04 2001`}, {`Thu Mar 15 02:14:05 2001`}, {`Tue Feb 10 17:32:01 1998`}, {`Tue Feb 10 17:32:01 1998`}, {`Tue Feb 10 17:32:00 1998`}, {`Tue Feb 10 17:32:01 1998`}, {`Tue Feb 10 17:32:01 1998`}, {`Tue Feb 10 17:32:01 1998`}, {`Tue Feb 10 17:32:01 1998`}, {`Tue Feb 10 17:32:01 1998`}, {`Tue Feb 10 17:32:01 1998`}, {`Tue Feb 10 17:32:01 1998`}, {`Tue Feb 10 17:32:01 1998`}, {`Tue Feb 10 17:32:01 1998`}, {`Wed Jun 10 18:32:01 1998`}, {`Tue Feb 10 17:32:01 1998`}, {`Wed Feb 11 17:32:01 1998`}, {`Thu Feb 12 17:32:01 1998`}, {`Fri Feb 13 17:32:01 1998`}, {`Sat Feb 14 17:32:01 1998`}, {`Sun Feb 15 17:32:01 1998`}, {`Mon Feb 16 17:32:01 1998`}, {`Thu Feb 16 17:32:01 0096 BC`}, {`Sun Feb 16 17:32:01 0098`}, {`Fri Feb 16 17:32:01 0598`}, {`Wed Feb 16 17:32:01 1098`}, {`Sun Feb 16 17:32:01 1698`}, {`Fri Feb 16 17:32:01 1798`}, {`Wed Feb 16 17:32:01 1898`}, {`Mon Feb 16 17:32:01 1998`}, {`Sun Feb 16 17:32:01 2098`}, {`Fri Feb 28 17:32:01 1997`}, {`Fri Feb 28 17:32:01 1997`}, {`Sat Mar 01 17:32:01 1997`}, {`Tue Dec 30 17:32:01 1997`}, {`Wed Dec 31 17:32:01 1997`}, {`Thu Jan 01 17:32:01 1998`}, {`Sat Feb 28 17:32:01 1998`}, {`Sun Mar 01 17:32:01 1998`}, {`Wed Dec 30 17:32:01 1998`}, {`Thu Dec 31 17:32:01 1998`}, {`Sun Dec 31 17:32:01 2000`}, {`Mon Jan 01 17:32:01 2001`}, {`Mon Dec 31 17:32:01 2001`}, {`Tue Jan 01 17:32:01 2002`}},
			},
			{
				Statement: `SELECT d1 - interval '1 year' AS one_year FROM TIMESTAMP_TBL;`,
				Results:   []sql.Row{{`-infinity`}, {`infinity`}, {`Wed Jan 01 00:00:00 1969`}, {`Sat Feb 10 17:32:01 1996`}, {`Sat Feb 10 17:32:01 1996`}, {`Sat Feb 10 17:32:02 1996`}, {`Sat Feb 10 17:32:01.4 1996`}, {`Sat Feb 10 17:32:01.5 1996`}, {`Sat Feb 10 17:32:01.6 1996`}, {`Tue Jan 02 00:00:00 1996`}, {`Tue Jan 02 03:04:05 1996`}, {`Sat Feb 10 17:32:01 1996`}, {`Sat Feb 10 17:32:01 1996`}, {`Sat Feb 10 17:32:01 1996`}, {`Sat Feb 10 17:32:01 1996`}, {`Mon Jun 10 17:32:01 1996`}, {`Fri Sep 22 18:19:20 2000`}, {`Mon Mar 15 08:14:01 1999`}, {`Mon Mar 15 13:14:02 1999`}, {`Mon Mar 15 12:14:03 1999`}, {`Mon Mar 15 03:14:04 1999`}, {`Mon Mar 15 02:14:05 1999`}, {`Sat Feb 10 17:32:01 1996`}, {`Sat Feb 10 17:32:01 1996`}, {`Sat Feb 10 17:32:00 1996`}, {`Sat Feb 10 17:32:01 1996`}, {`Sat Feb 10 17:32:01 1996`}, {`Sat Feb 10 17:32:01 1996`}, {`Sat Feb 10 17:32:01 1996`}, {`Sat Feb 10 17:32:01 1996`}, {`Sat Feb 10 17:32:01 1996`}, {`Sat Feb 10 17:32:01 1996`}, {`Sat Feb 10 17:32:01 1996`}, {`Sat Feb 10 17:32:01 1996`}, {`Mon Jun 10 18:32:01 1996`}, {`Sat Feb 10 17:32:01 1996`}, {`Sun Feb 11 17:32:01 1996`}, {`Mon Feb 12 17:32:01 1996`}, {`Tue Feb 13 17:32:01 1996`}, {`Wed Feb 14 17:32:01 1996`}, {`Thu Feb 15 17:32:01 1996`}, {`Fri Feb 16 17:32:01 1996`}, {`Mon Feb 16 17:32:01 0098 BC`}, {`Thu Feb 16 17:32:01 0096`}, {`Tue Feb 16 17:32:01 0596`}, {`Sun Feb 16 17:32:01 1096`}, {`Thu Feb 16 17:32:01 1696`}, {`Tue Feb 16 17:32:01 1796`}, {`Sun Feb 16 17:32:01 1896`}, {`Fri Feb 16 17:32:01 1996`}, {`Thu Feb 16 17:32:01 2096`}, {`Tue Feb 28 17:32:01 1995`}, {`Tue Feb 28 17:32:01 1995`}, {`Wed Mar 01 17:32:01 1995`}, {`Sat Dec 30 17:32:01 1995`}, {`Sun Dec 31 17:32:01 1995`}, {`Mon Jan 01 17:32:01 1996`}, {`Wed Feb 28 17:32:01 1996`}, {`Fri Mar 01 17:32:01 1996`}, {`Mon Dec 30 17:32:01 1996`}, {`Tue Dec 31 17:32:01 1996`}, {`Thu Dec 31 17:32:01 1998`}, {`Fri Jan 01 17:32:01 1999`}, {`Fri Dec 31 17:32:01 1999`}, {`Sat Jan 01 17:32:01 2000`}},
			},
			{
				Statement: `SELECT timestamp with time zone '1996-03-01' - interval '1 second' AS "Feb 29";`,
				Results:   []sql.Row{{`Thu Feb 29 23:59:59 1996 PST`}},
			},
			{
				Statement: `SELECT timestamp with time zone '1999-03-01' - interval '1 second' AS "Feb 28";`,
				Results:   []sql.Row{{`Sun Feb 28 23:59:59 1999 PST`}},
			},
			{
				Statement: `SELECT timestamp with time zone '2000-03-01' - interval '1 second' AS "Feb 29";`,
				Results:   []sql.Row{{`Tue Feb 29 23:59:59 2000 PST`}},
			},
			{
				Statement: `SELECT timestamp with time zone '1999-12-01' + interval '1 month - 1 second' AS "Dec 31";`,
				Results:   []sql.Row{{`Fri Dec 31 23:59:59 1999 PST`}},
			},
			{
				Statement: `SELECT (timestamp with time zone 'today' = (timestamp with time zone 'yesterday' + interval '1 day')) as "True";`,
				Results:   []sql.Row{{true}},
			},
			{
				Statement: `SELECT (timestamp with time zone 'today' = (timestamp with time zone 'tomorrow' - interval '1 day')) as "True";`,
				Results:   []sql.Row{{true}},
			},
			{
				Statement: `SELECT (timestamp with time zone 'tomorrow' = (timestamp with time zone 'yesterday' + interval '2 days')) as "True";`,
				Results:   []sql.Row{{true}},
			},
			{
				Statement: `SELECT (timestamp with time zone 'tomorrow' > 'now') as "True";`,
				Results:   []sql.Row{{true}},
			},
			{
				Statement: `SET TIME ZONE 'CST7CDT,M4.1.0,M10.5.0';`,
			},
			{
				Statement: `SELECT timestamp with time zone '2005-04-02 12:00-07' + interval '1 day' as "Apr 3, 12:00";`,
				Results:   []sql.Row{{`Sun Apr 03 12:00:00 2005 CDT`}},
			},
			{
				Statement: `SELECT timestamp with time zone '2005-04-02 12:00-07' + interval '24 hours' as "Apr 3, 13:00";`,
				Results:   []sql.Row{{`Sun Apr 03 13:00:00 2005 CDT`}},
			},
			{
				Statement: `SELECT timestamp with time zone '2005-04-03 12:00-06' - interval '1 day' as "Apr 2, 12:00";`,
				Results:   []sql.Row{{`Sat Apr 02 12:00:00 2005 CST`}},
			},
			{
				Statement: `SELECT timestamp with time zone '2005-04-03 12:00-06' - interval '24 hours' as "Apr 2, 11:00";`,
				Results:   []sql.Row{{`Sat Apr 02 11:00:00 2005 CST`}},
			},
			{
				Statement: `RESET TIME ZONE;`,
			},
			{
				Statement: `SELECT timestamptz(date '1994-01-01', time '11:00') AS "Jan_01_1994_10am";`,
				Results:   []sql.Row{{`Sat Jan 01 11:00:00 1994 PST`}},
			},
			{
				Statement: `SELECT timestamptz(date '1994-01-01', time '10:00') AS "Jan_01_1994_9am";`,
				Results:   []sql.Row{{`Sat Jan 01 10:00:00 1994 PST`}},
			},
			{
				Statement: `SELECT timestamptz(date '1994-01-01', time with time zone '11:00-8') AS "Jan_01_1994_11am";`,
				Results:   []sql.Row{{`Sat Jan 01 11:00:00 1994 PST`}},
			},
			{
				Statement: `SELECT timestamptz(date '1994-01-01', time with time zone '10:00-8') AS "Jan_01_1994_10am";`,
				Results:   []sql.Row{{`Sat Jan 01 10:00:00 1994 PST`}},
			},
			{
				Statement: `SELECT timestamptz(date '1994-01-01', time with time zone '11:00-5') AS "Jan_01_1994_8am";`,
				Results:   []sql.Row{{`Sat Jan 01 08:00:00 1994 PST`}},
			},
			{
				Statement: `SELECT d1 + interval '1 year' AS one_year FROM TIMESTAMPTZ_TBL;`,
				Results:   []sql.Row{{`-infinity`}, {`infinity`}, {`Thu Dec 31 16:00:00 1970 PST`}, {`Tue Feb 10 17:32:01 1998 PST`}, {`Tue Feb 10 17:32:01 1998 PST`}, {`Tue Feb 10 17:32:02 1998 PST`}, {`Tue Feb 10 17:32:01.4 1998 PST`}, {`Tue Feb 10 17:32:01.5 1998 PST`}, {`Tue Feb 10 17:32:01.6 1998 PST`}, {`Fri Jan 02 00:00:00 1998 PST`}, {`Fri Jan 02 03:04:05 1998 PST`}, {`Tue Feb 10 17:32:01 1998 PST`}, {`Tue Feb 10 17:32:01 1998 PST`}, {`Tue Feb 10 17:32:01 1998 PST`}, {`Tue Feb 10 17:32:01 1998 PST`}, {`Wed Jun 10 17:32:01 1998 PDT`}, {`Sun Sep 22 18:19:20 2002 PDT`}, {`Thu Mar 15 08:14:01 2001 PST`}, {`Thu Mar 15 04:14:02 2001 PST`}, {`Thu Mar 15 02:14:03 2001 PST`}, {`Thu Mar 15 03:14:04 2001 PST`}, {`Thu Mar 15 01:14:05 2001 PST`}, {`Tue Feb 10 17:32:01 1998 PST`}, {`Tue Feb 10 17:32:01 1998 PST`}, {`Tue Feb 10 17:32:00 1998 PST`}, {`Tue Feb 10 17:32:01 1998 PST`}, {`Tue Feb 10 17:32:01 1998 PST`}, {`Tue Feb 10 17:32:01 1998 PST`}, {`Tue Feb 10 17:32:01 1998 PST`}, {`Tue Feb 10 17:32:01 1998 PST`}, {`Tue Feb 10 09:32:01 1998 PST`}, {`Tue Feb 10 09:32:01 1998 PST`}, {`Tue Feb 10 09:32:01 1998 PST`}, {`Tue Feb 10 14:32:01 1998 PST`}, {`Fri Jul 10 14:32:01 1998 PDT`}, {`Wed Jun 10 18:32:01 1998 PDT`}, {`Tue Feb 10 17:32:01 1998 PST`}, {`Wed Feb 11 17:32:01 1998 PST`}, {`Thu Feb 12 17:32:01 1998 PST`}, {`Fri Feb 13 17:32:01 1998 PST`}, {`Sat Feb 14 17:32:01 1998 PST`}, {`Sun Feb 15 17:32:01 1998 PST`}, {`Mon Feb 16 17:32:01 1998 PST`}, {`Thu Feb 16 17:32:01 0096 PST BC`}, {`Sun Feb 16 17:32:01 0098 PST`}, {`Fri Feb 16 17:32:01 0598 PST`}, {`Wed Feb 16 17:32:01 1098 PST`}, {`Sun Feb 16 17:32:01 1698 PST`}, {`Fri Feb 16 17:32:01 1798 PST`}, {`Wed Feb 16 17:32:01 1898 PST`}, {`Mon Feb 16 17:32:01 1998 PST`}, {`Sun Feb 16 17:32:01 2098 PST`}, {`Fri Feb 28 17:32:01 1997 PST`}, {`Fri Feb 28 17:32:01 1997 PST`}, {`Sat Mar 01 17:32:01 1997 PST`}, {`Tue Dec 30 17:32:01 1997 PST`}, {`Wed Dec 31 17:32:01 1997 PST`}, {`Thu Jan 01 17:32:01 1998 PST`}, {`Sat Feb 28 17:32:01 1998 PST`}, {`Sun Mar 01 17:32:01 1998 PST`}, {`Wed Dec 30 17:32:01 1998 PST`}, {`Thu Dec 31 17:32:01 1998 PST`}, {`Sun Dec 31 17:32:01 2000 PST`}, {`Mon Jan 01 17:32:01 2001 PST`}, {`Mon Dec 31 17:32:01 2001 PST`}, {`Tue Jan 01 17:32:01 2002 PST`}},
			},
			{
				Statement: `SELECT d1 - interval '1 year' AS one_year FROM TIMESTAMPTZ_TBL;`,
				Results:   []sql.Row{{`-infinity`}, {`infinity`}, {`Tue Dec 31 16:00:00 1968 PST`}, {`Sat Feb 10 17:32:01 1996 PST`}, {`Sat Feb 10 17:32:01 1996 PST`}, {`Sat Feb 10 17:32:02 1996 PST`}, {`Sat Feb 10 17:32:01.4 1996 PST`}, {`Sat Feb 10 17:32:01.5 1996 PST`}, {`Sat Feb 10 17:32:01.6 1996 PST`}, {`Tue Jan 02 00:00:00 1996 PST`}, {`Tue Jan 02 03:04:05 1996 PST`}, {`Sat Feb 10 17:32:01 1996 PST`}, {`Sat Feb 10 17:32:01 1996 PST`}, {`Sat Feb 10 17:32:01 1996 PST`}, {`Sat Feb 10 17:32:01 1996 PST`}, {`Mon Jun 10 17:32:01 1996 PDT`}, {`Fri Sep 22 18:19:20 2000 PDT`}, {`Mon Mar 15 08:14:01 1999 PST`}, {`Mon Mar 15 04:14:02 1999 PST`}, {`Mon Mar 15 02:14:03 1999 PST`}, {`Mon Mar 15 03:14:04 1999 PST`}, {`Mon Mar 15 01:14:05 1999 PST`}, {`Sat Feb 10 17:32:01 1996 PST`}, {`Sat Feb 10 17:32:01 1996 PST`}, {`Sat Feb 10 17:32:00 1996 PST`}, {`Sat Feb 10 17:32:01 1996 PST`}, {`Sat Feb 10 17:32:01 1996 PST`}, {`Sat Feb 10 17:32:01 1996 PST`}, {`Sat Feb 10 17:32:01 1996 PST`}, {`Sat Feb 10 17:32:01 1996 PST`}, {`Sat Feb 10 09:32:01 1996 PST`}, {`Sat Feb 10 09:32:01 1996 PST`}, {`Sat Feb 10 09:32:01 1996 PST`}, {`Sat Feb 10 14:32:01 1996 PST`}, {`Wed Jul 10 14:32:01 1996 PDT`}, {`Mon Jun 10 18:32:01 1996 PDT`}, {`Sat Feb 10 17:32:01 1996 PST`}, {`Sun Feb 11 17:32:01 1996 PST`}, {`Mon Feb 12 17:32:01 1996 PST`}, {`Tue Feb 13 17:32:01 1996 PST`}, {`Wed Feb 14 17:32:01 1996 PST`}, {`Thu Feb 15 17:32:01 1996 PST`}, {`Fri Feb 16 17:32:01 1996 PST`}, {`Mon Feb 16 17:32:01 0098 PST BC`}, {`Thu Feb 16 17:32:01 0096 PST`}, {`Tue Feb 16 17:32:01 0596 PST`}, {`Sun Feb 16 17:32:01 1096 PST`}, {`Thu Feb 16 17:32:01 1696 PST`}, {`Tue Feb 16 17:32:01 1796 PST`}, {`Sun Feb 16 17:32:01 1896 PST`}, {`Fri Feb 16 17:32:01 1996 PST`}, {`Thu Feb 16 17:32:01 2096 PST`}, {`Tue Feb 28 17:32:01 1995 PST`}, {`Tue Feb 28 17:32:01 1995 PST`}, {`Wed Mar 01 17:32:01 1995 PST`}, {`Sat Dec 30 17:32:01 1995 PST`}, {`Sun Dec 31 17:32:01 1995 PST`}, {`Mon Jan 01 17:32:01 1996 PST`}, {`Wed Feb 28 17:32:01 1996 PST`}, {`Fri Mar 01 17:32:01 1996 PST`}, {`Mon Dec 30 17:32:01 1996 PST`}, {`Tue Dec 31 17:32:01 1996 PST`}, {`Thu Dec 31 17:32:01 1998 PST`}, {`Fri Jan 01 17:32:01 1999 PST`}, {`Fri Dec 31 17:32:01 1999 PST`}, {`Sat Jan 01 17:32:01 2000 PST`}},
			},
			{
				Statement: `SELECT CAST(time '01:02' AS interval) AS "+01:02";`,
				Results:   []sql.Row{{`@ 1 hour 2 mins`}},
			},
			{
				Statement: `SELECT CAST(interval '02:03' AS time) AS "02:03:00";`,
				Results:   []sql.Row{{`02:03:00`}},
			},
			{
				Statement: `SELECT time '01:30' + interval '02:01' AS "03:31:00";`,
				Results:   []sql.Row{{`03:31:00`}},
			},
			{
				Statement: `SELECT time '01:30' - interval '02:01' AS "23:29:00";`,
				Results:   []sql.Row{{`23:29:00`}},
			},
			{
				Statement: `SELECT time '02:30' + interval '36:01' AS "14:31:00";`,
				Results:   []sql.Row{{`14:31:00`}},
			},
			{
				Statement: `SELECT time '03:30' + interval '1 month 04:01' AS "07:31:00";`,
				Results:   []sql.Row{{`07:31:00`}},
			},
			{
				Statement:   `SELECT CAST(time with time zone '01:02-08' AS interval) AS "+00:01";`,
				ErrorString: `cannot cast type time with time zone to interval`,
			},
			{
				Statement:   `SELECT CAST(interval '02:03' AS time with time zone) AS "02:03:00-08";`,
				ErrorString: `cannot cast type interval to time with time zone`,
			},
			{
				Statement: `SELECT time with time zone '01:30-08' - interval '02:01' AS "23:29:00-08";`,
				Results:   []sql.Row{{`23:29:00-08`}},
			},
			{
				Statement: `SELECT time with time zone '02:30-08' + interval '36:01' AS "14:31:00-08";`,
				Results:   []sql.Row{{`14:31:00-08`}},
			},
			{
				Statement: `SELECT CAST(CAST(date 'today' + time with time zone '05:30'
            + interval '02:01' AS time with time zone) AS time) AS "07:31:00";`,
				Results: []sql.Row{{`07:31:00`}},
			},
			{
				Statement: `SELECT CAST(cast(date 'today' + time with time zone '03:30'
  + interval '1 month 04:01' as timestamp without time zone) AS time) AS "07:31:00";`,
				Results: []sql.Row{{`07:31:00`}},
			},
			{
				Statement: `SELECT t.d1 AS t, i.f1 AS i, t.d1 + i.f1 AS "add", t.d1 - i.f1 AS "subtract"
  FROM TIMESTAMP_TBL t, INTERVAL_TBL i
  WHERE t.d1 BETWEEN '1990-01-01' AND '2001-01-01'
    AND i.f1 BETWEEN '00:00' AND '23:00'
  ORDER BY 1,2;`,
				Results: []sql.Row{{`Wed Feb 28 17:32:01 1996`, `@ 1 min`, `Wed Feb 28 17:33:01 1996`, `Wed Feb 28 17:31:01 1996`}, {`Wed Feb 28 17:32:01 1996`, `@ 5 hours`, `Wed Feb 28 22:32:01 1996`, `Wed Feb 28 12:32:01 1996`}, {`Thu Feb 29 17:32:01 1996`, `@ 1 min`, `Thu Feb 29 17:33:01 1996`, `Thu Feb 29 17:31:01 1996`}, {`Thu Feb 29 17:32:01 1996`, `@ 5 hours`, `Thu Feb 29 22:32:01 1996`, `Thu Feb 29 12:32:01 1996`}, {`Fri Mar 01 17:32:01 1996`, `@ 1 min`, `Fri Mar 01 17:33:01 1996`, `Fri Mar 01 17:31:01 1996`}, {`Fri Mar 01 17:32:01 1996`, `@ 5 hours`, `Fri Mar 01 22:32:01 1996`, `Fri Mar 01 12:32:01 1996`}, {`Mon Dec 30 17:32:01 1996`, `@ 1 min`, `Mon Dec 30 17:33:01 1996`, `Mon Dec 30 17:31:01 1996`}, {`Mon Dec 30 17:32:01 1996`, `@ 5 hours`, `Mon Dec 30 22:32:01 1996`, `Mon Dec 30 12:32:01 1996`}, {`Tue Dec 31 17:32:01 1996`, `@ 1 min`, `Tue Dec 31 17:33:01 1996`, `Tue Dec 31 17:31:01 1996`}, {`Tue Dec 31 17:32:01 1996`, `@ 5 hours`, `Tue Dec 31 22:32:01 1996`, `Tue Dec 31 12:32:01 1996`}, {`Wed Jan 01 17:32:01 1997`, `@ 1 min`, `Wed Jan 01 17:33:01 1997`, `Wed Jan 01 17:31:01 1997`}, {`Wed Jan 01 17:32:01 1997`, `@ 5 hours`, `Wed Jan 01 22:32:01 1997`, `Wed Jan 01 12:32:01 1997`}, {`Thu Jan 02 00:00:00 1997`, `@ 1 min`, `Thu Jan 02 00:01:00 1997`, `Wed Jan 01 23:59:00 1997`}, {`Thu Jan 02 00:00:00 1997`, `@ 5 hours`, `Thu Jan 02 05:00:00 1997`, `Wed Jan 01 19:00:00 1997`}, {`Thu Jan 02 03:04:05 1997`, `@ 1 min`, `Thu Jan 02 03:05:05 1997`, `Thu Jan 02 03:03:05 1997`}, {`Thu Jan 02 03:04:05 1997`, `@ 5 hours`, `Thu Jan 02 08:04:05 1997`, `Wed Jan 01 22:04:05 1997`}, {`Mon Feb 10 17:32:00 1997`, `@ 1 min`, `Mon Feb 10 17:33:00 1997`, `Mon Feb 10 17:31:00 1997`}, {`Mon Feb 10 17:32:00 1997`, `@ 5 hours`, `Mon Feb 10 22:32:00 1997`, `Mon Feb 10 12:32:00 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 1 min`, `Mon Feb 10 17:33:01 1997`, `Mon Feb 10 17:31:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 1 min`, `Mon Feb 10 17:33:01 1997`, `Mon Feb 10 17:31:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 1 min`, `Mon Feb 10 17:33:01 1997`, `Mon Feb 10 17:31:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 1 min`, `Mon Feb 10 17:33:01 1997`, `Mon Feb 10 17:31:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 1 min`, `Mon Feb 10 17:33:01 1997`, `Mon Feb 10 17:31:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 1 min`, `Mon Feb 10 17:33:01 1997`, `Mon Feb 10 17:31:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 1 min`, `Mon Feb 10 17:33:01 1997`, `Mon Feb 10 17:31:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 1 min`, `Mon Feb 10 17:33:01 1997`, `Mon Feb 10 17:31:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 1 min`, `Mon Feb 10 17:33:01 1997`, `Mon Feb 10 17:31:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 1 min`, `Mon Feb 10 17:33:01 1997`, `Mon Feb 10 17:31:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 1 min`, `Mon Feb 10 17:33:01 1997`, `Mon Feb 10 17:31:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 1 min`, `Mon Feb 10 17:33:01 1997`, `Mon Feb 10 17:31:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 1 min`, `Mon Feb 10 17:33:01 1997`, `Mon Feb 10 17:31:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 1 min`, `Mon Feb 10 17:33:01 1997`, `Mon Feb 10 17:31:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 1 min`, `Mon Feb 10 17:33:01 1997`, `Mon Feb 10 17:31:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 1 min`, `Mon Feb 10 17:33:01 1997`, `Mon Feb 10 17:31:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 1 min`, `Mon Feb 10 17:33:01 1997`, `Mon Feb 10 17:31:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 1 min`, `Mon Feb 10 17:33:01 1997`, `Mon Feb 10 17:31:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 5 hours`, `Mon Feb 10 22:32:01 1997`, `Mon Feb 10 12:32:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 5 hours`, `Mon Feb 10 22:32:01 1997`, `Mon Feb 10 12:32:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 5 hours`, `Mon Feb 10 22:32:01 1997`, `Mon Feb 10 12:32:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 5 hours`, `Mon Feb 10 22:32:01 1997`, `Mon Feb 10 12:32:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 5 hours`, `Mon Feb 10 22:32:01 1997`, `Mon Feb 10 12:32:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 5 hours`, `Mon Feb 10 22:32:01 1997`, `Mon Feb 10 12:32:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 5 hours`, `Mon Feb 10 22:32:01 1997`, `Mon Feb 10 12:32:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 5 hours`, `Mon Feb 10 22:32:01 1997`, `Mon Feb 10 12:32:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 5 hours`, `Mon Feb 10 22:32:01 1997`, `Mon Feb 10 12:32:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 5 hours`, `Mon Feb 10 22:32:01 1997`, `Mon Feb 10 12:32:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 5 hours`, `Mon Feb 10 22:32:01 1997`, `Mon Feb 10 12:32:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 5 hours`, `Mon Feb 10 22:32:01 1997`, `Mon Feb 10 12:32:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 5 hours`, `Mon Feb 10 22:32:01 1997`, `Mon Feb 10 12:32:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 5 hours`, `Mon Feb 10 22:32:01 1997`, `Mon Feb 10 12:32:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 5 hours`, `Mon Feb 10 22:32:01 1997`, `Mon Feb 10 12:32:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 5 hours`, `Mon Feb 10 22:32:01 1997`, `Mon Feb 10 12:32:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 5 hours`, `Mon Feb 10 22:32:01 1997`, `Mon Feb 10 12:32:01 1997`}, {`Mon Feb 10 17:32:01 1997`, `@ 5 hours`, `Mon Feb 10 22:32:01 1997`, `Mon Feb 10 12:32:01 1997`}, {`Mon Feb 10 17:32:01.4 1997`, `@ 1 min`, `Mon Feb 10 17:33:01.4 1997`, `Mon Feb 10 17:31:01.4 1997`}, {`Mon Feb 10 17:32:01.4 1997`, `@ 5 hours`, `Mon Feb 10 22:32:01.4 1997`, `Mon Feb 10 12:32:01.4 1997`}, {`Mon Feb 10 17:32:01.5 1997`, `@ 1 min`, `Mon Feb 10 17:33:01.5 1997`, `Mon Feb 10 17:31:01.5 1997`}, {`Mon Feb 10 17:32:01.5 1997`, `@ 5 hours`, `Mon Feb 10 22:32:01.5 1997`, `Mon Feb 10 12:32:01.5 1997`}, {`Mon Feb 10 17:32:01.6 1997`, `@ 1 min`, `Mon Feb 10 17:33:01.6 1997`, `Mon Feb 10 17:31:01.6 1997`}, {`Mon Feb 10 17:32:01.6 1997`, `@ 5 hours`, `Mon Feb 10 22:32:01.6 1997`, `Mon Feb 10 12:32:01.6 1997`}, {`Mon Feb 10 17:32:02 1997`, `@ 1 min`, `Mon Feb 10 17:33:02 1997`, `Mon Feb 10 17:31:02 1997`}, {`Mon Feb 10 17:32:02 1997`, `@ 5 hours`, `Mon Feb 10 22:32:02 1997`, `Mon Feb 10 12:32:02 1997`}, {`Tue Feb 11 17:32:01 1997`, `@ 1 min`, `Tue Feb 11 17:33:01 1997`, `Tue Feb 11 17:31:01 1997`}, {`Tue Feb 11 17:32:01 1997`, `@ 5 hours`, `Tue Feb 11 22:32:01 1997`, `Tue Feb 11 12:32:01 1997`}, {`Wed Feb 12 17:32:01 1997`, `@ 1 min`, `Wed Feb 12 17:33:01 1997`, `Wed Feb 12 17:31:01 1997`}, {`Wed Feb 12 17:32:01 1997`, `@ 5 hours`, `Wed Feb 12 22:32:01 1997`, `Wed Feb 12 12:32:01 1997`}, {`Thu Feb 13 17:32:01 1997`, `@ 1 min`, `Thu Feb 13 17:33:01 1997`, `Thu Feb 13 17:31:01 1997`}, {`Thu Feb 13 17:32:01 1997`, `@ 5 hours`, `Thu Feb 13 22:32:01 1997`, `Thu Feb 13 12:32:01 1997`}, {`Fri Feb 14 17:32:01 1997`, `@ 1 min`, `Fri Feb 14 17:33:01 1997`, `Fri Feb 14 17:31:01 1997`}, {`Fri Feb 14 17:32:01 1997`, `@ 5 hours`, `Fri Feb 14 22:32:01 1997`, `Fri Feb 14 12:32:01 1997`}, {`Sat Feb 15 17:32:01 1997`, `@ 1 min`, `Sat Feb 15 17:33:01 1997`, `Sat Feb 15 17:31:01 1997`}, {`Sat Feb 15 17:32:01 1997`, `@ 5 hours`, `Sat Feb 15 22:32:01 1997`, `Sat Feb 15 12:32:01 1997`}, {`Sun Feb 16 17:32:01 1997`, `@ 1 min`, `Sun Feb 16 17:33:01 1997`, `Sun Feb 16 17:31:01 1997`}, {`Sun Feb 16 17:32:01 1997`, `@ 1 min`, `Sun Feb 16 17:33:01 1997`, `Sun Feb 16 17:31:01 1997`}, {`Sun Feb 16 17:32:01 1997`, `@ 5 hours`, `Sun Feb 16 22:32:01 1997`, `Sun Feb 16 12:32:01 1997`}, {`Sun Feb 16 17:32:01 1997`, `@ 5 hours`, `Sun Feb 16 22:32:01 1997`, `Sun Feb 16 12:32:01 1997`}, {`Fri Feb 28 17:32:01 1997`, `@ 1 min`, `Fri Feb 28 17:33:01 1997`, `Fri Feb 28 17:31:01 1997`}, {`Fri Feb 28 17:32:01 1997`, `@ 5 hours`, `Fri Feb 28 22:32:01 1997`, `Fri Feb 28 12:32:01 1997`}, {`Sat Mar 01 17:32:01 1997`, `@ 1 min`, `Sat Mar 01 17:33:01 1997`, `Sat Mar 01 17:31:01 1997`}, {`Sat Mar 01 17:32:01 1997`, `@ 5 hours`, `Sat Mar 01 22:32:01 1997`, `Sat Mar 01 12:32:01 1997`}, {`Tue Jun 10 17:32:01 1997`, `@ 1 min`, `Tue Jun 10 17:33:01 1997`, `Tue Jun 10 17:31:01 1997`}, {`Tue Jun 10 17:32:01 1997`, `@ 5 hours`, `Tue Jun 10 22:32:01 1997`, `Tue Jun 10 12:32:01 1997`}, {`Tue Jun 10 18:32:01 1997`, `@ 1 min`, `Tue Jun 10 18:33:01 1997`, `Tue Jun 10 18:31:01 1997`}, {`Tue Jun 10 18:32:01 1997`, `@ 5 hours`, `Tue Jun 10 23:32:01 1997`, `Tue Jun 10 13:32:01 1997`}, {`Tue Dec 30 17:32:01 1997`, `@ 1 min`, `Tue Dec 30 17:33:01 1997`, `Tue Dec 30 17:31:01 1997`}, {`Tue Dec 30 17:32:01 1997`, `@ 5 hours`, `Tue Dec 30 22:32:01 1997`, `Tue Dec 30 12:32:01 1997`}, {`Wed Dec 31 17:32:01 1997`, `@ 1 min`, `Wed Dec 31 17:33:01 1997`, `Wed Dec 31 17:31:01 1997`}, {`Wed Dec 31 17:32:01 1997`, `@ 5 hours`, `Wed Dec 31 22:32:01 1997`, `Wed Dec 31 12:32:01 1997`}, {`Fri Dec 31 17:32:01 1999`, `@ 1 min`, `Fri Dec 31 17:33:01 1999`, `Fri Dec 31 17:31:01 1999`}, {`Fri Dec 31 17:32:01 1999`, `@ 5 hours`, `Fri Dec 31 22:32:01 1999`, `Fri Dec 31 12:32:01 1999`}, {`Sat Jan 01 17:32:01 2000`, `@ 1 min`, `Sat Jan 01 17:33:01 2000`, `Sat Jan 01 17:31:01 2000`}, {`Sat Jan 01 17:32:01 2000`, `@ 5 hours`, `Sat Jan 01 22:32:01 2000`, `Sat Jan 01 12:32:01 2000`}, {`Wed Mar 15 02:14:05 2000`, `@ 1 min`, `Wed Mar 15 02:15:05 2000`, `Wed Mar 15 02:13:05 2000`}, {`Wed Mar 15 02:14:05 2000`, `@ 5 hours`, `Wed Mar 15 07:14:05 2000`, `Tue Mar 14 21:14:05 2000`}, {`Wed Mar 15 03:14:04 2000`, `@ 1 min`, `Wed Mar 15 03:15:04 2000`, `Wed Mar 15 03:13:04 2000`}, {`Wed Mar 15 03:14:04 2000`, `@ 5 hours`, `Wed Mar 15 08:14:04 2000`, `Tue Mar 14 22:14:04 2000`}, {`Wed Mar 15 08:14:01 2000`, `@ 1 min`, `Wed Mar 15 08:15:01 2000`, `Wed Mar 15 08:13:01 2000`}, {`Wed Mar 15 08:14:01 2000`, `@ 5 hours`, `Wed Mar 15 13:14:01 2000`, `Wed Mar 15 03:14:01 2000`}, {`Wed Mar 15 12:14:03 2000`, `@ 1 min`, `Wed Mar 15 12:15:03 2000`, `Wed Mar 15 12:13:03 2000`}, {`Wed Mar 15 12:14:03 2000`, `@ 5 hours`, `Wed Mar 15 17:14:03 2000`, `Wed Mar 15 07:14:03 2000`}, {`Wed Mar 15 13:14:02 2000`, `@ 1 min`, `Wed Mar 15 13:15:02 2000`, `Wed Mar 15 13:13:02 2000`}, {`Wed Mar 15 13:14:02 2000`, `@ 5 hours`, `Wed Mar 15 18:14:02 2000`, `Wed Mar 15 08:14:02 2000`}, {`Sun Dec 31 17:32:01 2000`, `@ 1 min`, `Sun Dec 31 17:33:01 2000`, `Sun Dec 31 17:31:01 2000`}, {`Sun Dec 31 17:32:01 2000`, `@ 5 hours`, `Sun Dec 31 22:32:01 2000`, `Sun Dec 31 12:32:01 2000`}},
			},
			{
				Statement: `SELECT t.f1 AS t, i.f1 AS i, t.f1 + i.f1 AS "add", t.f1 - i.f1 AS "subtract"
  FROM TIME_TBL t, INTERVAL_TBL i
  ORDER BY 1,2;`,
				Results: []sql.Row{{`00:00:00`, `@ 14 secs ago`, `23:59:46`, `00:00:14`}, {`00:00:00`, `@ 1 min`, `00:01:00`, `23:59:00`}, {`00:00:00`, `@ 5 hours`, `05:00:00`, `19:00:00`}, {`00:00:00`, `@ 1 day 2 hours 3 mins 4 secs`, `02:03:04`, `21:56:56`}, {`00:00:00`, `@ 10 days`, `00:00:00`, `00:00:00`}, {`00:00:00`, `@ 3 mons`, `00:00:00`, `00:00:00`}, {`00:00:00`, `@ 5 mons`, `00:00:00`, `00:00:00`}, {`00:00:00`, `@ 5 mons 12 hours`, `12:00:00`, `12:00:00`}, {`00:00:00`, `@ 6 years`, `00:00:00`, `00:00:00`}, {`00:00:00`, `@ 34 years`, `00:00:00`, `00:00:00`}, {`01:00:00`, `@ 14 secs ago`, `00:59:46`, `01:00:14`}, {`01:00:00`, `@ 1 min`, `01:01:00`, `00:59:00`}, {`01:00:00`, `@ 5 hours`, `06:00:00`, `20:00:00`}, {`01:00:00`, `@ 1 day 2 hours 3 mins 4 secs`, `03:03:04`, `22:56:56`}, {`01:00:00`, `@ 10 days`, `01:00:00`, `01:00:00`}, {`01:00:00`, `@ 3 mons`, `01:00:00`, `01:00:00`}, {`01:00:00`, `@ 5 mons`, `01:00:00`, `01:00:00`}, {`01:00:00`, `@ 5 mons 12 hours`, `13:00:00`, `13:00:00`}, {`01:00:00`, `@ 6 years`, `01:00:00`, `01:00:00`}, {`01:00:00`, `@ 34 years`, `01:00:00`, `01:00:00`}, {`02:03:00`, `@ 14 secs ago`, `02:02:46`, `02:03:14`}, {`02:03:00`, `@ 1 min`, `02:04:00`, `02:02:00`}, {`02:03:00`, `@ 5 hours`, `07:03:00`, `21:03:00`}, {`02:03:00`, `@ 1 day 2 hours 3 mins 4 secs`, `04:06:04`, `23:59:56`}, {`02:03:00`, `@ 10 days`, `02:03:00`, `02:03:00`}, {`02:03:00`, `@ 3 mons`, `02:03:00`, `02:03:00`}, {`02:03:00`, `@ 5 mons`, `02:03:00`, `02:03:00`}, {`02:03:00`, `@ 5 mons 12 hours`, `14:03:00`, `14:03:00`}, {`02:03:00`, `@ 6 years`, `02:03:00`, `02:03:00`}, {`02:03:00`, `@ 34 years`, `02:03:00`, `02:03:00`}, {`11:59:00`, `@ 14 secs ago`, `11:58:46`, `11:59:14`}, {`11:59:00`, `@ 1 min`, `12:00:00`, `11:58:00`}, {`11:59:00`, `@ 5 hours`, `16:59:00`, `06:59:00`}, {`11:59:00`, `@ 1 day 2 hours 3 mins 4 secs`, `14:02:04`, `09:55:56`}, {`11:59:00`, `@ 10 days`, `11:59:00`, `11:59:00`}, {`11:59:00`, `@ 3 mons`, `11:59:00`, `11:59:00`}, {`11:59:00`, `@ 5 mons`, `11:59:00`, `11:59:00`}, {`11:59:00`, `@ 5 mons 12 hours`, `23:59:00`, `23:59:00`}, {`11:59:00`, `@ 6 years`, `11:59:00`, `11:59:00`}, {`11:59:00`, `@ 34 years`, `11:59:00`, `11:59:00`}, {`12:00:00`, `@ 14 secs ago`, `11:59:46`, `12:00:14`}, {`12:00:00`, `@ 1 min`, `12:01:00`, `11:59:00`}, {`12:00:00`, `@ 5 hours`, `17:00:00`, `07:00:00`}, {`12:00:00`, `@ 1 day 2 hours 3 mins 4 secs`, `14:03:04`, `09:56:56`}, {`12:00:00`, `@ 10 days`, `12:00:00`, `12:00:00`}, {`12:00:00`, `@ 3 mons`, `12:00:00`, `12:00:00`}, {`12:00:00`, `@ 5 mons`, `12:00:00`, `12:00:00`}, {`12:00:00`, `@ 5 mons 12 hours`, `00:00:00`, `00:00:00`}, {`12:00:00`, `@ 6 years`, `12:00:00`, `12:00:00`}, {`12:00:00`, `@ 34 years`, `12:00:00`, `12:00:00`}, {`12:01:00`, `@ 14 secs ago`, `12:00:46`, `12:01:14`}, {`12:01:00`, `@ 1 min`, `12:02:00`, `12:00:00`}, {`12:01:00`, `@ 5 hours`, `17:01:00`, `07:01:00`}, {`12:01:00`, `@ 1 day 2 hours 3 mins 4 secs`, `14:04:04`, `09:57:56`}, {`12:01:00`, `@ 10 days`, `12:01:00`, `12:01:00`}, {`12:01:00`, `@ 3 mons`, `12:01:00`, `12:01:00`}, {`12:01:00`, `@ 5 mons`, `12:01:00`, `12:01:00`}, {`12:01:00`, `@ 5 mons 12 hours`, `00:01:00`, `00:01:00`}, {`12:01:00`, `@ 6 years`, `12:01:00`, `12:01:00`}, {`12:01:00`, `@ 34 years`, `12:01:00`, `12:01:00`}, {`15:36:39`, `@ 14 secs ago`, `15:36:25`, `15:36:53`}, {`15:36:39`, `@ 14 secs ago`, `15:36:25`, `15:36:53`}, {`15:36:39`, `@ 1 min`, `15:37:39`, `15:35:39`}, {`15:36:39`, `@ 1 min`, `15:37:39`, `15:35:39`}, {`15:36:39`, `@ 5 hours`, `20:36:39`, `10:36:39`}, {`15:36:39`, `@ 5 hours`, `20:36:39`, `10:36:39`}, {`15:36:39`, `@ 1 day 2 hours 3 mins 4 secs`, `17:39:43`, `13:33:35`}, {`15:36:39`, `@ 1 day 2 hours 3 mins 4 secs`, `17:39:43`, `13:33:35`}, {`15:36:39`, `@ 10 days`, `15:36:39`, `15:36:39`}, {`15:36:39`, `@ 10 days`, `15:36:39`, `15:36:39`}, {`15:36:39`, `@ 3 mons`, `15:36:39`, `15:36:39`}, {`15:36:39`, `@ 3 mons`, `15:36:39`, `15:36:39`}, {`15:36:39`, `@ 5 mons`, `15:36:39`, `15:36:39`}, {`15:36:39`, `@ 5 mons`, `15:36:39`, `15:36:39`}, {`15:36:39`, `@ 5 mons 12 hours`, `03:36:39`, `03:36:39`}, {`15:36:39`, `@ 5 mons 12 hours`, `03:36:39`, `03:36:39`}, {`15:36:39`, `@ 6 years`, `15:36:39`, `15:36:39`}, {`15:36:39`, `@ 6 years`, `15:36:39`, `15:36:39`}, {`15:36:39`, `@ 34 years`, `15:36:39`, `15:36:39`}, {`15:36:39`, `@ 34 years`, `15:36:39`, `15:36:39`}, {`23:59:00`, `@ 14 secs ago`, `23:58:46`, `23:59:14`}, {`23:59:00`, `@ 1 min`, `00:00:00`, `23:58:00`}, {`23:59:00`, `@ 5 hours`, `04:59:00`, `18:59:00`}, {`23:59:00`, `@ 1 day 2 hours 3 mins 4 secs`, `02:02:04`, `21:55:56`}, {`23:59:00`, `@ 10 days`, `23:59:00`, `23:59:00`}, {`23:59:00`, `@ 3 mons`, `23:59:00`, `23:59:00`}, {`23:59:00`, `@ 5 mons`, `23:59:00`, `23:59:00`}, {`23:59:00`, `@ 5 mons 12 hours`, `11:59:00`, `11:59:00`}, {`23:59:00`, `@ 6 years`, `23:59:00`, `23:59:00`}, {`23:59:00`, `@ 34 years`, `23:59:00`, `23:59:00`}, {`23:59:59.99`, `@ 14 secs ago`, `23:59:45.99`, `00:00:13.99`}, {`23:59:59.99`, `@ 1 min`, `00:00:59.99`, `23:58:59.99`}, {`23:59:59.99`, `@ 5 hours`, `04:59:59.99`, `18:59:59.99`}, {`23:59:59.99`, `@ 1 day 2 hours 3 mins 4 secs`, `02:03:03.99`, `21:56:55.99`}, {`23:59:59.99`, `@ 10 days`, `23:59:59.99`, `23:59:59.99`}, {`23:59:59.99`, `@ 3 mons`, `23:59:59.99`, `23:59:59.99`}, {`23:59:59.99`, `@ 5 mons`, `23:59:59.99`, `23:59:59.99`}, {`23:59:59.99`, `@ 5 mons 12 hours`, `11:59:59.99`, `11:59:59.99`}, {`23:59:59.99`, `@ 6 years`, `23:59:59.99`, `23:59:59.99`}, {`23:59:59.99`, `@ 34 years`, `23:59:59.99`, `23:59:59.99`}},
			},
			{
				Statement: `SELECT t.f1 AS t, i.f1 AS i, t.f1 + i.f1 AS "add", t.f1 - i.f1 AS "subtract"
  FROM TIMETZ_TBL t, INTERVAL_TBL i
  ORDER BY 1,2;`,
				Results: []sql.Row{{`00:01:00-07`, `@ 14 secs ago`, `00:00:46-07`, `00:01:14-07`}, {`00:01:00-07`, `@ 1 min`, `00:02:00-07`, `00:00:00-07`}, {`00:01:00-07`, `@ 5 hours`, `05:01:00-07`, `19:01:00-07`}, {`00:01:00-07`, `@ 1 day 2 hours 3 mins 4 secs`, `02:04:04-07`, `21:57:56-07`}, {`00:01:00-07`, `@ 10 days`, `00:01:00-07`, `00:01:00-07`}, {`00:01:00-07`, `@ 3 mons`, `00:01:00-07`, `00:01:00-07`}, {`00:01:00-07`, `@ 5 mons`, `00:01:00-07`, `00:01:00-07`}, {`00:01:00-07`, `@ 5 mons 12 hours`, `12:01:00-07`, `12:01:00-07`}, {`00:01:00-07`, `@ 6 years`, `00:01:00-07`, `00:01:00-07`}, {`00:01:00-07`, `@ 34 years`, `00:01:00-07`, `00:01:00-07`}, {`01:00:00-07`, `@ 14 secs ago`, `00:59:46-07`, `01:00:14-07`}, {`01:00:00-07`, `@ 1 min`, `01:01:00-07`, `00:59:00-07`}, {`01:00:00-07`, `@ 5 hours`, `06:00:00-07`, `20:00:00-07`}, {`01:00:00-07`, `@ 1 day 2 hours 3 mins 4 secs`, `03:03:04-07`, `22:56:56-07`}, {`01:00:00-07`, `@ 10 days`, `01:00:00-07`, `01:00:00-07`}, {`01:00:00-07`, `@ 3 mons`, `01:00:00-07`, `01:00:00-07`}, {`01:00:00-07`, `@ 5 mons`, `01:00:00-07`, `01:00:00-07`}, {`01:00:00-07`, `@ 5 mons 12 hours`, `13:00:00-07`, `13:00:00-07`}, {`01:00:00-07`, `@ 6 years`, `01:00:00-07`, `01:00:00-07`}, {`01:00:00-07`, `@ 34 years`, `01:00:00-07`, `01:00:00-07`}, {`02:03:00-07`, `@ 14 secs ago`, `02:02:46-07`, `02:03:14-07`}, {`02:03:00-07`, `@ 1 min`, `02:04:00-07`, `02:02:00-07`}, {`02:03:00-07`, `@ 5 hours`, `07:03:00-07`, `21:03:00-07`}, {`02:03:00-07`, `@ 1 day 2 hours 3 mins 4 secs`, `04:06:04-07`, `23:59:56-07`}, {`02:03:00-07`, `@ 10 days`, `02:03:00-07`, `02:03:00-07`}, {`02:03:00-07`, `@ 3 mons`, `02:03:00-07`, `02:03:00-07`}, {`02:03:00-07`, `@ 5 mons`, `02:03:00-07`, `02:03:00-07`}, {`02:03:00-07`, `@ 5 mons 12 hours`, `14:03:00-07`, `14:03:00-07`}, {`02:03:00-07`, `@ 6 years`, `02:03:00-07`, `02:03:00-07`}, {`02:03:00-07`, `@ 34 years`, `02:03:00-07`, `02:03:00-07`}, {`08:08:00-04`, `@ 14 secs ago`, `08:07:46-04`, `08:08:14-04`}, {`08:08:00-04`, `@ 1 min`, `08:09:00-04`, `08:07:00-04`}, {`08:08:00-04`, `@ 5 hours`, `13:08:00-04`, `03:08:00-04`}, {`08:08:00-04`, `@ 1 day 2 hours 3 mins 4 secs`, `10:11:04-04`, `06:04:56-04`}, {`08:08:00-04`, `@ 10 days`, `08:08:00-04`, `08:08:00-04`}, {`08:08:00-04`, `@ 3 mons`, `08:08:00-04`, `08:08:00-04`}, {`08:08:00-04`, `@ 5 mons`, `08:08:00-04`, `08:08:00-04`}, {`08:08:00-04`, `@ 5 mons 12 hours`, `20:08:00-04`, `20:08:00-04`}, {`08:08:00-04`, `@ 6 years`, `08:08:00-04`, `08:08:00-04`}, {`08:08:00-04`, `@ 34 years`, `08:08:00-04`, `08:08:00-04`}, {`07:07:00-08`, `@ 14 secs ago`, `07:06:46-08`, `07:07:14-08`}, {`07:07:00-08`, `@ 1 min`, `07:08:00-08`, `07:06:00-08`}, {`07:07:00-08`, `@ 5 hours`, `12:07:00-08`, `02:07:00-08`}, {`07:07:00-08`, `@ 1 day 2 hours 3 mins 4 secs`, `09:10:04-08`, `05:03:56-08`}, {`07:07:00-08`, `@ 10 days`, `07:07:00-08`, `07:07:00-08`}, {`07:07:00-08`, `@ 3 mons`, `07:07:00-08`, `07:07:00-08`}, {`07:07:00-08`, `@ 5 mons`, `07:07:00-08`, `07:07:00-08`}, {`07:07:00-08`, `@ 5 mons 12 hours`, `19:07:00-08`, `19:07:00-08`}, {`07:07:00-08`, `@ 6 years`, `07:07:00-08`, `07:07:00-08`}, {`07:07:00-08`, `@ 34 years`, `07:07:00-08`, `07:07:00-08`}, {`11:59:00-07`, `@ 14 secs ago`, `11:58:46-07`, `11:59:14-07`}, {`11:59:00-07`, `@ 1 min`, `12:00:00-07`, `11:58:00-07`}, {`11:59:00-07`, `@ 5 hours`, `16:59:00-07`, `06:59:00-07`}, {`11:59:00-07`, `@ 1 day 2 hours 3 mins 4 secs`, `14:02:04-07`, `09:55:56-07`}, {`11:59:00-07`, `@ 10 days`, `11:59:00-07`, `11:59:00-07`}, {`11:59:00-07`, `@ 3 mons`, `11:59:00-07`, `11:59:00-07`}, {`11:59:00-07`, `@ 5 mons`, `11:59:00-07`, `11:59:00-07`}, {`11:59:00-07`, `@ 5 mons 12 hours`, `23:59:00-07`, `23:59:00-07`}, {`11:59:00-07`, `@ 6 years`, `11:59:00-07`, `11:59:00-07`}, {`11:59:00-07`, `@ 34 years`, `11:59:00-07`, `11:59:00-07`}, {`12:00:00-07`, `@ 14 secs ago`, `11:59:46-07`, `12:00:14-07`}, {`12:00:00-07`, `@ 1 min`, `12:01:00-07`, `11:59:00-07`}, {`12:00:00-07`, `@ 5 hours`, `17:00:00-07`, `07:00:00-07`}, {`12:00:00-07`, `@ 1 day 2 hours 3 mins 4 secs`, `14:03:04-07`, `09:56:56-07`}, {`12:00:00-07`, `@ 10 days`, `12:00:00-07`, `12:00:00-07`}, {`12:00:00-07`, `@ 3 mons`, `12:00:00-07`, `12:00:00-07`}, {`12:00:00-07`, `@ 5 mons`, `12:00:00-07`, `12:00:00-07`}, {`12:00:00-07`, `@ 5 mons 12 hours`, `00:00:00-07`, `00:00:00-07`}, {`12:00:00-07`, `@ 6 years`, `12:00:00-07`, `12:00:00-07`}, {`12:00:00-07`, `@ 34 years`, `12:00:00-07`, `12:00:00-07`}, {`12:01:00-07`, `@ 14 secs ago`, `12:00:46-07`, `12:01:14-07`}, {`12:01:00-07`, `@ 1 min`, `12:02:00-07`, `12:00:00-07`}, {`12:01:00-07`, `@ 5 hours`, `17:01:00-07`, `07:01:00-07`}, {`12:01:00-07`, `@ 1 day 2 hours 3 mins 4 secs`, `14:04:04-07`, `09:57:56-07`}, {`12:01:00-07`, `@ 10 days`, `12:01:00-07`, `12:01:00-07`}, {`12:01:00-07`, `@ 3 mons`, `12:01:00-07`, `12:01:00-07`}, {`12:01:00-07`, `@ 5 mons`, `12:01:00-07`, `12:01:00-07`}, {`12:01:00-07`, `@ 5 mons 12 hours`, `00:01:00-07`, `00:01:00-07`}, {`12:01:00-07`, `@ 6 years`, `12:01:00-07`, `12:01:00-07`}, {`12:01:00-07`, `@ 34 years`, `12:01:00-07`, `12:01:00-07`}, {`15:36:39-04`, `@ 14 secs ago`, `15:36:25-04`, `15:36:53-04`}, {`15:36:39-04`, `@ 1 min`, `15:37:39-04`, `15:35:39-04`}, {`15:36:39-04`, `@ 5 hours`, `20:36:39-04`, `10:36:39-04`}, {`15:36:39-04`, `@ 1 day 2 hours 3 mins 4 secs`, `17:39:43-04`, `13:33:35-04`}, {`15:36:39-04`, `@ 10 days`, `15:36:39-04`, `15:36:39-04`}, {`15:36:39-04`, `@ 3 mons`, `15:36:39-04`, `15:36:39-04`}, {`15:36:39-04`, `@ 5 mons`, `15:36:39-04`, `15:36:39-04`}, {`15:36:39-04`, `@ 5 mons 12 hours`, `03:36:39-04`, `03:36:39-04`}, {`15:36:39-04`, `@ 6 years`, `15:36:39-04`, `15:36:39-04`}, {`15:36:39-04`, `@ 34 years`, `15:36:39-04`, `15:36:39-04`}, {`15:36:39-05`, `@ 14 secs ago`, `15:36:25-05`, `15:36:53-05`}, {`15:36:39-05`, `@ 1 min`, `15:37:39-05`, `15:35:39-05`}, {`15:36:39-05`, `@ 5 hours`, `20:36:39-05`, `10:36:39-05`}, {`15:36:39-05`, `@ 1 day 2 hours 3 mins 4 secs`, `17:39:43-05`, `13:33:35-05`}, {`15:36:39-05`, `@ 10 days`, `15:36:39-05`, `15:36:39-05`}, {`15:36:39-05`, `@ 3 mons`, `15:36:39-05`, `15:36:39-05`}, {`15:36:39-05`, `@ 5 mons`, `15:36:39-05`, `15:36:39-05`}, {`15:36:39-05`, `@ 5 mons 12 hours`, `03:36:39-05`, `03:36:39-05`}, {`15:36:39-05`, `@ 6 years`, `15:36:39-05`, `15:36:39-05`}, {`15:36:39-05`, `@ 34 years`, `15:36:39-05`, `15:36:39-05`}, {`23:59:00-07`, `@ 14 secs ago`, `23:58:46-07`, `23:59:14-07`}, {`23:59:00-07`, `@ 1 min`, `00:00:00-07`, `23:58:00-07`}, {`23:59:00-07`, `@ 5 hours`, `04:59:00-07`, `18:59:00-07`}, {`23:59:00-07`, `@ 1 day 2 hours 3 mins 4 secs`, `02:02:04-07`, `21:55:56-07`}, {`23:59:00-07`, `@ 10 days`, `23:59:00-07`, `23:59:00-07`}, {`23:59:00-07`, `@ 3 mons`, `23:59:00-07`, `23:59:00-07`}, {`23:59:00-07`, `@ 5 mons`, `23:59:00-07`, `23:59:00-07`}, {`23:59:00-07`, `@ 5 mons 12 hours`, `11:59:00-07`, `11:59:00-07`}, {`23:59:00-07`, `@ 6 years`, `23:59:00-07`, `23:59:00-07`}, {`23:59:00-07`, `@ 34 years`, `23:59:00-07`, `23:59:00-07`}, {`23:59:59.99-07`, `@ 14 secs ago`, `23:59:45.99-07`, `00:00:13.99-07`}, {`23:59:59.99-07`, `@ 1 min`, `00:00:59.99-07`, `23:58:59.99-07`}, {`23:59:59.99-07`, `@ 5 hours`, `04:59:59.99-07`, `18:59:59.99-07`}, {`23:59:59.99-07`, `@ 1 day 2 hours 3 mins 4 secs`, `02:03:03.99-07`, `21:56:55.99-07`}, {`23:59:59.99-07`, `@ 10 days`, `23:59:59.99-07`, `23:59:59.99-07`}, {`23:59:59.99-07`, `@ 3 mons`, `23:59:59.99-07`, `23:59:59.99-07`}, {`23:59:59.99-07`, `@ 5 mons`, `23:59:59.99-07`, `23:59:59.99-07`}, {`23:59:59.99-07`, `@ 5 mons 12 hours`, `11:59:59.99-07`, `11:59:59.99-07`}, {`23:59:59.99-07`, `@ 6 years`, `23:59:59.99-07`, `23:59:59.99-07`}, {`23:59:59.99-07`, `@ 34 years`, `23:59:59.99-07`, `23:59:59.99-07`}},
			},
			{
				Statement: `SELECT (timestamp with time zone '2000-11-27', timestamp with time zone '2000-11-28')
  OVERLAPS (timestamp with time zone '2000-11-27 12:00', timestamp with time zone '2000-11-30') AS "True";`,
				Results: []sql.Row{{true}},
			},
			{
				Statement: `SELECT (timestamp with time zone '2000-11-26', timestamp with time zone '2000-11-27')
  OVERLAPS (timestamp with time zone '2000-11-27 12:00', timestamp with time zone '2000-11-30') AS "False";`,
				Results: []sql.Row{{false}},
			},
			{
				Statement: `SELECT (timestamp with time zone '2000-11-27', timestamp with time zone '2000-11-28')
  OVERLAPS (timestamp with time zone '2000-11-27 12:00', interval '1 day') AS "True";`,
				Results: []sql.Row{{true}},
			},
			{
				Statement: `SELECT (timestamp with time zone '2000-11-27', interval '12 hours')
  OVERLAPS (timestamp with time zone '2000-11-27 12:00', timestamp with time zone '2000-11-30') AS "False";`,
				Results: []sql.Row{{false}},
			},
			{
				Statement: `SELECT (timestamp with time zone '2000-11-27', interval '12 hours')
  OVERLAPS (timestamp with time zone '2000-11-27', interval '12 hours') AS "True";`,
				Results: []sql.Row{{true}},
			},
			{
				Statement: `SELECT (timestamp with time zone '2000-11-27', interval '12 hours')
  OVERLAPS (timestamp with time zone '2000-11-27 12:00', interval '12 hours') AS "False";`,
				Results: []sql.Row{{false}},
			},
			{
				Statement: `SELECT (timestamp without time zone '2000-11-27', timestamp without time zone '2000-11-28')
  OVERLAPS (timestamp without time zone '2000-11-27 12:00', timestamp without time zone '2000-11-30') AS "True";`,
				Results: []sql.Row{{true}},
			},
			{
				Statement: `SELECT (timestamp without time zone '2000-11-26', timestamp without time zone '2000-11-27')
  OVERLAPS (timestamp without time zone '2000-11-27 12:00', timestamp without time zone '2000-11-30') AS "False";`,
				Results: []sql.Row{{false}},
			},
			{
				Statement: `SELECT (timestamp without time zone '2000-11-27', timestamp without time zone '2000-11-28')
  OVERLAPS (timestamp without time zone '2000-11-27 12:00', interval '1 day') AS "True";`,
				Results: []sql.Row{{true}},
			},
			{
				Statement: `SELECT (timestamp without time zone '2000-11-27', interval '12 hours')
  OVERLAPS (timestamp without time zone '2000-11-27 12:00', timestamp without time zone '2000-11-30') AS "False";`,
				Results: []sql.Row{{false}},
			},
			{
				Statement: `SELECT (timestamp without time zone '2000-11-27', interval '12 hours')
  OVERLAPS (timestamp without time zone '2000-11-27', interval '12 hours') AS "True";`,
				Results: []sql.Row{{true}},
			},
			{
				Statement: `SELECT (timestamp without time zone '2000-11-27', interval '12 hours')
  OVERLAPS (timestamp without time zone '2000-11-27 12:00', interval '12 hours') AS "False";`,
				Results: []sql.Row{{false}},
			},
			{
				Statement: `SELECT (time '00:00', time '01:00')
  OVERLAPS (time '00:30', time '01:30') AS "True";`,
				Results: []sql.Row{{true}},
			},
			{
				Statement: `SELECT (time '00:00', interval '1 hour')
  OVERLAPS (time '00:30', interval '1 hour') AS "True";`,
				Results: []sql.Row{{true}},
			},
			{
				Statement: `SELECT (time '00:00', interval '1 hour')
  OVERLAPS (time '01:30', interval '1 hour') AS "False";`,
				Results: []sql.Row{{false}},
			},
			{
				Statement: `SELECT (time '00:00', interval '1 hour')
  OVERLAPS (time '01:30', interval '1 day') AS "False";`,
				Results: []sql.Row{{false}},
			},
			{
				Statement: `CREATE TABLE TEMP_TIMESTAMP (f1 timestamp with time zone);`,
			},
			{
				Statement: `INSERT INTO TEMP_TIMESTAMP (f1)
  SELECT d1 FROM TIMESTAMP_TBL
  WHERE d1 BETWEEN '13-jun-1957' AND '1-jan-1997'
   OR d1 BETWEEN '1-jan-1999' AND '1-jan-2010';`,
			},
			{
				Statement: `SELECT f1 AS "timestamp"
  FROM TEMP_TIMESTAMP
  ORDER BY "timestamp";`,
				Results: []sql.Row{{`Thu Jan 01 00:00:00 1970 PST`}, {`Wed Feb 28 17:32:01 1996 PST`}, {`Thu Feb 29 17:32:01 1996 PST`}, {`Fri Mar 01 17:32:01 1996 PST`}, {`Mon Dec 30 17:32:01 1996 PST`}, {`Tue Dec 31 17:32:01 1996 PST`}, {`Fri Dec 31 17:32:01 1999 PST`}, {`Sat Jan 01 17:32:01 2000 PST`}, {`Wed Mar 15 02:14:05 2000 PST`}, {`Wed Mar 15 03:14:04 2000 PST`}, {`Wed Mar 15 08:14:01 2000 PST`}, {`Wed Mar 15 12:14:03 2000 PST`}, {`Wed Mar 15 13:14:02 2000 PST`}, {`Sun Dec 31 17:32:01 2000 PST`}, {`Mon Jan 01 17:32:01 2001 PST`}, {`Sat Sep 22 18:19:20 2001 PDT`}},
			},
			{
				Statement: `SELECT d.f1 AS "timestamp", t.f1 AS "interval", d.f1 + t.f1 AS plus
  FROM TEMP_TIMESTAMP d, INTERVAL_TBL t
  ORDER BY plus, "timestamp", "interval";`,
				Results: []sql.Row{{`Thu Jan 01 00:00:00 1970 PST`, `@ 14 secs ago`, `Wed Dec 31 23:59:46 1969 PST`}, {`Thu Jan 01 00:00:00 1970 PST`, `@ 1 min`, `Thu Jan 01 00:01:00 1970 PST`}, {`Thu Jan 01 00:00:00 1970 PST`, `@ 5 hours`, `Thu Jan 01 05:00:00 1970 PST`}, {`Thu Jan 01 00:00:00 1970 PST`, `@ 1 day 2 hours 3 mins 4 secs`, `Fri Jan 02 02:03:04 1970 PST`}, {`Thu Jan 01 00:00:00 1970 PST`, `@ 10 days`, `Sun Jan 11 00:00:00 1970 PST`}, {`Thu Jan 01 00:00:00 1970 PST`, `@ 3 mons`, `Wed Apr 01 00:00:00 1970 PST`}, {`Thu Jan 01 00:00:00 1970 PST`, `@ 5 mons`, `Mon Jun 01 00:00:00 1970 PDT`}, {`Thu Jan 01 00:00:00 1970 PST`, `@ 5 mons 12 hours`, `Mon Jun 01 12:00:00 1970 PDT`}, {`Thu Jan 01 00:00:00 1970 PST`, `@ 6 years`, `Thu Jan 01 00:00:00 1976 PST`}, {`Wed Feb 28 17:32:01 1996 PST`, `@ 14 secs ago`, `Wed Feb 28 17:31:47 1996 PST`}, {`Wed Feb 28 17:32:01 1996 PST`, `@ 1 min`, `Wed Feb 28 17:33:01 1996 PST`}, {`Wed Feb 28 17:32:01 1996 PST`, `@ 5 hours`, `Wed Feb 28 22:32:01 1996 PST`}, {`Thu Feb 29 17:32:01 1996 PST`, `@ 14 secs ago`, `Thu Feb 29 17:31:47 1996 PST`}, {`Thu Feb 29 17:32:01 1996 PST`, `@ 1 min`, `Thu Feb 29 17:33:01 1996 PST`}, {`Wed Feb 28 17:32:01 1996 PST`, `@ 1 day 2 hours 3 mins 4 secs`, `Thu Feb 29 19:35:05 1996 PST`}, {`Thu Feb 29 17:32:01 1996 PST`, `@ 5 hours`, `Thu Feb 29 22:32:01 1996 PST`}, {`Fri Mar 01 17:32:01 1996 PST`, `@ 14 secs ago`, `Fri Mar 01 17:31:47 1996 PST`}, {`Fri Mar 01 17:32:01 1996 PST`, `@ 1 min`, `Fri Mar 01 17:33:01 1996 PST`}, {`Thu Feb 29 17:32:01 1996 PST`, `@ 1 day 2 hours 3 mins 4 secs`, `Fri Mar 01 19:35:05 1996 PST`}, {`Fri Mar 01 17:32:01 1996 PST`, `@ 5 hours`, `Fri Mar 01 22:32:01 1996 PST`}, {`Fri Mar 01 17:32:01 1996 PST`, `@ 1 day 2 hours 3 mins 4 secs`, `Sat Mar 02 19:35:05 1996 PST`}, {`Wed Feb 28 17:32:01 1996 PST`, `@ 10 days`, `Sat Mar 09 17:32:01 1996 PST`}, {`Thu Feb 29 17:32:01 1996 PST`, `@ 10 days`, `Sun Mar 10 17:32:01 1996 PST`}, {`Fri Mar 01 17:32:01 1996 PST`, `@ 10 days`, `Mon Mar 11 17:32:01 1996 PST`}, {`Wed Feb 28 17:32:01 1996 PST`, `@ 3 mons`, `Tue May 28 17:32:01 1996 PDT`}, {`Thu Feb 29 17:32:01 1996 PST`, `@ 3 mons`, `Wed May 29 17:32:01 1996 PDT`}, {`Fri Mar 01 17:32:01 1996 PST`, `@ 3 mons`, `Sat Jun 01 17:32:01 1996 PDT`}, {`Wed Feb 28 17:32:01 1996 PST`, `@ 5 mons`, `Sun Jul 28 17:32:01 1996 PDT`}, {`Wed Feb 28 17:32:01 1996 PST`, `@ 5 mons 12 hours`, `Mon Jul 29 05:32:01 1996 PDT`}, {`Thu Feb 29 17:32:01 1996 PST`, `@ 5 mons`, `Mon Jul 29 17:32:01 1996 PDT`}, {`Thu Feb 29 17:32:01 1996 PST`, `@ 5 mons 12 hours`, `Tue Jul 30 05:32:01 1996 PDT`}, {`Fri Mar 01 17:32:01 1996 PST`, `@ 5 mons`, `Thu Aug 01 17:32:01 1996 PDT`}, {`Fri Mar 01 17:32:01 1996 PST`, `@ 5 mons 12 hours`, `Fri Aug 02 05:32:01 1996 PDT`}, {`Mon Dec 30 17:32:01 1996 PST`, `@ 14 secs ago`, `Mon Dec 30 17:31:47 1996 PST`}, {`Mon Dec 30 17:32:01 1996 PST`, `@ 1 min`, `Mon Dec 30 17:33:01 1996 PST`}, {`Mon Dec 30 17:32:01 1996 PST`, `@ 5 hours`, `Mon Dec 30 22:32:01 1996 PST`}, {`Tue Dec 31 17:32:01 1996 PST`, `@ 14 secs ago`, `Tue Dec 31 17:31:47 1996 PST`}, {`Tue Dec 31 17:32:01 1996 PST`, `@ 1 min`, `Tue Dec 31 17:33:01 1996 PST`}, {`Mon Dec 30 17:32:01 1996 PST`, `@ 1 day 2 hours 3 mins 4 secs`, `Tue Dec 31 19:35:05 1996 PST`}, {`Tue Dec 31 17:32:01 1996 PST`, `@ 5 hours`, `Tue Dec 31 22:32:01 1996 PST`}, {`Tue Dec 31 17:32:01 1996 PST`, `@ 1 day 2 hours 3 mins 4 secs`, `Wed Jan 01 19:35:05 1997 PST`}, {`Mon Dec 30 17:32:01 1996 PST`, `@ 10 days`, `Thu Jan 09 17:32:01 1997 PST`}, {`Tue Dec 31 17:32:01 1996 PST`, `@ 10 days`, `Fri Jan 10 17:32:01 1997 PST`}, {`Mon Dec 30 17:32:01 1996 PST`, `@ 3 mons`, `Sun Mar 30 17:32:01 1997 PST`}, {`Tue Dec 31 17:32:01 1996 PST`, `@ 3 mons`, `Mon Mar 31 17:32:01 1997 PST`}, {`Mon Dec 30 17:32:01 1996 PST`, `@ 5 mons`, `Fri May 30 17:32:01 1997 PDT`}, {`Mon Dec 30 17:32:01 1996 PST`, `@ 5 mons 12 hours`, `Sat May 31 05:32:01 1997 PDT`}, {`Tue Dec 31 17:32:01 1996 PST`, `@ 5 mons`, `Sat May 31 17:32:01 1997 PDT`}, {`Tue Dec 31 17:32:01 1996 PST`, `@ 5 mons 12 hours`, `Sun Jun 01 05:32:01 1997 PDT`}, {`Fri Dec 31 17:32:01 1999 PST`, `@ 14 secs ago`, `Fri Dec 31 17:31:47 1999 PST`}, {`Fri Dec 31 17:32:01 1999 PST`, `@ 1 min`, `Fri Dec 31 17:33:01 1999 PST`}, {`Fri Dec 31 17:32:01 1999 PST`, `@ 5 hours`, `Fri Dec 31 22:32:01 1999 PST`}, {`Sat Jan 01 17:32:01 2000 PST`, `@ 14 secs ago`, `Sat Jan 01 17:31:47 2000 PST`}, {`Sat Jan 01 17:32:01 2000 PST`, `@ 1 min`, `Sat Jan 01 17:33:01 2000 PST`}, {`Fri Dec 31 17:32:01 1999 PST`, `@ 1 day 2 hours 3 mins 4 secs`, `Sat Jan 01 19:35:05 2000 PST`}, {`Sat Jan 01 17:32:01 2000 PST`, `@ 5 hours`, `Sat Jan 01 22:32:01 2000 PST`}, {`Sat Jan 01 17:32:01 2000 PST`, `@ 1 day 2 hours 3 mins 4 secs`, `Sun Jan 02 19:35:05 2000 PST`}, {`Fri Dec 31 17:32:01 1999 PST`, `@ 10 days`, `Mon Jan 10 17:32:01 2000 PST`}, {`Sat Jan 01 17:32:01 2000 PST`, `@ 10 days`, `Tue Jan 11 17:32:01 2000 PST`}, {`Wed Mar 15 02:14:05 2000 PST`, `@ 14 secs ago`, `Wed Mar 15 02:13:51 2000 PST`}, {`Wed Mar 15 02:14:05 2000 PST`, `@ 1 min`, `Wed Mar 15 02:15:05 2000 PST`}, {`Wed Mar 15 03:14:04 2000 PST`, `@ 14 secs ago`, `Wed Mar 15 03:13:50 2000 PST`}, {`Wed Mar 15 03:14:04 2000 PST`, `@ 1 min`, `Wed Mar 15 03:15:04 2000 PST`}, {`Wed Mar 15 02:14:05 2000 PST`, `@ 5 hours`, `Wed Mar 15 07:14:05 2000 PST`}, {`Wed Mar 15 08:14:01 2000 PST`, `@ 14 secs ago`, `Wed Mar 15 08:13:47 2000 PST`}, {`Wed Mar 15 03:14:04 2000 PST`, `@ 5 hours`, `Wed Mar 15 08:14:04 2000 PST`}, {`Wed Mar 15 08:14:01 2000 PST`, `@ 1 min`, `Wed Mar 15 08:15:01 2000 PST`}, {`Wed Mar 15 12:14:03 2000 PST`, `@ 14 secs ago`, `Wed Mar 15 12:13:49 2000 PST`}, {`Wed Mar 15 12:14:03 2000 PST`, `@ 1 min`, `Wed Mar 15 12:15:03 2000 PST`}, {`Wed Mar 15 13:14:02 2000 PST`, `@ 14 secs ago`, `Wed Mar 15 13:13:48 2000 PST`}, {`Wed Mar 15 08:14:01 2000 PST`, `@ 5 hours`, `Wed Mar 15 13:14:01 2000 PST`}, {`Wed Mar 15 13:14:02 2000 PST`, `@ 1 min`, `Wed Mar 15 13:15:02 2000 PST`}, {`Wed Mar 15 12:14:03 2000 PST`, `@ 5 hours`, `Wed Mar 15 17:14:03 2000 PST`}, {`Wed Mar 15 13:14:02 2000 PST`, `@ 5 hours`, `Wed Mar 15 18:14:02 2000 PST`}, {`Wed Mar 15 02:14:05 2000 PST`, `@ 1 day 2 hours 3 mins 4 secs`, `Thu Mar 16 04:17:09 2000 PST`}, {`Wed Mar 15 03:14:04 2000 PST`, `@ 1 day 2 hours 3 mins 4 secs`, `Thu Mar 16 05:17:08 2000 PST`}, {`Wed Mar 15 08:14:01 2000 PST`, `@ 1 day 2 hours 3 mins 4 secs`, `Thu Mar 16 10:17:05 2000 PST`}, {`Wed Mar 15 12:14:03 2000 PST`, `@ 1 day 2 hours 3 mins 4 secs`, `Thu Mar 16 14:17:07 2000 PST`}, {`Wed Mar 15 13:14:02 2000 PST`, `@ 1 day 2 hours 3 mins 4 secs`, `Thu Mar 16 15:17:06 2000 PST`}, {`Wed Mar 15 02:14:05 2000 PST`, `@ 10 days`, `Sat Mar 25 02:14:05 2000 PST`}, {`Wed Mar 15 03:14:04 2000 PST`, `@ 10 days`, `Sat Mar 25 03:14:04 2000 PST`}, {`Wed Mar 15 08:14:01 2000 PST`, `@ 10 days`, `Sat Mar 25 08:14:01 2000 PST`}, {`Wed Mar 15 12:14:03 2000 PST`, `@ 10 days`, `Sat Mar 25 12:14:03 2000 PST`}, {`Wed Mar 15 13:14:02 2000 PST`, `@ 10 days`, `Sat Mar 25 13:14:02 2000 PST`}, {`Fri Dec 31 17:32:01 1999 PST`, `@ 3 mons`, `Fri Mar 31 17:32:01 2000 PST`}, {`Sat Jan 01 17:32:01 2000 PST`, `@ 3 mons`, `Sat Apr 01 17:32:01 2000 PST`}, {`Fri Dec 31 17:32:01 1999 PST`, `@ 5 mons`, `Wed May 31 17:32:01 2000 PDT`}, {`Fri Dec 31 17:32:01 1999 PST`, `@ 5 mons 12 hours`, `Thu Jun 01 05:32:01 2000 PDT`}, {`Sat Jan 01 17:32:01 2000 PST`, `@ 5 mons`, `Thu Jun 01 17:32:01 2000 PDT`}, {`Sat Jan 01 17:32:01 2000 PST`, `@ 5 mons 12 hours`, `Fri Jun 02 05:32:01 2000 PDT`}, {`Wed Mar 15 02:14:05 2000 PST`, `@ 3 mons`, `Thu Jun 15 02:14:05 2000 PDT`}, {`Wed Mar 15 03:14:04 2000 PST`, `@ 3 mons`, `Thu Jun 15 03:14:04 2000 PDT`}, {`Wed Mar 15 08:14:01 2000 PST`, `@ 3 mons`, `Thu Jun 15 08:14:01 2000 PDT`}, {`Wed Mar 15 12:14:03 2000 PST`, `@ 3 mons`, `Thu Jun 15 12:14:03 2000 PDT`}, {`Wed Mar 15 13:14:02 2000 PST`, `@ 3 mons`, `Thu Jun 15 13:14:02 2000 PDT`}, {`Wed Mar 15 02:14:05 2000 PST`, `@ 5 mons`, `Tue Aug 15 02:14:05 2000 PDT`}, {`Wed Mar 15 03:14:04 2000 PST`, `@ 5 mons`, `Tue Aug 15 03:14:04 2000 PDT`}, {`Wed Mar 15 08:14:01 2000 PST`, `@ 5 mons`, `Tue Aug 15 08:14:01 2000 PDT`}, {`Wed Mar 15 12:14:03 2000 PST`, `@ 5 mons`, `Tue Aug 15 12:14:03 2000 PDT`}, {`Wed Mar 15 13:14:02 2000 PST`, `@ 5 mons`, `Tue Aug 15 13:14:02 2000 PDT`}, {`Wed Mar 15 02:14:05 2000 PST`, `@ 5 mons 12 hours`, `Tue Aug 15 14:14:05 2000 PDT`}, {`Wed Mar 15 03:14:04 2000 PST`, `@ 5 mons 12 hours`, `Tue Aug 15 15:14:04 2000 PDT`}, {`Wed Mar 15 08:14:01 2000 PST`, `@ 5 mons 12 hours`, `Tue Aug 15 20:14:01 2000 PDT`}, {`Wed Mar 15 12:14:03 2000 PST`, `@ 5 mons 12 hours`, `Wed Aug 16 00:14:03 2000 PDT`}, {`Wed Mar 15 13:14:02 2000 PST`, `@ 5 mons 12 hours`, `Wed Aug 16 01:14:02 2000 PDT`}, {`Sun Dec 31 17:32:01 2000 PST`, `@ 14 secs ago`, `Sun Dec 31 17:31:47 2000 PST`}, {`Sun Dec 31 17:32:01 2000 PST`, `@ 1 min`, `Sun Dec 31 17:33:01 2000 PST`}, {`Sun Dec 31 17:32:01 2000 PST`, `@ 5 hours`, `Sun Dec 31 22:32:01 2000 PST`}, {`Mon Jan 01 17:32:01 2001 PST`, `@ 14 secs ago`, `Mon Jan 01 17:31:47 2001 PST`}, {`Mon Jan 01 17:32:01 2001 PST`, `@ 1 min`, `Mon Jan 01 17:33:01 2001 PST`}, {`Sun Dec 31 17:32:01 2000 PST`, `@ 1 day 2 hours 3 mins 4 secs`, `Mon Jan 01 19:35:05 2001 PST`}, {`Mon Jan 01 17:32:01 2001 PST`, `@ 5 hours`, `Mon Jan 01 22:32:01 2001 PST`}, {`Mon Jan 01 17:32:01 2001 PST`, `@ 1 day 2 hours 3 mins 4 secs`, `Tue Jan 02 19:35:05 2001 PST`}, {`Sun Dec 31 17:32:01 2000 PST`, `@ 10 days`, `Wed Jan 10 17:32:01 2001 PST`}, {`Mon Jan 01 17:32:01 2001 PST`, `@ 10 days`, `Thu Jan 11 17:32:01 2001 PST`}, {`Sun Dec 31 17:32:01 2000 PST`, `@ 3 mons`, `Sat Mar 31 17:32:01 2001 PST`}, {`Mon Jan 01 17:32:01 2001 PST`, `@ 3 mons`, `Sun Apr 01 17:32:01 2001 PDT`}, {`Sun Dec 31 17:32:01 2000 PST`, `@ 5 mons`, `Thu May 31 17:32:01 2001 PDT`}, {`Sun Dec 31 17:32:01 2000 PST`, `@ 5 mons 12 hours`, `Fri Jun 01 05:32:01 2001 PDT`}, {`Mon Jan 01 17:32:01 2001 PST`, `@ 5 mons`, `Fri Jun 01 17:32:01 2001 PDT`}, {`Mon Jan 01 17:32:01 2001 PST`, `@ 5 mons 12 hours`, `Sat Jun 02 05:32:01 2001 PDT`}, {`Sat Sep 22 18:19:20 2001 PDT`, `@ 14 secs ago`, `Sat Sep 22 18:19:06 2001 PDT`}, {`Sat Sep 22 18:19:20 2001 PDT`, `@ 1 min`, `Sat Sep 22 18:20:20 2001 PDT`}, {`Sat Sep 22 18:19:20 2001 PDT`, `@ 5 hours`, `Sat Sep 22 23:19:20 2001 PDT`}, {`Sat Sep 22 18:19:20 2001 PDT`, `@ 1 day 2 hours 3 mins 4 secs`, `Sun Sep 23 20:22:24 2001 PDT`}, {`Sat Sep 22 18:19:20 2001 PDT`, `@ 10 days`, `Tue Oct 02 18:19:20 2001 PDT`}, {`Sat Sep 22 18:19:20 2001 PDT`, `@ 3 mons`, `Sat Dec 22 18:19:20 2001 PST`}, {`Sat Sep 22 18:19:20 2001 PDT`, `@ 5 mons`, `Fri Feb 22 18:19:20 2002 PST`}, {`Sat Sep 22 18:19:20 2001 PDT`, `@ 5 mons 12 hours`, `Sat Feb 23 06:19:20 2002 PST`}, {`Wed Feb 28 17:32:01 1996 PST`, `@ 6 years`, `Thu Feb 28 17:32:01 2002 PST`}, {`Thu Feb 29 17:32:01 1996 PST`, `@ 6 years`, `Thu Feb 28 17:32:01 2002 PST`}, {`Fri Mar 01 17:32:01 1996 PST`, `@ 6 years`, `Fri Mar 01 17:32:01 2002 PST`}, {`Mon Dec 30 17:32:01 1996 PST`, `@ 6 years`, `Mon Dec 30 17:32:01 2002 PST`}, {`Tue Dec 31 17:32:01 1996 PST`, `@ 6 years`, `Tue Dec 31 17:32:01 2002 PST`}, {`Thu Jan 01 00:00:00 1970 PST`, `@ 34 years`, `Thu Jan 01 00:00:00 2004 PST`}, {`Fri Dec 31 17:32:01 1999 PST`, `@ 6 years`, `Sat Dec 31 17:32:01 2005 PST`}, {`Sat Jan 01 17:32:01 2000 PST`, `@ 6 years`, `Sun Jan 01 17:32:01 2006 PST`}, {`Wed Mar 15 02:14:05 2000 PST`, `@ 6 years`, `Wed Mar 15 02:14:05 2006 PST`}, {`Wed Mar 15 03:14:04 2000 PST`, `@ 6 years`, `Wed Mar 15 03:14:04 2006 PST`}, {`Wed Mar 15 08:14:01 2000 PST`, `@ 6 years`, `Wed Mar 15 08:14:01 2006 PST`}, {`Wed Mar 15 12:14:03 2000 PST`, `@ 6 years`, `Wed Mar 15 12:14:03 2006 PST`}, {`Wed Mar 15 13:14:02 2000 PST`, `@ 6 years`, `Wed Mar 15 13:14:02 2006 PST`}, {`Sun Dec 31 17:32:01 2000 PST`, `@ 6 years`, `Sun Dec 31 17:32:01 2006 PST`}, {`Mon Jan 01 17:32:01 2001 PST`, `@ 6 years`, `Mon Jan 01 17:32:01 2007 PST`}, {`Sat Sep 22 18:19:20 2001 PDT`, `@ 6 years`, `Sat Sep 22 18:19:20 2007 PDT`}, {`Wed Feb 28 17:32:01 1996 PST`, `@ 34 years`, `Thu Feb 28 17:32:01 2030 PST`}, {`Thu Feb 29 17:32:01 1996 PST`, `@ 34 years`, `Thu Feb 28 17:32:01 2030 PST`}, {`Fri Mar 01 17:32:01 1996 PST`, `@ 34 years`, `Fri Mar 01 17:32:01 2030 PST`}, {`Mon Dec 30 17:32:01 1996 PST`, `@ 34 years`, `Mon Dec 30 17:32:01 2030 PST`}, {`Tue Dec 31 17:32:01 1996 PST`, `@ 34 years`, `Tue Dec 31 17:32:01 2030 PST`}, {`Fri Dec 31 17:32:01 1999 PST`, `@ 34 years`, `Sat Dec 31 17:32:01 2033 PST`}, {`Sat Jan 01 17:32:01 2000 PST`, `@ 34 years`, `Sun Jan 01 17:32:01 2034 PST`}, {`Wed Mar 15 02:14:05 2000 PST`, `@ 34 years`, `Wed Mar 15 02:14:05 2034 PDT`}, {`Wed Mar 15 03:14:04 2000 PST`, `@ 34 years`, `Wed Mar 15 03:14:04 2034 PDT`}, {`Wed Mar 15 08:14:01 2000 PST`, `@ 34 years`, `Wed Mar 15 08:14:01 2034 PDT`}, {`Wed Mar 15 12:14:03 2000 PST`, `@ 34 years`, `Wed Mar 15 12:14:03 2034 PDT`}, {`Wed Mar 15 13:14:02 2000 PST`, `@ 34 years`, `Wed Mar 15 13:14:02 2034 PDT`}, {`Sun Dec 31 17:32:01 2000 PST`, `@ 34 years`, `Sun Dec 31 17:32:01 2034 PST`}, {`Mon Jan 01 17:32:01 2001 PST`, `@ 34 years`, `Mon Jan 01 17:32:01 2035 PST`}, {`Sat Sep 22 18:19:20 2001 PDT`, `@ 34 years`, `Sat Sep 22 18:19:20 2035 PDT`}},
			},
			{
				Statement: `SELECT d.f1 AS "timestamp", t.f1 AS "interval", d.f1 - t.f1 AS minus
  FROM TEMP_TIMESTAMP d, INTERVAL_TBL t
  WHERE isfinite(d.f1)
  ORDER BY minus, "timestamp", "interval";`,
				Results: []sql.Row{{`Thu Jan 01 00:00:00 1970 PST`, `@ 34 years`, `Wed Jan 01 00:00:00 1936 PST`}, {`Wed Feb 28 17:32:01 1996 PST`, `@ 34 years`, `Wed Feb 28 17:32:01 1962 PST`}, {`Thu Feb 29 17:32:01 1996 PST`, `@ 34 years`, `Wed Feb 28 17:32:01 1962 PST`}, {`Fri Mar 01 17:32:01 1996 PST`, `@ 34 years`, `Thu Mar 01 17:32:01 1962 PST`}, {`Mon Dec 30 17:32:01 1996 PST`, `@ 34 years`, `Sun Dec 30 17:32:01 1962 PST`}, {`Tue Dec 31 17:32:01 1996 PST`, `@ 34 years`, `Mon Dec 31 17:32:01 1962 PST`}, {`Thu Jan 01 00:00:00 1970 PST`, `@ 6 years`, `Wed Jan 01 00:00:00 1964 PST`}, {`Fri Dec 31 17:32:01 1999 PST`, `@ 34 years`, `Fri Dec 31 17:32:01 1965 PST`}, {`Sat Jan 01 17:32:01 2000 PST`, `@ 34 years`, `Sat Jan 01 17:32:01 1966 PST`}, {`Wed Mar 15 02:14:05 2000 PST`, `@ 34 years`, `Tue Mar 15 02:14:05 1966 PST`}, {`Wed Mar 15 03:14:04 2000 PST`, `@ 34 years`, `Tue Mar 15 03:14:04 1966 PST`}, {`Wed Mar 15 08:14:01 2000 PST`, `@ 34 years`, `Tue Mar 15 08:14:01 1966 PST`}, {`Wed Mar 15 12:14:03 2000 PST`, `@ 34 years`, `Tue Mar 15 12:14:03 1966 PST`}, {`Wed Mar 15 13:14:02 2000 PST`, `@ 34 years`, `Tue Mar 15 13:14:02 1966 PST`}, {`Sun Dec 31 17:32:01 2000 PST`, `@ 34 years`, `Sat Dec 31 17:32:01 1966 PST`}, {`Mon Jan 01 17:32:01 2001 PST`, `@ 34 years`, `Sun Jan 01 17:32:01 1967 PST`}, {`Sat Sep 22 18:19:20 2001 PDT`, `@ 34 years`, `Fri Sep 22 18:19:20 1967 PDT`}, {`Thu Jan 01 00:00:00 1970 PST`, `@ 5 mons 12 hours`, `Thu Jul 31 12:00:00 1969 PDT`}, {`Thu Jan 01 00:00:00 1970 PST`, `@ 5 mons`, `Fri Aug 01 00:00:00 1969 PDT`}, {`Thu Jan 01 00:00:00 1970 PST`, `@ 3 mons`, `Wed Oct 01 00:00:00 1969 PDT`}, {`Thu Jan 01 00:00:00 1970 PST`, `@ 10 days`, `Mon Dec 22 00:00:00 1969 PST`}, {`Thu Jan 01 00:00:00 1970 PST`, `@ 1 day 2 hours 3 mins 4 secs`, `Tue Dec 30 21:56:56 1969 PST`}, {`Thu Jan 01 00:00:00 1970 PST`, `@ 5 hours`, `Wed Dec 31 19:00:00 1969 PST`}, {`Thu Jan 01 00:00:00 1970 PST`, `@ 1 min`, `Wed Dec 31 23:59:00 1969 PST`}, {`Thu Jan 01 00:00:00 1970 PST`, `@ 14 secs ago`, `Thu Jan 01 00:00:14 1970 PST`}, {`Wed Feb 28 17:32:01 1996 PST`, `@ 6 years`, `Wed Feb 28 17:32:01 1990 PST`}, {`Thu Feb 29 17:32:01 1996 PST`, `@ 6 years`, `Wed Feb 28 17:32:01 1990 PST`}, {`Fri Mar 01 17:32:01 1996 PST`, `@ 6 years`, `Thu Mar 01 17:32:01 1990 PST`}, {`Mon Dec 30 17:32:01 1996 PST`, `@ 6 years`, `Sun Dec 30 17:32:01 1990 PST`}, {`Tue Dec 31 17:32:01 1996 PST`, `@ 6 years`, `Mon Dec 31 17:32:01 1990 PST`}, {`Fri Dec 31 17:32:01 1999 PST`, `@ 6 years`, `Fri Dec 31 17:32:01 1993 PST`}, {`Sat Jan 01 17:32:01 2000 PST`, `@ 6 years`, `Sat Jan 01 17:32:01 1994 PST`}, {`Wed Mar 15 02:14:05 2000 PST`, `@ 6 years`, `Tue Mar 15 02:14:05 1994 PST`}, {`Wed Mar 15 03:14:04 2000 PST`, `@ 6 years`, `Tue Mar 15 03:14:04 1994 PST`}, {`Wed Mar 15 08:14:01 2000 PST`, `@ 6 years`, `Tue Mar 15 08:14:01 1994 PST`}, {`Wed Mar 15 12:14:03 2000 PST`, `@ 6 years`, `Tue Mar 15 12:14:03 1994 PST`}, {`Wed Mar 15 13:14:02 2000 PST`, `@ 6 years`, `Tue Mar 15 13:14:02 1994 PST`}, {`Sun Dec 31 17:32:01 2000 PST`, `@ 6 years`, `Sat Dec 31 17:32:01 1994 PST`}, {`Mon Jan 01 17:32:01 2001 PST`, `@ 6 years`, `Sun Jan 01 17:32:01 1995 PST`}, {`Sat Sep 22 18:19:20 2001 PDT`, `@ 6 years`, `Fri Sep 22 18:19:20 1995 PDT`}, {`Wed Feb 28 17:32:01 1996 PST`, `@ 5 mons 12 hours`, `Thu Sep 28 05:32:01 1995 PDT`}, {`Wed Feb 28 17:32:01 1996 PST`, `@ 5 mons`, `Thu Sep 28 17:32:01 1995 PDT`}, {`Thu Feb 29 17:32:01 1996 PST`, `@ 5 mons 12 hours`, `Fri Sep 29 05:32:01 1995 PDT`}, {`Thu Feb 29 17:32:01 1996 PST`, `@ 5 mons`, `Fri Sep 29 17:32:01 1995 PDT`}, {`Fri Mar 01 17:32:01 1996 PST`, `@ 5 mons 12 hours`, `Sun Oct 01 05:32:01 1995 PDT`}, {`Fri Mar 01 17:32:01 1996 PST`, `@ 5 mons`, `Sun Oct 01 17:32:01 1995 PDT`}, {`Wed Feb 28 17:32:01 1996 PST`, `@ 3 mons`, `Tue Nov 28 17:32:01 1995 PST`}, {`Thu Feb 29 17:32:01 1996 PST`, `@ 3 mons`, `Wed Nov 29 17:32:01 1995 PST`}, {`Fri Mar 01 17:32:01 1996 PST`, `@ 3 mons`, `Fri Dec 01 17:32:01 1995 PST`}, {`Wed Feb 28 17:32:01 1996 PST`, `@ 10 days`, `Sun Feb 18 17:32:01 1996 PST`}, {`Thu Feb 29 17:32:01 1996 PST`, `@ 10 days`, `Mon Feb 19 17:32:01 1996 PST`}, {`Fri Mar 01 17:32:01 1996 PST`, `@ 10 days`, `Tue Feb 20 17:32:01 1996 PST`}, {`Wed Feb 28 17:32:01 1996 PST`, `@ 1 day 2 hours 3 mins 4 secs`, `Tue Feb 27 15:28:57 1996 PST`}, {`Wed Feb 28 17:32:01 1996 PST`, `@ 5 hours`, `Wed Feb 28 12:32:01 1996 PST`}, {`Thu Feb 29 17:32:01 1996 PST`, `@ 1 day 2 hours 3 mins 4 secs`, `Wed Feb 28 15:28:57 1996 PST`}, {`Wed Feb 28 17:32:01 1996 PST`, `@ 1 min`, `Wed Feb 28 17:31:01 1996 PST`}, {`Wed Feb 28 17:32:01 1996 PST`, `@ 14 secs ago`, `Wed Feb 28 17:32:15 1996 PST`}, {`Thu Feb 29 17:32:01 1996 PST`, `@ 5 hours`, `Thu Feb 29 12:32:01 1996 PST`}, {`Fri Mar 01 17:32:01 1996 PST`, `@ 1 day 2 hours 3 mins 4 secs`, `Thu Feb 29 15:28:57 1996 PST`}, {`Thu Feb 29 17:32:01 1996 PST`, `@ 1 min`, `Thu Feb 29 17:31:01 1996 PST`}, {`Thu Feb 29 17:32:01 1996 PST`, `@ 14 secs ago`, `Thu Feb 29 17:32:15 1996 PST`}, {`Fri Mar 01 17:32:01 1996 PST`, `@ 5 hours`, `Fri Mar 01 12:32:01 1996 PST`}, {`Fri Mar 01 17:32:01 1996 PST`, `@ 1 min`, `Fri Mar 01 17:31:01 1996 PST`}, {`Fri Mar 01 17:32:01 1996 PST`, `@ 14 secs ago`, `Fri Mar 01 17:32:15 1996 PST`}, {`Mon Dec 30 17:32:01 1996 PST`, `@ 5 mons 12 hours`, `Tue Jul 30 05:32:01 1996 PDT`}, {`Mon Dec 30 17:32:01 1996 PST`, `@ 5 mons`, `Tue Jul 30 17:32:01 1996 PDT`}, {`Tue Dec 31 17:32:01 1996 PST`, `@ 5 mons 12 hours`, `Wed Jul 31 05:32:01 1996 PDT`}, {`Tue Dec 31 17:32:01 1996 PST`, `@ 5 mons`, `Wed Jul 31 17:32:01 1996 PDT`}, {`Mon Dec 30 17:32:01 1996 PST`, `@ 3 mons`, `Mon Sep 30 17:32:01 1996 PDT`}, {`Tue Dec 31 17:32:01 1996 PST`, `@ 3 mons`, `Mon Sep 30 17:32:01 1996 PDT`}, {`Mon Dec 30 17:32:01 1996 PST`, `@ 10 days`, `Fri Dec 20 17:32:01 1996 PST`}, {`Tue Dec 31 17:32:01 1996 PST`, `@ 10 days`, `Sat Dec 21 17:32:01 1996 PST`}, {`Mon Dec 30 17:32:01 1996 PST`, `@ 1 day 2 hours 3 mins 4 secs`, `Sun Dec 29 15:28:57 1996 PST`}, {`Mon Dec 30 17:32:01 1996 PST`, `@ 5 hours`, `Mon Dec 30 12:32:01 1996 PST`}, {`Tue Dec 31 17:32:01 1996 PST`, `@ 1 day 2 hours 3 mins 4 secs`, `Mon Dec 30 15:28:57 1996 PST`}, {`Mon Dec 30 17:32:01 1996 PST`, `@ 1 min`, `Mon Dec 30 17:31:01 1996 PST`}, {`Mon Dec 30 17:32:01 1996 PST`, `@ 14 secs ago`, `Mon Dec 30 17:32:15 1996 PST`}, {`Tue Dec 31 17:32:01 1996 PST`, `@ 5 hours`, `Tue Dec 31 12:32:01 1996 PST`}, {`Tue Dec 31 17:32:01 1996 PST`, `@ 1 min`, `Tue Dec 31 17:31:01 1996 PST`}, {`Tue Dec 31 17:32:01 1996 PST`, `@ 14 secs ago`, `Tue Dec 31 17:32:15 1996 PST`}, {`Fri Dec 31 17:32:01 1999 PST`, `@ 5 mons 12 hours`, `Sat Jul 31 05:32:01 1999 PDT`}, {`Fri Dec 31 17:32:01 1999 PST`, `@ 5 mons`, `Sat Jul 31 17:32:01 1999 PDT`}, {`Sat Jan 01 17:32:01 2000 PST`, `@ 5 mons 12 hours`, `Sun Aug 01 05:32:01 1999 PDT`}, {`Sat Jan 01 17:32:01 2000 PST`, `@ 5 mons`, `Sun Aug 01 17:32:01 1999 PDT`}, {`Fri Dec 31 17:32:01 1999 PST`, `@ 3 mons`, `Thu Sep 30 17:32:01 1999 PDT`}, {`Sat Jan 01 17:32:01 2000 PST`, `@ 3 mons`, `Fri Oct 01 17:32:01 1999 PDT`}, {`Wed Mar 15 02:14:05 2000 PST`, `@ 5 mons 12 hours`, `Thu Oct 14 14:14:05 1999 PDT`}, {`Wed Mar 15 03:14:04 2000 PST`, `@ 5 mons 12 hours`, `Thu Oct 14 15:14:04 1999 PDT`}, {`Wed Mar 15 08:14:01 2000 PST`, `@ 5 mons 12 hours`, `Thu Oct 14 20:14:01 1999 PDT`}, {`Wed Mar 15 12:14:03 2000 PST`, `@ 5 mons 12 hours`, `Fri Oct 15 00:14:03 1999 PDT`}, {`Wed Mar 15 13:14:02 2000 PST`, `@ 5 mons 12 hours`, `Fri Oct 15 01:14:02 1999 PDT`}, {`Wed Mar 15 02:14:05 2000 PST`, `@ 5 mons`, `Fri Oct 15 02:14:05 1999 PDT`}, {`Wed Mar 15 03:14:04 2000 PST`, `@ 5 mons`, `Fri Oct 15 03:14:04 1999 PDT`}, {`Wed Mar 15 08:14:01 2000 PST`, `@ 5 mons`, `Fri Oct 15 08:14:01 1999 PDT`}, {`Wed Mar 15 12:14:03 2000 PST`, `@ 5 mons`, `Fri Oct 15 12:14:03 1999 PDT`}, {`Wed Mar 15 13:14:02 2000 PST`, `@ 5 mons`, `Fri Oct 15 13:14:02 1999 PDT`}, {`Wed Mar 15 02:14:05 2000 PST`, `@ 3 mons`, `Wed Dec 15 02:14:05 1999 PST`}, {`Wed Mar 15 03:14:04 2000 PST`, `@ 3 mons`, `Wed Dec 15 03:14:04 1999 PST`}, {`Wed Mar 15 08:14:01 2000 PST`, `@ 3 mons`, `Wed Dec 15 08:14:01 1999 PST`}, {`Wed Mar 15 12:14:03 2000 PST`, `@ 3 mons`, `Wed Dec 15 12:14:03 1999 PST`}, {`Wed Mar 15 13:14:02 2000 PST`, `@ 3 mons`, `Wed Dec 15 13:14:02 1999 PST`}, {`Fri Dec 31 17:32:01 1999 PST`, `@ 10 days`, `Tue Dec 21 17:32:01 1999 PST`}, {`Sat Jan 01 17:32:01 2000 PST`, `@ 10 days`, `Wed Dec 22 17:32:01 1999 PST`}, {`Fri Dec 31 17:32:01 1999 PST`, `@ 1 day 2 hours 3 mins 4 secs`, `Thu Dec 30 15:28:57 1999 PST`}, {`Fri Dec 31 17:32:01 1999 PST`, `@ 5 hours`, `Fri Dec 31 12:32:01 1999 PST`}, {`Sat Jan 01 17:32:01 2000 PST`, `@ 1 day 2 hours 3 mins 4 secs`, `Fri Dec 31 15:28:57 1999 PST`}, {`Fri Dec 31 17:32:01 1999 PST`, `@ 1 min`, `Fri Dec 31 17:31:01 1999 PST`}, {`Fri Dec 31 17:32:01 1999 PST`, `@ 14 secs ago`, `Fri Dec 31 17:32:15 1999 PST`}, {`Sat Jan 01 17:32:01 2000 PST`, `@ 5 hours`, `Sat Jan 01 12:32:01 2000 PST`}, {`Sat Jan 01 17:32:01 2000 PST`, `@ 1 min`, `Sat Jan 01 17:31:01 2000 PST`}, {`Sat Jan 01 17:32:01 2000 PST`, `@ 14 secs ago`, `Sat Jan 01 17:32:15 2000 PST`}, {`Wed Mar 15 02:14:05 2000 PST`, `@ 10 days`, `Sun Mar 05 02:14:05 2000 PST`}, {`Wed Mar 15 03:14:04 2000 PST`, `@ 10 days`, `Sun Mar 05 03:14:04 2000 PST`}, {`Wed Mar 15 08:14:01 2000 PST`, `@ 10 days`, `Sun Mar 05 08:14:01 2000 PST`}, {`Wed Mar 15 12:14:03 2000 PST`, `@ 10 days`, `Sun Mar 05 12:14:03 2000 PST`}, {`Wed Mar 15 13:14:02 2000 PST`, `@ 10 days`, `Sun Mar 05 13:14:02 2000 PST`}, {`Wed Mar 15 02:14:05 2000 PST`, `@ 1 day 2 hours 3 mins 4 secs`, `Tue Mar 14 00:11:01 2000 PST`}, {`Wed Mar 15 03:14:04 2000 PST`, `@ 1 day 2 hours 3 mins 4 secs`, `Tue Mar 14 01:11:00 2000 PST`}, {`Wed Mar 15 08:14:01 2000 PST`, `@ 1 day 2 hours 3 mins 4 secs`, `Tue Mar 14 06:10:57 2000 PST`}, {`Wed Mar 15 12:14:03 2000 PST`, `@ 1 day 2 hours 3 mins 4 secs`, `Tue Mar 14 10:10:59 2000 PST`}, {`Wed Mar 15 13:14:02 2000 PST`, `@ 1 day 2 hours 3 mins 4 secs`, `Tue Mar 14 11:10:58 2000 PST`}, {`Wed Mar 15 02:14:05 2000 PST`, `@ 5 hours`, `Tue Mar 14 21:14:05 2000 PST`}, {`Wed Mar 15 03:14:04 2000 PST`, `@ 5 hours`, `Tue Mar 14 22:14:04 2000 PST`}, {`Wed Mar 15 02:14:05 2000 PST`, `@ 1 min`, `Wed Mar 15 02:13:05 2000 PST`}, {`Wed Mar 15 02:14:05 2000 PST`, `@ 14 secs ago`, `Wed Mar 15 02:14:19 2000 PST`}, {`Wed Mar 15 03:14:04 2000 PST`, `@ 1 min`, `Wed Mar 15 03:13:04 2000 PST`}, {`Wed Mar 15 08:14:01 2000 PST`, `@ 5 hours`, `Wed Mar 15 03:14:01 2000 PST`}, {`Wed Mar 15 03:14:04 2000 PST`, `@ 14 secs ago`, `Wed Mar 15 03:14:18 2000 PST`}, {`Wed Mar 15 12:14:03 2000 PST`, `@ 5 hours`, `Wed Mar 15 07:14:03 2000 PST`}, {`Wed Mar 15 08:14:01 2000 PST`, `@ 1 min`, `Wed Mar 15 08:13:01 2000 PST`}, {`Wed Mar 15 13:14:02 2000 PST`, `@ 5 hours`, `Wed Mar 15 08:14:02 2000 PST`}, {`Wed Mar 15 08:14:01 2000 PST`, `@ 14 secs ago`, `Wed Mar 15 08:14:15 2000 PST`}, {`Wed Mar 15 12:14:03 2000 PST`, `@ 1 min`, `Wed Mar 15 12:13:03 2000 PST`}, {`Wed Mar 15 12:14:03 2000 PST`, `@ 14 secs ago`, `Wed Mar 15 12:14:17 2000 PST`}, {`Wed Mar 15 13:14:02 2000 PST`, `@ 1 min`, `Wed Mar 15 13:13:02 2000 PST`}, {`Wed Mar 15 13:14:02 2000 PST`, `@ 14 secs ago`, `Wed Mar 15 13:14:16 2000 PST`}, {`Sun Dec 31 17:32:01 2000 PST`, `@ 5 mons 12 hours`, `Mon Jul 31 05:32:01 2000 PDT`}, {`Sun Dec 31 17:32:01 2000 PST`, `@ 5 mons`, `Mon Jul 31 17:32:01 2000 PDT`}, {`Mon Jan 01 17:32:01 2001 PST`, `@ 5 mons 12 hours`, `Tue Aug 01 05:32:01 2000 PDT`}, {`Mon Jan 01 17:32:01 2001 PST`, `@ 5 mons`, `Tue Aug 01 17:32:01 2000 PDT`}, {`Sun Dec 31 17:32:01 2000 PST`, `@ 3 mons`, `Sat Sep 30 17:32:01 2000 PDT`}, {`Mon Jan 01 17:32:01 2001 PST`, `@ 3 mons`, `Sun Oct 01 17:32:01 2000 PDT`}, {`Sun Dec 31 17:32:01 2000 PST`, `@ 10 days`, `Thu Dec 21 17:32:01 2000 PST`}, {`Mon Jan 01 17:32:01 2001 PST`, `@ 10 days`, `Fri Dec 22 17:32:01 2000 PST`}, {`Sun Dec 31 17:32:01 2000 PST`, `@ 1 day 2 hours 3 mins 4 secs`, `Sat Dec 30 15:28:57 2000 PST`}, {`Sun Dec 31 17:32:01 2000 PST`, `@ 5 hours`, `Sun Dec 31 12:32:01 2000 PST`}, {`Mon Jan 01 17:32:01 2001 PST`, `@ 1 day 2 hours 3 mins 4 secs`, `Sun Dec 31 15:28:57 2000 PST`}, {`Sun Dec 31 17:32:01 2000 PST`, `@ 1 min`, `Sun Dec 31 17:31:01 2000 PST`}, {`Sun Dec 31 17:32:01 2000 PST`, `@ 14 secs ago`, `Sun Dec 31 17:32:15 2000 PST`}, {`Mon Jan 01 17:32:01 2001 PST`, `@ 5 hours`, `Mon Jan 01 12:32:01 2001 PST`}, {`Mon Jan 01 17:32:01 2001 PST`, `@ 1 min`, `Mon Jan 01 17:31:01 2001 PST`}, {`Mon Jan 01 17:32:01 2001 PST`, `@ 14 secs ago`, `Mon Jan 01 17:32:15 2001 PST`}, {`Sat Sep 22 18:19:20 2001 PDT`, `@ 5 mons 12 hours`, `Sun Apr 22 06:19:20 2001 PDT`}, {`Sat Sep 22 18:19:20 2001 PDT`, `@ 5 mons`, `Sun Apr 22 18:19:20 2001 PDT`}, {`Sat Sep 22 18:19:20 2001 PDT`, `@ 3 mons`, `Fri Jun 22 18:19:20 2001 PDT`}, {`Sat Sep 22 18:19:20 2001 PDT`, `@ 10 days`, `Wed Sep 12 18:19:20 2001 PDT`}, {`Sat Sep 22 18:19:20 2001 PDT`, `@ 1 day 2 hours 3 mins 4 secs`, `Fri Sep 21 16:16:16 2001 PDT`}, {`Sat Sep 22 18:19:20 2001 PDT`, `@ 5 hours`, `Sat Sep 22 13:19:20 2001 PDT`}, {`Sat Sep 22 18:19:20 2001 PDT`, `@ 1 min`, `Sat Sep 22 18:18:20 2001 PDT`}, {`Sat Sep 22 18:19:20 2001 PDT`, `@ 14 secs ago`, `Sat Sep 22 18:19:34 2001 PDT`}},
			},
			{
				Statement: `SELECT d.f1 AS "timestamp",
   timestamp with time zone '1980-01-06 00:00 GMT' AS gpstime_zero,
   d.f1 - timestamp with time zone '1980-01-06 00:00 GMT' AS difference
  FROM TEMP_TIMESTAMP d
  ORDER BY difference;`,
				Results: []sql.Row{{`Thu Jan 01 00:00:00 1970 PST`, `Sat Jan 05 16:00:00 1980 PST`, `@ 3656 days 16 hours ago`}, {`Wed Feb 28 17:32:01 1996 PST`, `Sat Jan 05 16:00:00 1980 PST`, `@ 5898 days 1 hour 32 mins 1 sec`}, {`Thu Feb 29 17:32:01 1996 PST`, `Sat Jan 05 16:00:00 1980 PST`, `@ 5899 days 1 hour 32 mins 1 sec`}, {`Fri Mar 01 17:32:01 1996 PST`, `Sat Jan 05 16:00:00 1980 PST`, `@ 5900 days 1 hour 32 mins 1 sec`}, {`Mon Dec 30 17:32:01 1996 PST`, `Sat Jan 05 16:00:00 1980 PST`, `@ 6204 days 1 hour 32 mins 1 sec`}, {`Tue Dec 31 17:32:01 1996 PST`, `Sat Jan 05 16:00:00 1980 PST`, `@ 6205 days 1 hour 32 mins 1 sec`}, {`Fri Dec 31 17:32:01 1999 PST`, `Sat Jan 05 16:00:00 1980 PST`, `@ 7300 days 1 hour 32 mins 1 sec`}, {`Sat Jan 01 17:32:01 2000 PST`, `Sat Jan 05 16:00:00 1980 PST`, `@ 7301 days 1 hour 32 mins 1 sec`}, {`Wed Mar 15 02:14:05 2000 PST`, `Sat Jan 05 16:00:00 1980 PST`, `@ 7374 days 10 hours 14 mins 5 secs`}, {`Wed Mar 15 03:14:04 2000 PST`, `Sat Jan 05 16:00:00 1980 PST`, `@ 7374 days 11 hours 14 mins 4 secs`}, {`Wed Mar 15 08:14:01 2000 PST`, `Sat Jan 05 16:00:00 1980 PST`, `@ 7374 days 16 hours 14 mins 1 sec`}, {`Wed Mar 15 12:14:03 2000 PST`, `Sat Jan 05 16:00:00 1980 PST`, `@ 7374 days 20 hours 14 mins 3 secs`}, {`Wed Mar 15 13:14:02 2000 PST`, `Sat Jan 05 16:00:00 1980 PST`, `@ 7374 days 21 hours 14 mins 2 secs`}, {`Sun Dec 31 17:32:01 2000 PST`, `Sat Jan 05 16:00:00 1980 PST`, `@ 7666 days 1 hour 32 mins 1 sec`}, {`Mon Jan 01 17:32:01 2001 PST`, `Sat Jan 05 16:00:00 1980 PST`, `@ 7667 days 1 hour 32 mins 1 sec`}, {`Sat Sep 22 18:19:20 2001 PDT`, `Sat Jan 05 16:00:00 1980 PST`, `@ 7931 days 1 hour 19 mins 20 secs`}},
			},
			{
				Statement: `SELECT d1.f1 AS timestamp1, d2.f1 AS timestamp2, d1.f1 - d2.f1 AS difference
  FROM TEMP_TIMESTAMP d1, TEMP_TIMESTAMP d2
  ORDER BY timestamp1, timestamp2, difference;`,
				Results: []sql.Row{{`Thu Jan 01 00:00:00 1970 PST`, `Thu Jan 01 00:00:00 1970 PST`, `@ 0`}, {`Thu Jan 01 00:00:00 1970 PST`, `Wed Feb 28 17:32:01 1996 PST`, `@ 9554 days 17 hours 32 mins 1 sec ago`}, {`Thu Jan 01 00:00:00 1970 PST`, `Thu Feb 29 17:32:01 1996 PST`, `@ 9555 days 17 hours 32 mins 1 sec ago`}, {`Thu Jan 01 00:00:00 1970 PST`, `Fri Mar 01 17:32:01 1996 PST`, `@ 9556 days 17 hours 32 mins 1 sec ago`}, {`Thu Jan 01 00:00:00 1970 PST`, `Mon Dec 30 17:32:01 1996 PST`, `@ 9860 days 17 hours 32 mins 1 sec ago`}, {`Thu Jan 01 00:00:00 1970 PST`, `Tue Dec 31 17:32:01 1996 PST`, `@ 9861 days 17 hours 32 mins 1 sec ago`}, {`Thu Jan 01 00:00:00 1970 PST`, `Fri Dec 31 17:32:01 1999 PST`, `@ 10956 days 17 hours 32 mins 1 sec ago`}, {`Thu Jan 01 00:00:00 1970 PST`, `Sat Jan 01 17:32:01 2000 PST`, `@ 10957 days 17 hours 32 mins 1 sec ago`}, {`Thu Jan 01 00:00:00 1970 PST`, `Wed Mar 15 02:14:05 2000 PST`, `@ 11031 days 2 hours 14 mins 5 secs ago`}, {`Thu Jan 01 00:00:00 1970 PST`, `Wed Mar 15 03:14:04 2000 PST`, `@ 11031 days 3 hours 14 mins 4 secs ago`}, {`Thu Jan 01 00:00:00 1970 PST`, `Wed Mar 15 08:14:01 2000 PST`, `@ 11031 days 8 hours 14 mins 1 sec ago`}, {`Thu Jan 01 00:00:00 1970 PST`, `Wed Mar 15 12:14:03 2000 PST`, `@ 11031 days 12 hours 14 mins 3 secs ago`}, {`Thu Jan 01 00:00:00 1970 PST`, `Wed Mar 15 13:14:02 2000 PST`, `@ 11031 days 13 hours 14 mins 2 secs ago`}, {`Thu Jan 01 00:00:00 1970 PST`, `Sun Dec 31 17:32:01 2000 PST`, `@ 11322 days 17 hours 32 mins 1 sec ago`}, {`Thu Jan 01 00:00:00 1970 PST`, `Mon Jan 01 17:32:01 2001 PST`, `@ 11323 days 17 hours 32 mins 1 sec ago`}, {`Thu Jan 01 00:00:00 1970 PST`, `Sat Sep 22 18:19:20 2001 PDT`, `@ 11587 days 17 hours 19 mins 20 secs ago`}, {`Wed Feb 28 17:32:01 1996 PST`, `Thu Jan 01 00:00:00 1970 PST`, `@ 9554 days 17 hours 32 mins 1 sec`}, {`Wed Feb 28 17:32:01 1996 PST`, `Wed Feb 28 17:32:01 1996 PST`, `@ 0`}, {`Wed Feb 28 17:32:01 1996 PST`, `Thu Feb 29 17:32:01 1996 PST`, `@ 1 day ago`}, {`Wed Feb 28 17:32:01 1996 PST`, `Fri Mar 01 17:32:01 1996 PST`, `@ 2 days ago`}, {`Wed Feb 28 17:32:01 1996 PST`, `Mon Dec 30 17:32:01 1996 PST`, `@ 306 days ago`}, {`Wed Feb 28 17:32:01 1996 PST`, `Tue Dec 31 17:32:01 1996 PST`, `@ 307 days ago`}, {`Wed Feb 28 17:32:01 1996 PST`, `Fri Dec 31 17:32:01 1999 PST`, `@ 1402 days ago`}, {`Wed Feb 28 17:32:01 1996 PST`, `Sat Jan 01 17:32:01 2000 PST`, `@ 1403 days ago`}, {`Wed Feb 28 17:32:01 1996 PST`, `Wed Mar 15 02:14:05 2000 PST`, `@ 1476 days 8 hours 42 mins 4 secs ago`}, {`Wed Feb 28 17:32:01 1996 PST`, `Wed Mar 15 03:14:04 2000 PST`, `@ 1476 days 9 hours 42 mins 3 secs ago`}, {`Wed Feb 28 17:32:01 1996 PST`, `Wed Mar 15 08:14:01 2000 PST`, `@ 1476 days 14 hours 42 mins ago`}, {`Wed Feb 28 17:32:01 1996 PST`, `Wed Mar 15 12:14:03 2000 PST`, `@ 1476 days 18 hours 42 mins 2 secs ago`}, {`Wed Feb 28 17:32:01 1996 PST`, `Wed Mar 15 13:14:02 2000 PST`, `@ 1476 days 19 hours 42 mins 1 sec ago`}, {`Wed Feb 28 17:32:01 1996 PST`, `Sun Dec 31 17:32:01 2000 PST`, `@ 1768 days ago`}, {`Wed Feb 28 17:32:01 1996 PST`, `Mon Jan 01 17:32:01 2001 PST`, `@ 1769 days ago`}, {`Wed Feb 28 17:32:01 1996 PST`, `Sat Sep 22 18:19:20 2001 PDT`, `@ 2032 days 23 hours 47 mins 19 secs ago`}, {`Thu Feb 29 17:32:01 1996 PST`, `Thu Jan 01 00:00:00 1970 PST`, `@ 9555 days 17 hours 32 mins 1 sec`}, {`Thu Feb 29 17:32:01 1996 PST`, `Wed Feb 28 17:32:01 1996 PST`, `@ 1 day`}, {`Thu Feb 29 17:32:01 1996 PST`, `Thu Feb 29 17:32:01 1996 PST`, `@ 0`}, {`Thu Feb 29 17:32:01 1996 PST`, `Fri Mar 01 17:32:01 1996 PST`, `@ 1 day ago`}, {`Thu Feb 29 17:32:01 1996 PST`, `Mon Dec 30 17:32:01 1996 PST`, `@ 305 days ago`}, {`Thu Feb 29 17:32:01 1996 PST`, `Tue Dec 31 17:32:01 1996 PST`, `@ 306 days ago`}, {`Thu Feb 29 17:32:01 1996 PST`, `Fri Dec 31 17:32:01 1999 PST`, `@ 1401 days ago`}, {`Thu Feb 29 17:32:01 1996 PST`, `Sat Jan 01 17:32:01 2000 PST`, `@ 1402 days ago`}, {`Thu Feb 29 17:32:01 1996 PST`, `Wed Mar 15 02:14:05 2000 PST`, `@ 1475 days 8 hours 42 mins 4 secs ago`}, {`Thu Feb 29 17:32:01 1996 PST`, `Wed Mar 15 03:14:04 2000 PST`, `@ 1475 days 9 hours 42 mins 3 secs ago`}, {`Thu Feb 29 17:32:01 1996 PST`, `Wed Mar 15 08:14:01 2000 PST`, `@ 1475 days 14 hours 42 mins ago`}, {`Thu Feb 29 17:32:01 1996 PST`, `Wed Mar 15 12:14:03 2000 PST`, `@ 1475 days 18 hours 42 mins 2 secs ago`}, {`Thu Feb 29 17:32:01 1996 PST`, `Wed Mar 15 13:14:02 2000 PST`, `@ 1475 days 19 hours 42 mins 1 sec ago`}, {`Thu Feb 29 17:32:01 1996 PST`, `Sun Dec 31 17:32:01 2000 PST`, `@ 1767 days ago`}, {`Thu Feb 29 17:32:01 1996 PST`, `Mon Jan 01 17:32:01 2001 PST`, `@ 1768 days ago`}, {`Thu Feb 29 17:32:01 1996 PST`, `Sat Sep 22 18:19:20 2001 PDT`, `@ 2031 days 23 hours 47 mins 19 secs ago`}, {`Fri Mar 01 17:32:01 1996 PST`, `Thu Jan 01 00:00:00 1970 PST`, `@ 9556 days 17 hours 32 mins 1 sec`}, {`Fri Mar 01 17:32:01 1996 PST`, `Wed Feb 28 17:32:01 1996 PST`, `@ 2 days`}, {`Fri Mar 01 17:32:01 1996 PST`, `Thu Feb 29 17:32:01 1996 PST`, `@ 1 day`}, {`Fri Mar 01 17:32:01 1996 PST`, `Fri Mar 01 17:32:01 1996 PST`, `@ 0`}, {`Fri Mar 01 17:32:01 1996 PST`, `Mon Dec 30 17:32:01 1996 PST`, `@ 304 days ago`}, {`Fri Mar 01 17:32:01 1996 PST`, `Tue Dec 31 17:32:01 1996 PST`, `@ 305 days ago`}, {`Fri Mar 01 17:32:01 1996 PST`, `Fri Dec 31 17:32:01 1999 PST`, `@ 1400 days ago`}, {`Fri Mar 01 17:32:01 1996 PST`, `Sat Jan 01 17:32:01 2000 PST`, `@ 1401 days ago`}, {`Fri Mar 01 17:32:01 1996 PST`, `Wed Mar 15 02:14:05 2000 PST`, `@ 1474 days 8 hours 42 mins 4 secs ago`}, {`Fri Mar 01 17:32:01 1996 PST`, `Wed Mar 15 03:14:04 2000 PST`, `@ 1474 days 9 hours 42 mins 3 secs ago`}, {`Fri Mar 01 17:32:01 1996 PST`, `Wed Mar 15 08:14:01 2000 PST`, `@ 1474 days 14 hours 42 mins ago`}, {`Fri Mar 01 17:32:01 1996 PST`, `Wed Mar 15 12:14:03 2000 PST`, `@ 1474 days 18 hours 42 mins 2 secs ago`}, {`Fri Mar 01 17:32:01 1996 PST`, `Wed Mar 15 13:14:02 2000 PST`, `@ 1474 days 19 hours 42 mins 1 sec ago`}, {`Fri Mar 01 17:32:01 1996 PST`, `Sun Dec 31 17:32:01 2000 PST`, `@ 1766 days ago`}, {`Fri Mar 01 17:32:01 1996 PST`, `Mon Jan 01 17:32:01 2001 PST`, `@ 1767 days ago`}, {`Fri Mar 01 17:32:01 1996 PST`, `Sat Sep 22 18:19:20 2001 PDT`, `@ 2030 days 23 hours 47 mins 19 secs ago`}, {`Mon Dec 30 17:32:01 1996 PST`, `Thu Jan 01 00:00:00 1970 PST`, `@ 9860 days 17 hours 32 mins 1 sec`}, {`Mon Dec 30 17:32:01 1996 PST`, `Wed Feb 28 17:32:01 1996 PST`, `@ 306 days`}, {`Mon Dec 30 17:32:01 1996 PST`, `Thu Feb 29 17:32:01 1996 PST`, `@ 305 days`}, {`Mon Dec 30 17:32:01 1996 PST`, `Fri Mar 01 17:32:01 1996 PST`, `@ 304 days`}, {`Mon Dec 30 17:32:01 1996 PST`, `Mon Dec 30 17:32:01 1996 PST`, `@ 0`}, {`Mon Dec 30 17:32:01 1996 PST`, `Tue Dec 31 17:32:01 1996 PST`, `@ 1 day ago`}, {`Mon Dec 30 17:32:01 1996 PST`, `Fri Dec 31 17:32:01 1999 PST`, `@ 1096 days ago`}, {`Mon Dec 30 17:32:01 1996 PST`, `Sat Jan 01 17:32:01 2000 PST`, `@ 1097 days ago`}, {`Mon Dec 30 17:32:01 1996 PST`, `Wed Mar 15 02:14:05 2000 PST`, `@ 1170 days 8 hours 42 mins 4 secs ago`}, {`Mon Dec 30 17:32:01 1996 PST`, `Wed Mar 15 03:14:04 2000 PST`, `@ 1170 days 9 hours 42 mins 3 secs ago`}, {`Mon Dec 30 17:32:01 1996 PST`, `Wed Mar 15 08:14:01 2000 PST`, `@ 1170 days 14 hours 42 mins ago`}, {`Mon Dec 30 17:32:01 1996 PST`, `Wed Mar 15 12:14:03 2000 PST`, `@ 1170 days 18 hours 42 mins 2 secs ago`}, {`Mon Dec 30 17:32:01 1996 PST`, `Wed Mar 15 13:14:02 2000 PST`, `@ 1170 days 19 hours 42 mins 1 sec ago`}, {`Mon Dec 30 17:32:01 1996 PST`, `Sun Dec 31 17:32:01 2000 PST`, `@ 1462 days ago`}, {`Mon Dec 30 17:32:01 1996 PST`, `Mon Jan 01 17:32:01 2001 PST`, `@ 1463 days ago`}, {`Mon Dec 30 17:32:01 1996 PST`, `Sat Sep 22 18:19:20 2001 PDT`, `@ 1726 days 23 hours 47 mins 19 secs ago`}, {`Tue Dec 31 17:32:01 1996 PST`, `Thu Jan 01 00:00:00 1970 PST`, `@ 9861 days 17 hours 32 mins 1 sec`}, {`Tue Dec 31 17:32:01 1996 PST`, `Wed Feb 28 17:32:01 1996 PST`, `@ 307 days`}, {`Tue Dec 31 17:32:01 1996 PST`, `Thu Feb 29 17:32:01 1996 PST`, `@ 306 days`}, {`Tue Dec 31 17:32:01 1996 PST`, `Fri Mar 01 17:32:01 1996 PST`, `@ 305 days`}, {`Tue Dec 31 17:32:01 1996 PST`, `Mon Dec 30 17:32:01 1996 PST`, `@ 1 day`}, {`Tue Dec 31 17:32:01 1996 PST`, `Tue Dec 31 17:32:01 1996 PST`, `@ 0`}, {`Tue Dec 31 17:32:01 1996 PST`, `Fri Dec 31 17:32:01 1999 PST`, `@ 1095 days ago`}, {`Tue Dec 31 17:32:01 1996 PST`, `Sat Jan 01 17:32:01 2000 PST`, `@ 1096 days ago`}, {`Tue Dec 31 17:32:01 1996 PST`, `Wed Mar 15 02:14:05 2000 PST`, `@ 1169 days 8 hours 42 mins 4 secs ago`}, {`Tue Dec 31 17:32:01 1996 PST`, `Wed Mar 15 03:14:04 2000 PST`, `@ 1169 days 9 hours 42 mins 3 secs ago`}, {`Tue Dec 31 17:32:01 1996 PST`, `Wed Mar 15 08:14:01 2000 PST`, `@ 1169 days 14 hours 42 mins ago`}, {`Tue Dec 31 17:32:01 1996 PST`, `Wed Mar 15 12:14:03 2000 PST`, `@ 1169 days 18 hours 42 mins 2 secs ago`}, {`Tue Dec 31 17:32:01 1996 PST`, `Wed Mar 15 13:14:02 2000 PST`, `@ 1169 days 19 hours 42 mins 1 sec ago`}, {`Tue Dec 31 17:32:01 1996 PST`, `Sun Dec 31 17:32:01 2000 PST`, `@ 1461 days ago`}, {`Tue Dec 31 17:32:01 1996 PST`, `Mon Jan 01 17:32:01 2001 PST`, `@ 1462 days ago`}, {`Tue Dec 31 17:32:01 1996 PST`, `Sat Sep 22 18:19:20 2001 PDT`, `@ 1725 days 23 hours 47 mins 19 secs ago`}, {`Fri Dec 31 17:32:01 1999 PST`, `Thu Jan 01 00:00:00 1970 PST`, `@ 10956 days 17 hours 32 mins 1 sec`}, {`Fri Dec 31 17:32:01 1999 PST`, `Wed Feb 28 17:32:01 1996 PST`, `@ 1402 days`}, {`Fri Dec 31 17:32:01 1999 PST`, `Thu Feb 29 17:32:01 1996 PST`, `@ 1401 days`}, {`Fri Dec 31 17:32:01 1999 PST`, `Fri Mar 01 17:32:01 1996 PST`, `@ 1400 days`}, {`Fri Dec 31 17:32:01 1999 PST`, `Mon Dec 30 17:32:01 1996 PST`, `@ 1096 days`}, {`Fri Dec 31 17:32:01 1999 PST`, `Tue Dec 31 17:32:01 1996 PST`, `@ 1095 days`}, {`Fri Dec 31 17:32:01 1999 PST`, `Fri Dec 31 17:32:01 1999 PST`, `@ 0`}, {`Fri Dec 31 17:32:01 1999 PST`, `Sat Jan 01 17:32:01 2000 PST`, `@ 1 day ago`}, {`Fri Dec 31 17:32:01 1999 PST`, `Wed Mar 15 02:14:05 2000 PST`, `@ 74 days 8 hours 42 mins 4 secs ago`}, {`Fri Dec 31 17:32:01 1999 PST`, `Wed Mar 15 03:14:04 2000 PST`, `@ 74 days 9 hours 42 mins 3 secs ago`}, {`Fri Dec 31 17:32:01 1999 PST`, `Wed Mar 15 08:14:01 2000 PST`, `@ 74 days 14 hours 42 mins ago`}, {`Fri Dec 31 17:32:01 1999 PST`, `Wed Mar 15 12:14:03 2000 PST`, `@ 74 days 18 hours 42 mins 2 secs ago`}, {`Fri Dec 31 17:32:01 1999 PST`, `Wed Mar 15 13:14:02 2000 PST`, `@ 74 days 19 hours 42 mins 1 sec ago`}, {`Fri Dec 31 17:32:01 1999 PST`, `Sun Dec 31 17:32:01 2000 PST`, `@ 366 days ago`}, {`Fri Dec 31 17:32:01 1999 PST`, `Mon Jan 01 17:32:01 2001 PST`, `@ 367 days ago`}, {`Fri Dec 31 17:32:01 1999 PST`, `Sat Sep 22 18:19:20 2001 PDT`, `@ 630 days 23 hours 47 mins 19 secs ago`}, {`Sat Jan 01 17:32:01 2000 PST`, `Thu Jan 01 00:00:00 1970 PST`, `@ 10957 days 17 hours 32 mins 1 sec`}, {`Sat Jan 01 17:32:01 2000 PST`, `Wed Feb 28 17:32:01 1996 PST`, `@ 1403 days`}, {`Sat Jan 01 17:32:01 2000 PST`, `Thu Feb 29 17:32:01 1996 PST`, `@ 1402 days`}, {`Sat Jan 01 17:32:01 2000 PST`, `Fri Mar 01 17:32:01 1996 PST`, `@ 1401 days`}, {`Sat Jan 01 17:32:01 2000 PST`, `Mon Dec 30 17:32:01 1996 PST`, `@ 1097 days`}, {`Sat Jan 01 17:32:01 2000 PST`, `Tue Dec 31 17:32:01 1996 PST`, `@ 1096 days`}, {`Sat Jan 01 17:32:01 2000 PST`, `Fri Dec 31 17:32:01 1999 PST`, `@ 1 day`}, {`Sat Jan 01 17:32:01 2000 PST`, `Sat Jan 01 17:32:01 2000 PST`, `@ 0`}, {`Sat Jan 01 17:32:01 2000 PST`, `Wed Mar 15 02:14:05 2000 PST`, `@ 73 days 8 hours 42 mins 4 secs ago`}, {`Sat Jan 01 17:32:01 2000 PST`, `Wed Mar 15 03:14:04 2000 PST`, `@ 73 days 9 hours 42 mins 3 secs ago`}, {`Sat Jan 01 17:32:01 2000 PST`, `Wed Mar 15 08:14:01 2000 PST`, `@ 73 days 14 hours 42 mins ago`}, {`Sat Jan 01 17:32:01 2000 PST`, `Wed Mar 15 12:14:03 2000 PST`, `@ 73 days 18 hours 42 mins 2 secs ago`}, {`Sat Jan 01 17:32:01 2000 PST`, `Wed Mar 15 13:14:02 2000 PST`, `@ 73 days 19 hours 42 mins 1 sec ago`}, {`Sat Jan 01 17:32:01 2000 PST`, `Sun Dec 31 17:32:01 2000 PST`, `@ 365 days ago`}, {`Sat Jan 01 17:32:01 2000 PST`, `Mon Jan 01 17:32:01 2001 PST`, `@ 366 days ago`}, {`Sat Jan 01 17:32:01 2000 PST`, `Sat Sep 22 18:19:20 2001 PDT`, `@ 629 days 23 hours 47 mins 19 secs ago`}, {`Wed Mar 15 02:14:05 2000 PST`, `Thu Jan 01 00:00:00 1970 PST`, `@ 11031 days 2 hours 14 mins 5 secs`}, {`Wed Mar 15 02:14:05 2000 PST`, `Wed Feb 28 17:32:01 1996 PST`, `@ 1476 days 8 hours 42 mins 4 secs`}, {`Wed Mar 15 02:14:05 2000 PST`, `Thu Feb 29 17:32:01 1996 PST`, `@ 1475 days 8 hours 42 mins 4 secs`}, {`Wed Mar 15 02:14:05 2000 PST`, `Fri Mar 01 17:32:01 1996 PST`, `@ 1474 days 8 hours 42 mins 4 secs`}, {`Wed Mar 15 02:14:05 2000 PST`, `Mon Dec 30 17:32:01 1996 PST`, `@ 1170 days 8 hours 42 mins 4 secs`}, {`Wed Mar 15 02:14:05 2000 PST`, `Tue Dec 31 17:32:01 1996 PST`, `@ 1169 days 8 hours 42 mins 4 secs`}, {`Wed Mar 15 02:14:05 2000 PST`, `Fri Dec 31 17:32:01 1999 PST`, `@ 74 days 8 hours 42 mins 4 secs`}, {`Wed Mar 15 02:14:05 2000 PST`, `Sat Jan 01 17:32:01 2000 PST`, `@ 73 days 8 hours 42 mins 4 secs`}, {`Wed Mar 15 02:14:05 2000 PST`, `Wed Mar 15 02:14:05 2000 PST`, `@ 0`}, {`Wed Mar 15 02:14:05 2000 PST`, `Wed Mar 15 03:14:04 2000 PST`, `@ 59 mins 59 secs ago`}, {`Wed Mar 15 02:14:05 2000 PST`, `Wed Mar 15 08:14:01 2000 PST`, `@ 5 hours 59 mins 56 secs ago`}, {`Wed Mar 15 02:14:05 2000 PST`, `Wed Mar 15 12:14:03 2000 PST`, `@ 9 hours 59 mins 58 secs ago`}, {`Wed Mar 15 02:14:05 2000 PST`, `Wed Mar 15 13:14:02 2000 PST`, `@ 10 hours 59 mins 57 secs ago`}, {`Wed Mar 15 02:14:05 2000 PST`, `Sun Dec 31 17:32:01 2000 PST`, `@ 291 days 15 hours 17 mins 56 secs ago`}, {`Wed Mar 15 02:14:05 2000 PST`, `Mon Jan 01 17:32:01 2001 PST`, `@ 292 days 15 hours 17 mins 56 secs ago`}, {`Wed Mar 15 02:14:05 2000 PST`, `Sat Sep 22 18:19:20 2001 PDT`, `@ 556 days 15 hours 5 mins 15 secs ago`}, {`Wed Mar 15 03:14:04 2000 PST`, `Thu Jan 01 00:00:00 1970 PST`, `@ 11031 days 3 hours 14 mins 4 secs`}, {`Wed Mar 15 03:14:04 2000 PST`, `Wed Feb 28 17:32:01 1996 PST`, `@ 1476 days 9 hours 42 mins 3 secs`}, {`Wed Mar 15 03:14:04 2000 PST`, `Thu Feb 29 17:32:01 1996 PST`, `@ 1475 days 9 hours 42 mins 3 secs`}, {`Wed Mar 15 03:14:04 2000 PST`, `Fri Mar 01 17:32:01 1996 PST`, `@ 1474 days 9 hours 42 mins 3 secs`}, {`Wed Mar 15 03:14:04 2000 PST`, `Mon Dec 30 17:32:01 1996 PST`, `@ 1170 days 9 hours 42 mins 3 secs`}, {`Wed Mar 15 03:14:04 2000 PST`, `Tue Dec 31 17:32:01 1996 PST`, `@ 1169 days 9 hours 42 mins 3 secs`}, {`Wed Mar 15 03:14:04 2000 PST`, `Fri Dec 31 17:32:01 1999 PST`, `@ 74 days 9 hours 42 mins 3 secs`}, {`Wed Mar 15 03:14:04 2000 PST`, `Sat Jan 01 17:32:01 2000 PST`, `@ 73 days 9 hours 42 mins 3 secs`}, {`Wed Mar 15 03:14:04 2000 PST`, `Wed Mar 15 02:14:05 2000 PST`, `@ 59 mins 59 secs`}, {`Wed Mar 15 03:14:04 2000 PST`, `Wed Mar 15 03:14:04 2000 PST`, `@ 0`}, {`Wed Mar 15 03:14:04 2000 PST`, `Wed Mar 15 08:14:01 2000 PST`, `@ 4 hours 59 mins 57 secs ago`}, {`Wed Mar 15 03:14:04 2000 PST`, `Wed Mar 15 12:14:03 2000 PST`, `@ 8 hours 59 mins 59 secs ago`}, {`Wed Mar 15 03:14:04 2000 PST`, `Wed Mar 15 13:14:02 2000 PST`, `@ 9 hours 59 mins 58 secs ago`}, {`Wed Mar 15 03:14:04 2000 PST`, `Sun Dec 31 17:32:01 2000 PST`, `@ 291 days 14 hours 17 mins 57 secs ago`}, {`Wed Mar 15 03:14:04 2000 PST`, `Mon Jan 01 17:32:01 2001 PST`, `@ 292 days 14 hours 17 mins 57 secs ago`}, {`Wed Mar 15 03:14:04 2000 PST`, `Sat Sep 22 18:19:20 2001 PDT`, `@ 556 days 14 hours 5 mins 16 secs ago`}, {`Wed Mar 15 08:14:01 2000 PST`, `Thu Jan 01 00:00:00 1970 PST`, `@ 11031 days 8 hours 14 mins 1 sec`}, {`Wed Mar 15 08:14:01 2000 PST`, `Wed Feb 28 17:32:01 1996 PST`, `@ 1476 days 14 hours 42 mins`}, {`Wed Mar 15 08:14:01 2000 PST`, `Thu Feb 29 17:32:01 1996 PST`, `@ 1475 days 14 hours 42 mins`}, {`Wed Mar 15 08:14:01 2000 PST`, `Fri Mar 01 17:32:01 1996 PST`, `@ 1474 days 14 hours 42 mins`}, {`Wed Mar 15 08:14:01 2000 PST`, `Mon Dec 30 17:32:01 1996 PST`, `@ 1170 days 14 hours 42 mins`}, {`Wed Mar 15 08:14:01 2000 PST`, `Tue Dec 31 17:32:01 1996 PST`, `@ 1169 days 14 hours 42 mins`}, {`Wed Mar 15 08:14:01 2000 PST`, `Fri Dec 31 17:32:01 1999 PST`, `@ 74 days 14 hours 42 mins`}, {`Wed Mar 15 08:14:01 2000 PST`, `Sat Jan 01 17:32:01 2000 PST`, `@ 73 days 14 hours 42 mins`}, {`Wed Mar 15 08:14:01 2000 PST`, `Wed Mar 15 02:14:05 2000 PST`, `@ 5 hours 59 mins 56 secs`}, {`Wed Mar 15 08:14:01 2000 PST`, `Wed Mar 15 03:14:04 2000 PST`, `@ 4 hours 59 mins 57 secs`}, {`Wed Mar 15 08:14:01 2000 PST`, `Wed Mar 15 08:14:01 2000 PST`, `@ 0`}, {`Wed Mar 15 08:14:01 2000 PST`, `Wed Mar 15 12:14:03 2000 PST`, `@ 4 hours 2 secs ago`}, {`Wed Mar 15 08:14:01 2000 PST`, `Wed Mar 15 13:14:02 2000 PST`, `@ 5 hours 1 sec ago`}, {`Wed Mar 15 08:14:01 2000 PST`, `Sun Dec 31 17:32:01 2000 PST`, `@ 291 days 9 hours 18 mins ago`}, {`Wed Mar 15 08:14:01 2000 PST`, `Mon Jan 01 17:32:01 2001 PST`, `@ 292 days 9 hours 18 mins ago`}, {`Wed Mar 15 08:14:01 2000 PST`, `Sat Sep 22 18:19:20 2001 PDT`, `@ 556 days 9 hours 5 mins 19 secs ago`}, {`Wed Mar 15 12:14:03 2000 PST`, `Thu Jan 01 00:00:00 1970 PST`, `@ 11031 days 12 hours 14 mins 3 secs`}, {`Wed Mar 15 12:14:03 2000 PST`, `Wed Feb 28 17:32:01 1996 PST`, `@ 1476 days 18 hours 42 mins 2 secs`}, {`Wed Mar 15 12:14:03 2000 PST`, `Thu Feb 29 17:32:01 1996 PST`, `@ 1475 days 18 hours 42 mins 2 secs`}, {`Wed Mar 15 12:14:03 2000 PST`, `Fri Mar 01 17:32:01 1996 PST`, `@ 1474 days 18 hours 42 mins 2 secs`}, {`Wed Mar 15 12:14:03 2000 PST`, `Mon Dec 30 17:32:01 1996 PST`, `@ 1170 days 18 hours 42 mins 2 secs`}, {`Wed Mar 15 12:14:03 2000 PST`, `Tue Dec 31 17:32:01 1996 PST`, `@ 1169 days 18 hours 42 mins 2 secs`}, {`Wed Mar 15 12:14:03 2000 PST`, `Fri Dec 31 17:32:01 1999 PST`, `@ 74 days 18 hours 42 mins 2 secs`}, {`Wed Mar 15 12:14:03 2000 PST`, `Sat Jan 01 17:32:01 2000 PST`, `@ 73 days 18 hours 42 mins 2 secs`}, {`Wed Mar 15 12:14:03 2000 PST`, `Wed Mar 15 02:14:05 2000 PST`, `@ 9 hours 59 mins 58 secs`}, {`Wed Mar 15 12:14:03 2000 PST`, `Wed Mar 15 03:14:04 2000 PST`, `@ 8 hours 59 mins 59 secs`}, {`Wed Mar 15 12:14:03 2000 PST`, `Wed Mar 15 08:14:01 2000 PST`, `@ 4 hours 2 secs`}, {`Wed Mar 15 12:14:03 2000 PST`, `Wed Mar 15 12:14:03 2000 PST`, `@ 0`}, {`Wed Mar 15 12:14:03 2000 PST`, `Wed Mar 15 13:14:02 2000 PST`, `@ 59 mins 59 secs ago`}, {`Wed Mar 15 12:14:03 2000 PST`, `Sun Dec 31 17:32:01 2000 PST`, `@ 291 days 5 hours 17 mins 58 secs ago`}, {`Wed Mar 15 12:14:03 2000 PST`, `Mon Jan 01 17:32:01 2001 PST`, `@ 292 days 5 hours 17 mins 58 secs ago`}, {`Wed Mar 15 12:14:03 2000 PST`, `Sat Sep 22 18:19:20 2001 PDT`, `@ 556 days 5 hours 5 mins 17 secs ago`}, {`Wed Mar 15 13:14:02 2000 PST`, `Thu Jan 01 00:00:00 1970 PST`, `@ 11031 days 13 hours 14 mins 2 secs`}, {`Wed Mar 15 13:14:02 2000 PST`, `Wed Feb 28 17:32:01 1996 PST`, `@ 1476 days 19 hours 42 mins 1 sec`}, {`Wed Mar 15 13:14:02 2000 PST`, `Thu Feb 29 17:32:01 1996 PST`, `@ 1475 days 19 hours 42 mins 1 sec`}, {`Wed Mar 15 13:14:02 2000 PST`, `Fri Mar 01 17:32:01 1996 PST`, `@ 1474 days 19 hours 42 mins 1 sec`}, {`Wed Mar 15 13:14:02 2000 PST`, `Mon Dec 30 17:32:01 1996 PST`, `@ 1170 days 19 hours 42 mins 1 sec`}, {`Wed Mar 15 13:14:02 2000 PST`, `Tue Dec 31 17:32:01 1996 PST`, `@ 1169 days 19 hours 42 mins 1 sec`}, {`Wed Mar 15 13:14:02 2000 PST`, `Fri Dec 31 17:32:01 1999 PST`, `@ 74 days 19 hours 42 mins 1 sec`}, {`Wed Mar 15 13:14:02 2000 PST`, `Sat Jan 01 17:32:01 2000 PST`, `@ 73 days 19 hours 42 mins 1 sec`}, {`Wed Mar 15 13:14:02 2000 PST`, `Wed Mar 15 02:14:05 2000 PST`, `@ 10 hours 59 mins 57 secs`}, {`Wed Mar 15 13:14:02 2000 PST`, `Wed Mar 15 03:14:04 2000 PST`, `@ 9 hours 59 mins 58 secs`}, {`Wed Mar 15 13:14:02 2000 PST`, `Wed Mar 15 08:14:01 2000 PST`, `@ 5 hours 1 sec`}, {`Wed Mar 15 13:14:02 2000 PST`, `Wed Mar 15 12:14:03 2000 PST`, `@ 59 mins 59 secs`}, {`Wed Mar 15 13:14:02 2000 PST`, `Wed Mar 15 13:14:02 2000 PST`, `@ 0`}, {`Wed Mar 15 13:14:02 2000 PST`, `Sun Dec 31 17:32:01 2000 PST`, `@ 291 days 4 hours 17 mins 59 secs ago`}, {`Wed Mar 15 13:14:02 2000 PST`, `Mon Jan 01 17:32:01 2001 PST`, `@ 292 days 4 hours 17 mins 59 secs ago`}, {`Wed Mar 15 13:14:02 2000 PST`, `Sat Sep 22 18:19:20 2001 PDT`, `@ 556 days 4 hours 5 mins 18 secs ago`}, {`Sun Dec 31 17:32:01 2000 PST`, `Thu Jan 01 00:00:00 1970 PST`, `@ 11322 days 17 hours 32 mins 1 sec`}, {`Sun Dec 31 17:32:01 2000 PST`, `Wed Feb 28 17:32:01 1996 PST`, `@ 1768 days`}, {`Sun Dec 31 17:32:01 2000 PST`, `Thu Feb 29 17:32:01 1996 PST`, `@ 1767 days`}, {`Sun Dec 31 17:32:01 2000 PST`, `Fri Mar 01 17:32:01 1996 PST`, `@ 1766 days`}, {`Sun Dec 31 17:32:01 2000 PST`, `Mon Dec 30 17:32:01 1996 PST`, `@ 1462 days`}, {`Sun Dec 31 17:32:01 2000 PST`, `Tue Dec 31 17:32:01 1996 PST`, `@ 1461 days`}, {`Sun Dec 31 17:32:01 2000 PST`, `Fri Dec 31 17:32:01 1999 PST`, `@ 366 days`}, {`Sun Dec 31 17:32:01 2000 PST`, `Sat Jan 01 17:32:01 2000 PST`, `@ 365 days`}, {`Sun Dec 31 17:32:01 2000 PST`, `Wed Mar 15 02:14:05 2000 PST`, `@ 291 days 15 hours 17 mins 56 secs`}, {`Sun Dec 31 17:32:01 2000 PST`, `Wed Mar 15 03:14:04 2000 PST`, `@ 291 days 14 hours 17 mins 57 secs`}, {`Sun Dec 31 17:32:01 2000 PST`, `Wed Mar 15 08:14:01 2000 PST`, `@ 291 days 9 hours 18 mins`}, {`Sun Dec 31 17:32:01 2000 PST`, `Wed Mar 15 12:14:03 2000 PST`, `@ 291 days 5 hours 17 mins 58 secs`}, {`Sun Dec 31 17:32:01 2000 PST`, `Wed Mar 15 13:14:02 2000 PST`, `@ 291 days 4 hours 17 mins 59 secs`}, {`Sun Dec 31 17:32:01 2000 PST`, `Sun Dec 31 17:32:01 2000 PST`, `@ 0`}, {`Sun Dec 31 17:32:01 2000 PST`, `Mon Jan 01 17:32:01 2001 PST`, `@ 1 day ago`}, {`Sun Dec 31 17:32:01 2000 PST`, `Sat Sep 22 18:19:20 2001 PDT`, `@ 264 days 23 hours 47 mins 19 secs ago`}, {`Mon Jan 01 17:32:01 2001 PST`, `Thu Jan 01 00:00:00 1970 PST`, `@ 11323 days 17 hours 32 mins 1 sec`}, {`Mon Jan 01 17:32:01 2001 PST`, `Wed Feb 28 17:32:01 1996 PST`, `@ 1769 days`}, {`Mon Jan 01 17:32:01 2001 PST`, `Thu Feb 29 17:32:01 1996 PST`, `@ 1768 days`}, {`Mon Jan 01 17:32:01 2001 PST`, `Fri Mar 01 17:32:01 1996 PST`, `@ 1767 days`}, {`Mon Jan 01 17:32:01 2001 PST`, `Mon Dec 30 17:32:01 1996 PST`, `@ 1463 days`}, {`Mon Jan 01 17:32:01 2001 PST`, `Tue Dec 31 17:32:01 1996 PST`, `@ 1462 days`}, {`Mon Jan 01 17:32:01 2001 PST`, `Fri Dec 31 17:32:01 1999 PST`, `@ 367 days`}, {`Mon Jan 01 17:32:01 2001 PST`, `Sat Jan 01 17:32:01 2000 PST`, `@ 366 days`}, {`Mon Jan 01 17:32:01 2001 PST`, `Wed Mar 15 02:14:05 2000 PST`, `@ 292 days 15 hours 17 mins 56 secs`}, {`Mon Jan 01 17:32:01 2001 PST`, `Wed Mar 15 03:14:04 2000 PST`, `@ 292 days 14 hours 17 mins 57 secs`}, {`Mon Jan 01 17:32:01 2001 PST`, `Wed Mar 15 08:14:01 2000 PST`, `@ 292 days 9 hours 18 mins`}, {`Mon Jan 01 17:32:01 2001 PST`, `Wed Mar 15 12:14:03 2000 PST`, `@ 292 days 5 hours 17 mins 58 secs`}, {`Mon Jan 01 17:32:01 2001 PST`, `Wed Mar 15 13:14:02 2000 PST`, `@ 292 days 4 hours 17 mins 59 secs`}, {`Mon Jan 01 17:32:01 2001 PST`, `Sun Dec 31 17:32:01 2000 PST`, `@ 1 day`}, {`Mon Jan 01 17:32:01 2001 PST`, `Mon Jan 01 17:32:01 2001 PST`, `@ 0`}, {`Mon Jan 01 17:32:01 2001 PST`, `Sat Sep 22 18:19:20 2001 PDT`, `@ 263 days 23 hours 47 mins 19 secs ago`}, {`Sat Sep 22 18:19:20 2001 PDT`, `Thu Jan 01 00:00:00 1970 PST`, `@ 11587 days 17 hours 19 mins 20 secs`}, {`Sat Sep 22 18:19:20 2001 PDT`, `Wed Feb 28 17:32:01 1996 PST`, `@ 2032 days 23 hours 47 mins 19 secs`}, {`Sat Sep 22 18:19:20 2001 PDT`, `Thu Feb 29 17:32:01 1996 PST`, `@ 2031 days 23 hours 47 mins 19 secs`}, {`Sat Sep 22 18:19:20 2001 PDT`, `Fri Mar 01 17:32:01 1996 PST`, `@ 2030 days 23 hours 47 mins 19 secs`}, {`Sat Sep 22 18:19:20 2001 PDT`, `Mon Dec 30 17:32:01 1996 PST`, `@ 1726 days 23 hours 47 mins 19 secs`}, {`Sat Sep 22 18:19:20 2001 PDT`, `Tue Dec 31 17:32:01 1996 PST`, `@ 1725 days 23 hours 47 mins 19 secs`}, {`Sat Sep 22 18:19:20 2001 PDT`, `Fri Dec 31 17:32:01 1999 PST`, `@ 630 days 23 hours 47 mins 19 secs`}, {`Sat Sep 22 18:19:20 2001 PDT`, `Sat Jan 01 17:32:01 2000 PST`, `@ 629 days 23 hours 47 mins 19 secs`}, {`Sat Sep 22 18:19:20 2001 PDT`, `Wed Mar 15 02:14:05 2000 PST`, `@ 556 days 15 hours 5 mins 15 secs`}, {`Sat Sep 22 18:19:20 2001 PDT`, `Wed Mar 15 03:14:04 2000 PST`, `@ 556 days 14 hours 5 mins 16 secs`}, {`Sat Sep 22 18:19:20 2001 PDT`, `Wed Mar 15 08:14:01 2000 PST`, `@ 556 days 9 hours 5 mins 19 secs`}, {`Sat Sep 22 18:19:20 2001 PDT`, `Wed Mar 15 12:14:03 2000 PST`, `@ 556 days 5 hours 5 mins 17 secs`}, {`Sat Sep 22 18:19:20 2001 PDT`, `Wed Mar 15 13:14:02 2000 PST`, `@ 556 days 4 hours 5 mins 18 secs`}, {`Sat Sep 22 18:19:20 2001 PDT`, `Sun Dec 31 17:32:01 2000 PST`, `@ 264 days 23 hours 47 mins 19 secs`}, {`Sat Sep 22 18:19:20 2001 PDT`, `Mon Jan 01 17:32:01 2001 PST`, `@ 263 days 23 hours 47 mins 19 secs`}, {`Sat Sep 22 18:19:20 2001 PDT`, `Sat Sep 22 18:19:20 2001 PDT`, `@ 0`}},
			},
			{
				Statement: `SELECT f1 AS "timestamp", date(f1) AS date
  FROM TEMP_TIMESTAMP
  WHERE f1 <> timestamp 'now'
  ORDER BY date, "timestamp";`,
				Results: []sql.Row{{`Thu Jan 01 00:00:00 1970 PST`, `01-01-1970`}, {`Wed Feb 28 17:32:01 1996 PST`, `02-28-1996`}, {`Thu Feb 29 17:32:01 1996 PST`, `02-29-1996`}, {`Fri Mar 01 17:32:01 1996 PST`, `03-01-1996`}, {`Mon Dec 30 17:32:01 1996 PST`, `12-30-1996`}, {`Tue Dec 31 17:32:01 1996 PST`, `12-31-1996`}, {`Fri Dec 31 17:32:01 1999 PST`, `12-31-1999`}, {`Sat Jan 01 17:32:01 2000 PST`, `01-01-2000`}, {`Wed Mar 15 02:14:05 2000 PST`, `03-15-2000`}, {`Wed Mar 15 03:14:04 2000 PST`, `03-15-2000`}, {`Wed Mar 15 08:14:01 2000 PST`, `03-15-2000`}, {`Wed Mar 15 12:14:03 2000 PST`, `03-15-2000`}, {`Wed Mar 15 13:14:02 2000 PST`, `03-15-2000`}, {`Sun Dec 31 17:32:01 2000 PST`, `12-31-2000`}, {`Mon Jan 01 17:32:01 2001 PST`, `01-01-2001`}, {`Sat Sep 22 18:19:20 2001 PDT`, `09-22-2001`}},
			},
			{
				Statement: `DROP TABLE TEMP_TIMESTAMP;`,
			},
			{
				Statement: `---
SELECT '2202020-10-05'::date::timestamp;  -- fail`,
				ErrorString: `date out of range for timestamp`,
			},
			{
				Statement: `SELECT '2202020-10-05'::date > '2020-10-05'::timestamp as t;`,
				Results:   []sql.Row{{true}},
			},
			{
				Statement: `SELECT '2020-10-05'::timestamp > '2202020-10-05'::date as f;`,
				Results:   []sql.Row{{false}},
			},
			{
				Statement:   `SELECT '2202020-10-05'::date::timestamptz;  -- fail`,
				ErrorString: `date out of range for timestamp`,
			},
			{
				Statement: `SELECT '2202020-10-05'::date > '2020-10-05'::timestamptz as t;`,
				Results:   []sql.Row{{true}},
			},
			{
				Statement: `SELECT '2020-10-05'::timestamptz > '2202020-10-05'::date as f;`,
				Results:   []sql.Row{{false}},
			},
			{
				Statement: `SELECT '4714-11-24 BC'::date::timestamptz;`,
				Results:   []sql.Row{{`Mon Nov 24 00:00:00 4714 PST BC`}},
			},
			{
				Statement: `SET TimeZone = 'UTC-2';`,
			},
			{
				Statement:   `SELECT '4714-11-24 BC'::date::timestamptz;  -- fail`,
				ErrorString: `date out of range for timestamp`,
			},
			{
				Statement: `SELECT '4714-11-24 BC'::date < '2020-10-05'::timestamptz as t;`,
				Results:   []sql.Row{{true}},
			},
			{
				Statement: `SELECT '2020-10-05'::timestamptz >= '4714-11-24 BC'::date as t;`,
				Results:   []sql.Row{{true}},
			},
			{
				Statement: `SELECT '4714-11-24 BC'::timestamp < '2020-10-05'::timestamptz as t;`,
				Results:   []sql.Row{{true}},
			},
			{
				Statement: `SELECT '2020-10-05'::timestamptz >= '4714-11-24 BC'::timestamp as t;`,
				Results:   []sql.Row{{true}},
			},
			{
				Statement: `RESET TimeZone;`,
			},
			{
				Statement: `explain (costs off)
select count(*) from date_tbl
  where f1 between '1997-01-01' and '1998-01-01';`,
				Results: []sql.Row{{`Aggregate`}, {`->  Seq Scan on date_tbl`}, {`Filter: ((f1 >= '01-01-1997'::date) AND (f1 <= '01-01-1998'::date))`}},
			},
			{
				Statement: `select count(*) from date_tbl
  where f1 between '1997-01-01' and '1998-01-01';`,
				Results: []sql.Row{{3}},
			},
			{
				Statement: `explain (costs off)
select count(*) from date_tbl
  where f1 not between '1997-01-01' and '1998-01-01';`,
				Results: []sql.Row{{`Aggregate`}, {`->  Seq Scan on date_tbl`}, {`Filter: ((f1 < '01-01-1997'::date) OR (f1 > '01-01-1998'::date))`}},
			},
			{
				Statement: `select count(*) from date_tbl
  where f1 not between '1997-01-01' and '1998-01-01';`,
				Results: []sql.Row{{13}},
			},
			{
				Statement: `explain (costs off)
select count(*) from date_tbl
  where f1 between symmetric '1997-01-01' and '1998-01-01';`,
				Results: []sql.Row{{`Aggregate`}, {`->  Seq Scan on date_tbl`}, {`Filter: (((f1 >= '01-01-1997'::date) AND (f1 <= '01-01-1998'::date)) OR ((f1 >= '01-01-1998'::date) AND (f1 <= '01-01-1997'::date)))`}},
			},
			{
				Statement: `select count(*) from date_tbl
  where f1 between symmetric '1997-01-01' and '1998-01-01';`,
				Results: []sql.Row{{3}},
			},
			{
				Statement: `explain (costs off)
select count(*) from date_tbl
  where f1 not between symmetric '1997-01-01' and '1998-01-01';`,
				Results: []sql.Row{{`Aggregate`}, {`->  Seq Scan on date_tbl`}, {`Filter: (((f1 < '01-01-1997'::date) OR (f1 > '01-01-1998'::date)) AND ((f1 < '01-01-1998'::date) OR (f1 > '01-01-1997'::date)))`}},
			},
			{
				Statement: `select count(*) from date_tbl
  where f1 not between symmetric '1997-01-01' and '1998-01-01';`,
				Results: []sql.Row{{13}},
			},
			{
				Statement: `SET DateStyle TO 'US,Postgres';`,
			},
			{
				Statement: `SHOW DateStyle;`,
				Results:   []sql.Row{{`Postgres, MDY`}},
			},
			{
				Statement: `SELECT d1 AS us_postgres FROM TIMESTAMP_TBL;`,
				Results:   []sql.Row{{`-infinity`}, {`infinity`}, {`Thu Jan 01 00:00:00 1970`}, {`Mon Feb 10 17:32:01 1997`}, {`Mon Feb 10 17:32:01 1997`}, {`Mon Feb 10 17:32:02 1997`}, {`Mon Feb 10 17:32:01.4 1997`}, {`Mon Feb 10 17:32:01.5 1997`}, {`Mon Feb 10 17:32:01.6 1997`}, {`Thu Jan 02 00:00:00 1997`}, {`Thu Jan 02 03:04:05 1997`}, {`Mon Feb 10 17:32:01 1997`}, {`Mon Feb 10 17:32:01 1997`}, {`Mon Feb 10 17:32:01 1997`}, {`Mon Feb 10 17:32:01 1997`}, {`Tue Jun 10 17:32:01 1997`}, {`Sat Sep 22 18:19:20 2001`}, {`Wed Mar 15 08:14:01 2000`}, {`Wed Mar 15 13:14:02 2000`}, {`Wed Mar 15 12:14:03 2000`}, {`Wed Mar 15 03:14:04 2000`}, {`Wed Mar 15 02:14:05 2000`}, {`Mon Feb 10 17:32:01 1997`}, {`Mon Feb 10 17:32:01 1997`}, {`Mon Feb 10 17:32:00 1997`}, {`Mon Feb 10 17:32:01 1997`}, {`Mon Feb 10 17:32:01 1997`}, {`Mon Feb 10 17:32:01 1997`}, {`Mon Feb 10 17:32:01 1997`}, {`Mon Feb 10 17:32:01 1997`}, {`Mon Feb 10 17:32:01 1997`}, {`Mon Feb 10 17:32:01 1997`}, {`Mon Feb 10 17:32:01 1997`}, {`Mon Feb 10 17:32:01 1997`}, {`Tue Jun 10 18:32:01 1997`}, {`Mon Feb 10 17:32:01 1997`}, {`Tue Feb 11 17:32:01 1997`}, {`Wed Feb 12 17:32:01 1997`}, {`Thu Feb 13 17:32:01 1997`}, {`Fri Feb 14 17:32:01 1997`}, {`Sat Feb 15 17:32:01 1997`}, {`Sun Feb 16 17:32:01 1997`}, {`Tue Feb 16 17:32:01 0097 BC`}, {`Sat Feb 16 17:32:01 0097`}, {`Thu Feb 16 17:32:01 0597`}, {`Tue Feb 16 17:32:01 1097`}, {`Sat Feb 16 17:32:01 1697`}, {`Thu Feb 16 17:32:01 1797`}, {`Tue Feb 16 17:32:01 1897`}, {`Sun Feb 16 17:32:01 1997`}, {`Sat Feb 16 17:32:01 2097`}, {`Wed Feb 28 17:32:01 1996`}, {`Thu Feb 29 17:32:01 1996`}, {`Fri Mar 01 17:32:01 1996`}, {`Mon Dec 30 17:32:01 1996`}, {`Tue Dec 31 17:32:01 1996`}, {`Wed Jan 01 17:32:01 1997`}, {`Fri Feb 28 17:32:01 1997`}, {`Sat Mar 01 17:32:01 1997`}, {`Tue Dec 30 17:32:01 1997`}, {`Wed Dec 31 17:32:01 1997`}, {`Fri Dec 31 17:32:01 1999`}, {`Sat Jan 01 17:32:01 2000`}, {`Sun Dec 31 17:32:01 2000`}, {`Mon Jan 01 17:32:01 2001`}},
			},
			{
				Statement: `SET DateStyle TO 'US,ISO';`,
			},
			{
				Statement: `SELECT d1 AS us_iso FROM TIMESTAMP_TBL;`,
				Results:   []sql.Row{{`-infinity`}, {`infinity`}, {`1970-01-01 00:00:00`}, {`1997-02-10 17:32:01`}, {`1997-02-10 17:32:01`}, {`1997-02-10 17:32:02`}, {`1997-02-10 17:32:01.4`}, {`1997-02-10 17:32:01.5`}, {`1997-02-10 17:32:01.6`}, {`1997-01-02 00:00:00`}, {`1997-01-02 03:04:05`}, {`1997-02-10 17:32:01`}, {`1997-02-10 17:32:01`}, {`1997-02-10 17:32:01`}, {`1997-02-10 17:32:01`}, {`1997-06-10 17:32:01`}, {`2001-09-22 18:19:20`}, {`2000-03-15 08:14:01`}, {`2000-03-15 13:14:02`}, {`2000-03-15 12:14:03`}, {`2000-03-15 03:14:04`}, {`2000-03-15 02:14:05`}, {`1997-02-10 17:32:01`}, {`1997-02-10 17:32:01`}, {`1997-02-10 17:32:00`}, {`1997-02-10 17:32:01`}, {`1997-02-10 17:32:01`}, {`1997-02-10 17:32:01`}, {`1997-02-10 17:32:01`}, {`1997-02-10 17:32:01`}, {`1997-02-10 17:32:01`}, {`1997-02-10 17:32:01`}, {`1997-02-10 17:32:01`}, {`1997-02-10 17:32:01`}, {`1997-06-10 18:32:01`}, {`1997-02-10 17:32:01`}, {`1997-02-11 17:32:01`}, {`1997-02-12 17:32:01`}, {`1997-02-13 17:32:01`}, {`1997-02-14 17:32:01`}, {`1997-02-15 17:32:01`}, {`1997-02-16 17:32:01`}, {`0097-02-16 17:32:01 BC`}, {`0097-02-16 17:32:01`}, {`0597-02-16 17:32:01`}, {`1097-02-16 17:32:01`}, {`1697-02-16 17:32:01`}, {`1797-02-16 17:32:01`}, {`1897-02-16 17:32:01`}, {`1997-02-16 17:32:01`}, {`2097-02-16 17:32:01`}, {`1996-02-28 17:32:01`}, {`1996-02-29 17:32:01`}, {`1996-03-01 17:32:01`}, {`1996-12-30 17:32:01`}, {`1996-12-31 17:32:01`}, {`1997-01-01 17:32:01`}, {`1997-02-28 17:32:01`}, {`1997-03-01 17:32:01`}, {`1997-12-30 17:32:01`}, {`1997-12-31 17:32:01`}, {`1999-12-31 17:32:01`}, {`2000-01-01 17:32:01`}, {`2000-12-31 17:32:01`}, {`2001-01-01 17:32:01`}},
			},
			{
				Statement: `SET DateStyle TO 'US,SQL';`,
			},
			{
				Statement: `SHOW DateStyle;`,
				Results:   []sql.Row{{`SQL, MDY`}},
			},
			{
				Statement: `SELECT d1 AS us_sql FROM TIMESTAMP_TBL;`,
				Results:   []sql.Row{{`-infinity`}, {`infinity`}, {`01/01/1970 00:00:00`}, {`02/10/1997 17:32:01`}, {`02/10/1997 17:32:01`}, {`02/10/1997 17:32:02`}, {`02/10/1997 17:32:01.4`}, {`02/10/1997 17:32:01.5`}, {`02/10/1997 17:32:01.6`}, {`01/02/1997 00:00:00`}, {`01/02/1997 03:04:05`}, {`02/10/1997 17:32:01`}, {`02/10/1997 17:32:01`}, {`02/10/1997 17:32:01`}, {`02/10/1997 17:32:01`}, {`06/10/1997 17:32:01`}, {`09/22/2001 18:19:20`}, {`03/15/2000 08:14:01`}, {`03/15/2000 13:14:02`}, {`03/15/2000 12:14:03`}, {`03/15/2000 03:14:04`}, {`03/15/2000 02:14:05`}, {`02/10/1997 17:32:01`}, {`02/10/1997 17:32:01`}, {`02/10/1997 17:32:00`}, {`02/10/1997 17:32:01`}, {`02/10/1997 17:32:01`}, {`02/10/1997 17:32:01`}, {`02/10/1997 17:32:01`}, {`02/10/1997 17:32:01`}, {`02/10/1997 17:32:01`}, {`02/10/1997 17:32:01`}, {`02/10/1997 17:32:01`}, {`02/10/1997 17:32:01`}, {`06/10/1997 18:32:01`}, {`02/10/1997 17:32:01`}, {`02/11/1997 17:32:01`}, {`02/12/1997 17:32:01`}, {`02/13/1997 17:32:01`}, {`02/14/1997 17:32:01`}, {`02/15/1997 17:32:01`}, {`02/16/1997 17:32:01`}, {`02/16/0097 17:32:01 BC`}, {`02/16/0097 17:32:01`}, {`02/16/0597 17:32:01`}, {`02/16/1097 17:32:01`}, {`02/16/1697 17:32:01`}, {`02/16/1797 17:32:01`}, {`02/16/1897 17:32:01`}, {`02/16/1997 17:32:01`}, {`02/16/2097 17:32:01`}, {`02/28/1996 17:32:01`}, {`02/29/1996 17:32:01`}, {`03/01/1996 17:32:01`}, {`12/30/1996 17:32:01`}, {`12/31/1996 17:32:01`}, {`01/01/1997 17:32:01`}, {`02/28/1997 17:32:01`}, {`03/01/1997 17:32:01`}, {`12/30/1997 17:32:01`}, {`12/31/1997 17:32:01`}, {`12/31/1999 17:32:01`}, {`01/01/2000 17:32:01`}, {`12/31/2000 17:32:01`}, {`01/01/2001 17:32:01`}},
			},
			{
				Statement: `SET DateStyle TO 'European,Postgres';`,
			},
			{
				Statement: `SHOW DateStyle;`,
				Results:   []sql.Row{{`Postgres, DMY`}},
			},
			{
				Statement: `INSERT INTO TIMESTAMP_TBL VALUES('13/06/1957');`,
			},
			{
				Statement: `SELECT count(*) as one FROM TIMESTAMP_TBL WHERE d1 = 'Jun 13 1957';`,
				Results:   []sql.Row{{1}},
			},
			{
				Statement: `SELECT d1 AS european_postgres FROM TIMESTAMP_TBL;`,
				Results:   []sql.Row{{`-infinity`}, {`infinity`}, {`Thu 01 Jan 00:00:00 1970`}, {`Mon 10 Feb 17:32:01 1997`}, {`Mon 10 Feb 17:32:01 1997`}, {`Mon 10 Feb 17:32:02 1997`}, {`Mon 10 Feb 17:32:01.4 1997`}, {`Mon 10 Feb 17:32:01.5 1997`}, {`Mon 10 Feb 17:32:01.6 1997`}, {`Thu 02 Jan 00:00:00 1997`}, {`Thu 02 Jan 03:04:05 1997`}, {`Mon 10 Feb 17:32:01 1997`}, {`Mon 10 Feb 17:32:01 1997`}, {`Mon 10 Feb 17:32:01 1997`}, {`Mon 10 Feb 17:32:01 1997`}, {`Tue 10 Jun 17:32:01 1997`}, {`Sat 22 Sep 18:19:20 2001`}, {`Wed 15 Mar 08:14:01 2000`}, {`Wed 15 Mar 13:14:02 2000`}, {`Wed 15 Mar 12:14:03 2000`}, {`Wed 15 Mar 03:14:04 2000`}, {`Wed 15 Mar 02:14:05 2000`}, {`Mon 10 Feb 17:32:01 1997`}, {`Mon 10 Feb 17:32:01 1997`}, {`Mon 10 Feb 17:32:00 1997`}, {`Mon 10 Feb 17:32:01 1997`}, {`Mon 10 Feb 17:32:01 1997`}, {`Mon 10 Feb 17:32:01 1997`}, {`Mon 10 Feb 17:32:01 1997`}, {`Mon 10 Feb 17:32:01 1997`}, {`Mon 10 Feb 17:32:01 1997`}, {`Mon 10 Feb 17:32:01 1997`}, {`Mon 10 Feb 17:32:01 1997`}, {`Mon 10 Feb 17:32:01 1997`}, {`Tue 10 Jun 18:32:01 1997`}, {`Mon 10 Feb 17:32:01 1997`}, {`Tue 11 Feb 17:32:01 1997`}, {`Wed 12 Feb 17:32:01 1997`}, {`Thu 13 Feb 17:32:01 1997`}, {`Fri 14 Feb 17:32:01 1997`}, {`Sat 15 Feb 17:32:01 1997`}, {`Sun 16 Feb 17:32:01 1997`}, {`Tue 16 Feb 17:32:01 0097 BC`}, {`Sat 16 Feb 17:32:01 0097`}, {`Thu 16 Feb 17:32:01 0597`}, {`Tue 16 Feb 17:32:01 1097`}, {`Sat 16 Feb 17:32:01 1697`}, {`Thu 16 Feb 17:32:01 1797`}, {`Tue 16 Feb 17:32:01 1897`}, {`Sun 16 Feb 17:32:01 1997`}, {`Sat 16 Feb 17:32:01 2097`}, {`Wed 28 Feb 17:32:01 1996`}, {`Thu 29 Feb 17:32:01 1996`}, {`Fri 01 Mar 17:32:01 1996`}, {`Mon 30 Dec 17:32:01 1996`}, {`Tue 31 Dec 17:32:01 1996`}, {`Wed 01 Jan 17:32:01 1997`}, {`Fri 28 Feb 17:32:01 1997`}, {`Sat 01 Mar 17:32:01 1997`}, {`Tue 30 Dec 17:32:01 1997`}, {`Wed 31 Dec 17:32:01 1997`}, {`Fri 31 Dec 17:32:01 1999`}, {`Sat 01 Jan 17:32:01 2000`}, {`Sun 31 Dec 17:32:01 2000`}, {`Mon 01 Jan 17:32:01 2001`}, {`Thu 13 Jun 00:00:00 1957`}},
			},
			{
				Statement: `SET DateStyle TO 'European,ISO';`,
			},
			{
				Statement: `SHOW DateStyle;`,
				Results:   []sql.Row{{`ISO, DMY`}},
			},
			{
				Statement: `SELECT d1 AS european_iso FROM TIMESTAMP_TBL;`,
				Results:   []sql.Row{{`-infinity`}, {`infinity`}, {`1970-01-01 00:00:00`}, {`1997-02-10 17:32:01`}, {`1997-02-10 17:32:01`}, {`1997-02-10 17:32:02`}, {`1997-02-10 17:32:01.4`}, {`1997-02-10 17:32:01.5`}, {`1997-02-10 17:32:01.6`}, {`1997-01-02 00:00:00`}, {`1997-01-02 03:04:05`}, {`1997-02-10 17:32:01`}, {`1997-02-10 17:32:01`}, {`1997-02-10 17:32:01`}, {`1997-02-10 17:32:01`}, {`1997-06-10 17:32:01`}, {`2001-09-22 18:19:20`}, {`2000-03-15 08:14:01`}, {`2000-03-15 13:14:02`}, {`2000-03-15 12:14:03`}, {`2000-03-15 03:14:04`}, {`2000-03-15 02:14:05`}, {`1997-02-10 17:32:01`}, {`1997-02-10 17:32:01`}, {`1997-02-10 17:32:00`}, {`1997-02-10 17:32:01`}, {`1997-02-10 17:32:01`}, {`1997-02-10 17:32:01`}, {`1997-02-10 17:32:01`}, {`1997-02-10 17:32:01`}, {`1997-02-10 17:32:01`}, {`1997-02-10 17:32:01`}, {`1997-02-10 17:32:01`}, {`1997-02-10 17:32:01`}, {`1997-06-10 18:32:01`}, {`1997-02-10 17:32:01`}, {`1997-02-11 17:32:01`}, {`1997-02-12 17:32:01`}, {`1997-02-13 17:32:01`}, {`1997-02-14 17:32:01`}, {`1997-02-15 17:32:01`}, {`1997-02-16 17:32:01`}, {`0097-02-16 17:32:01 BC`}, {`0097-02-16 17:32:01`}, {`0597-02-16 17:32:01`}, {`1097-02-16 17:32:01`}, {`1697-02-16 17:32:01`}, {`1797-02-16 17:32:01`}, {`1897-02-16 17:32:01`}, {`1997-02-16 17:32:01`}, {`2097-02-16 17:32:01`}, {`1996-02-28 17:32:01`}, {`1996-02-29 17:32:01`}, {`1996-03-01 17:32:01`}, {`1996-12-30 17:32:01`}, {`1996-12-31 17:32:01`}, {`1997-01-01 17:32:01`}, {`1997-02-28 17:32:01`}, {`1997-03-01 17:32:01`}, {`1997-12-30 17:32:01`}, {`1997-12-31 17:32:01`}, {`1999-12-31 17:32:01`}, {`2000-01-01 17:32:01`}, {`2000-12-31 17:32:01`}, {`2001-01-01 17:32:01`}, {`1957-06-13 00:00:00`}},
			},
			{
				Statement: `SET DateStyle TO 'European,SQL';`,
			},
			{
				Statement: `SHOW DateStyle;`,
				Results:   []sql.Row{{`SQL, DMY`}},
			},
			{
				Statement: `SELECT d1 AS european_sql FROM TIMESTAMP_TBL;`,
				Results:   []sql.Row{{`-infinity`}, {`infinity`}, {`01/01/1970 00:00:00`}, {`10/02/1997 17:32:01`}, {`10/02/1997 17:32:01`}, {`10/02/1997 17:32:02`}, {`10/02/1997 17:32:01.4`}, {`10/02/1997 17:32:01.5`}, {`10/02/1997 17:32:01.6`}, {`02/01/1997 00:00:00`}, {`02/01/1997 03:04:05`}, {`10/02/1997 17:32:01`}, {`10/02/1997 17:32:01`}, {`10/02/1997 17:32:01`}, {`10/02/1997 17:32:01`}, {`10/06/1997 17:32:01`}, {`22/09/2001 18:19:20`}, {`15/03/2000 08:14:01`}, {`15/03/2000 13:14:02`}, {`15/03/2000 12:14:03`}, {`15/03/2000 03:14:04`}, {`15/03/2000 02:14:05`}, {`10/02/1997 17:32:01`}, {`10/02/1997 17:32:01`}, {`10/02/1997 17:32:00`}, {`10/02/1997 17:32:01`}, {`10/02/1997 17:32:01`}, {`10/02/1997 17:32:01`}, {`10/02/1997 17:32:01`}, {`10/02/1997 17:32:01`}, {`10/02/1997 17:32:01`}, {`10/02/1997 17:32:01`}, {`10/02/1997 17:32:01`}, {`10/02/1997 17:32:01`}, {`10/06/1997 18:32:01`}, {`10/02/1997 17:32:01`}, {`11/02/1997 17:32:01`}, {`12/02/1997 17:32:01`}, {`13/02/1997 17:32:01`}, {`14/02/1997 17:32:01`}, {`15/02/1997 17:32:01`}, {`16/02/1997 17:32:01`}, {`16/02/0097 17:32:01 BC`}, {`16/02/0097 17:32:01`}, {`16/02/0597 17:32:01`}, {`16/02/1097 17:32:01`}, {`16/02/1697 17:32:01`}, {`16/02/1797 17:32:01`}, {`16/02/1897 17:32:01`}, {`16/02/1997 17:32:01`}, {`16/02/2097 17:32:01`}, {`28/02/1996 17:32:01`}, {`29/02/1996 17:32:01`}, {`01/03/1996 17:32:01`}, {`30/12/1996 17:32:01`}, {`31/12/1996 17:32:01`}, {`01/01/1997 17:32:01`}, {`28/02/1997 17:32:01`}, {`01/03/1997 17:32:01`}, {`30/12/1997 17:32:01`}, {`31/12/1997 17:32:01`}, {`31/12/1999 17:32:01`}, {`01/01/2000 17:32:01`}, {`31/12/2000 17:32:01`}, {`01/01/2001 17:32:01`}, {`13/06/1957 00:00:00`}},
			},
			{
				Statement: `RESET DateStyle;`,
			},
			{
				Statement: `SELECT to_timestamp('0097/Feb/16 --> 08:14:30', 'YYYY/Mon/DD --> HH:MI:SS');`,
				Results:   []sql.Row{{`Sat Feb 16 08:14:30 0097 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('97/2/16 8:14:30', 'FMYYYY/FMMM/FMDD FMHH:FMMI:FMSS');`,
				Results:   []sql.Row{{`Sat Feb 16 08:14:30 0097 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('2011$03!18 23_38_15', 'YYYY-MM-DD HH24:MI:SS');`,
				Results:   []sql.Row{{`Fri Mar 18 23:38:15 2011 PDT`}},
			},
			{
				Statement: `SELECT to_timestamp('1985 January 12', 'YYYY FMMonth DD');`,
				Results:   []sql.Row{{`Sat Jan 12 00:00:00 1985 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('1985 FMMonth 12', 'YYYY "FMMonth" DD');`,
				Results:   []sql.Row{{`Sat Jan 12 00:00:00 1985 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('1985 \ 12', 'YYYY \\ DD');`,
				Results:   []sql.Row{{`Sat Jan 12 00:00:00 1985 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('My birthday-> Year: 1976, Month: May, Day: 16',
                    '"My birthday-> Year:" YYYY, "Month:" FMMonth, "Day:" DD');`,
				Results: []sql.Row{{`Sun May 16 00:00:00 1976 PDT`}},
			},
			{
				Statement: `SELECT to_timestamp('1,582nd VIII 21', 'Y,YYYth FMRM DD');`,
				Results:   []sql.Row{{`Sat Aug 21 00:00:00 1582 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('15 "text between quote marks" 98 54 45',
                    E'HH24 "\\"text between quote marks\\"" YY MI SS');`,
				Results: []sql.Row{{`Thu Jan 01 15:54:45 1998 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('05121445482000', 'MMDDHH24MISSYYYY');`,
				Results:   []sql.Row{{`Fri May 12 14:45:48 2000 PDT`}},
			},
			{
				Statement: `SELECT to_timestamp('2000January09Sunday', 'YYYYFMMonthDDFMDay');`,
				Results:   []sql.Row{{`Sun Jan 09 00:00:00 2000 PST`}},
			},
			{
				Statement:   `SELECT to_timestamp('97/Feb/16', 'YYMonDD');`,
				ErrorString: `invalid value "/Feb/16" for "Mon"`,
			},
			{
				Statement: `SELECT to_timestamp('97/Feb/16', 'YY:Mon:DD');`,
				Results:   []sql.Row{{`Sun Feb 16 00:00:00 1997 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('97/Feb/16', 'FXYY:Mon:DD');`,
				Results:   []sql.Row{{`Sun Feb 16 00:00:00 1997 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('97/Feb/16', 'FXYY/Mon/DD');`,
				Results:   []sql.Row{{`Sun Feb 16 00:00:00 1997 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('19971116', 'YYYYMMDD');`,
				Results:   []sql.Row{{`Sun Nov 16 00:00:00 1997 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('20000-1116', 'YYYY-MMDD');`,
				Results:   []sql.Row{{`Thu Nov 16 00:00:00 20000 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('1997 AD 11 16', 'YYYY BC MM DD');`,
				Results:   []sql.Row{{`Sun Nov 16 00:00:00 1997 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('1997 BC 11 16', 'YYYY BC MM DD');`,
				Results:   []sql.Row{{`Tue Nov 16 00:00:00 1997 PST BC`}},
			},
			{
				Statement: `SELECT to_timestamp('1997 A.D. 11 16', 'YYYY B.C. MM DD');`,
				Results:   []sql.Row{{`Sun Nov 16 00:00:00 1997 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('1997 B.C. 11 16', 'YYYY B.C. MM DD');`,
				Results:   []sql.Row{{`Tue Nov 16 00:00:00 1997 PST BC`}},
			},
			{
				Statement: `SELECT to_timestamp('9-1116', 'Y-MMDD');`,
				Results:   []sql.Row{{`Mon Nov 16 00:00:00 2009 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('95-1116', 'YY-MMDD');`,
				Results:   []sql.Row{{`Thu Nov 16 00:00:00 1995 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('995-1116', 'YYY-MMDD');`,
				Results:   []sql.Row{{`Thu Nov 16 00:00:00 1995 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('2005426', 'YYYYWWD');`,
				Results:   []sql.Row{{`Sat Oct 15 00:00:00 2005 PDT`}},
			},
			{
				Statement: `SELECT to_timestamp('2005300', 'YYYYDDD');`,
				Results:   []sql.Row{{`Thu Oct 27 00:00:00 2005 PDT`}},
			},
			{
				Statement: `SELECT to_timestamp('2005527', 'IYYYIWID');`,
				Results:   []sql.Row{{`Sun Jan 01 00:00:00 2006 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('005527', 'IYYIWID');`,
				Results:   []sql.Row{{`Sun Jan 01 00:00:00 2006 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('05527', 'IYIWID');`,
				Results:   []sql.Row{{`Sun Jan 01 00:00:00 2006 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('5527', 'IIWID');`,
				Results:   []sql.Row{{`Sun Jan 01 00:00:00 2006 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('2005364', 'IYYYIDDD');`,
				Results:   []sql.Row{{`Sun Jan 01 00:00:00 2006 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('20050302', 'YYYYMMDD');`,
				Results:   []sql.Row{{`Wed Mar 02 00:00:00 2005 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('2005 03 02', 'YYYYMMDD');`,
				Results:   []sql.Row{{`Wed Mar 02 00:00:00 2005 PST`}},
			},
			{
				Statement: `SELECT to_timestamp(' 2005 03 02', 'YYYYMMDD');`,
				Results:   []sql.Row{{`Wed Mar 02 00:00:00 2005 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('  20050302', 'YYYYMMDD');`,
				Results:   []sql.Row{{`Wed Mar 02 00:00:00 2005 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('2011-12-18 11:38 AM', 'YYYY-MM-DD HH12:MI PM');`,
				Results:   []sql.Row{{`Sun Dec 18 11:38:00 2011 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('2011-12-18 11:38 PM', 'YYYY-MM-DD HH12:MI PM');`,
				Results:   []sql.Row{{`Sun Dec 18 23:38:00 2011 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('2011-12-18 11:38 A.M.', 'YYYY-MM-DD HH12:MI P.M.');`,
				Results:   []sql.Row{{`Sun Dec 18 11:38:00 2011 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('2011-12-18 11:38 P.M.', 'YYYY-MM-DD HH12:MI P.M.');`,
				Results:   []sql.Row{{`Sun Dec 18 23:38:00 2011 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('2011-12-18 11:38 +05',    'YYYY-MM-DD HH12:MI TZH');`,
				Results:   []sql.Row{{`Sat Dec 17 22:38:00 2011 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('2011-12-18 11:38 -05',    'YYYY-MM-DD HH12:MI TZH');`,
				Results:   []sql.Row{{`Sun Dec 18 08:38:00 2011 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('2011-12-18 11:38 +05:20', 'YYYY-MM-DD HH12:MI TZH:TZM');`,
				Results:   []sql.Row{{`Sat Dec 17 22:18:00 2011 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('2011-12-18 11:38 -05:20', 'YYYY-MM-DD HH12:MI TZH:TZM');`,
				Results:   []sql.Row{{`Sun Dec 18 08:58:00 2011 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('2011-12-18 11:38 20',     'YYYY-MM-DD HH12:MI TZM');`,
				Results:   []sql.Row{{`Sun Dec 18 03:18:00 2011 PST`}},
			},
			{
				Statement:   `SELECT to_timestamp('2011-12-18 11:38 PST', 'YYYY-MM-DD HH12:MI TZ');  -- NYI`,
				ErrorString: `formatting field "TZ" is only supported in to_char`,
			},
			{
				Statement: `SELECT to_timestamp('2018-11-02 12:34:56.025', 'YYYY-MM-DD HH24:MI:SS.MS');`,
				Results:   []sql.Row{{`Fri Nov 02 12:34:56.025 2018 PDT`}},
			},
			{
				Statement: `SELECT i, to_timestamp('2018-11-02 12:34:56', 'YYYY-MM-DD HH24:MI:SS.FF' || i) FROM generate_series(1, 6) i;`,
				Results:   []sql.Row{{1, `Fri Nov 02 12:34:56 2018 PDT`}, {2, `Fri Nov 02 12:34:56 2018 PDT`}, {3, `Fri Nov 02 12:34:56 2018 PDT`}, {4, `Fri Nov 02 12:34:56 2018 PDT`}, {5, `Fri Nov 02 12:34:56 2018 PDT`}, {6, `Fri Nov 02 12:34:56 2018 PDT`}},
			},
			{
				Statement: `SELECT i, to_timestamp('2018-11-02 12:34:56.1', 'YYYY-MM-DD HH24:MI:SS.FF' || i) FROM generate_series(1, 6) i;`,
				Results:   []sql.Row{{1, `Fri Nov 02 12:34:56.1 2018 PDT`}, {2, `Fri Nov 02 12:34:56.1 2018 PDT`}, {3, `Fri Nov 02 12:34:56.1 2018 PDT`}, {4, `Fri Nov 02 12:34:56.1 2018 PDT`}, {5, `Fri Nov 02 12:34:56.1 2018 PDT`}, {6, `Fri Nov 02 12:34:56.1 2018 PDT`}},
			},
			{
				Statement: `SELECT i, to_timestamp('2018-11-02 12:34:56.12', 'YYYY-MM-DD HH24:MI:SS.FF' || i) FROM generate_series(1, 6) i;`,
				Results:   []sql.Row{{1, `Fri Nov 02 12:34:56.1 2018 PDT`}, {2, `Fri Nov 02 12:34:56.12 2018 PDT`}, {3, `Fri Nov 02 12:34:56.12 2018 PDT`}, {4, `Fri Nov 02 12:34:56.12 2018 PDT`}, {5, `Fri Nov 02 12:34:56.12 2018 PDT`}, {6, `Fri Nov 02 12:34:56.12 2018 PDT`}},
			},
			{
				Statement: `SELECT i, to_timestamp('2018-11-02 12:34:56.123', 'YYYY-MM-DD HH24:MI:SS.FF' || i) FROM generate_series(1, 6) i;`,
				Results:   []sql.Row{{1, `Fri Nov 02 12:34:56.1 2018 PDT`}, {2, `Fri Nov 02 12:34:56.12 2018 PDT`}, {3, `Fri Nov 02 12:34:56.123 2018 PDT`}, {4, `Fri Nov 02 12:34:56.123 2018 PDT`}, {5, `Fri Nov 02 12:34:56.123 2018 PDT`}, {6, `Fri Nov 02 12:34:56.123 2018 PDT`}},
			},
			{
				Statement: `SELECT i, to_timestamp('2018-11-02 12:34:56.1234', 'YYYY-MM-DD HH24:MI:SS.FF' || i) FROM generate_series(1, 6) i;`,
				Results:   []sql.Row{{1, `Fri Nov 02 12:34:56.1 2018 PDT`}, {2, `Fri Nov 02 12:34:56.12 2018 PDT`}, {3, `Fri Nov 02 12:34:56.123 2018 PDT`}, {4, `Fri Nov 02 12:34:56.1234 2018 PDT`}, {5, `Fri Nov 02 12:34:56.1234 2018 PDT`}, {6, `Fri Nov 02 12:34:56.1234 2018 PDT`}},
			},
			{
				Statement: `SELECT i, to_timestamp('2018-11-02 12:34:56.12345', 'YYYY-MM-DD HH24:MI:SS.FF' || i) FROM generate_series(1, 6) i;`,
				Results:   []sql.Row{{1, `Fri Nov 02 12:34:56.1 2018 PDT`}, {2, `Fri Nov 02 12:34:56.12 2018 PDT`}, {3, `Fri Nov 02 12:34:56.123 2018 PDT`}, {4, `Fri Nov 02 12:34:56.1235 2018 PDT`}, {5, `Fri Nov 02 12:34:56.12345 2018 PDT`}, {6, `Fri Nov 02 12:34:56.12345 2018 PDT`}},
			},
			{
				Statement: `SELECT i, to_timestamp('2018-11-02 12:34:56.123456', 'YYYY-MM-DD HH24:MI:SS.FF' || i) FROM generate_series(1, 6) i;`,
				Results:   []sql.Row{{1, `Fri Nov 02 12:34:56.1 2018 PDT`}, {2, `Fri Nov 02 12:34:56.12 2018 PDT`}, {3, `Fri Nov 02 12:34:56.123 2018 PDT`}, {4, `Fri Nov 02 12:34:56.1235 2018 PDT`}, {5, `Fri Nov 02 12:34:56.12346 2018 PDT`}, {6, `Fri Nov 02 12:34:56.123456 2018 PDT`}},
			},
			{
				Statement:   `SELECT i, to_timestamp('2018-11-02 12:34:56.123456789', 'YYYY-MM-DD HH24:MI:SS.FF' || i) FROM generate_series(1, 6) i;`,
				ErrorString: `date/time field value out of range: "2018-11-02 12:34:56.123456789"`,
			},
			{
				Statement: `SELECT to_date('1 4 1902', 'Q MM YYYY');  -- Q is ignored`,
				Results:   []sql.Row{{`04-01-1902`}},
			},
			{
				Statement: `SELECT to_date('3 4 21 01', 'W MM CC YY');`,
				Results:   []sql.Row{{`04-15-2001`}},
			},
			{
				Statement: `SELECT to_date('2458872', 'J');`,
				Results:   []sql.Row{{`01-23-2020`}},
			},
			{
				Statement: `SELECT to_date('44-02-01 BC','YYYY-MM-DD BC');`,
				Results:   []sql.Row{{`02-01-0044 BC`}},
			},
			{
				Statement: `SELECT to_date('-44-02-01','YYYY-MM-DD');`,
				Results:   []sql.Row{{`02-01-0044 BC`}},
			},
			{
				Statement: `SELECT to_date('-44-02-01 BC','YYYY-MM-DD BC');`,
				Results:   []sql.Row{{`02-01-0044`}},
			},
			{
				Statement: `SELECT to_timestamp('44-02-01 11:12:13 BC','YYYY-MM-DD HH24:MI:SS BC');`,
				Results:   []sql.Row{{`Fri Feb 01 11:12:13 0044 PST BC`}},
			},
			{
				Statement: `SELECT to_timestamp('-44-02-01 11:12:13','YYYY-MM-DD HH24:MI:SS');`,
				Results:   []sql.Row{{`Fri Feb 01 11:12:13 0044 PST BC`}},
			},
			{
				Statement: `SELECT to_timestamp('-44-02-01 11:12:13 BC','YYYY-MM-DD HH24:MI:SS BC');`,
				Results:   []sql.Row{{`Mon Feb 01 11:12:13 0044 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('2011-12-18 23:38:15', 'YYYY-MM-DD  HH24:MI:SS');`,
				Results:   []sql.Row{{`Sun Dec 18 23:38:15 2011 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('2011-12-18  23:38:15', 'YYYY-MM-DD  HH24:MI:SS');`,
				Results:   []sql.Row{{`Sun Dec 18 23:38:15 2011 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('2011-12-18   23:38:15', 'YYYY-MM-DD  HH24:MI:SS');`,
				Results:   []sql.Row{{`Sun Dec 18 23:38:15 2011 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('2011-12-18  23:38:15', 'YYYY-MM-DD HH24:MI:SS');`,
				Results:   []sql.Row{{`Sun Dec 18 23:38:15 2011 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('2011-12-18  23:38:15', 'YYYY-MM-DD  HH24:MI:SS');`,
				Results:   []sql.Row{{`Sun Dec 18 23:38:15 2011 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('2011-12-18  23:38:15', 'YYYY-MM-DD   HH24:MI:SS');`,
				Results:   []sql.Row{{`Sun Dec 18 23:38:15 2011 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('2000+   JUN', 'YYYY/MON');`,
				Results:   []sql.Row{{`Thu Jun 01 00:00:00 2000 PDT`}},
			},
			{
				Statement: `SELECT to_timestamp('  2000 +JUN', 'YYYY/MON');`,
				Results:   []sql.Row{{`Thu Jun 01 00:00:00 2000 PDT`}},
			},
			{
				Statement: `SELECT to_timestamp(' 2000 +JUN', 'YYYY//MON');`,
				Results:   []sql.Row{{`Thu Jun 01 00:00:00 2000 PDT`}},
			},
			{
				Statement: `SELECT to_timestamp('2000  +JUN', 'YYYY//MON');`,
				Results:   []sql.Row{{`Thu Jun 01 00:00:00 2000 PDT`}},
			},
			{
				Statement: `SELECT to_timestamp('2000 + JUN', 'YYYY MON');`,
				Results:   []sql.Row{{`Thu Jun 01 00:00:00 2000 PDT`}},
			},
			{
				Statement: `SELECT to_timestamp('2000 ++ JUN', 'YYYY  MON');`,
				Results:   []sql.Row{{`Thu Jun 01 00:00:00 2000 PDT`}},
			},
			{
				Statement:   `SELECT to_timestamp('2000 + + JUN', 'YYYY  MON');`,
				ErrorString: `invalid value "+" for "MON"`,
			},
			{
				Statement: `SELECT to_timestamp('2000 + + JUN', 'YYYY   MON');`,
				Results:   []sql.Row{{`Thu Jun 01 00:00:00 2000 PDT`}},
			},
			{
				Statement: `SELECT to_timestamp('2000 -10', 'YYYY TZH');`,
				Results:   []sql.Row{{`Sat Jan 01 02:00:00 2000 PST`}},
			},
			{
				Statement: `SELECT to_timestamp('2000 -10', 'YYYY  TZH');`,
				Results:   []sql.Row{{`Fri Dec 31 06:00:00 1999 PST`}},
			},
			{
				Statement: `SELECT to_date('2011 12  18', 'YYYY MM DD');`,
				Results:   []sql.Row{{`12-18-2011`}},
			},
			{
				Statement: `SELECT to_date('2011 12  18', 'YYYY MM  DD');`,
				Results:   []sql.Row{{`12-18-2011`}},
			},
			{
				Statement: `SELECT to_date('2011 12  18', 'YYYY MM   DD');`,
				Results:   []sql.Row{{`12-18-2011`}},
			},
			{
				Statement: `SELECT to_date('2011 12 18', 'YYYY  MM DD');`,
				Results:   []sql.Row{{`12-18-2011`}},
			},
			{
				Statement: `SELECT to_date('2011  12 18', 'YYYY  MM DD');`,
				Results:   []sql.Row{{`12-18-2011`}},
			},
			{
				Statement: `SELECT to_date('2011   12 18', 'YYYY  MM DD');`,
				Results:   []sql.Row{{`12-18-2011`}},
			},
			{
				Statement: `SELECT to_date('2011 12 18', 'YYYYxMMxDD');`,
				Results:   []sql.Row{{`12-18-2011`}},
			},
			{
				Statement: `SELECT to_date('2011x 12x 18', 'YYYYxMMxDD');`,
				Results:   []sql.Row{{`12-18-2011`}},
			},
			{
				Statement:   `SELECT to_date('2011 x12 x18', 'YYYYxMMxDD');`,
				ErrorString: `invalid value "x1" for "MM"`,
			},
			{
				Statement:   `SELECT to_timestamp('2005527', 'YYYYIWID');`,
				ErrorString: `invalid combination of date conventions`,
			},
			{
				Statement:   `SELECT to_timestamp('19971', 'YYYYMMDD');`,
				ErrorString: `source string too short for "MM" formatting field`,
			},
			{
				Statement:   `SELECT to_timestamp('19971)24', 'YYYYMMDD');`,
				ErrorString: `invalid value "1)" for "MM"`,
			},
			{
				Statement:   `SELECT to_timestamp('Friday 1-January-1999', 'DY DD MON YYYY');`,
				ErrorString: `invalid value "da" for "DD"`,
			},
			{
				Statement:   `SELECT to_timestamp('Fri 1-January-1999', 'DY DD MON YYYY');`,
				ErrorString: `invalid value "uary" for "YYYY"`,
			},
			{
				Statement: `SELECT to_timestamp('Fri 1-Jan-1999', 'DY DD MON YYYY');  -- ok`,
				Results:   []sql.Row{{`Fri Jan 01 00:00:00 1999 PST`}},
			},
			{
				Statement:   `SELECT to_timestamp('1997-11-Jan-16', 'YYYY-MM-Mon-DD');`,
				ErrorString: `conflicting values for "Mon" field in formatting string`,
			},
			{
				Statement:   `SELECT to_timestamp('199711xy', 'YYYYMMDD');`,
				ErrorString: `invalid value "xy" for "DD"`,
			},
			{
				Statement:   `SELECT to_timestamp('10000000000', 'FMYYYY');`,
				ErrorString: `value for "YYYY" in source string is out of range`,
			},
			{
				Statement:   `SELECT to_timestamp('2016-06-13 25:00:00', 'YYYY-MM-DD HH24:MI:SS');`,
				ErrorString: `date/time field value out of range: "2016-06-13 25:00:00"`,
			},
			{
				Statement:   `SELECT to_timestamp('2016-06-13 15:60:00', 'YYYY-MM-DD HH24:MI:SS');`,
				ErrorString: `date/time field value out of range: "2016-06-13 15:60:00"`,
			},
			{
				Statement:   `SELECT to_timestamp('2016-06-13 15:50:60', 'YYYY-MM-DD HH24:MI:SS');`,
				ErrorString: `date/time field value out of range: "2016-06-13 15:50:60"`,
			},
			{
				Statement: `SELECT to_timestamp('2016-06-13 15:50:55', 'YYYY-MM-DD HH24:MI:SS');  -- ok`,
				Results:   []sql.Row{{`Mon Jun 13 15:50:55 2016 PDT`}},
			},
			{
				Statement:   `SELECT to_timestamp('2016-06-13 15:50:55', 'YYYY-MM-DD HH:MI:SS');`,
				ErrorString: `hour "15" is invalid for the 12-hour clock`,
			},
			{
				Statement:   `SELECT to_timestamp('2016-13-01 15:50:55', 'YYYY-MM-DD HH24:MI:SS');`,
				ErrorString: `date/time field value out of range: "2016-13-01 15:50:55"`,
			},
			{
				Statement:   `SELECT to_timestamp('2016-02-30 15:50:55', 'YYYY-MM-DD HH24:MI:SS');`,
				ErrorString: `date/time field value out of range: "2016-02-30 15:50:55"`,
			},
			{
				Statement: `SELECT to_timestamp('2016-02-29 15:50:55', 'YYYY-MM-DD HH24:MI:SS');  -- ok`,
				Results:   []sql.Row{{`Mon Feb 29 15:50:55 2016 PST`}},
			},
			{
				Statement:   `SELECT to_timestamp('2015-02-29 15:50:55', 'YYYY-MM-DD HH24:MI:SS');`,
				ErrorString: `date/time field value out of range: "2015-02-29 15:50:55"`,
			},
			{
				Statement: `SELECT to_timestamp('2015-02-11 86000', 'YYYY-MM-DD SSSS');  -- ok`,
				Results:   []sql.Row{{`Wed Feb 11 23:53:20 2015 PST`}},
			},
			{
				Statement:   `SELECT to_timestamp('2015-02-11 86400', 'YYYY-MM-DD SSSS');`,
				ErrorString: `date/time field value out of range: "2015-02-11 86400"`,
			},
			{
				Statement: `SELECT to_timestamp('2015-02-11 86000', 'YYYY-MM-DD SSSSS');  -- ok`,
				Results:   []sql.Row{{`Wed Feb 11 23:53:20 2015 PST`}},
			},
			{
				Statement:   `SELECT to_timestamp('2015-02-11 86400', 'YYYY-MM-DD SSSSS');`,
				ErrorString: `date/time field value out of range: "2015-02-11 86400"`,
			},
			{
				Statement:   `SELECT to_date('2016-13-10', 'YYYY-MM-DD');`,
				ErrorString: `date/time field value out of range: "2016-13-10"`,
			},
			{
				Statement:   `SELECT to_date('2016-02-30', 'YYYY-MM-DD');`,
				ErrorString: `date/time field value out of range: "2016-02-30"`,
			},
			{
				Statement: `SELECT to_date('2016-02-29', 'YYYY-MM-DD');  -- ok`,
				Results:   []sql.Row{{`02-29-2016`}},
			},
			{
				Statement:   `SELECT to_date('2015-02-29', 'YYYY-MM-DD');`,
				ErrorString: `date/time field value out of range: "2015-02-29"`,
			},
			{
				Statement: `SELECT to_date('2015 365', 'YYYY DDD');  -- ok`,
				Results:   []sql.Row{{`12-31-2015`}},
			},
			{
				Statement:   `SELECT to_date('2015 366', 'YYYY DDD');`,
				ErrorString: `date/time field value out of range: "2015 366"`,
			},
			{
				Statement: `SELECT to_date('2016 365', 'YYYY DDD');  -- ok`,
				Results:   []sql.Row{{`12-30-2016`}},
			},
			{
				Statement: `SELECT to_date('2016 366', 'YYYY DDD');  -- ok`,
				Results:   []sql.Row{{`12-31-2016`}},
			},
			{
				Statement:   `SELECT to_date('2016 367', 'YYYY DDD');`,
				ErrorString: `date/time field value out of range: "2016 367"`,
			},
			{
				Statement: `SELECT to_date('0000-02-01','YYYY-MM-DD');  -- allowed, though it shouldn't be`,
				Results:   []sql.Row{{`02-01-0001 BC`}},
			},
			{
				Statement: `SET TIME ZONE 'America/New_York';`,
			},
			{
				Statement: `SET TIME ZONE '-1.5';`,
			},
			{
				Statement: `SHOW TIME ZONE;`,
				Results:   []sql.Row{{`<-01:30>+01:30`}},
			},
			{
				Statement: `SELECT '2012-12-12 12:00'::timestamptz;`,
				Results:   []sql.Row{{`Wed Dec 12 12:00:00 2012 -01:30`}},
			},
			{
				Statement: `SELECT '2012-12-12 12:00 America/New_York'::timestamptz;`,
				Results:   []sql.Row{{`Wed Dec 12 15:30:00 2012 -01:30`}},
			},
			{
				Statement: `SELECT to_char('2012-12-12 12:00'::timestamptz, 'YYYY-MM-DD HH:MI:SS TZ');`,
				Results:   []sql.Row{{`2012-12-12 12:00:00 -01:30`}},
			},
			{
				Statement: `SELECT to_char('2012-12-12 12:00'::timestamptz, 'YYYY-MM-DD SSSS');`,
				Results:   []sql.Row{{`2012-12-12 43200`}},
			},
			{
				Statement: `SELECT to_char('2012-12-12 12:00'::timestamptz, 'YYYY-MM-DD SSSSS');`,
				Results:   []sql.Row{{`2012-12-12 43200`}},
			},
			{
				Statement: `RESET TIME ZONE;`,
			},
		},
	})
}
