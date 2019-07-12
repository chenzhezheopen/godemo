package main

import (
	"errors"
	"fmt"
)

func main() {
	stack := make([]int, 0, 10)
	push := func(x int) error {
		n := len(stack)
		if n == cap(stack) {
			return errors.New("stack is full")
		}
		stack = stack[:n+1]
		stack[n] = x
		return nil
	}
	for i := 0; i < 7; i++ {
		fmt.Println("push: %v, %v\n", push(i), stack)
	}
	fmt.Println(stack[3:9])
}
