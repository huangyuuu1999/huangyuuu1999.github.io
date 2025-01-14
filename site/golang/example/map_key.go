package main

import "fmt"

func main() {
	// 整数作为键
	mapInt := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	fmt.Printf("mapInt: %v\n", mapInt)

	// 字符串作为键
	m1 := map[string]int{"one": 1, "two": 2}
	fmt.Printf("m1: %v\n", m1)

	// 浮点数作为键（不推荐，因为浮点数的比较可能会因精度问题导致不准确）
	m2 := map[float64]bool{1.2: false, 3.14: true}
	fmt.Printf("m2: %v\n", m2)

	// 数组类型作为键，因为数组是长度固定的，可以比较两个数组是否相等
	m3 := make(map[[2]int]string, 4)
	m3[[...]int{1, 2}] = "array1"
	m3[[2]int{9, 7}] = "array2"
	fmt.Printf("m3: %v\n", m3)
	a := [3]int{2, 3, 4}
	b := [3]int{2, 3, 4}
	equal := a == b
	fmt.Printf("equal: %v\n", equal)

	// 布尔值作为键
	mapBool := map[bool]string{
		true:  "true",
		false: "false",
	}
	fmt.Printf("mapBool: %v\n", mapBool)
}

/*
mapInt: map[1:one 2:two 3:three]
m1: map[one:1 two:2]
m2: map[1.2:false 3.14:true]
m3: map[[1 2]:array1 [9 7]:array2]
equal: true
mapBool: map[false:false true:true]
*/
