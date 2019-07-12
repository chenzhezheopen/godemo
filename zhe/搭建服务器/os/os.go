package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	// var file, err = os.Open("config.json")
	// if err != nil {
	// 	fmt.Println("加载失败")
	// }
	// // file.Close()
	// var b []byte = make([]byte, 4096)
	// n, err := file.Read(b)
	// if err != nil {
	// 	fmt.Println("Open file Failed", err)
	// }
	// data := string(b[:n])
	// fmt.Println(data)
	// fmt.Println(n)
	http.HandleFunc("/", newfun)
	http.ListenAndServe(":20000", nil)

}
func newfun(w http.ResponseWriter, req *http.Request) {
	buf, err := ioutil.ReadFile("./InterBlack.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// w.Header().Set("Content-Type", mime)
	w.Write(buf)
}
