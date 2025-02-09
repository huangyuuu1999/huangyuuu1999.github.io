package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{4,5,6,7,0,1,2}
	n := len(nums)
    a := sort.Search(n, func(i int) bool {
        if i == 0 {
            return false
        }
        return nums[i] < nums[i-1]
    })
    fmt.Printf("a: %v\n", a)
}