package frm

import (
	"fghpdf.com/mysqlMigration/errors"
	"fmt"
	"strconv"
	"strings"
)

type byteSlice []byte

func (b byteSlice) readData(offset uint64, length uint64) []byte {
	return b[offset : offset+length]
}

func (b byteSlice) convertRangeToNumber(offset uint64, length uint64) uint64 {
	data := b.readData(offset, length)

	dataLen := len(data)

	// must be even
	if dataLen%2 != 0 {
		panic(errors.BYTE_LEN_MUST_BE_EVEN.New("byte length must be even"))
	}

	var sb strings.Builder

	for i := dataLen - 1; i >= 0; i = i - 2 {
		hexStr := fmt.Sprintf("%02x%02x", data[i], data[i-1])
		sb.WriteString(hexStr)
	}

	result, err := strconv.ParseUint(sb.String(), 16, 32)
	if err != nil {
		panic(errors.WrapFormat(err, "[convertByteSliceToString] strconv.ParseUint error"))
	}

	return result
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
