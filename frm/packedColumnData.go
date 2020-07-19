package frm

type packedColumnData struct {
	Count     uint64
	NullCount uint64
	Metadata  []byte
	Names     []byte
	Labels    []byte
	Comments  []byte
	Defaults  []byte
}
