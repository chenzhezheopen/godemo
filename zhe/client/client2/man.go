package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	// 解析url传递的参数
	fmt.Println("r", r)
	bod, _ := ioutil.ReadAll(r.Body)
	fmt.Println("string", string(bod))
	client := &http.Client{}
	r.ParseForm()
	req, err := http.NewRequest("POST", "http://localhost:5566/jj", strings.NewReader("username=12156&password=aa"))
	if err != nil {
		// handle error
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "Name=zheee")
	// req.Header.Set("Cookie", "name=anny")
	resp, _ := client.Do(req)
	fmt.Println("resp", resp)
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
	data := make(map[string][]string)
	data["key"] = []string{"this is key"}
	data["value"] = []string{"this is value"}
	req, err := http.PostForm("http://127.0.0.1:5566/jj", data)
	if err != nil {
		// handle error
	}
	fmt.Println(req)
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		// func (t *Template) Execute(wr io.Writer, data interface{}) error {
		t.Execute(w, nil)
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
