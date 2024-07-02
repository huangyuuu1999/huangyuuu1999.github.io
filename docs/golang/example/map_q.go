package main

import "fmt"


func main() {
	m := map[int]int{1: 2, 3: 9}
	fmt.Printf("m[5]: %v\n", m[5]) // m[5]: 0

	// 对于不存在于m的键，你甚至可以对它进行++操作
	m[4]++
	fmt.Printf("m: %v\n", m) // m: map[1:2 3:9 4:1]

	
	// 可以检查 某个元素是否存在于map中
	if val, exist := m[8]; exist {
		fmt.Println("8存在", val)
		} else {
			fmt.Println("8不存在") // 执行这一行
		}
		
	var sz = len(m)
	fmt.Printf("sz: %v\n", sz)
	// 此时可以 m[8]++, 然后8就存在了
	m[8]++
	if val, exist := m[8]; exist {
		fmt.Println("8存在", val) // 执行这一行
	} else {
		fmt.Println("8不存在") 
	}

	sz = len(m)
	fmt.Printf("sz: %v\n", sz)
}