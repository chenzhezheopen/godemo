package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(5)
	fmt.Println(n)

	// file, err := os.Open("gless.txt")
	// defer file.Close()
	// if err != nil {
	// 	return
	// }
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }
	// f, _ := ioutil.ReadFile("gless.txt")
	// str := string(f)
	// fmt.Println(str)
	// fileone, _ := os.OpenFile("gless.txt", os.O_RDWR|os.O_APPEND, 0)
	// fmt.Println(fileone)
}
