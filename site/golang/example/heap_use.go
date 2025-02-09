package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type maxHeap struct {
	sort.IntSlice
}

// Less 方法反转比较逻辑，实现最大堆
func (mh maxHeap) Less(i, j int) bool {
	return mh.IntSlice[i] > mh.IntSlice[j]
}

func (mh *maxHeap) Pop() any {
	old := mh.IntSlice
	n := len(old)
	x := old[n-1]
	mh.IntSlice = old[0 : n-1]
	return x
}

func (mh *maxHeap) Push(x any) {
	mh.IntSlice = append(mh.IntSlice, x.(int))
}

func main() {
	testPointerReciver() // 指针可以直接使用 值接受者的方法

	a := &maxHeap{sort.IntSlice{3, 2, 5, 1, 8, 9}}
	heap.Init(a)
	for _ = range 3 {
		popRes := heap.Pop(a)
		fmt.Printf("popRes: %v\n", popRes)
	}
}

type User struct {
	name string
}

func (u User) basicInfo() {
	fmt.Printf("u.name: %v\n", u.name)
}

func testPointerReciver() {
	u := &User{"杨戬"}
	u.basicInfo()
}
