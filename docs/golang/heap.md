在go中使用堆 需要使用到 container 下面的 heap

并且需要知道一些 “接口” 的特性。

go没有类，但是可以通过组合和转发特性来实现面向对象。

结构体嵌入，匿名字段

转发方法

## 结构体嵌入

在定义结构体的时候，可以在成员中写其他结构体，这个就叫结构体的嵌入。

### 具名字段
结构体的字段可以是另外的结构体类型，这在其他的语言中也很常见。
```go
// struct_embedding.go 演示结构体的字段是另外的结构体类型
package main

import "fmt"

type BodyInfo struct {
	height float64 // 身高cm
	weight float64 // 体重kg
}

func (bi BodyInfo) calBMI() float64 {
	heightInMeters := bi.height / 100 // 将身高从厘米转换为米
	return bi.weight / (heightInMeters * heightInMeters)
}

type User struct {
	bodyInfo BodyInfo // 这里就是结构体字段嵌入
	name     string
	age      int
}

func main() {
	wang := User{
		BodyInfo{175, 70}, "王小明", 22,
	}
	wangBMI := wang.bodyInfo.calBMI()
	fmt.Printf("wangBMI: %.2f\n", wangBMI) // wangBMI: 22.86
}

```

### 匿名字段
定义结构体的时候，字段可以不起名字，这时候，这个字段的类型实现的方法，就转发到了外面的结构体身上。
```go
package main

import "fmt"

type BodyInfo struct {
	height float64 // 身高cm
	weight float64 // 体重kg
}

func (bi BodyInfo) calBMI() float64 {
	heightInMeters := bi.height / 100 // 将身高从厘米转换为米
	return bi.weight / (heightInMeters * heightInMeters)
}

type User struct {
	BodyInfo // 字段是另外的结构体，但是不起名字
	name     string
	age      int
}

func main() {
	wang := User{
		BodyInfo{175, 70}, "王小明", 22,
	}
	wangBMI := wang.calBMI()               // 发现匿名的字段，的方法，直接可以被User使用，也就是说相当于User实现了calBMI方法
	fmt.Printf("wangBMI: %.2f\n", wangBMI) // wangBMI: 22.86
}

```

## 值接受者和指针接收者
值接受者的函数，指针接收者可以直接使用，内部会转换。
指针接收者的函数，

main函数中的user是指针，但是可以调用selfInfo。
```go
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

```
下面的代码展示了：指针接收者的函数，可以被值直接调用。
```go
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

```