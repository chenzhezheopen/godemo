package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	server := &http.Server{
		Addr: ":8585",
	}
	http.HandleFunc("/accept", accept)
	server.ListenAndServe()
}

func accept(w http.ResponseWriter, req *http.Request) {
	// Cookie := http.Cookie{Name: "zhe", Value: "chen", Path: "/", MaxAge: 3600}
	// http.SetCookie(w, &Cookie)
	// resp, _ := http.Get("http://www.baidu.com")
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	// url :=
	// proxy, _ := url.Parse("http://127.0.0.1:5656/daili")
	// tr := &http.Transport{
	// 	Proxy:           http.ProxyURL(proxy),
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }
	re := req.Body
	s, _ := ioutil.ReadAll(re)
	str := string(s)
	reqs, _ := http.NewRequest("POST", "http://localhost:5656/daili", strings.NewReader(str))
	reqs.Header.Set("Cookie", "zhe=chen")
	client := &http.Client{}

	resp, err := client.Do(reqs)
	if err != nil {
		fmt.Println("出错了", err)
		return
	}
	defer resp.Body.Close()
	// resp, _ := http.Get(webUrl)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	coo, err := req.Cookie("zhe")
	fmt.Println(coo)
	if err != nil {
		fmt.Println("没收到cookie")
	} else {
		fmt.Println(coo)
	}
}
