package main

import (
	"fmt"
	"sync"
)

func hello(wait sync.WaitGroup) {
	fmt.Println("hello")
	wait.Done() // 错误
}

func main() {
	var mainWait sync.WaitGroup
	mainWait.Add(1)
	hello(mainWait)
	mainWait.Wait()
	fmt.Println("end")
}
/*
错误提示所有的协程都已经退出，但主协程依旧在等待，
这就形成了死锁，
因为hello函数内部对一个形参WaitGroup调用Done并不会作用到原来的mainWait上，
所以应该使用指针来进行传递。
*/