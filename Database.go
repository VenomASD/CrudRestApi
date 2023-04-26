package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var urlDSN = "root:root@tcp(127.0.0.1:3307)/mydb?parseTime=true"

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
