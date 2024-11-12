package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	// 缓冲区为10 的管道，写入5个值，没问题
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
	// 关闭之后，仍然可以读
	/*
		n: 0 ok: true
		n: 1 ok: true
		n: 2 ok: true
		n: 3 ok: true
		n: 4 ok: true
		n: 0 ok: false
	*/
	for i := 0; i < 6; i++ { // 只有第6次读取会失败，因为已经没有数据可以读了
		n, ok := <-ch
		fmt.Printf("n: %v ", n)
		fmt.Printf("ok: %v\n", ok)
	}
	// 如果没有关闭的话，第六次读就会阻塞，入股关闭了的话，就会直接返回false 告诉你不能读到
}
