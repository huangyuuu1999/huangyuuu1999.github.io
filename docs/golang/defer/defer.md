# `defer` 关键字

defer是go语言的关键字，用于在函数结束（函数调用栈，当前帧出栈，pc指向return address）之前做一些资源回收、错误处理等收尾工作。

常用于资源清理、日志记录、错误处理等场景。

## defer 的使用场景
### defer 用来关闭文件描述符
最直观的使用就是用来关闭一些打开的资源。
```go
// filename defer_basic.go
// show basic usage of defer

package main

import (
	"fmt"
	"os"
)

func main() {
	deferCloseFile()
}

func deferCloseFile() {
	file, _ := os.Open("example.txt")
	defer file.Close()
	data := make([]byte, 20)
	n, _ := file.Read(data)
	fmt.Println("Data read from file:\n", string(data[:n]))
	fmt.Printf("\n%v\n", "other code...")
}
```
### defer 记录日志
可以用于在函数退出时记录日志，无论函数是正常退出还是因为错误退出。
```go
// filename defer_basic.go
// show basic usage of defer

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	deferLog()
}

func deferLog() {
	start := time.Now()
	defer func() {
		duration := time.Since(start)
		fmt.Printf("Function took: %v\n", duration)
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("Function completed")
}
```

### defer 用来捕获错误
在defer中捕获可能的panic这是defer的一个重要的用途。
```go
// filename defer_basic.go
// show basic usage of defer

package main

import (
	"fmt"
	"os"
)

func main() {
	deferRecover()
}


func mightPanic() {
	panic("a problem")
}

func deferRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("revovered a panic: %v\n", r)
		}
	}()
	mightPanic()
	fmt.Println("after panic.")
}
```
## defer的执行顺序和时机

每一个协程都有自己的defer链表，在runtime中，协程是g struct，在g中，就保存了defer链表。
```go
type g struct {
    ...
	_panic    *_panic // innermost panic - offset known to liblink
	_defer    *_defer // innermost defer
    ...
}
```

### 多个defer LIFO
后面的defer先执行，实际上是一个defer链表，采用头插法，后面生命的defer放在最前面。
```go
// defer_sequence.go
// 演示多个 defer 的执行顺序

package main

import "fmt"

func main() {
	someFunc()
}

func someFunc() {
	defer func() { fmt.Println("1") }()
	defer func() { fmt.Println("2") }()
	defer func() { fmt.Println("3") }()
	a := 1
	a++
	var m map[int]string
	b, e := m[666]
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v e: %v \n", b, e)
}
/*
b:  e: false
3
2
1
*/
```

### defer 在 return 后执行
注意，return语句对应到汇编上是有丰富的内容的，包括：把返回值的内容放到约定的返回值寄存器（或内存某处），调整stack pointer指针为return address，等等。

实际上在go的func里面写的return，只是对应了第一步，也就是把要返回的值写到返回地址。
return关键字不是代表完整的“return”过程，而只是return的一小步。

所以defer实际上是可以操作return放置好的结果的，这是巧妙也很tricky的一点。
```go
// defer_return.go
// 演示 defer 和 return 谁先执行
package main

import "fmt"

func main() {
	// 观察先输出什么内容
	wrapper()
	ans := wrapper2()
	fmt.Printf("ans: %v\n", ans)
}

func deferFunc(i int) {
	for j := range i {
		fmt.Println(j, "defer was called.")
	}
}

func returnFunc() int {
	fmt.Println("return was called.")
	return 1
}

func wrapper() int { // 说服力有限
	defer deferFunc(3)
	return returnFunc()
}

func wrapper2() (ans int) { // 能够有力证明 defer 就是在 return语句之后 启动的
	defer func() {
		fmt.Printf("ans: %v\n", ans)
		ans += 10
	}()
	return ans + 3
}
```

## defer和panic配合
defer能执行，要么是return已经执行了。要么就是发生了panic。

发生panic时，遍历本协程的defer链表，一个个执行，如果遇到recover，就停止当前的panic，返回recover处继续往下执行。

如果没有遇到recover，就遍历完本协程的defer链表之后，向stderr抛出panic信息。

### panic 不被捕获的情况

没有revover的情况，直接执行defer链表，直达全部结束，最后抛出panic给stderr。

```go
// defer_panic.go
// panic导致defer链表执行，但是defer里面没有recover的情形

package main

import (
	"fmt"
	"unicode/utf8"
)


func main() {
	panicFireDefer()
}

func panicFireDefer() {
	defer fmt.Println("1")
	defer fmt.Println(utf8.RuneCountInString("👍🏿")) // 2
	defer fmt.Println(3)
	panic("some problem happened T^T...")
}
```

### panic 被捕获的情况

当panic触发defer，且defer中出现recover的时候，就会终止当前的panic，继续执行。

recover 相当于打断了panic过程。【panic过程是指从发生panic到进程结束抛出stderr的整个过程】

```go
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
```

## defer中有panic

defer 的回调函数内部也可以写panic，那么panic会表现出什么行为？

g协程的内部，除了维护defer链表，还维护了panic链表。defer链表的每个节点就是defer结构体，defer结构体内部定义了started表示是否开始执行了，还定义了_panic表示是哪个panic启动的这个defer。
```go
// 伪代码，仅用于说明defer内部和g内部的成员
type g struct {
	_defer // deferlist defer链表
	_panic // paniclist panic链表
} 

type Defer struct {
	started bool
	_panic // panic 表示谁启动了这个 defer
}
```
当发生panic的时候，把panic链表的头上加上当前panic，然后就开始执行defer链表，注意后注册的defer越靠近链表头，开始执行一个defer的时候，把这个defer的started设置为true，_panic设置为当前的panic。

如果现在defer中又出现了panic，那么就在panic链表的头上再插入一个新的panic节点，然后再从defer链表的头节点开始执行，当发现已经started并且_panic不是自己的时候，【这个_panic就标记为结束？】就继续往后执行下一个defer节点。

当全部的defer节点执行完毕之后，就要进行panic的打印了，此时是从g的panic链表尾部向头部打印，也就是panic出现的顺序打印的。

### defer中有panic，但是不recover
```go
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
```
可以看到三个panic都打印出来了，并且是按照先后顺序打印的，这也是协程g的panic链表的从尾到头的打印顺序。

### 嵌套 panic 测试题
看下面的代码，输出是什么？来源是 [golang中文学习文档](https://www.dstgo.cn/essential/senior/95.error.html#%E5%96%84%E5%90%8E) 的错误-panic章节。
```go
func main() {
	defer fmt.Println("A") // 1
	defer func() { //2
		func() {
			panic("panicA")
			defer fmt.Println("E")
		}()
	}()
	fmt.Println("C")
	dangerOp()
	defer fmt.Println("D")
}

func dangerOp() {
	defer fmt.Println(1) //3
	defer fmt.Println(2) //4
	panic("panicB")
	defer fmt.Println(3)
}
```
首先再main协程中注册了两个defer，然后打印出'C'，随后进入dangerOp函数内（注意还是本协程，没有开新的协程），再次注册了两个函数，注册的defer在代码中注释为 1，2，3，4。在defer链上的首尾顺序是 4321。然后panicB在main-g的panic链表挂上panicB，触发defer。输出2、1。
然后执行defer2，立马panicA，在main-g的panic链的panicB前面挂上panicA，再去执行defer链。此时注意到defer2还在链上，且状态为started，但是_panic为panicB因为是panicB触发的他，那么就跳过这个defer，再去执行defer1了，注意此时也会标记panicA为done。到最后defer全部执行完了。就输出panic信息，顺序是panicB，panicA。

所以总的答案是：
```txt
C 
2 
1 
A 
panicB 
panicA
```

### defer中有panic，进行recover
捕捉到的是后面出现的panic。

recover捕捉到的只有最后一个出现的panic。

```go
// panic_in_defer_recover.go
// defer中有panic，进行recover

package main

import "fmt"


func main() {
	func2() // 捕捉到的是: panic in defer2
}

func func2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("捕捉到的是: %v\n", err)
		}
	}()

	defer func() {
		panic("panic in defer2")
	}()

	panic("func1 panic")
}
```

## recover 的作用

recover 把当前执行的 panic 的 recovered 字段置为 true。
```go
// recover_panic.go
// 展示recover的作用

package main

import "fmt"

func main() {
	A()
}

func A() {
	defer A1()
	defer A2()
	panic("panicA")
}

func A1() {
	fmt.Println("A1")
}

func A2() {
	p := recover()
	fmt.Printf("p: %v\n", p)
}
```
当执行到A的panicA时，g的defer链表挂着A1，A2，panic链表挂上panicA，然后开始执行defer，先执行A2，recover把当前的panicA设置为recovered，继续往下执行，当每个defer执行完了以后，都会检查当前正在执行的panic是否被恢复了，如果已经recovered，就把他从链表中移除。