// 演示具有返回值的内置函数，不可以用go开启协程
package main

import "time"

func main() {
	// go discards result of make([]int, 10) (value of type []int)
	go make([]int, 10)

	// Sleep has no return value, so it's ok
	go time.Sleep(time.Second)
}