// cls.go
// 闭包函数定义时的上下文

package main

import "fmt"

func outer() func() int {
	val := 3
	a := func() int {
		val++
		return val
	}
	return a
}

func main() {
	val := 9
	_ = val
	f := outer()
	for _ = range 3 {
		res := f()
		fmt.Printf("res: %v\n", res)
	}
}