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
