golang的闭包函数，可以访问到相同地位的兄弟变量，同一作用域的其他变量

```go
func innerfunc2() {
	x := 20
	this_func_can_visit_x := func() {
		fmt.Println("this_func_can_visit_x", x)
	}
	this_func_can_visit_x() // this_func_can_visit_x 20
}
```