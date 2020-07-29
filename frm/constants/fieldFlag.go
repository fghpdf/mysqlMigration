package constants

type FieldFlag struct {
	Code uint64
	Name string
}

var stringToFieldFlagMap = map[string]FieldFlag{
	"DECIMAL": {
		Code: 1,
		Name: "DECIMAL",
	},
	"BINARY": {
		Code: 1,
		Name: "BINARY",
	},
	"NUMBER": {
		Code: 2,
		Name: "NUMBER",
	},
	"ZEROFILL": {
		Code: 4,
		Name: "ZEROFILL",
	},
	"PACK": {
		Code: 120,
		Name: "PACK",
	},
	"INTERVAL": {
		Code: 256,
		Name: "INTERVAL",
	},
	"BITFIELD": {
		Code: 512,
		Name: "BITFIELD",
	},
	"BLOB": {
		Code: 1024,
		Name: "BLOB",
	},
	"GEOM": {
		Code: 2048,
		Name: "GEOM",
	},
	"TREAT_BIT_AS_CHAR": {
		Code: 4096,
		Name: "TREAT_BIT_AS_CHAR",
	},
	"NO_DEFAULT": {
		Code: 16384,
		Name: "NO_DEFAULT",
	},
	"MAYBE_NULL": {
		Code: 32768,
		Name: "MAYBE_NULL",
	},
	"HEX_ESCAPE": {
		Code: 65536,
		Name: "HEX_ESCAPE",
	},
	"PACK_SHIFT": {
		Code: 3,
		Name: "PACK_SHIFT",
	},
	"DEC_SHIFT": {
		Code: 8,
		Name: "DEC_SHIFT",
	},
	"MAX_DEC": {
		Code: 31,
		Name: "MAX_DEC",
	},
	"NUM_SCREEN_TYPE": {
		Code: 32513,
		Name: "NUM_SCREEN_TYPE",
	},
	"ALFA_SCREEN_TYPE": {
		Code: 30720,
		Name: "ALFA_SCREEN_TYPE",
	},
}

func GetFieldFlagFromCode(code uint64) *[]FieldFlag {
	var res []FieldFlag
	for _, field := range stringToFieldFlagMap {
		if code&field.Code != 0 {
			res = append(res, field)
		}
	}

	return &res
}
