package constants

type haOption uint

const (
	PACK_RECORD          = 1
	PACK_KEYS            = 2
	COMPRESS_RECORD      = 4
	LONG_BLOB_PTR        = 8
	TMP_TABLE            = 16
	CHECKSUM             = 32
	DELAY_KEY_WRITE      = 64
	NO_PACK_KEYS         = 128
	CREATE_FROM_ENGINE   = 256
	RELIES_ON_SQL_LAYER  = 512
	NULL_FIELDS          = 1024
	PAGE_CHECKSUM        = 2048
	STATS_PERSISTENT     = 4096
	NO_STATS_PERSISTENT  = 8192
	TEMP_COMPRESS_RECORD = 16384
	READ_ONLY_DATA       = 32768
)
