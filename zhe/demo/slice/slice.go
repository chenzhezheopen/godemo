package main

import "fmt"

func main() {
	a := make([]string, 5, 6)
	a[1] = "123"
	fmt.Println(a[1])
}
