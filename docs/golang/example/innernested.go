package main

import "fmt"

func main() {
	// 需要定义一个递归函数
	// f1 := func(n int) int {
	// 	if n < 0 {
	// 		return 1
	// 	}
	// 	return n + f1(n-1) // 错误的方式1
	// }
	// var f2 = func(n int) int {
	// 	if n < 0 {
	// 		return 1
	// 	}
	// 	return n + f(n-1) // 错误的方式2
	// }
	var f3 func(int) int
	f3 = func(n int) int {
		if n < 0 {
			return 1
		}
		return n + f3(n-1) // 错误的方式
	}
	fmt.Printf("f3(): %v\n", f3(4))
}