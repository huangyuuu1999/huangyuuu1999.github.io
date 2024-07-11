package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	b := a[5:] // panic: runtime error: slice bounds out of range [5:4]
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)

	inorder := []int{3}
	right_part_inorder := inorder[1:]
	fmt.Printf("inorder: %v\n", inorder)
	fmt.Printf("right_part_inorder: %v\n", right_part_inorder)
	// inorder: [3]
	// right_part_inorder: []
}