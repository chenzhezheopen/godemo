package main

import "fmt"

func main() {

	// ticker := time.NewTicker(500 * time.Millisecond)
	// go func() {
	// 	for t := range ticker.C {
	// 		fmt.Println("Tick at", t)
	// 	}
	// }()

	// time.Sleep(3100 * time.Millisecond)
	// ticker.Stop()
	// fmt.Println("Ticker stopped")
	b := make(chan bool)
	go func() {
		b <- true
	}()
	fmt.Println("123")
	<-b
	fmt.Println("123")
}
