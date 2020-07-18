package main

import (
	"fghpdf.com/mysqlMigration/frm"
	"log"
	"path"
)

func main() {
	filePath := path.Join("c:", "workspace", "data_622", "data", "pics", "templates.frm")

	err := frm.ParsePath(filePath)
	if err != nil {
		log.Fatalln(err)
	}
}
