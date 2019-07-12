package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(5)
	fmt.Println(runtime.GOROOT())
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.NumCPU())
}
