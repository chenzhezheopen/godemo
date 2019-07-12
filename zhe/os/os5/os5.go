package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {

	binary, lookErr := exec.LookPath("ll.txt")
	if lookErr != nil {
		panic(lookErr)
	}
	fmt.Println("123")
	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
