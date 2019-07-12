package main

import "fmt"

func New(text string) temp {
	return &errorString{text}
}

type errorString struct {
	s string
}

type temp interface {
	Error()
}

func (e *errorString) Error() string {
	fmt.Println("Error")
	return e.s
}

func main() {
	fmt.Println(New("nihao"))
	arr := errorString{"hello"}
	fmt.Println(arr.Error())
}
