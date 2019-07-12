package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hehe", wait)
	serve := &http.Server{
		Addr: ":4000",
	}
	serve.ListenAndServe()
	// fmt.Println([]byte("aaaaaa"))
	// fmt.Println([]rune("aaaaaa"))
	// fmt.Println(len([]rune("aaaa你a")))
	// fmt.Println(len([]byte("aaaa你a")))
	// fmt.Println([]byte("aaaa你a"))
	// fmt.Println([]rune("aaaa你a"))
	// fmt.Println([]byte{97, 97, 97, 97, 228, 189, 160, 97})
}

func wait(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr)
}
