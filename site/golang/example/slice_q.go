package main

import "fmt"

func main() {
    a := []int{2, 3, 5, 7}
    b := a
    b[0] = 999
    fmt.Println(a, b) // [999 3 5 7] [999 3 5 7]
	
	a = append(a, 666)

    fmt.Println(a, b) // [999 3 5 7 666] [999 3 5 7]
	
	a[1] = 7777
    fmt.Println(a, b) // [999 7777 5 7 666] [999 3 5 7]
	test1()
}

func test1() {
	s1 := []int{1, 2}
	s2 := s1
	s2 = append(s2, 3)
	fmt.Printf("s2: %v\n", s2) // s2: [1 2 3]
	fmt.Printf("s1: %v\n", s1) // s1: [1 2]
	SliceRise(s1)
	SliceRise(s2)
	fmt.Printf("s2: %v\n", s2) // s2: [2 3 4]
	fmt.Printf("s1: %v\n", s1) // s1: [1 2]
}

func SliceRise(s []int) {
	s = append(s, 0)
	for i := range s {
		s[i] += 1
	}
}