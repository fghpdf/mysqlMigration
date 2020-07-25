package frm

type packedColumnData struct {
	Count     uint64
	NullCount uint64
	Metadata  byteSlice
	Names     byteSlice
	Labels    byteSlice
	Comments  byteSlice
	Defaults  byteSlice
}
