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
