package constants

import "fghpdf.com/mysqlMigration/errors"

type UType struct {
	Code uint64
	Name string
}

var codeToUTypeMap = map[uint64]UType{
	0: {
		Code: 0,
		Name: "NONE",
	},
	1: {
		Code: 1,
		Name: "DATE",
	},
	2: {
		Code: 2,
		Name: "SHIELD",
	},
	3: {
		Code: 3,
		Name: "NOEMPTY",
	},
	4: {
		Code: 4,
		Name: "CASEUP",
	},
	5: {
		Code: 5,
		Name: "PNR",
	},
	6: {
		Code: 6,
		Name: "BGNR",
	},
	7: {
		Code: 7,
		Name: "PGNR",
	},
	8: {
		Code: 8,
		Name: "YES",
	},
	9: {
		Code: 9,
		Name: "NO",
	},
	10: {
		Code: 10,
		Name: "REL",
	},
	11: {
		Code: 11,
		Name: "CHECK",
	},
	12: {
		Code: 12,
		Name: "EMPTY",
	},
	13: {
		Code: 13,
		Name: "UNKNOWN_FIELD",
	},
	14: {
		Code: 14,
		Name: "CASEDN",
	},
	15: {
		Code: 15,
		Name: "NEXT_NUMBER",
	},
	16: {
		Code: 16,
		Name: "INTERVAL_FIELD",
	},
	17: {
		Code: 17,
		Name: "BIT_FIELD",
	},
	18: {
		Code: 18,
		Name: "TIMESTAMP_OLD_FIELD",
	},
	19: {
		Code: 19,
		Name: "CAPITALIZE",
	},
	20: {
		Code: 20,
		Name: "BLOB_FIELD",
	},
	21: {
		Code: 21,
		Name: "TIMESTAMP_DN_FIELD",
	},
	22: {
		Code: 22,
		Name: "TIMESTAMP_UN_FIELD",
	},
	23: {
		Code: 23,
		Name: "TIMESTAMP_DNUN_FIELD",
	},
}

func GetUTypeFromCode(code uint64) *UType {
	uType, ok := codeToUTypeMap[code]
	if !ok {
		panic(errors.NewFormat("uType code '%d' not found", code))
	}

	return &uType
}
