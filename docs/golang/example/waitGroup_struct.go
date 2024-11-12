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