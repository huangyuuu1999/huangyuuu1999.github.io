### b = a 会发生什么?

在`python`中，若`a`是一个`list`，`b=a`会让`a`获得原来列表的修改权
```python
a = [2, 3, 5, 7]
b = a
b[0] = '>'
print(a, b)
```

在golang中这个小demo的表现是否一样？

```go
func main() {
    a := []int{2, 3, 5, 7}
    b := a
    b[0] = 999
    fmt.Println(a, b) // [999 3 5 7] [999 3 5 7]
}
```

目前看来表现的现象完全一样，但真的有那么简单吗？

### 不同的append

```python
a = [2, 3, 5, 7]
b = a
b[0] = '>'
print(a, b)  # ['>', 3, 5, 7] ['>', 3, 5, 7]

a.append(666)

print(a, b)  # ['>', 3, 5, 7, 666] ['>', 3, 5, 7, 666]
```
这次的输出不再一样了！

```go
func main() {
    a := []int{2, 3, 5, 7}
    b := a
    b[0] = 999
    fmt.Println(a, b) // [999 3 5 7] [999 3 5 7]
	
	a = append(a, 666)

    fmt.Println(a, b) // [999 3 5 7 666] [999 3 5 7]
}
```
虽然都叫`append`但是两个函数所做的事情不一样。
`golang`的`append`函数，在发现底层数组的`cap`不够填充新的元素的时候，重新在别的内存区域上开更大的底层数组，并将这个新的底层数组的某一处（可能是首地址）的指针返回。
所以当 `a = append(a, 666)`执行之后，`a`和原来的`b`已经不指向同一内存了，二者的关系只是前4个元素相同而已。此时修改一个也不会影响另一个了。

```go
func main() {
    a := []int{2, 3, 5, 7}
    b := a
    b[0] = 999
    fmt.Println(a, b) // [999 3 5 7] [999 3 5 7]
	
	a = append(a, 666)

    fmt.Println(a, b) // [999 3 5 7 666] [999 3 5 7]
	
	a[1] = 7777 // 修改 a 对 b 没有影响
    fmt.Println(a, b) // [999 7777 5 7 666] [999 3 5 7]
}
```
### 扩容？扩多少？
### 解释下面的代码
```go
func test1() {
	s1 := []int{1, 2}
	s2 := s1
	s2 = append(s2, 3)
	fmt.Printf("s2: %v\n", s2) // s2: [1 2 3]
	fmt.Printf("s1: %v\n", s1) // s1: [1 2]
	SliceRise(s1)
	SliceRise(s2)
	fmt.Printf("s2: %v\n", s2) // s2: [2 3 4]
	fmt.Printf("s1: %v\n", s1) // s1: [1 2]
}

func SliceRise(s []int) {
	s = append(s, 0)
	for i := range s {
		s[i] += 1
	}
}
```

### append函数导致的扩容行为
```go
package main

import "fmt"

func main() {
	s1 := []int{2, 3, 5}
	test1(s1)
	fmt.Printf("s1: %v\n", s1) // s1: [2 3 8]

	test2(s1)
	fmt.Printf("s1: %v\n", s1) // s1: [2 3 8], test2 内部改的，不是s1的数组

	fmt.Printf("s1 len: %v cap %v\n", len(s1), cap(s1)) // s1 len: 3 cap 3

	s1 = append(s1, 408) 
	fmt.Printf("s1: %v\n", s1) // s1: [2 3 8 408]
	fmt.Printf("s1 len: %v cap %v\n", len(s1), cap(s1))  // s1 len: 4 cap 6

	test2(s1)
	fmt.Printf("s1: %v\n", s1) // s1: [-3 3 8 408], 这次 test2 内部改的，是s1的数组
	fmt.Printf("s1: %p\n", s1)
}

func test1(s []int) {  // 切片作为参数，是传引用
	s[2] = 8
	fmt.Printf("s: %v\n", s) // s: [2 3 8]
}

func test2(s []int) {
	s = append(s, 42)  // append之后，新的s可能和原先的s1 公用数组，也可能不共用数组
	s[0] = -3
	fmt.Printf("s: %v\n", s) // s: [-3 3 8 42]
	fmt.Printf("s1: %p\n", s)
}

```