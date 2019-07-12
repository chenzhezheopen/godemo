package main

import (
	"fmt"
)

func main() {
	done := make(chan bool)
	data := make(chan int)

	go producer(data)
	go consumer(data, done)

	<-done
}
func consumer(data chan int, done chan bool) {
	for x := range data {
		fmt.Println("recv:", x)
	}
	done <- true
}
func producer(data chan int) {
	for i := 0; i < 4; i++ {
		data <- i
	}
	close(data)
}
