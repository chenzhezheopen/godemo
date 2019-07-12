package main

import "fmt"

func mapp() []int {
	m := make([]int, 5, 10)
	m[0] = 1
	println("2", m)
	fmt.Println("1", &m)

	return m
}

func main() {
	m := mapp()
	fmt.Println(m)

	fmt.Println("fmt.Println")
	println("println")
	println(m)
	x := 100
	fmt.Println("q=", &x)
	q := (*int)(&x)
	fmt.Println("q=", q)

}
