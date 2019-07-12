package main

import "os"
import "fmt"

func main() {

	// argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	fmt.Println(argsWithoutProg)
	// arg := os.Args[3]
	// fmt.Println("arg", arg)
	// fmt.Println("argsWithProgaaaaa", argsWithProg)
	for k, _ := range argsWithoutProg {
		// fmt.Println(k)
		// fmt.Println("vasklgjosid")
		fmt.Println(argsWithoutProg[k])
	}

	// fmt.Println("argsWithoutProg", argsWithoutProg)

}
