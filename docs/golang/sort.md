# 排序API

## 内置数据结构

### 对各类切片排序
这里只是演示了整数切片和字符串切片的排序，使用`sort.Ints`和`sort.Strings`两个函数。

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b1"}
	sort.Strings(strs)
	fmt.Printf("strs: %v\n", strs)

	ints := []int{7, 2, 5, 0}
	sort.Ints(ints)
	fmt.Printf("ints: %v\n", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Printf("Is ints slice sorted?: %v\n", s)
	
	s = sort.StringsAreSorted(strs)
	fmt.Printf("Is ints slice sorted?: %v\n", s)
}
```

## 自定义排序

默认的排序方法是升序的（n increasing order）。

有时候我们可能需要按照别的方式来排序

1. 整数切片降序
2. 根据字符串的长度升序|降序

sort.Slice()
