package frm

import (
	"fghpdf.com/mysqlMigration/errors"
	"fmt"
	"strconv"
	"strings"
)

func parse(fileData []byte) {
	versionIdData := fileData[51:55]

	keyInfoOffset := convertByteSliceToString(fileData[6:8])
	keyInfoLength := convertByteSliceToString(fileData[14:16])
	// 65535
	if keyInfoLength == 0xffff {
		keyInfoLength = convertByteSliceToString(fileData[47:51])
	}
	keyInfo := fileData[keyInfoOffset : keyInfoLength+keyInfoOffset]
	fmt.Println(keyInfo)

	mySQLVersion := getMySQLVersionFromByte(versionIdData)
	fmt.Println(mySQLVersion)
}

// MySQL version encoded as a 4-byte integer in little endian format.
// This is the value MYSQL_VERSION_ID from include/mysql_version.h in the mysql source tree.
// Example: ‘xb6xc5x00x00’ 0x0000c5b6 => 50614 => MySQL v5.6.14
func getMySQLVersionFromByte(versionIdData []byte) string {
	versionId := convertByteSliceToString(versionIdData)
	return GetMySQLVersionFromId(uint32(versionId)).Format()
}

func convertByteSliceToString(data []byte) uint64 {
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
