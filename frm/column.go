package frm

import (
	"bytes"
	"fghpdf.com/mysqlMigration/frm/constants"
	"fmt"
)

type Column struct {
	Name         string
	Length       uint64
	TypeCode     *constants.MySQLType
	TypeName     string
	DefaultValue string
	Attributes   byteSlice
	Charset      *constants.Charset
	Comment      string
}

func parseColumnData(data packedColumnData, table Table) *[]Column {
	var res []Column

	names := getNames(data.Names)
	labels := *getLabels(data.Labels)

	nullBit := 1
	for _, value := range table.TableOptions.HandlerOptions {
		if value.Name == "PACK_RECORD" {
			nullBit = 0
		}
	}

	nullBytesLength := (data.NullCount + 1 + 7) / 8
	nullBytes := data.Defaults.readData(0, nullBytesLength)

	metadataOffset := uint64(0)
	for index, name := range *names {
		length := data.Metadata.convertRangeToNumber(metadataOffset+3, 2)
		typeCode := constants.GetMySQLTypeFromCode(uint64(data.Metadata[metadataOffset+13]))
		fieldFlags := constants.GetFieldFlagFromCode(data.Metadata.convertRangeToNumber(metadataOffset+8, 2))
		uniregCheck := constants.GetUTypeFromCode(uint64(data.Metadata[metadataOffset+10]))

		if StringInSlice(typeCode.Name, []string{"ENUM", "SET"}) {
			labelId := uint64(data.Metadata[metadataOffset+12]) - 1
			labels = []string{labels[labelId]}
		} else {
			labels = nil
		}

		defaultsOffset := data.Metadata.convertRangeToNumber(metadataOffset+5, 3) - 1
		commentLength := data.Metadata.convertRangeToNumber(metadataOffset+15, 2)

		subtypeCode := uint64(0)
		charsetId := uint64(0)
		if typeCode.Name != "GEOMETRY" {
			charsetId = (uint64(data.Metadata[metadataOffset+11]) << 8) +
				data.Metadata.convertRangeToNumber(metadataOffset+14, 2)
			subtypeCode = 0
		} else {
			charsetId = 63
			subtypeCode = uint64(data.Metadata[metadataOffset+14])
		}
		subType := constants.GetGeometryTypeFromCode(subtypeCode)

		fmt.Println(index, name, length, typeCode, fieldFlags, uniregCheck)
		metadataOffset += 17

		charset := constants.Lookup(charsetId)

		if labels != nil {
			if StringInSlice(charset.Name, []string{"ucs2", "utf16", "utf16le", "utf32"}) {
				// clear
				for index, _ := range labels {
					labels[index] = ""
				}
			}
		}

		comment := string(data.Comments.readData(0, commentLength))
		defaultDataOpts := getDefaultDataOptions{
			typeCode:    typeCode,
			flags:       fieldFlags,
			nullBit:     nullBit,
			nullBytes:   nullBytes,
			uniregCheck: uniregCheck,
		}
		defaultValue := getDefaultData(data.Defaults[defaultsOffset:], defaultDataOpts)

		column := Column{
			Name:         name,
			Length:       length,
			TypeCode:     typeCode,
			TypeName:     "",
			DefaultValue: defaultValue,
			Charset:      charset,
			Comment:      comment,
		}

		res = append(res, column)
	}

	return &res
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

type getDefaultDataOptions struct {
	typeCode    *constants.MySQLType
	flags       *[]constants.FieldFlag
	nullBit     int
	nullBytes   byteSlice
	uniregCheck *constants.UType
}

func getDefaultData(data byteSlice, options getDefaultDataOptions) string {
	hasDefaultValue := true
	isMayBeNull := false
	isDecimal := false
	for _, flag := range *options.flags {
		if flag.Name == "NO_DEFAULT" {
			hasDefaultValue = false
		}

		if flag.Name == "MAYBE_NULL" {
			isMayBeNull = true
		}

		if flag.Name == "DECIMAL" {
			isDecimal = true
		}
	}

	isAutoIncr := options.uniregCheck.Name == "NEXT_NUMBER"

	if isAutoIncr || !hasDefaultValue {
		return ""
	}

	if isMayBeNull {
		nullMap := options.nullBytes
		offset := options.nullBit / 8
		nullByte := nullMap[offset]
		nullBit := options.nullBit % 8
		if nullByte&(1<<nullBit) && options.uniregCheck.Name != "BLOB_FIELD" {
			return "NULL"
		}
	}

	if options.uniregCheck.Name == "BLOB_FIELD" {
		return ""
	}

	return constants.GetTypeDefault(data, isDecimal, options.typeCode)
}
