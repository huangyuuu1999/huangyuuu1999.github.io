错误处理

error就是一个定义好的接口
```go
type error interface {
	Error() string
}
```
这个接口只需要实现一个函数，就是Error函数
Error函数不需要参数并且返回一个字符串

```go
package main

import (
	"fmt"
	"strings"
)

type someStructWithErrorFunc struct {
	msg    string
	repeat int
}

/*这个结构体实现了Error这个函数，那么这个结构体就已经是一个合格的error接口，这个接口就可以被当做错误类型来使用*/
func (sswe *someStructWithErrorFunc) Error() string {
	// repeat msg repeat times.
	return strings.Repeat(sswe.msg, sswe.repeat)
}

func someFuncUseError(num int) (int, error) {
	if num < 0 {
		return 0, &someStructWithErrorFunc{"should bigger than zero", 1}
	}
	return 2 * num, nil
}

func main() {
	obj := someStructWithErrorFunc{
		msg:    "msg.",
		repeat: 3,
	}
	ans := obj.Error()
	fmt.Println(ans)

	res, err := someFuncUseError(2)
	if err != nil {
		fmt.Println(err.Error())
	}
	res, err = someFuncUseError(-1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("res =", res)
}

```
为什么上面的代码里面，返回的时候要返回 &someStructWithErrorFunc{"should bigger than zero", 1}？
因为实现Error函数的就是 *someStructWithErrorFunc 而不是 someStructWithErrorFunc

看下面的代码
```go
package main

import (
	"fmt"
)

type duck interface {
	Walk(int)
	Hello(int)
	Error() string
}

type someAnimal struct {
	name string
	age  int
}

func (a someAnimal) Walk(num int) {
	fmt.Println("walk", num, "miles.")
}

func (a someAnimal) Hello(num int) (int, error) {
	if num < 0 {
		// 这个结构体实现了Error函数，所以符合error接口，并且实现Error的就是结构体而不是结构体指针
		return -1, someAnimal{
			name: "wrong",
			age:  -1,
		}
	}
	for _ = range num {
		fmt.Println("hello I'am ", a.name, "I'am ", a.age, "years old")
	}
	return num, nil
}

func (a someAnimal) Error() string {
	return "something went wrong!"
}

func main() {
	a := someAnimal{
		name: "kiki",
		age:  5,
	}
	a.Walk(2)
	a.Hello(3)
}

```
