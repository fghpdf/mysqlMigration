package frm

import (
	"io/ioutil"
	"path"
	"strings"
)
import "fghpdf.com/mysqlMigration/errors"

func ParsePath(filePath string) error {
	// check path
	if path.Ext(filePath) != ".frm" {
		return errors.FILE_TYPE_INVALID.NewFormat("file type must be .frm but get %s", path.Ext(filePath))
	}

	// open file
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return errors.WrapFormat(err, "open file err")
	}

	// read magic section
	magic := string(fileData[0:2])
	if strings.HasPrefix(magic, "\xfe\x01") {
		parse(fileData[:])
	}
	return nil
}
