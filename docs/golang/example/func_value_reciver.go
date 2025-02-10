// 指针接收者 和 值接受者
// func_value_reciver.go
package main

import "fmt"

type User struct {
	name string
	age  int
}

// 值接受者的函数
func (u User) selfInfo() {
	fmt.Printf("u.name: %v\n", u.name)
	fmt.Printf("u.age: %v\n", u.age)
}

func main() {
	user := &User{"小苗", 23}
	user.selfInfo() // 这里的user是指针，但是可以调用selfInfo
}
