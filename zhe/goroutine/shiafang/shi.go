package main

import (
	"fmt"
)

func main() {
	data := make(chan int, 2)
	bol := make(chan bool)
	data <- 1
	data <- 2
	close(data)
	go func() {
		if i, ok := <-data; ok {
			fmt.Printf("%d %v\n", i, <-data)
		} else {
			bol <- true
		}
	}()
	// go func() {
	// 	for range time.Tick(time.Second * 1) {
	// 		fmt.Printf("123")
	// 	}
	// }()
	<-bol
	select {}
}
