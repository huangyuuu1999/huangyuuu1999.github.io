// chan_buffer.go
// 有缓冲的 chan
package main

import "fmt"

func main() {
	chanWithBuffer()
	// blockOrNot()
	blockUntilSend()
}

func chanWithBuffer() {
	ch := make(chan int, 3)
	for i := range 5 {
		go func() {
			ch <- i
		}()
	}

	for _ = range 5 {
		// 输出的顺序是不固定的，不知道哪个协程先写
		fmt.Println(<-ch)
	}
}

// 如果从有缓冲区的管道中读，会阻塞吗？
// fatal error: all goroutines are asleep - deadlock!
func blockOrNot() {
	ch := make(chan int, 3)
	x := <-ch
	fmt.Printf("x: %v\n", x)
}

func blockUntilSend() {
	ch := make(chan string, 3)
	for _ = range 3 {
		go func() {
			ch <- "不要回答"
		}()
	}
	x := <-ch // 只读取1次，管道内剩余的没被处理
	fmt.Printf("x: %v\n", x)
}