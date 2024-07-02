package main

import "fmt"

func main() {
	s1 := []int{2, 3, 5}
	test1(s1)
	fmt.Printf("s1: %v\n", s1) // s1: [2 3 8]

	test2(s1)
	fmt.Printf("s1: %v\n", s1) // s1: [2 3 8], test2 内部改的，不是s1的数组

	fmt.Printf("s1 len: %v cap %v\n", len(s1), cap(s1)) // s1 len: 3 cap 3

	s1 = append(s1, 408) 
	fmt.Printf("s1: %v\n", s1) // s1: [2 3 8 408]
	fmt.Printf("s1 len: %v cap %v\n", len(s1), cap(s1))  // s1 len: 4 cap 6

	test2(s1)
	fmt.Printf("s1: %v\n", s1) // s1: [-3 3 8 408], 这次 test2 内部改的，是s1的数组
	fmt.Printf("s1: %p\n", s1)
}

func test1(s []int) {  // 切片作为参数，是传引用
	s[2] = 8
	fmt.Printf("s: %v\n", s) // s: [2 3 8]
}

func test2(s []int) {
	s = append(s, 42)  // append之后，新的s可能和原先的s1 公用数组，也可能不共用数组
	s[0] = -3
	fmt.Printf("s: %v\n", s) // s: [-3 3 8 42]
	fmt.Printf("s1: %p\n", s)
}

