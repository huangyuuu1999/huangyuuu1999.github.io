package main

import (
	"fmt"
	"unsafe"
)

type Player struct {
	id       string
	winRound int
}

func main() {
	var i int
	fmt.Printf("i: %v\n", i) //  0

	var s string
	fmt.Printf("s: %v\n", s) // ""

	var b bool
	fmt.Printf("b: %v\n", b) // false

	var f func(int) int
	fmt.Printf("f: %v\n", f) // f: <nil>

	var sl []int
	fmt.Printf("sl: %v\n", sl) // sl: [], 打印出来是[], 但是是nil
	if sl == nil {
		fmt.Println("sl是nil")
	}

	confusingSlice()
	testStruct()
	testSize()
}

func confusingSlice() {
	var s1 []int
	fmt.Printf("s1: %v\n", s1) // s1: []

	s2 := []int{}
	fmt.Printf("s2: %v\n", s2) // s2: [] 打印出来都是[] 但是二者是完全不同的
}

func testStruct() {
	var p1 Player
	fmt.Printf("p1: %v\n", p1)
}

func testSize() {
	s := []int{1, 2, 3, 4, 5}
	println("slice struct size: ", unsafe.Sizeof(s))
	var po unsafe.Pointer
	println(unsafe.Sizeof(po))
}
