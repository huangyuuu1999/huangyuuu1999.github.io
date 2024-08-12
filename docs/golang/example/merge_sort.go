package main

import "fmt"

func partition(nums []int) {
	pivot := nums[0]
	l, r := 0, len(nums)-1
	for l < r {
		for l < r && nums[l] > pivot {
			r--
		}
		nums[l] = nums[r]
		l++
		for l < r && nums[l] <= pivot {
			l++
		}
		nums[r] = nums[l]
		r--
	}
	nums[l] = pivot
}

func test_p() {
	a := []int{6, 3, 5, 7, 8, 1, 2, 0, 9}
	fmt.Printf("a: %v\n", a)
	partition(a)
	fmt.Printf("a: %v\n", a)
}

func main() {
	test_p()
}
