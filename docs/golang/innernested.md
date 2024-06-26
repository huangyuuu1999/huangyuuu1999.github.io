要在一个A函数内部，定义另一个B函数，且要求B是递归函数，观察下面几种写法。
```go
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
```
错误写法1和2，都是由于在函数B内部作用域内，找不到名称f1和f2导致的。