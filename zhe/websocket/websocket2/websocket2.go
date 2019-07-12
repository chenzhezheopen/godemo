package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"code.google.com/p/go.net/websocket"
	// "golang.org/x/net/websocket"
)

func main() {
	http.Handle("/nihao", websocket.Handler(Echo))
	http.HandleFunc("/", hello)
	http.ListenAndServe(":5555", nil)
}
func hello(w http.ResponseWriter, req *http.Request) {
	f, _ := ioutil.ReadFile("." + req.RequestURI)
	w.Write(f)
}

func Echo(w *websocket.Conn) {
	fmt.Println("lllllllllllllll")
	var error error
	// go func(){
	// 	w.Accept();
	// }
	for {
		var reply string
		if error = websocket.Message.Receive(w, &reply); error != nil {
			// fmt.Println("不能够接受消息 error==", error)
			fmt.Println(w)
			break
		}
		msg := "我已经收到消息 Received:" + reply
		if error = websocket.Message.Send(w, msg); error != nil {
			fmt.Println("不能够发送消息 悲催哦")
			break
		}
	}
}
