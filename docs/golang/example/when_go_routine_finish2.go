package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	for i:=range 10{
		go fmt.Println(i) 
	}
	time.Sleep(time.Millisecond) // Println们: 太好啦，是sleep函数，我们有救了！
	fmt.Println("end")
}
// start
// 1
// 9
// 0
// 4
// 5
// 6
// 7
// 8
// 3
// 2
// end