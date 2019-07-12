package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	var AppPath string
	AppPath = "nini"
	// arr:=map["chen","zhe","zhe"]
	// `join`.strings.Join("Aaaaaaaaa", "sssssssssss")
	url := filepath.Join(AppPath, "conf", "app.conf")
	fmt.Println(AppPath)
	fmt.Println(url)
}
