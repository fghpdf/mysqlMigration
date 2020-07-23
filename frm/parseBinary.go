package frm

import (
	"fghpdf.com/mysqlMigration/errors"
	"fghpdf.com/mysqlMigration/frm/constants"
	"fmt"
	"strconv"
	"strings"
)

func parse(fileData byteSlice) {
	versionIdData := fileData[51:55]

	// get MySQL version
	mySQLVersion := getMySQLVersionFromByte(versionIdData)

	// get key info section
	keyInfoOffset := convertByteSliceToString(fileData[6:8])
	keyInfoLength := convertByteSliceToString(fileData[14:16])
	// 65535
	if keyInfoLength == 0xffff {
		keyInfoLength = convertByteSliceToString(fileData[47:51])
	}
	keyInfo := fileData.readData(keyInfoOffset, keyInfoLength)

	// get column defaults section
	columnDefaultsOffset := keyInfoOffset + keyInfoLength
	columnDefaultsLength := convertByteSliceToString(fileData[16:18])
	columnDefaults := fileData.readData(columnDefaultsOffset, columnDefaultsLength)

	// get table extra / attributes section
	extraInfoOffset := columnDefaultsOffset + columnDefaultsLength
	extraInfoLength := convertByteSliceToString(fileData[0x0037:0x003b])
	extraInfo := fileData.readData(extraInfoOffset, extraInfoLength)

	// get column info section offset and length
	namesLength := convertByteSliceToString(fileData[0x0004:0x0006])
	headerSize := uint64(64)
	formInfoOffset := convertByteSliceToString(fileData[headerSize+namesLength : headerSize+namesLength+4])
	formInfoLength := uint64(288)

	// get screens section
	screensLength := convertByteSliceToString(fileData[formInfoOffset+260 : formInfoOffset+262])

	// Column data
	nullFields := convertByteSliceToString(fileData[formInfoOffset+282 : formInfoOffset+284])
	columnCount := convertByteSliceToString(fileData[formInfoOffset+258 : formInfoOffset+260])

	// 17 bytes of metadata per column
	metaDataOffset := formInfoOffset + formInfoLength + screensLength
	metaDataLength := 17 * columnCount

	namesLength = convertByteSliceToString(fileData[formInfoOffset+268 : formInfoOffset+270])
	namesOffset := metaDataOffset + metaDataLength

	labelsLength := convertByteSliceToString(fileData[formInfoOffset+274 : formInfoOffset+276])
	labelOffset := namesOffset + namesLength

	commentsLength := convertByteSliceToString(fileData[formInfoOffset+284 : formInfoOffset+286])
	commentsOffset := labelOffset + labelsLength

	columnData := packedColumnData{
		Count:     columnCount,
		NullCount: nullFields,
		Metadata:  fileData.readData(metaDataOffset, metaDataLength),
		Names:     fileData.readData(namesOffset, namesLength),
		Labels:    fileData.readData(labelOffset, labelsLength),
		Comments:  fileData.readData(commentsOffset, commentsLength),
		Defaults:  fileData.readData(columnDefaultsOffset, columnDefaultsLength),
	}

	packedFrmData := packedFrmData{
		MySQLVersion: mySQLVersion,
		KeyInfo:      keyInfo,
		Defaults:     columnDefaults,
		ExtraInfo:    extraInfo,
		Columns:      columnData,
	}

	connection, engine, partitionInfo := packedFrmData.ExtraInfo.parseExtraInfo()
	fmt.Println(connection, engine, partitionInfo)

	// get table engine
	engine = constants.GetLegacyDBTypeFromCode(uint(fileData[0x0003])).Name

	tableOpts := TableOptions{
		Engine:       engine,
		Charset:      constants.Lookup(uint(fileData[0x0026])),
		MinRows:      convertByteSliceToString(fileData[0x0016:0x001a]),
		MaxRows:      convertByteSliceToString(fileData[0x0012:0x0016]),
		AvgRowLength: convertByteSliceToString(fileData[0x0022:0x0026]),
		HaOption:     0,
		RowFormat:    *constants.GetHaRowTypeFromCode(uint(fileData[0x0028])),
		KeyBlockSize: convertByteSliceToString(fileData[0x003e:0x0040]),
		Comment:      "",
	}

	table := Table{
		TableOptions: tableOpts,
	}

	fmt.Println(table.TableOptions.Charset)
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
