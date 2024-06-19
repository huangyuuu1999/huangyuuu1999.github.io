python的变量，在使用前不需要先声明，在循环的时候，可以直接使用变量i来承接迭代器给出的值

```python
i = 9
for i in range(5):
    print('x')
print(i) # i = ?
```

在golang中的表现
```golang
package main

import "fmt"

func main() {
    x := []int{1, 2, 3, 4}
    for i := range(x) {
        fmt.Println("x", i)
    }
    // fmt.Printf("i: %v\n", i) // undefined: i
}
```

这段代码展示了一个事实：循环内部的i和外部定义的变量i是无关的，这表明go的循环内部变量是有作用域的
```golang
package main

import "fmt"

func main() {
    i := 20
    x := []int{1, 2, 3, 4}
    for i := range(x) {
        fmt.Println("x", i)
    }
    fmt.Printf("i: %v\n", i) // i: 20
}
```