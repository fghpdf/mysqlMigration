package constants

import "fghpdf.com/mysqlMigration/errors"

type GeometryType struct {
	Code uint64
	Name string
}

var codeToGeometryTypeMap = map[uint64]GeometryType{
	0: {
		Code: 0,
		Name: "GEOMETRY",
	},
	1: {
		Code: 1,
		Name: "POINT",
	},
	2: {
		Code: 2,
		Name: "LINESTRING",
	},
	3: {
		Code: 3,
		Name: "POLYGON",
	},
	4: {
		Code: 4,
		Name: "MULTIPOINT",
	},
	5: {
		Code: 5,
		Name: "MULTILINESTRING",
	},
	6: {
		Code: 6,
		Name: "MULTIPOLYGON",
	},
	7: {
		Code: 7,
		Name: "GEOMETRYCOLLECTION",
	},
}

func GetGeometryTypeFromCode(code uint64) *GeometryType {
	geo, ok := codeToGeometryTypeMap[code]
	if !ok {
		panic(errors.NewFormat("GeometryType code '%d' not found", code))
	}

	return &geo
}
