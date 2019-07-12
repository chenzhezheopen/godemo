package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	a, b := make(chan int), make(chan int)
	wg.Add(2)
	// time.Sleep(time.Second * 5)
	go func() {
		defer wg.Done()
		for {
			var (
				name string
				x    int
				ok   bool
			)
			select {
			case x, ok = <-a:
				name = "a"
				a = nil
			case x, ok = <-b:
				name = "b"
			}
			if !ok {
				return
			}
			fmt.Println(name, x)
		}
	}()
	go func() {
		defer wg.Done()
		defer close(a)
		defer close(b)
		for i := 0; i < 10; i++ {
			select {
			case a <- i:
			case b <- i:
			}
		}
	}()
	wg.Wait()
}
