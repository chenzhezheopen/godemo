package main

import "fmt"

type a struct {
	name string
	age  int
}
type b struct {
	name string
	age  int
}

func main() {
	type qq a
	cc := a{"ni", 1}
	mm := qq{"ni", 1}
	nn := b{"ni", 1}
	fmt.Println(mm)
	fmt.Println(nn)
	fmt.Println("%T\n", cc)
}
