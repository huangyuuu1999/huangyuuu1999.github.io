package main

import (
	"fmt"
	"time"
)

func test_go_routine() {
	fmt.Println("hello goroutine!")
}

func main() {
	go test_go_routine()    // 只需要在函数调用前面加一个go关键字即可
	time.Sleep(time.Second) // 主进程阻塞1s，以便于让协程完成运行
}
