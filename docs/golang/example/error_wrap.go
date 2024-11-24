package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("原始的错误")
	// 这里必须使用 %w 这个占位符
	wrapErr := fmt.Errorf("包装错误 %w", err)
	// wrapErr: *fmt.wrapError
	// wrapErr: 包装错误 原始的错误
	fmt.Printf("wrapErr: %T\n", wrapErr)
	fmt.Printf("wrapErr: %v\n", wrapErr)

	// res: 原始的错误
	// res: *errors.errorString
	res := errors.Unwrap(wrapErr)
	fmt.Printf("res: %v\n", res)
	fmt.Printf("res: %T\n", res)
}
