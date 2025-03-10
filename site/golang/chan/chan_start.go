// chan_start.go  直接deadlock
package main

import "fmt"

func main() {
	// 创建了一个管道，能传输int数据
	ch := make(chan int)
	defer close(ch)

	ch <- 1
	x := <-ch
	fmt.Printf("x: %v\n", x)
}