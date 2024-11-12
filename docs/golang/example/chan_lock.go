package main

import (
	"fmt"
	"time"
)

var count int = 0

var lock = make(chan struct{}, 1) // 缓冲区大小为1

func Add() {
	lock <- struct{}{} // acquire lock
	count++
	fmt.Printf("count: %v\n", count)
	<-lock // release lock, 可以再被其他协程获得
}

func Sub() {
	lock <- struct{}{}
	count--
	fmt.Printf("count: %v\n", count)
	<-lock
}

func main() {
	defer close(lock)
	for i := range 20 {
		if i % 2 == 0 {
			go Add()
		} else {
			go Sub()
		}
	}
	time.Sleep(time.Second)
}