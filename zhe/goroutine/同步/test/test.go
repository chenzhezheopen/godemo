package main

import (
	"fmt"
	"zhe/goroutine/同步/lib"
)

func main() {

	d := lib.Ling()
	d.Name = "123"
	fmt.Println(d)
	k := lib.Data{}
	k.Name = "123"
	fmt.Println(k)
}
