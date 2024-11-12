package main

import "fmt"

func main() {
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)
	defer func() {
		close(chA)
		close(chB)
		close(chC)
	}()
	select {
	case n, ok := <-chA:
		fmt.Printf("n: %v ok: %v\n", n, ok)
	case n, ok := <-chB:
		fmt.Printf("n: %v ok: %v\n", n, ok)
	case n, ok := <-chC:
		fmt.Printf("n: %v ok: %v\n", n, ok)
	default:
		fmt.Println("所有管道都不可用.")
	}
}
