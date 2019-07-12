package main

import (
	"fmt"
	"net/http"
	"text/template"

	"golang.org/x/net/websocket"
)

func main() {
	http.HandleFunc("/", hello)
	// http.HandleFunc("/ws/", hhh)
	http.Handle("/ws", websocket.Handler(handleWS))
	http.ListenAndServe(":1999", nil)
}

var tmpl *template.Template

func handleWS(w *websocket.Conn) {
	fmt.Println("11111111111111111111111111111111")
	// tmpl := `<!DOCTYPE html>
	// <html>
	// 	<head>
	// 		<meta http-equiv="Content-Type" content="text/html; charset=utf-8"> <title>Go Web Programming</title>
	// 	</head>
	// 	<body>
	// 		{{ . }}
	// 	</body>
	// </html>`

	// t := template.New("layout.html")
	// t, _ = t.Parse(tmpl)
	// t.Execute(w, "jjjjjjjjjjjjjj")
}
func hello(w http.ResponseWriter, r *http.Request) {
	wsURL := fmt.Sprintf("ws://%s/ws", r.Host)
	fmt.Println(wsURL)
	tmpl.Execute(w, wsURL)

}
