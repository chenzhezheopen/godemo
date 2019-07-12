package main

import (
	"fmt"
	"reflect"
)

type yes struct {
	name string
	sex  byte
}

func main() {
	u := yes{"Tom", 1}
	v := reflect.ValueOf(u)
	t := v.Type()
	for i, n := 0, t.NumField(); i < n; i++ {
		fmt.Printf("%v\n", t.Field(i))
	}
}
