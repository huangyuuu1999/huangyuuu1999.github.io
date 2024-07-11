使用[x:y]进行切片时，何时会panic？
答案：当x<0 或者y>len(a)时
### 示例1
```go
package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	b := a[5:] // panic: runtime error: slice bounds out of range [5:4]
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)

    inorder := []int{3}
	right_part_inorder := inorder[1:]
	fmt.Printf("inorder: %v\n", inorder)
	fmt.Printf("right_part_inorder: %v\n", right_part_inorder)
	// inorder: [3]
	// right_part_inorder: []
}
```
### 示例2
下面的代码，在只有一个元素 inorder = [3], preorder = [3]时，会panic吗？
```go
import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	return construct(preorder, inorder)
}

func findIndex(inorder []int, target int) int {
    for i := range inorder {
        if inorder[i] != target {
            continue
        }
        return i
    }
    return -1
}

func construct(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}

	mid := findIndex(inorder, preorder[0])

	left_inorder := inorder[:mid]
	left_n := len(left_inorder)

	right_inorder := inorder[mid+1:]
	right_n := len(right_inorder)

	left_preorder := preorder[1 : 1+left_n]
	right_preorder := preorder[1+left_n : 1+left_n+right_n]

	root.Left = construct(left_preorder, left_inorder)
	root.Right = construct(right_preorder, right_inorder)
	return root
}
```