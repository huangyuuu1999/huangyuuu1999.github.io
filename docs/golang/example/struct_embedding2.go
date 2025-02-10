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
