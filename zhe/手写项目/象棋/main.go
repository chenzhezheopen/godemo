package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"

	"golang.org/x/net/websocket"
)

type diff struct {
	sync.Mutex
	contact chan adopt
	num     int
	conn    []*websocket.Conn
	man     []adopt
}
type adopt struct {
	name string
	age  int
	sex  string
}

func (man *diff) block() bool {
	return len(man.conn) < 10
}

func (man *diff) judge() string {
	return strconv.Itoa(man.num)
}

func (man *diff) En() {
	man.num += 1
}

func (man *diff) Read() {
	man.Lock()
	defer man.Unlock()
	var msg = Ent.judge()
	var error error
	for i := 0; i < len(Ent.conn); i++ {
		if error = websocket.Message.Send(Ent.conn[i], msg); error != nil {
			return
		}
	}
}

// var tmpl *template.Template
var done = make(chan bool, 1)
var Ent diff

// var hhh = make(chan bool, 1)

// var fla = make([]*websocket.Conn, 0, 10)

func main() {
	http.Handle("/xqto", websocket.Handler(wait))
	http.HandleFunc("/", star)
	// http.Handle("/num", websocket.Handler(number))
	serve := &http.Server{
		Addr:    ":6363",
		Handler: nil,
	}
	serve.ListenAndServe()
}

func star(w http.ResponseWriter, req *http.Request) {
	fmt.Println("req", req.RemoteAddr)
	if Ent.block() {
		f, _ := ioutil.ReadFile("." + req.RequestURI)
		w.Write(f)
		if req.RequestURI == "/xqtop.html" {
			Ent.En()
		}
		go func() {
			<-done
			Ent.Read()
			websocket.Message.Send(Ent.conn[len(Ent.conn)-1], Ent.judge())
		}()
	} else {
		w.Write([]byte("服务器过载"))
	}
	// wait(fla)
	// go func() {
	// 	<-hhh
	// 	wsURL := fmt.Sprintf("ws://%s/ws", req.Host)
	// 	fmt.Println(wsURL)
	// 	tmpl.Execute(w, "."+req.RequestURI)
	// }()
}

func wait(ws *websocket.Conn) {
	fmt.Println("111111111111")
	var msg = Ent.judge()
	var error error
	Ent.conn = append(Ent.conn, ws)
	done <- true
	// websocket.Message.Send(ws, ms)
	for {
		var reply string
		if error = websocket.Message.Receive(ws, &reply); error != nil {
			fmt.Println("不能够接受消息 error==", error)
			break
		} else {
			fmt.Println("客户端传来一条消息")

		}
		msg = Ent.judge()
		if error = websocket.Message.Send(ws, msg); error != nil {
			fmt.Println("不能够发送消息")
			return
		} else {
			fmt.Println("发送了一条消息给客户端")
		}
		// hhh <- true

	}
}
