// test_rune.go
// go version go1.22.1 darwin/arm64

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s1 := "你是这个👍🏻"
	s2 := "你是这个👍"
	l1, l2 := len(s1), len(s2)
	fmt.Printf("l1: %v\n", l1)
	fmt.Printf("l2: %v\n", l2)
	r1, r2 := utf8.RuneCountInString(s1), utf8.RuneCountInString(s2)
	fmt.Printf("r1: %v\n", r1)
	fmt.Printf("r2: %v\n", r2)
}

/*
❯ go run test_rune.go
l1: 20
l2: 16
r1: 6
r2: 5
*/