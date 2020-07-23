package frm

import "fghpdf.com/mysqlMigration/frm/constants"

type TableOptions struct {
	Connection    byteSlice
	Engine        string
	Charset       constants.Charset
	MinRows       uint64
	MaxRows       uint64
	AvgRowLength  uint64
	HaOption      int
	RowFormat     constants.HaRowType
	KeyBlockSize  uint64
	Comment       string
	PartitionInfo byteSlice
}
