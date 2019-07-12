package main

import (
	"fmt"
	"sync"
)

type user interface {
}
type kind struct {
	sync.RWMutex
}

func (kind *kind) hee(t func()) {
	kind.Lock()
	defer kind.Unlock()
	fmt.Println("mmmm")
	table.addedItem=t
}
func main() {
	var k kind
	k.hee(func() {
		fmt.Println("nnnnnnnn")
	})
}
