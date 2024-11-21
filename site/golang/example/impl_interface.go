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
