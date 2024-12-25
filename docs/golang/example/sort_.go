package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b1"}
	sort.Strings(strs)
	fmt.Printf("strs: %v\n", strs)

	ints := []int{7, 2, 5, 0}
	sort.Ints(ints)
	fmt.Printf("ints: %v\n", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Printf("Is ints slice sorted?: %v\n", s)
	
	s = sort.StringsAreSorted(strs)
	fmt.Printf("Is ints slice sorted?: %v\n", s)
}

func sortableInterface() {
	// sort.Interface
}