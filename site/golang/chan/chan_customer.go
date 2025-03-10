// chan_customer.go
// 生产者消费者模型

package main

import (
	"fmt"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	// jobs是一个只读管道 results是只写管道
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		// do some cal jobs
		results <- j * 2 - 1
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// 开启三个消费者
	for w := 1; w<=3; w++ {
		go worker(w, jobs, results)
	}
	for j := 1; j<=9; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= 9; a++ {
		<-results
	}
}
// 在这个例子中，我们使用通道来协调三个 worker goroutine 之间的任务处理。
// 每个 worker goroutine 从 jobs 通道中获取任务，
// 并将处理结果发送到 results 通道中。
// 主函数负责将所有任务发送到 jobs 通道中，并等待所有任务的结果返回。