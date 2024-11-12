package main

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

// 现实中的 协程 运行的时间不会很固定（一定是1ms），可能是完全不确定的

func main() {
	fmt.Println("start")
	for i := 0; i < 10; i++ {
		go hello(i)
		time.Sleep(time.Millisecond)
	}
	time.Sleep(time.Millisecond)
	fmt.Println("end")
}

func hello(i int) {
	// 模拟随机耗时
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	fmt.Println(i)
}
