package main

import (
	"fmt"
	"time"
)


func main() {
	// After(d time.Duration) <-chan time.Time
	timeCh := time.After(time.Second)
	ans := <-timeCh
	fmt.Printf("ans: %v\n", ans)
	fmt.Printf("ans: %T\n", ans)  // Time.time
}