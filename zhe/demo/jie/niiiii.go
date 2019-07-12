package main

import "fmt"

func main() {
	fmt.Println(Boss("123"))
}

type hello interface {
	User() string
	String() string //Error>String>User
	Error() string
}

type temp struct {
	price string
}

func (temp *temp) User() string {
	return "BBBB"
}
func (temp *temp) Error() string {
	return "cccc"
}

func (temp *temp) String() string {
	return "aaaa"
}
func Boss(text string) hello {
	return &temp{text}
}
