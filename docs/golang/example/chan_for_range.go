package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i * i
		}
		close(ch)
	}()
	for n := range ch {  // for range遍历管道时，当无法成功读取数据时，便会退出循环
		fmt.Println(n)
	}
}