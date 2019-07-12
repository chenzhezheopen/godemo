package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println(div())
	for i := 0; i < 20; i++ {
		go fmt.Println(i)
	}
	p := temp{"chenzhe"}
	(&p).name = "zhe"
	fmt.Println("p=", p)

	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println("&arr=", &arr)
	var number = 1
	fmt.Println("number=", &number)
}

type temp struct {
	name string
}

func fun() {}
func div() (temp, error) {
	var hj = errors.New("zheshiyigeerr")
	var ji = temp{"ni"}
	return ji, hj
}
func (he *temp) use() string {
	return he.name
}
