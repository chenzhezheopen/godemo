package main

import (
	"net"
	"time"

	. "github.com/soekchl/myUtils"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		Error(err)
		return
	}
	var b []byte = make([]byte, 1024)

	n, err := conn.Read(b)
	if err != nil {
		Error(err)
		return
	}

	Notice(string(b[:n]))
	time.Sleep(time.Second * 10)
	conn.Close()
}
