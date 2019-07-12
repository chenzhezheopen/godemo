package main

import (
	"fmt"
	"os"
	"os/signal"
)

//signal包中提供了两个函数
//Notifyf()用于监听信号
//Stop()用于停止监听

func main() {
	ch := make(chan os.Signal)
	//notify用于监听信号
	//参数1表示接收信号的channel
	//参数2及后面的表示要监听的信号
	//os.Interrupt 表示中断
	//os.Kill 杀死退出进程
	signal.Notify(ch, os.Interrupt, os.Kill)

	//获取信号，如果没有会一直阻塞在这里。
	s := <-ch
	//我们通过Ctrl+C或用taskkill /pid -t -f来杀死进程，查看效果。
	fmt.Println("信号：", s)
}
