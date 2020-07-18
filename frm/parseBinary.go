package frm

import (
	"fmt"
	"strconv"
)

func parse(fileData []byte) {
	versionIdData := fileData[51:55]

	mySQLVersion := getMySQLVersionFromByte(versionIdData)
	fmt.Println(mySQLVersion)
}

// MySQL version encoded as a 4-byte integer in little endian format.
// This is the value MYSQL_VERSION_ID from include/mysql_version.h in the mysql source tree.
// Example: ‘xb6xc5x00x00’ 0x0000c5b6 => 50614 => MySQL v5.6.14
func getMySQLVersionFromByte(versionIdData []byte) string {
	versionIdHex := fmt.Sprintf("%x%x%x%x", versionIdData[2], versionIdData[3], versionIdData[1], versionIdData[0])
	versionId, _ := strconv.ParseUint(versionIdHex, 16, 32)
	return GetMySQLVersionFromId(uint32(versionId)).Format()
}
