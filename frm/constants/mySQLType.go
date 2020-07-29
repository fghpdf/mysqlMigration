package constants

import (
	"encoding/binary"
	"fghpdf.com/mysqlMigration/frm/utils"
	"fmt"
	"strings"
)

type MySQLType struct {
	Code uint
	Name string
}

var (
	DECIMAL     = MySQLType{Code: 0, Name: "DECIMAL"}
	TINY        = MySQLType{Code: 1, Name: "TINY"}
	SHORT       = MySQLType{Code: 2, Name: "SHORT"}
	LONG        = MySQLType{Code: 3, Name: "LONG"}
	FLOAT       = MySQLType{Code: 4, Name: "FLOAT"}
	DOUBLE      = MySQLType{Code: 5, Name: "DOUBLE"}
	NULL        = MySQLType{Code: 6, Name: "NULL"}
	TIMESTAMP   = MySQLType{Code: 7, Name: "TIMESTAMP"}
	LONGLONG    = MySQLType{Code: 8, Name: "LONGLONG"}
	INT24       = MySQLType{Code: 9, Name: "INT24"}
	DATE        = MySQLType{Code: 10, Name: "DATE"}
	TIME        = MySQLType{Code: 11, Name: "TIME"}
	DATETIME    = MySQLType{Code: 12, Name: "DATETIME"}
	YEAR        = MySQLType{Code: 13, Name: "YEAR"}
	NEWDATE     = MySQLType{Code: 14, Name: "NEWDATE"}
	VARCHAR     = MySQLType{Code: 15, Name: "VARCHAR"}
	BIT         = MySQLType{Code: 16, Name: "BIT"}
	TIMESTAMP2  = MySQLType{Code: 17, Name: "TIMESTAMP2"}
	DATETIME2   = MySQLType{Code: 18, Name: "DATETIME2"}
	TIME2       = MySQLType{Code: 19, Name: "TIME2"}
	NEWDECIMAL  = MySQLType{Code: 246, Name: "NEWDECIMAL"}
	ENUM        = MySQLType{Code: 247, Name: "ENUM"}
	SET         = MySQLType{Code: 248, Name: "SET"}
	TINY_BLOB   = MySQLType{Code: 249, Name: "TINY_BLOB"}
	MEDIUM_BLOB = MySQLType{Code: 250, Name: "MEDIUM_BLOB"}
	LONG_BLOB   = MySQLType{Code: 251, Name: "LONG_BLOB"}
	BLOB        = MySQLType{Code: 252, Name: "BLOB"}
	VAR_STRING  = MySQLType{Code: 253, Name: "VAR_STRING"}
	STRING      = MySQLType{Code: 254, Name: "STRING"}
	GEOMETRY    = MySQLType{Code: 255, Name: "GEOMETRY"}
)

var codeToMySQLTypeMap = map[uint64]MySQLType{
	0:   DECIMAL,
	1:   TINY,
	2:   SHORT,
	3:   LONG,
	4:   FLOAT,
	5:   DOUBLE,
	6:   NULL,
	7:   TIMESTAMP,
	8:   LONGLONG,
	9:   INT24,
	10:  DATE,
	11:  TIME,
	12:  DATETIME,
	13:  YEAR,
	14:  NEWDATE,
	15:  VARCHAR,
	16:  BIT,
	17:  TIMESTAMP2,
	18:  DATETIME2,
	19:  TIME2,
	246: NEWDECIMAL,
	247: ENUM,
	248: SET,
	249: TINY_BLOB,
	250: MEDIUM_BLOB,
	251: LONG_BLOB,
	252: BLOB,
	253: VAR_STRING,
	254: STRING,
	255: GEOMETRY,
}

func GetMySQLTypeFromCode(code uint64) *MySQLType {
	mysqlType := codeToMySQLTypeMap[code]
	return &mysqlType
}

func formatDefaultValue(value uint64) string {
	return fmt.Sprintf("/'%d/'", value)
}

func GetDefaultValue(data []byte, isDecimal bool, sqlType *MySQLType) string {
	if *sqlType == TINY {
		if isDecimal {
			x := int8(data[0])
			return fmt.Sprintf("/'%d/'", x)
		} else {
			x := uint64(data[0])
			return formatDefaultValue(x)
		}
	}

	if *sqlType == SHORT {
		x := binary.LittleEndian.Uint16(data)
		if isDecimal {
			return fmt.Sprintf("/'%d/'", int16(x))
		} else {
			return formatDefaultValue(uint64(x))
		}
	}

	if *sqlType == INT24 {
		// TODO: int24
	}

	if *sqlType == LONG {
		x := binary.LittleEndian.Uint32(data)
		if isDecimal {
			return fmt.Sprintf("/'%d/'", int32(x))
		} else {
			return formatDefaultValue(uint64(x))
		}
	}

	if *sqlType == LONGLONG {
		x := binary.LittleEndian.Uint64(data)
		if isDecimal {
			return fmt.Sprintf("/'%d/'", int64(x))
		} else {
			return formatDefaultValue(uint64(x))
		}
	}

	return ""
}

type FormatTypeOptions struct {
	Length      uint64
	Flags       *[]FieldFlag
	uniregCheck *UType
}

func (sqlType *MySQLType) FormatType(opts FormatTypeOptions) string {
	name := strings.ToLower(sqlType.Name)

	if utils.StringInSlice(name, []string{"tinyint", "smallint", "mediumint", "int", "bigint"}) {
		return formatNumber(name, opts)
	}

	if name == "newdecimal" {
		// TODO
	}

	return ""
}

func formatNumber(name string, opts FormatTypeOptions) string {
	res := name
	isDecimal := false
	isZeroFill := false

	for _, flag := range *opts.Flags {
		if flag.Name == "DECIMAL" {
			isDecimal = true
		}

		if flag.Name == "ZEROFILL" {
			isZeroFill = true
		}
	}

	if opts.Length != 0 {
		res += fmt.Sprintf("({%d})", opts.Length)
	}

	if !isDecimal {
		res += " unsigned"
	}

	if isZeroFill {
		res += " zerofill"
	}

	return res
}
