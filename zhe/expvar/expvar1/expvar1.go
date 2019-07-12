package main

import (
	"expvar"

	"fmt"

	"net/http"
)

// var stats = expvar.NewMap("tcp")

var visits = expvar.NewInt("visits")

func handler(w http.ResponseWriter, r *http.Request) {

	visits.Add(1)
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

}

// var requests, requestsFailed expvar.Int

// func init() {

// 	stats.Set("requests", &requests)

// 	stats.Set("requests_failed", &requestsFailed)

// }

func main() {

	http.HandleFunc("/", handler)

	http.ListenAndServe(":1818", nil)
}
