package main

import (
	"fmt"
	"time"
)

type us struct {
	name string
	sex  string
	age  int
}

func main() {
	// var n us
	// l := reflect.TypeOf(n)
	// fmt.Println(l)
	// for i := 0; i < l.NumField(); i++ {
	// 	fmt.Println(l.Field(i).Name)
	// }
	for range time.Tick(2 * time.Second) {
		fmt.Println("aaaaaa")
	}

}
