package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Replace:   ", strings.Replace("foooooo", "o", "0", -1))
	fmt.Println("Replace:   ", strings.Replace("fooooooo", "o", "0", 2))
	http.HandleFunc("/he", head)
	http.ListenAndServe(":8099", nil)
}
