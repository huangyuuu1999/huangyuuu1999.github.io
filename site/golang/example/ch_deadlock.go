package main

import "fmt"


// fatal error: all goroutines are asleep - deadlock!
func main() {
	ch := make(chan int)
	defer close(ch)
	ch <- 123
	n := <- ch
	fmt.Println(n)
}