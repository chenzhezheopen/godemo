package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// type ResFolderItem struct {
// 	Path       string   `json:"path"`
// 	ExtAllowed []string `json:"ext"`
// }

// //Config --
// type Config struct {
// 	MimeType         map[string]string `json:"mimetype"` //
// 	StaticResFolders []ResFolderItem   `json:"static"`   //
// }

// var cfg *Config

func LastIndexByte(s string, c byte) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			return i
		}
	}
	return -1
}
func SayHello(w http.ResponseWriter, req *http.Request) {
	// var config = &Config{}
	// jso, _ := ioutil.ReadFile("config.json")
	// json.Unmarshal(jso, config)
	pos := strings.LastIndexByte(req.RequestURI, '.')
	uri := req.RequestURI

	ext := uri[pos+1:]

	// mime := config.MimeType[ext]
	if ext == "htm" || ext == "html" {
		w.Header().Set("Content-Type", "text/html")
	} else if ext == "css" {
		w.Header().Set("Content-Type", "text/css")
	} else if ext == "js" {
		w.Header().Set("Content-Type", "application/x-javascript")
	}
	buf, err := ioutil.ReadFile("." + req.RequestURI)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// w.Header().Set("Content-Type", mime)
	w.Write(buf)

	// ion, _ := ioutil.ReadFile("hello/InterBlack.html")
	// w.Header().Set("Content-Type", "text/html")
	// w.Header().Set("Content-Type", "text/html")
	// w.Header().Set("Content-Type", "text/html")
	// w.Write(ion)
	// fmt.Println(req.Form["id"])
	// //获取我们set进去的nginx的请求时间，并返回给客户端
	// os.Open("hello/InterBlack.html")
	// w.Write(f)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// // defer f.Close()
	// io.Copy(w, f)
	// fmt.Println(req.Form)
	// fmt.Println(req.Form("name"))
	// fmt.Println(w)
	// fmt.Println(req.RequestURI)
	// fmt.Println("xuny")
	// fmt.Println(req.Method)
	// w.Write([]byte("Hello"))
}

func main() {
	//访问的URL以及对应的方法
	// server := http.Server{
	// 	Addr: "localhost:8002",
	// }
	// http.Handle("/xcv/", http.FileServer(http.Dir("hello")))

	http.HandleFunc("/hello/", SayHello)
	// http.HandleFunc("/hello/", SayHello)
	//设置监控端口号
	http.ListenAndServe(":1999", nil)
	// http.ListenAndServe(":8001", nil)

}
