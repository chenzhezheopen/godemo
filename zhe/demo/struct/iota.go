package main

import "fmt"

// type flags byte

// const (
// 	read  flags = 7
// 	write       = 3 << iota
// 	exec
// 	down
// 	up
// )

// func main() {
// 	f := read | exec
// 	fmt.Println(read)
// 	fmt.Println(write)
// 	fmt.Println(exec)
// 	fmt.Println(down)
// 	fmt.Println(up)
// 	fmt.Println("%b\n", f)
// }

type us struct {
	name string
}

func (us) kind() {
	fmt.Println("aaaa")
}
func main() {
	var k = us{"she"}
	us.kind(k)
}
