package main

import (
	"sync"
	"time"
)

type temp struct {
	sync.Mutex
}

func (d temp) num(s string) {
	d.Lock()
	defer d.Unlock()
	for i := 0; i < 10; i++ {
		println(s, i)
		time.Sleep(time.Second)
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	var kind temp
	go func() {
		defer wg.Done()
		kind.num("第一个线程")
	}()
	go func() {
		defer wg.Done()
		kind.num("第二个线程")
	}()
	wg.Wait()
}
