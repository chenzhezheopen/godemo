package main

import (
	"fmt"
	"time"
)

func main() {

	// strs := []string{"c", "a", "b"}
	// sort.Strings(strs)
	// fmt.Println("Strings:", strs)

	// ints := []int{7, 2, 4}
	// sort.Ints(ints)
	// fmt.Println("Ints:   ", ints)

	// s := sort.IntsAreSorted(ints)
	// fmt.Println("Sorted: ", s)
	c := make(chan int, 1)
	c <- 1
	go func() {
		d := <-c
		fmt.Println(d)
	}()
	time.Sleep(1000)
}
