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