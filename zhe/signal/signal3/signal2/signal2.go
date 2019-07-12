package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	var sl []net.Conn
	for i := 1; i < 1000; i++ {
		conn, _ := net.Dial("tcp", ":8080")
		if conn != nil {
			sl = append(sl, conn)
		}
		fmt.Println(i)
	}
	time.Sleep(time.Second * 10000)
}
