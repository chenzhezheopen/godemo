package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	t := time.Now()
	// time1 := time.Time()
	// time.Sleep(time.Second * 2)
	// time2 := time.Now()
	timestamp := strconv.FormatInt(t.UTC().UnixNano(), 10)

	fmt.Println(timestamp)
}
