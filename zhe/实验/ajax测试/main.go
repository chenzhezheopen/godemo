package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"code.google.com/p/go.net/websocket"
)

func main() {
	file, _ := os.Open("papa/1.jpg")
	file1, _ := os.OpenFile("index.jpg", os.O_WRONLY|os.O_CREATE, 0666)
	b1 := make([]byte, 1024)
	n1, _ := file.Read(b1)
	fmt.Println(b1[:n1])
	io.Copy(file, file1)
	var conn = make(map[string]int)
	conn["l"] = 5
	fmt.Println(conn)
	delete(conn, "l")
	fmt.Println(conn)

	http.HandleFunc("/", ji)
	http.HandleFunc("/img", img)
	server := http.Server{
		Addr: ":7878",
	}
	http.Handle("/add", websocket.Handler(add))
	server.ListenAndServe()
}
func img(w http.ResponseWriter, req *http.Request) {
	file, _ := os.Open("papa/1.jpg")
	file1, _ := os.Open("index.jpg")
	io.Copy(file, file1)
	// req.ParseForm()
	// uploadFile, handle, _ := req.FormFile("img")
	// fmt.Println(handle.Filename)
	// // os.Mkdir("index/", 0777)
	// saveFile, _ := os.OpenFile("index/"+handle.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	io.Copy(file, file1)

	// defer uploadFile.Close()
	// defer saveFile.Close()

}
func ji(w http.ResponseWriter, req *http.Request) {
	// fmt.Println("aaaa")
	pos := strings.LastIndexByte(req.RequestURI, '.')
	uri := req.RequestURI

	ext := uri[pos+1:]

	// mime := config.MimeType[ext]
	if ext == "htm" || ext == "html" {
		w.Header().Set("Content-Type", "text/html")
	} else if ext == "css" {
		w.Header().Set("Content-Type", "text/css")
	} else if ext == "js" {
		w.Header().Set("Content-Type", "application/x-javascript")
	}
	buf, err := ioutil.ReadFile("." + req.RequestURI)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// w.Header().Set("Content-Type", mime)
	w.Write(buf)
}

func add(w *websocket.Conn) {
	fmt.Println("JIJIJIJIJ")
	var error error
	for {
		var reply string
		if error = websocket.Message.Receive(w, &reply); error != nil {
			fmt.Println("不能够接受消息 error==", error)
			break
		}
		fmt.Println(reply)
		msg := "我已经收到消息 Received:" + reply
		if error = websocket.Message.Send(w, msg); error != nil {
			fmt.Println("不能够发送消息 悲催哦55")
			break
		}
	}
}
