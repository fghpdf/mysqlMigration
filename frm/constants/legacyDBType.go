package constants

type legacyDBType struct {
	Code uint
	Name string
}

var (
	UNKNOWN_LEGACY_DB_TYPE = legacyDBType{Code: 0, Name: "UNKNOWN"}
	DIAB_ISAM              = legacyDBType{Code: 1, Name: "DIAB_ISAM"}
	HASH                   = legacyDBType{Code: 2, Name: "HASH"}
	MISAM                  = legacyDBType{Code: 3, Name: "MISAM"}
	PISAM                  = legacyDBType{Code: 4, Name: "PISAM"}
	RMS_ISAM               = legacyDBType{Code: 5, Name: "RMS_ISAM"}
	HEAP                   = legacyDBType{Code: 6, Name: "HEAP"}
	ISAM                   = legacyDBType{Code: 7, Name: "ISAM"}
	MRG_ISAM               = legacyDBType{Code: 8, Name: "MRG_ISAM"}
	MyISAM                 = legacyDBType{Code: 9, Name: "MyISAM"}
	MRG_MYISAM             = legacyDBType{Code: 10, Name: "MRG_MYISAM"}
	BERKELEYDB             = legacyDBType{Code: 11, Name: "BERKELEYDB"}
	InnoDB                 = legacyDBType{Code: 12, Name: "InnoDB"}
	GEMINI                 = legacyDBType{Code: 13, Name: "GEMINI"}
	NDBCLUSTER             = legacyDBType{Code: 14, Name: "NDBCLUSTER"}
	EXAMPLE_DB             = legacyDBType{Code: 15, Name: "EXAMPLE_DB"}
	ARCHIVE_DB             = legacyDBType{Code: 16, Name: "ARCHIVE_DB"}
	CSV                    = legacyDBType{Code: 17, Name: "CSV"}
	FEDERATED              = legacyDBType{Code: 18, Name: "FEDERATED"}
	BLACKHOLE              = legacyDBType{Code: 19, Name: "BLACKHOLE"}
	PARTITION_DB           = legacyDBType{Code: 20, Name: "PARTITION_DB"}
	BINLOG                 = legacyDBType{Code: 21, Name: "BINLOG"}
	SOLID                  = legacyDBType{Code: 22, Name: "SOLID"}
	PBXT                   = legacyDBType{Code: 23, Name: "PBXT"}
	TABLE_FUNCTION         = legacyDBType{Code: 24, Name: "TABLE_FUNCTION"}
	MEMCACHE               = legacyDBType{Code: 25, Name: "MEMCACHE"}
	FALCON                 = legacyDBType{Code: 26, Name: "FALCON"}
	MARIA                  = legacyDBType{Code: 27, Name: "MARIA"}
	PERFORMANCE_SCHEMA     = legacyDBType{Code: 28, Name: "PERFORMANCE_SCHEMA"}
	FIRST_DYNAMIC          = legacyDBType{Code: 42, Name: "FIRST_DYNAMIC"}
	DEFAULT_LEGACY_DB_TYPE = legacyDBType{Code: 127, Name: "DEFAULT"}
)

var codeToLegacyDBTypeMap = map[uint]legacyDBType{
	0:  UNKNOWN_LEGACY_DB_TYPE,
	1:  DIAB_ISAM,
	2:  HASH,
	3:  MISAM,
	4:  PISAM,
	5:  RMS_ISAM,
	6:  HEAP,
	7:  ISAM,
	8:  MRG_ISAM,
	9:  MyISAM,
	10: MRG_MYISAM,
	11: BERKELEYDB,
	12: InnoDB,
	13: GEMINI,
	14: NDBCLUSTER,
	15: EXAMPLE_DB,
	16: ARCHIVE_DB,
	17: CSV,
	18: FEDERATED,
	19: BLACKHOLE,
	20: PARTITION_DB,
	21: BINLOG,
	22: SOLID,
	23: PBXT,
	24: TABLE_FUNCTION,
	25: MEMCACHE,
	26: FALCON,
	27: MARIA,
	28: PERFORMANCE_SCHEMA,
	29: FIRST_DYNAMIC,
	30: DEFAULT_LEGACY_DB_TYPE,
}

func GetLegacyDBTypeFromCode(code uint) *legacyDBType {
	dbType, ok := codeToLegacyDBTypeMap[code]
	if !ok {
		return &UNKNOWN_LEGACY_DB_TYPE
	}

	return &dbType
}
