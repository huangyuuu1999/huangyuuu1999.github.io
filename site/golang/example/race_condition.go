package main

import (
	"fmt"
	"sync"
)

var count uint32 = 0

func main() {
	var wait sync.WaitGroup
	for i:= range 16 {
		go func(count *uint32) {
			defer wait.Done()
			wait.Add(1)
			tmp := *count
			fmt.Printf("before task%v count: %v  ==> ", i, tmp)
			*count = tmp + 1
			fmt.Printf("after  task%v count: %v\n", i, *count)
		}(&count)
	}
	wait.Wait()
	fmt.Printf("finally count: %v\n", count)
}