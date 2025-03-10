// chan_simple_usage.go
// 演示 chan 的简单使用场景

package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		ch <- "hello"
	}()

	word := <-ch
	fmt.Printf("word: %v\n", word)

	dropData()

	readBlock()
}

func dropData() {
	ch := make(chan bool)
	go func() {
		ch <- true
	}()
	<-ch // 直接丢弃数据
	fmt.Println("dropped")
}

func readBlock() {
	ch := make(chan float64)
	<-ch
	go func() {
		ch <- 3.14
	}()
	fmt.Println("yes")
}
