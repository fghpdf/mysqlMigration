package frm

import "fghpdf.com/mysqlMigration/frm/constants"

type TableOptions struct {
	Connection    string
	Engine        string
	Charset       constants.Charset
	MinRows       uint64
	MaxRows       uint64
	AvgRowLength  uint64
	HandlerOption []constants.HaOption
	RowFormat     constants.HaRowType
	KeyBlockSize  uint64
	Comment       string
	PartitionInfo string
}
