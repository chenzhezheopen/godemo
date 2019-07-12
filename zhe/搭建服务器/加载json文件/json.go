package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//定义配置文件解析后的结构
type MongoConfig struct {
	MongoAddr      string
	MongoPoolLimit int
	MongoDb        string
	MongoCol       string
}

type Config struct {
	Add   string `json:"Addr"`
	Mongo MongoConfig
}

func main() {
	file, _ := os.Open("InterBlack.html")

	var b []byte = make([]byte, 4096)
	n, rr := file.Read(b)
	fmt.Println("n", n)
	if rr != nil {
		fmt.Println("Open file Failed", rr)
	}
	da := string(b[:n])
	fmt.Println("da", da)

	var JsonParse = NewJsonStruct()
	v := Config{}
	m := Config{}
	//下面使用的是相对路径，config.json文件和main.go文件处于同一目录下
	JsonParse.Load("config.json", &v)
	fmt.Println(v.Add)
	fmt.Println(v.Mongo.MongoDb)

	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		return
	}
	json.Unmarshal(data, &m)
	fmt.Println(m.Add)
	fmt.Println(m.Mongo.MongoDb)
}

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func (jst *JsonStruct) Load(filename string, v interface{}) {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	fmt.Println(string(data))
	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}
