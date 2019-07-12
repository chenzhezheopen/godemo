package main

import (
	"fmt"
)

func main() {
	a := make(map[int]map[string]string, 5)
	a[0] = make(map[string]string)
	a[0]["name"] = "hello"
	fmt.Println(len(a))

	aa := make([]int, 0, 5)
	n := len(aa)
	fmt.Println("nn=", n)
	if n == cap(aa) {
		return
	}
	aa = aa[:n+1]

	aa[n] = 11111
	fmt.Println(cap(aa))
	fmt.Println("aa=", aa)
	fmt.Println("n=", len(aa))
}
