package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr: ":5656",
	}
	http.HandleFunc("/daili", daili)
	server.ListenAndServe()
}
func daili(w http.ResponseWriter, req *http.Request) {
	fmt.Println("接收代理")
	coo, err := req.Cookie("zhe")
	fmt.Println(coo)
	if err != nil {
		fmt.Println("没收到cookie")
	} else {
		fmt.Println(coo)
	}
	value := req.Body

	body, _ := ioutil.ReadAll(value)
	fmt.Println("这是服务器发送给代理服务器的消息", body)
	w.Write([]byte("这是代理返回的信息"))
}
