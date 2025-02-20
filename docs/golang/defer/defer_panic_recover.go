// defer_panic_recover.go
// panic触发了defer，而defer中有recover

package main

import "fmt"

func main() {
	deferWithRecover()
}

func deferWithRecover() {
	defer func() {
		fmt.Println("尝试捕获panic")
		if err := recover(); err != nil {
			fmt.Printf("err: %v\n", err)
		}
		fmt.Println("会继续执行吗？")
	}()
	defer fmt.Println("这个defer不捕捉panic")
	panic("出错了😇")
	fmt.Println("这里不会执行")
}