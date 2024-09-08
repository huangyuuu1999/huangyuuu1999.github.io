package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{3, 2, 4, 1, 5}
	sort.Ints(a)
	fmt.Printf("a: %v\n", a)
}