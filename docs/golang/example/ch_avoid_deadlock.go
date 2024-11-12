package main

import "fmt"

// 不会发生死锁
func main() {
	ch := make(chan int)
	go func() {
		ch <- 123123
	}() // 立即创建 这个函数的协程
	n := <-ch
	fmt.Println(n)
}
