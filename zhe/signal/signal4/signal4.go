package main

import (
	"fmt"
	"net"
)

func main() {
	// stop_chan := make(chan os.Signal) // 接收系统中断信号
	// signal.Notify(stop_chan, os.Interrupt)
	// listen, _ := net.Listen("tcp", ":8080")
	// for {
	// 	// time.Sleep(time.Second * 10)
	// 	listen.Accept()
	// 	fmt.Println("打开一个443端口")
	// }
	tcpaddr, _ := net.ResolveTCPAddr("tcp4", "127.0.0.1:8080")
	tcplisten, _ := net.ListenTCP("tcp", tcpaddr)
	//死循环的处理客户端请求
	for {

		//等待客户的连接
		_, err3 := tcplisten.Accept()
		//如果有错误直接跳过
		if err3 != nil {
			continue
		}

		//通过goroutine来处理用户的请求
		fmt.Println("1111111111")
	}
}
