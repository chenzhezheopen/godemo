package main

import (
	"fmt"
	"time"
)

func main() {
	// j := make(chan struct{})
	// go func() {
	// 	time.Sleep(time.Second)
	// 	j <- struct{}{}
	// }()
	// fmt.Println(123)
	// <-(chan struct{})(nil)

	// c := make(chan int, 10)
	// d := make(chan struct{})
	// go func() {
	// 	fmt.Println(123)
	// 	for x := range c {
	// 		fmt.Println(11)
	// 		fmt.Println(x)
	// 	}
	// 	close(d)
	// }()
	// for i := 0; i < 50; i++ {					//普通模式对数据打包
	// 	fmt.Println(22)
	// 	c <- i
	// }
	// fmt.Println(1)
	// close(c)
	// fmt.Println(123)
	// <-d

	d := make(chan struct{})
	c := make(chan [5]int, 10)
	go func() {
		count := 0
		for x := range c {
			time.Sleep(time.Second * 2)

			for _, a := range x {
				count += a
			}
			fmt.Println(count)
		}
		close(d)
	}()
	for i := 0; i < 50; i++ { //数组模式对数据打包
		fmt.Println(22)
		c <- [5]int{1, 2, 3, 4, 5}

	}
	close(c)
	fmt.Println(123)
	<-d
}
