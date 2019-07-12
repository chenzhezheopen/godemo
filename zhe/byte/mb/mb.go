package main

import (
	"fmt"
)

// 420612182106225
func main() {
	// var m string
	// fmt.Scan(&m)
	// fmt.Println(m)
	// rand.Seed(time.Now().Unix())
	// number := rand.Intn(1200)
	// fmt.Println(number)
	data := new([10][10 * 2]byte)
	for i := range data {
		for x := 0; x < len(data[i]); x++ {

			data[i][x] = 10
		}
	}
	fmt.Println(data)
}
