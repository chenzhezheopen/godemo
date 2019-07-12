package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	s := "a Hello 世界! "
	ts := strings.TrimLeft(s, " Helo")
	fmt.Printf("%q\n", ts) // "世界! "
	subdir := filepath.Dir("C:/gopath/src/zhe/实验")
	fmt.Println(subdir)
	fmt.Println(" :", path.Dir(""))
	fmt.Println(". :", path.Dir("."))
	os.OpenFile()
	fmt.Println("a :", path.Dir("a"))
	ioutil.ReadAll()
}
