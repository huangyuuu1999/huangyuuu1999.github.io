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
