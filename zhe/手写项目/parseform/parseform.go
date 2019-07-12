package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Println("r.host", r.Host)
	fmt.Println("w", r.PostFormValue("password"))
	fmt.Println("地址", r.RemoteAddr[:strings.Index(r.RemoteAddr, ":")])
	fmt.Println("地", strings.Index(r.RemoteAddr, ":"))
	fmt.Println("jinru")
	// 解析url传递的参数
	r.ParseForm()
	for k, v := range r.Form {
		fmt.Println("key:", k)
		// join() 方法用于把数组中的所有元素放入一个字符串。
		// 元素是通过指定的分隔符进行分隔的
		fmt.Println("val:", strings.Join(v, ""))
	}
	// 输出到客户端
	name := r.Form["username"]
	pass := r.Form["password"]
	fmt.Println(name)
	for i, v := range name {
		fmt.Println(i)
		fmt.Fprintf(w, "用户名:%v\n", v)
	}
	for k, n := range pass {
		fmt.Println(k)
		fmt.Fprintf(w, "密码:%v\n", n)
	}
}
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("one.html")
		fmt.Println("t", t)
		// func (t *Template) Execute(wr io.Writer, data interface{}) error {
		t.Execute(w, nil)
		fmt.Println("w", w)
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}
func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndserve:", err)
	}
}
