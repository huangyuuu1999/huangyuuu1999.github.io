package main

import "fmt"

// 一个长度为n的数组a，现在需要将它分割为一些非空子数组，有多少种分割方法？
// [1, 2, 3]
// [[1], [2], [3]] [[1], [2, 3]], [[1, 2, 3]] 4 种
// 答案的个数是 2^(n-1)个
// 请返回所有的分割结果

func f(a []int) [][][]int {
	n := len(a)
	ans := [][][]int{}
	path := []int{} // 表示划线的位置

	var back func(i int)
	back = func(i int) {
		if i == n {
			tmp := [][]int{}
			lastIndex := 0
			for _, split := range path {
				tmp = append(tmp, a[lastIndex:split])
				lastIndex = split
			}
			tmp = append(tmp, a[lastIndex:])
			ans = append(ans, tmp)
			return
		}
		path = append(path, i)
		back(i + 1)
		path = path[:len(path)-1]

		back(i + 1)
	}
	back(1)
	return ans
}

func main() {
	a := []int{7, 5, 3, 9}

	fmt.Printf("nums: %v\n", a)

	ans := f(a)
	for _, v := range ans {
		fmt.Printf("v: %v\n", v)
	}
}
