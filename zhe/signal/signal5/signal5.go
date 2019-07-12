package main

import (
	"log"
	"net/http"
)

// func main() {
// 	list, err := net.Listen("tcp", ":8080")
// 	if err != nil {
// 		log.Println("error listen:", err)
// 		return
// 	}
// 	defer list.Close()
// 	for {
// 		conn, _ := list.Accept()
// 		fmt.Println(conn)
// 		head(conn)
// 	}
// }
// func head(conn net.Conn) {
// 	fmt.Println("connc", conn)
// }
func main() {
	var mux http.ServeMux
	mux.Handle("/a", &myHandler{})
	mux.HandleFunc("/bye", sayBye)

	log.Println("Starting v2 httpserver")
	http.ListenAndServe(":1210", nil)
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is version 2"))
}
func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bye bye ,this is v2 httpServer"))
}
