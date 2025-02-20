// defer_return.go
// 演示 defer 和 return 谁先执行
package main

import "fmt"

func main() {
	// 观察先输出什么内容
	wrapper()
	ans := wrapper2()
	fmt.Printf("ans: %v\n", ans)
}

func deferFunc(i int) {
	for j := range i {
		fmt.Println(j, "defer was called.")
	}
}

func returnFunc() int {
	fmt.Println("return was called.")
	return 1
}

func wrapper() int { // 说服力有限
	defer deferFunc(3)
	return returnFunc()
}

func wrapper2() (ans int) { // 能够有力证明 defer 就是在 return语句之后 启动的
	defer func() {
		fmt.Printf("ans: %v\n", ans)
		ans += 10
	}()
	return ans + 3
}