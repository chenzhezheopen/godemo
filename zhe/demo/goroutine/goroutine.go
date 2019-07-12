package main

import (
	"fmt"
	"time"
)

var c int

func counter(a int) int {
	fmt.Println(a, c)
	c++
	return c
}
func main() {
	a := 100

	go func(y int) {
		time.Sleep(time.Second)
		println("go:1", a, y)
	}(counter(2))

	go func(y int) {
		time.Sleep(time.Second)
		println("go:2", a, y)
	}(counter(3))

	go func(y int) {
		time.Sleep(time.Second)
		println("go:3", a, y)
	}(counter(4))

	time.Sleep(1000)
	println("main:")
	a += 100
	println("main:", a, counter(7))
	time.Sleep(time.Second * 10)
}
