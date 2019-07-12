package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	var start, end int
	var city string
	fmt.Printf("请输入起始页")
	fmt.Scan(&start)
	fmt.Printf("请输入结束页")
	fmt.Scan(&end)
	fmt.Println("请输入查找城市")
	fmt.Scan(&city)
	ur(start, end, city)
	// fmt.Println(gu)
	// f.WriteString(result)
	// fmt.Println(result)
}
func ur(start int, end int, city string) {
	ch := make(chan int)
	data := make(chan bool, 20)
	file, _ := os.Create(city + ".txt")
	for i := start; i <= end; i++ {
		select {
		case data <- true:
			go zi(i, ch, city, data, file)
		case <-ch:
			<-data
		}

	}
	// for j := start; j < end; j++ {
	// 	<-ch
	// }
}
func zi(u int, ch chan<- int, city string, data chan<- bool, file *os.File) {

	res, _ := http.Get("https://hotels.ctrip.com/international/" + strconv.Itoa((u - 1)) + ".html") //https://hotels.ctrip.com/hotel/
	buf := make([]byte, 4*1024)
	var result string
	for true {
		n, err := res.Body.Read(buf)
		if err != nil {
			if err == io.EOF {
				// fmt.Println("文件读取完毕")
				break
			} else {
				// fmt.Println("resp.Body.Read err = ", err)
				break
			}
		}
		result += string(buf[:n])
	}
	ch <- u
	rrr := strings.Index(result, "en_n")
	if rrr <= 1 {
		return
	}

	gu := result[rrr-50 : rrr-29]
	r := []rune(gu)
	//fmt.Println("rune=", r)
	strSlice := []string{}
	cnstr := ""
	for i := 0; i < len(r); i++ {
		if r[i] <= 40869 && r[i] >= 19968 {
			cnstr = cnstr + string(r[i])
		}
		//fmt.Println("r[", i, "]=", r[i], "string=", string(r[i]))
	}
	strSlice = append(strSlice, cnstr)
	// if 0 == len(strSlice) {
	// 	//无中文，需要跳过，后面再找规律
	// }
	if len(cnstr) < 2 {
		return
	}
	in := strings.Index(cnstr, city)
	if in < 0 {
		return
	}
	fmt.Println(cnstr)
	file.WriteString("\n" + cnstr)
	return
	// fmt.Println(gu)
}
