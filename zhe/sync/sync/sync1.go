package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func main() {

	for i, _ := range make([]string, 10) {
		once.Do(onces)
		fmt.Println(i)
	}

	once.Do(onced)

}
func onces() {
	fmt.Println("onces")
}
func onced() {
	fmt.Println("onced")
}
