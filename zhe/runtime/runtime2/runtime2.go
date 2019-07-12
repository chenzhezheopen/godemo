package main

import (
	"fmt"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(1)
}

func say(s string) {
	for i := 0; i < 2; i++ {
		fmt.Println(s)
		runtime.Gosched()
		fmt.Println(i)
		fmt.Println(s)
	}
}
func main() {
	go say("world")
	say("hello")
}
