package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err != nil {
		err = err1
		return
	}
	buf := make([]byte, 4*1024)
	for true {
		n, err := resp.Body.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件读取完毕")
				break
			} else {
				fmt.Println("resp.Body.Read err = ", err)
				break
			}
		}
		result += string(buf[:n])

	}
	fmt.Println(strings.Contains(result, "itemprop="))
	return
}

func SpiderPage(i int, page chan<- int) {
	url := "https://hotels.ctrip.com/hotel/" + strconv.Itoa((i - 1)) + ".html"
	fmt.Printf("正在爬取第%d个网页\n", i)
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("http.Get err = ", err)
		return
	}
	filename := strconv.Itoa((i - 1)) + ".html"
	f, err1 := os.Create(filename)
	if err1 != nil {
		fmt.Println("os.Create err = ", err1)
		return
	}
	f.WriteString(result)
	f.Close()
	page <- i
}

func DoWork(start, end int) {
	fmt.Printf("正在爬取第%d页到%d页\n", start, end)
	page := make(chan int)
	for i := start; i <= end; i++ {
		go SpiderPage(i, page)
	}
	for i := start; i <= end; i++ {
		fmt.Printf("第%d个页面爬取完成\n", <-page)
	}
}

func main() {
	var start, end int
	fmt.Printf("请输入起始页>=1：> ")
	fmt.Scan(&start)
	fmt.Printf("请输入结束页：> ")
	fmt.Scan(&end)
	DoWork(start, end)
}
