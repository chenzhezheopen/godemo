package main

import (
	"fmt"
	"net"
	"time"
	// . "github.com/soekchl/myUtils"
)

func main() {
	for i := 0; i < 200; i++ {
		net.DialTimeout("tcp", ":8088", time.Second*10)
		// if conn != nil {
		// 	return
		// }
		// defer conn.Close()
		fmt.Println(i)
		// var b []byte = make([]byte, 1024)

		// n, err := conn.Read(b)
		// if err != nil {
		// 	Error(err)
		// 	return
		// }

	}
	// Notice(string(b[:n]))
	time.Sleep(time.Second * 100000)
}
