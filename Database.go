package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var urlDSN = "root:root@tcp(localhost:3306)/mydb?parseTime=true"

//"root:<mypassword>@tcp(localhost:3306)/<dbName>"
var err error

func DataMigration() {
	Database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Connection failed!!")
	}
	log.Println("sucessfully connected to mysql server!!")

	// Create necessary tables
	err = CreateTables()
	if err!=nil{
		log.Fatal(err)
		panic("Failed to create myDB Tables!!")
	}
}

func CreateTables() error{
	err := Database.Raw(CreateEmployeeTable)
	log.Println(CreateEmployeeTable)
	if err.Error != nil {
		log.Println(err.Error)
		return err.Error
	}
	return nil
}