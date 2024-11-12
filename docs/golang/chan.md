## 使用go的协程

- go中的协程，对比了go的协程和py的协程的表现行为；看看协程的控制，使用Sleep来让协程运行。最后引出了协程的控制问题：后面会了解的管道，锁和Group。

### 使用 go 关键字开启协程
go开启协程只需要一个关键字go，不像其他语言中可能需要async await
对于func定义的函数，都可以在前面加一个go关键字来开启协程


```go
// 1.演示如何开启一个协程 go_routine.go
package main

import "fmt"


func test_go_routine() {
	fmt.Println("hello goroutine!")
}

func main() {
    // 只需要在函数调用前面加一个go关键字即可
	go test_go_routine()
}
```
运行上面的代码会发现，极大概率没有输出任何内容，这是因为，go关键字只是用来开启一个协程，而并不是像await那样，等待一个协程运行完毕。
```python
# 演示python的协程和go的协程的区别 go_routine.py
import asyncio

async def test_go_routine():
    print("hello goroutine!")
    return 42


async def main():
    res = await test_go_routine()
    print("in main: ", res)

if __name__ == '__main__':
    asyncio.run(main())

# hello goroutine!
# in main:  42
```
实际上上面的go代码中的协程，没有开始运行，主进程就已经结束了。
那么要如何让上面的go代码中的协程运行？只需要在主函数里面想办法“等一段时间”，例如使用time.Sleep函数

```go
// 演示 主进程阻塞 以等待协程运行完毕
package main

import (
	"fmt"
	"time"
)


func test_go_routine() {
	fmt.Println("hello goroutine!")
}

func main() {
	go test_go_routine() // 只需要在函数调用前面加一个go关键字即可
	time.Sleep(time.Second) // 主进程阻塞1s，以便于让协程完成运行
}
```

### 开启协程的更多方式
go的后面必须是一个函数的调用，例如下面的立即执行函数的方式
```go
// 使用匿名函数立即执行来开启协程
package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("go + anonymous func()")
	}()
	time.Sleep(time.Second)
}

```
但是不可以是有返回值的内置函数
```go
// 演示具有返回值的内置函数，不可以用go开启协程
package main

import "time"

func main() {
	// go discards result of make([]int, 10) (value of type []int)
	go make([]int, 10)

	// Sleep has no return value, so it's ok
	go time.Sleep(time.Second)
}
```
下面的方式开启协程都是可以的
```go
// 使用3种方式来开启协程
func main() {
	go fmt.Println("hello world!")
	go hello()
	go func() {
		fmt.Println("hello world!")
	}()
}

func hello() {
	fmt.Println("hello world!")
}
```
以上三种开启协程的方式都是可以的，但是其实这个例子执行过后在大部分情况下什么都不会输出，协程是并发执行的，系统创建协程需要时间，而在此之前，主协程早已运行结束，一旦主线程退出，其他子协程也就自然退出了。

### 协程开启后，万类霜天竞自由
一旦协程被go关键字激活，他就会在合适的时机运行，但也有可能不运行（在限定的时间内），例如下面的例子，甚至可能出现这样的结果
```shell
start
end
2
```
也并不意味着先激活的协程就一定先运行。
```golang
// 演示协程开启之后，是否运行，哪一个先运行，是不可预知的
package main

import "fmt"

func main() {
	fmt.Println("start")
	for i:=range 10{
		go fmt.Println(i) 
	}
	fmt.Println("end")
}
```
这是一个在循环体中开启协程的例子，永远也无法精准的预判到它到底会输出什么。可能子协程还没开始运行，主协程就已经结束了
又或者只有一部分子协程在主协程退出前成功运行。
那么我们就让主进程来等一等吧，这样所有的协程就都能执行完毕了，不是吗？
```go
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
// ...
// 3
// 2
// end
```
当把Sleep的时间设置的足够的时候，就可以保证所有的协程都可以完成了。但是注意输出的内容不是有序的哦。即协程什么时候执行，并不能由程序员来安排，而是由GPM模型自行调度。

我们可以先开启一个，等一段时间，再开启第二个，这样的方式，让先开启的尽可能先被处理。
```go
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
// ...
// 9
// end
```

上面的例子中结果输出很完美，那么并发的问题解决了吗，不，一点也没有。对于并发的程序而言，不可控的因素非常多，执行的时机，先后顺序，执行过程的耗时等等，倘若循环中子协程的工作不只是一个简单的输出数字，而是一个非常巨大复杂的任务，耗时的不确定的，那么依旧会重现之前的问题。

```go
package main

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

// 现实中的 协程 运行的时间不会很固定（一定是1ms），可能是完全不确定的

func main() {
	fmt.Println("start")
	for i := 0; i < 10; i++ {
		go hello(i)
		time.Sleep(time.Millisecond)
	}
	time.Sleep(time.Millisecond)
	fmt.Println("end")
}

func hello(i int) {
	// 模拟随机耗时
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	fmt.Println(i)
}

```
因此time.Sleep并不是一种良好的解决办法.

### 并发控制方式？
幸运的是Go提供了非常多的并发控制手段，常用的并发控制方法有三种：

channel：管道
WaitGroup：信号量
Context：上下文

三种方法有着不同的适用情况，WaitGroup可以动态的控制一组指定数量的协程，Context更适合子孙协程嵌套层级更深的情况，管道更适合协程间通信。对于较为传统的锁控制，Go也对此提供了支持：

Mutex：互斥锁
RWMutex ：读写互斥锁