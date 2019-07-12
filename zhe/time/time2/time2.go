package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var time1 = time.Now()
	wg.Add(1)
	time.AfterFunc(5*time.Second, func() {
		fmt.Println("555555555555")
		wg.Done()
		var now = time.Now()
		fmt.Println(now.Sub(time1))
	})

	wg.Wait()
}
