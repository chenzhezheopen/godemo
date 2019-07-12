package main

import "fmt"

func main() {
	fmt.Println(New("nihao"))
	jjjj := New("nihao")
	fmt.Println("lllllllllllllll")
	fmt.Println(jjjj)
}

type errorString struct {
	s string
}
type error interface {
	String() string
}

func (e *errorString) String() string {
	fmt.Println("String...")
	return "OKOK"
}

func New(text string) error {
	fmt.Println("eeeeeeeeeee")
	return &errorString{text}
}

func (e *errorString) Error() string {
	fmt.Println("Error")
	return e.s
}
