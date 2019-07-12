package main

import (
	"fmt"
	"runtime"
)

func main() {
	defer runtime.Goexit()
	go func() {
		fmt.Println("lllll")
	}()

	go func() {
		fmt.Println("222222")
	}()
	go func() {
		fmt.Println("33333333")
	}()
	go func() {
		fmt.Println("4444444")
	}()

}
