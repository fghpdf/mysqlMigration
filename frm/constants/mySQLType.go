package constants

type mySQLType struct {
	Code uint
	Name string
}

var (
	DECIMAL     = mySQLType{Code: 0, Name: "DECIMAL"}
	TINY        = mySQLType{Code: 1, Name: "TINY"}
	SHORT       = mySQLType{Code: 2, Name: "SHORT"}
	LONG        = mySQLType{Code: 3, Name: "LONG"}
	FLOAT       = mySQLType{Code: 4, Name: "FLOAT"}
	DOUBLE      = mySQLType{Code: 5, Name: "DOUBLE"}
	NULL        = mySQLType{Code: 6, Name: "NULL"}
	TIMESTAMP   = mySQLType{Code: 7, Name: "TIMESTAMP"}
	LONGLONG    = mySQLType{Code: 8, Name: "LONGLONG"}
	INT24       = mySQLType{Code: 9, Name: "INT24"}
	DATE        = mySQLType{Code: 10, Name: "DATE"}
	TIME        = mySQLType{Code: 11, Name: "TIME"}
	DATETIME    = mySQLType{Code: 12, Name: "DATETIME"}
	YEAR        = mySQLType{Code: 13, Name: "YEAR"}
	NEWDATE     = mySQLType{Code: 14, Name: "NEWDATE"}
	VARCHAR     = mySQLType{Code: 15, Name: "VARCHAR"}
	BIT         = mySQLType{Code: 16, Name: "BIT"}
	TIMESTAMP2  = mySQLType{Code: 17, Name: "TIMESTAMP2"}
	DATETIME2   = mySQLType{Code: 18, Name: "DATETIME2"}
	TIME2       = mySQLType{Code: 19, Name: "TIME2"}
	NEWDECIMAL  = mySQLType{Code: 246, Name: "NEWDECIMAL"}
	ENUM        = mySQLType{Code: 247, Name: "ENUM"}
	SET         = mySQLType{Code: 248, Name: "SET"}
	TINY_BLOB   = mySQLType{Code: 249, Name: "TINY_BLOB"}
	MEDIUM_BLOB = mySQLType{Code: 250, Name: "MEDIUM_BLOB"}
	LONG_BLOB   = mySQLType{Code: 251, Name: "LONG_BLOB"}
	BLOB        = mySQLType{Code: 252, Name: "BLOB"}
	VAR_STRING  = mySQLType{Code: 253, Name: "VAR_STRING"}
	STRING      = mySQLType{Code: 254, Name: "STRING"}
	GEOMETRY    = mySQLType{Code: 255, Name: "GEOMETRY"}
)

var codeToMySQLTypeMap = map[uint]mySQLType{
	0:   DECIMAL,
	1:   TINY,
	2:   SHORT,
	3:   LONG,
	4:   FLOAT,
	5:   DOUBLE,
	6:   NULL,
	7:   TIMESTAMP,
	8:   LONGLONG,
	9:   INT24,
	10:  DATE,
	11:  TIME,
	12:  DATETIME,
	13:  YEAR,
	14:  NEWDATE,
	15:  VARCHAR,
	16:  BIT,
	17:  TIMESTAMP2,
	18:  DATETIME2,
	19:  TIME2,
	246: NEWDECIMAL,
	247: ENUM,
	248: SET,
	249: TINY_BLOB,
	250: MEDIUM_BLOB,
	251: LONG_BLOB,
	252: BLOB,
	253: VAR_STRING,
	254: STRING,
	255: GEOMETRY,
}

func GetMySQLTypeFromCode(code uint) *mySQLType {
	mysqlType := codeToMySQLTypeMap[code]
	return &mysqlType
}
