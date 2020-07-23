package frm

type byteSlice []byte

func (b byteSlice) readData(offset uint64, length uint64) []byte {
	return b[offset : offset+length]
}

/**
parseExtraInfo get data from extra info
return:
	connection engine partitionInfo
*/
func (extraInfo byteSlice) parseExtraInfo() (string, string, string) {
	extraInfoSize := uint64(len(extraInfo))
	var connection, engine, partitionInfo string

	offset := uint64(0)

	if extraInfoSize > offset {
		connectionLength := extraInfo.getExtraInfoDataLength(offset)
		offset += 2
		connection = string(extraInfo.readData(offset, connectionLength))
		offset += connectionLength
	}

	if extraInfoSize > offset {
		engineLength := extraInfo.getExtraInfoDataLength(offset)
		offset += 2
		engine = string(extraInfo.readData(offset, engineLength))
		offset += engineLength
	}

	if extraInfoSize > offset {
		partLength := extraInfo.getExtraInfoDataLength(offset)
		offset += 2
		partitionInfo = string(extraInfo.readData(offset, partLength))
		offset += partLength
	}

	return connection, engine, partitionInfo
}

func (b byteSlice) getExtraInfoDataLength(offset uint64) uint64 {
	return uint64(b.readData(offset, 2)[0])
}
