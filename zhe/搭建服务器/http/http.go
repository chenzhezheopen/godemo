package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func httpGet() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
func main() {
	httpGet()
}
