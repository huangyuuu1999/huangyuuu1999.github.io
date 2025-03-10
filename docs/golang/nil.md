# 

### 切片单独声明，值是nil

```go
package main
import "fmt"

func main() {
    var s []int
    if s == nil {
    fmt.Println("is nil")
    }
}

```

不可以直接给某个赋值为nil 会警告 untyped bil