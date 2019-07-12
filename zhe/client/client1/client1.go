package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	// client := &http.Client{}

	req, err := http.PostForm("http://127.0.0.1:5566/jj", url.Values{"key": {"value"}, "id": {"123"}})
	if err != nil {
		// handle error
	}
	fmt.Println(req)
	// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Set("Cookie", "Name=zheee")
	// // req.Header.Set("Cookie", "name=anny")
	// resp, err := client.Do(req)
	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	// handle error
	// }
	// fmt.Println(string(body))
}
