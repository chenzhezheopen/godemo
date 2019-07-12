package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type use struct {
	q string
	p bool
	s []string `CHARSET:"UTF-8"`
}
type sug struct {
	sug use
}

func main() {
	var m, err = http.Get("https://sp0.baidu.com/5a1Fazu8AA54nxGko9WTAnF6hhy/su?wd=aaa&charset=utf8")

	m.Header.Add("content-type", "application/x-www-form-urlencoded")
	m.Header.Add("cache-control", "no-cache")
	m.Header.Add("CHARSET", "utf8")
	if err != nil {
		// handle error
		fmt.Println(123)
	}
	// m = Charset("UTF-8")
	body, err := ioutil.ReadAll(m.Body)

	if err != nil {
		// handle error
		fmt.Println(123)
	}
	number := string(body)
	fmt.Println("number", number)

	fmt.Println(string(body))
	var shuai = sug{}
	json.Unmarshal(body, &shuai)
	fmt.Println(456)
	fmt.Println(shuai.sug.p)
	fmt.Println(789)
}
