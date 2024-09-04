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

func TestPortals(t *testing.T) {
	t.Skip()
	_ = RunTests(t, RegressionFileName_portals)
}

func init() {
	RegisterRegressionFile(RegressionFile{
		RegressionFileName: RegressionFileName_portals,
		DependsOn:          []RegressionFileName{RegressionFileName_test_setup},
		Statements: []RegressionFileStatement{
			{
				Statement: `BEGIN;`,
			},
			{
				Statement: `DECLARE foo1 SCROLL CURSOR FOR SELECT * FROM tenk1 ORDER BY unique2;`,
			},
			{
				Statement: `DECLARE foo2 SCROLL CURSOR FOR SELECT * FROM tenk2;`,
			},
			{
				Statement: `DECLARE foo3 SCROLL CURSOR FOR SELECT * FROM tenk1 ORDER BY unique2;`,
			},
			{
				Statement: `DECLARE foo4 SCROLL CURSOR FOR SELECT * FROM tenk2;`,
			},
			{
				Statement: `DECLARE foo5 SCROLL CURSOR FOR SELECT * FROM tenk1 ORDER BY unique2;`,
			},
			{
				Statement: `DECLARE foo6 SCROLL CURSOR FOR SELECT * FROM tenk2;`,
			},
			{
				Statement: `DECLARE foo7 SCROLL CURSOR FOR SELECT * FROM tenk1 ORDER BY unique2;`,
			},
			{
				Statement: `DECLARE foo8 SCROLL CURSOR FOR SELECT * FROM tenk2;`,
			},
			{
				Statement: `DECLARE foo9 SCROLL CURSOR FOR SELECT * FROM tenk1 ORDER BY unique2;`,
			},
			{
				Statement: `DECLARE foo10 SCROLL CURSOR FOR SELECT * FROM tenk2;`,
			},
			{
				Statement: `DECLARE foo11 SCROLL CURSOR FOR SELECT * FROM tenk1 ORDER BY unique2;`,
			},
			{
				Statement: `DECLARE foo12 SCROLL CURSOR FOR SELECT * FROM tenk2;`,
			},
			{
				Statement: `DECLARE foo13 SCROLL CURSOR FOR SELECT * FROM tenk1 ORDER BY unique2;`,
			},
			{
				Statement: `DECLARE foo14 SCROLL CURSOR FOR SELECT * FROM tenk2;`,
			},
			{
				Statement: `DECLARE foo15 SCROLL CURSOR FOR SELECT * FROM tenk1 ORDER BY unique2;`,
			},
			{
				Statement: `DECLARE foo16 SCROLL CURSOR FOR SELECT * FROM tenk2;`,
			},
			{
				Statement: `DECLARE foo17 SCROLL CURSOR FOR SELECT * FROM tenk1 ORDER BY unique2;`,
			},
			{
				Statement: `DECLARE foo18 SCROLL CURSOR FOR SELECT * FROM tenk2;`,
			},
			{
				Statement: `DECLARE foo19 SCROLL CURSOR FOR SELECT * FROM tenk1 ORDER BY unique2;`,
			},
			{
				Statement: `DECLARE foo20 SCROLL CURSOR FOR SELECT * FROM tenk2;`,
			},
			{
				Statement: `DECLARE foo21 SCROLL CURSOR FOR SELECT * FROM tenk1 ORDER BY unique2;`,
			},
			{
				Statement: `DECLARE foo22 SCROLL CURSOR FOR SELECT * FROM tenk2;`,
			},
			{
				Statement: `DECLARE foo23 SCROLL CURSOR FOR SELECT * FROM tenk1 ORDER BY unique2;`,
			},
			{
				Statement: `FETCH 1 in foo1;`,
				Results:   []sql.Row{{8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}},
			},
			{
				Statement: `FETCH 2 in foo2;`,
				Results:   []sql.Row{{8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}, {1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}},
			},
			{
				Statement: `FETCH 3 in foo3;`,
				Results:   []sql.Row{{8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}, {1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}, {3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}},
			},
			{
				Statement: `FETCH 4 in foo4;`,
				Results:   []sql.Row{{8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}, {1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}, {3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}, {9850, 3, 0, 2, 0, 10, 50, 850, 1850, 4850, 9850, 100, 101, `WOAAAA`, `DAAAAA`, `VVVVxx`}},
			},
			{
				Statement: `FETCH 5 in foo5;`,
				Results:   []sql.Row{{8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}, {1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}, {3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}, {9850, 3, 0, 2, 0, 10, 50, 850, 1850, 4850, 9850, 100, 101, `WOAAAA`, `DAAAAA`, `VVVVxx`}, {7164, 4, 0, 0, 4, 4, 64, 164, 1164, 2164, 7164, 128, 129, `OPAAAA`, `EAAAAA`, `AAAAxx`}},
			},
			{
				Statement: `FETCH 6 in foo6;`,
				Results:   []sql.Row{{8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}, {1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}, {3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}, {9850, 3, 0, 2, 0, 10, 50, 850, 1850, 4850, 9850, 100, 101, `WOAAAA`, `DAAAAA`, `VVVVxx`}, {7164, 4, 0, 0, 4, 4, 64, 164, 1164, 2164, 7164, 128, 129, `OPAAAA`, `EAAAAA`, `AAAAxx`}, {8009, 5, 1, 1, 9, 9, 9, 9, 9, 3009, 8009, 18, 19, `BWAAAA`, `FAAAAA`, `HHHHxx`}},
			},
			{
				Statement: `FETCH 7 in foo7;`,
				Results:   []sql.Row{{8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}, {1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}, {3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}, {9850, 3, 0, 2, 0, 10, 50, 850, 1850, 4850, 9850, 100, 101, `WOAAAA`, `DAAAAA`, `VVVVxx`}, {7164, 4, 0, 0, 4, 4, 64, 164, 1164, 2164, 7164, 128, 129, `OPAAAA`, `EAAAAA`, `AAAAxx`}, {8009, 5, 1, 1, 9, 9, 9, 9, 9, 3009, 8009, 18, 19, `BWAAAA`, `FAAAAA`, `HHHHxx`}, {5057, 6, 1, 1, 7, 17, 57, 57, 1057, 57, 5057, 114, 115, `NMAAAA`, `GAAAAA`, `OOOOxx`}},
			},
			{
				Statement: `FETCH 8 in foo8;`,
				Results:   []sql.Row{{8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}, {1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}, {3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}, {9850, 3, 0, 2, 0, 10, 50, 850, 1850, 4850, 9850, 100, 101, `WOAAAA`, `DAAAAA`, `VVVVxx`}, {7164, 4, 0, 0, 4, 4, 64, 164, 1164, 2164, 7164, 128, 129, `OPAAAA`, `EAAAAA`, `AAAAxx`}, {8009, 5, 1, 1, 9, 9, 9, 9, 9, 3009, 8009, 18, 19, `BWAAAA`, `FAAAAA`, `HHHHxx`}, {5057, 6, 1, 1, 7, 17, 57, 57, 1057, 57, 5057, 114, 115, `NMAAAA`, `GAAAAA`, `OOOOxx`}, {6701, 7, 1, 1, 1, 1, 1, 701, 701, 1701, 6701, 2, 3, `TXAAAA`, `HAAAAA`, `VVVVxx`}},
			},
			{
				Statement: `FETCH 9 in foo9;`,
				Results:   []sql.Row{{8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}, {1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}, {3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}, {9850, 3, 0, 2, 0, 10, 50, 850, 1850, 4850, 9850, 100, 101, `WOAAAA`, `DAAAAA`, `VVVVxx`}, {7164, 4, 0, 0, 4, 4, 64, 164, 1164, 2164, 7164, 128, 129, `OPAAAA`, `EAAAAA`, `AAAAxx`}, {8009, 5, 1, 1, 9, 9, 9, 9, 9, 3009, 8009, 18, 19, `BWAAAA`, `FAAAAA`, `HHHHxx`}, {5057, 6, 1, 1, 7, 17, 57, 57, 1057, 57, 5057, 114, 115, `NMAAAA`, `GAAAAA`, `OOOOxx`}, {6701, 7, 1, 1, 1, 1, 1, 701, 701, 1701, 6701, 2, 3, `TXAAAA`, `HAAAAA`, `VVVVxx`}, {4321, 8, 1, 1, 1, 1, 21, 321, 321, 4321, 4321, 42, 43, `FKAAAA`, `IAAAAA`, `AAAAxx`}},
			},
			{
				Statement: `FETCH 10 in foo10;`,
				Results:   []sql.Row{{8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}, {1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}, {3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}, {9850, 3, 0, 2, 0, 10, 50, 850, 1850, 4850, 9850, 100, 101, `WOAAAA`, `DAAAAA`, `VVVVxx`}, {7164, 4, 0, 0, 4, 4, 64, 164, 1164, 2164, 7164, 128, 129, `OPAAAA`, `EAAAAA`, `AAAAxx`}, {8009, 5, 1, 1, 9, 9, 9, 9, 9, 3009, 8009, 18, 19, `BWAAAA`, `FAAAAA`, `HHHHxx`}, {5057, 6, 1, 1, 7, 17, 57, 57, 1057, 57, 5057, 114, 115, `NMAAAA`, `GAAAAA`, `OOOOxx`}, {6701, 7, 1, 1, 1, 1, 1, 701, 701, 1701, 6701, 2, 3, `TXAAAA`, `HAAAAA`, `VVVVxx`}, {4321, 8, 1, 1, 1, 1, 21, 321, 321, 4321, 4321, 42, 43, `FKAAAA`, `IAAAAA`, `AAAAxx`}, {3043, 9, 1, 3, 3, 3, 43, 43, 1043, 3043, 3043, 86, 87, `BNAAAA`, `JAAAAA`, `HHHHxx`}},
			},
			{
				Statement: `FETCH 11 in foo11;`,
				Results:   []sql.Row{{8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}, {1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}, {3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}, {9850, 3, 0, 2, 0, 10, 50, 850, 1850, 4850, 9850, 100, 101, `WOAAAA`, `DAAAAA`, `VVVVxx`}, {7164, 4, 0, 0, 4, 4, 64, 164, 1164, 2164, 7164, 128, 129, `OPAAAA`, `EAAAAA`, `AAAAxx`}, {8009, 5, 1, 1, 9, 9, 9, 9, 9, 3009, 8009, 18, 19, `BWAAAA`, `FAAAAA`, `HHHHxx`}, {5057, 6, 1, 1, 7, 17, 57, 57, 1057, 57, 5057, 114, 115, `NMAAAA`, `GAAAAA`, `OOOOxx`}, {6701, 7, 1, 1, 1, 1, 1, 701, 701, 1701, 6701, 2, 3, `TXAAAA`, `HAAAAA`, `VVVVxx`}, {4321, 8, 1, 1, 1, 1, 21, 321, 321, 4321, 4321, 42, 43, `FKAAAA`, `IAAAAA`, `AAAAxx`}, {3043, 9, 1, 3, 3, 3, 43, 43, 1043, 3043, 3043, 86, 87, `BNAAAA`, `JAAAAA`, `HHHHxx`}, {1314, 10, 0, 2, 4, 14, 14, 314, 1314, 1314, 1314, 28, 29, `OYAAAA`, `KAAAAA`, `OOOOxx`}},
			},
			{
				Statement: `FETCH 12 in foo12;`,
				Results:   []sql.Row{{8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}, {1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}, {3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}, {9850, 3, 0, 2, 0, 10, 50, 850, 1850, 4850, 9850, 100, 101, `WOAAAA`, `DAAAAA`, `VVVVxx`}, {7164, 4, 0, 0, 4, 4, 64, 164, 1164, 2164, 7164, 128, 129, `OPAAAA`, `EAAAAA`, `AAAAxx`}, {8009, 5, 1, 1, 9, 9, 9, 9, 9, 3009, 8009, 18, 19, `BWAAAA`, `FAAAAA`, `HHHHxx`}, {5057, 6, 1, 1, 7, 17, 57, 57, 1057, 57, 5057, 114, 115, `NMAAAA`, `GAAAAA`, `OOOOxx`}, {6701, 7, 1, 1, 1, 1, 1, 701, 701, 1701, 6701, 2, 3, `TXAAAA`, `HAAAAA`, `VVVVxx`}, {4321, 8, 1, 1, 1, 1, 21, 321, 321, 4321, 4321, 42, 43, `FKAAAA`, `IAAAAA`, `AAAAxx`}, {3043, 9, 1, 3, 3, 3, 43, 43, 1043, 3043, 3043, 86, 87, `BNAAAA`, `JAAAAA`, `HHHHxx`}, {1314, 10, 0, 2, 4, 14, 14, 314, 1314, 1314, 1314, 28, 29, `OYAAAA`, `KAAAAA`, `OOOOxx`}, {1504, 11, 0, 0, 4, 4, 4, 504, 1504, 1504, 1504, 8, 9, `WFAAAA`, `LAAAAA`, `VVVVxx`}},
			},
			{
				Statement: `FETCH 13 in foo13;`,
				Results:   []sql.Row{{8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}, {1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}, {3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}, {9850, 3, 0, 2, 0, 10, 50, 850, 1850, 4850, 9850, 100, 101, `WOAAAA`, `DAAAAA`, `VVVVxx`}, {7164, 4, 0, 0, 4, 4, 64, 164, 1164, 2164, 7164, 128, 129, `OPAAAA`, `EAAAAA`, `AAAAxx`}, {8009, 5, 1, 1, 9, 9, 9, 9, 9, 3009, 8009, 18, 19, `BWAAAA`, `FAAAAA`, `HHHHxx`}, {5057, 6, 1, 1, 7, 17, 57, 57, 1057, 57, 5057, 114, 115, `NMAAAA`, `GAAAAA`, `OOOOxx`}, {6701, 7, 1, 1, 1, 1, 1, 701, 701, 1701, 6701, 2, 3, `TXAAAA`, `HAAAAA`, `VVVVxx`}, {4321, 8, 1, 1, 1, 1, 21, 321, 321, 4321, 4321, 42, 43, `FKAAAA`, `IAAAAA`, `AAAAxx`}, {3043, 9, 1, 3, 3, 3, 43, 43, 1043, 3043, 3043, 86, 87, `BNAAAA`, `JAAAAA`, `HHHHxx`}, {1314, 10, 0, 2, 4, 14, 14, 314, 1314, 1314, 1314, 28, 29, `OYAAAA`, `KAAAAA`, `OOOOxx`}, {1504, 11, 0, 0, 4, 4, 4, 504, 1504, 1504, 1504, 8, 9, `WFAAAA`, `LAAAAA`, `VVVVxx`}, {5222, 12, 0, 2, 2, 2, 22, 222, 1222, 222, 5222, 44, 45, `WSAAAA`, `MAAAAA`, `AAAAxx`}},
			},
			{
				Statement: `FETCH 14 in foo14;`,
				Results:   []sql.Row{{8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}, {1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}, {3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}, {9850, 3, 0, 2, 0, 10, 50, 850, 1850, 4850, 9850, 100, 101, `WOAAAA`, `DAAAAA`, `VVVVxx`}, {7164, 4, 0, 0, 4, 4, 64, 164, 1164, 2164, 7164, 128, 129, `OPAAAA`, `EAAAAA`, `AAAAxx`}, {8009, 5, 1, 1, 9, 9, 9, 9, 9, 3009, 8009, 18, 19, `BWAAAA`, `FAAAAA`, `HHHHxx`}, {5057, 6, 1, 1, 7, 17, 57, 57, 1057, 57, 5057, 114, 115, `NMAAAA`, `GAAAAA`, `OOOOxx`}, {6701, 7, 1, 1, 1, 1, 1, 701, 701, 1701, 6701, 2, 3, `TXAAAA`, `HAAAAA`, `VVVVxx`}, {4321, 8, 1, 1, 1, 1, 21, 321, 321, 4321, 4321, 42, 43, `FKAAAA`, `IAAAAA`, `AAAAxx`}, {3043, 9, 1, 3, 3, 3, 43, 43, 1043, 3043, 3043, 86, 87, `BNAAAA`, `JAAAAA`, `HHHHxx`}, {1314, 10, 0, 2, 4, 14, 14, 314, 1314, 1314, 1314, 28, 29, `OYAAAA`, `KAAAAA`, `OOOOxx`}, {1504, 11, 0, 0, 4, 4, 4, 504, 1504, 1504, 1504, 8, 9, `WFAAAA`, `LAAAAA`, `VVVVxx`}, {5222, 12, 0, 2, 2, 2, 22, 222, 1222, 222, 5222, 44, 45, `WSAAAA`, `MAAAAA`, `AAAAxx`}, {6243, 13, 1, 3, 3, 3, 43, 243, 243, 1243, 6243, 86, 87, `DGAAAA`, `NAAAAA`, `HHHHxx`}},
			},
			{
				Statement: `FETCH 15 in foo15;`,
				Results:   []sql.Row{{8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}, {1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}, {3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}, {9850, 3, 0, 2, 0, 10, 50, 850, 1850, 4850, 9850, 100, 101, `WOAAAA`, `DAAAAA`, `VVVVxx`}, {7164, 4, 0, 0, 4, 4, 64, 164, 1164, 2164, 7164, 128, 129, `OPAAAA`, `EAAAAA`, `AAAAxx`}, {8009, 5, 1, 1, 9, 9, 9, 9, 9, 3009, 8009, 18, 19, `BWAAAA`, `FAAAAA`, `HHHHxx`}, {5057, 6, 1, 1, 7, 17, 57, 57, 1057, 57, 5057, 114, 115, `NMAAAA`, `GAAAAA`, `OOOOxx`}, {6701, 7, 1, 1, 1, 1, 1, 701, 701, 1701, 6701, 2, 3, `TXAAAA`, `HAAAAA`, `VVVVxx`}, {4321, 8, 1, 1, 1, 1, 21, 321, 321, 4321, 4321, 42, 43, `FKAAAA`, `IAAAAA`, `AAAAxx`}, {3043, 9, 1, 3, 3, 3, 43, 43, 1043, 3043, 3043, 86, 87, `BNAAAA`, `JAAAAA`, `HHHHxx`}, {1314, 10, 0, 2, 4, 14, 14, 314, 1314, 1314, 1314, 28, 29, `OYAAAA`, `KAAAAA`, `OOOOxx`}, {1504, 11, 0, 0, 4, 4, 4, 504, 1504, 1504, 1504, 8, 9, `WFAAAA`, `LAAAAA`, `VVVVxx`}, {5222, 12, 0, 2, 2, 2, 22, 222, 1222, 222, 5222, 44, 45, `WSAAAA`, `MAAAAA`, `AAAAxx`}, {6243, 13, 1, 3, 3, 3, 43, 243, 243, 1243, 6243, 86, 87, `DGAAAA`, `NAAAAA`, `HHHHxx`}, {5471, 14, 1, 3, 1, 11, 71, 471, 1471, 471, 5471, 142, 143, `LCAAAA`, `OAAAAA`, `OOOOxx`}},
			},
			{
				Statement: `FETCH 16 in foo16;`,
				Results:   []sql.Row{{8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}, {1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}, {3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}, {9850, 3, 0, 2, 0, 10, 50, 850, 1850, 4850, 9850, 100, 101, `WOAAAA`, `DAAAAA`, `VVVVxx`}, {7164, 4, 0, 0, 4, 4, 64, 164, 1164, 2164, 7164, 128, 129, `OPAAAA`, `EAAAAA`, `AAAAxx`}, {8009, 5, 1, 1, 9, 9, 9, 9, 9, 3009, 8009, 18, 19, `BWAAAA`, `FAAAAA`, `HHHHxx`}, {5057, 6, 1, 1, 7, 17, 57, 57, 1057, 57, 5057, 114, 115, `NMAAAA`, `GAAAAA`, `OOOOxx`}, {6701, 7, 1, 1, 1, 1, 1, 701, 701, 1701, 6701, 2, 3, `TXAAAA`, `HAAAAA`, `VVVVxx`}, {4321, 8, 1, 1, 1, 1, 21, 321, 321, 4321, 4321, 42, 43, `FKAAAA`, `IAAAAA`, `AAAAxx`}, {3043, 9, 1, 3, 3, 3, 43, 43, 1043, 3043, 3043, 86, 87, `BNAAAA`, `JAAAAA`, `HHHHxx`}, {1314, 10, 0, 2, 4, 14, 14, 314, 1314, 1314, 1314, 28, 29, `OYAAAA`, `KAAAAA`, `OOOOxx`}, {1504, 11, 0, 0, 4, 4, 4, 504, 1504, 1504, 1504, 8, 9, `WFAAAA`, `LAAAAA`, `VVVVxx`}, {5222, 12, 0, 2, 2, 2, 22, 222, 1222, 222, 5222, 44, 45, `WSAAAA`, `MAAAAA`, `AAAAxx`}, {6243, 13, 1, 3, 3, 3, 43, 243, 243, 1243, 6243, 86, 87, `DGAAAA`, `NAAAAA`, `HHHHxx`}, {5471, 14, 1, 3, 1, 11, 71, 471, 1471, 471, 5471, 142, 143, `LCAAAA`, `OAAAAA`, `OOOOxx`}, {5006, 15, 0, 2, 6, 6, 6, 6, 1006, 6, 5006, 12, 13, `OKAAAA`, `PAAAAA`, `VVVVxx`}},
			},
			{
				Statement: `FETCH 17 in foo17;`,
				Results:   []sql.Row{{8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}, {1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}, {3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}, {9850, 3, 0, 2, 0, 10, 50, 850, 1850, 4850, 9850, 100, 101, `WOAAAA`, `DAAAAA`, `VVVVxx`}, {7164, 4, 0, 0, 4, 4, 64, 164, 1164, 2164, 7164, 128, 129, `OPAAAA`, `EAAAAA`, `AAAAxx`}, {8009, 5, 1, 1, 9, 9, 9, 9, 9, 3009, 8009, 18, 19, `BWAAAA`, `FAAAAA`, `HHHHxx`}, {5057, 6, 1, 1, 7, 17, 57, 57, 1057, 57, 5057, 114, 115, `NMAAAA`, `GAAAAA`, `OOOOxx`}, {6701, 7, 1, 1, 1, 1, 1, 701, 701, 1701, 6701, 2, 3, `TXAAAA`, `HAAAAA`, `VVVVxx`}, {4321, 8, 1, 1, 1, 1, 21, 321, 321, 4321, 4321, 42, 43, `FKAAAA`, `IAAAAA`, `AAAAxx`}, {3043, 9, 1, 3, 3, 3, 43, 43, 1043, 3043, 3043, 86, 87, `BNAAAA`, `JAAAAA`, `HHHHxx`}, {1314, 10, 0, 2, 4, 14, 14, 314, 1314, 1314, 1314, 28, 29, `OYAAAA`, `KAAAAA`, `OOOOxx`}, {1504, 11, 0, 0, 4, 4, 4, 504, 1504, 1504, 1504, 8, 9, `WFAAAA`, `LAAAAA`, `VVVVxx`}, {5222, 12, 0, 2, 2, 2, 22, 222, 1222, 222, 5222, 44, 45, `WSAAAA`, `MAAAAA`, `AAAAxx`}, {6243, 13, 1, 3, 3, 3, 43, 243, 243, 1243, 6243, 86, 87, `DGAAAA`, `NAAAAA`, `HHHHxx`}, {5471, 14, 1, 3, 1, 11, 71, 471, 1471, 471, 5471, 142, 143, `LCAAAA`, `OAAAAA`, `OOOOxx`}, {5006, 15, 0, 2, 6, 6, 6, 6, 1006, 6, 5006, 12, 13, `OKAAAA`, `PAAAAA`, `VVVVxx`}, {5387, 16, 1, 3, 7, 7, 87, 387, 1387, 387, 5387, 174, 175, `FZAAAA`, `QAAAAA`, `AAAAxx`}},
			},
			{
				Statement: `FETCH 18 in foo18;`,
				Results:   []sql.Row{{8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}, {1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}, {3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}, {9850, 3, 0, 2, 0, 10, 50, 850, 1850, 4850, 9850, 100, 101, `WOAAAA`, `DAAAAA`, `VVVVxx`}, {7164, 4, 0, 0, 4, 4, 64, 164, 1164, 2164, 7164, 128, 129, `OPAAAA`, `EAAAAA`, `AAAAxx`}, {8009, 5, 1, 1, 9, 9, 9, 9, 9, 3009, 8009, 18, 19, `BWAAAA`, `FAAAAA`, `HHHHxx`}, {5057, 6, 1, 1, 7, 17, 57, 57, 1057, 57, 5057, 114, 115, `NMAAAA`, `GAAAAA`, `OOOOxx`}, {6701, 7, 1, 1, 1, 1, 1, 701, 701, 1701, 6701, 2, 3, `TXAAAA`, `HAAAAA`, `VVVVxx`}, {4321, 8, 1, 1, 1, 1, 21, 321, 321, 4321, 4321, 42, 43, `FKAAAA`, `IAAAAA`, `AAAAxx`}, {3043, 9, 1, 3, 3, 3, 43, 43, 1043, 3043, 3043, 86, 87, `BNAAAA`, `JAAAAA`, `HHHHxx`}, {1314, 10, 0, 2, 4, 14, 14, 314, 1314, 1314, 1314, 28, 29, `OYAAAA`, `KAAAAA`, `OOOOxx`}, {1504, 11, 0, 0, 4, 4, 4, 504, 1504, 1504, 1504, 8, 9, `WFAAAA`, `LAAAAA`, `VVVVxx`}, {5222, 12, 0, 2, 2, 2, 22, 222, 1222, 222, 5222, 44, 45, `WSAAAA`, `MAAAAA`, `AAAAxx`}, {6243, 13, 1, 3, 3, 3, 43, 243, 243, 1243, 6243, 86, 87, `DGAAAA`, `NAAAAA`, `HHHHxx`}, {5471, 14, 1, 3, 1, 11, 71, 471, 1471, 471, 5471, 142, 143, `LCAAAA`, `OAAAAA`, `OOOOxx`}, {5006, 15, 0, 2, 6, 6, 6, 6, 1006, 6, 5006, 12, 13, `OKAAAA`, `PAAAAA`, `VVVVxx`}, {5387, 16, 1, 3, 7, 7, 87, 387, 1387, 387, 5387, 174, 175, `FZAAAA`, `QAAAAA`, `AAAAxx`}, {5785, 17, 1, 1, 5, 5, 85, 785, 1785, 785, 5785, 170, 171, `NOAAAA`, `RAAAAA`, `HHHHxx`}},
			},
			{
				Statement: `FETCH 19 in foo19;`,
				Results:   []sql.Row{{8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}, {1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}, {3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}, {9850, 3, 0, 2, 0, 10, 50, 850, 1850, 4850, 9850, 100, 101, `WOAAAA`, `DAAAAA`, `VVVVxx`}, {7164, 4, 0, 0, 4, 4, 64, 164, 1164, 2164, 7164, 128, 129, `OPAAAA`, `EAAAAA`, `AAAAxx`}, {8009, 5, 1, 1, 9, 9, 9, 9, 9, 3009, 8009, 18, 19, `BWAAAA`, `FAAAAA`, `HHHHxx`}, {5057, 6, 1, 1, 7, 17, 57, 57, 1057, 57, 5057, 114, 115, `NMAAAA`, `GAAAAA`, `OOOOxx`}, {6701, 7, 1, 1, 1, 1, 1, 701, 701, 1701, 6701, 2, 3, `TXAAAA`, `HAAAAA`, `VVVVxx`}, {4321, 8, 1, 1, 1, 1, 21, 321, 321, 4321, 4321, 42, 43, `FKAAAA`, `IAAAAA`, `AAAAxx`}, {3043, 9, 1, 3, 3, 3, 43, 43, 1043, 3043, 3043, 86, 87, `BNAAAA`, `JAAAAA`, `HHHHxx`}, {1314, 10, 0, 2, 4, 14, 14, 314, 1314, 1314, 1314, 28, 29, `OYAAAA`, `KAAAAA`, `OOOOxx`}, {1504, 11, 0, 0, 4, 4, 4, 504, 1504, 1504, 1504, 8, 9, `WFAAAA`, `LAAAAA`, `VVVVxx`}, {5222, 12, 0, 2, 2, 2, 22, 222, 1222, 222, 5222, 44, 45, `WSAAAA`, `MAAAAA`, `AAAAxx`}, {6243, 13, 1, 3, 3, 3, 43, 243, 243, 1243, 6243, 86, 87, `DGAAAA`, `NAAAAA`, `HHHHxx`}, {5471, 14, 1, 3, 1, 11, 71, 471, 1471, 471, 5471, 142, 143, `LCAAAA`, `OAAAAA`, `OOOOxx`}, {5006, 15, 0, 2, 6, 6, 6, 6, 1006, 6, 5006, 12, 13, `OKAAAA`, `PAAAAA`, `VVVVxx`}, {5387, 16, 1, 3, 7, 7, 87, 387, 1387, 387, 5387, 174, 175, `FZAAAA`, `QAAAAA`, `AAAAxx`}, {5785, 17, 1, 1, 5, 5, 85, 785, 1785, 785, 5785, 170, 171, `NOAAAA`, `RAAAAA`, `HHHHxx`}, {6621, 18, 1, 1, 1, 1, 21, 621, 621, 1621, 6621, 42, 43, `RUAAAA`, `SAAAAA`, `OOOOxx`}},
			},
			{
				Statement: `FETCH 20 in foo20;`,
				Results:   []sql.Row{{8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}, {1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}, {3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}, {9850, 3, 0, 2, 0, 10, 50, 850, 1850, 4850, 9850, 100, 101, `WOAAAA`, `DAAAAA`, `VVVVxx`}, {7164, 4, 0, 0, 4, 4, 64, 164, 1164, 2164, 7164, 128, 129, `OPAAAA`, `EAAAAA`, `AAAAxx`}, {8009, 5, 1, 1, 9, 9, 9, 9, 9, 3009, 8009, 18, 19, `BWAAAA`, `FAAAAA`, `HHHHxx`}, {5057, 6, 1, 1, 7, 17, 57, 57, 1057, 57, 5057, 114, 115, `NMAAAA`, `GAAAAA`, `OOOOxx`}, {6701, 7, 1, 1, 1, 1, 1, 701, 701, 1701, 6701, 2, 3, `TXAAAA`, `HAAAAA`, `VVVVxx`}, {4321, 8, 1, 1, 1, 1, 21, 321, 321, 4321, 4321, 42, 43, `FKAAAA`, `IAAAAA`, `AAAAxx`}, {3043, 9, 1, 3, 3, 3, 43, 43, 1043, 3043, 3043, 86, 87, `BNAAAA`, `JAAAAA`, `HHHHxx`}, {1314, 10, 0, 2, 4, 14, 14, 314, 1314, 1314, 1314, 28, 29, `OYAAAA`, `KAAAAA`, `OOOOxx`}, {1504, 11, 0, 0, 4, 4, 4, 504, 1504, 1504, 1504, 8, 9, `WFAAAA`, `LAAAAA`, `VVVVxx`}, {5222, 12, 0, 2, 2, 2, 22, 222, 1222, 222, 5222, 44, 45, `WSAAAA`, `MAAAAA`, `AAAAxx`}, {6243, 13, 1, 3, 3, 3, 43, 243, 243, 1243, 6243, 86, 87, `DGAAAA`, `NAAAAA`, `HHHHxx`}, {5471, 14, 1, 3, 1, 11, 71, 471, 1471, 471, 5471, 142, 143, `LCAAAA`, `OAAAAA`, `OOOOxx`}, {5006, 15, 0, 2, 6, 6, 6, 6, 1006, 6, 5006, 12, 13, `OKAAAA`, `PAAAAA`, `VVVVxx`}, {5387, 16, 1, 3, 7, 7, 87, 387, 1387, 387, 5387, 174, 175, `FZAAAA`, `QAAAAA`, `AAAAxx`}, {5785, 17, 1, 1, 5, 5, 85, 785, 1785, 785, 5785, 170, 171, `NOAAAA`, `RAAAAA`, `HHHHxx`}, {6621, 18, 1, 1, 1, 1, 21, 621, 621, 1621, 6621, 42, 43, `RUAAAA`, `SAAAAA`, `OOOOxx`}, {6969, 19, 1, 1, 9, 9, 69, 969, 969, 1969, 6969, 138, 139, `BIAAAA`, `TAAAAA`, `VVVVxx`}},
			},
			{
				Statement: `FETCH 21 in foo21;`,
				Results:   []sql.Row{{8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}, {1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}, {3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}, {9850, 3, 0, 2, 0, 10, 50, 850, 1850, 4850, 9850, 100, 101, `WOAAAA`, `DAAAAA`, `VVVVxx`}, {7164, 4, 0, 0, 4, 4, 64, 164, 1164, 2164, 7164, 128, 129, `OPAAAA`, `EAAAAA`, `AAAAxx`}, {8009, 5, 1, 1, 9, 9, 9, 9, 9, 3009, 8009, 18, 19, `BWAAAA`, `FAAAAA`, `HHHHxx`}, {5057, 6, 1, 1, 7, 17, 57, 57, 1057, 57, 5057, 114, 115, `NMAAAA`, `GAAAAA`, `OOOOxx`}, {6701, 7, 1, 1, 1, 1, 1, 701, 701, 1701, 6701, 2, 3, `TXAAAA`, `HAAAAA`, `VVVVxx`}, {4321, 8, 1, 1, 1, 1, 21, 321, 321, 4321, 4321, 42, 43, `FKAAAA`, `IAAAAA`, `AAAAxx`}, {3043, 9, 1, 3, 3, 3, 43, 43, 1043, 3043, 3043, 86, 87, `BNAAAA`, `JAAAAA`, `HHHHxx`}, {1314, 10, 0, 2, 4, 14, 14, 314, 1314, 1314, 1314, 28, 29, `OYAAAA`, `KAAAAA`, `OOOOxx`}, {1504, 11, 0, 0, 4, 4, 4, 504, 1504, 1504, 1504, 8, 9, `WFAAAA`, `LAAAAA`, `VVVVxx`}, {5222, 12, 0, 2, 2, 2, 22, 222, 1222, 222, 5222, 44, 45, `WSAAAA`, `MAAAAA`, `AAAAxx`}, {6243, 13, 1, 3, 3, 3, 43, 243, 243, 1243, 6243, 86, 87, `DGAAAA`, `NAAAAA`, `HHHHxx`}, {5471, 14, 1, 3, 1, 11, 71, 471, 1471, 471, 5471, 142, 143, `LCAAAA`, `OAAAAA`, `OOOOxx`}, {5006, 15, 0, 2, 6, 6, 6, 6, 1006, 6, 5006, 12, 13, `OKAAAA`, `PAAAAA`, `VVVVxx`}, {5387, 16, 1, 3, 7, 7, 87, 387, 1387, 387, 5387, 174, 175, `FZAAAA`, `QAAAAA`, `AAAAxx`}, {5785, 17, 1, 1, 5, 5, 85, 785, 1785, 785, 5785, 170, 171, `NOAAAA`, `RAAAAA`, `HHHHxx`}, {6621, 18, 1, 1, 1, 1, 21, 621, 621, 1621, 6621, 42, 43, `RUAAAA`, `SAAAAA`, `OOOOxx`}, {6969, 19, 1, 1, 9, 9, 69, 969, 969, 1969, 6969, 138, 139, `BIAAAA`, `TAAAAA`, `VVVVxx`}, {9460, 20, 0, 0, 0, 0, 60, 460, 1460, 4460, 9460, 120, 121, `WZAAAA`, `UAAAAA`, `AAAAxx`}},
			},
			{
				Statement: `FETCH 22 in foo22;`,
				Results:   []sql.Row{{8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}, {1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}, {3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}, {9850, 3, 0, 2, 0, 10, 50, 850, 1850, 4850, 9850, 100, 101, `WOAAAA`, `DAAAAA`, `VVVVxx`}, {7164, 4, 0, 0, 4, 4, 64, 164, 1164, 2164, 7164, 128, 129, `OPAAAA`, `EAAAAA`, `AAAAxx`}, {8009, 5, 1, 1, 9, 9, 9, 9, 9, 3009, 8009, 18, 19, `BWAAAA`, `FAAAAA`, `HHHHxx`}, {5057, 6, 1, 1, 7, 17, 57, 57, 1057, 57, 5057, 114, 115, `NMAAAA`, `GAAAAA`, `OOOOxx`}, {6701, 7, 1, 1, 1, 1, 1, 701, 701, 1701, 6701, 2, 3, `TXAAAA`, `HAAAAA`, `VVVVxx`}, {4321, 8, 1, 1, 1, 1, 21, 321, 321, 4321, 4321, 42, 43, `FKAAAA`, `IAAAAA`, `AAAAxx`}, {3043, 9, 1, 3, 3, 3, 43, 43, 1043, 3043, 3043, 86, 87, `BNAAAA`, `JAAAAA`, `HHHHxx`}, {1314, 10, 0, 2, 4, 14, 14, 314, 1314, 1314, 1314, 28, 29, `OYAAAA`, `KAAAAA`, `OOOOxx`}, {1504, 11, 0, 0, 4, 4, 4, 504, 1504, 1504, 1504, 8, 9, `WFAAAA`, `LAAAAA`, `VVVVxx`}, {5222, 12, 0, 2, 2, 2, 22, 222, 1222, 222, 5222, 44, 45, `WSAAAA`, `MAAAAA`, `AAAAxx`}, {6243, 13, 1, 3, 3, 3, 43, 243, 243, 1243, 6243, 86, 87, `DGAAAA`, `NAAAAA`, `HHHHxx`}, {5471, 14, 1, 3, 1, 11, 71, 471, 1471, 471, 5471, 142, 143, `LCAAAA`, `OAAAAA`, `OOOOxx`}, {5006, 15, 0, 2, 6, 6, 6, 6, 1006, 6, 5006, 12, 13, `OKAAAA`, `PAAAAA`, `VVVVxx`}, {5387, 16, 1, 3, 7, 7, 87, 387, 1387, 387, 5387, 174, 175, `FZAAAA`, `QAAAAA`, `AAAAxx`}, {5785, 17, 1, 1, 5, 5, 85, 785, 1785, 785, 5785, 170, 171, `NOAAAA`, `RAAAAA`, `HHHHxx`}, {6621, 18, 1, 1, 1, 1, 21, 621, 621, 1621, 6621, 42, 43, `RUAAAA`, `SAAAAA`, `OOOOxx`}, {6969, 19, 1, 1, 9, 9, 69, 969, 969, 1969, 6969, 138, 139, `BIAAAA`, `TAAAAA`, `VVVVxx`}, {9460, 20, 0, 0, 0, 0, 60, 460, 1460, 4460, 9460, 120, 121, `WZAAAA`, `UAAAAA`, `AAAAxx`}, {59, 21, 1, 3, 9, 19, 59, 59, 59, 59, 59, 118, 119, `HCAAAA`, `VAAAAA`, `HHHHxx`}},
			},
			{
				Statement: `FETCH 23 in foo23;`,
				Results:   []sql.Row{{8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}, {1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}, {3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}, {9850, 3, 0, 2, 0, 10, 50, 850, 1850, 4850, 9850, 100, 101, `WOAAAA`, `DAAAAA`, `VVVVxx`}, {7164, 4, 0, 0, 4, 4, 64, 164, 1164, 2164, 7164, 128, 129, `OPAAAA`, `EAAAAA`, `AAAAxx`}, {8009, 5, 1, 1, 9, 9, 9, 9, 9, 3009, 8009, 18, 19, `BWAAAA`, `FAAAAA`, `HHHHxx`}, {5057, 6, 1, 1, 7, 17, 57, 57, 1057, 57, 5057, 114, 115, `NMAAAA`, `GAAAAA`, `OOOOxx`}, {6701, 7, 1, 1, 1, 1, 1, 701, 701, 1701, 6701, 2, 3, `TXAAAA`, `HAAAAA`, `VVVVxx`}, {4321, 8, 1, 1, 1, 1, 21, 321, 321, 4321, 4321, 42, 43, `FKAAAA`, `IAAAAA`, `AAAAxx`}, {3043, 9, 1, 3, 3, 3, 43, 43, 1043, 3043, 3043, 86, 87, `BNAAAA`, `JAAAAA`, `HHHHxx`}, {1314, 10, 0, 2, 4, 14, 14, 314, 1314, 1314, 1314, 28, 29, `OYAAAA`, `KAAAAA`, `OOOOxx`}, {1504, 11, 0, 0, 4, 4, 4, 504, 1504, 1504, 1504, 8, 9, `WFAAAA`, `LAAAAA`, `VVVVxx`}, {5222, 12, 0, 2, 2, 2, 22, 222, 1222, 222, 5222, 44, 45, `WSAAAA`, `MAAAAA`, `AAAAxx`}, {6243, 13, 1, 3, 3, 3, 43, 243, 243, 1243, 6243, 86, 87, `DGAAAA`, `NAAAAA`, `HHHHxx`}, {5471, 14, 1, 3, 1, 11, 71, 471, 1471, 471, 5471, 142, 143, `LCAAAA`, `OAAAAA`, `OOOOxx`}, {5006, 15, 0, 2, 6, 6, 6, 6, 1006, 6, 5006, 12, 13, `OKAAAA`, `PAAAAA`, `VVVVxx`}, {5387, 16, 1, 3, 7, 7, 87, 387, 1387, 387, 5387, 174, 175, `FZAAAA`, `QAAAAA`, `AAAAxx`}, {5785, 17, 1, 1, 5, 5, 85, 785, 1785, 785, 5785, 170, 171, `NOAAAA`, `RAAAAA`, `HHHHxx`}, {6621, 18, 1, 1, 1, 1, 21, 621, 621, 1621, 6621, 42, 43, `RUAAAA`, `SAAAAA`, `OOOOxx`}, {6969, 19, 1, 1, 9, 9, 69, 969, 969, 1969, 6969, 138, 139, `BIAAAA`, `TAAAAA`, `VVVVxx`}, {9460, 20, 0, 0, 0, 0, 60, 460, 1460, 4460, 9460, 120, 121, `WZAAAA`, `UAAAAA`, `AAAAxx`}, {59, 21, 1, 3, 9, 19, 59, 59, 59, 59, 59, 118, 119, `HCAAAA`, `VAAAAA`, `HHHHxx`}, {8020, 22, 0, 0, 0, 0, 20, 20, 20, 3020, 8020, 40, 41, `MWAAAA`, `WAAAAA`, `OOOOxx`}},
			},
			{
				Statement: `FETCH backward 1 in foo23;`,
				Results:   []sql.Row{{59, 21, 1, 3, 9, 19, 59, 59, 59, 59, 59, 118, 119, `HCAAAA`, `VAAAAA`, `HHHHxx`}},
			},
			{
				Statement: `FETCH backward 2 in foo22;`,
				Results:   []sql.Row{{9460, 20, 0, 0, 0, 0, 60, 460, 1460, 4460, 9460, 120, 121, `WZAAAA`, `UAAAAA`, `AAAAxx`}, {6969, 19, 1, 1, 9, 9, 69, 969, 969, 1969, 6969, 138, 139, `BIAAAA`, `TAAAAA`, `VVVVxx`}},
			},
			{
				Statement: `FETCH backward 3 in foo21;`,
				Results:   []sql.Row{{6969, 19, 1, 1, 9, 9, 69, 969, 969, 1969, 6969, 138, 139, `BIAAAA`, `TAAAAA`, `VVVVxx`}, {6621, 18, 1, 1, 1, 1, 21, 621, 621, 1621, 6621, 42, 43, `RUAAAA`, `SAAAAA`, `OOOOxx`}, {5785, 17, 1, 1, 5, 5, 85, 785, 1785, 785, 5785, 170, 171, `NOAAAA`, `RAAAAA`, `HHHHxx`}},
			},
			{
				Statement: `FETCH backward 4 in foo20;`,
				Results:   []sql.Row{{6621, 18, 1, 1, 1, 1, 21, 621, 621, 1621, 6621, 42, 43, `RUAAAA`, `SAAAAA`, `OOOOxx`}, {5785, 17, 1, 1, 5, 5, 85, 785, 1785, 785, 5785, 170, 171, `NOAAAA`, `RAAAAA`, `HHHHxx`}, {5387, 16, 1, 3, 7, 7, 87, 387, 1387, 387, 5387, 174, 175, `FZAAAA`, `QAAAAA`, `AAAAxx`}, {5006, 15, 0, 2, 6, 6, 6, 6, 1006, 6, 5006, 12, 13, `OKAAAA`, `PAAAAA`, `VVVVxx`}},
			},
			{
				Statement: `FETCH backward 5 in foo19;`,
				Results:   []sql.Row{{5785, 17, 1, 1, 5, 5, 85, 785, 1785, 785, 5785, 170, 171, `NOAAAA`, `RAAAAA`, `HHHHxx`}, {5387, 16, 1, 3, 7, 7, 87, 387, 1387, 387, 5387, 174, 175, `FZAAAA`, `QAAAAA`, `AAAAxx`}, {5006, 15, 0, 2, 6, 6, 6, 6, 1006, 6, 5006, 12, 13, `OKAAAA`, `PAAAAA`, `VVVVxx`}, {5471, 14, 1, 3, 1, 11, 71, 471, 1471, 471, 5471, 142, 143, `LCAAAA`, `OAAAAA`, `OOOOxx`}, {6243, 13, 1, 3, 3, 3, 43, 243, 243, 1243, 6243, 86, 87, `DGAAAA`, `NAAAAA`, `HHHHxx`}},
			},
			{
				Statement: `FETCH backward 6 in foo18;`,
				Results:   []sql.Row{{5387, 16, 1, 3, 7, 7, 87, 387, 1387, 387, 5387, 174, 175, `FZAAAA`, `QAAAAA`, `AAAAxx`}, {5006, 15, 0, 2, 6, 6, 6, 6, 1006, 6, 5006, 12, 13, `OKAAAA`, `PAAAAA`, `VVVVxx`}, {5471, 14, 1, 3, 1, 11, 71, 471, 1471, 471, 5471, 142, 143, `LCAAAA`, `OAAAAA`, `OOOOxx`}, {6243, 13, 1, 3, 3, 3, 43, 243, 243, 1243, 6243, 86, 87, `DGAAAA`, `NAAAAA`, `HHHHxx`}, {5222, 12, 0, 2, 2, 2, 22, 222, 1222, 222, 5222, 44, 45, `WSAAAA`, `MAAAAA`, `AAAAxx`}, {1504, 11, 0, 0, 4, 4, 4, 504, 1504, 1504, 1504, 8, 9, `WFAAAA`, `LAAAAA`, `VVVVxx`}},
			},
			{
				Statement: `FETCH backward 7 in foo17;`,
				Results:   []sql.Row{{5006, 15, 0, 2, 6, 6, 6, 6, 1006, 6, 5006, 12, 13, `OKAAAA`, `PAAAAA`, `VVVVxx`}, {5471, 14, 1, 3, 1, 11, 71, 471, 1471, 471, 5471, 142, 143, `LCAAAA`, `OAAAAA`, `OOOOxx`}, {6243, 13, 1, 3, 3, 3, 43, 243, 243, 1243, 6243, 86, 87, `DGAAAA`, `NAAAAA`, `HHHHxx`}, {5222, 12, 0, 2, 2, 2, 22, 222, 1222, 222, 5222, 44, 45, `WSAAAA`, `MAAAAA`, `AAAAxx`}, {1504, 11, 0, 0, 4, 4, 4, 504, 1504, 1504, 1504, 8, 9, `WFAAAA`, `LAAAAA`, `VVVVxx`}, {1314, 10, 0, 2, 4, 14, 14, 314, 1314, 1314, 1314, 28, 29, `OYAAAA`, `KAAAAA`, `OOOOxx`}, {3043, 9, 1, 3, 3, 3, 43, 43, 1043, 3043, 3043, 86, 87, `BNAAAA`, `JAAAAA`, `HHHHxx`}},
			},
			{
				Statement: `FETCH backward 8 in foo16;`,
				Results:   []sql.Row{{5471, 14, 1, 3, 1, 11, 71, 471, 1471, 471, 5471, 142, 143, `LCAAAA`, `OAAAAA`, `OOOOxx`}, {6243, 13, 1, 3, 3, 3, 43, 243, 243, 1243, 6243, 86, 87, `DGAAAA`, `NAAAAA`, `HHHHxx`}, {5222, 12, 0, 2, 2, 2, 22, 222, 1222, 222, 5222, 44, 45, `WSAAAA`, `MAAAAA`, `AAAAxx`}, {1504, 11, 0, 0, 4, 4, 4, 504, 1504, 1504, 1504, 8, 9, `WFAAAA`, `LAAAAA`, `VVVVxx`}, {1314, 10, 0, 2, 4, 14, 14, 314, 1314, 1314, 1314, 28, 29, `OYAAAA`, `KAAAAA`, `OOOOxx`}, {3043, 9, 1, 3, 3, 3, 43, 43, 1043, 3043, 3043, 86, 87, `BNAAAA`, `JAAAAA`, `HHHHxx`}, {4321, 8, 1, 1, 1, 1, 21, 321, 321, 4321, 4321, 42, 43, `FKAAAA`, `IAAAAA`, `AAAAxx`}, {6701, 7, 1, 1, 1, 1, 1, 701, 701, 1701, 6701, 2, 3, `TXAAAA`, `HAAAAA`, `VVVVxx`}},
			},
			{
				Statement: `FETCH backward 9 in foo15;`,
				Results:   []sql.Row{{6243, 13, 1, 3, 3, 3, 43, 243, 243, 1243, 6243, 86, 87, `DGAAAA`, `NAAAAA`, `HHHHxx`}, {5222, 12, 0, 2, 2, 2, 22, 222, 1222, 222, 5222, 44, 45, `WSAAAA`, `MAAAAA`, `AAAAxx`}, {1504, 11, 0, 0, 4, 4, 4, 504, 1504, 1504, 1504, 8, 9, `WFAAAA`, `LAAAAA`, `VVVVxx`}, {1314, 10, 0, 2, 4, 14, 14, 314, 1314, 1314, 1314, 28, 29, `OYAAAA`, `KAAAAA`, `OOOOxx`}, {3043, 9, 1, 3, 3, 3, 43, 43, 1043, 3043, 3043, 86, 87, `BNAAAA`, `JAAAAA`, `HHHHxx`}, {4321, 8, 1, 1, 1, 1, 21, 321, 321, 4321, 4321, 42, 43, `FKAAAA`, `IAAAAA`, `AAAAxx`}, {6701, 7, 1, 1, 1, 1, 1, 701, 701, 1701, 6701, 2, 3, `TXAAAA`, `HAAAAA`, `VVVVxx`}, {5057, 6, 1, 1, 7, 17, 57, 57, 1057, 57, 5057, 114, 115, `NMAAAA`, `GAAAAA`, `OOOOxx`}, {8009, 5, 1, 1, 9, 9, 9, 9, 9, 3009, 8009, 18, 19, `BWAAAA`, `FAAAAA`, `HHHHxx`}},
			},
			{
				Statement: `FETCH backward 10 in foo14;`,
				Results:   []sql.Row{{5222, 12, 0, 2, 2, 2, 22, 222, 1222, 222, 5222, 44, 45, `WSAAAA`, `MAAAAA`, `AAAAxx`}, {1504, 11, 0, 0, 4, 4, 4, 504, 1504, 1504, 1504, 8, 9, `WFAAAA`, `LAAAAA`, `VVVVxx`}, {1314, 10, 0, 2, 4, 14, 14, 314, 1314, 1314, 1314, 28, 29, `OYAAAA`, `KAAAAA`, `OOOOxx`}, {3043, 9, 1, 3, 3, 3, 43, 43, 1043, 3043, 3043, 86, 87, `BNAAAA`, `JAAAAA`, `HHHHxx`}, {4321, 8, 1, 1, 1, 1, 21, 321, 321, 4321, 4321, 42, 43, `FKAAAA`, `IAAAAA`, `AAAAxx`}, {6701, 7, 1, 1, 1, 1, 1, 701, 701, 1701, 6701, 2, 3, `TXAAAA`, `HAAAAA`, `VVVVxx`}, {5057, 6, 1, 1, 7, 17, 57, 57, 1057, 57, 5057, 114, 115, `NMAAAA`, `GAAAAA`, `OOOOxx`}, {8009, 5, 1, 1, 9, 9, 9, 9, 9, 3009, 8009, 18, 19, `BWAAAA`, `FAAAAA`, `HHHHxx`}, {7164, 4, 0, 0, 4, 4, 64, 164, 1164, 2164, 7164, 128, 129, `OPAAAA`, `EAAAAA`, `AAAAxx`}, {9850, 3, 0, 2, 0, 10, 50, 850, 1850, 4850, 9850, 100, 101, `WOAAAA`, `DAAAAA`, `VVVVxx`}},
			},
			{
				Statement: `FETCH backward 11 in foo13;`,
				Results:   []sql.Row{{1504, 11, 0, 0, 4, 4, 4, 504, 1504, 1504, 1504, 8, 9, `WFAAAA`, `LAAAAA`, `VVVVxx`}, {1314, 10, 0, 2, 4, 14, 14, 314, 1314, 1314, 1314, 28, 29, `OYAAAA`, `KAAAAA`, `OOOOxx`}, {3043, 9, 1, 3, 3, 3, 43, 43, 1043, 3043, 3043, 86, 87, `BNAAAA`, `JAAAAA`, `HHHHxx`}, {4321, 8, 1, 1, 1, 1, 21, 321, 321, 4321, 4321, 42, 43, `FKAAAA`, `IAAAAA`, `AAAAxx`}, {6701, 7, 1, 1, 1, 1, 1, 701, 701, 1701, 6701, 2, 3, `TXAAAA`, `HAAAAA`, `VVVVxx`}, {5057, 6, 1, 1, 7, 17, 57, 57, 1057, 57, 5057, 114, 115, `NMAAAA`, `GAAAAA`, `OOOOxx`}, {8009, 5, 1, 1, 9, 9, 9, 9, 9, 3009, 8009, 18, 19, `BWAAAA`, `FAAAAA`, `HHHHxx`}, {7164, 4, 0, 0, 4, 4, 64, 164, 1164, 2164, 7164, 128, 129, `OPAAAA`, `EAAAAA`, `AAAAxx`}, {9850, 3, 0, 2, 0, 10, 50, 850, 1850, 4850, 9850, 100, 101, `WOAAAA`, `DAAAAA`, `VVVVxx`}, {3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}, {1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}},
			},
			{
				Statement: `FETCH backward 12 in foo12;`,
				Results:   []sql.Row{{1314, 10, 0, 2, 4, 14, 14, 314, 1314, 1314, 1314, 28, 29, `OYAAAA`, `KAAAAA`, `OOOOxx`}, {3043, 9, 1, 3, 3, 3, 43, 43, 1043, 3043, 3043, 86, 87, `BNAAAA`, `JAAAAA`, `HHHHxx`}, {4321, 8, 1, 1, 1, 1, 21, 321, 321, 4321, 4321, 42, 43, `FKAAAA`, `IAAAAA`, `AAAAxx`}, {6701, 7, 1, 1, 1, 1, 1, 701, 701, 1701, 6701, 2, 3, `TXAAAA`, `HAAAAA`, `VVVVxx`}, {5057, 6, 1, 1, 7, 17, 57, 57, 1057, 57, 5057, 114, 115, `NMAAAA`, `GAAAAA`, `OOOOxx`}, {8009, 5, 1, 1, 9, 9, 9, 9, 9, 3009, 8009, 18, 19, `BWAAAA`, `FAAAAA`, `HHHHxx`}, {7164, 4, 0, 0, 4, 4, 64, 164, 1164, 2164, 7164, 128, 129, `OPAAAA`, `EAAAAA`, `AAAAxx`}, {9850, 3, 0, 2, 0, 10, 50, 850, 1850, 4850, 9850, 100, 101, `WOAAAA`, `DAAAAA`, `VVVVxx`}, {3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}, {1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}, {8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}},
			},
			{
				Statement: `FETCH backward 13 in foo11;`,
				Results:   []sql.Row{{3043, 9, 1, 3, 3, 3, 43, 43, 1043, 3043, 3043, 86, 87, `BNAAAA`, `JAAAAA`, `HHHHxx`}, {4321, 8, 1, 1, 1, 1, 21, 321, 321, 4321, 4321, 42, 43, `FKAAAA`, `IAAAAA`, `AAAAxx`}, {6701, 7, 1, 1, 1, 1, 1, 701, 701, 1701, 6701, 2, 3, `TXAAAA`, `HAAAAA`, `VVVVxx`}, {5057, 6, 1, 1, 7, 17, 57, 57, 1057, 57, 5057, 114, 115, `NMAAAA`, `GAAAAA`, `OOOOxx`}, {8009, 5, 1, 1, 9, 9, 9, 9, 9, 3009, 8009, 18, 19, `BWAAAA`, `FAAAAA`, `HHHHxx`}, {7164, 4, 0, 0, 4, 4, 64, 164, 1164, 2164, 7164, 128, 129, `OPAAAA`, `EAAAAA`, `AAAAxx`}, {9850, 3, 0, 2, 0, 10, 50, 850, 1850, 4850, 9850, 100, 101, `WOAAAA`, `DAAAAA`, `VVVVxx`}, {3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}, {1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}, {8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}},
			},
			{
				Statement: `FETCH backward 14 in foo10;`,
				Results:   []sql.Row{{4321, 8, 1, 1, 1, 1, 21, 321, 321, 4321, 4321, 42, 43, `FKAAAA`, `IAAAAA`, `AAAAxx`}, {6701, 7, 1, 1, 1, 1, 1, 701, 701, 1701, 6701, 2, 3, `TXAAAA`, `HAAAAA`, `VVVVxx`}, {5057, 6, 1, 1, 7, 17, 57, 57, 1057, 57, 5057, 114, 115, `NMAAAA`, `GAAAAA`, `OOOOxx`}, {8009, 5, 1, 1, 9, 9, 9, 9, 9, 3009, 8009, 18, 19, `BWAAAA`, `FAAAAA`, `HHHHxx`}, {7164, 4, 0, 0, 4, 4, 64, 164, 1164, 2164, 7164, 128, 129, `OPAAAA`, `EAAAAA`, `AAAAxx`}, {9850, 3, 0, 2, 0, 10, 50, 850, 1850, 4850, 9850, 100, 101, `WOAAAA`, `DAAAAA`, `VVVVxx`}, {3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}, {1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}, {8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}},
			},
			{
				Statement: `FETCH backward 15 in foo9;`,
				Results:   []sql.Row{{6701, 7, 1, 1, 1, 1, 1, 701, 701, 1701, 6701, 2, 3, `TXAAAA`, `HAAAAA`, `VVVVxx`}, {5057, 6, 1, 1, 7, 17, 57, 57, 1057, 57, 5057, 114, 115, `NMAAAA`, `GAAAAA`, `OOOOxx`}, {8009, 5, 1, 1, 9, 9, 9, 9, 9, 3009, 8009, 18, 19, `BWAAAA`, `FAAAAA`, `HHHHxx`}, {7164, 4, 0, 0, 4, 4, 64, 164, 1164, 2164, 7164, 128, 129, `OPAAAA`, `EAAAAA`, `AAAAxx`}, {9850, 3, 0, 2, 0, 10, 50, 850, 1850, 4850, 9850, 100, 101, `WOAAAA`, `DAAAAA`, `VVVVxx`}, {3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}, {1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}, {8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}},
			},
			{
				Statement: `FETCH backward 16 in foo8;`,
				Results:   []sql.Row{{5057, 6, 1, 1, 7, 17, 57, 57, 1057, 57, 5057, 114, 115, `NMAAAA`, `GAAAAA`, `OOOOxx`}, {8009, 5, 1, 1, 9, 9, 9, 9, 9, 3009, 8009, 18, 19, `BWAAAA`, `FAAAAA`, `HHHHxx`}, {7164, 4, 0, 0, 4, 4, 64, 164, 1164, 2164, 7164, 128, 129, `OPAAAA`, `EAAAAA`, `AAAAxx`}, {9850, 3, 0, 2, 0, 10, 50, 850, 1850, 4850, 9850, 100, 101, `WOAAAA`, `DAAAAA`, `VVVVxx`}, {3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}, {1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}, {8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}},
			},
			{
				Statement: `FETCH backward 17 in foo7;`,
				Results:   []sql.Row{{8009, 5, 1, 1, 9, 9, 9, 9, 9, 3009, 8009, 18, 19, `BWAAAA`, `FAAAAA`, `HHHHxx`}, {7164, 4, 0, 0, 4, 4, 64, 164, 1164, 2164, 7164, 128, 129, `OPAAAA`, `EAAAAA`, `AAAAxx`}, {9850, 3, 0, 2, 0, 10, 50, 850, 1850, 4850, 9850, 100, 101, `WOAAAA`, `DAAAAA`, `VVVVxx`}, {3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}, {1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}, {8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}},
			},
			{
				Statement: `FETCH backward 18 in foo6;`,
				Results:   []sql.Row{{7164, 4, 0, 0, 4, 4, 64, 164, 1164, 2164, 7164, 128, 129, `OPAAAA`, `EAAAAA`, `AAAAxx`}, {9850, 3, 0, 2, 0, 10, 50, 850, 1850, 4850, 9850, 100, 101, `WOAAAA`, `DAAAAA`, `VVVVxx`}, {3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}, {1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}, {8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}},
			},
			{
				Statement: `FETCH backward 19 in foo5;`,
				Results:   []sql.Row{{9850, 3, 0, 2, 0, 10, 50, 850, 1850, 4850, 9850, 100, 101, `WOAAAA`, `DAAAAA`, `VVVVxx`}, {3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}, {1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}, {8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}},
			},
			{
				Statement: `FETCH backward 20 in foo4;`,
				Results:   []sql.Row{{3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}, {1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}, {8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}},
			},
			{
				Statement: `FETCH backward 21 in foo3;`,
				Results:   []sql.Row{{1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}, {8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}},
			},
			{
				Statement: `FETCH backward 22 in foo2;`,
				Results:   []sql.Row{{8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}},
			},
			{
				Statement: `FETCH backward 23 in foo1;`,
				Results:   []sql.Row{},
			},
			{
				Statement: `CLOSE foo1;`,
			},
			{
				Statement: `CLOSE foo2;`,
			},
			{
				Statement: `CLOSE foo3;`,
			},
			{
				Statement: `CLOSE foo4;`,
			},
			{
				Statement: `CLOSE foo5;`,
			},
			{
				Statement: `CLOSE foo6;`,
			},
			{
				Statement: `CLOSE foo7;`,
			},
			{
				Statement: `CLOSE foo8;`,
			},
			{
				Statement: `CLOSE foo9;`,
			},
			{
				Statement: `CLOSE foo10;`,
			},
			{
				Statement: `CLOSE foo11;`,
			},
			{
				Statement: `CLOSE foo12;`,
			},
			{
				Statement: `SELECT name, statement, is_holdable, is_binary, is_scrollable FROM pg_cursors ORDER BY 1;`,
				Results:   []sql.Row{{`foo13`, `DECLARE foo13 SCROLL CURSOR FOR SELECT * FROM tenk1 ORDER BY unique2;`, false, false, true}, {`foo14`, `DECLARE foo14 SCROLL CURSOR FOR SELECT * FROM tenk2;`, false, false, true}, {`foo15`, `DECLARE foo15 SCROLL CURSOR FOR SELECT * FROM tenk1 ORDER BY unique2;`, false, false, true}, {`foo16`, `DECLARE foo16 SCROLL CURSOR FOR SELECT * FROM tenk2;`, false, false, true}, {`foo17`, `DECLARE foo17 SCROLL CURSOR FOR SELECT * FROM tenk1 ORDER BY unique2;`, false, false, true}, {`foo18`, `DECLARE foo18 SCROLL CURSOR FOR SELECT * FROM tenk2;`, false, false, true}, {`foo19`, `DECLARE foo19 SCROLL CURSOR FOR SELECT * FROM tenk1 ORDER BY unique2;`, false, false, true}, {`foo20`, `DECLARE foo20 SCROLL CURSOR FOR SELECT * FROM tenk2;`, false, false, true}, {`foo21`, `DECLARE foo21 SCROLL CURSOR FOR SELECT * FROM tenk1 ORDER BY unique2;`, false, false, true}, {`foo22`, `DECLARE foo22 SCROLL CURSOR FOR SELECT * FROM tenk2;`, false, false, true}, {`foo23`, `DECLARE foo23 SCROLL CURSOR FOR SELECT * FROM tenk1 ORDER BY unique2;`, false, false, true}},
			},
			{
				Statement: `END;`,
			},
			{
				Statement: `SELECT name, statement, is_holdable, is_binary, is_scrollable FROM pg_cursors;`,
				Results:   []sql.Row{},
			},
			{
				Statement: `BEGIN;`,
			},
			{
				Statement: `DECLARE foo24 NO SCROLL CURSOR FOR SELECT * FROM tenk1 ORDER BY unique2;`,
			},
			{
				Statement: `FETCH 1 FROM foo24;`,
				Results:   []sql.Row{{8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}},
			},
			{
				Statement:   `FETCH BACKWARD 1 FROM foo24; -- should fail`,
				ErrorString: `cursor can only scan forward`,
			},
			{
				Statement: `END;`,
			},
			{
				Statement: `BEGIN;`,
			},
			{
				Statement: `DECLARE foo24 NO SCROLL CURSOR FOR SELECT * FROM tenk1 ORDER BY unique2;`,
			},
			{
				Statement: `FETCH 1 FROM foo24;`,
				Results:   []sql.Row{{8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}},
			},
			{
				Statement: `FETCH ABSOLUTE 2 FROM foo24; -- allowed`,
				Results:   []sql.Row{{1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}},
			},
			{
				Statement:   `FETCH ABSOLUTE 1 FROM foo24; -- should fail`,
				ErrorString: `cursor can only scan forward`,
			},
			{
				Statement: `END;`,
			},
			{
				Statement: `SELECT name, statement, is_holdable, is_binary, is_scrollable FROM pg_cursors;`,
				Results:   []sql.Row{},
			},
			{
				Statement: `BEGIN;`,
			},
			{
				Statement: `DECLARE foo25 SCROLL CURSOR WITH HOLD FOR SELECT * FROM tenk2;`,
			},
			{
				Statement: `FETCH FROM foo25;`,
				Results:   []sql.Row{{8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}},
			},
			{
				Statement: `FETCH FROM foo25;`,
				Results:   []sql.Row{{1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}},
			},
			{
				Statement: `COMMIT;`,
			},
			{
				Statement: `FETCH FROM foo25;`,
				Results:   []sql.Row{{3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}},
			},
			{
				Statement: `FETCH BACKWARD FROM foo25;`,
				Results:   []sql.Row{{1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}},
			},
			{
				Statement: `FETCH ABSOLUTE -1 FROM foo25;`,
				Results:   []sql.Row{{2968, 9999, 0, 0, 8, 8, 68, 968, 968, 2968, 2968, 136, 137, `EKAAAA`, `PUOAAA`, `VVVVxx`}},
			},
			{
				Statement: `SELECT name, statement, is_holdable, is_binary, is_scrollable FROM pg_cursors;`,
				Results:   []sql.Row{{`foo25`, `DECLARE foo25 SCROLL CURSOR WITH HOLD FOR SELECT * FROM tenk2;`, true, false, true}},
			},
			{
				Statement: `CLOSE foo25;`,
			},
			{
				Statement: `BEGIN;`,
			},
			{
				Statement: `DECLARE foo25ns NO SCROLL CURSOR WITH HOLD FOR SELECT * FROM tenk2;`,
			},
			{
				Statement: `FETCH FROM foo25ns;`,
				Results:   []sql.Row{{8800, 0, 0, 0, 0, 0, 0, 800, 800, 3800, 8800, 0, 1, `MAAAAA`, `AAAAAA`, `AAAAxx`}},
			},
			{
				Statement: `FETCH FROM foo25ns;`,
				Results:   []sql.Row{{1891, 1, 1, 3, 1, 11, 91, 891, 1891, 1891, 1891, 182, 183, `TUAAAA`, `BAAAAA`, `HHHHxx`}},
			},
			{
				Statement: `COMMIT;`,
			},
			{
				Statement: `FETCH FROM foo25ns;`,
				Results:   []sql.Row{{3420, 2, 0, 0, 0, 0, 20, 420, 1420, 3420, 3420, 40, 41, `OBAAAA`, `CAAAAA`, `OOOOxx`}},
			},
			{
				Statement: `FETCH ABSOLUTE 4 FROM foo25ns;`,
				Results:   []sql.Row{{9850, 3, 0, 2, 0, 10, 50, 850, 1850, 4850, 9850, 100, 101, `WOAAAA`, `DAAAAA`, `VVVVxx`}},
			},
			{
				Statement:   `FETCH ABSOLUTE 4 FROM foo25ns; -- fail`,
				ErrorString: `cursor can only scan forward`,
			},
			{
				Statement: `SELECT name, statement, is_holdable, is_binary, is_scrollable FROM pg_cursors;`,
				Results:   []sql.Row{{`foo25ns`, `DECLARE foo25ns NO SCROLL CURSOR WITH HOLD FOR SELECT * FROM tenk2;`, true, false, false}},
			},
			{
				Statement: `CLOSE foo25ns;`,
			},
			{
				Statement: `BEGIN;`,
			},
			{
				Statement: `DECLARE foo26 CURSOR WITH HOLD FOR SELECT * FROM tenk1 ORDER BY unique2;`,
			},
			{
				Statement: `ROLLBACK;`,
			},
			{
				Statement:   `FETCH FROM foo26;`,
				ErrorString: `cursor "foo26" does not exist`,
			},
			{
				Statement: `BEGIN;`,
			},
			{
				Statement: `CREATE FUNCTION declares_cursor(text)
   RETURNS void
   AS 'DECLARE c CURSOR FOR SELECT stringu1 FROM tenk1 WHERE stringu1 LIKE $1;'
   LANGUAGE SQL;`,
			},
			{
				Statement: `SELECT declares_cursor('AB%');`,
				Results:   []sql.Row{{``}},
			},
			{
				Statement: `FETCH ALL FROM c;`,
				Results:   []sql.Row{{`ABAAAA`}, {`ABAAAA`}, {`ABAAAA`}, {`ABAAAA`}, {`ABAAAA`}, {`ABAAAA`}, {`ABAAAA`}, {`ABAAAA`}, {`ABAAAA`}, {`ABAAAA`}, {`ABAAAA`}, {`ABAAAA`}, {`ABAAAA`}, {`ABAAAA`}, {`ABAAAA`}},
			},
			{
				Statement: `ROLLBACK;`,
			},
			{
				Statement: `create temp table tt1(f1 int);`,
			},
			{
				Statement: `create function count_tt1_v() returns int8 as
'select count(*) from tt1' language sql volatile;`,
			},
			{
				Statement: `create function count_tt1_s() returns int8 as
'select count(*) from tt1' language sql stable;`,
			},
			{
				Statement: `begin;`,
			},
			{
				Statement: `insert into tt1 values(1);`,
			},
			{
				Statement: `declare c1 cursor for select count_tt1_v(), count_tt1_s();`,
			},
			{
				Statement: `insert into tt1 values(2);`,
			},
			{
				Statement: `fetch all from c1;`,
				Results:   []sql.Row{{2, 1}},
			},
			{
				Statement: `rollback;`,
			},
			{
				Statement: `begin;`,
			},
			{
				Statement: `insert into tt1 values(1);`,
			},
			{
				Statement: `declare c2 cursor with hold for select count_tt1_v(), count_tt1_s();`,
			},
			{
				Statement: `insert into tt1 values(2);`,
			},
			{
				Statement: `commit;`,
			},
			{
				Statement: `delete from tt1;`,
			},
			{
				Statement: `fetch all from c2;`,
				Results:   []sql.Row{{2, 1}},
			},
			{
				Statement: `drop function count_tt1_v();`,
			},
			{
				Statement: `drop function count_tt1_s();`,
			},
			{
				Statement: `BEGIN;`,
			},
			{
				Statement: `SELECT name, statement, is_holdable, is_binary, is_scrollable FROM pg_cursors;`,
				Results:   []sql.Row{{`c2`, `declare c2 cursor with hold for select count_tt1_v(), count_tt1_s();`, true, false, false}},
			},
			{
				Statement: `DECLARE bc BINARY CURSOR FOR SELECT * FROM tenk1;`,
			},
			{
				Statement: `SELECT name, statement, is_holdable, is_binary, is_scrollable FROM pg_cursors ORDER BY 1;`,
				Results:   []sql.Row{{`bc`, `DECLARE bc BINARY CURSOR FOR SELECT * FROM tenk1;`, false, true, true}, {`c2`, `declare c2 cursor with hold for select count_tt1_v(), count_tt1_s();`, true, false, false}},
			},
			{
				Statement: `ROLLBACK;`,
			},
			{
				Statement: `PREPARE cprep AS
  SELECT name, statement, is_holdable, is_binary, is_scrollable FROM pg_cursors;`,
			},
			{
				Statement: `EXECUTE cprep;`,
				Results:   []sql.Row{{`c2`, `declare c2 cursor with hold for select count_tt1_v(), count_tt1_s();`, true, false, false}},
			},
			{
				Statement: `SELECT name FROM pg_cursors ORDER BY 1;`,
				Results:   []sql.Row{{`c2`}},
			},
			{
				Statement: `CLOSE ALL;`,
			},
			{
				Statement: `SELECT name FROM pg_cursors ORDER BY 1;`,
				Results:   []sql.Row{},
			},
			{
				Statement: `BEGIN;`,
			},
			{
				Statement: `DECLARE foo1 CURSOR WITH HOLD FOR SELECT 1;`,
			},
			{
				Statement: `DECLARE foo2 CURSOR WITHOUT HOLD FOR SELECT 1;`,
			},
			{
				Statement: `SELECT name FROM pg_cursors ORDER BY 1;`,
				Results:   []sql.Row{{`foo1`}, {`foo2`}},
			},
			{
				Statement: `CLOSE ALL;`,
			},
			{
				Statement: `SELECT name FROM pg_cursors ORDER BY 1;`,
				Results:   []sql.Row{},
			},
			{
				Statement: `COMMIT;`,
			},
			{
				Statement: `CREATE TEMP TABLE uctest(f1 int, f2 text);`,
			},
			{
				Statement: `INSERT INTO uctest VALUES (1, 'one'), (2, 'two'), (3, 'three');`,
			},
			{
				Statement: `SELECT * FROM uctest;`,
				Results:   []sql.Row{{1, `one`}, {2, `two`}, {3, `three`}},
			},
			{
				Statement: `BEGIN;`,
			},
			{
				Statement: `DECLARE c1 CURSOR FOR SELECT * FROM uctest;`,
			},
			{
				Statement: `FETCH 2 FROM c1;`,
				Results:   []sql.Row{{1, `one`}, {2, `two`}},
			},
			{
				Statement: `DELETE FROM uctest WHERE CURRENT OF c1;`,
			},
			{
				Statement: `SELECT * FROM uctest;`,
				Results:   []sql.Row{{1, `one`}, {3, `three`}},
			},
			{
				Statement: `FETCH ALL FROM c1;`,
				Results:   []sql.Row{{3, `three`}},
			},
			{
				Statement: `MOVE BACKWARD ALL IN c1;`,
			},
			{
				Statement: `FETCH ALL FROM c1;`,
				Results:   []sql.Row{{1, `one`}, {2, `two`}, {3, `three`}},
			},
			{
				Statement: `COMMIT;`,
			},
			{
				Statement: `SELECT * FROM uctest;`,
				Results:   []sql.Row{{1, `one`}, {3, `three`}},
			},
			{
				Statement: `BEGIN;`,
			},
			{
				Statement: `DECLARE c1 CURSOR FOR SELECT * FROM uctest FOR UPDATE;`,
			},
			{
				Statement: `FETCH c1;`,
				Results:   []sql.Row{{1, `one`}},
			},
			{
				Statement: `UPDATE uctest SET f1 = 8 WHERE CURRENT OF c1;`,
			},
			{
				Statement: `SELECT * FROM uctest;`,
				Results:   []sql.Row{{3, `three`}, {8, `one`}},
			},
			{
				Statement: `COMMIT;`,
			},
			{
				Statement: `SELECT * FROM uctest;`,
				Results:   []sql.Row{{3, `three`}, {8, `one`}},
			},
			{
				Statement: `BEGIN;`,
			},
			{
				Statement: `DECLARE c1 CURSOR FOR SELECT * FROM uctest;`,
			},
			{
				Statement: `FETCH c1;`,
				Results:   []sql.Row{{3, `three`}},
			},
			{
				Statement: `UPDATE uctest SET f1 = f1 + 10 WHERE CURRENT OF c1;`,
			},
			{
				Statement: `SELECT * FROM uctest;`,
				Results:   []sql.Row{{8, `one`}, {13, `three`}},
			},
			{
				Statement: `UPDATE uctest SET f1 = f1 + 10 WHERE CURRENT OF c1;`,
			},
			{
				Statement: `SELECT * FROM uctest;`,
				Results:   []sql.Row{{8, `one`}, {23, `three`}},
			},
			{
				Statement: `FETCH RELATIVE 0 FROM c1;`,
				Results:   []sql.Row{{3, `three`}},
			},
			{
				Statement: `DELETE FROM uctest WHERE CURRENT OF c1;`,
			},
			{
				Statement: `SELECT * FROM uctest;`,
				Results:   []sql.Row{{8, `one`}},
			},
			{
				Statement: `DELETE FROM uctest WHERE CURRENT OF c1; -- no-op`,
			},
			{
				Statement: `SELECT * FROM uctest;`,
				Results:   []sql.Row{{8, `one`}},
			},
			{
				Statement: `UPDATE uctest SET f1 = f1 + 10 WHERE CURRENT OF c1; -- no-op`,
			},
			{
				Statement: `SELECT * FROM uctest;`,
				Results:   []sql.Row{{8, `one`}},
			},
			{
				Statement: `FETCH RELATIVE 0 FROM c1;`,
				Results:   []sql.Row{{3, `three`}},
			},
			{
				Statement: `ROLLBACK;`,
			},
			{
				Statement: `SELECT * FROM uctest;`,
				Results:   []sql.Row{{3, `three`}, {8, `one`}},
			},
			{
				Statement: `BEGIN;`,
			},
			{
				Statement: `DECLARE c1 CURSOR FOR SELECT * FROM uctest FOR UPDATE;`,
			},
			{
				Statement: `FETCH c1;`,
				Results:   []sql.Row{{3, `three`}},
			},
			{
				Statement: `UPDATE uctest SET f1 = f1 + 10 WHERE CURRENT OF c1;`,
			},
			{
				Statement: `SELECT * FROM uctest;`,
				Results:   []sql.Row{{8, `one`}, {13, `three`}},
			},
			{
				Statement: `UPDATE uctest SET f1 = f1 + 10 WHERE CURRENT OF c1;`,
			},
			{
				Statement: `SELECT * FROM uctest;`,
				Results:   []sql.Row{{8, `one`}, {23, `three`}},
			},
			{
				Statement: `DELETE FROM uctest WHERE CURRENT OF c1;`,
			},
			{
				Statement: `SELECT * FROM uctest;`,
				Results:   []sql.Row{{8, `one`}},
			},
			{
				Statement: `DELETE FROM uctest WHERE CURRENT OF c1; -- no-op`,
			},
			{
				Statement: `SELECT * FROM uctest;`,
				Results:   []sql.Row{{8, `one`}},
			},
			{
				Statement: `UPDATE uctest SET f1 = f1 + 10 WHERE CURRENT OF c1; -- no-op`,
			},
			{
				Statement: `SELECT * FROM uctest;`,
				Results:   []sql.Row{{8, `one`}},
			},
			{
				Statement:   `FETCH RELATIVE 0 FROM c1;`,
				ErrorString: `cursor can only scan forward`,
			},
			{
				Statement: `ROLLBACK;`,
			},
			{
				Statement: `SELECT * FROM uctest;`,
				Results:   []sql.Row{{3, `three`}, {8, `one`}},
			},
			{
				Statement: `BEGIN;`,
			},
			{
				Statement: `DECLARE c1 INSENSITIVE CURSOR FOR SELECT * FROM uctest;`,
			},
			{
				Statement: `INSERT INTO uctest VALUES (10, 'ten');`,
			},
			{
				Statement: `FETCH NEXT FROM c1;`,
				Results:   []sql.Row{{3, `three`}},
			},
			{
				Statement: `FETCH NEXT FROM c1;`,
				Results:   []sql.Row{{8, `one`}},
			},
			{
				Statement: `FETCH NEXT FROM c1;  -- insert not visible`,
				Results:   []sql.Row{},
			},
			{
				Statement: `COMMIT;`,
			},
			{
				Statement: `SELECT * FROM uctest;`,
				Results:   []sql.Row{{3, `three`}, {8, `one`}, {10, `ten`}},
			},
			{
				Statement: `DELETE FROM uctest WHERE f1 = 10;  -- restore test table state`,
			},
			{
				Statement: `CREATE TEMP TABLE ucchild () inherits (uctest);`,
			},
			{
				Statement: `INSERT INTO ucchild values(100, 'hundred');`,
			},
			{
				Statement: `SELECT * FROM uctest;`,
				Results:   []sql.Row{{3, `three`}, {8, `one`}, {100, `hundred`}},
			},
			{
				Statement: `BEGIN;`,
			},
			{
				Statement: `DECLARE c1 CURSOR FOR SELECT * FROM uctest FOR UPDATE;`,
			},
			{
				Statement: `FETCH 1 FROM c1;`,
				Results:   []sql.Row{{3, `three`}},
			},
			{
				Statement: `UPDATE uctest SET f1 = f1 + 10 WHERE CURRENT OF c1;`,
			},
			{
				Statement: `FETCH 1 FROM c1;`,
				Results:   []sql.Row{{8, `one`}},
			},
			{
				Statement: `UPDATE uctest SET f1 = f1 + 10 WHERE CURRENT OF c1;`,
			},
			{
				Statement: `FETCH 1 FROM c1;`,
				Results:   []sql.Row{{100, `hundred`}},
			},
			{
				Statement: `UPDATE uctest SET f1 = f1 + 10 WHERE CURRENT OF c1;`,
			},
			{
				Statement: `FETCH 1 FROM c1;`,
				Results:   []sql.Row{},
			},
			{
				Statement: `COMMIT;`,
			},
			{
				Statement: `SELECT * FROM uctest;`,
				Results:   []sql.Row{{13, `three`}, {18, `one`}, {110, `hundred`}},
			},
			{
				Statement: `BEGIN;`,
			},
			{
				Statement: `DECLARE c1 CURSOR FOR SELECT * FROM uctest a, uctest b WHERE a.f1 = b.f1 + 5;`,
			},
			{
				Statement: `FETCH 1 FROM c1;`,
				Results:   []sql.Row{{18, `one`, 13, `three`}},
			},
			{
				Statement:   `UPDATE uctest SET f1 = f1 + 10 WHERE CURRENT OF c1;  -- fail`,
				ErrorString: `cursor "c1" is not a simply updatable scan of table "uctest"`,
			},
			{
				Statement: `ROLLBACK;`,
			},
			{
				Statement: `BEGIN;`,
			},
			{
				Statement: `DECLARE c1 CURSOR FOR SELECT * FROM uctest a, uctest b WHERE a.f1 = b.f1 + 5 FOR UPDATE;`,
			},
			{
				Statement: `FETCH 1 FROM c1;`,
				Results:   []sql.Row{{18, `one`, 13, `three`}},
			},
			{
				Statement:   `UPDATE uctest SET f1 = f1 + 10 WHERE CURRENT OF c1;  -- fail`,
				ErrorString: `cursor "c1" has multiple FOR UPDATE/SHARE references to table "uctest"`,
			},
			{
				Statement: `ROLLBACK;`,
			},
			{
				Statement: `BEGIN;`,
			},
			{
				Statement: `DECLARE c1 CURSOR FOR SELECT * FROM uctest a, uctest b WHERE a.f1 = b.f1 + 5 FOR SHARE OF a;`,
			},
			{
				Statement: `FETCH 1 FROM c1;`,
				Results:   []sql.Row{{18, `one`, 13, `three`}},
			},
			{
				Statement: `UPDATE uctest SET f1 = f1 + 10 WHERE CURRENT OF c1;`,
			},
			{
				Statement: `SELECT * FROM uctest;`,
				Results:   []sql.Row{{13, `three`}, {28, `one`}, {110, `hundred`}},
			},
			{
				Statement: `ROLLBACK;`,
			},
			{
				Statement:   `DELETE FROM uctest WHERE CURRENT OF c1;  -- fail, no such cursor`,
				ErrorString: `cursor "c1" does not exist`,
			},
			{
				Statement: `DECLARE cx CURSOR WITH HOLD FOR SELECT * FROM uctest;`,
			},
			{
				Statement:   `DELETE FROM uctest WHERE CURRENT OF cx;  -- fail, can't use held cursor`,
				ErrorString: `cursor "cx" is held from a previous transaction`,
			},
			{
				Statement: `BEGIN;`,
			},
			{
				Statement: `DECLARE c CURSOR FOR SELECT * FROM tenk2;`,
			},
			{
				Statement:   `DELETE FROM uctest WHERE CURRENT OF c;  -- fail, cursor on wrong table`,
				ErrorString: `cursor "c" is not a simply updatable scan of table "uctest"`,
			},
			{
				Statement: `ROLLBACK;`,
			},
			{
				Statement: `BEGIN;`,
			},
			{
				Statement: `DECLARE c CURSOR FOR SELECT * FROM tenk2 FOR SHARE;`,
			},
			{
				Statement:   `DELETE FROM uctest WHERE CURRENT OF c;  -- fail, cursor on wrong table`,
				ErrorString: `cursor "c" does not have a FOR UPDATE/SHARE reference to table "uctest"`,
			},
			{
				Statement: `ROLLBACK;`,
			},
			{
				Statement: `BEGIN;`,
			},
			{
				Statement: `DECLARE c CURSOR FOR SELECT * FROM tenk1 JOIN tenk2 USING (unique1);`,
			},
			{
				Statement:   `DELETE FROM tenk1 WHERE CURRENT OF c;  -- fail, cursor is on a join`,
				ErrorString: `cursor "c" is not a simply updatable scan of table "tenk1"`,
			},
			{
				Statement: `ROLLBACK;`,
			},
			{
				Statement: `BEGIN;`,
			},
			{
				Statement: `DECLARE c CURSOR FOR SELECT f1,count(*) FROM uctest GROUP BY f1;`,
			},
			{
				Statement:   `DELETE FROM uctest WHERE CURRENT OF c;  -- fail, cursor is on aggregation`,
				ErrorString: `cursor "c" is not a simply updatable scan of table "uctest"`,
			},
			{
				Statement: `ROLLBACK;`,
			},
			{
				Statement: `BEGIN;`,
			},
			{
				Statement: `DECLARE c1 CURSOR FOR SELECT * FROM uctest;`,
			},
			{
				Statement:   `DELETE FROM uctest WHERE CURRENT OF c1; -- fail, no current row`,
				ErrorString: `cursor "c1" is not positioned on a row`,
			},
			{
				Statement: `ROLLBACK;`,
			},
			{
				Statement: `BEGIN;`,
			},
			{
				Statement:   `DECLARE c1 CURSOR FOR SELECT MIN(f1) FROM uctest FOR UPDATE;`,
				ErrorString: `FOR UPDATE is not allowed with aggregate functions`,
			},
			{
				Statement: `ROLLBACK;`,
			},
			{
				Statement: `CREATE TEMP VIEW ucview AS SELECT * FROM uctest;`,
			},
			{
				Statement: `CREATE RULE ucrule AS ON DELETE TO ucview DO INSTEAD
  DELETE FROM uctest WHERE f1 = OLD.f1;`,
			},
			{
				Statement: `BEGIN;`,
			},
			{
				Statement: `DECLARE c1 CURSOR FOR SELECT * FROM ucview;`,
			},
			{
				Statement: `FETCH FROM c1;`,
				Results:   []sql.Row{{13, `three`}},
			},
			{
				Statement:   `DELETE FROM ucview WHERE CURRENT OF c1; -- fail, views not supported`,
				ErrorString: `WHERE CURRENT OF on a view is not implemented`,
			},
			{
				Statement: `ROLLBACK;`,
			},
			{
				Statement: `BEGIN;`,
			},
			{
				Statement: `EXPLAIN (costs off)
DECLARE c1 CURSOR FOR SELECT stringu1 FROM onek WHERE stringu1 = 'DZAAAA';`,
				Results: []sql.Row{{`Index Only Scan using onek_stringu1 on onek`}, {`Index Cond: (stringu1 = 'DZAAAA'::name)`}},
			},
			{
				Statement: `DECLARE c1 CURSOR FOR SELECT stringu1 FROM onek WHERE stringu1 = 'DZAAAA';`,
			},
			{
				Statement: `FETCH FROM c1;`,
				Results:   []sql.Row{{`DZAAAA`}},
			},
			{
				Statement: `DELETE FROM onek WHERE CURRENT OF c1;`,
			},
			{
				Statement: `SELECT stringu1 FROM onek WHERE stringu1 = 'DZAAAA';`,
				Results:   []sql.Row{},
			},
			{
				Statement: `ROLLBACK;`,
			},
			{
				Statement: `BEGIN;`,
			},
			{
				Statement: `CREATE TABLE current_check (currentid int, payload text);`,
			},
			{
				Statement: `CREATE TABLE current_check_1 () INHERITS (current_check);`,
			},
			{
				Statement: `CREATE TABLE current_check_2 () INHERITS (current_check);`,
			},
			{
				Statement: `INSERT INTO current_check_1 SELECT i, 'p' || i FROM generate_series(1,9) i;`,
			},
			{
				Statement: `INSERT INTO current_check_2 SELECT i, 'P' || i FROM generate_series(10,19) i;`,
			},
			{
				Statement: `DECLARE c1 SCROLL CURSOR FOR SELECT * FROM current_check;`,
			},
			{
				Statement: `FETCH ABSOLUTE 12 FROM c1;`,
				Results:   []sql.Row{{12, `P12`}},
			},
			{
				Statement: `FETCH ABSOLUTE 8 FROM c1;`,
				Results:   []sql.Row{{8, `p8`}},
			},
			{
				Statement: `DELETE FROM current_check WHERE CURRENT OF c1 RETURNING *;`,
				Results:   []sql.Row{{8, `p8`}},
			},
			{
				Statement: `FETCH ABSOLUTE 13 FROM c1;`,
				Results:   []sql.Row{{13, `P13`}},
			},
			{
				Statement: `FETCH ABSOLUTE 1 FROM c1;`,
				Results:   []sql.Row{{1, `p1`}},
			},
			{
				Statement: `DELETE FROM current_check WHERE CURRENT OF c1 RETURNING *;`,
				Results:   []sql.Row{{1, `p1`}},
			},
			{
				Statement: `SELECT * FROM current_check;`,
				Results:   []sql.Row{{2, `p2`}, {3, `p3`}, {4, `p4`}, {5, `p5`}, {6, `p6`}, {7, `p7`}, {9, `p9`}, {10, `P10`}, {11, `P11`}, {12, `P12`}, {13, `P13`}, {14, `P14`}, {15, `P15`}, {16, `P16`}, {17, `P17`}, {18, `P18`}, {19, `P19`}},
			},
			{
				Statement: `ROLLBACK;`,
			},
			{
				Statement: `BEGIN;`,
			},
			{
				Statement: `SET TRANSACTION ISOLATION LEVEL SERIALIZABLE;`,
			},
			{
				Statement: `CREATE TABLE cursor (a int);`,
			},
			{
				Statement: `INSERT INTO cursor VALUES (1);`,
			},
			{
				Statement: `DECLARE c1 NO SCROLL CURSOR FOR SELECT * FROM cursor FOR UPDATE;`,
			},
			{
				Statement: `UPDATE cursor SET a = 2;`,
			},
			{
				Statement: `FETCH ALL FROM c1;`,
				Results:   []sql.Row{},
			},
			{
				Statement: `COMMIT;`,
			},
			{
				Statement: `DROP TABLE cursor;`,
			},
			{
				Statement: `begin;`,
			},
			{
				Statement: `create function nochange(int) returns int
  as 'select $1 limit 1' language sql stable;`,
			},
			{
				Statement: `declare c cursor for select * from int8_tbl limit nochange(3);`,
			},
			{
				Statement: `fetch all from c;`,
				Results:   []sql.Row{{123, 456}, {123, 4567890123456789}, {4567890123456789, 123}},
			},
			{
				Statement: `move backward all in c;`,
			},
			{
				Statement: `fetch all from c;`,
				Results:   []sql.Row{{123, 456}, {123, 4567890123456789}, {4567890123456789, 123}},
			},
			{
				Statement: `rollback;`,
			},
			{
				Statement: `begin;`,
			},
			{
				Statement: `explain (costs off) declare c1 cursor for select (select 42) as x;`,
				Results:   []sql.Row{{`Result`}, {`InitPlan 1 (returns $0)`}, {`->  Result`}},
			},
			{
				Statement: `explain (costs off) declare c1 scroll cursor for select (select 42) as x;`,
				Results:   []sql.Row{{`Materialize`}, {`InitPlan 1 (returns $0)`}, {`->  Result`}, {`->  Result`}},
			},
			{
				Statement: `declare c1 scroll cursor for select (select 42) as x;`,
			},
			{
				Statement: `fetch all in c1;`,
				Results:   []sql.Row{{42}},
			},
			{
				Statement: `fetch backward all in c1;`,
				Results:   []sql.Row{{42}},
			},
			{
				Statement: `rollback;`,
			},
			{
				Statement: `begin;`,
			},
			{
				Statement: `explain (costs off) declare c2 cursor for select generate_series(1,3) as g;`,
				Results:   []sql.Row{{`ProjectSet`}, {`->  Result`}},
			},
			{
				Statement: `explain (costs off) declare c2 scroll cursor for select generate_series(1,3) as g;`,
				Results:   []sql.Row{{`Materialize`}, {`->  ProjectSet`}, {`->  Result`}},
			},
			{
				Statement: `declare c2 scroll cursor for select generate_series(1,3) as g;`,
			},
			{
				Statement: `fetch all in c2;`,
				Results:   []sql.Row{{1}, {2}, {3}},
			},
			{
				Statement: `fetch backward all in c2;`,
				Results:   []sql.Row{{3}, {2}, {1}},
			},
			{
				Statement: `rollback;`,
			},
			{
				Statement: `begin;`,
			},
			{
				Statement: `set default_toast_compression = 'pglz';`,
			},
			{
				Statement: `create table toasted_data (f1 int[]);`,
			},
			{
				Statement: `insert into toasted_data
  select array_agg(i) from generate_series(12345678, 12345678 + 1000) i;`,
			},
			{
				Statement: `declare local_portal cursor for select * from toasted_data;`,
			},
			{
				Statement: `fetch all in local_portal;`,
				Results:   []sql.Row{{`{12345678,12345679,12345680,12345681,12345682,12345683,12345684,12345685,12345686,12345687,12345688,12345689,12345690,12345691,12345692,12345693,12345694,12345695,12345696,12345697,12345698,12345699,12345700,12345701,12345702,12345703,12345704,12345705,12345706,12345707,12345708,12345709,12345710,12345711,12345712,12345713,12345714,12345715,12345716,12345717,12345718,12345719,12345720,12345721,12345722,12345723,12345724,12345725,12345726,12345727,12345728,12345729,12345730,12345731,12345732,12345733,12345734,12345735,12345736,12345737,12345738,12345739,12345740,12345741,12345742,12345743,12345744,12345745,12345746,12345747,12345748,12345749,12345750,12345751,12345752,12345753,12345754,12345755,12345756,12345757,12345758,12345759,12345760,12345761,12345762,12345763,12345764,12345765,12345766,12345767,12345768,12345769,12345770,12345771,12345772,12345773,12345774,12345775,12345776,12345777,12345778,12345779,12345780,12345781,12345782,12345783,12345784,12345785,12345786,12345787,12345788,12345789,12345790,12345791,12345792,12345793,12345794,12345795,12345796,12345797,12345798,12345799,12345800,12345801,12345802,12345803,12345804,12345805,12345806,12345807,12345808,12345809,12345810,12345811,12345812,12345813,12345814,12345815,12345816,12345817,12345818,12345819,12345820,12345821,12345822,12345823,12345824,12345825,12345826,12345827,12345828,12345829,12345830,12345831,12345832,12345833,12345834,12345835,12345836,12345837,12345838,12345839,12345840,12345841,12345842,12345843,12345844,12345845,12345846,12345847,12345848,12345849,12345850,12345851,12345852,12345853,12345854,12345855,12345856,12345857,12345858,12345859,12345860,12345861,12345862,12345863,12345864,12345865,12345866,12345867,12345868,12345869,12345870,12345871,12345872,12345873,12345874,12345875,12345876,12345877,12345878,12345879,12345880,12345881,12345882,12345883,12345884,12345885,12345886,12345887,12345888,12345889,12345890,12345891,12345892,12345893,12345894,12345895,12345896,12345897,12345898,12345899,12345900,12345901,12345902,12345903,12345904,12345905,12345906,12345907,12345908,12345909,12345910,12345911,12345912,12345913,12345914,12345915,12345916,12345917,12345918,12345919,12345920,12345921,12345922,12345923,12345924,12345925,12345926,12345927,12345928,12345929,12345930,12345931,12345932,12345933,12345934,12345935,12345936,12345937,12345938,12345939,12345940,12345941,12345942,12345943,12345944,12345945,12345946,12345947,12345948,12345949,12345950,12345951,12345952,12345953,12345954,12345955,12345956,12345957,12345958,12345959,12345960,12345961,12345962,12345963,12345964,12345965,12345966,12345967,12345968,12345969,12345970,12345971,12345972,12345973,12345974,12345975,12345976,12345977,12345978,12345979,12345980,12345981,12345982,12345983,12345984,12345985,12345986,12345987,12345988,12345989,12345990,12345991,12345992,12345993,12345994,12345995,12345996,12345997,12345998,12345999,12346000,12346001,12346002,12346003,12346004,12346005,12346006,12346007,12346008,12346009,12346010,12346011,12346012,12346013,12346014,12346015,12346016,12346017,12346018,12346019,12346020,12346021,12346022,12346023,12346024,12346025,12346026,12346027,12346028,12346029,12346030,12346031,12346032,12346033,12346034,12346035,12346036,12346037,12346038,12346039,12346040,12346041,12346042,12346043,12346044,12346045,12346046,12346047,12346048,12346049,12346050,12346051,12346052,12346053,12346054,12346055,12346056,12346057,12346058,12346059,12346060,12346061,12346062,12346063,12346064,12346065,12346066,12346067,12346068,12346069,12346070,12346071,12346072,12346073,12346074,12346075,12346076,12346077,12346078,12346079,12346080,12346081,12346082,12346083,12346084,12346085,12346086,12346087,12346088,12346089,12346090,12346091,12346092,12346093,12346094,12346095,12346096,12346097,12346098,12346099,12346100,12346101,12346102,12346103,12346104,12346105,12346106,12346107,12346108,12346109,12346110,12346111,12346112,12346113,12346114,12346115,12346116,12346117,12346118,12346119,12346120,12346121,12346122,12346123,12346124,12346125,12346126,12346127,12346128,12346129,12346130,12346131,12346132,12346133,12346134,12346135,12346136,12346137,12346138,12346139,12346140,12346141,12346142,12346143,12346144,12346145,12346146,12346147,12346148,12346149,12346150,12346151,12346152,12346153,12346154,12346155,12346156,12346157,12346158,12346159,12346160,12346161,12346162,12346163,12346164,12346165,12346166,12346167,12346168,12346169,12346170,12346171,12346172,12346173,12346174,12346175,12346176,12346177,12346178,12346179,12346180,12346181,12346182,12346183,12346184,12346185,12346186,12346187,12346188,12346189,12346190,12346191,12346192,12346193,12346194,12346195,12346196,12346197,12346198,12346199,12346200,12346201,12346202,12346203,12346204,12346205,12346206,12346207,12346208,12346209,12346210,12346211,12346212,12346213,12346214,12346215,12346216,12346217,12346218,12346219,12346220,12346221,12346222,12346223,12346224,12346225,12346226,12346227,12346228,12346229,12346230,12346231,12346232,12346233,12346234,12346235,12346236,12346237,12346238,12346239,12346240,12346241,12346242,12346243,12346244,12346245,12346246,12346247,12346248,12346249,12346250,12346251,12346252,12346253,12346254,12346255,12346256,12346257,12346258,12346259,12346260,12346261,12346262,12346263,12346264,12346265,12346266,12346267,12346268,12346269,12346270,12346271,12346272,12346273,12346274,12346275,12346276,12346277,12346278,12346279,12346280,12346281,12346282,12346283,12346284,12346285,12346286,12346287,12346288,12346289,12346290,12346291,12346292,12346293,12346294,12346295,12346296,12346297,12346298,12346299,12346300,12346301,12346302,12346303,12346304,12346305,12346306,12346307,12346308,12346309,12346310,12346311,12346312,12346313,12346314,12346315,12346316,12346317,12346318,12346319,12346320,12346321,12346322,12346323,12346324,12346325,12346326,12346327,12346328,12346329,12346330,12346331,12346332,12346333,12346334,12346335,12346336,12346337,12346338,12346339,12346340,12346341,12346342,12346343,12346344,12346345,12346346,12346347,12346348,12346349,12346350,12346351,12346352,12346353,12346354,12346355,12346356,12346357,12346358,12346359,12346360,12346361,12346362,12346363,12346364,12346365,12346366,12346367,12346368,12346369,12346370,12346371,12346372,12346373,12346374,12346375,12346376,12346377,12346378,12346379,12346380,12346381,12346382,12346383,12346384,12346385,12346386,12346387,12346388,12346389,12346390,12346391,12346392,12346393,12346394,12346395,12346396,12346397,12346398,12346399,12346400,12346401,12346402,12346403,12346404,12346405,12346406,12346407,12346408,12346409,12346410,12346411,12346412,12346413,12346414,12346415,12346416,12346417,12346418,12346419,12346420,12346421,12346422,12346423,12346424,12346425,12346426,12346427,12346428,12346429,12346430,12346431,12346432,12346433,12346434,12346435,12346436,12346437,12346438,12346439,12346440,12346441,12346442,12346443,12346444,12346445,12346446,12346447,12346448,12346449,12346450,12346451,12346452,12346453,12346454,12346455,12346456,12346457,12346458,12346459,12346460,12346461,12346462,12346463,12346464,12346465,12346466,12346467,12346468,12346469,12346470,12346471,12346472,12346473,12346474,12346475,12346476,12346477,12346478,12346479,12346480,12346481,12346482,12346483,12346484,12346485,12346486,12346487,12346488,12346489,12346490,12346491,12346492,12346493,12346494,12346495,12346496,12346497,12346498,12346499,12346500,12346501,12346502,12346503,12346504,12346505,12346506,12346507,12346508,12346509,12346510,12346511,12346512,12346513,12346514,12346515,12346516,12346517,12346518,12346519,12346520,12346521,12346522,12346523,12346524,12346525,12346526,12346527,12346528,12346529,12346530,12346531,12346532,12346533,12346534,12346535,12346536,12346537,12346538,12346539,12346540,12346541,12346542,12346543,12346544,12346545,12346546,12346547,12346548,12346549,12346550,12346551,12346552,12346553,12346554,12346555,12346556,12346557,12346558,12346559,12346560,12346561,12346562,12346563,12346564,12346565,12346566,12346567,12346568,12346569,12346570,12346571,12346572,12346573,12346574,12346575,12346576,12346577,12346578,12346579,12346580,12346581,12346582,12346583,12346584,12346585,12346586,12346587,12346588,12346589,12346590,12346591,12346592,12346593,12346594,12346595,12346596,12346597,12346598,12346599,12346600,12346601,12346602,12346603,12346604,12346605,12346606,12346607,12346608,12346609,12346610,12346611,12346612,12346613,12346614,12346615,12346616,12346617,12346618,12346619,12346620,12346621,12346622,12346623,12346624,12346625,12346626,12346627,12346628,12346629,12346630,12346631,12346632,12346633,12346634,12346635,12346636,12346637,12346638,12346639,12346640,12346641,12346642,12346643,12346644,12346645,12346646,12346647,12346648,12346649,12346650,12346651,12346652,12346653,12346654,12346655,12346656,12346657,12346658,12346659,12346660,12346661,12346662,12346663,12346664,12346665,12346666,12346667,12346668,12346669,12346670,12346671,12346672,12346673,12346674,12346675,12346676,12346677,12346678}`}},
			},
			{
				Statement: `declare held_portal cursor with hold for select * from toasted_data;`,
			},
			{
				Statement: `commit;`,
			},
			{
				Statement: `drop table toasted_data;`,
			},
			{
				Statement: `fetch all in held_portal;`,
				Results:   []sql.Row{{`{12345678,12345679,12345680,12345681,12345682,12345683,12345684,12345685,12345686,12345687,12345688,12345689,12345690,12345691,12345692,12345693,12345694,12345695,12345696,12345697,12345698,12345699,12345700,12345701,12345702,12345703,12345704,12345705,12345706,12345707,12345708,12345709,12345710,12345711,12345712,12345713,12345714,12345715,12345716,12345717,12345718,12345719,12345720,12345721,12345722,12345723,12345724,12345725,12345726,12345727,12345728,12345729,12345730,12345731,12345732,12345733,12345734,12345735,12345736,12345737,12345738,12345739,12345740,12345741,12345742,12345743,12345744,12345745,12345746,12345747,12345748,12345749,12345750,12345751,12345752,12345753,12345754,12345755,12345756,12345757,12345758,12345759,12345760,12345761,12345762,12345763,12345764,12345765,12345766,12345767,12345768,12345769,12345770,12345771,12345772,12345773,12345774,12345775,12345776,12345777,12345778,12345779,12345780,12345781,12345782,12345783,12345784,12345785,12345786,12345787,12345788,12345789,12345790,12345791,12345792,12345793,12345794,12345795,12345796,12345797,12345798,12345799,12345800,12345801,12345802,12345803,12345804,12345805,12345806,12345807,12345808,12345809,12345810,12345811,12345812,12345813,12345814,12345815,12345816,12345817,12345818,12345819,12345820,12345821,12345822,12345823,12345824,12345825,12345826,12345827,12345828,12345829,12345830,12345831,12345832,12345833,12345834,12345835,12345836,12345837,12345838,12345839,12345840,12345841,12345842,12345843,12345844,12345845,12345846,12345847,12345848,12345849,12345850,12345851,12345852,12345853,12345854,12345855,12345856,12345857,12345858,12345859,12345860,12345861,12345862,12345863,12345864,12345865,12345866,12345867,12345868,12345869,12345870,12345871,12345872,12345873,12345874,12345875,12345876,12345877,12345878,12345879,12345880,12345881,12345882,12345883,12345884,12345885,12345886,12345887,12345888,12345889,12345890,12345891,12345892,12345893,12345894,12345895,12345896,12345897,12345898,12345899,12345900,12345901,12345902,12345903,12345904,12345905,12345906,12345907,12345908,12345909,12345910,12345911,12345912,12345913,12345914,12345915,12345916,12345917,12345918,12345919,12345920,12345921,12345922,12345923,12345924,12345925,12345926,12345927,12345928,12345929,12345930,12345931,12345932,12345933,12345934,12345935,12345936,12345937,12345938,12345939,12345940,12345941,12345942,12345943,12345944,12345945,12345946,12345947,12345948,12345949,12345950,12345951,12345952,12345953,12345954,12345955,12345956,12345957,12345958,12345959,12345960,12345961,12345962,12345963,12345964,12345965,12345966,12345967,12345968,12345969,12345970,12345971,12345972,12345973,12345974,12345975,12345976,12345977,12345978,12345979,12345980,12345981,12345982,12345983,12345984,12345985,12345986,12345987,12345988,12345989,12345990,12345991,12345992,12345993,12345994,12345995,12345996,12345997,12345998,12345999,12346000,12346001,12346002,12346003,12346004,12346005,12346006,12346007,12346008,12346009,12346010,12346011,12346012,12346013,12346014,12346015,12346016,12346017,12346018,12346019,12346020,12346021,12346022,12346023,12346024,12346025,12346026,12346027,12346028,12346029,12346030,12346031,12346032,12346033,12346034,12346035,12346036,12346037,12346038,12346039,12346040,12346041,12346042,12346043,12346044,12346045,12346046,12346047,12346048,12346049,12346050,12346051,12346052,12346053,12346054,12346055,12346056,12346057,12346058,12346059,12346060,12346061,12346062,12346063,12346064,12346065,12346066,12346067,12346068,12346069,12346070,12346071,12346072,12346073,12346074,12346075,12346076,12346077,12346078,12346079,12346080,12346081,12346082,12346083,12346084,12346085,12346086,12346087,12346088,12346089,12346090,12346091,12346092,12346093,12346094,12346095,12346096,12346097,12346098,12346099,12346100,12346101,12346102,12346103,12346104,12346105,12346106,12346107,12346108,12346109,12346110,12346111,12346112,12346113,12346114,12346115,12346116,12346117,12346118,12346119,12346120,12346121,12346122,12346123,12346124,12346125,12346126,12346127,12346128,12346129,12346130,12346131,12346132,12346133,12346134,12346135,12346136,12346137,12346138,12346139,12346140,12346141,12346142,12346143,12346144,12346145,12346146,12346147,12346148,12346149,12346150,12346151,12346152,12346153,12346154,12346155,12346156,12346157,12346158,12346159,12346160,12346161,12346162,12346163,12346164,12346165,12346166,12346167,12346168,12346169,12346170,12346171,12346172,12346173,12346174,12346175,12346176,12346177,12346178,12346179,12346180,12346181,12346182,12346183,12346184,12346185,12346186,12346187,12346188,12346189,12346190,12346191,12346192,12346193,12346194,12346195,12346196,12346197,12346198,12346199,12346200,12346201,12346202,12346203,12346204,12346205,12346206,12346207,12346208,12346209,12346210,12346211,12346212,12346213,12346214,12346215,12346216,12346217,12346218,12346219,12346220,12346221,12346222,12346223,12346224,12346225,12346226,12346227,12346228,12346229,12346230,12346231,12346232,12346233,12346234,12346235,12346236,12346237,12346238,12346239,12346240,12346241,12346242,12346243,12346244,12346245,12346246,12346247,12346248,12346249,12346250,12346251,12346252,12346253,12346254,12346255,12346256,12346257,12346258,12346259,12346260,12346261,12346262,12346263,12346264,12346265,12346266,12346267,12346268,12346269,12346270,12346271,12346272,12346273,12346274,12346275,12346276,12346277,12346278,12346279,12346280,12346281,12346282,12346283,12346284,12346285,12346286,12346287,12346288,12346289,12346290,12346291,12346292,12346293,12346294,12346295,12346296,12346297,12346298,12346299,12346300,12346301,12346302,12346303,12346304,12346305,12346306,12346307,12346308,12346309,12346310,12346311,12346312,12346313,12346314,12346315,12346316,12346317,12346318,12346319,12346320,12346321,12346322,12346323,12346324,12346325,12346326,12346327,12346328,12346329,12346330,12346331,12346332,12346333,12346334,12346335,12346336,12346337,12346338,12346339,12346340,12346341,12346342,12346343,12346344,12346345,12346346,12346347,12346348,12346349,12346350,12346351,12346352,12346353,12346354,12346355,12346356,12346357,12346358,12346359,12346360,12346361,12346362,12346363,12346364,12346365,12346366,12346367,12346368,12346369,12346370,12346371,12346372,12346373,12346374,12346375,12346376,12346377,12346378,12346379,12346380,12346381,12346382,12346383,12346384,12346385,12346386,12346387,12346388,12346389,12346390,12346391,12346392,12346393,12346394,12346395,12346396,12346397,12346398,12346399,12346400,12346401,12346402,12346403,12346404,12346405,12346406,12346407,12346408,12346409,12346410,12346411,12346412,12346413,12346414,12346415,12346416,12346417,12346418,12346419,12346420,12346421,12346422,12346423,12346424,12346425,12346426,12346427,12346428,12346429,12346430,12346431,12346432,12346433,12346434,12346435,12346436,12346437,12346438,12346439,12346440,12346441,12346442,12346443,12346444,12346445,12346446,12346447,12346448,12346449,12346450,12346451,12346452,12346453,12346454,12346455,12346456,12346457,12346458,12346459,12346460,12346461,12346462,12346463,12346464,12346465,12346466,12346467,12346468,12346469,12346470,12346471,12346472,12346473,12346474,12346475,12346476,12346477,12346478,12346479,12346480,12346481,12346482,12346483,12346484,12346485,12346486,12346487,12346488,12346489,12346490,12346491,12346492,12346493,12346494,12346495,12346496,12346497,12346498,12346499,12346500,12346501,12346502,12346503,12346504,12346505,12346506,12346507,12346508,12346509,12346510,12346511,12346512,12346513,12346514,12346515,12346516,12346517,12346518,12346519,12346520,12346521,12346522,12346523,12346524,12346525,12346526,12346527,12346528,12346529,12346530,12346531,12346532,12346533,12346534,12346535,12346536,12346537,12346538,12346539,12346540,12346541,12346542,12346543,12346544,12346545,12346546,12346547,12346548,12346549,12346550,12346551,12346552,12346553,12346554,12346555,12346556,12346557,12346558,12346559,12346560,12346561,12346562,12346563,12346564,12346565,12346566,12346567,12346568,12346569,12346570,12346571,12346572,12346573,12346574,12346575,12346576,12346577,12346578,12346579,12346580,12346581,12346582,12346583,12346584,12346585,12346586,12346587,12346588,12346589,12346590,12346591,12346592,12346593,12346594,12346595,12346596,12346597,12346598,12346599,12346600,12346601,12346602,12346603,12346604,12346605,12346606,12346607,12346608,12346609,12346610,12346611,12346612,12346613,12346614,12346615,12346616,12346617,12346618,12346619,12346620,12346621,12346622,12346623,12346624,12346625,12346626,12346627,12346628,12346629,12346630,12346631,12346632,12346633,12346634,12346635,12346636,12346637,12346638,12346639,12346640,12346641,12346642,12346643,12346644,12346645,12346646,12346647,12346648,12346649,12346650,12346651,12346652,12346653,12346654,12346655,12346656,12346657,12346658,12346659,12346660,12346661,12346662,12346663,12346664,12346665,12346666,12346667,12346668,12346669,12346670,12346671,12346672,12346673,12346674,12346675,12346676,12346677,12346678}`}},
			},
			{
				Statement: `reset default_toast_compression;`,
			},
		},
	})
}
