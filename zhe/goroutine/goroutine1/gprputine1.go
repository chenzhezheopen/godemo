package main

import "fmt"
import "time"
import "sync/atomic"

func main() {

	var ops uint64

	for i := 0; i < 50; i++ {
		go func() {
			for {
				// fmt.Println(i)
				// atomic.AddUint64(&ops, 1000)

				time.Sleep(time.Millisecond * 500)
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
