package frm

import "fghpdf.com/mysqlMigration/frm/constants"

type Column struct {
	Name       string
	Length     uint64
	TypeCode   constants.MySQLType
	TypeName   string
	Default    byteSlice
	Attributes byteSlice
	Charset    string
	Comment    string
}
