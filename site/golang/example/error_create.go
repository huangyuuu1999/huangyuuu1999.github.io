package main

import (
	"errors"
	"fmt"
)

func positiveSum(upper int) (int, error) {
	if upper <= 0 {
		return 0, errors.New("必须是正整数")
	}
	sum := 0
	for i := 0; i <= upper; i++ {
		sum += i
	}
	return sum, nil
}

func testErrorString() {
	a := errors.New("错误") // 这里返回的就是一个 errorString结构体，这个结构体实现了Error方法
	fmt.Printf("a: %v\n", a)
}

func testWrapError() {
	err := errors.New("这是一个原始错误")
	wrapErr := fmt.Errorf("错误，%w", err)

	fmt.Printf("wrapErr: %v\n", wrapErr)
	fmt.Printf("wrapErr: %T\n", wrapErr)
}

func main() {
	res, err := positiveSum(5)
	if err != nil {

		fmt.Println(err.Error())
	}
	fmt.Println(res)
	res, err = positiveSum(-2)
	if err != nil {

		fmt.Println(err.Error())
	}

	testErrorString()
	testWrapError()
	// wrapErr: 错误，这是一个原始错误
	// wrapErr: *fmt.wrapError
}
