package frm

import (
	"fghpdf.com/mysqlMigration/frm/constants"
	"fmt"
)

func parse(fileData byteSlice) {
	versionIdData := fileData[0x0033:0x0037]

	// get MySQL version
	mySQLVersion := getMySQLVersionFromByte(versionIdData)

	// get key info section
	keyInfoOffset := fileData.convertRangeToNumber(0x0006, 2)
	keyInfoLength := fileData.convertRangeToNumber(0x000e, 2)
	// 65535
	if keyInfoLength == 0xffff {
		keyInfoLength = fileData.convertRangeToNumber(0x002f, 4)
	}
	keyInfo := fileData.readData(keyInfoOffset, keyInfoLength)

	// get column defaults section
	columnDefaultsOffset := keyInfoOffset + keyInfoLength
	columnDefaultsLength := fileData.convertRangeToNumber(0x0010, 2)
	columnDefaults := fileData.readData(columnDefaultsOffset, columnDefaultsLength)

	// get table extra / attributes section
	extraInfoOffset := columnDefaultsOffset + columnDefaultsLength
	extraInfoLength := fileData.convertRangeToNumber(0x0037, 4)
	extraInfo := fileData.readData(extraInfoOffset, extraInfoLength)

	// get column info section offset and length
	namesLength := fileData.convertRangeToNumber(0x0004, 2)
	headerSize := uint64(64)
	formInfoOffset := fileData.convertRangeToNumber(headerSize+namesLength, 4)
	formInfoLength := uint64(288)

	// get screens section
	screensLength := fileData.convertRangeToNumber(formInfoOffset+260, 2)

	// Column data
	nullFields := fileData.convertRangeToNumber(formInfoOffset+282, 2)
	columnCount := fileData.convertRangeToNumber(formInfoOffset+258, 2)

	// 17 bytes of metadata per column
	metaDataOffset := formInfoOffset + formInfoLength + screensLength
	metaDataLength := 17 * columnCount

	namesLength = fileData.convertRangeToNumber(formInfoOffset+268, 2)
	namesOffset := metaDataOffset + metaDataLength

	labelsLength := fileData.convertRangeToNumber(formInfoOffset+274, 2)
	labelOffset := namesOffset + namesLength

	commentsLength := fileData.convertRangeToNumber(formInfoOffset+284, 2)
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

	// get table engine
	if len(engine) == 0 {
		engine = constants.GetLegacyDBTypeFromCode(uint(fileData[0x0003])).Name
	} else if engine == "partition" {
		engine = constants.GetLegacyDBTypeFromCode(uint(fileData[0x003d])).Name
	}

	charset := constants.Lookup(uint(fileData[0x0026]))

	tableOpts := TableOptions{
		Connection:    connection,
		Engine:        engine,
		Charset:       *charset,
		MinRows:       fileData.convertRangeToNumber(0x0016, 4),
		MaxRows:       fileData.convertRangeToNumber(0x0012, 4),
		AvgRowLength:  fileData.convertRangeToNumber(0x0022, 4),
		HandlerOption: *constants.GetHaOptionsFromCode(fileData.convertRangeToNumber(0x001e, 2)),
		RowFormat:     *constants.GetHaRowTypeFromCode(uint(fileData[0x0028])),
		KeyBlockSize:  fileData.convertRangeToNumber(0x003e, 2),
		Comment:       "",
		PartitionInfo: partitionInfo,
	}

	table := Table{
		MySQLVersion: mySQLVersion,
		TableOptions: tableOpts,
		Charset:      *charset,
	}

	fmt.Println(table.TableOptions.HandlerOption)
}

// MySQL version encoded as a 4-byte integer in little endian format.
// This is the value MYSQL_VERSION_ID from include/mysql_version.h in the mysql source tree.
// Example: ‘xb6xc5x00x00’ 0x0000c5b6 => 50614 => MySQL v5.6.14
func getMySQLVersionFromByte(versionIdData byteSlice) MySQLVersion {
	versionId := versionIdData.convertRangeToNumber(0, uint64(len(versionIdData)))
	return GetMySQLVersionFromId(versionId)
}
