package main

/*
快速选择算法，比较各种partition方式的差异
https://leetcode.cn/problems/kth-largest-element-in-an-array/description/

*/

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"slices"
	"time"
)

func generateRandomArr(n int, minVal, maxVal int) []int {
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())

	array := make([]int, n)
	for i := 0; i < n; i++ {
		// 生成随机整数
		array[i] = rand.Intn(maxVal-minVal+1) + minVal
	}
	return array
}

// 检查m下标位置 的左边是不是都小于它，右边是不是都大于他
func testPartitionResult(arr []int, p, r, m int) bool {
	for i := p; i <= r; i++ {
		if i < m && arr[i] > arr[m] {
			return false
		}
		if i > m && arr[i] < arr[m] {
			return false
		}
	}
	return true
}

func testQuickSort(partitionFunc func([]int, int, int) int) {
	fv := reflect.ValueOf(partitionFunc)
	pc := fv.Pointer()
	fn := runtime.FuncForPC(pc)
	fName := fn.Name()
	fmt.Printf("****** testing quicksort func %v *****\n", fName)
	for i := range 10 {
		fmt.Printf("test[%2d|10]", i+1)
		arr := generateRandomArr(35, -2, 100)
		quickSort(arr, 4, 32, partitionFunc)
		res := slices.IsSorted(arr[4:33])
		var flag string
		if res {
			flag = "🟢"
		} else {
			flag = "🔴"
		}
		fmt.Printf(": %v\n", flag)
	}
}

// lomuto scheme，选左边第一个作为pivot; j是探测指针，i维护小于x区域的右边界
func partitionLomuto(arr []int, p, r int) int {
	x := arr[p]
	i := p
	for j := p + 1; j <= r; j++ {
		if arr[j] < x {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// 最后把pivot换到中间，也就是小于x区域的最右侧元素的位置
	arr[p], arr[i] = arr[i], arr[p]
	return i
}

// slot 方法，左边挖空右边找一个填空，左右交替
func partitionSlots(arr []int, p, r int) int {
	x := arr[p]
	l, r := p, r
	for l < r {
		for l < r && arr[r] >= x {
			r--
		}
		arr[l] = arr[r]
		for l < r && arr[l] < x {
			l++
		}
		arr[r] = arr[l]
	}
	arr[l] = x // l == r
	return l
}

// 左右找，一次交互, 比挖空slot交换的次数少
func partitionHoare(arr []int, p, r int) int {
	pivotIndex := p
	x := arr[pivotIndex]
	l, r := p, r
	for l < r {
		for l < r && arr[r] >= x {
			r--
		}
		for l < r && arr[l] <= x {
			l++
		}
		arr[l], arr[r] = arr[r], arr[l]
	}
	arr[pivotIndex], arr[r] = arr[r], arr[pivotIndex]
	return r
}

// 3.3 写快速排序 新写了一个排序算法
func partition3_3(nums []int, p, r int) int {
	pivot := nums[p]
	l, r := p, r
	for l < r {
		for l < r && nums[r] >= pivot { // 这里必须要写等于
			r--
		}
		for l < r && nums[l] <= pivot { // 这里也是必须要写等于
			l++
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	nums[p], nums[r] = nums[r], nums[p]
	return r
}

// 更类似于hoare原始写法的，pivot的值的位置可以变。do while写法
// https://blog.csdn.net/qq_33919450/article/details/127095084
// https://blog.csdn.net/2201_75314884/article/details/137834523
// https://en.wikipedia.org/wiki/Quicksort#Hoare_partition_scheme
// 关键在于，最终pivot的位置，不是在中间该在的位置的，所以不能用上面的测试案例去测，但本身是对的
// https://www.bilibili.com/video/BV15A411R7P4/?spm_id_from=333.337.search-card.all.click&vd_source=d3681ed3219ed1f8a753e351f7feb904
func partitionHoare1(arr []int, p, r int) int {
	x := arr[p]
	lo, hi := p-1, r+1
	for lo < hi {
		for lo++; arr[lo] < x; lo++ {
		}
		for hi--; arr[hi] > x; hi-- {
		}
		if lo < hi {
			arr[lo], arr[hi] = arr[hi], arr[lo]
		} else {
			break
		}
	}
	return hi
}

func testPartitionFunctions(f func([]int, int, int) int) {
	fv := reflect.ValueOf(f)
	pc := fv.Pointer()
	fn := runtime.FuncForPC(pc)
	fName := fn.Name()
	fmt.Printf("****** testing partition func %v *****\n", fName)
	for i := range 10 {
		fmt.Printf("test[%2d|10]", i+1)
		arr := generateRandomArr(5, 0, 100)
		m := f(arr, 0, 4)
		res := testPartitionResult(arr, 0, 4, m)
		var flag string
		if res {
			flag = "🟢"
		} else {
			flag = "🔴"
		}
		fmt.Printf(": %v\n", flag)
	}
}

func quickSort(arr []int, p, r int, partitionFunc func([]int, int, int) int) {
	if r-p+1 <= 1 {
		return
	}
	m := partitionFunc(arr, p, r)
	quickSort(arr, p, m-1, partitionFunc)
	quickSort(arr, m+1, r, partitionFunc)
}

func main() {
	testPartitionFunctions(partitionLomuto)
	testPartitionFunctions(partitionSlots)
	testPartitionFunctions(partitionHoare)
	testPartitionFunctions(partition3_3)
	// testPartitionFunctions(partitionHoare1) // not suit this test mode, but it's right?
	testQuickSort(partitionLomuto)
	testQuickSort(partitionSlots)
	testQuickSort(partitionHoare)
}

// 荷兰国旗问题
// https://leetcode.cn/problems/sort-colors/description/
func Dutchflag(nums []int, p, r int, x int) (int, int) {
	less, more := p-1, r+1 // 小于区域 和 大于区域的 边界
	for i := p; i < more; i++ {
		if nums[i] == x {
			continue
		} else if nums[i] < x {
			less++
			nums[less], nums[i] = nums[i], nums[less]
		} else {
			more--
			nums[more], nums[i] = nums[i], nums[more]
			i--
		}
	}
	return less, more
}

/*
在快速排序的分区操作中，元素的视角和缝隙的视角分别对应不同的实现策略，主要区别在于分界点的确定方式和元素的处理逻辑：

---

### **1. 元素的视角（如 Lomuto 分区法）**
- **核心思想**：通过逐个交换元素，将基准值放置到正确的位置，分界点即基准值的最终索引。
- **实现步骤**：
  1. 选择最右侧元素作为基准（pivot）。
  2. 维护指针 `i`，表示小于基准的子数组的末尾。
  3. 遍历数组，将小于基准的元素交换到 `i` 的左侧。
  4. 最后将基准交换到 `i+1` 的位置，此时分界点为 `i+1`。
- **特点**：
  - 基准元素最终位于正确位置。
  - 分界点是明确的元素索引。
  - 实现简单，但交换次数较多。

```python
def lomuto_partition(arr, low, high):
    pivot = arr[high]
    i = low - 1
    for j in range(low, high):
        if arr[j] <= pivot:
            i += 1
            arr[i], arr[j] = arr[j], arr[i]
    arr[i+1], arr[high] = arr[high], arr[i+1]
    return i + 1  # 分界点是基准的最终位置
```

---

### **2. 缝隙的视角（如 Hoare 分区法）**
- **核心思想**：通过双指针从两端向中间扫描，找到分界点的“缝隙”，分割数组而不固定基准位置。
- **实现步骤**：
  1. 选择最左侧元素作为基准。
  2. 左指针 `i` 向右找大于基准的元素，右指针 `j` 向左找小于基准的元素。
  3. 交换不符合条件的元素，直到指针相遇，返回相遇点 `j` 作为分界。
- **特点**：
  - 分界点是一个虚拟的“缝隙”（如 `j` 的位置），左右子数组为 `[low, j]` 和 `[j+1, high]`。
  - 基准不一定在最终位置，但分界点将数组划分为两部分。
  - 交换次数较少，效率更高，但逻辑较复杂。

```python
def hoare_partition(arr, low, high):
    pivot = arr[low]
    i = low - 1
    j = high + 1
    while True:
        i += 1
        while arr[i] < pivot:
            i += 1
        j -= 1
        while arr[j] > pivot:
            j -= 1
        if i >= j:
            return j  # 分界点是缝隙 j，分割为 [low, j] 和 [j+1, high]
        arr[i], arr[j] = arr[j], arr[i]
```

---

### **对比总结**
| **特性**         | **元素的视角（Lomuto）**               | **缝隙的视角（Hoare）**               |
|------------------|--------------------------------------|--------------------------------------|
| **分界点**       | 基准元素的最终索引（具体元素位置）       | 指针相遇的虚拟缝隙（不固定元素位置）    |
| **基准位置**     | 最终位于分界点                         | 可能不在分界点                        |
| **交换次数**     | 较多（每次小元素都交换）                | 较少（仅交换不匹配的较大块）           |
| **实现复杂度**   | 简单直观                               | 较复杂，需处理指针越界                 |
| **适用场景**     | 教学或简单实现                         | 高性能需求场景                        |

---

通过这两种视角的分区操作，快速排序可以灵活平衡代码可读性与性能，理解其差异有助于在不同场景下选择合适的实现方式。

*/
