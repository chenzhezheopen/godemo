package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	em := make(chan int, 2)
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func(id int) {
			defer func() {
				fmt.Println(1)
				<-em
			}()
			defer func() {
				fmt.Println(3)
				wg.Done()
			}()
			fmt.Println(2)
			em <- 1,

			time.Sleep(time.Second * 2)
			fmt.Println(id, time.Now())
		}(i)
	}
	wg.Wait()
}
