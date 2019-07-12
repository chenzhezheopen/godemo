package main

import "fmt"

// type use struct {
// 	name string
// 	age  int
// }

// var m = make(map[string]*use)

// func main() {
// 	a, err := m["tab"]
// 	if !err {
// 		fmt.Println("创建对象失败")
// 		a = &use{
// 			name: "cha",
// 			age:  15,
// 		}
// 		// m["tab"] = a
// 	}
// 	b, err := m["tab"]
// 	if !err {
// 		fmt.Println("创建对象失败")
// 		b = &use{
// 			name: "cha",
// 			age:  15,
// 		}
// 		m["tab"] = b
// 	}
// 	fmt.Println(a)

// }1111111111
// type us struct {
// 	name string
// 	fu   func(n string) string
// }

// func (us *us) uuu(pri string) {
// 	us.name = pri
// 	fmt.Println(us.name)
// 	uss := us.fu
// 	if uss != nil {
// 		fmt.Println("nil")
// 	}
// 	fmt.Println("you", uss)

// }
// func main() {
// 	var tn us
// 	tn.uuu("pppppppppp")

// }

var cache = make(map[string]string)

func main() {
	_, ok := cache["zhe"]
	if !ok {
		fmt.Println("空")
	} else {
		fmt.Println("有")
	}
}
