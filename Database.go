package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var urlDSN = "sql6523101:qpt2c1wn5i@tcp(sql6.freesqldatabase.com:3306)/sql6523101?parseTime=true"

//"root:<mypassword>@tcp(localhost:3306)/<dbName>"
var err error

func DataMigration() {
	Database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Connection failed!!")
	}
	log.Println("sucessfully connected to mysql server!!")
}
