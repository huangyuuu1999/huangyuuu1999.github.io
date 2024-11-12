package main

import (
	"fmt"
	"sync"
)

var count uint32 = 0
var lock sync.Mutex // 实现了 sync.Locker 接口的结构体

func main() {
	// testSlice()

	var wait sync.WaitGroup
	for i:= range 10 {
		go func(count *uint32) {
			defer wait.Done()
			wait.Add(1)

			lock.Lock()
			tmp := *count
			fmt.Printf("before task%v count: %v  ==> ", i, tmp)
			*count = tmp + 1
			fmt.Printf("after  task%v count: %v\n", i, *count)
			lock.Unlock()

		}(&count)
	}
	wait.Wait()
	fmt.Printf("finally count: %v\n", count)
}

func testSlice() {
	var a []int
	fmt.Printf("a: %v\n", a)
	a = append(a, 2)
	fmt.Printf("a: %v\n", a)
}