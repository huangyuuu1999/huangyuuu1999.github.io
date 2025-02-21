# 闭包
先看一个案例。
```go
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
```
函数是一等公民，可以被赋值、传递，返回。outer函数在内部定义了一个闭包函数，并返回出来。
但是注意，当闭包函数在外面被接受时，其中的val变量，仍然是定义他的时候的上下文中的val。

在外面调用闭包函数，修改的也是定义的时候上下文中的val。而不是main当中的val=9。