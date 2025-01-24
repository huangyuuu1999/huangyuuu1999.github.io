package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "你6"
	l := len(s)
	fmt.Printf("l: %v\n", l) // 应该是4
	rn := utf8.RuneCountInString(s)
	fmt.Printf("rn: %v\n", rn)
	for i := range l {
		fmt.Printf("%b \n", s[i])
	}

	s1 := "👍🏻" // 这个就含有两个 unicode 代码点
	fmt.Printf("s1: %v\n", s1)
	l1 := len(s1)
	fmt.Printf("l1: %v\n", l1) // 8
	rn1 := utf8.RuneCountInString(s1) // 2
	fmt.Printf("rn1: %v\n", rn1)
}