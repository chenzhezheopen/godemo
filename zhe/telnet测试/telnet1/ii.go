package main

import (
	"fmt"
	"net/http"
)

func head(w http.ResponseWriter, req *http.Request) {
	fmt.Println(http.StatusInternalServerError)
	w.WriteHeader(http.StatusInternalServerError)
	// fmt.Println("5555544444")
	// con, _ := ioutil.ReadAll(req.Body)
	// var mon []byte = []byte()
	// iou, _ := ioutil.ReadFile("./xqtop.html")
	w.Write([]byte("1111111111111111"))
	// fmt.Println(con)
}
