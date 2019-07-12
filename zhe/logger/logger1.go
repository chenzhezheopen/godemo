package main

import (
	"fmt"
	"log"
	"os"
)

var logger *log.Logger

func main() {
	logfile, _ := os.OpenFile("test.txt", os.O_RDWR|os.O_APPEND, 0)
	logger = log.New(logfile, "\r\n", log.Ldate|log.Lshortfile)
	by := make([]byte, 1024)
	n, _ := logfile.Read(by)
	nei := string(by[:n])
	fmt.Println(by[:n])
	fmt.Println(nei)
	logger.Println("lllllll")
}
