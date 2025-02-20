// panic_in_defer.go
// 如果发生了panic，触发defer，而defer里面又有panic 会怎么样呢？

package main

func main() {
	func1()
}

func func1() {
	defer func() {
		panic("panic in defer1")
	}()
	defer func() {
		panic("panic in defer2")
	}()
	panic("func1 panic")
}

// panic: func1 panic
//         panic: panic in defer2
//         panic: panic in defer1

// goroutine 1 [running]:
// main.func1.func1()
//         C:/Users/Administrator/Desktop/codes/defer/panic_in_defer.go:12 +0x25
// panic({0xe7f5e0?, 0xea8fb8?})
//         D:/golang/src/runtime/panic.go:785 +0x132
// main.func1.func2()
//         C:/Users/Administrator/Desktop/codes/defer/panic_in_defer.go:15 +0x25
// panic({0xe7f5e0?, 0xea8f98?})
//         D:/golang/src/runtime/panic.go:785 +0x132
// main.func1()
//         C:/Users/Administrator/Desktop/codes/defer/panic_in_defer.go:17 +0x4e
// main.main()
//         C:/Users/Administrator/Desktop/codes/defer/panic_in_defer.go:7 +0xf
// exit status 2