package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("gless.txt")
	fmt.Println(file)
	defer file.Close()
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
