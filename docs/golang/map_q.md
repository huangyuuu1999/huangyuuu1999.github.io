# 试图访问map中不存在的键会怎样

### 默认返回零值
Go的map是引用类型，并且是无序的。如果你尝试访问一个不存在的键，你将得到该类型的零值（例如，对于整数类型是0，对于字符串类型是""）。以下是Go中使用map的一个例子：

```golang
package main

import "fmt"

func main() {
    m := map[string]int{
        "apple": 1,
        "banana": 2,
    }

    // 访问存在的键
    fmt.Println(m["apple"]) // 输出: 1

    // 访问不存在的键，将得到int类型的零值0
    fmt.Println(m["orange"]) // 输出: 0
}
```
### 和python defaultdict的相似和不同

在go中，你甚至可以对map里不存在的键进行运算
```go
package main

import "fmt"


func main() {
	m := map[int]int{1: 2, 3: 9}
	fmt.Printf("m[5]: %v\n", m[5]) // m[5]: 0

	// 对于不存在于m的键，你甚至可以对它进行++操作
	m[4]++
	fmt.Printf("m: %v\n", m) // m: map[1:2 3:9 4:1]

	
	// 可以检查 某个元素是否存在于map中
	if val, exist := m[8]; exist {
		fmt.Println("8存在", val)
		} else {
			fmt.Println("8不存在") // 执行这一行
		}
		
	var sz = len(m)
	fmt.Printf("sz: %v\n", sz)
	// 此时可以 m[8]++, 然后8就存在了
	m[8]++
	if val, exist := m[8]; exist {
		fmt.Println("8存在", val) // 执行这一行
	} else {
		fmt.Println("8不存在") 
	}

	sz = len(m)
	fmt.Printf("sz: %v\n", sz)
}
```