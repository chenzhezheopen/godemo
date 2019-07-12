package main

import "fmt"

type mm struct {
	sss string
}

type ma struct {
	sex string
	age int //map[string]session
}
type inter interface {
	name()
}

func main() {
	var ma ma
	ma.name()
	var mm mm
	mm.name()

}
func (*ma) name() {
	fmt.Println("ma")
}

func (*mm) name() {
	fmt.Println("mm")
}

// type session struct {
// 	local string
// 	id    int
// }
