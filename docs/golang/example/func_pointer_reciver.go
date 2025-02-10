// 指针接收者 和 值接受者
// func_pointer_reciver.go
package main

import "fmt"

type User struct {
	name string
	age  int
}

// 值接受者的函数
func (u *User) selfInfo() {
	fmt.Printf("u.name: %v\n", u.name)
	fmt.Printf("u.age: %v\n", u.age)
}

func main() {
	user := &User{"小淑", 23}
	user.selfInfo() // user是指针，使用指针
	user1 := User{"王保", 45}
	user1.selfInfo()
}
