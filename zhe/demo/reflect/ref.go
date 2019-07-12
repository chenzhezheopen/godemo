package main

import (
	"fmt"
	"reflect"
)

type use struct {
	name string
	sex  string
	age  int
}

func (arr use) namee() {
	fmt.Println("123")
	fmt.Println(arr.name)
}
func main() {
	arr := use{"你", "是", 55}
	hee := use{"你sa", "是", 54}
	arr.namee()
	hee.namee()
	info(arr)
}
func info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("typeOf:", t)
}
