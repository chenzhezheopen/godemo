package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	test()
	fmt.Println(123)
	for {
		time.Sleep(time.Second)
		runtime.GC()
	}
}
func test() {
	var c = make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			<-c
		}()
	}
}
