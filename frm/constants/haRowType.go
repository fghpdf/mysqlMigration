package constants

type HaRowType struct {
	Code uint
	Name string
}

var (
	DEFAULT_HA_ROW_TYPE = HaRowType{Code: 0, Name: "DEFAULT"}
	FIXED               = HaRowType{Code: 1, Name: "FIXED"}
	DYNAMIC             = HaRowType{Code: 2, Name: "DYNAMIC"}
	COMPRESSED          = HaRowType{Code: 3, Name: "COMPRESSED"}
	REDUNDANT           = HaRowType{Code: 4, Name: "REDUNDANT"}
	COMPACT             = HaRowType{Code: 5, Name: "COMPACT"}
	UNKNOWN_6           = HaRowType{Code: 6, Name: "UNKNOWN_6"}
	TOKUDB_UNCOMPRESSED = HaRowType{Code: 7, Name: "TOKUDB_UNCOMPRESSED"}
	TOKUDB_ZLIB         = HaRowType{Code: 8, Name: "TOKUDB_ZLIB"}
	TOKUDB_SNAPPY       = HaRowType{Code: 9, Name: "TOKUDB_SNAPPY"}
	TOKUDB_QUICKLZ      = HaRowType{Code: 10, Name: "TOKUDB_QUICKLZ"}
	TOKUDB_LZMA         = HaRowType{Code: 11, Name: "TOKUDB_LZMA"}
	TOKUDB_FAST         = HaRowType{Code: 12, Name: "TOKUDB_FAST"}
	TOKUDB_SMALL        = HaRowType{Code: 13, Name: "TOKUDB_SMALL"}
	TOKUDB_DEFAULT      = HaRowType{Code: 14, Name: "TOKUDB_DEFAULT"}
	UNKNOWN_15          = HaRowType{Code: 15, Name: "UNKNOWN_15"}
	UNKNOWN_16          = HaRowType{Code: 16, Name: "UNKNOWN_16"}
	UNKNOWN_17          = HaRowType{Code: 17, Name: "UNKNOWN_17"}
	UNKNOWN_18          = HaRowType{Code: 18, Name: "UNKNOWN_18"}
)

var codeToHaRowTypeMap = map[uint]HaRowType{
	0:  DEFAULT_HA_ROW_TYPE,
	1:  FIXED,
	2:  DYNAMIC,
	3:  COMPRESSED,
	4:  REDUNDANT,
	5:  COMPACT,
	6:  UNKNOWN_6,
	7:  TOKUDB_UNCOMPRESSED,
	8:  TOKUDB_ZLIB,
	9:  TOKUDB_SNAPPY,
	10: TOKUDB_QUICKLZ,
	11: TOKUDB_LZMA,
	12: TOKUDB_FAST,
	13: TOKUDB_SMALL,
	14: TOKUDB_DEFAULT,
	15: UNKNOWN_15,
	16: UNKNOWN_16,
	17: UNKNOWN_17,
	18: UNKNOWN_18,
}

func GetHaRowTypeFromCode(code uint) *HaRowType {
	rowType, ok := codeToHaRowTypeMap[code]
	if !ok {
		return &DEFAULT_HA_ROW_TYPE
	}

	if TOKUDB_DEFAULT == rowType {
		return &TOKUDB_ZLIB
	}

	if TOKUDB_FAST == rowType {
		return &TOKUDB_QUICKLZ
	}

	if TOKUDB_SMALL == rowType {
		return &TOKUDB_LZMA
	}

	return &rowType
}
