package frm

import "fmt"

type MySQLVersion struct {
	Major   uint64
	Minor   uint64
	Release uint64
}

func (version MySQLVersion) Format() string {
	if version.Major == 0 && version.Minor == 0 && version.Release == 0 {
		return "< 5.0"
	} else {
		return fmt.Sprintf("%d.%d.%d", version.Major, version.Minor, version.Release)
	}
}

func GetMySQLVersionFromId(versionId uint64) MySQLVersion {
	return MySQLVersion{
		Major:   versionId / 10000,
		Minor:   versionId % 1000 / 100,
		Release: versionId % 100,
	}
}
