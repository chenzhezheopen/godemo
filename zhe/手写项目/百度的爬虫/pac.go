package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	var look string
	fmt.Println("输入想要查找的关键字")
	fmt.Scan(&look)
	lookup(look)
}

func lookup(look string) {
	ur := "https://www.baidu.com/s?wd=" + look
	HttpGet(ur)
}

func HttpGet(url string) (result string, err error) {
	file, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	buf := make([]byte, 4*1024)
	for true {
		f, err := file.Body.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件读取完毕")
				break
			} else {
				fmt.Println("resp.Body.Read err = ", err)
				break
			}
		}
		result += string(buf[:f])
	}

}
