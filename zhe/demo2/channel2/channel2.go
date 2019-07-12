package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)
	c := make(chan int)
	var send chan<- int = c
	var recv <-chan int = c
	go func() {
		defer wg.Done()

		defer close(c)
		for x := range recv {
			println(x)
		}
	}()
	go func() {
		defer wg.Done()
		fmt.Println(1233)

		for i := 0; i < 6; i++ {

			send <- i
			fmt.Println(123)
		}
	}()
	wg.Wait()
}
