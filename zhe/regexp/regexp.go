package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"regexp"
)

func main() {

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println("1", match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println("2", r.MatchString("peach"))

	fmt.Println("3", r.FindString("peach punch"))

	fmt.Println("4", r.FindStringIndex("peach punch"))

	fmt.Println("5", r.FindStringSubmatch("peach punch"))

	fmt.Println("6", r.FindStringSubmatchIndex("peach punch"))

	fmt.Println("7", r.FindAllString("peach punch pinch", -1))

	fmt.Println("8", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	fmt.Println(r.FindAllString("peach punch pinch", 2))

	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
	// rand.Seed(time.Now().Unix())
	fmt.Println(rand.Intn(1000))
	fmt.Println(rand.Intn(1000))
}
