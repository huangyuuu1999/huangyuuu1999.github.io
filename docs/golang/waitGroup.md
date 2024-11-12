# WaitDGroup

### WaitGroup是什么
sync.WaitGroup是sync包下提供的一个结构体
```go
// A WaitGroup waits for a collection of goroutines to finish.
// The main goroutine calls Add to set the number of
// goroutines to wait for. Then each of the goroutines
// runs and calls Done when finished. At the same time,
// Wait can be used to block until all goroutines have finished.
//
// A WaitGroup must not be copied after first use.
//
// In the terminology of the Go memory model, a call to Done
// “synchronizes before” the return of any Wait call that it unblocks.
type WaitGroup struct {
	noCopy noCopy

	state atomic.Uint64 // high 32 bits are counter, low 32 bits are waiter count.
	sema  uint32
}
```
它用于跟踪一组 goroutine 的状态，允许主程序等待这些 goroutine 完成。WaitGroup 通过计数器来实现，计数器的值表示正在运行的 goroutine 的数量。

这个结构体上面就只是实现了三个方法 Add，Done，Wait


程序开始时调用Add初始化计数，每当一个协程执行完毕时调用Done，计数就-1，直到减为0，而在此期间，主协程调用Wait 会一直阻塞直到全部计数减为0，然后才会被唤醒

这是一个简单的例子
```go
func testWaitGroup() {
	var wait sync.WaitGroup
	wait.Add(1) // 指定 需要等待 两个协程完成
	go func() {
		fmt.Println("一个协程执行完了")
		wait.Done()
	}()
	fmt.Println("我是主进程 我现在使用WaitGroup等待 *所有的子进程* 运行结束 再继续往下运行")
	wait.Wait()
	fmt.Println("所有的子进程执行完了。")
}
```

### WaitGroup怎么使用
在协程的开启的介绍文章中，很可能会出现这个例子：
```go
func main() {
	fmt.Println("start")
	for i := 0; i < 10; i++ {
		go fmt.Println(i)
	}
    // 暂停1ms
	time.Sleep(time.Millisecond)
	fmt.Println("end")
}
```
这个例子有一个问题，就是，你不知道主线程结束之后，会有几个协程运行结束，你也不知道这些协程运行的先后顺序。现在有了waitGroup，就有了解决的转机。

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	fmt.Printf("wg: %v\n", wg)
	fmt.Printf("wg: %T\n", wg)
	testWaitGroup()
	testOldQuestion()
	testWorker()
}

/*
wg: {{} {{} {} 0} 0}
wg: sync.WaitGroup
*/

func testWaitGroup() {
	var wait sync.WaitGroup
	wait.Add(1) // 指定 需要等待 两个协程完成
	go func() {
		fmt.Println("一个协程执行完了")
		wait.Done()
	}()
	fmt.Println("我是主进程 我现在使用WaitGroup等待 *所有的子进程* 运行结束 再继续往下运行")
	wait.Wait()
	fmt.Println("所有的子进程执行完了。")
}

func testOldQuestion() {
	var mainWait sync.WaitGroup // 用来控制 10个协程结束，主进程才结束
	var wait sync.WaitGroup     // 用来控制，一个add运行完，在运行下一个add
	// 计数10
	mainWait.Add(10) // 有10个协程需要等待
	fmt.Println("start")
	for i := 0; i < 10; i++ {
		wait.Add(1) // 需要等待一个协程完毕，再开启下一个协程，实际上变成串行了
		go func() {
			fmt.Println(i)
			wait.Done()
			mainWait.Done()
		}()
		// 等待当前循环的协程执行完毕
		wait.Wait()
	}
	// 等待所有的协程执行完毕
	mainWait.Wait()
	fmt.Println("end")
}

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // 在函数结束时调用 Done()
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second) // 模拟工作
    fmt.Printf("Worker %d done\n", id)
}

func testWorker() {
	// 注意这个函数 只用了一个waitgroup，所以跟上面的函数不一样，
	// 只会保证运行完5个，但是5个谁快谁慢不保证
    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1) // 增加计数器, 不像上面的test一样，一次性增加10个，也是可以的
        go worker(i, &wg) // 启动 goroutine
    }

    wg.Wait() // 等待所有 goroutine 完成
    fmt.Println("All workers done")
}
```

### WaitGroup的使用场景

- 并发任务管理：当你需要并发执行多个任务，并且希望在所有任务完成后再继续执行后续操作时，WaitGroup 是一个理想的选择。
- 并行处理：在处理大量数据时，可以将数据分成多个部分并行处理，使用 WaitGroup 等待所有处理完成。
- 网络请求：在发起多个并发的网络请求时，可以使用 WaitGroup 等待所有请求完成，确保在处理响应时所有请求都已完成。
- 资源清理：在程序结束时，确保所有 goroutine 完成后再进行资源清理或关闭操作。

### WaitGroup 的陷阱
WaitGroup通常适用于可动态调整协程数量的时候，例如事先知晓协程的数量，又或者在运行过程中需要动态调整。
WaitGroup的值不应该被复制，复制后的值也不应该继续使用，尤其是将其作为函数参数传递时，因该传递指针而不是值。倘若使用复制的值，计数完全无法作用到真正的WaitGroup上，这可能会导致主协程一直阻塞等待，程序将无法正常运行。

```go
package main

import (
	"fmt"
	"sync"
)

func hello(wait sync.WaitGroup) {
	fmt.Println("hello")
	wait.Done() // 错误
}

func main() {
	var mainWait sync.WaitGroup
	mainWait.Add(1)
	hello(mainWait)
	mainWait.Wait()
	fmt.Println("end")
}
/*
错误提示所有的协程都已经退出，但主协程依旧在等待，
这就形成了死锁，
因为hello函数内部对一个形参WaitGroup调用Done并不会作用到原来的mainWait上，
所以应该使用指针来进行传递。
*/
```