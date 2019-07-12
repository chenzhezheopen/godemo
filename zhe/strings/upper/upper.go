package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"
)

const (
	Ldate         = 1 << iota     //日期示例： 2009/01/23
	Ltime                         //时间示例: 01:23:23
	Lmicroseconds                 //毫秒示例: 01:23:23.123123.
	Llongfile                     //绝对路径和行号: /a/b/c/d.go:23
	Lshortfile                    //文件和行号: d.go:23.
	LUTC                          //日期时间转为0时区的
	LstdFlags     = Ldate | Ltime //Go提供的标准抬头信息
)

func main() {
	var king = "nihao"
	fmt.Println(strings.ToUpper(king)) //NIHAO
	var kind = "NIHAO"
	fmt.Println(strings.ToLower(kind)) //nihao
	var title = "niHao"
	fmt.Println(strings.ToTitle(title))   //nihao
	fmt.Println(strings.Trim(title, "n")) //iHao
	fmt.Println(strings.Trim(title, "o")) //niHa
	fmt.Println(strings.Trim(title, "H")) //niHao

	fmt.Println("Split", strings.Split(title, "")) //[n i H a o]

	var join = []string{"1", "2", "3", "4", "5", "6"}
	var jjjj = [4]int{1, 2, 3, 4}
	fmt.Println("Join", strings.Join(join, "")) //123456
	fmt.Println(reflect.TypeOf(join).Kind())
	fmt.Println(reflect.TypeOf(jjjj).Kind())

	log.SetFlags(log.Ldate | log.Lshortfile)
	s := strings.NewReader("hello world!")
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))

	var t try

	fmt.Println(t)
	fmt.Println(reflect.TypeOf(t).Kind())

	var d data = 5
	fmt.Println(reflect.TypeOf(d))
	var x interface{} = d
	fmt.Println("d", reflect.TypeOf(x).Name())

	if ww, a := x.(fmt.Stringer); a {
		fmt.Println("123", reflect.TypeOf(ww))
	}

	n, _ := x.(fmt.Stringer)
	fmt.Println(reflect.TypeOf(n))

	var jong data = 3
	var xing interface{} = jong
	fmt.Println("jong.name()", jong.name())
	fmt.Println("1", reflect.TypeOf(xing))
	pppp, _ := xing.(uuuu)
	fmt.Println(reflect.TypeOf(pppp))

	t5 := time.Now()
	fmt.Println(t5.Day())
	fmt.Println(t5.Weekday())
	fmt.Println(t5.Month())
}

type uuuu interface {
	name() string
}

func (data) name() string {
	return "456"
}

type try func() string
type data int

func (try) String() string {
	return "heloi"
}
