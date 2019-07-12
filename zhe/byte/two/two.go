package main

import (
	"fmt"
	"time"
)

func read(ch1 chan int) {
	fmt.Println("123123")
	for range time.Tick(time.Second * 1) {
		fmt.Printf("read a int is %d\n", <-ch1)
	}
}

func write(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch1 := make(chan int)

	go write(ch1)
	go read(ch1)
	select {}
}
