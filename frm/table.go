package frm

import "fghpdf.com/mysqlMigration/frm/constants"

type Table struct {
	Name         string
	MySQLVersion MySQLVersion
	Charset      constants.Charset
	TableOptions TableOptions
	Columns      []Column
}
