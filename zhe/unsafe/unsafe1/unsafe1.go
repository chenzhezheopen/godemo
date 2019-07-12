package main

import (
	"fmt"
	"unsafe"
)

type T struct {
	t1 byte
	t2 int32
	t3 int64
	t4 string
	t5 bool
}

func main() {
	fmt.Println("----------unsafe.Pointer---------")
	t := &T{1, 2, 3, "this is a example", true}
	ptr := unsafe.Pointer(t)
	fmt.Println(&ptr)
	t1 := (*byte)(ptr)
	fmt.Println(*t1)
	t2 := (*int32)(unsafe.Pointer(uintptr(ptr) + unsafe.Offsetof(t.t2)))
	*t2 = 99
	fmt.Println(t)
	t3 := (*int64)(unsafe.Pointer(uintptr(ptr) + unsafe.Offsetof(t.t3)))
	fmt.Println(*t3)
	*t3 = 123
	fmt.Println(t)
	str := "hello 你好"
	fmt.Println(len([]rune(str)))
	// r := '我'
	// fmt.Printf("%T \n", r)
	// fmt.Println(r)
	// fmt.Printf("%d \n", r)
}
