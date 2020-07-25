package frm

import (
	"bytes"
	"fghpdf.com/mysqlMigration/frm/constants"
	"fmt"
)

type Column struct {
	Name       string
	Length     uint64
	TypeCode   constants.MySQLType
	TypeName   string
	Default    byteSlice
	Attributes byteSlice
	Charset    string
	Comment    string
}

func parseColumnData(data packedColumnData, table Table) *[]Column {
	names := getNames(data.Names)
	labels := getLabels(data.Labels)

	nullBit := 1
	for _, value := range table.TableOptions.HandlerOptions {
		if value.Name == "PACK_RECORD" {
			nullBit = 0
		}
	}

	nullBytesLength := (data.NullCount + 1 + 7) / 8
	nullBytes := data.Defaults.readData(0, nullBytesLength)
	fmt.Println(nullBytes, nullBit, labels)

	metadataOffset := uint64(0)
	for index, name := range *names {
		length := data.Metadata.convertRangeToNumber(metadataOffset+3, 2)
		typeCode := constants.GetMySQLTypeFromCode(uint64(data.Metadata[metadataOffset+13]))

		fmt.Println(index, name, length, typeCode)
		metadataOffset += 17
	}

	return nil
}

func getNames(names byteSlice) *[]string {
	var res []string
	if len(names) <= 2 {
		return &res
	}

	// delete first char \xff and last char \xff\x00
	nameList := bytes.Split(names[1:len(names)-2], []byte{0xff})
	for _, name := range nameList {
		res = append(res, string(name))
	}

	return &res
}

func getLabels(labels byteSlice) *[]string {
	var res []string
	if len(labels) <= 1 {
		return &res
	}

	labelList := bytes.Split(labels[:len(labels)-1], []byte{0x00})
	for _, group := range labelList {
		groupList := bytes.Split(group[1:len(group)-1], []byte{0xff})
		for _, name := range groupList {
			res = append(res, string(name))
		}
	}

	return &res
}
