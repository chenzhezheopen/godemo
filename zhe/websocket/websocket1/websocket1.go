package main

import (
    "fmt"
    "net/http"
    "code.google.com/p/go.net/websocket"
)

const queueSize = 20

type Input struct {
    Cmd    string
}

type Output struct {
    Cmd    string
}

func Handler(ws *websocket.Conn) {

    msgWrite := make(chan *Output, queueSize)
    var in Input

    go writeHandler(ws, msgWrite)

    for {

        err := websocket.JSON.Receive(ws, &in)

        if err != nil {
            fmt.Println(err)
            break
        } else {
            msgWrite <- &Output{Cmd: "Thanks for your message: " + in.Cmd}
        }
    }
}

func writeHandler(ws *websocket.Conn, out chan *Output) {
    var d *Output
    for {
        select {
        case d = <-out:
            if err := c(ws, &d); err != nil {
                fmt.Println(err.Error())
            } else {
                fmt.Println("> ", d.Cmd)
            }
        }
    }
}

func main() {
    http.Handle("/echo", websocket.Handler(Handler));
    err := http.ListenAndServe(":1235", nil);
    if err != nil {
        panic("ListenAndServe: " + err.Error())
    }
    fmt.Println("Server running")
}