// defer_arg.go
// 演示defer的参数
package main

import "fmt"


func main() {
	// test()
	test2()
}

func test() int {
	defer func(arg int) {
		fmt.Printf("arg: %v\n", arg)
	}(42)
	defer fmt.Println(1)
	return 0
}

func compute(a, b int) int {
	return a + b
}

func test2() {
	a, b := 1, 2
	defer func(x int) {
		fmt.Printf("x: %v\n", x)
	}(compute(a, b)) // defer的参数也可以是函数调用的结果
	a, b = 3, 4
	panic("wrong")
}