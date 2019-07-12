package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"

	. "github.com/soekchl/myUtils"
)

var (
	wg sync.WaitGroup = sync.WaitGroup{} // 等待各个socket连接处理
)

func main() {

	stop_chan := make(chan os.Signal) // 接收系统中断信号
	signal.Notify(stop_chan, os.Interrupt)

	listen, err := net.Listen("tcp", ":8088")
	if err != nil {
		Error(err)
		return
	}

	go func() {
		<-stop_chan
		// Warn("Get Stop Command. Now Stoping...")
		fmt.Println("开始中断")
		if err = listen.Close(); err != nil {
			Error(err)
		}
	}()

	Notice("Start listen :8080 ... ")
	for {
		time.Sleep(time.Second * 10)
		// time.Sleep(time.Second * 10)
		conn, err := listen.Accept()

		if err != nil {
			if strings.Contains(err.Error(), "use of closed network connection") {
				break
			}
			Error(err)
			continue
		}
		Info("Accept ", conn.RemoteAddr())
		wg.Add(1)
		// fmt.Println("4aaaaaaaaaaaa")
		// fmt.Println(conn)
		go Handler(conn)
	}

	wg.Wait() // 等待是否有未处理完socket处理
}

func Handler(conn net.Conn) {
	defer wg.Done()
	defer conn.Close()

	time.Sleep(5 * time.Second)

	conn.Write([]byte("Hello!"))
	Info("Send hello")
}
