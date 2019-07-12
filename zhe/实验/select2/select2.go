package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c1 := make(chan int, 2)
	c1 <- 1
	c1 <- 2

	rand.Seed(time.Now().Unix())
	for range time.Tick(time.Second * 1) {
		select {
		case c1 <- rand.Intn(100):
			fmt.Println("写入")
		case x, _ := <-c1:
			fmt.Println(x)
		default:
			fmt.Println("default")
		}
	}
}
