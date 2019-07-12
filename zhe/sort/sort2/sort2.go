package main

import "sort"
import "fmt"

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	// var byLength
	fruits := []string{"peach", "banana", "kiwi"}
	v := 0 << 3
	fmt.Println(v)
	fmt.Println(len(fruits))
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
