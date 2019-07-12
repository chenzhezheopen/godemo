package main

import (
	"fmt"
)

func main() {
	// var wg sync.WaitGroup
	// wg.Add(1)
	done := make(chan int)
	// don := make(chan bool, 1)
	go func() {
		// time.Sleep(time.Second * 3)
		// for i := 0; i < 4; i++ {

		// 	if _, ok := <-done; !ok {
		// 		return
		// 	}
		<-done
		fmt.Println("lllllllllllllllllllllllllllllllllllllllllllllllllll")
		done <- 1
		// }

		// don <- true
		// defer wg.Done()
	}()
	// fmt.Println("222222222")
	// for i := 0; i < 3; i++ {
	done <- 1
	fmt.Println(123)
	<-done
	// }
	// <-don
	// fmt.Println("lllll")
	// fmt.Println("lllll")
	// fmt.Println("lllll")
	// fmt.Println("lllll")
	// fmt.Println("lllll")
	// fmt.Println("lllll")
	// fmt.Println("lllll")
	// fmt.Println("lllll")
	// fmt.Println("lllll")
	// wg.Wait()
}
