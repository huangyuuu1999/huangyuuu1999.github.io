## Go的Mutex

### 锁的必要性
race-condition案例，多个协程读取到相同的count，先后写入相同的新值，导致结果大概率小于理想结果10。
```go
package main

import (
	"fmt"
	"sync"
)

var count uint32 = 0

func main() {
	var wait sync.WaitGroup
	for i:= range 16 {
		go func(count *uint32) {
			defer wait.Done()
			wait.Add(1)
			tmp := *count
			fmt.Printf("before task%v count: %v  ==> ", i, tmp)
			*count = tmp + 1
			fmt.Printf("after  task%v count: %v\n", i, *count)
		}(&count)
	}
	wait.Wait()
	fmt.Printf("finally count: %v\n", count)
}
```
使用默认的锁保护数据access代码
```go
package main

import (
	"fmt"
	"sync"
)

var count uint32 = 0
var lock sync.Mutex // 实现了 sync.Locker 接口的结构体

func main() {
	var wait sync.WaitGroup
	for i:= range 10 {
		go func(count *uint32) {
			defer wait.Done()
			wait.Add(1)

			lock.Lock()
			tmp := *count
			fmt.Printf("before task%v count: %v  ==> ", i, tmp)
			*count = tmp + 1
			fmt.Printf("after  task%v count: %v\n", i, *count)
			lock.Unlock()
		}(&count)
	}
	wait.Wait()
	fmt.Printf("finally count: %v\n", count)
}
```

### 不可重入锁/非递归锁


### 读写锁
