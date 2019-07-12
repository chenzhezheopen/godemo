package main

import "fmt"

func main() {
	var str string = "test"

	var data []byte = []byte(str)
	fmt.Println(data)
	fmt.Println(string(data[:]))
}
