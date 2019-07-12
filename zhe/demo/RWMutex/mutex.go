package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.RWMutex

func main() {
	done := make(chan bool)

	m := make(map[string]int)
	go mut(1, m, done)
	go mut(2, m, done)
	// go mut(3, m, done)
	select {}
}
func mut(i int, m map[string]int, done chan bool) {
	for range time.Tick(100 * time.Millisecond) {
		lock.Lock()
		m["a"] += 1
		fmt.Println(" map:", i, m)
		lock.Unlock()
	}
}
