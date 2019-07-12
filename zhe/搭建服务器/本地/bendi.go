package main

import (
	"net/http"
)

func main() {
	server := http.Server{
		Addr: ":8001",
	}
	http.Handle("/hello", http.FileServer(http.Dir("hello")))

	server.ListenAndServe()
}
