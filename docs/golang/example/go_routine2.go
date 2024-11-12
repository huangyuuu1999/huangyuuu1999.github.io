// 使用匿名函数立即执行来开启协程
package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("go + anonymous func()")
	}()
	time.Sleep(time.Second)
}
