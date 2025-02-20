// defer_sequence.go
// 演示多个 defer 的执行顺序

package main

import "fmt"

func main() {
	someFunc()
}

func someFunc() {
	defer func() { fmt.Println("1") }()
	defer func() { fmt.Println("2") }()
	defer func() { fmt.Println("3") }()
	a := 1
	a++
	var m map[int]string
	b, e := m[666]
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v e: %v \n", b, e)
}
