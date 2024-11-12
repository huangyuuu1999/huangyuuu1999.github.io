package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	for i:=range 10{
		go fmt.Println(i) 
		time.Sleep(time.Millisecond)
	}
	time.Sleep(time.Millisecond) // Println们: 太好啦，是sleep函数，我们有救了！
	fmt.Println("end")
}

// start
// 0
// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
// end