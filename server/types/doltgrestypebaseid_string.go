// Code generated by "stringer -type=DoltgresTypeBaseID"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DoltgresTypeBaseID_Any-8192]
	_ = x[DoltgresTypeBaseID_AnyElement-8193]
	_ = x[DoltgresTypeBaseID_AnyArray-8194]
	_ = x[DoltgresTypeBaseID_AnyNonArray-8195]
	_ = x[DoltgresTypeBaseID_AnyEnum-8196]
	_ = x[DoltgresTypeBaseID_AnyRange-8197]
	_ = x[DoltgresTypeBaseID_AnyMultirange-8198]
	_ = x[DoltgresTypeBaseID_AnyCompatible-8199]
	_ = x[DoltgresTypeBaseID_AnyCompatibleArray-8200]
	_ = x[DoltgresTypeBaseID_AnyCompatibleNonArray-8201]
	_ = x[DoltgresTypeBaseID_AnyCompatibleRange-8202]
	_ = x[DoltgresTypeBaseID_AnyCompatibleMultirange-8203]
	_ = x[DoltgresTypeBaseID_CString-8204]
	_ = x[DoltgresTypeBaseID_Internal-8205]
	_ = x[DoltgresTypeBaseID_Language_Handler-8206]
	_ = x[DoltgresTypeBaseID_FDW_Handler-8207]
	_ = x[DoltgresTypeBaseID_Table_AM_Handler-8208]
	_ = x[DoltgresTypeBaseID_Index_AM_Handler-8209]
	_ = x[DoltgresTypeBaseID_TSM_Handler-8210]
	_ = x[DoltgresTypeBaseID_Record-8211]
	_ = x[DoltgresTypeBaseID_Trigger-8212]
	_ = x[DoltgresTypeBaseID_Event_Trigger-8213]
	_ = x[DoltgresTypeBaseID_PG_DDL_Command-8214]
	_ = x[DoltgresTypeBaseID_Void-8215]
	_ = x[DoltgresTypeBaseID_Unknown-8216]
	_ = x[DoltgresTypeBaseID_Int16Serial-8217]
	_ = x[DoltgresTypeBaseID_Int32Serial-8218]
	_ = x[DoltgresTypeBaseID_Int64Serial-8219]
	_ = x[DoltgresTypeBaseID_Regclass-8220]
	_ = x[DoltgresTypeBaseID_Regcollation-8221]
	_ = x[DoltgresTypeBaseID_Regconfig-8222]
	_ = x[DoltgresTypeBaseID_Regdictionary-8223]
	_ = x[DoltgresTypeBaseID_Regnamespace-8224]
	_ = x[DoltgresTypeBaseID_Regoper-8225]
	_ = x[DoltgresTypeBaseID_Regoperator-8226]
	_ = x[DoltgresTypeBaseID_Regproc-8227]
	_ = x[DoltgresTypeBaseID_Regprocedure-8228]
	_ = x[DoltgresTypeBaseID_Regrole-8229]
	_ = x[DoltgresTypeBaseID_Regtype-8230]
	_ = x[DoltgresTypeBaseID_Bool-3]
	_ = x[DoltgresTypeBaseID_Bytea-7]
	_ = x[DoltgresTypeBaseID_Char-9]
	_ = x[DoltgresTypeBaseID_Date-15]
	_ = x[DoltgresTypeBaseID_Float32-21]
	_ = x[DoltgresTypeBaseID_Float64-23]
	_ = x[DoltgresTypeBaseID_Int16-27]
	_ = x[DoltgresTypeBaseID_Int32-29]
	_ = x[DoltgresTypeBaseID_Int64-33]
	_ = x[DoltgresTypeBaseID_Json-39]
	_ = x[DoltgresTypeBaseID_JsonB-41]
	_ = x[DoltgresTypeBaseID_Name-90]
	_ = x[DoltgresTypeBaseID_Null-53]
	_ = x[DoltgresTypeBaseID_Numeric-54]
	_ = x[DoltgresTypeBaseID_Oid-92]
	_ = x[DoltgresTypeBaseID_Text-64]
	_ = x[DoltgresTypeBaseID_Time-66]
	_ = x[DoltgresTypeBaseID_Timestamp-70]
	_ = x[DoltgresTypeBaseID_TimestampTZ-74]
	_ = x[DoltgresTypeBaseID_TimeTZ-68]
	_ = x[DoltgresTypeBaseID_Uuid-82]
	_ = x[DoltgresTypeBaseID_VarChar-86]
	_ = x[DoltgresTypeBaseID_Xid-94]
}

const _DoltgresTypeBaseID_name = "DoltgresTypeBaseID_BoolDoltgresTypeBaseID_ByteaDoltgresTypeBaseID_CharDoltgresTypeBaseID_DateDoltgresTypeBaseID_Float32DoltgresTypeBaseID_Float64DoltgresTypeBaseID_Int16DoltgresTypeBaseID_Int32DoltgresTypeBaseID_Int64DoltgresTypeBaseID_JsonDoltgresTypeBaseID_JsonBDoltgresTypeBaseID_NullDoltgresTypeBaseID_NumericDoltgresTypeBaseID_TextDoltgresTypeBaseID_TimeDoltgresTypeBaseID_TimeTZDoltgresTypeBaseID_TimestampDoltgresTypeBaseID_TimestampTZDoltgresTypeBaseID_UuidDoltgresTypeBaseID_VarCharDoltgresTypeBaseID_NameDoltgresTypeBaseID_OidDoltgresTypeBaseID_XidDoltgresTypeBaseID_AnyDoltgresTypeBaseID_AnyElementDoltgresTypeBaseID_AnyArrayDoltgresTypeBaseID_AnyNonArrayDoltgresTypeBaseID_AnyEnumDoltgresTypeBaseID_AnyRangeDoltgresTypeBaseID_AnyMultirangeDoltgresTypeBaseID_AnyCompatibleDoltgresTypeBaseID_AnyCompatibleArrayDoltgresTypeBaseID_AnyCompatibleNonArrayDoltgresTypeBaseID_AnyCompatibleRangeDoltgresTypeBaseID_AnyCompatibleMultirangeDoltgresTypeBaseID_CStringDoltgresTypeBaseID_InternalDoltgresTypeBaseID_Language_HandlerDoltgresTypeBaseID_FDW_HandlerDoltgresTypeBaseID_Table_AM_HandlerDoltgresTypeBaseID_Index_AM_HandlerDoltgresTypeBaseID_TSM_HandlerDoltgresTypeBaseID_RecordDoltgresTypeBaseID_TriggerDoltgresTypeBaseID_Event_TriggerDoltgresTypeBaseID_PG_DDL_CommandDoltgresTypeBaseID_VoidDoltgresTypeBaseID_UnknownDoltgresTypeBaseID_Int16SerialDoltgresTypeBaseID_Int32SerialDoltgresTypeBaseID_Int64SerialDoltgresTypeBaseID_RegclassDoltgresTypeBaseID_RegcollationDoltgresTypeBaseID_RegconfigDoltgresTypeBaseID_RegdictionaryDoltgresTypeBaseID_RegnamespaceDoltgresTypeBaseID_RegoperDoltgresTypeBaseID_RegoperatorDoltgresTypeBaseID_RegprocDoltgresTypeBaseID_RegprocedureDoltgresTypeBaseID_RegroleDoltgresTypeBaseID_Regtype"

var _DoltgresTypeBaseID_map = map[DoltgresTypeBaseID]string{
	3:    _DoltgresTypeBaseID_name[0:23],
	7:    _DoltgresTypeBaseID_name[23:47],
	9:    _DoltgresTypeBaseID_name[47:70],
	15:   _DoltgresTypeBaseID_name[70:93],
	21:   _DoltgresTypeBaseID_name[93:119],
	23:   _DoltgresTypeBaseID_name[119:145],
	27:   _DoltgresTypeBaseID_name[145:169],
	29:   _DoltgresTypeBaseID_name[169:193],
	33:   _DoltgresTypeBaseID_name[193:217],
	39:   _DoltgresTypeBaseID_name[217:240],
	41:   _DoltgresTypeBaseID_name[240:264],
	53:   _DoltgresTypeBaseID_name[264:287],
	54:   _DoltgresTypeBaseID_name[287:313],
	64:   _DoltgresTypeBaseID_name[313:336],
	66:   _DoltgresTypeBaseID_name[336:359],
	68:   _DoltgresTypeBaseID_name[359:384],
	70:   _DoltgresTypeBaseID_name[384:412],
	74:   _DoltgresTypeBaseID_name[412:442],
	82:   _DoltgresTypeBaseID_name[442:465],
	86:   _DoltgresTypeBaseID_name[465:491],
	90:   _DoltgresTypeBaseID_name[491:514],
	92:   _DoltgresTypeBaseID_name[514:536],
	94:   _DoltgresTypeBaseID_name[536:558],
	8192: _DoltgresTypeBaseID_name[558:580],
	8193: _DoltgresTypeBaseID_name[580:609],
	8194: _DoltgresTypeBaseID_name[609:636],
	8195: _DoltgresTypeBaseID_name[636:666],
	8196: _DoltgresTypeBaseID_name[666:692],
	8197: _DoltgresTypeBaseID_name[692:719],
	8198: _DoltgresTypeBaseID_name[719:751],
	8199: _DoltgresTypeBaseID_name[751:783],
	8200: _DoltgresTypeBaseID_name[783:820],
	8201: _DoltgresTypeBaseID_name[820:860],
	8202: _DoltgresTypeBaseID_name[860:897],
	8203: _DoltgresTypeBaseID_name[897:939],
	8204: _DoltgresTypeBaseID_name[939:965],
	8205: _DoltgresTypeBaseID_name[965:992],
	8206: _DoltgresTypeBaseID_name[992:1027],
	8207: _DoltgresTypeBaseID_name[1027:1057],
	8208: _DoltgresTypeBaseID_name[1057:1092],
	8209: _DoltgresTypeBaseID_name[1092:1127],
	8210: _DoltgresTypeBaseID_name[1127:1157],
	8211: _DoltgresTypeBaseID_name[1157:1182],
	8212: _DoltgresTypeBaseID_name[1182:1208],
	8213: _DoltgresTypeBaseID_name[1208:1240],
	8214: _DoltgresTypeBaseID_name[1240:1273],
	8215: _DoltgresTypeBaseID_name[1273:1296],
	8216: _DoltgresTypeBaseID_name[1296:1322],
	8217: _DoltgresTypeBaseID_name[1322:1352],
	8218: _DoltgresTypeBaseID_name[1352:1382],
	8219: _DoltgresTypeBaseID_name[1382:1412],
	8220: _DoltgresTypeBaseID_name[1412:1439],
	8221: _DoltgresTypeBaseID_name[1439:1470],
	8222: _DoltgresTypeBaseID_name[1470:1498],
	8223: _DoltgresTypeBaseID_name[1498:1530],
	8224: _DoltgresTypeBaseID_name[1530:1561],
	8225: _DoltgresTypeBaseID_name[1561:1587],
	8226: _DoltgresTypeBaseID_name[1587:1617],
	8227: _DoltgresTypeBaseID_name[1617:1643],
	8228: _DoltgresTypeBaseID_name[1643:1674],
	8229: _DoltgresTypeBaseID_name[1674:1700],
	8230: _DoltgresTypeBaseID_name[1700:1726],
}

func (i DoltgresTypeBaseID) String() string {
	if str, ok := _DoltgresTypeBaseID_map[i]; ok {
		return str
	}
	return "DoltgresTypeBaseID(" + strconv.FormatInt(int64(i), 10) + ")"
}
