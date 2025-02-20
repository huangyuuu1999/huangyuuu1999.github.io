// panic_in_defer_recover.go
// defer中有panic，进行recover

package main

import "fmt"


func main() {
	func2() // 捕捉到的是: panic in defer2
}

func func2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("捕捉到的是: %v\n", err)
		}
	}()

	defer func() {
		panic("panic in defer2")
	}()

	panic("func1 panic")
}