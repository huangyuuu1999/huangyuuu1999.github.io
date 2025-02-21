// recover_panic.go
// 展示recover的作用

package main

import "fmt"

func main() {
	A()
}

func A() {
	defer A1()
	defer A2()
	panic("panicA")
}

func A1() {
	fmt.Println("A1")
}

func A2() {
	p := recover()
	fmt.Printf("p: %v\n", p)
}