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

func TestNumeric(t *testing.T) {
	t.Skip()
	_ = RunTests(t, RegressionFileName_numeric)
}

func init() {
	RegisterRegressionFile(RegressionFile{
		RegressionFileName: RegressionFileName_numeric,
		DependsOn:          []RegressionFileName{RegressionFileName_test_setup},
		Statements: []RegressionFileStatement{
			{
				Statement: `CREATE TABLE num_data (id int4, val numeric(210,10));`,
			},
			{
				Statement: `CREATE TABLE num_exp_add (id1 int4, id2 int4, expected numeric(210,10));`,
			},
			{
				Statement: `CREATE TABLE num_exp_sub (id1 int4, id2 int4, expected numeric(210,10));`,
			},
			{
				Statement: `CREATE TABLE num_exp_div (id1 int4, id2 int4, expected numeric(210,10));`,
			},
			{
				Statement: `CREATE TABLE num_exp_mul (id1 int4, id2 int4, expected numeric(210,10));`,
			},
			{
				Statement: `CREATE TABLE num_exp_sqrt (id int4, expected numeric(210,10));`,
			},
			{
				Statement: `CREATE TABLE num_exp_ln (id int4, expected numeric(210,10));`,
			},
			{
				Statement: `CREATE TABLE num_exp_log10 (id int4, expected numeric(210,10));`,
			},
			{
				Statement: `CREATE TABLE num_exp_power_10_ln (id int4, expected numeric(210,10));`,
			},
			{
				Statement: `CREATE TABLE num_result (id1 int4, id2 int4, result numeric(210,10));`,
			},
			{
				Statement: `BEGIN TRANSACTION;`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (0,0,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (0,0,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (0,0,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (0,0,'NaN');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (0,1,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (0,1,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (0,1,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (0,1,'NaN');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (0,2,'-34338492.215397047');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (0,2,'34338492.215397047');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (0,2,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (0,2,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (0,3,'4.31');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (0,3,'-4.31');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (0,3,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (0,3,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (0,4,'7799461.4119');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (0,4,'-7799461.4119');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (0,4,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (0,4,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (0,5,'16397.038491');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (0,5,'-16397.038491');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (0,5,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (0,5,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (0,6,'93901.57763026');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (0,6,'-93901.57763026');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (0,6,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (0,6,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (0,7,'-83028485');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (0,7,'83028485');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (0,7,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (0,7,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (0,8,'74881');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (0,8,'-74881');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (0,8,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (0,8,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (0,9,'-24926804.045047420');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (0,9,'24926804.045047420');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (0,9,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (0,9,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (1,0,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (1,0,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (1,0,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (1,0,'NaN');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (1,1,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (1,1,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (1,1,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (1,1,'NaN');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (1,2,'-34338492.215397047');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (1,2,'34338492.215397047');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (1,2,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (1,2,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (1,3,'4.31');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (1,3,'-4.31');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (1,3,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (1,3,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (1,4,'7799461.4119');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (1,4,'-7799461.4119');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (1,4,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (1,4,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (1,5,'16397.038491');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (1,5,'-16397.038491');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (1,5,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (1,5,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (1,6,'93901.57763026');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (1,6,'-93901.57763026');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (1,6,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (1,6,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (1,7,'-83028485');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (1,7,'83028485');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (1,7,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (1,7,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (1,8,'74881');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (1,8,'-74881');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (1,8,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (1,8,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (1,9,'-24926804.045047420');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (1,9,'24926804.045047420');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (1,9,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (1,9,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (2,0,'-34338492.215397047');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (2,0,'-34338492.215397047');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (2,0,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (2,0,'NaN');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (2,1,'-34338492.215397047');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (2,1,'-34338492.215397047');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (2,1,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (2,1,'NaN');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (2,2,'-68676984.430794094');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (2,2,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (2,2,'1179132047626883.596862135856320209');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (2,2,'1.00000000000000000000');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (2,3,'-34338487.905397047');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (2,3,'-34338496.525397047');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (2,3,'-147998901.44836127257');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (2,3,'-7967167.56737750510440835266');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (2,4,'-26539030.803497047');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (2,4,'-42137953.627297047');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (2,4,'-267821744976817.8111137106593');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (2,4,'-4.40267480046830116685');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (2,5,'-34322095.176906047');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (2,5,'-34354889.253888047');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (2,5,'-563049578578.769242506736077');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (2,5,'-2094.18866914563535496429');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (2,6,'-34244590.637766787');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (2,6,'-34432393.793027307');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (2,6,'-3224438592470.18449811926184222');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (2,6,'-365.68599891479766440940');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (2,7,'-117366977.215397047');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (2,7,'48689992.784602953');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (2,7,'2851072985828710.485883795');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (2,7,'.41357483778485235518');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (2,8,'-34263611.215397047');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (2,8,'-34413373.215397047');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (2,8,'-2571300635581.146276407');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (2,8,'-458.57416721727870888476');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (2,9,'-59265296.260444467');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (2,9,'-9411688.170349627');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (2,9,'855948866655588.453741509242968740');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (2,9,'1.37757299946438931811');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (3,0,'4.31');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (3,0,'4.31');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (3,0,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (3,0,'NaN');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (3,1,'4.31');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (3,1,'4.31');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (3,1,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (3,1,'NaN');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (3,2,'-34338487.905397047');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (3,2,'34338496.525397047');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (3,2,'-147998901.44836127257');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (3,2,'-.00000012551512084352');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (3,3,'8.62');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (3,3,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (3,3,'18.5761');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (3,3,'1.00000000000000000000');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (3,4,'7799465.7219');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (3,4,'-7799457.1019');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (3,4,'33615678.685289');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (3,4,'.00000055260225961552');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (3,5,'16401.348491');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (3,5,'-16392.728491');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (3,5,'70671.23589621');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (3,5,'.00026285234387695504');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (3,6,'93905.88763026');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (3,6,'-93897.26763026');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (3,6,'404715.7995864206');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (3,6,'.00004589912234457595');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (3,7,'-83028480.69');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (3,7,'83028489.31');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (3,7,'-357852770.35');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (3,7,'-.00000005190989574240');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (3,8,'74885.31');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (3,8,'-74876.69');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (3,8,'322737.11');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (3,8,'.00005755799201399553');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (3,9,'-24926799.735047420');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (3,9,'24926808.355047420');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (3,9,'-107434525.43415438020');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (3,9,'-.00000017290624149854');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (4,0,'7799461.4119');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (4,0,'7799461.4119');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (4,0,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (4,0,'NaN');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (4,1,'7799461.4119');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (4,1,'7799461.4119');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (4,1,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (4,1,'NaN');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (4,2,'-26539030.803497047');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (4,2,'42137953.627297047');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (4,2,'-267821744976817.8111137106593');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (4,2,'-.22713465002993920385');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (4,3,'7799465.7219');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (4,3,'7799457.1019');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (4,3,'33615678.685289');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (4,3,'1809619.81714617169373549883');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (4,4,'15598922.8238');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (4,4,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (4,4,'60831598315717.14146161');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (4,4,'1.00000000000000000000');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (4,5,'7815858.450391');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (4,5,'7783064.373409');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (4,5,'127888068979.9935054429');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (4,5,'475.66281046305802686061');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (4,6,'7893362.98953026');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (4,6,'7705559.83426974');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (4,6,'732381731243.745115764094');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (4,6,'83.05996138436129499606');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (4,7,'-75229023.5881');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (4,7,'90827946.4119');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (4,7,'-647577464846017.9715');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (4,7,'-.09393717604145131637');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (4,8,'7874342.4119');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (4,8,'7724580.4119');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (4,8,'584031469984.4839');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (4,8,'104.15808298366741897143');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (4,9,'-17127342.633147420');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (4,9,'32726265.456947420');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (4,9,'-194415646271340.1815956522980');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (4,9,'-.31289456112403769409');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (5,0,'16397.038491');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (5,0,'16397.038491');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (5,0,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (5,0,'NaN');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (5,1,'16397.038491');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (5,1,'16397.038491');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (5,1,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (5,1,'NaN');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (5,2,'-34322095.176906047');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (5,2,'34354889.253888047');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (5,2,'-563049578578.769242506736077');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (5,2,'-.00047751189505192446');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (5,3,'16401.348491');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (5,3,'16392.728491');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (5,3,'70671.23589621');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (5,3,'3804.41728329466357308584');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (5,4,'7815858.450391');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (5,4,'-7783064.373409');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (5,4,'127888068979.9935054429');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (5,4,'.00210232958726897192');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (5,5,'32794.076982');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (5,5,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (5,5,'268862871.275335557081');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (5,5,'1.00000000000000000000');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (5,6,'110298.61612126');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (5,6,'-77504.53913926');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (5,6,'1539707782.76899778633766');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (5,6,'.17461941433576102689');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (5,7,'-83012087.961509');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (5,7,'83044882.038491');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (5,7,'-1361421264394.416135');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (5,7,'-.00019748690453643710');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (5,8,'91278.038491');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (5,8,'-58483.961509');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (5,8,'1227826639.244571');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (5,8,'.21897461960978085228');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (5,9,'-24910407.006556420');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (5,9,'24943201.083538420');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (5,9,'-408725765384.257043660243220');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (5,9,'-.00065780749354660427');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (6,0,'93901.57763026');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (6,0,'93901.57763026');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (6,0,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (6,0,'NaN');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (6,1,'93901.57763026');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (6,1,'93901.57763026');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (6,1,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (6,1,'NaN');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (6,2,'-34244590.637766787');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (6,2,'34432393.793027307');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (6,2,'-3224438592470.18449811926184222');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (6,2,'-.00273458651128995823');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (6,3,'93905.88763026');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (6,3,'93897.26763026');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (6,3,'404715.7995864206');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (6,3,'21786.90896293735498839907');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (6,4,'7893362.98953026');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (6,4,'-7705559.83426974');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (6,4,'732381731243.745115764094');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (6,4,'.01203949512295682469');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (6,5,'110298.61612126');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (6,5,'77504.53913926');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (6,5,'1539707782.76899778633766');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (6,5,'5.72674008674192359679');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (6,6,'187803.15526052');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (6,6,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (6,6,'8817506281.4517452372676676');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (6,6,'1.00000000000000000000');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (6,7,'-82934583.42236974');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (6,7,'83122386.57763026');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (6,7,'-7796505729750.37795610');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (6,7,'-.00113095617281538980');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (6,8,'168782.57763026');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (6,8,'19020.57763026');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (6,8,'7031444034.53149906');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (6,8,'1.25401073209839612184');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (6,9,'-24832902.467417160');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (6,9,'25020705.622677680');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (6,9,'-2340666225110.29929521292692920');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (6,9,'-.00376709254265256789');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (7,0,'-83028485');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (7,0,'-83028485');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (7,0,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (7,0,'NaN');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (7,1,'-83028485');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (7,1,'-83028485');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (7,1,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (7,1,'NaN');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (7,2,'-117366977.215397047');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (7,2,'-48689992.784602953');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (7,2,'2851072985828710.485883795');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (7,2,'2.41794207151503385700');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (7,3,'-83028480.69');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (7,3,'-83028489.31');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (7,3,'-357852770.35');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (7,3,'-19264149.65197215777262180974');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (7,4,'-75229023.5881');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (7,4,'-90827946.4119');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (7,4,'-647577464846017.9715');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (7,4,'-10.64541262725136247686');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (7,5,'-83012087.961509');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (7,5,'-83044882.038491');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (7,5,'-1361421264394.416135');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (7,5,'-5063.62688881730941836574');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (7,6,'-82934583.42236974');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (7,6,'-83122386.57763026');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (7,6,'-7796505729750.37795610');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (7,6,'-884.20756174009028770294');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (7,7,'-166056970');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (7,7,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (7,7,'6893729321395225');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (7,7,'1.00000000000000000000');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (7,8,'-82953604');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (7,8,'-83103366');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (7,8,'-6217255985285');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (7,8,'-1108.80577182462841041118');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (7,9,'-107955289.045047420');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (7,9,'-58101680.954952580');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (7,9,'2069634775752159.035758700');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (7,9,'3.33089171198810413382');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (8,0,'74881');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (8,0,'74881');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (8,0,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (8,0,'NaN');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (8,1,'74881');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (8,1,'74881');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (8,1,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (8,1,'NaN');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (8,2,'-34263611.215397047');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (8,2,'34413373.215397047');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (8,2,'-2571300635581.146276407');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (8,2,'-.00218067233500788615');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (8,3,'74885.31');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (8,3,'74876.69');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (8,3,'322737.11');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (8,3,'17373.78190255220417633410');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (8,4,'7874342.4119');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (8,4,'-7724580.4119');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (8,4,'584031469984.4839');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (8,4,'.00960079113741758956');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (8,5,'91278.038491');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (8,5,'58483.961509');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (8,5,'1227826639.244571');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (8,5,'4.56673929509287019456');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (8,6,'168782.57763026');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (8,6,'-19020.57763026');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (8,6,'7031444034.53149906');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (8,6,'.79744134113322314424');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (8,7,'-82953604');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (8,7,'83103366');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (8,7,'-6217255985285');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (8,7,'-.00090187120721280172');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (8,8,'149762');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (8,8,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (8,8,'5607164161');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (8,8,'1.00000000000000000000');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (8,9,'-24851923.045047420');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (8,9,'25001685.045047420');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (8,9,'-1866544013697.195857020');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (8,9,'-.00300403532938582735');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (9,0,'-24926804.045047420');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (9,0,'-24926804.045047420');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (9,0,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (9,0,'NaN');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (9,1,'-24926804.045047420');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (9,1,'-24926804.045047420');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (9,1,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (9,1,'NaN');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (9,2,'-59265296.260444467');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (9,2,'9411688.170349627');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (9,2,'855948866655588.453741509242968740');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (9,2,'.72591434384152961526');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (9,3,'-24926799.735047420');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (9,3,'-24926808.355047420');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (9,3,'-107434525.43415438020');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (9,3,'-5783481.21694835730858468677');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (9,4,'-17127342.633147420');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (9,4,'-32726265.456947420');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (9,4,'-194415646271340.1815956522980');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (9,4,'-3.19596478892958416484');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (9,5,'-24910407.006556420');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (9,5,'-24943201.083538420');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (9,5,'-408725765384.257043660243220');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (9,5,'-1520.20159364322004505807');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (9,6,'-24832902.467417160');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (9,6,'-25020705.622677680');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (9,6,'-2340666225110.29929521292692920');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (9,6,'-265.45671195426965751280');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (9,7,'-107955289.045047420');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (9,7,'58101680.954952580');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (9,7,'2069634775752159.035758700');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (9,7,'.30021990699995814689');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (9,8,'-24851923.045047420');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (9,8,'-25001685.045047420');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (9,8,'-1866544013697.195857020');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (9,8,'-332.88556569820675471748');`,
			},
			{
				Statement: `INSERT INTO num_exp_add VALUES (9,9,'-49853608.090094840');`,
			},
			{
				Statement: `INSERT INTO num_exp_sub VALUES (9,9,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_mul VALUES (9,9,'621345559900192.420120630048656400');`,
			},
			{
				Statement: `INSERT INTO num_exp_div VALUES (9,9,'1.00000000000000000000');`,
			},
			{
				Statement: `COMMIT TRANSACTION;`,
			},
			{
				Statement: `BEGIN TRANSACTION;`,
			},
			{
				Statement: `INSERT INTO num_exp_sqrt VALUES (0,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_sqrt VALUES (1,'0');`,
			},
			{
				Statement: `INSERT INTO num_exp_sqrt VALUES (2,'5859.90547836712524903505');`,
			},
			{
				Statement: `INSERT INTO num_exp_sqrt VALUES (3,'2.07605394920266944396');`,
			},
			{
				Statement: `INSERT INTO num_exp_sqrt VALUES (4,'2792.75158435189147418923');`,
			},
			{
				Statement: `INSERT INTO num_exp_sqrt VALUES (5,'128.05092147657509145473');`,
			},
			{
				Statement: `INSERT INTO num_exp_sqrt VALUES (6,'306.43364311096782703406');`,
			},
			{
				Statement: `INSERT INTO num_exp_sqrt VALUES (7,'9111.99676251039939975230');`,
			},
			{
				Statement: `INSERT INTO num_exp_sqrt VALUES (8,'273.64392922189960397542');`,
			},
			{
				Statement: `INSERT INTO num_exp_sqrt VALUES (9,'4992.67503899937593364766');`,
			},
			{
				Statement: `COMMIT TRANSACTION;`,
			},
			{
				Statement: `BEGIN TRANSACTION;`,
			},
			{
				Statement: `INSERT INTO num_exp_ln VALUES (0,'NaN');`,
			},
			{
				Statement: `INSERT INTO num_exp_ln VALUES (1,'NaN');`,
			},
			{
				Statement: `INSERT INTO num_exp_ln VALUES (2,'17.35177750493897715514');`,
			},
			{
				Statement: `INSERT INTO num_exp_ln VALUES (3,'1.46093790411565641971');`,
			},
			{
				Statement: `INSERT INTO num_exp_ln VALUES (4,'15.86956523951936572464');`,
			},
			{
				Statement: `INSERT INTO num_exp_ln VALUES (5,'9.70485601768871834038');`,
			},
			{
				Statement: `INSERT INTO num_exp_ln VALUES (6,'11.45000246622944403127');`,
			},
			{
				Statement: `INSERT INTO num_exp_ln VALUES (7,'18.23469429965478772991');`,
			},
			{
				Statement: `INSERT INTO num_exp_ln VALUES (8,'11.22365546576315513668');`,
			},
			{
				Statement: `INSERT INTO num_exp_ln VALUES (9,'17.03145425013166006962');`,
			},
			{
				Statement: `COMMIT TRANSACTION;`,
			},
			{
				Statement: `BEGIN TRANSACTION;`,
			},
			{
				Statement: `INSERT INTO num_exp_log10 VALUES (0,'NaN');`,
			},
			{
				Statement: `INSERT INTO num_exp_log10 VALUES (1,'NaN');`,
			},
			{
				Statement: `INSERT INTO num_exp_log10 VALUES (2,'7.53578122160797276459');`,
			},
			{
				Statement: `INSERT INTO num_exp_log10 VALUES (3,'.63447727016073160075');`,
			},
			{
				Statement: `INSERT INTO num_exp_log10 VALUES (4,'6.89206461372691743345');`,
			},
			{
				Statement: `INSERT INTO num_exp_log10 VALUES (5,'4.21476541614777768626');`,
			},
			{
				Statement: `INSERT INTO num_exp_log10 VALUES (6,'4.97267288886207207671');`,
			},
			{
				Statement: `INSERT INTO num_exp_log10 VALUES (7,'7.91922711353275546914');`,
			},
			{
				Statement: `INSERT INTO num_exp_log10 VALUES (8,'4.87437163556421004138');`,
			},
			{
				Statement: `INSERT INTO num_exp_log10 VALUES (9,'7.39666659961986567059');`,
			},
			{
				Statement: `COMMIT TRANSACTION;`,
			},
			{
				Statement: `BEGIN TRANSACTION;`,
			},
			{
				Statement: `INSERT INTO num_exp_power_10_ln VALUES (0,'NaN');`,
			},
			{
				Statement: `INSERT INTO num_exp_power_10_ln VALUES (1,'NaN');`,
			},
			{
				Statement: `INSERT INTO num_exp_power_10_ln VALUES (2,'224790267919917955.13261618583642653184');`,
			},
			{
				Statement: `INSERT INTO num_exp_power_10_ln VALUES (3,'28.90266599445155957393');`,
			},
			{
				Statement: `INSERT INTO num_exp_power_10_ln VALUES (4,'7405685069594999.07733999469386277636');`,
			},
			{
				Statement: `INSERT INTO num_exp_power_10_ln VALUES (5,'5068226527.32127265408584640098');`,
			},
			{
				Statement: `INSERT INTO num_exp_power_10_ln VALUES (6,'281839893606.99372343357047819067');`,
			},
			{
				Statement: `INSERT INTO num_exp_power_10_ln VALUES (7,'1716699575118597095.42330819910640247627');`,
			},
			{
				Statement: `INSERT INTO num_exp_power_10_ln VALUES (8,'167361463828.07491320069016125952');`,
			},
			{
				Statement: `INSERT INTO num_exp_power_10_ln VALUES (9,'107511333880052007.04141124673540337457');`,
			},
			{
				Statement: `COMMIT TRANSACTION;`,
			},
			{
				Statement: `BEGIN TRANSACTION;`,
			},
			{
				Statement: `INSERT INTO num_data VALUES (0, '0');`,
			},
			{
				Statement: `INSERT INTO num_data VALUES (1, '0');`,
			},
			{
				Statement: `INSERT INTO num_data VALUES (2, '-34338492.215397047');`,
			},
			{
				Statement: `INSERT INTO num_data VALUES (3, '4.31');`,
			},
			{
				Statement: `INSERT INTO num_data VALUES (4, '7799461.4119');`,
			},
			{
				Statement: `INSERT INTO num_data VALUES (5, '16397.038491');`,
			},
			{
				Statement: `INSERT INTO num_data VALUES (6, '93901.57763026');`,
			},
			{
				Statement: `INSERT INTO num_data VALUES (7, '-83028485');`,
			},
			{
				Statement: `INSERT INTO num_data VALUES (8, '74881');`,
			},
			{
				Statement: `INSERT INTO num_data VALUES (9, '-24926804.045047420');`,
			},
			{
				Statement: `COMMIT TRANSACTION;`,
			},
			{
				Statement: `CREATE UNIQUE INDEX num_exp_add_idx ON num_exp_add (id1, id2);`,
			},
			{
				Statement: `CREATE UNIQUE INDEX num_exp_sub_idx ON num_exp_sub (id1, id2);`,
			},
			{
				Statement: `CREATE UNIQUE INDEX num_exp_div_idx ON num_exp_div (id1, id2);`,
			},
			{
				Statement: `CREATE UNIQUE INDEX num_exp_mul_idx ON num_exp_mul (id1, id2);`,
			},
			{
				Statement: `CREATE UNIQUE INDEX num_exp_sqrt_idx ON num_exp_sqrt (id);`,
			},
			{
				Statement: `CREATE UNIQUE INDEX num_exp_ln_idx ON num_exp_ln (id);`,
			},
			{
				Statement: `CREATE UNIQUE INDEX num_exp_log10_idx ON num_exp_log10 (id);`,
			},
			{
				Statement: `CREATE UNIQUE INDEX num_exp_power_10_ln_idx ON num_exp_power_10_ln (id);`,
			},
			{
				Statement: `VACUUM ANALYZE num_exp_add;`,
			},
			{
				Statement: `VACUUM ANALYZE num_exp_sub;`,
			},
			{
				Statement: `VACUUM ANALYZE num_exp_div;`,
			},
			{
				Statement: `VACUUM ANALYZE num_exp_mul;`,
			},
			{
				Statement: `VACUUM ANALYZE num_exp_sqrt;`,
			},
			{
				Statement: `VACUUM ANALYZE num_exp_ln;`,
			},
			{
				Statement: `VACUUM ANALYZE num_exp_log10;`,
			},
			{
				Statement: `VACUUM ANALYZE num_exp_power_10_ln;`,
			},
			{
				Statement: `DELETE FROM num_result;`,
			},
			{
				Statement: `INSERT INTO num_result SELECT t1.id, t2.id, t1.val + t2.val
    FROM num_data t1, num_data t2;`,
			},
			{
				Statement: `SELECT t1.id1, t1.id2, t1.result, t2.expected
    FROM num_result t1, num_exp_add t2
    WHERE t1.id1 = t2.id1 AND t1.id2 = t2.id2
    AND t1.result != t2.expected;`,
				Results: []sql.Row{},
			},
			{
				Statement: `DELETE FROM num_result;`,
			},
			{
				Statement: `INSERT INTO num_result SELECT t1.id, t2.id, round(t1.val + t2.val, 10)
    FROM num_data t1, num_data t2;`,
			},
			{
				Statement: `SELECT t1.id1, t1.id2, t1.result, round(t2.expected, 10) as expected
    FROM num_result t1, num_exp_add t2
    WHERE t1.id1 = t2.id1 AND t1.id2 = t2.id2
    AND t1.result != round(t2.expected, 10);`,
				Results: []sql.Row{},
			},
			{
				Statement: `DELETE FROM num_result;`,
			},
			{
				Statement: `INSERT INTO num_result SELECT t1.id, t2.id, t1.val - t2.val
    FROM num_data t1, num_data t2;`,
			},
			{
				Statement: `SELECT t1.id1, t1.id2, t1.result, t2.expected
    FROM num_result t1, num_exp_sub t2
    WHERE t1.id1 = t2.id1 AND t1.id2 = t2.id2
    AND t1.result != t2.expected;`,
				Results: []sql.Row{},
			},
			{
				Statement: `DELETE FROM num_result;`,
			},
			{
				Statement: `INSERT INTO num_result SELECT t1.id, t2.id, round(t1.val - t2.val, 40)
    FROM num_data t1, num_data t2;`,
			},
			{
				Statement: `SELECT t1.id1, t1.id2, t1.result, round(t2.expected, 40)
    FROM num_result t1, num_exp_sub t2
    WHERE t1.id1 = t2.id1 AND t1.id2 = t2.id2
    AND t1.result != round(t2.expected, 40);`,
				Results: []sql.Row{},
			},
			{
				Statement: `DELETE FROM num_result;`,
			},
			{
				Statement: `INSERT INTO num_result SELECT t1.id, t2.id, t1.val * t2.val
    FROM num_data t1, num_data t2;`,
			},
			{
				Statement: `SELECT t1.id1, t1.id2, t1.result, t2.expected
    FROM num_result t1, num_exp_mul t2
    WHERE t1.id1 = t2.id1 AND t1.id2 = t2.id2
    AND t1.result != t2.expected;`,
				Results: []sql.Row{},
			},
			{
				Statement: `DELETE FROM num_result;`,
			},
			{
				Statement: `INSERT INTO num_result SELECT t1.id, t2.id, round(t1.val * t2.val, 30)
    FROM num_data t1, num_data t2;`,
			},
			{
				Statement: `SELECT t1.id1, t1.id2, t1.result, round(t2.expected, 30) as expected
    FROM num_result t1, num_exp_mul t2
    WHERE t1.id1 = t2.id1 AND t1.id2 = t2.id2
    AND t1.result != round(t2.expected, 30);`,
				Results: []sql.Row{},
			},
			{
				Statement: `DELETE FROM num_result;`,
			},
			{
				Statement: `INSERT INTO num_result SELECT t1.id, t2.id, t1.val / t2.val
    FROM num_data t1, num_data t2
    WHERE t2.val != '0.0';`,
			},
			{
				Statement: `SELECT t1.id1, t1.id2, t1.result, t2.expected
    FROM num_result t1, num_exp_div t2
    WHERE t1.id1 = t2.id1 AND t1.id2 = t2.id2
    AND t1.result != t2.expected;`,
				Results: []sql.Row{},
			},
			{
				Statement: `DELETE FROM num_result;`,
			},
			{
				Statement: `INSERT INTO num_result SELECT t1.id, t2.id, round(t1.val / t2.val, 80)
    FROM num_data t1, num_data t2
    WHERE t2.val != '0.0';`,
			},
			{
				Statement: `SELECT t1.id1, t1.id2, t1.result, round(t2.expected, 80) as expected
    FROM num_result t1, num_exp_div t2
    WHERE t1.id1 = t2.id1 AND t1.id2 = t2.id2
    AND t1.result != round(t2.expected, 80);`,
				Results: []sql.Row{},
			},
			{
				Statement: `DELETE FROM num_result;`,
			},
			{
				Statement: `INSERT INTO num_result SELECT id, 0, SQRT(ABS(val))
    FROM num_data;`,
			},
			{
				Statement: `SELECT t1.id1, t1.result, t2.expected
    FROM num_result t1, num_exp_sqrt t2
    WHERE t1.id1 = t2.id
    AND t1.result != t2.expected;`,
				Results: []sql.Row{},
			},
			{
				Statement: `DELETE FROM num_result;`,
			},
			{
				Statement: `INSERT INTO num_result SELECT id, 0, LN(ABS(val))
    FROM num_data
    WHERE val != '0.0';`,
			},
			{
				Statement: `SELECT t1.id1, t1.result, t2.expected
    FROM num_result t1, num_exp_ln t2
    WHERE t1.id1 = t2.id
    AND t1.result != t2.expected;`,
				Results: []sql.Row{},
			},
			{
				Statement: `DELETE FROM num_result;`,
			},
			{
				Statement: `INSERT INTO num_result SELECT id, 0, LOG(numeric '10', ABS(val))
    FROM num_data
    WHERE val != '0.0';`,
			},
			{
				Statement: `SELECT t1.id1, t1.result, t2.expected
    FROM num_result t1, num_exp_log10 t2
    WHERE t1.id1 = t2.id
    AND t1.result != t2.expected;`,
				Results: []sql.Row{},
			},
			{
				Statement: `DELETE FROM num_result;`,
			},
			{
				Statement: `INSERT INTO num_result SELECT id, 0, POWER(numeric '10', LN(ABS(round(val,200))))
    FROM num_data
    WHERE val != '0.0';`,
			},
			{
				Statement: `SELECT t1.id1, t1.result, t2.expected
    FROM num_result t1, num_exp_power_10_ln t2
    WHERE t1.id1 = t2.id
    AND t1.result != t2.expected;`,
				Results: []sql.Row{},
			},
			{
				Statement: `WITH v(x) AS
  (VALUES('0'::numeric),('1'),('-1'),('4.2'),('inf'),('-inf'),('nan'))
SELECT x1, x2,
  x1 + x2 AS sum,
  x1 - x2 AS diff,
  x1 * x2 AS prod
FROM v AS v1(x1), v AS v2(x2);`,
				Results: []sql.Row{{0, 0, 0, 0, 0}, {0, 1, 1, -1, 0}, {0, -1, -1, 1, 0}, {0, 4.2, 4.2, -4.2, 0.0}, {0, `Infinity`, `Infinity`, `-Infinity`, `NaN`}, {0, `-Infinity`, `-Infinity`, `Infinity`, `NaN`}, {0, `NaN`, `NaN`, `NaN`, `NaN`}, {1, 0, 1, 1, 0}, {1, 1, 2, 0, 1}, {1, -1, 0, 2, -1}, {1, 4.2, 5.2, -3.2, 4.2}, {1, `Infinity`, `Infinity`, `-Infinity`, `Infinity`}, {1, `-Infinity`, `-Infinity`, `Infinity`, `-Infinity`}, {1, `NaN`, `NaN`, `NaN`, `NaN`}, {-1, 0, -1, -1, 0}, {-1, 1, 0, -2, -1}, {-1, -1, -2, 0, 1}, {-1, 4.2, 3.2, -5.2, -4.2}, {-1, `Infinity`, `Infinity`, `-Infinity`, `-Infinity`}, {-1, `-Infinity`, `-Infinity`, `Infinity`, `Infinity`}, {-1, `NaN`, `NaN`, `NaN`, `NaN`}, {4.2, 0, 4.2, 4.2, 0.0}, {4.2, 1, 5.2, 3.2, 4.2}, {4.2, -1, 3.2, 5.2, -4.2}, {4.2, 4.2, 8.4, 0.0, 17.64}, {4.2, `Infinity`, `Infinity`, `-Infinity`, `Infinity`}, {4.2, `-Infinity`, `-Infinity`, `Infinity`, `-Infinity`}, {4.2, `NaN`, `NaN`, `NaN`, `NaN`}, {`Infinity`, 0, `Infinity`, `Infinity`, `NaN`}, {`Infinity`, 1, `Infinity`, `Infinity`, `Infinity`}, {`Infinity`, -1, `Infinity`, `Infinity`, `-Infinity`}, {`Infinity`, 4.2, `Infinity`, `Infinity`, `Infinity`}, {`Infinity`, `Infinity`, `Infinity`, `NaN`, `Infinity`}, {`Infinity`, `-Infinity`, `NaN`, `Infinity`, `-Infinity`}, {`Infinity`, `NaN`, `NaN`, `NaN`, `NaN`}, {`-Infinity`, 0, `-Infinity`, `-Infinity`, `NaN`}, {`-Infinity`, 1, `-Infinity`, `-Infinity`, `-Infinity`}, {`-Infinity`, -1, `-Infinity`, `-Infinity`, `Infinity`}, {`-Infinity`, 4.2, `-Infinity`, `-Infinity`, `-Infinity`}, {`-Infinity`, `Infinity`, `NaN`, `-Infinity`, `-Infinity`}, {`-Infinity`, `-Infinity`, `-Infinity`, `NaN`, `Infinity`}, {`-Infinity`, `NaN`, `NaN`, `NaN`, `NaN`}, {`NaN`, 0, `NaN`, `NaN`, `NaN`}, {`NaN`, 1, `NaN`, `NaN`, `NaN`}, {`NaN`, -1, `NaN`, `NaN`, `NaN`}, {`NaN`, 4.2, `NaN`, `NaN`, `NaN`}, {`NaN`, `Infinity`, `NaN`, `NaN`, `NaN`}, {`NaN`, `-Infinity`, `NaN`, `NaN`, `NaN`}, {`NaN`, `NaN`, `NaN`, `NaN`, `NaN`}},
			},
			{
				Statement: `WITH v(x) AS
  (VALUES('0'::numeric),('1'),('-1'),('4.2'),('inf'),('-inf'),('nan'))
SELECT x1, x2,
  x1 / x2 AS quot,
  x1 % x2 AS mod,
  div(x1, x2) AS div
FROM v AS v1(x1), v AS v2(x2) WHERE x2 != 0;`,
				Results: []sql.Row{{0, 1, 0.00000000000000000000, 0, 0}, {1, 1, 1.00000000000000000000, 0, 1}, {-1, 1, -1.00000000000000000000, 0, -1}, {4.2, 1, 4.2000000000000000, 0.2, 4}, {`Infinity`, 1, `Infinity`, `NaN`, `Infinity`}, {`-Infinity`, 1, `-Infinity`, `NaN`, `-Infinity`}, {`NaN`, 1, `NaN`, `NaN`, `NaN`}, {0, -1, 0.00000000000000000000, 0, 0}, {1, -1, -1.00000000000000000000, 0, -1}, {-1, -1, 1.00000000000000000000, 0, 1}, {4.2, -1, -4.2000000000000000, 0.2, -4}, {`Infinity`, -1, `-Infinity`, `NaN`, `-Infinity`}, {`-Infinity`, -1, `Infinity`, `NaN`, `Infinity`}, {`NaN`, -1, `NaN`, `NaN`, `NaN`}, {0, 4.2, 0.00000000000000000000, 0.0, 0}, {1, 4.2, 0.23809523809523809524, 1.0, 0}, {-1, 4.2, -0.23809523809523809524, -1.0, 0}, {4.2, 4.2, 1.00000000000000000000, 0.0, 1}, {`Infinity`, 4.2, `Infinity`, `NaN`, `Infinity`}, {`-Infinity`, 4.2, `-Infinity`, `NaN`, `-Infinity`}, {`NaN`, 4.2, `NaN`, `NaN`, `NaN`}, {0, `Infinity`, 0, 0, 0}, {1, `Infinity`, 0, 1, 0}, {-1, `Infinity`, 0, -1, 0}, {4.2, `Infinity`, 0, 4.2, 0}, {`Infinity`, `Infinity`, `NaN`, `NaN`, `NaN`}, {`-Infinity`, `Infinity`, `NaN`, `NaN`, `NaN`}, {`NaN`, `Infinity`, `NaN`, `NaN`, `NaN`}, {0, `-Infinity`, 0, 0, 0}, {1, `-Infinity`, 0, 1, 0}, {-1, `-Infinity`, 0, -1, 0}, {4.2, `-Infinity`, 0, 4.2, 0}, {`Infinity`, `-Infinity`, `NaN`, `NaN`, `NaN`}, {`-Infinity`, `-Infinity`, `NaN`, `NaN`, `NaN`}, {`NaN`, `-Infinity`, `NaN`, `NaN`, `NaN`}, {0, `NaN`, `NaN`, `NaN`, `NaN`}, {1, `NaN`, `NaN`, `NaN`, `NaN`}, {-1, `NaN`, `NaN`, `NaN`, `NaN`}, {4.2, `NaN`, `NaN`, `NaN`, `NaN`}, {`Infinity`, `NaN`, `NaN`, `NaN`, `NaN`}, {`-Infinity`, `NaN`, `NaN`, `NaN`, `NaN`}, {`NaN`, `NaN`, `NaN`, `NaN`, `NaN`}},
			},
			{
				Statement:   `SELECT 'inf'::numeric / '0';`,
				ErrorString: `division by zero`,
			},
			{
				Statement:   `SELECT '-inf'::numeric / '0';`,
				ErrorString: `division by zero`,
			},
			{
				Statement: `SELECT 'nan'::numeric / '0';`,
				Results:   []sql.Row{{`NaN`}},
			},
			{
				Statement:   `SELECT '0'::numeric / '0';`,
				ErrorString: `division by zero`,
			},
			{
				Statement:   `SELECT 'inf'::numeric % '0';`,
				ErrorString: `division by zero`,
			},
			{
				Statement:   `SELECT '-inf'::numeric % '0';`,
				ErrorString: `division by zero`,
			},
			{
				Statement: `SELECT 'nan'::numeric % '0';`,
				Results:   []sql.Row{{`NaN`}},
			},
			{
				Statement:   `SELECT '0'::numeric % '0';`,
				ErrorString: `division by zero`,
			},
			{
				Statement:   `SELECT div('inf'::numeric, '0');`,
				ErrorString: `division by zero`,
			},
			{
				Statement:   `SELECT div('-inf'::numeric, '0');`,
				ErrorString: `division by zero`,
			},
			{
				Statement: `SELECT div('nan'::numeric, '0');`,
				Results:   []sql.Row{{`NaN`}},
			},
			{
				Statement:   `SELECT div('0'::numeric, '0');`,
				ErrorString: `division by zero`,
			},
			{
				Statement: `WITH v(x) AS
  (VALUES('0'::numeric),('1'),('-1'),('4.2'),('-7.777'),('inf'),('-inf'),('nan'))
SELECT x, -x as minusx, abs(x), floor(x), ceil(x), sign(x), numeric_inc(x) as inc
FROM v;`,
				Results: []sql.Row{{0, 0, 0, 0, 0, 0, 1}, {1, -1, 1, 1, 1, 1, 2}, {-1, 1, 1, -1, -1, -1, 0}, {4.2, -4.2, 4.2, 4, 5, 1, 5.2}, {-7.777, 7.777, 7.777, -8, -7, -1, -6.777}, {`Infinity`, `-Infinity`, `Infinity`, `Infinity`, `Infinity`, 1, `Infinity`}, {`-Infinity`, `Infinity`, `Infinity`, `-Infinity`, `-Infinity`, -1, `-Infinity`}, {`NaN`, `NaN`, `NaN`, `NaN`, `NaN`, `NaN`, `NaN`}},
			},
			{
				Statement: `WITH v(x) AS
  (VALUES('0'::numeric),('1'),('-1'),('4.2'),('-7.777'),('inf'),('-inf'),('nan'))
SELECT x, round(x), round(x,1) as round1, trunc(x), trunc(x,1) as trunc1
FROM v;`,
				Results: []sql.Row{{0, 0, 0.0, 0, 0.0}, {1, 1, 1.0, 1, 1.0}, {-1, -1, -1.0, -1, -1.0}, {4.2, 4, 4.2, 4, 4.2}, {-7.777, -8, -7.8, -7, -7.7}, {`Infinity`, `Infinity`, `Infinity`, `Infinity`, `Infinity`}, {`-Infinity`, `-Infinity`, `-Infinity`, `-Infinity`, `-Infinity`}, {`NaN`, `NaN`, `NaN`, `NaN`, `NaN`}},
			},
			{
				Statement: `WITH v(x) AS
  (VALUES('0'::numeric),('1'),('-1'),('4.2'),('-7.777'),('1e340'),('-1e340'),
         ('inf'),('-inf'),('nan'),
         ('inf'),('-inf'),('nan'))
SELECT substring(x::text, 1, 32)
FROM v ORDER BY x;`,
				Results: []sql.Row{{`-Infinity`}, {`-Infinity`}, {-1000000000000000000000000000000.0}, {-7.777}, {-1}, {0}, {1}, {4.2}, {10000000000000000000000000000000.0}, {`Infinity`}, {`Infinity`}, {`NaN`}, {`NaN`}},
			},
			{
				Statement: `WITH v(x) AS
  (VALUES('0'::numeric),('1'),('4.2'),('inf'),('nan'))
SELECT x, sqrt(x)
FROM v;`,
				Results: []sql.Row{{0, 0.000000000000000}, {1, 1.000000000000000}, {4.2, 2.049390153191920}, {`Infinity`, `Infinity`}, {`NaN`, `NaN`}},
			},
			{
				Statement:   `SELECT sqrt('-1'::numeric);`,
				ErrorString: `cannot take square root of a negative number`,
			},
			{
				Statement:   `SELECT sqrt('-inf'::numeric);`,
				ErrorString: `cannot take square root of a negative number`,
			},
			{
				Statement: `WITH v(x) AS
  (VALUES('1'::numeric),('4.2'),('inf'),('nan'))
SELECT x,
  log(x),
  log10(x),
  ln(x)
FROM v;`,
				Results: []sql.Row{{1, 0.0000000000000000, 0.0000000000000000, 0.0000000000000000}, {4.2, 0.6232492903979005, 0.6232492903979005, 1.4350845252893226}, {`Infinity`, `Infinity`, `Infinity`, `Infinity`}, {`NaN`, `NaN`, `NaN`, `NaN`}},
			},
			{
				Statement:   `SELECT ln('0'::numeric);`,
				ErrorString: `cannot take logarithm of zero`,
			},
			{
				Statement:   `SELECT ln('-1'::numeric);`,
				ErrorString: `cannot take logarithm of a negative number`,
			},
			{
				Statement:   `SELECT ln('-inf'::numeric);`,
				ErrorString: `cannot take logarithm of a negative number`,
			},
			{
				Statement: `WITH v(x) AS
  (VALUES('2'::numeric),('4.2'),('inf'),('nan'))
SELECT x1, x2,
  log(x1, x2)
FROM v AS v1(x1), v AS v2(x2);`,
				Results: []sql.Row{{2, 2, 1.0000000000000000}, {2, 4.2, 2.0703893278913979}, {2, `Infinity`, `Infinity`}, {2, `NaN`, `NaN`}, {4.2, 2, 0.4830009440873890}, {4.2, 4.2, 1.0000000000000000}, {4.2, `Infinity`, `Infinity`}, {4.2, `NaN`, `NaN`}, {`Infinity`, 2, 0}, {`Infinity`, 4.2, 0}, {`Infinity`, `Infinity`, `NaN`}, {`Infinity`, `NaN`, `NaN`}, {`NaN`, 2, `NaN`}, {`NaN`, 4.2, `NaN`}, {`NaN`, `Infinity`, `NaN`}, {`NaN`, `NaN`, `NaN`}},
			},
			{
				Statement:   `SELECT log('0'::numeric, '10');`,
				ErrorString: `cannot take logarithm of zero`,
			},
			{
				Statement:   `SELECT log('10'::numeric, '0');`,
				ErrorString: `cannot take logarithm of zero`,
			},
			{
				Statement:   `SELECT log('-inf'::numeric, '10');`,
				ErrorString: `cannot take logarithm of a negative number`,
			},
			{
				Statement:   `SELECT log('10'::numeric, '-inf');`,
				ErrorString: `cannot take logarithm of a negative number`,
			},
			{
				Statement:   `SELECT log('inf'::numeric, '0');`,
				ErrorString: `cannot take logarithm of zero`,
			},
			{
				Statement:   `SELECT log('inf'::numeric, '-inf');`,
				ErrorString: `cannot take logarithm of a negative number`,
			},
			{
				Statement:   `SELECT log('-inf'::numeric, 'inf');`,
				ErrorString: `cannot take logarithm of a negative number`,
			},
			{
				Statement: `WITH v(x) AS
  (VALUES('0'::numeric),('1'),('2'),('4.2'),('inf'),('nan'))
SELECT x1, x2,
  power(x1, x2)
FROM v AS v1(x1), v AS v2(x2) WHERE x1 != 0 OR x2 >= 0;`,
				Results: []sql.Row{{0, 0, 1.0000000000000000}, {0, 1, 0.0000000000000000}, {0, 2, 0.0000000000000000}, {0, 4.2, 0.0000000000000000}, {0, `Infinity`, 0}, {0, `NaN`, `NaN`}, {1, 0, 1.0000000000000000}, {1, 1, 1.0000000000000000}, {1, 2, 1.0000000000000000}, {1, 4.2, 1.0000000000000000}, {1, `Infinity`, 1}, {1, `NaN`, 1}, {2, 0, 1.0000000000000000}, {2, 1, 2.0000000000000000}, {2, 2, 4.0000000000000000}, {2, 4.2, 18.379173679952560}, {2, `Infinity`, `Infinity`}, {2, `NaN`, `NaN`}, {4.2, 0, 1.0000000000000000}, {4.2, 1, 4.2000000000000000}, {4.2, 2, 17.6400000000000000}, {4.2, 4.2, 414.61691860129675}, {4.2, `Infinity`, `Infinity`}, {4.2, `NaN`, `NaN`}, {`Infinity`, 0, 1}, {`Infinity`, 1, `Infinity`}, {`Infinity`, 2, `Infinity`}, {`Infinity`, 4.2, `Infinity`}, {`Infinity`, `Infinity`, `Infinity`}, {`Infinity`, `NaN`, `NaN`}, {`NaN`, 0, 1}, {`NaN`, 1, `NaN`}, {`NaN`, 2, `NaN`}, {`NaN`, 4.2, `NaN`}, {`NaN`, `Infinity`, `NaN`}, {`NaN`, `NaN`, `NaN`}},
			},
			{
				Statement:   `SELECT power('0'::numeric, '-1');`,
				ErrorString: `zero raised to a negative power is undefined`,
			},
			{
				Statement:   `SELECT power('0'::numeric, '-inf');`,
				ErrorString: `zero raised to a negative power is undefined`,
			},
			{
				Statement: `SELECT power('-1'::numeric, 'inf');`,
				Results:   []sql.Row{{1}},
			},
			{
				Statement: `SELECT power('-2'::numeric, '3');`,
				Results:   []sql.Row{{-8.0000000000000000}},
			},
			{
				Statement:   `SELECT power('-2'::numeric, '3.3');`,
				ErrorString: `a negative number raised to a non-integer power yields a complex result`,
			},
			{
				Statement: `SELECT power('-2'::numeric, '-1');`,
				Results:   []sql.Row{{-0.5000000000000000}},
			},
			{
				Statement:   `SELECT power('-2'::numeric, '-1.5');`,
				ErrorString: `a negative number raised to a non-integer power yields a complex result`,
			},
			{
				Statement: `SELECT power('-2'::numeric, 'inf');`,
				Results:   []sql.Row{{`Infinity`}},
			},
			{
				Statement: `SELECT power('-2'::numeric, '-inf');`,
				Results:   []sql.Row{{0}},
			},
			{
				Statement: `SELECT power('inf'::numeric, '-2');`,
				Results:   []sql.Row{{0}},
			},
			{
				Statement: `SELECT power('inf'::numeric, '-inf');`,
				Results:   []sql.Row{{0}},
			},
			{
				Statement: `SELECT power('-inf'::numeric, '2');`,
				Results:   []sql.Row{{`Infinity`}},
			},
			{
				Statement: `SELECT power('-inf'::numeric, '3');`,
				Results:   []sql.Row{{`-Infinity`}},
			},
			{
				Statement:   `SELECT power('-inf'::numeric, '4.5');`,
				ErrorString: `a negative number raised to a non-integer power yields a complex result`,
			},
			{
				Statement: `SELECT power('-inf'::numeric, '-2');`,
				Results:   []sql.Row{{0}},
			},
			{
				Statement: `SELECT power('-inf'::numeric, '-3');`,
				Results:   []sql.Row{{0}},
			},
			{
				Statement: `SELECT power('-inf'::numeric, '0');`,
				Results:   []sql.Row{{1}},
			},
			{
				Statement: `SELECT power('-inf'::numeric, 'inf');`,
				Results:   []sql.Row{{`Infinity`}},
			},
			{
				Statement: `SELECT power('-inf'::numeric, '-inf');`,
				Results:   []sql.Row{{0}},
			},
			{
				Statement: `SELECT AVG(val) FROM num_data;`,
				Results:   []sql.Row{{-13430913.592242320700}},
			},
			{
				Statement: `SELECT MAX(val) FROM num_data;`,
				Results:   []sql.Row{{7799461.4119000000}},
			},
			{
				Statement: `SELECT MIN(val) FROM num_data;`,
				Results:   []sql.Row{{-83028485.0000000000}},
			},
			{
				Statement: `SELECT STDDEV(val) FROM num_data;`,
				Results:   []sql.Row{{27791203.28758835329805617386}},
			},
			{
				Statement: `SELECT VARIANCE(val) FROM num_data;`,
				Results:   []sql.Row{{772350980172061.69659105821915863601}},
			},
			{
				Statement: `CREATE TABLE fract_only (id int, val numeric(4,4));`,
			},
			{
				Statement: `INSERT INTO fract_only VALUES (1, '0.0');`,
			},
			{
				Statement: `INSERT INTO fract_only VALUES (2, '0.1');`,
			},
			{
				Statement:   `INSERT INTO fract_only VALUES (3, '1.0');	-- should fail`,
				ErrorString: `numeric field overflow`,
			},
			{
				Statement: `INSERT INTO fract_only VALUES (4, '-0.9999');`,
			},
			{
				Statement: `INSERT INTO fract_only VALUES (5, '0.99994');`,
			},
			{
				Statement:   `INSERT INTO fract_only VALUES (6, '0.99995');  -- should fail`,
				ErrorString: `numeric field overflow`,
			},
			{
				Statement: `INSERT INTO fract_only VALUES (7, '0.00001');`,
			},
			{
				Statement: `INSERT INTO fract_only VALUES (8, '0.00017');`,
			},
			{
				Statement: `INSERT INTO fract_only VALUES (9, 'NaN');`,
			},
			{
				Statement:   `INSERT INTO fract_only VALUES (10, 'Inf');	-- should fail`,
				ErrorString: `numeric field overflow`,
			},
			{
				Statement:   `INSERT INTO fract_only VALUES (11, '-Inf');	-- should fail`,
				ErrorString: `numeric field overflow`,
			},
			{
				Statement: `SELECT * FROM fract_only;`,
				Results:   []sql.Row{{1, 0.0000}, {2, 0.1000}, {4, -0.9999}, {5, 0.9999}, {7, 0.0000}, {8, 0.0002}, {9, `NaN`}},
			},
			{
				Statement: `DROP TABLE fract_only;`,
			},
			{
				Statement:   `SELECT (-9223372036854775808.5)::int8; -- should fail`,
				ErrorString: `bigint out of range`,
			},
			{
				Statement: `SELECT (-9223372036854775808.4)::int8; -- ok`,
				Results:   []sql.Row{{-9223372036854775808}},
			},
			{
				Statement: `SELECT 9223372036854775807.4::int8; -- ok`,
				Results:   []sql.Row{{9223372036854775807}},
			},
			{
				Statement:   `SELECT 9223372036854775807.5::int8; -- should fail`,
				ErrorString: `bigint out of range`,
			},
			{
				Statement:   `SELECT (-2147483648.5)::int4; -- should fail`,
				ErrorString: `integer out of range`,
			},
			{
				Statement: `SELECT (-2147483648.4)::int4; -- ok`,
				Results:   []sql.Row{{-2147483648}},
			},
			{
				Statement: `SELECT 2147483647.4::int4; -- ok`,
				Results:   []sql.Row{{2147483647}},
			},
			{
				Statement:   `SELECT 2147483647.5::int4; -- should fail`,
				ErrorString: `integer out of range`,
			},
			{
				Statement:   `SELECT (-32768.5)::int2; -- should fail`,
				ErrorString: `smallint out of range`,
			},
			{
				Statement: `SELECT (-32768.4)::int2; -- ok`,
				Results:   []sql.Row{{-32768}},
			},
			{
				Statement: `SELECT 32767.4::int2; -- ok`,
				Results:   []sql.Row{{32767}},
			},
			{
				Statement:   `SELECT 32767.5::int2; -- should fail`,
				ErrorString: `smallint out of range`,
			},
			{
				Statement: `SELECT 'NaN'::float8::numeric;`,
				Results:   []sql.Row{{`NaN`}},
			},
			{
				Statement: `SELECT 'Infinity'::float8::numeric;`,
				Results:   []sql.Row{{`Infinity`}},
			},
			{
				Statement: `SELECT '-Infinity'::float8::numeric;`,
				Results:   []sql.Row{{`-Infinity`}},
			},
			{
				Statement: `SELECT 'NaN'::numeric::float8;`,
				Results:   []sql.Row{{`NaN`}},
			},
			{
				Statement: `SELECT 'Infinity'::numeric::float8;`,
				Results:   []sql.Row{{`Infinity`}},
			},
			{
				Statement: `SELECT '-Infinity'::numeric::float8;`,
				Results:   []sql.Row{{`-Infinity`}},
			},
			{
				Statement: `SELECT 'NaN'::float4::numeric;`,
				Results:   []sql.Row{{`NaN`}},
			},
			{
				Statement: `SELECT 'Infinity'::float4::numeric;`,
				Results:   []sql.Row{{`Infinity`}},
			},
			{
				Statement: `SELECT '-Infinity'::float4::numeric;`,
				Results:   []sql.Row{{`-Infinity`}},
			},
			{
				Statement: `SELECT 'NaN'::numeric::float4;`,
				Results:   []sql.Row{{`NaN`}},
			},
			{
				Statement: `SELECT 'Infinity'::numeric::float4;`,
				Results:   []sql.Row{{`Infinity`}},
			},
			{
				Statement: `SELECT '-Infinity'::numeric::float4;`,
				Results:   []sql.Row{{`-Infinity`}},
			},
			{
				Statement: `SELECT '42'::int2::numeric;`,
				Results:   []sql.Row{{42}},
			},
			{
				Statement:   `SELECT 'NaN'::numeric::int2;`,
				ErrorString: `cannot convert NaN to smallint`,
			},
			{
				Statement:   `SELECT 'Infinity'::numeric::int2;`,
				ErrorString: `cannot convert infinity to smallint`,
			},
			{
				Statement:   `SELECT '-Infinity'::numeric::int2;`,
				ErrorString: `cannot convert infinity to smallint`,
			},
			{
				Statement:   `SELECT 'NaN'::numeric::int4;`,
				ErrorString: `cannot convert NaN to integer`,
			},
			{
				Statement:   `SELECT 'Infinity'::numeric::int4;`,
				ErrorString: `cannot convert infinity to integer`,
			},
			{
				Statement:   `SELECT '-Infinity'::numeric::int4;`,
				ErrorString: `cannot convert infinity to integer`,
			},
			{
				Statement:   `SELECT 'NaN'::numeric::int8;`,
				ErrorString: `cannot convert NaN to bigint`,
			},
			{
				Statement:   `SELECT 'Infinity'::numeric::int8;`,
				ErrorString: `cannot convert infinity to bigint`,
			},
			{
				Statement:   `SELECT '-Infinity'::numeric::int8;`,
				ErrorString: `cannot convert infinity to bigint`,
			},
			{
				Statement: `CREATE TABLE ceil_floor_round (a numeric);`,
			},
			{
				Statement: `INSERT INTO ceil_floor_round VALUES ('-5.5');`,
			},
			{
				Statement: `INSERT INTO ceil_floor_round VALUES ('-5.499999');`,
			},
			{
				Statement: `INSERT INTO ceil_floor_round VALUES ('9.5');`,
			},
			{
				Statement: `INSERT INTO ceil_floor_round VALUES ('9.4999999');`,
			},
			{
				Statement: `INSERT INTO ceil_floor_round VALUES ('0.0');`,
			},
			{
				Statement: `INSERT INTO ceil_floor_round VALUES ('0.0000001');`,
			},
			{
				Statement: `INSERT INTO ceil_floor_round VALUES ('-0.000001');`,
			},
			{
				Statement: `SELECT a, ceil(a), ceiling(a), floor(a), round(a) FROM ceil_floor_round;`,
				Results:   []sql.Row{{-5.5, -5, -5, -6, -6}, {-5.499999, -5, -5, -6, -5}, {9.5, 10, 10, 9, 10}, {9.4999999, 10, 10, 9, 9}, {0.0, 0, 0, 0, 0}, {0.0000001, 1, 1, 0, 0}, {-0.000001, 0, 0, -1, 0}},
			},
			{
				Statement: `DROP TABLE ceil_floor_round;`,
			},
			{
				Statement: `SELECT i as pow,
	round((-2.5 * 10 ^ i)::numeric, -i),
	round((-1.5 * 10 ^ i)::numeric, -i),
	round((-0.5 * 10 ^ i)::numeric, -i),
	round((0.5 * 10 ^ i)::numeric, -i),
	round((1.5 * 10 ^ i)::numeric, -i),
	round((2.5 * 10 ^ i)::numeric, -i)
FROM generate_series(-5,5) AS t(i);`,
				Results: []sql.Row{{-5, -0.00003, -0.00002, -0.00001, 0.00001, 0.00002, 0.00003}, {-4, -0.0003, -0.0002, -0.0001, 0.0001, 0.0002, 0.0003}, {-3, -0.003, -0.002, -0.001, 0.001, 0.002, 0.003}, {-2, -0.03, -0.02, -0.01, 0.01, 0.02, 0.03}, {-1, -0.3, -0.2, -0.1, 0.1, 0.2, 0.3}, {0, -3, -2, -1, 1, 2, 3}, {1, -30, -20, -10, 10, 20, 30}, {2, -300, -200, -100, 100, 200, 300}, {3, -3000, -2000, -1000, 1000, 2000, 3000}, {4, -30000, -20000, -10000, 10000, 20000, 30000}, {5, -300000, -200000, -100000, 100000, 200000, 300000}},
			},
			{
				Statement:   `SELECT width_bucket(5.0, 3.0, 4.0, 0);`,
				ErrorString: `count must be greater than zero`,
			},
			{
				Statement:   `SELECT width_bucket(5.0, 3.0, 4.0, -5);`,
				ErrorString: `count must be greater than zero`,
			},
			{
				Statement:   `SELECT width_bucket(3.5, 3.0, 3.0, 888);`,
				ErrorString: `lower bound cannot equal upper bound`,
			},
			{
				Statement:   `SELECT width_bucket(5.0::float8, 3.0::float8, 4.0::float8, 0);`,
				ErrorString: `count must be greater than zero`,
			},
			{
				Statement:   `SELECT width_bucket(5.0::float8, 3.0::float8, 4.0::float8, -5);`,
				ErrorString: `count must be greater than zero`,
			},
			{
				Statement:   `SELECT width_bucket(3.5::float8, 3.0::float8, 3.0::float8, 888);`,
				ErrorString: `lower bound cannot equal upper bound`,
			},
			{
				Statement:   `SELECT width_bucket('NaN', 3.0, 4.0, 888);`,
				ErrorString: `operand, lower bound, and upper bound cannot be NaN`,
			},
			{
				Statement:   `SELECT width_bucket(0::float8, 'NaN', 4.0::float8, 888);`,
				ErrorString: `operand, lower bound, and upper bound cannot be NaN`,
			},
			{
				Statement:   `SELECT width_bucket(2.0, 3.0, '-inf', 888);`,
				ErrorString: `lower and upper bounds must be finite`,
			},
			{
				Statement:   `SELECT width_bucket(0::float8, '-inf', 4.0::float8, 888);`,
				ErrorString: `lower and upper bounds must be finite`,
			},
			{
				Statement: `CREATE TABLE width_bucket_test (operand_num numeric, operand_f8 float8);`,
			},
			{
				Statement: `COPY width_bucket_test (operand_num) FROM stdin;`,
			},
			{
				Statement: `UPDATE width_bucket_test SET operand_f8 = operand_num::float8;`,
			},
			{
				Statement: `SELECT
    operand_num,
    width_bucket(operand_num, 0, 10, 5) AS wb_1,
    width_bucket(operand_f8, 0, 10, 5) AS wb_1f,
    width_bucket(operand_num, 10, 0, 5) AS wb_2,
    width_bucket(operand_f8, 10, 0, 5) AS wb_2f,
    width_bucket(operand_num, 2, 8, 4) AS wb_3,
    width_bucket(operand_f8, 2, 8, 4) AS wb_3f,
    width_bucket(operand_num, 5.0, 5.5, 20) AS wb_4,
    width_bucket(operand_f8, 5.0, 5.5, 20) AS wb_4f,
    width_bucket(operand_num, -25, 25, 10) AS wb_5,
    width_bucket(operand_f8, -25, 25, 10) AS wb_5f
    FROM width_bucket_test;`,
				Results: []sql.Row{{-5.2, 0, 0, 6, 6, 0, 0, 0, 0, 4, 4}, {-0.0000000001, 0, 0, 6, 6, 0, 0, 0, 0, 5, 5}, {0.000000000001, 1, 1, 5, 5, 0, 0, 0, 0, 6, 6}, {1, 1, 1, 5, 5, 0, 0, 0, 0, 6, 6}, {1.99999999999999, 1, 1, 5, 5, 0, 0, 0, 0, 6, 6}, {2, 2, 2, 5, 5, 1, 1, 0, 0, 6, 6}, {2.00000000000001, 2, 2, 4, 4, 1, 1, 0, 0, 6, 6}, {3, 2, 2, 4, 4, 1, 1, 0, 0, 6, 6}, {4, 3, 3, 4, 4, 2, 2, 0, 0, 6, 6}, {4.5, 3, 3, 3, 3, 2, 2, 0, 0, 6, 6}, {5, 3, 3, 3, 3, 3, 3, 1, 1, 7, 7}, {5.5, 3, 3, 3, 3, 3, 3, 21, 21, 7, 7}, {6, 4, 4, 3, 3, 3, 3, 21, 21, 7, 7}, {7, 4, 4, 2, 2, 4, 4, 21, 21, 7, 7}, {8, 5, 5, 2, 2, 5, 5, 21, 21, 7, 7}, {9, 5, 5, 1, 1, 5, 5, 21, 21, 7, 7}, {9.99999999999999, 5, 5, 1, 1, 5, 5, 21, 21, 7, 7}, {10, 6, 6, 1, 1, 5, 5, 21, 21, 8, 8}, {10.0000000000001, 6, 6, 0, 0, 5, 5, 21, 21, 8, 8}},
			},
			{
				Statement:   `SELECT width_bucket(0.0::numeric, 'Infinity'::numeric, 5, 10); -- error`,
				ErrorString: `lower and upper bounds must be finite`,
			},
			{
				Statement:   `SELECT width_bucket(0.0::numeric, 5, '-Infinity'::numeric, 20); -- error`,
				ErrorString: `lower and upper bounds must be finite`,
			},
			{
				Statement: `SELECT width_bucket('Infinity'::numeric, 1, 10, 10),
       width_bucket('-Infinity'::numeric, 1, 10, 10);`,
				Results: []sql.Row{{11, 0}},
			},
			{
				Statement:   `SELECT width_bucket(0.0::float8, 'Infinity'::float8, 5, 10); -- error`,
				ErrorString: `lower and upper bounds must be finite`,
			},
			{
				Statement:   `SELECT width_bucket(0.0::float8, 5, '-Infinity'::float8, 20); -- error`,
				ErrorString: `lower and upper bounds must be finite`,
			},
			{
				Statement: `SELECT width_bucket('Infinity'::float8, 1, 10, 10),
       width_bucket('-Infinity'::float8, 1, 10, 10);`,
				Results: []sql.Row{{11, 0}},
			},
			{
				Statement: `DROP TABLE width_bucket_test;`,
			},
			{
				Statement: `SELECT x, width_bucket(x::float8, 10, 100, 9) as flt,
       width_bucket(x::numeric, 10, 100, 9) as num
FROM generate_series(0, 110, 10) x;`,
				Results: []sql.Row{{0, 0, 0}, {10, 1, 1}, {20, 2, 2}, {30, 3, 3}, {40, 4, 4}, {50, 5, 5}, {60, 6, 6}, {70, 7, 7}, {80, 8, 8}, {90, 9, 9}, {100, 10, 10}, {110, 10, 10}},
			},
			{
				Statement: `SELECT x, width_bucket(x::float8, 100, 10, 9) as flt,
       width_bucket(x::numeric, 100, 10, 9) as num
FROM generate_series(0, 110, 10) x;`,
				Results: []sql.Row{{0, 10, 10}, {10, 10, 10}, {20, 9, 9}, {30, 8, 8}, {40, 7, 7}, {50, 6, 6}, {60, 5, 5}, {70, 4, 4}, {80, 3, 3}, {90, 2, 2}, {100, 1, 1}, {110, 0, 0}},
			},
			{
				Statement: `SELECT to_char(val, '9G999G999G999G999G999')
	FROM num_data;`,
				Results: []sql.Row{{0}, {0}, {`-34,338,492`}, {4}, {`7,799,461`}, {`16,397`}, {`93,902`}, {`-83,028,485`}, {`74,881`}, {`-24,926,804`}},
			},
			{
				Statement: `SELECT to_char(val, '9G999G999G999G999G999D999G999G999G999G999')
	FROM num_data;`,
				Results: []sql.Row{{`.000,000,000,000,000`}, {`.000,000,000,000,000`}, {`-34,338,492.215,397,047,000,000`}, {`4.310,000,000,000,000`}, {`7,799,461.411,900,000,000,000`}, {`16,397.038,491,000,000,000`}, {`93,901.577,630,260,000,000`}, {`-83,028,485.000,000,000,000,000`}, {`74,881.000,000,000,000,000`}, {`-24,926,804.045,047,420,000,000`}},
			},
			{
				Statement: `SELECT to_char(val, '9999999999999999.999999999999999PR')
	FROM num_data;`,
				Results: []sql.Row{{.000000000000000}, {.000000000000000}, {`<34338492.215397047000000>`}, {4.310000000000000}, {7799461.411900000000000}, {16397.038491000000000}, {93901.577630260000000}, {`<83028485.000000000000000>`}, {74881.000000000000000}, {`<24926804.045047420000000>`}},
			},
			{
				Statement: `SELECT to_char(val, '9999999999999999.999999999999999S')
	FROM num_data;`,
				Results: []sql.Row{{`.000000000000000+`}, {`.000000000000000+`}, {`34338492.215397047000000-`}, {`4.310000000000000+`}, {`7799461.411900000000000+`}, {`16397.038491000000000+`}, {`93901.577630260000000+`}, {`83028485.000000000000000-`}, {`74881.000000000000000+`}, {`24926804.045047420000000-`}},
			},
			{
				Statement: `SELECT to_char(val, 'MI9999999999999999.999999999999999')     FROM num_data;`,
				Results:   []sql.Row{{.000000000000000}, {.000000000000000}, {`-        34338492.215397047000000`}, {4.310000000000000}, {7799461.411900000000000}, {16397.038491000000000}, {93901.577630260000000}, {`-        83028485.000000000000000`}, {74881.000000000000000}, {`-        24926804.045047420000000`}},
			},
			{
				Statement: `SELECT to_char(val, 'FMS9999999999999999.999999999999999')    FROM num_data;`,
				Results:   []sql.Row{{+0.}, {+0.}, {-34338492.215397047}, {+4.31}, {+7799461.4119}, {+16397.038491}, {+93901.57763026}, {-83028485.}, {+74881.}, {-24926804.04504742}},
			},
			{
				Statement: `SELECT to_char(val, 'FM9999999999999999.999999999999999THPR') FROM num_data;`,
				Results:   []sql.Row{{0.}, {0.}, {`<34338492.215397047>`}, {4.31}, {7799461.4119}, {16397.038491}, {93901.57763026}, {`<83028485.>`}, {74881.}, {`<24926804.04504742>`}},
			},
			{
				Statement: `SELECT to_char(val, 'SG9999999999999999.999999999999999th')   FROM num_data;`,
				Results:   []sql.Row{{`+                .000000000000000`}, {`+                .000000000000000`}, {`-        34338492.215397047000000`}, {`+               4.310000000000000`}, {`+         7799461.411900000000000`}, {`+           16397.038491000000000`}, {`+           93901.577630260000000`}, {`-        83028485.000000000000000`}, {`+           74881.000000000000000`}, {`-        24926804.045047420000000`}},
			},
			{
				Statement: `SELECT to_char(val, '0999999999999999.999999999999999')       FROM num_data;`,
				Results:   []sql.Row{{0000000000000000.000000000000000}, {0000000000000000.000000000000000}, {-0000000034338492.215397047000000}, {0000000000000004.310000000000000}, {0000000007799461.411900000000000}, {0000000000016397.038491000000000}, {0000000000093901.577630260000000}, {-0000000083028485.000000000000000}, {0000000000074881.000000000000000}, {-0000000024926804.045047420000000}},
			},
			{
				Statement: `SELECT to_char(val, 'S0999999999999999.999999999999999')      FROM num_data;`,
				Results:   []sql.Row{{+0000000000000000.000000000000000}, {+0000000000000000.000000000000000}, {-0000000034338492.215397047000000}, {+0000000000000004.310000000000000}, {+0000000007799461.411900000000000}, {+0000000000016397.038491000000000}, {+0000000000093901.577630260000000}, {-0000000083028485.000000000000000}, {+0000000000074881.000000000000000}, {-0000000024926804.045047420000000}},
			},
			{
				Statement: `SELECT to_char(val, 'FM0999999999999999.999999999999999')     FROM num_data;`,
				Results:   []sql.Row{{0000000000000000.}, {0000000000000000.}, {-0000000034338492.215397047}, {0000000000000004.31}, {0000000007799461.4119}, {0000000000016397.038491}, {0000000000093901.57763026}, {-0000000083028485.}, {0000000000074881.}, {-0000000024926804.04504742}},
			},
			{
				Statement: `SELECT to_char(val, 'FM9999999999999999.099999999999999') 	FROM num_data;`,
				Results:   []sql.Row{{.0}, {.0}, {-34338492.215397047}, {4.31}, {7799461.4119}, {16397.038491}, {93901.57763026}, {-83028485.0}, {74881.0}, {-24926804.04504742}},
			},
			{
				Statement: `SELECT to_char(val, 'FM9999999999990999.990999999999999') 	FROM num_data;`,
				Results:   []sql.Row{{0000.000}, {0000.000}, {-34338492.215397047}, {0004.310}, {7799461.4119}, {16397.038491}, {93901.57763026}, {-83028485.000}, {74881.000}, {-24926804.04504742}},
			},
			{
				Statement: `SELECT to_char(val, 'FM0999999999999999.999909999999999') 	FROM num_data;`,
				Results:   []sql.Row{{0000000000000000.00000}, {0000000000000000.00000}, {-0000000034338492.215397047}, {0000000000000004.31000}, {0000000007799461.41190}, {0000000000016397.038491}, {0000000000093901.57763026}, {-0000000083028485.00000}, {0000000000074881.00000}, {-0000000024926804.04504742}},
			},
			{
				Statement: `SELECT to_char(val, 'FM9999999990999999.099999999999999') 	FROM num_data;`,
				Results:   []sql.Row{{0000000.0}, {0000000.0}, {-34338492.215397047}, {0000004.31}, {7799461.4119}, {0016397.038491}, {0093901.57763026}, {-83028485.0}, {0074881.0}, {-24926804.04504742}},
			},
			{
				Statement: `SELECT to_char(val, 'L9999999999999999.099999999999999')	FROM num_data;`,
				Results:   []sql.Row{{.000000000000000}, {.000000000000000}, {-34338492.215397047000000}, {4.310000000000000}, {7799461.411900000000000}, {16397.038491000000000}, {93901.577630260000000}, {-83028485.000000000000000}, {74881.000000000000000}, {-24926804.045047420000000}},
			},
			{
				Statement: `SELECT to_char(val, 'FM9999999999999999.99999999999999')	FROM num_data;`,
				Results:   []sql.Row{{0.}, {0.}, {-34338492.215397047}, {4.31}, {7799461.4119}, {16397.038491}, {93901.57763026}, {-83028485.}, {74881.}, {-24926804.04504742}},
			},
			{
				Statement: `SELECT to_char(val, 'S 9 9 9 9 9 9 9 9 9 9 9 9 9 9 9 9 . 9 9 9 9 9 9 9 9 9 9 9 9 9 9 9 9 9') FROM num_data;`,
				Results:   []sql.Row{{`+. 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0`}, {`+. 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0`}, {`-3 4 3 3 8 4 9 2 . 2 1 5 3 9 7 0 4 7 0 0 0 0 0 0 0 0`}, {`+4 . 3 1 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0`}, {`+7 7 9 9 4 6 1 . 4 1 1 9 0 0 0 0 0 0 0 0 0 0 0 0 0`}, {`+1 6 3 9 7 . 0 3 8 4 9 1 0 0 0 0 0 0 0 0 0 0 0`}, {`+9 3 9 0 1 . 5 7 7 6 3 0 2 6 0 0 0 0 0 0 0 0 0`}, {`-8 3 0 2 8 4 8 5 . 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0`}, {`+7 4 8 8 1 . 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0`}, {`-2 4 9 2 6 8 0 4 . 0 4 5 0 4 7 4 2 0 0 0 0 0 0 0 0 0`}},
			},
			{
				Statement: `SELECT to_char(val, 'FMS 9 9 9 9 9 9 9 9 9 9 9 9 9 9 9 9 . 9 9 9 9 9 9 9 9 9 9 9 9 9 9 9 9 9') FROM num_data;`,
				Results:   []sql.Row{{`+0 .`}, {`+0 .`}, {`-3 4 3 3 8 4 9 2 . 2 1 5 3 9 7 0 4 7`}, {`+4 . 3 1`}, {`+7 7 9 9 4 6 1 . 4 1 1 9`}, {`+1 6 3 9 7 . 0 3 8 4 9 1`}, {`+9 3 9 0 1 . 5 7 7 6 3 0 2 6`}, {`-8 3 0 2 8 4 8 5 .`}, {`+7 4 8 8 1 .`}, {`-2 4 9 2 6 8 0 4 . 0 4 5 0 4 7 4 2`}},
			},
			{
				Statement: `SELECT to_char(val, E'99999 "text" 9999 "9999" 999 "\\"text between quote marks\\"" 9999') FROM num_data;`,
				Results:   []sql.Row{{`text      9999     "text between quote marks"     0`}, {`text      9999     "text between quote marks"     0`}, {`text    -3 9999 433 "text between quote marks" 8492`}, {`text      9999     "text between quote marks"     4`}, {`text      9999  779 "text between quote marks" 9461`}, {`text      9999    1 "text between quote marks" 6397`}, {`text      9999    9 "text between quote marks" 3902`}, {`text    -8 9999 302 "text between quote marks" 8485`}, {`text      9999    7 "text between quote marks" 4881`}, {`text    -2 9999 492 "text between quote marks" 6804`}},
			},
			{
				Statement: `SELECT to_char(val, '999999SG9999999999')			FROM num_data;`,
				Results:   []sql.Row{{`+         0`}, {`+         0`}, {`-  34338492`}, {`+         4`}, {`+   7799461`}, {`+     16397`}, {`+     93902`}, {`-  83028485`}, {`+     74881`}, {`-  24926804`}},
			},
			{
				Statement: `SELECT to_char(val, 'FM9999999999999999.999999999999999')	FROM num_data;`,
				Results:   []sql.Row{{0.}, {0.}, {-34338492.215397047}, {4.31}, {7799461.4119}, {16397.038491}, {93901.57763026}, {-83028485.}, {74881.}, {-24926804.04504742}},
			},
			{
				Statement: `SELECT to_char(val, '9.999EEEE')				FROM num_data;`,
				Results:   []sql.Row{{0.000e+00}, {0.000e+00}, {-3.434e+07}, {4.310e+00}, {7.799e+06}, {1.640e+04}, {9.390e+04}, {-8.303e+07}, {7.488e+04}, {-2.493e+07}},
			},
			{
				Statement: `WITH v(val) AS
  (VALUES('0'::numeric),('-4.2'),('4.2e9'),('1.2e-5'),('inf'),('-inf'),('nan'))
SELECT val,
  to_char(val, '9.999EEEE') as numeric,
  to_char(val::float8, '9.999EEEE') as float8,
  to_char(val::float4, '9.999EEEE') as float4
FROM v;`,
				Results: []sql.Row{{0, 0.000e+00, 0.000e+00, 0.000e+00}, {-4.2, -4.200e+00, -4.200e+00, -4.200e+00}, {4200000000, 4.200e+09, 4.200e+09, 4.200e+09}, {0.000012, 1.200e-05, 1.200e-05, 1.200e-05}, {`Infinity`, `#.#######`, `#.#######`, `#.#######`}, {`-Infinity`, `#.#######`, `#.#######`, `#.#######`}, {`NaN`, `#.#######`, `#.#######`, `#.#######`}},
			},
			{
				Statement: `WITH v(exp) AS
  (VALUES(-16379),(-16378),(-1234),(-789),(-45),(-5),(-4),(-3),(-2),(-1),(0),
         (1),(2),(3),(4),(5),(38),(275),(2345),(45678),(131070),(131071))
SELECT exp,
  to_char(('1.2345e'||exp)::numeric, '9.999EEEE') as numeric
FROM v;`,
				Results: []sql.Row{{-16379, 1.235e-16379}, {-16378, 1.235e-16378}, {-1234, 1.235e-1234}, {-789, 1.235e-789}, {-45, 1.235e-45}, {-5, 1.235e-05}, {-4, 1.235e-04}, {-3, 1.235e-03}, {-2, 1.235e-02}, {-1, 1.235e-01}, {0, 1.235e+00}, {1, 1.235e+01}, {2, 1.235e+02}, {3, 1.235e+03}, {4, 1.235e+04}, {5, 1.235e+05}, {38, 1.235e+38}, {275, 1.235e+275}, {2345, "1.235e+2345"}, {45678, "1.235e+45678"}, {131070, "1.235e+131070"}, {131071, "1.235e+131071"}},
			},
			{
				Statement: `WITH v(val) AS
  (VALUES('0'::numeric),('-4.2'),('4.2e9'),('1.2e-5'),('inf'),('-inf'),('nan'))
SELECT val,
  to_char(val, 'MI9999999999.99') as numeric,
  to_char(val::float8, 'MI9999999999.99') as float8,
  to_char(val::float4, 'MI9999999999.99') as float4
FROM v;`,
				Results: []sql.Row{{0, .00, .00, .00}, {-4.2, `-         4.20`, `-         4.20`, `-         4.20`}, {4200000000, 4200000000.00, 4200000000.00, 4200000000}, {0.000012, .00, .00, .00}, {`Infinity`, `Infinity`, `Infinity`, `Infinity`}, {`-Infinity`, `-  Infinity`, `-  Infinity`, `-  Infinity`}, {`NaN`, `NaN`, `NaN`, `NaN`}},
			},
			{
				Statement: `WITH v(val) AS
  (VALUES('0'::numeric),('-4.2'),('4.2e9'),('1.2e-5'),('inf'),('-inf'),('nan'))
SELECT val,
  to_char(val, 'MI99.99') as numeric,
  to_char(val::float8, 'MI99.99') as float8,
  to_char(val::float4, 'MI99.99') as float4
FROM v;`,
				Results: []sql.Row{{0, .00, .00, .00}, {-4.2, `- 4.20`, `- 4.20`, `- 4.20`}, {4200000000, `##.##`, `##.##`, `##.`}, {0.000012, .00, .00, .00}, {`Infinity`, `##.##`, `##.##`, `##.`}, {`-Infinity`, `-##.##`, `-##.##`, `-##.`}, {`NaN`, `##.##`, `##.##`, `##.##`}},
			},
			{
				Statement: `SELECT to_char('100'::numeric, 'FM999.9');`,
				Results:   []sql.Row{{100.}},
			},
			{
				Statement: `SELECT to_char('100'::numeric, 'FM999.');`,
				Results:   []sql.Row{{100}},
			},
			{
				Statement: `SELECT to_char('100'::numeric, 'FM999');`,
				Results:   []sql.Row{{100}},
			},
			{
				Statement: `SELECT to_char('12345678901'::float8, 'FM9999999999D9999900000000000000000');`,
				Results:   []sql.Row{{`##########.####`}},
			},
			{
				Statement: `SELECT to_char('100'::numeric, 'foo999');`,
				Results:   []sql.Row{{`foo 100`}},
			},
			{
				Statement: `SELECT to_char('100'::numeric, 'f\oo999');`,
				Results:   []sql.Row{{`f\oo 100`}},
			},
			{
				Statement: `SELECT to_char('100'::numeric, 'f\\oo999');`,
				Results:   []sql.Row{{`f\\oo 100`}},
			},
			{
				Statement: `SELECT to_char('100'::numeric, 'f\"oo999');`,
				Results:   []sql.Row{{`f"oo 100`}},
			},
			{
				Statement: `SELECT to_char('100'::numeric, 'f\\"oo999');`,
				Results:   []sql.Row{{`f\"oo 100`}},
			},
			{
				Statement: `SELECT to_char('100'::numeric, 'f"ool"999');`,
				Results:   []sql.Row{{`fool 100`}},
			},
			{
				Statement: `SELECT to_char('100'::numeric, 'f"\ool"999');`,
				Results:   []sql.Row{{`fool 100`}},
			},
			{
				Statement: `SELECT to_char('100'::numeric, 'f"\\ool"999');`,
				Results:   []sql.Row{{`f\ool 100`}},
			},
			{
				Statement: `SELECT to_char('100'::numeric, 'f"ool\"999');`,
				Results:   []sql.Row{{`fool"999`}},
			},
			{
				Statement: `SELECT to_char('100'::numeric, 'f"ool\\"999');`,
				Results:   []sql.Row{{`fool\ 100`}},
			},
			{
				Statement: `SET lc_numeric = 'C';`,
			},
			{
				Statement: `SELECT to_number('-34,338,492', '99G999G999');`,
				Results:   []sql.Row{{-34338492}},
			},
			{
				Statement: `SELECT to_number('-34,338,492.654,878', '99G999G999D999G999');`,
				Results:   []sql.Row{{-34338492.654878}},
			},
			{
				Statement: `SELECT to_number('<564646.654564>', '999999.999999PR');`,
				Results:   []sql.Row{{-564646.654564}},
			},
			{
				Statement: `SELECT to_number('0.00001-', '9.999999S');`,
				Results:   []sql.Row{{-0.00001}},
			},
			{
				Statement: `SELECT to_number('5.01-', 'FM9.999999S');`,
				Results:   []sql.Row{{-5.01}},
			},
			{
				Statement: `SELECT to_number('5.01-', 'FM9.999999MI');`,
				Results:   []sql.Row{{-5.01}},
			},
			{
				Statement: `SELECT to_number('5 4 4 4 4 8 . 7 8', '9 9 9 9 9 9 . 9 9');`,
				Results:   []sql.Row{{544448.78}},
			},
			{
				Statement: `SELECT to_number('.01', 'FM9.99');`,
				Results:   []sql.Row{{0.01}},
			},
			{
				Statement: `SELECT to_number('.0', '99999999.99999999');`,
				Results:   []sql.Row{{0.0}},
			},
			{
				Statement: `SELECT to_number('0', '99.99');`,
				Results:   []sql.Row{{0}},
			},
			{
				Statement: `SELECT to_number('.-01', 'S99.99');`,
				Results:   []sql.Row{{-0.01}},
			},
			{
				Statement: `SELECT to_number('.01-', '99.99S');`,
				Results:   []sql.Row{{-0.01}},
			},
			{
				Statement: `SELECT to_number(' . 0 1-', ' 9 9 . 9 9 S');`,
				Results:   []sql.Row{{-0.01}},
			},
			{
				Statement: `SELECT to_number('34,50','999,99');`,
				Results:   []sql.Row{{3450}},
			},
			{
				Statement: `SELECT to_number('123,000','999G');`,
				Results:   []sql.Row{{123}},
			},
			{
				Statement: `SELECT to_number('123456','999G999');`,
				Results:   []sql.Row{{123456}},
			},
			{
				Statement: `SELECT to_number('$1234.56','L9,999.99');`,
				Results:   []sql.Row{{1234.56}},
			},
			{
				Statement: `SELECT to_number('$1234.56','L99,999.99');`,
				Results:   []sql.Row{{1234.56}},
			},
			{
				Statement: `SELECT to_number('$1,234.56','L99,999.99');`,
				Results:   []sql.Row{{1234.56}},
			},
			{
				Statement: `SELECT to_number('1234.56','L99,999.99');`,
				Results:   []sql.Row{{1234.56}},
			},
			{
				Statement: `SELECT to_number('1,234.56','L99,999.99');`,
				Results:   []sql.Row{{1234.56}},
			},
			{
				Statement: `SELECT to_number('42nd', '99th');`,
				Results:   []sql.Row{{42}},
			},
			{
				Statement: `RESET lc_numeric;`,
			},
			{
				Statement: `CREATE TABLE num_input_test (n1 numeric);`,
			},
			{
				Statement: `INSERT INTO num_input_test(n1) VALUES (' 123');`,
			},
			{
				Statement: `INSERT INTO num_input_test(n1) VALUES ('   3245874    ');`,
			},
			{
				Statement: `INSERT INTO num_input_test(n1) VALUES ('  -93853');`,
			},
			{
				Statement: `INSERT INTO num_input_test(n1) VALUES ('555.50');`,
			},
			{
				Statement: `INSERT INTO num_input_test(n1) VALUES ('-555.50');`,
			},
			{
				Statement: `INSERT INTO num_input_test(n1) VALUES ('NaN ');`,
			},
			{
				Statement: `INSERT INTO num_input_test(n1) VALUES ('        nan');`,
			},
			{
				Statement: `INSERT INTO num_input_test(n1) VALUES (' inf ');`,
			},
			{
				Statement: `INSERT INTO num_input_test(n1) VALUES (' +inf ');`,
			},
			{
				Statement: `INSERT INTO num_input_test(n1) VALUES (' -inf ');`,
			},
			{
				Statement: `INSERT INTO num_input_test(n1) VALUES (' Infinity ');`,
			},
			{
				Statement: `INSERT INTO num_input_test(n1) VALUES (' +inFinity ');`,
			},
			{
				Statement: `INSERT INTO num_input_test(n1) VALUES (' -INFINITY ');`,
			},
			{
				Statement:   `INSERT INTO num_input_test(n1) VALUES ('     ');`,
				ErrorString: `invalid input syntax for type numeric: "     "`,
			},
			{
				Statement:   `INSERT INTO num_input_test(n1) VALUES ('   1234   %');`,
				ErrorString: `invalid input syntax for type numeric: "   1234   %"`,
			},
			{
				Statement:   `INSERT INTO num_input_test(n1) VALUES ('xyz');`,
				ErrorString: `invalid input syntax for type numeric: "xyz"`,
			},
			{
				Statement:   `INSERT INTO num_input_test(n1) VALUES ('- 1234');`,
				ErrorString: `invalid input syntax for type numeric: "- 1234"`,
			},
			{
				Statement:   `INSERT INTO num_input_test(n1) VALUES ('5 . 0');`,
				ErrorString: `invalid input syntax for type numeric: "5 . 0"`,
			},
			{
				Statement:   `INSERT INTO num_input_test(n1) VALUES ('5. 0   ');`,
				ErrorString: `invalid input syntax for type numeric: "5. 0   "`,
			},
			{
				Statement:   `INSERT INTO num_input_test(n1) VALUES ('');`,
				ErrorString: `invalid input syntax for type numeric: ""`,
			},
			{
				Statement:   `INSERT INTO num_input_test(n1) VALUES (' N aN ');`,
				ErrorString: `invalid input syntax for type numeric: " N aN "`,
			},
			{
				Statement:   `INSERT INTO num_input_test(n1) VALUES ('+ infinity');`,
				ErrorString: `invalid input syntax for type numeric: "+ infinity"`,
			},
			{
				Statement: `SELECT * FROM num_input_test;`,
				Results:   []sql.Row{{123}, {3245874}, {-93853}, {555.50}, {-555.50}, {`NaN`}, {`NaN`}, {`Infinity`}, {`Infinity`}, {`-Infinity`}, {`Infinity`}, {`Infinity`}, {`-Infinity`}},
			},
			{
				Statement: `CREATE TABLE num_typemod_test (
  millions numeric(3, -6),
  thousands numeric(3, -3),
  units numeric(3, 0),
  thousandths numeric(3, 3),
  millionths numeric(3, 6)
);`,
			},
			{
				Statement: `\d num_typemod_test
               Table "public.num_typemod_test"
   Column    |     Type      | Collation | Nullable | Default 
-------------+---------------+-----------+----------+---------
 millions    | numeric(3,-6) |           |          | 
 thousands   | numeric(3,-3) |           |          | 
 units       | numeric(3,0)  |           |          | 
 thousandths | numeric(3,3)  |           |          | 
 millionths  | numeric(3,6)  |           |          | 
INSERT INTO num_typemod_test VALUES (123456, 123, 0.123, 0.000123, 0.000000123);`,
			},
			{
				Statement: `INSERT INTO num_typemod_test VALUES (654321, 654, 0.654, 0.000654, 0.000000654);`,
			},
			{
				Statement: `INSERT INTO num_typemod_test VALUES (2345678, 2345, 2.345, 0.002345, 0.000002345);`,
			},
			{
				Statement: `INSERT INTO num_typemod_test VALUES (7654321, 7654, 7.654, 0.007654, 0.000007654);`,
			},
			{
				Statement: `INSERT INTO num_typemod_test VALUES (12345678, 12345, 12.345, 0.012345, 0.000012345);`,
			},
			{
				Statement: `INSERT INTO num_typemod_test VALUES (87654321, 87654, 87.654, 0.087654, 0.000087654);`,
			},
			{
				Statement: `INSERT INTO num_typemod_test VALUES (123456789, 123456, 123.456, 0.123456, 0.000123456);`,
			},
			{
				Statement: `INSERT INTO num_typemod_test VALUES (987654321, 987654, 987.654, 0.987654, 0.000987654);`,
			},
			{
				Statement: `INSERT INTO num_typemod_test VALUES ('NaN', 'NaN', 'NaN', 'NaN', 'NaN');`,
			},
			{
				Statement: `SELECT scale(millions), * FROM num_typemod_test ORDER BY millions;`,
				Results:   []sql.Row{{0, 0, 0, 0, 0.000, 0.000000}, {0, 1000000, 1000, 1, 0.001, 0.000001}, {0, 2000000, 2000, 2, 0.002, 0.000002}, {0, 8000000, 8000, 8, 0.008, 0.000008}, {0, 12000000, 12000, 12, 0.012, 0.000012}, {0, 88000000, 88000, 88, 0.088, 0.000088}, {0, 123000000, 123000, 123, 0.123, 0.000123}, {0, 988000000, 988000, 988, 0.988, 0.000988}, {``, `NaN`, `NaN`, `NaN`, `NaN`, `NaN`}},
			},
			{
				Statement:   `INSERT INTO num_typemod_test (millions) VALUES ('inf');`,
				ErrorString: `numeric field overflow`,
			},
			{
				Statement:   `INSERT INTO num_typemod_test (millions) VALUES (999500000);`,
				ErrorString: `numeric field overflow`,
			},
			{
				Statement:   `INSERT INTO num_typemod_test (thousands) VALUES (999500);`,
				ErrorString: `numeric field overflow`,
			},
			{
				Statement:   `INSERT INTO num_typemod_test (units) VALUES (999.5);`,
				ErrorString: `numeric field overflow`,
			},
			{
				Statement:   `INSERT INTO num_typemod_test (thousandths) VALUES (0.9995);`,
				ErrorString: `numeric field overflow`,
			},
			{
				Statement:   `INSERT INTO num_typemod_test (millionths) VALUES (0.0009995);`,
				ErrorString: `numeric field overflow`,
			},
			{
				Statement: `select 4790999999999999999999999999999999999999999999999999999999999999999999999999999999999999 * 9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999;`,
				Results:   []sql.Row{{47909999999999999999999999999999999999999999999999999999999999999999999999999999999999985209000000000000000000000000000000000000000000000000000000000000000000000000000000000001.0}},
			},
			{
				Statement: `select 4789999999999999999999999999999999999999999999999999999999999999999999999999999999999999 * 9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999;`,
				Results:   []sql.Row{{47899999999999999999999999999999999999999999999999999999999999999999999999999999999999985210000000000000000000000000000000000000000000000000000000000000000000000000000000000001.0}},
			},
			{
				Statement: `select 4770999999999999999999999999999999999999999999999999999999999999999999999999999999999999 * 9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999;`,
				Results:   []sql.Row{{47709999999999999999999999999999999999999999999999999999999999999999999999999999999999985229000000000000000000000000000000000000000000000000000000000000000000000000000000000001.0}},
			},
			{
				Statement: `select 4769999999999999999999999999999999999999999999999999999999999999999999999999999999999999 * 9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999;`,
				Results:   []sql.Row{{47699999999999999999999999999999999999999999999999999999999999999999999999999999999999985230000000000000000000000000000000000000000000000000000000000000000000000000000000000001.0}},
			},
			{
				Statement: `select trim_scale((0.1 - 2e-16383) * (0.1 - 3e-16383));`,
				Results:   []sql.Row{{0.01}},
			},
			{
				Statement: `select 999999999999999999999::numeric/1000000000000000000000;`,
				Results:   []sql.Row{{1.00000000000000000000}},
			},
			{
				Statement: `select div(999999999999999999999::numeric,1000000000000000000000);`,
				Results:   []sql.Row{{0}},
			},
			{
				Statement: `select mod(999999999999999999999::numeric,1000000000000000000000);`,
				Results:   []sql.Row{{999999999999999999999.0}},
			},
			{
				Statement: `select div(-9999999999999999999999::numeric,1000000000000000000000);`,
				Results:   []sql.Row{{-9}},
			},
			{
				Statement: `select mod(-9999999999999999999999::numeric,1000000000000000000000);`,
				Results:   []sql.Row{{-999999999999999999999.0}},
			},
			{
				Statement: `select div(-9999999999999999999999::numeric,1000000000000000000000)*1000000000000000000000 + mod(-9999999999999999999999::numeric,1000000000000000000000);`,
				Results:   []sql.Row{{-9999999999999999999999.0}},
			},
			{
				Statement: `select mod (70.0,70) ;`,
				Results:   []sql.Row{{0.0}},
			},
			{
				Statement: `select div (70.0,70) ;`,
				Results:   []sql.Row{{1}},
			},
			{
				Statement: `select 70.0 / 70 ;`,
				Results:   []sql.Row{{1.00000000000000000000}},
			},
			{
				Statement: `select 12345678901234567890 % 123;`,
				Results:   []sql.Row{{78}},
			},
			{
				Statement: `select 12345678901234567890 / 123;`,
				Results:   []sql.Row{{100371373180768845}},
			},
			{
				Statement: `select div(12345678901234567890, 123);`,
				Results:   []sql.Row{{100371373180768844}},
			},
			{
				Statement: `select div(12345678901234567890, 123) * 123 + 12345678901234567890 % 123;`,
				Results:   []sql.Row{{12345678901234567890.0}},
			},
			{
				Statement: `select sqrt(1.000000000000003::numeric);`,
				Results:   []sql.Row{{1.000000000000001}},
			},
			{
				Statement: `select sqrt(1.000000000000004::numeric);`,
				Results:   []sql.Row{{1.000000000000002}},
			},
			{
				Statement: `select sqrt(96627521408608.56340355805::numeric);`,
				Results:   []sql.Row{{9829929.87811248648}},
			},
			{
				Statement: `select sqrt(96627521408608.56340355806::numeric);`,
				Results:   []sql.Row{{9829929.87811248649}},
			},
			{
				Statement: `select sqrt(515549506212297735.073688290367::numeric);`,
				Results:   []sql.Row{{718017761.766585921184}},
			},
			{
				Statement: `select sqrt(515549506212297735.073688290368::numeric);`,
				Results:   []sql.Row{{718017761.766585921185}},
			},
			{
				Statement: `select sqrt(8015491789940783531003294973900306::numeric);`,
				Results:   []sql.Row{{89529278953540017}},
			},
			{
				Statement: `select sqrt(8015491789940783531003294973900307::numeric);`,
				Results:   []sql.Row{{89529278953540018}},
			},
			{
				Statement: `select 10.0 ^ -2147483648 as rounds_to_zero;`,
				Results:   []sql.Row{{0.0000000000000000}},
			},
			{
				Statement: `select 10.0 ^ -2147483647 as rounds_to_zero;`,
				Results:   []sql.Row{{0.0000000000000000}},
			},
			{
				Statement:   `select 10.0 ^ 2147483647 as overflows;`,
				ErrorString: `value overflows numeric format`,
			},
			{
				Statement:   `select 117743296169.0 ^ 1000000000 as overflows;`,
				ErrorString: `value overflows numeric format`,
			},
			{
				Statement: `select 3.789 ^ 21;`,
				Results:   []sql.Row{{1409343026052.8716016316022141}},
			},
			{
				Statement: `select 3.789 ^ 35;`,
				Results:   []sql.Row{{177158169650516670809.3820586142670135}},
			},
			{
				Statement: `select 1.2 ^ 345;`,
				Results:   []sql.Row{{2077446682327378559843444695.5827049735727869}},
			},
			{
				Statement: `select 0.12 ^ (-20);`,
				Results:   []sql.Row{{2608405330458882702.5529619561355838}},
			},
			{
				Statement: `select 1.000000000123 ^ (-2147483648);`,
				Results:   []sql.Row{{0.7678656556403084}},
			},
			{
				Statement: `select coalesce(nullif(0.9999999999 ^ 23300000000000, 0), 0) as rounds_to_zero;`,
				Results:   []sql.Row{{0}},
			},
			{
				Statement: `select round(((1 - 1.500012345678e-1000) ^ 1.45e1003) * 1e1000);`,
				Results:   []sql.Row{{25218976308958387188077465658068501556514992509509282366.0}},
			},
			{
				Statement: `select 0.12 ^ (-25);`,
				Results:   []sql.Row{{104825960103961013959336.4983657883169110}},
			},
			{
				Statement: `select 0.5678 ^ (-85);`,
				Results:   []sql.Row{{782333637740774446257.7719390061997396}},
			},
			{
				Statement: `select coalesce(nullif(0.9999999999 ^ 70000000000000, 0), 0) as underflows;`,
				Results:   []sql.Row{{0}},
			},
			{
				Statement: `select (-1.0) ^ 2147483646;`,
				Results:   []sql.Row{{1.0000000000000000}},
			},
			{
				Statement: `select (-1.0) ^ 2147483647;`,
				Results:   []sql.Row{{-1.0000000000000000}},
			},
			{
				Statement: `select (-1.0) ^ 2147483648;`,
				Results:   []sql.Row{{1.0000000000000000}},
			},
			{
				Statement: `select (-1.0) ^ 1000000000000000;`,
				Results:   []sql.Row{{1.0000000000000000}},
			},
			{
				Statement: `select (-1.0) ^ 1000000000000001;`,
				Results:   []sql.Row{{-1.0000000000000000}},
			},
			{
				Statement: `select 0.0 ^ 0.0;`,
				Results:   []sql.Row{{1.0000000000000000}},
			},
			{
				Statement: `select (-12.34) ^ 0.0;`,
				Results:   []sql.Row{{1.0000000000000000}},
			},
			{
				Statement: `select 12.34 ^ 0.0;`,
				Results:   []sql.Row{{1.0000000000000000}},
			},
			{
				Statement: `select 0.0 ^ 12.34;`,
				Results:   []sql.Row{{0.0000000000000000}},
			},
			{
				Statement: `select 'NaN'::numeric ^ 'NaN'::numeric;`,
				Results:   []sql.Row{{`NaN`}},
			},
			{
				Statement: `select 'NaN'::numeric ^ 0;`,
				Results:   []sql.Row{{1}},
			},
			{
				Statement: `select 'NaN'::numeric ^ 1;`,
				Results:   []sql.Row{{`NaN`}},
			},
			{
				Statement: `select 0 ^ 'NaN'::numeric;`,
				Results:   []sql.Row{{`NaN`}},
			},
			{
				Statement: `select 1 ^ 'NaN'::numeric;`,
				Results:   []sql.Row{{1}},
			},
			{
				Statement:   `select 0.0 ^ (-12.34);`,
				ErrorString: `zero raised to a negative power is undefined`,
			},
			{
				Statement:   `select (-12.34) ^ 1.2;`,
				ErrorString: `a negative number raised to a non-integer power yields a complex result`,
			},
			{
				Statement: `select 32.1 ^ 9.8;`,
				Results:   []sql.Row{{580429286790711.10}},
			},
			{
				Statement: `select 32.1 ^ (-9.8);`,
				Results:   []sql.Row{{0.000000000000001722862754788209}},
			},
			{
				Statement: `select 12.3 ^ 45.6;`,
				Results:   []sql.Row{{50081010321492803393171165777624533697036806969694.9}},
			},
			{
				Statement: `select 12.3 ^ (-45.6);`,
				Results:   []sql.Row{{0.00000000000000000000000000000000000000000000000001996764828785491}},
			},
			{
				Statement: `select 1.234 ^ 5678;`,
				Results:   []sql.Row{{"307239295662090741644584872593956173493568238595074141254349565406661439636598896798876823220904084953233015553994854875890890858118656468658643918169805277399402542281777901029346337707622181574346585989613344285010764501017625366742865066948856161360224801370482171458030533346309750557140549621313515752078638620714732831815297168231790779296290266207315344008883935010274044001522606235576584215999260117523114297033944018699691024106823438431754073086813382242140602291215149759520833200152654884259619588924545324.5973362312547382"}},
			},
			{
				Statement: `select exp(0.0);`,
				Results:   []sql.Row{{1.0000000000000000}},
			},
			{
				Statement: `select exp(1.0);`,
				Results:   []sql.Row{{2.7182818284590452}},
			},
			{
				Statement: `select exp(1.0::numeric(71,70));`,
				Results:   []sql.Row{{2.7182818284590452353602874713526624977572470936999595749669676277240766}},
			},
			{
				Statement: `select exp('nan'::numeric);`,
				Results:   []sql.Row{{`NaN`}},
			},
			{
				Statement: `select exp('inf'::numeric);`,
				Results:   []sql.Row{{`Infinity`}},
			},
			{
				Statement: `select exp('-inf'::numeric);`,
				Results:   []sql.Row{{0}},
			},
			{
				Statement: `select coalesce(nullif(exp(-5000::numeric), 0), 0) as rounds_to_zero;`,
				Results:   []sql.Row{{0}},
			},
			{
				Statement: `select coalesce(nullif(exp(-10000::numeric), 0), 0) as underflows;`,
				Results:   []sql.Row{{0}},
			},
			{
				Statement: `select exp(32.999);`,
				Results:   []sql.Row{{214429043492155.053}},
			},
			{
				Statement: `select exp(-32.999);`,
				Results:   []sql.Row{{0.000000000000004663547361468248}},
			},
			{
				Statement: `select exp(123.456);`,
				Results:   []sql.Row{{413294435277809344957685441227343146614594393746575438.725}},
			},
			{
				Statement: `select exp(-123.456);`,
				Results:   []sql.Row{{0.000000000000000000000000000000000000000000000000000002419582541264601}},
			},
			{
				Statement: `select exp(1234.5678);`,
				Results:   []sql.Row{{"146549072930959479983482138503979804217622199675223653966270157446954995433819741094410764947112047906012815540251009949604426069672532417736057033099274204598385314594846509975629046864798765888104789074984927709616261452461385220475510438783429612447831614003668421849727379202555580791042606170523016207262965336641214601082882495255771621327088265411334088968112458492660609809762865582162764292604697957813514621259353683899630997077707406305730694385703091201347848855199354307506425820147289848677003277208302716466011827836279231.9667"}},
			},
			{
				Statement: `select * from generate_series(0.0::numeric, 4.0::numeric);`,
				Results:   []sql.Row{{0.0}, {1.0}, {2.0}, {3.0}, {4.0}},
			},
			{
				Statement: `select * from generate_series(0.1::numeric, 4.0::numeric, 1.3::numeric);`,
				Results:   []sql.Row{{0.1}, {1.4}, {2.7}, {4.0}},
			},
			{
				Statement: `select * from generate_series(4.0::numeric, -1.5::numeric, -2.2::numeric);`,
				Results:   []sql.Row{{4.0}, {1.8}, {-0.4}},
			},
			{
				Statement:   `select * from generate_series(-100::numeric, 100::numeric, 0::numeric);`,
				ErrorString: `step size cannot equal zero`,
			},
			{
				Statement:   `select * from generate_series(-100::numeric, 100::numeric, 'nan'::numeric);`,
				ErrorString: `step size cannot be NaN`,
			},
			{
				Statement:   `select * from generate_series('nan'::numeric, 100::numeric, 10::numeric);`,
				ErrorString: `start value cannot be NaN`,
			},
			{
				Statement:   `select * from generate_series(0::numeric, 'nan'::numeric, 10::numeric);`,
				ErrorString: `stop value cannot be NaN`,
			},
			{
				Statement:   `select * from generate_series('inf'::numeric, 'inf'::numeric, 10::numeric);`,
				ErrorString: `start value cannot be infinity`,
			},
			{
				Statement:   `select * from generate_series(0::numeric, 'inf'::numeric, 10::numeric);`,
				ErrorString: `stop value cannot be infinity`,
			},
			{
				Statement:   `select * from generate_series(0::numeric, '42'::numeric, '-inf'::numeric);`,
				ErrorString: `step size cannot be infinity`,
			},
			{
				Statement: `select (i / (10::numeric ^ 131071))::numeric(1,0)
	from generate_series(6 * (10::numeric ^ 131071),
			     9 * (10::numeric ^ 131071),
			     10::numeric ^ 131071) as a(i);`,
				Results: []sql.Row{{6}, {7}, {8}, {9}},
			},
			{
				Statement: `select * from generate_series(1::numeric, 3::numeric) i, generate_series(i,3) j;`,
				Results:   []sql.Row{{1, 1}, {1, 2}, {1, 3}, {2, 2}, {2, 3}, {3, 3}},
			},
			{
				Statement: `select * from generate_series(1::numeric, 3::numeric) i, generate_series(1,i) j;`,
				Results:   []sql.Row{{1, 1}, {2, 1}, {2, 2}, {3, 1}, {3, 2}, {3, 3}},
			},
			{
				Statement: `select * from generate_series(1::numeric, 3::numeric) i, generate_series(1,5,i) j;`,
				Results:   []sql.Row{{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}, {2, 1}, {2, 3}, {2, 5}, {3, 1}, {3, 4}},
			},
			{
				Statement:   `select ln(-12.34);`,
				ErrorString: `cannot take logarithm of a negative number`,
			},
			{
				Statement:   `select ln(0.0);`,
				ErrorString: `cannot take logarithm of zero`,
			},
			{
				Statement: `select ln(1.2345678e-28);`,
				Results:   []sql.Row{{-64.26166165451762991204894255882820859}},
			},
			{
				Statement: `select ln(0.0456789);`,
				Results:   []sql.Row{{-3.0861187944847439}},
			},
			{
				Statement: `select ln(0.349873948359354029493948309745709580730482050975);`,
				Results:   []sql.Row{{-1.050182336912082775693991697979750253056317885460}},
			},
			{
				Statement: `select ln(0.99949452);`,
				Results:   []sql.Row{{-0.00050560779808326467}},
			},
			{
				Statement: `select ln(1.00049687395);`,
				Results:   []sql.Row{{0.00049675054901370394}},
			},
			{
				Statement: `select ln(1234.567890123456789);`,
				Results:   []sql.Row{{7.1184763012977896}},
			},
			{
				Statement: `select ln(5.80397490724e5);`,
				Results:   []sql.Row{{13.271468476626518}},
			},
			{
				Statement: `select ln(9.342536355e34);`,
				Results:   []sql.Row{{80.522470935524187}},
			},
			{
				Statement:   `select log(-12.34);`,
				ErrorString: `cannot take logarithm of a negative number`,
			},
			{
				Statement: `CONTEXT:  SQL function "log" statement 1
select log(0.0);`,
				ErrorString: `cannot take logarithm of zero`,
			},
			{
				Statement: `CONTEXT:  SQL function "log" statement 1
select log(1.234567e-89);`,
				Results: []sql.Row{{-88.90848533591373725637496492944925187293052336306443143312825869985819779294142441287021741054275}},
			},
			{
				Statement: `select log(3.4634998359873254962349856073435545);`,
				Results:   []sql.Row{{0.5395151714070134409152404011959981}},
			},
			{
				Statement: `select log(9.999999999999999999);`,
				Results:   []sql.Row{{1.000000000000000000}},
			},
			{
				Statement: `select log(10.00000000000000000);`,
				Results:   []sql.Row{{1.00000000000000000}},
			},
			{
				Statement: `select log(10.00000000000000001);`,
				Results:   []sql.Row{{1.00000000000000000}},
			},
			{
				Statement: `select log(590489.45235237);`,
				Results:   []sql.Row{{5.771212144411727}},
			},
			{
				Statement:   `select log(-12.34, 56.78);`,
				ErrorString: `cannot take logarithm of a negative number`,
			},
			{
				Statement:   `select log(-12.34, -56.78);`,
				ErrorString: `cannot take logarithm of a negative number`,
			},
			{
				Statement:   `select log(12.34, -56.78);`,
				ErrorString: `cannot take logarithm of a negative number`,
			},
			{
				Statement:   `select log(0.0, 12.34);`,
				ErrorString: `cannot take logarithm of zero`,
			},
			{
				Statement:   `select log(12.34, 0.0);`,
				ErrorString: `cannot take logarithm of zero`,
			},
			{
				Statement:   `select log(1.0, 12.34);`,
				ErrorString: `division by zero`,
			},
			{
				Statement: `select log(1.23e-89, 6.4689e45);`,
				Results:   []sql.Row{{-0.5152489207781856983977054971756484879653568168479201885425588841094788842469115325262329756}},
			},
			{
				Statement: `select log(0.99923, 4.58934e34);`,
				Results:   []sql.Row{{-103611.55579544132}},
			},
			{
				Statement: `select log(1.000016, 8.452010e18);`,
				Results:   []sql.Row{{2723830.2877097365}},
			},
			{
				Statement: `select log(3.1954752e47, 9.4792021e-73);`,
				Results:   []sql.Row{{-1.51613372350688302142917386143459361608600157692779164475351842333265418126982165}},
			},
			{
				Statement: `select scale(numeric 'NaN');`,
				Results:   []sql.Row{{``}},
			},
			{
				Statement: `select scale(numeric 'inf');`,
				Results:   []sql.Row{{``}},
			},
			{
				Statement: `select scale(NULL::numeric);`,
				Results:   []sql.Row{{``}},
			},
			{
				Statement: `select scale(1.12);`,
				Results:   []sql.Row{{2}},
			},
			{
				Statement: `select scale(0);`,
				Results:   []sql.Row{{0}},
			},
			{
				Statement: `select scale(0.00);`,
				Results:   []sql.Row{{2}},
			},
			{
				Statement: `select scale(1.12345);`,
				Results:   []sql.Row{{5}},
			},
			{
				Statement: `select scale(110123.12475871856128);`,
				Results:   []sql.Row{{14}},
			},
			{
				Statement: `select scale(-1123.12471856128);`,
				Results:   []sql.Row{{11}},
			},
			{
				Statement: `select scale(-13.000000000000000);`,
				Results:   []sql.Row{{15}},
			},
			{
				Statement: `select min_scale(numeric 'NaN') is NULL; -- should be true`,
				Results:   []sql.Row{{true}},
			},
			{
				Statement: `select min_scale(numeric 'inf') is NULL; -- should be true`,
				Results:   []sql.Row{{true}},
			},
			{
				Statement: `select min_scale(0);                     -- no digits`,
				Results:   []sql.Row{{0}},
			},
			{
				Statement: `select min_scale(0.00);                  -- no digits again`,
				Results:   []sql.Row{{0}},
			},
			{
				Statement: `select min_scale(1.0);                   -- no scale`,
				Results:   []sql.Row{{0}},
			},
			{
				Statement: `select min_scale(1.1);                   -- scale 1`,
				Results:   []sql.Row{{1}},
			},
			{
				Statement: `select min_scale(1.12);                  -- scale 2`,
				Results:   []sql.Row{{2}},
			},
			{
				Statement: `select min_scale(1.123);                 -- scale 3`,
				Results:   []sql.Row{{3}},
			},
			{
				Statement: `select min_scale(1.1234);                -- scale 4, filled digit`,
				Results:   []sql.Row{{4}},
			},
			{
				Statement: `select min_scale(1.12345);               -- scale 5, 2 NDIGITS`,
				Results:   []sql.Row{{5}},
			},
			{
				Statement: `select min_scale(1.1000);                -- 1 pos in NDIGITS`,
				Results:   []sql.Row{{1}},
			},
			{
				Statement: `select min_scale(1e100);                 -- very big number`,
				Results:   []sql.Row{{0}},
			},
			{
				Statement: `select trim_scale(numeric 'NaN');`,
				Results:   []sql.Row{{`NaN`}},
			},
			{
				Statement: `select trim_scale(numeric 'inf');`,
				Results:   []sql.Row{{`Infinity`}},
			},
			{
				Statement: `select trim_scale(1.120);`,
				Results:   []sql.Row{{1.12}},
			},
			{
				Statement: `select trim_scale(0);`,
				Results:   []sql.Row{{0}},
			},
			{
				Statement: `select trim_scale(0.00);`,
				Results:   []sql.Row{{0}},
			},
			{
				Statement: `select trim_scale(1.1234500);`,
				Results:   []sql.Row{{1.12345}},
			},
			{
				Statement: `select trim_scale(110123.12475871856128000);`,
				Results:   []sql.Row{{110123.12475871856128}},
			},
			{
				Statement: `select trim_scale(-1123.124718561280000000);`,
				Results:   []sql.Row{{-1123.12471856128}},
			},
			{
				Statement: `select trim_scale(-13.00000000000000000000);`,
				Results:   []sql.Row{{-13}},
			},
			{
				Statement: `select trim_scale(1e100);`,
				Results:   []sql.Row{{10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000.0}},
			},
			{
				Statement: `SELECT SUM(9999::numeric) FROM generate_series(1, 100000);`,
				Results:   []sql.Row{{999900000}},
			},
			{
				Statement: `SELECT SUM((-9999)::numeric) FROM generate_series(1, 100000);`,
				Results:   []sql.Row{{-999900000}},
			},
			{
				Statement: `CREATE TABLE num_variance (a numeric);`,
			},
			{
				Statement: `INSERT INTO num_variance VALUES (0);`,
			},
			{
				Statement: `INSERT INTO num_variance VALUES (3e-500);`,
			},
			{
				Statement: `INSERT INTO num_variance VALUES (-3e-500);`,
			},
			{
				Statement: `INSERT INTO num_variance VALUES (4e-500 - 1e-16383);`,
			},
			{
				Statement: `INSERT INTO num_variance VALUES (-4e-500 + 1e-16383);`,
			},
			{
				Statement: `SELECT trim_scale(variance(a) * 1e1000) FROM num_variance;`,
				Results:   []sql.Row{{12}},
			},
			{
				Statement: `BEGIN;`,
			},
			{
				Statement: `ALTER TABLE num_variance SET (parallel_workers = 4);`,
			},
			{
				Statement: `SET LOCAL parallel_setup_cost = 0;`,
			},
			{
				Statement: `SET LOCAL max_parallel_workers_per_gather = 4;`,
			},
			{
				Statement: `SELECT trim_scale(variance(a) * 1e1000) FROM num_variance;`,
				Results:   []sql.Row{{12}},
			},
			{
				Statement: `ROLLBACK;`,
			},
			{
				Statement: `DELETE FROM num_variance;`,
			},
			{
				Statement: `INSERT INTO num_variance SELECT 9e131071 + x FROM generate_series(1, 5) x;`,
			},
			{
				Statement: `SELECT variance(a) FROM num_variance;`,
				Results:   []sql.Row{{2.5000000000000000}},
			},
			{
				Statement: `BEGIN;`,
			},
			{
				Statement: `ALTER TABLE num_variance SET (parallel_workers = 4);`,
			},
			{
				Statement: `SET LOCAL parallel_setup_cost = 0;`,
			},
			{
				Statement: `SET LOCAL max_parallel_workers_per_gather = 4;`,
			},
			{
				Statement: `SELECT variance(a) FROM num_variance;`,
				Results:   []sql.Row{{2.5000000000000000}},
			},
			{
				Statement: `ROLLBACK;`,
			},
			{
				Statement: `DROP TABLE num_variance;`,
			},
			{
				Statement: `SELECT a, b, gcd(a, b), gcd(a, -b), gcd(-b, a), gcd(-b, -a)
FROM (VALUES (0::numeric, 0::numeric),
             (0::numeric, numeric 'NaN'),
             (0::numeric, 46375::numeric),
             (433125::numeric, 46375::numeric),
             (43312.5::numeric, 4637.5::numeric),
             (4331.250::numeric, 463.75000::numeric),
             ('inf', '0'),
             ('inf', '42'),
             ('inf', 'inf')
     ) AS v(a, b);`,
				Results: []sql.Row{{0, 0, 0, 0, 0, 0}, {0, `NaN`, `NaN`, `NaN`, `NaN`, `NaN`}, {0, 46375, 46375, 46375, 46375, 46375}, {433125, 46375, 875, 875, 875, 875}, {43312.5, 4637.5, 87.5, 87.5, 87.5, 87.5}, {4331.250, 463.75000, 8.75000, 8.75000, 8.75000, 8.75000}, {`Infinity`, 0, `NaN`, `NaN`, `NaN`, `NaN`}, {`Infinity`, 42, `NaN`, `NaN`, `NaN`, `NaN`}, {`Infinity`, `Infinity`, `NaN`, `NaN`, `NaN`, `NaN`}},
			},
			{
				Statement: `SELECT a,b, lcm(a, b), lcm(a, -b), lcm(-b, a), lcm(-b, -a)
FROM (VALUES (0::numeric, 0::numeric),
             (0::numeric, numeric 'NaN'),
             (0::numeric, 13272::numeric),
             (13272::numeric, 13272::numeric),
             (423282::numeric, 13272::numeric),
             (42328.2::numeric, 1327.2::numeric),
             (4232.820::numeric, 132.72000::numeric),
             ('inf', '0'),
             ('inf', '42'),
             ('inf', 'inf')
     ) AS v(a, b);`,
				Results: []sql.Row{{0, 0, 0, 0, 0, 0}, {0, `NaN`, `NaN`, `NaN`, `NaN`, `NaN`}, {0, 13272, 0, 0, 0, 0}, {13272, 13272, 13272, 13272, 13272, 13272}, {423282, 13272, 11851896, 11851896, 11851896, 11851896}, {42328.2, 1327.2, 1185189.6, 1185189.6, 1185189.6, 1185189.6}, {4232.820, 132.72000, 118518.96000, 118518.96000, 118518.96000, 118518.96000}, {`Infinity`, 0, `NaN`, `NaN`, `NaN`, `NaN`}, {`Infinity`, 42, `NaN`, `NaN`, `NaN`, `NaN`}, {`Infinity`, `Infinity`, `NaN`, `NaN`, `NaN`, `NaN`}},
			},
			{
				Statement:   `SELECT lcm(9999 * (10::numeric)^131068 + (10::numeric^131068 - 1), 2); -- overflow`,
				ErrorString: `value overflows numeric format`,
			},
			{
				Statement: `SELECT factorial(4);`,
				Results:   []sql.Row{{24}},
			},
			{
				Statement: `SELECT factorial(15);`,
				Results:   []sql.Row{{1307674368000}},
			},
			{
				Statement:   `SELECT factorial(100000);`,
				ErrorString: `value overflows numeric format`,
			},
			{
				Statement: `SELECT factorial(0);`,
				Results:   []sql.Row{{1}},
			},
			{
				Statement:   `SELECT factorial(-4);`,
				ErrorString: `factorial of a negative number is undefined`,
			},
			{
				Statement: `SELECT pg_lsn(23783416::numeric);`,
				Results:   []sql.Row{{`0/16AE7F8`}},
			},
			{
				Statement: `SELECT pg_lsn(0::numeric);`,
				Results:   []sql.Row{{`0/0`}},
			},
			{
				Statement: `SELECT pg_lsn(18446744073709551615::numeric);`,
				Results:   []sql.Row{{`FFFFFFFF/FFFFFFFF`}},
			},
			{
				Statement:   `SELECT pg_lsn(-1::numeric);`,
				ErrorString: `pg_lsn out of range`,
			},
			{
				Statement:   `SELECT pg_lsn(18446744073709551616::numeric);`,
				ErrorString: `pg_lsn out of range`,
			},
			{
				Statement:   `SELECT pg_lsn('NaN'::numeric);`,
				ErrorString: `cannot convert NaN to pg_lsn`,
			},
		},
	})
}
