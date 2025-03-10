# chan
管道

## 使用无缓冲chan

- chan 要使用 make来创建，make(chan, int) 表示创建一个无缓冲的 channel，能够传输的类型是int。
- 无缓冲的效果

下面的例子就是无缓冲的阻塞效果，导致死锁。

```go
// chan_start.go  直接deadlock
// • 向无缓冲 channel 写数据，如果读协程没有准备好，会阻塞
package main

import "fmt"

func main() {
	// 创建了一个管道，能传输int数据
	ch := make(chan int)
	ch <- 1 // 这是无缓冲的chan，没有人来取的话，我这个goroutine就会阻塞
	x := <-ch
	fmt.Printf("x: %v\n", x)
}
```

开一个协程来向chan发数据，在主协程收数据。
```go
// chan_simple_usage.go
// 演示 chan 的简单使用场景

package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		ch <- "hello"
	}()

	word := <-ch
	fmt.Printf("word: %v\n", word)
}
```
也可以直接丢弃 <-ch

```go
// chan_simple_usage.go
// 演示 chan 的简单使用场景

package main

import "fmt"

func main() {
	dropData()
}

func dropData() {
	ch := make(chan bool)
	go func() {
		ch <- true
	}()
	<-ch // 直接丢弃数据
	fmt.Println("dropped")
}

```

无缓冲的chan，读写都会阻塞的
无缓冲 channel 在读和写的过程中是都会阻塞，由于阻塞的存在，所以使用 channel 时特别注意使用方法，防止死锁和协程泄漏的产生

```go
// fatal error: all goroutines are asleep - deadlock!
// chan_simple_usage.go
// 演示 chan 的简单使用场景

package main

import "fmt"

func main() {
	readBlock()
}

func readBlock() {
    // 从无缓冲 channel 读数据，如果写协程没有准备好，会阻塞
	ch := make(chan float64)
	<-ch
	go func() {
		ch <- 3.14
	}()
	fmt.Println("yes")
}
```

## 有缓冲的 chan
```go
// chan_buffer.go
// 有缓冲的 chan
package main

import "fmt"

func main() {
	chanWithBuffer()
}

func chanWithBuffer() {
	ch := make(chan int, 3)
	for i := range 5 {
		go func() {
			ch <- i
		}()
	}

	for _ = range 5 {
		// 输出的顺序是不固定的，不知道哪个协程先写
		fmt.Println(<-ch)
	}
}

```

读有缓冲的chan，但是没有数据发送方，也会阻塞
```go
// chan_buffer.go
// 有缓冲的 chan
package main

import "fmt"

func main() {
	chanWithBuffer()
	// blockOrNot()
	blockUntilSend()
}

func chanWithBuffer() {
	ch := make(chan int, 3)
	for i := range 5 {
		go func() {
			ch <- i
		}()
	}

	for _ = range 5 {
		// 输出的顺序是不固定的，不知道哪个协程先写
		fmt.Println(<-ch)
	}
}

// 如果从有缓冲区的管道中读，会阻塞吗？
// fatal error: all goroutines are asleep - deadlock!
func blockOrNot() {
	ch := make(chan int, 3)
	x := <-ch
	fmt.Printf("x: %v\n", x)
}

func blockUntilSend() {
	ch := make(chan string, 3)
	for _ = range 3 {
		go func() {
			ch <- "不要回答"
		}()
	}
	x := <-ch // 只读取1次，管道内剩余的没被处理
	fmt.Printf("x: %v\n", x)
}
```


## 阻塞

Channel 各种操作导致阻塞和协程泄漏的场景

写操作，什么时候会被阻塞？

• 向 nil 通道发送数据会被阻塞
• 向无缓冲 channel 写数据，如果读协程没有准备好，会阻塞
• 无缓冲 channel ，必须要有读有写，写了数据之后，必须要读出来，否则导致 channel 阻塞，从而使得协程阻塞而使得协程泄漏
• 一个无缓冲 channel，如果每次来一个请求就开一个 go 协程往里面写数据，但是一直没有被读取，那么就会导致这个 chan 一直阻塞，使得写这个 chan 的 go 协程一直无法释放从而协程泄漏。
• 向有缓冲 channel 写数据，如果缓冲已满，会阻塞
• 有缓冲的 channel，在缓冲 buffer 之内，不读取也不会导致阻塞，当然也就不会使得协程泄漏，但是如果写数据超过了 buffer 还没有读取，那么继续写的时候就会阻塞了。如果往有缓冲的 channel 写了数据但是一直没有读取就直接退出协程的话，一样会导致 channel 阻塞，从而使得协程阻塞并泄漏。
读操作，什么时候会被阻塞？

• 从 nil 通道接收数据会被阻塞
• 从无缓冲 channel 读数据，如果写协程没有准备好，会阻塞
• 从有缓冲 channel 读数据，如果缓冲为空，会阻塞
close 操作，什么时候会被阻塞？

• close channel 对 channel 阻塞是没有任何效果的，写了数据但是不读，直接 close，还是会阻塞的。

## close chan
关闭chan

## chan range
使用range 处理切片，可以依次得到切片的值。那么range chan会怎样呢？

1. 通道必须关闭：使用 range 遍历时，通道必须被关闭，否则循环可能永远不会结束。
2. 避免死锁：如果通道未关闭且没有值发送，range 会阻塞，可能导致程序无法正常退出。
3. 通道关闭后不能发送值：一旦通道被关闭，再次向通道发送值会导致运行时错误。
总之，range 是一种方便的方式，用于从通道中接收值，直到通道关闭。
```go
// chan_range.go
// 使用 range 依次从chan取数据

package main

import (
	"fmt"
	"time"
)

func main() {
	withClose()
	withoutClose()
}

func withClose() {
	// 创建一个通道
	ch := make(chan int)

	// 启动一个协程向通道发送值
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(1 * time.Second) // 模拟耗时操作
		}
		close(ch) // 关闭通道
	}()

	// 使用 range 遍历通道
	for v := range ch {
		fmt.Println("Received:", v)
	}
	fmt.Println("Channel closed, exiting.")
}

func withoutClose() {
	// 创建一个通道
	ch := make(chan int)

	// 启动一个协程向通道发送值
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(1 * time.Second) // 模拟耗时操作
		}
	}()

	for v := range ch { // 要读 但是没人写，会死锁
		fmt.Println("Received:", v)
	}
	fmt.Println("Channel closed, exiting.")
}

```

### 多个goroutine range同一个chan？

## 单向的chan

## chan当做参数传递

## 使用场景

### 生产者消费者
这里是一个示例程序，忽略了一些错误和极端的情况.
```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// 统计字节频率
func Counter(filePath string) [256]int {
	bytes, _ := os.ReadFile(filePath)
	ans := [256]int{}
	for i := 0; i < len(bytes); i++ {
		ans[bytes[i]]++
	}
	return ans
}

func worker(jobs <-chan []byte, results chan<- [256]int, wg *sync.WaitGroup) {
	for bs := range jobs {
		ans := [256]int{}
		for i := 0; i < len(bs); i++ {
			ans[bs[i]]++
		}
		results <- ans
		wg.Done()
	}
}

func Counter2(filePath string) [256]int {
	bufferSize := 1024 * 1024 * 1

	info, _ := os.Stat(filePath)
	chanCnt := info.Size()/int64(bufferSize) + 1

	jobs, results := make(chan []byte, chanCnt), make(chan [256]int, chanCnt)
	var wg sync.WaitGroup

	for _ = range 8 {
		go worker(jobs, results, &wg)
	}
	file, _ := os.Open(filePath)
	defer file.Close()

	reader := bufio.NewReader(file)

	buffer := make([]byte, bufferSize)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		// 复制数据到新切片以避免被覆盖
		chunk := make([]byte, n)
		copy(chunk, buffer[:n])
		jobs <- chunk // 发送正确的数据块
		wg.Add(1)
	}
	close(jobs)

	wg.Wait()
	close(results)

	ans := [256]int{}
	for m := range results {
		for k, v := range m {
			ans[k] += v
		}
	}
	return ans
}

func main() {
	filePath := "三体.txt"

	// 测试 Counter 方法
	start := time.Now()
	result1 := Counter(filePath)
	duration1 := time.Since(start)
	fmt.Println("Counter result:", result1)
	fmt.Printf("Counter took %v\n", duration1)

	// 测试 Counter2 方法
	start = time.Now()
	result2 := Counter2(filePath)
	duration2 := time.Since(start)
	fmt.Println("Counter2 result:", result2)
	fmt.Printf("Counter2 took %v\n", duration2)
}

```
