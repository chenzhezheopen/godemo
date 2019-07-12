package main

import (
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	// gorm.Model
	User_id   int `gorm:"primary_key"` //指定主键并自增
	Name      string
	Pwd       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
type ret struct {
	User_id   int
	Name      string
	Pwd       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
type iii struct {
	Value        ret
	Erroe        string
	RowsAffected int
}

func main() {
	//创建连接
	db, err := gorm.Open("mysql", "root:root@/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()
	// 自动迁移模式
	//db.AutoMigrate(&User{})
	user := User{}
	// db.CreateTable(&User{})
	if !db.HasTable(&User{}) {
		if err := db.Set("gorm:table_options", "CHARSET=utf8").CreateTable(&User{}).Error; err != nil {
			panic(err)
		}
	}
	// 添加
	db.Create(&User{Name: "qwwqw", Pwd: "qwqwqwqwqw"})

	// 查询
	f := db.Where("user_id=3").First(&user)
	// fmt.Printf("%v", f)
	oi, _ := json.Marshal(f)
	var ret iii
	json.Unmarshal(oi, &ret)
	fmt.Println(ret.Value.CreatedAt)
	// he := json.Marshal(f)
	// var file User
	// json.Unmarshal(f, file)
	// fmt.Println(file)
	// 修改
	user.Name = "jinzhu 2"
	db.Save(&user)
	db.Model(&User{}).ModifyColumn("id=3", "aaaaaaaaaaaaa")
	// db.First(&user, 1)
	// 删除
	// db.Where("user_id = ?", 2).Delete(User{})
	// db.DropTable(&User{})

}
