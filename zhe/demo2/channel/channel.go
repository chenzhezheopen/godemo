package main

func main() {
	c, d := make(chan int, 3), make(chan int)
	var e chan int
	c <- 1
	c <- 2
	c <- 3
	println(<-c)
	println(<-c)
	println(<-c)
	println(c == d)
	println(d == e)
	println(e == nil)
	println(d == nil)

}
