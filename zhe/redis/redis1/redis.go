package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	http.Handle("/wat", websocket.Handler(wait))
	http.HandleFunc("/", syt)
	server := http.Server{
		Addr: ":7777",
	}
	server.ListenAndServe()
}

func syt(rw http.ResponseWriter, req *http.Request) {
	// fmt.Println("连接成功") 
	file, _ := ioutil.ReadFile("sed.html")
	rw.Write(file)
}
func wait(w *websocket.Conn) {
	var error error
	fmt.Println("有一个client连接")
	for {
		fmt.Println("进入for循环")
		var str string
		if error = websocket.Message.Receive(w, &str); error != nil {
			break
		} else {
			fmt.Println("str", str)
		}
		fmt.Println(str)
		var ui = "服务器发送的消息" + str
		if error = websocket.Message.Send(w, ui); error != nil {
			break
		}
	}
}
