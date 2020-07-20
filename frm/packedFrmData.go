package frm

type packedFrmData struct {
	MySQLVersion string
	FilePath     string
	KeyInfo      byteSlice
	Defaults     byteSlice
	ExtraInfo    byteSlice
	Columns      packedColumnData
}
