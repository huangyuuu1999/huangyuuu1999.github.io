// chan_range.go
// 使用 range 依次从chan取数据

package main

import (
	"fmt"
	"time"
)

func main() {
	withClose()
	withoutClose()
}

func withClose() {
	// 创建一个通道
	ch := make(chan int)

	// 启动一个协程向通道发送值
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(1 * time.Second) // 模拟耗时操作
		}
		close(ch) // 关闭通道
	}()

	// 使用 range 遍历通道
	for v := range ch {
		fmt.Println("Received:", v)
	}
	fmt.Println("Channel closed, exiting.")
}

func withoutClose() {
	// 创建一个通道
	ch := make(chan int)

	// 启动一个协程向通道发送值
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(1 * time.Second) // 模拟耗时操作
		}
	}()

	for v := range ch { // 要读 但是没人写，会死锁
		fmt.Println("Received:", v)
	}
	fmt.Println("Channel closed, exiting.")
}
