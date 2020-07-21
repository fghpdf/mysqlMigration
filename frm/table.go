package frm

import "fghpdf.com/mysqlMigration/frm/constants"

type table struct {
	Name         string
	MySQLVersion MySQLVersion
	Charset      string
	TableOptions tableOptions
	Columns      []column
}

type tableOptions struct {
	Connection    byteSlice
	Engine        string
	Charset       string
	MinRows       uint64
	MaxRows       uint64
	AvgRowLength  uint64
	HaOption      int
	RowFormat     constants.HaRowType
	keyBlockSize  uint64
	Comment       string
	PartitionInfo byteSlice
}

type column struct {
	Name       string
	Length     uint64
	TypeCode   constants.MySQLType
	TypeName   string
	Default    byteSlice
	Attributes byteSlice
	Charset    string
	Comment    string
}
