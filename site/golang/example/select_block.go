/*
  循环中使用 select 来读取多个管道
*/

package main

import (
	"fmt"
	"time"
)

func Send(ch chan<- int) {
	for i := range 3 {
		time.Sleep(time.Millisecond)
		ch <- i
	}
}

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)
	ch3 := make(chan int, 2)

	go Send(ch1)
	go Send(ch2)
	go Send(ch3)

	for {
		select {
		case n, ok := <-ch1:
			fmt.Printf("n: %v ok: %v\n", n, ok)
		case n, ok := <-ch2:
			fmt.Printf("n: %v ok: %v\n", n, ok)
		case n, ok := <-ch3:
			fmt.Printf("n: %v ok: %v\n", n, ok)
		}
	}
	// 当所有数据都读取完成之后，select还会等待
	// fatal error: all goroutines are asleep - deadlock!
}
