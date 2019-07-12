package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	data := make(chan bool, 200)
	for i := 0; i < 201000; i++ {
		select {
		case data <- true:

			go yunx(i, ch)
		case <-ch:
			// time.Sleep(time.Second * 1)
			<-data
		}
	}
}
func yunx(i int, ch chan<- int) {
	time.Sleep(time.Second * 1)
	fmt.Println(i)
	ch <- i
}
