package constants

import (
	"math"
)

type HaOption struct {
	Code uint64
	Name string
}

var codeToHaOptionMap = map[uint64]HaOption{
	1: {
		Code: 1,
		Name: "PACK_RECORD",
	},
	2: {
		Code: 2,
		Name: "PACK_KEYS",
	},
	4: {
		Code: 4,
		Name: "COMPRESS_RECORD",
	},
	8: {
		Code: 8,
		Name: "LONG_BLOB_PTR",
	},
	16: {
		Code: 16,
		Name: "TMP_TABLE",
	},
	32: {
		Code: 32,
		Name: "CHECKSUM",
	},
	64: {
		Code: 64,
		Name: "DELAY_KEY_WRITE",
	},
	128: {
		Code: 128,
		Name: "NO_PACK_KEYS",
	},
	256: {
		Code: 256,
		Name: "CREATE_FROM_ENGINE",
	},
	512: {
		Code: 512,
		Name: "RELIES_ON_SQL_LAYER",
	},
	1024: {
		Code: 1024,
		Name: "NULL_FIELDS",
	},
	2048: {
		Code: 2048,
		Name: "PAGE_CHECKSUM",
	},
	4096: {
		Code: 4096,
		Name: "STATS_PERSISTENT",
	},
	8192: {
		Code: 8192,
		Name: "NO_STATS_PERSISTENT",
	},
	16384: {
		Code: 16384,
		Name: "TEMP_COMPRESS_RECORD",
	},
	32768: {
		Code: 32768,
		Name: "READ_ONLY_DATA",
	},
}

func GetHaOptionsFromCode(code uint64) *[]HaOption {
	codes := *sliceCode(code)

	var options []HaOption

	for _, value := range codes {
		options = append(options, codeToHaOptionMap[value])
	}

	return &options
}

func sliceCode(code uint64) *[]uint64 {
	var res []uint64
	remain := code

	for i := 16; remain >= 0 && i >= 0; i-- {
		pow := uint64(math.Pow(2, float64(i)))
		if remain >= pow {
			res = append(res, pow)
			remain -= pow
		}
	}

	return &res
}
