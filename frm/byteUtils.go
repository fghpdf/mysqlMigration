package frm

type byteSlice []byte

func (b byteSlice) readData(offset uint64, length uint64) []byte {
	return b[offset : offset+length]
}
