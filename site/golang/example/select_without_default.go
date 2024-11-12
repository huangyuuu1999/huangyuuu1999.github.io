package main

import "fmt"

/*
select 等待从多个管道读取数据；没有default，并且管道都没有数据
*/

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	select {
	case n, k := <- ch1:
		fmt.Printf("n: %v k: %v\n", n, k)
	case n, k := <- ch2:
		fmt.Printf("n: %v k: %v\n", n, k)
	}
} // fatal error: all goroutines are asleep - deadlock!