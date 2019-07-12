package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type coo struct {
	Name  string
	Value string

	Path       string    // optional
	Domain     string    // optional
	Expires    time.Time // optional
	RawExpires string    // for reading cookies only

	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// MaxAge>0 means Max-Age attribute present and given in seconds
	MaxAge   int
	Secure   bool
	HttpOnly bool
	Raw      string
	Unparsed []string // Raw text of unparsed attribute-value pairs
}

func main() {
	http.HandleFunc("/jj", cook)
	Server := &http.Server{
		Addr: ":5566",
	}
	Server.ListenAndServe()
}
func cook(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	// var coo coo
	bod, _ := ioutil.ReadAll(req.Body)
	fmt.Println("string(bod)", string(bod))

	// io, _ := json.Marshal(req.Cookie)
	// json.Unmarshal(io, &coo)
	// fmt.Println("string(io)", coo)
	// fmt.Println("连接成功")
	Cookie := http.Cookie{Name: "zheee", Value: "chen", Path: "/", MaxAge: 3600}
	http.SetCookie(w, &Cookie)
	w.Write([]byte("222222222我设置了一个Cookie"))

	req.ParseForm()
	fmt.Println("ewq", req.Form["key"])
	for k, v := range req.Form {
		fmt.Println("key:", k)
		// join() 方法用于把数组中的所有元素放入一个字符串。
		// 元素是通过指定的分隔符进行分隔的
		fmt.Println("val:", strings.Join(v, ""))
	}
}
