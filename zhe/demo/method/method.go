package main

import "fmt"

type N int

func (n N) test() {
	fmt.Printf("test=%p,%d\n", &n, n)
}
func main() {
	var n N = 25
	fmt.Printf("main=%p,%d\n", &n, n)
	f1 := N.test
	f1(n)
	n.test()
	n.test()

	f2 := (*N).test
	f2(&n)
}
