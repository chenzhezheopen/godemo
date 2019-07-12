package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Like struct {
	ID        int    `gorm:"primary_key"`
	Ip        string `gorm:"type:varchar(20);not null;index:ip_idx"`
	Ua        string `gorm:"type:varchar(256);not null;"`
	Title     string `gorm:"type:varchar(128);not null;index:title_idx"`
	Hash      int    `gorm:"unique_index:hash_idx;"`
	CreatedAt time.Time
}

func main() {
	db, err := gorm.Open("mysql", "root:root@/jubao?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("数据库连接失败")
	}

	// ket := temp{id: 1, name: "chen"}

	if !db.HasTable(&Like{}) {
		if err := db.Set("gorm:table_options", "CHARSET=utf8").CreateTable(&Like{}).Error; err != nil {
			panic(err)
		}
	}
	// currentTime := time.Now()
	// rand.Seed(time.Now().Unix())
	// fmt.Println(rand.Intn(1000))
	// // rand := rand.Seed
	// like := &Like{
	// 	Ip:        "dd",
	// 	Ua:        "aa",
	// 	Title:     "bb",
	// 	Hash:      rand.Intn(1000),
	// 	CreatedAt: currentTime,
	// }
	// // db.Where(&Like{Hash: 122}).Delete(Like{})
	// db.Create(like)

	// // fmt.Println(db.Model(&Like{}).Where(&Like{Hash: 12}).Count(&like))

	// fmt.Println(*db.First(like, 1))
	// fmt.Println(os.Getwd())
	// fmt.Println(os.Getenv("GOPATH"))
	// if err := db.Create(like).Error; err != nil {
	// 	fmt.Println("chenggong")
	// }
	// db.DropTable("likes")
}
