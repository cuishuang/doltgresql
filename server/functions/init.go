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

package functions

// Init initializes all functions in this package.
func Init() {
	initAbs()
	initAcos()
	initAcosd()
	initAcosh()
	initAge()
	initArrayAppend()
	initArrayToString()
	initAscii()
	initAsin()
	initAsind()
	initAsinh()
	initAtan()
	initAtan2()
	initAtan2d()
	initAtand()
	initAtanh()
	initBitLength()
	initBtrim()
	initCbrt()
	initCeil()
	initCharLength()
	initChr()
	initColDescription()
	initCos()
	initCosd()
	initCosh()
	initCot()
	initCotd()
	initCurrentDatabase()
	initCurrentSchema()
	initCurrentSchemas()
	initDegrees()
	initDiv()
	initDoltProcedures()
	initExp()
	initExtract()
	initFactorial()
	initFloor()
	initFormatType()
	initGcd()
	initInitcap()
	initLcm()
	initLeft()
	initLength()
	initLn()
	initLog()
	initLog10()
	initLower()
	initLpad()
	initLtrim()
	initMd5()
	initMinScale()
	initMod()
	initNextVal()
	initObjDescription()
	initOctetLength()
	initPgEncodingToChar()
	initPgFunctionIsVisible()
	initPgGetConstraintdef()
	initPgGetExpr()
	initPgGetFunctionIdentityArguments()
	initPgGetFunctionDef()
	initPgGetIndexDef()
	initPgGetPartKeyDef()
	initPgGetTriggerDef()
	initPgGetUserbyid()
	initPgGetViewDef()
	initPgIndexesSize()
	initPgIsInRecovery()
	initPgPostmasterStartTime()
	initPgRelationSize()
	initPgStatGetNumscans()
	initPgTableIsVisible()
	initPgTableSize()
	initPgTablespaceLocation()
	initPgTotalRelationSize()
	initPi()
	initPower()
	initQuoteIdent()
	initRadians()
	initRandom()
	initRepeat()
	initReplace()
	initReverse()
	initRight()
	initRound()
	initRpad()
	initRtrim()
	initScale()
	initSetVal()
	initShobjDescription()
	initSign()
	initSin()
	initSind()
	initSinh()
	initSplitPart()
	initSqrt()
	initStrpos()
	initSubstr()
	initTan()
	initTand()
	initTanh()
	initTimezone()
	initToHex()
	initToRegclass()
	initToRegproc()
	initToRegtype()
	initTranslate()
	initTrimScale()
	initTrunc()
	initTxidCurrent()
	initUnnest()
	initUpper()
	initVersion()
	initWidthBucket()
}
