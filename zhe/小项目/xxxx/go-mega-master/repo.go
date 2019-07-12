package main

import "fmt"

func main() {

	c := new([10]*[1024 * 1024]byte)
	for i := range c {
		for j := 0; i < len(c[i]); j++ {
			c[i][j] = 4
		}
	}
	fmt.Println(c)
}
