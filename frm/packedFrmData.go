package frm

type packedFrmData struct {
	MySQLVersion MySQLVersion
	FilePath     string
	KeyInfo      byteSlice
	Defaults     byteSlice
	ExtraInfo    byteSlice
	Columns      packedColumnData
}
