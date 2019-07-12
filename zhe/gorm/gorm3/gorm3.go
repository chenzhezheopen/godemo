package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Loggle *log.Logger

func main() {
	db, _ := gorm.Open("mysql", "root/root@/gorm?charset:utf8&&parseTime=true&&loc=local")
	fmt.Println(db)
	file, _ := os.OpenFile("test.txt", os.O_RDWR|os.O_APPEND, 0)
	Loggle = log.New(file, "\r\n", log.Ldate|log.Lshortfile)
	Loggle.Println("kkkkkkk")
}

// func intn() {
// 	var a string
// 	a = "123"
// }
