// defer_panic.go
// panic导致defer链表执行，但是defer里面没有recover的情形

package main

import (
	"fmt"
	"unicode/utf8"
)


func main() {
	panicFireDefer()
}

func panicFireDefer() {
	defer fmt.Println("1")
	defer fmt.Println(utf8.RuneCountInString("👍🏿")) // 2
	defer fmt.Println(3)
	panic("some problem happened T^T...")
}
