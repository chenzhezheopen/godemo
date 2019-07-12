package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Us struct {
	UsId      int `gorm:"primary_key"`
	profile   Profile
	ProfileID int
}
type Profile struct {
	gorm.Model
	UserID uint
	Name   string
}

func main() {
	db, err := gorm.Open("mysql", "root:root@/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("数据库连接失败")
	}
	db.CreateTable(&Us{})
	db.CreateTable(&Profile{})

}
