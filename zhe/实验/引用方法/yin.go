package main

import "fmt"

type user struct {
	name string
	sex  string
	age  int
}
type user1 struct {
	ty user
}

func (us *user) assist() {
	fmt.Println("user")
}
func (uu *user1) assist() {
	fmt.Println("user1")
}
func main() {
	var man user1
	man.assist()
	fmt.Println(this)
}
