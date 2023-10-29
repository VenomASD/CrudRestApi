package main

import (
	"fmt"
	"log"
	"reflect"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

// var urlDSN = "root:root@tcp(localhost:3306)/mydb?parseTime=true"

// "root:<mypassword>@tcp(localhost:3306)/<dbName>"
var err error

func DataMigration(databaseUrl string) {
	Database, err = gorm.Open(mysql.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Connection failed!!")
	}
	log.Println("sucessfully connected to mysql server!!")

	// Create necessary tables
	Migration(Actor{})
}

func Migration(tableSchema interface{}) {
	logEntry := fmt.Sprintf("Auto Migrating %s...", reflect.TypeOf(tableSchema))
	// Create Table in SQL DB corresponding to schema
	db := Database.AutoMigrate(tableSchema)
	if db != nil {
		//We have an error
		log.Fatal(fmt.Sprintf("%s %s with error %s", logEntry, "Failed", db.Error()))
	}
	log.Println(fmt.Sprintf("%s %s", logEntry, "Success"))
}
