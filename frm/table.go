package frm

type Table struct {
	Name         string
	MySQLVersion MySQLVersion
	Charset      string
	TableOptions TableOptions
	Columns      []Column
}
