package main

import "fmt"

type user struct {
	name string
}
type inter interface {
	string()
	test()
}

func (user) string() {
	fmt.Println("string")
}
func (user) test() {
	fmt.Println("test")
}
func main() {
	var hello user
	hello.test()
	var ggg inter = hello
	ggg.string()
}
