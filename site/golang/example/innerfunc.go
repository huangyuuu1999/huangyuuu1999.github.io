package main

import "fmt"

// func main() {
// 	fmt.Println("Outer function")
// 	var innerFunction = func() {
//         fmt.Println("Inner function")
//     }
// 	innerFunction()
// }

func main() {
	fmt.Println("Outer function")
	innerFunction := func() {
        fmt.Println("Inner function")
    }
	innerFunction() // 没有这一句会出错，因为innerFunction定义而不使用

	innerfunc2()
	call_it_now()
}

func innerfunc2() {
	x := 20
	this_func_can_visit_x := func() {
		fmt.Println("this_func_can_visit_x", x)
	}
	this_func_can_visit_x() // this_func_can_visit_x 20
}

func call_it_now() {
	// 立即调用函数
	ans := func(n int) int {
		return 5
	}(2)
	fmt.Println(ans)
}