package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func main() {
	resp, err := http.PostForm("http://14.215.177.39/s?", url.Values{"wd": {"111"}})
	if err != nil {
		fmt.Println(err)
	}
	f, err := os.Create("index.html")
	if err != nil {
		fmt.Println("文件创建失败")
	}
	buf := make([]byte, 4*1024)
	for true {
		n, err := resp.Body.Read(buf)
		f.Write(buf[:n])
		if err != nil {
			if err == io.EOF {
				// fmt.Println("文件读取完毕")
				break
			} else {
				// fmt.Println("resp.Body.Read err = ", err)
				break
			}
		}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))

}
