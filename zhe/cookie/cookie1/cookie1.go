package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/write/", writecookie)
	server := &http.Server{
		Addr: ":8899",
	}
	server.ListenAndServe()
}
func writecookie(w http.ResponseWriter, res *http.Request) {
	fmt.Println(res.Cookie)
	
	Cookie := http.Cookie{Name: "name", Value: "chen", Path: "/", MaxAge: 60}
	http.SetCookie(w, &Cookie)
	fmt.Println(Cookie.String())
	w.Write([]byte("write cookie ok"))
}
