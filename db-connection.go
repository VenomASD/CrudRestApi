package main

import (
		"fmt"
		"gorm.io/driver/mysql"
		"gorm.io/gorm")

var Database *gorm.DB
var urlDSN = "root:root@tcp(localhost:3306)/godb?parseTime=true"
var err error

func DataMigration() {
	Database, err = gorm.Open(mysql.Open(urlDSN) , &gorm.Config{})
	//db, err := sql.Open("mysql", "root:mypassword@tcp(127.0.0.1:3306)/test")
	if err!=nil{
		fmt.Print(err.Error())
		panic("Connection failed!!")
	}
	Database.AutoMigrate(&Employee{})
}