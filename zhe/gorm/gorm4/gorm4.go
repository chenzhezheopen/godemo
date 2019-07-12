package main

import (
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Title struct {
	Id       int       `gorm:"primary_key"`
	ToDate   time.Time `gorm:"column:to_date" json:"to_date"`
	EmpNo    int       `gorm:"column:emp_no" json:"emp_no"`
	Title    string    `gorm:"column:title" json:"title"`
	FromDate time.Time `gorm:"column:from_date" json:"from_date"`
}
type total struct {
	Value        City
	RowsAffected int
}

// type err struct {
// 	Number  int
// 	Message string
// }
type City struct {
	Id        int `gorm:"primary_key"`
	Name      string
	Pid       int
	Is_delete int
}
type Citytwo struct {
	Id        int `gorm:"primary_key"`
	Name      string
	Pid       int
	Is_delete int
}
type likes struct {
	id int `gorm:"primary_key"`
}

var like likes
var country total
var city City
var title Title
var citytwo Citytwo

func main() {

	DB, err := gorm.Open("mysql", "root:root@/jubao?charset=utf8&parseTime=True&loc=Local")
	//    gorm.Open("mysql", "root:root@/gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("zheshiyigecuowuwuuuuuuuuuuuuuuuuuuuuuuuuuuuu")
	}
	DB.SingularTable(true)
	// DB.CreateTable(&citythree)

	// queryStr := `select name from  city`

	// rows, _ := DB.Self.Raw(queryStr).Rows() // (*sql.Rows, error)
	// defer rows.Close()

	// d := &DataTest{}
	// for rows.Next() {
	// 	rows.Scan(&d.Data) //这样是无法得到结果的，不知道应该怎么操作。
	// }
	// number := DB.Where("Id=1").Find(&City{})
	number := DB.Select("id,name,pid").Table("city").Where("id=55").Scan(&city)
	// number := DB.Where("id=5").Find(&city)
	number1, _ := json.Marshal(number)
	json.Unmarshal(number1, city)
	// vim := string(number1)
	fmt.Println(city)
	// DB.Updata("citytwo").Set("pid=100").Where("id=4").Scan(&citytwo)
}
