package main

import "fmt"

// 有一个n个元素的集合 {1, 2, 3, 4 ...n} 现在需要把他划分成一些非空子集 总共有几种划分方式？
// example: n = 4
// { {1} {2} {3} {4} }
// { {1, 2} {3, 4} }
// { {1} {2, 3, 4} }
// { {2} {1, 3, 4} }
// ...

// {1 .. a} a 个元素，划分成 b 个部分，有多少种方案
func F(a int, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	if a == 1 && b == 1 {
		return 1
	}
	return F(a-1, b-1) + F(a-1, b)*b
}

func main() {
	a := 4
	for i := 1; i <= 4; i++ {
		fmt.Printf("F(a, 1): %v\n", F(a, i))
	}
}
