package main

import (
	"fmt"
	"strings"
)

func permute(nums []string) []string {
    ans := []string{}
    n := len(nums)
    var dfs func(i int)
    dfs = func(i int) {
        if i == n {
            tmp := strings.Join(nums, "~")
            ans = append(ans, tmp)
            return
        }
        for j := i; j<n; j++ {
            nums[i], nums[j] = nums[j], nums[i]
            dfs(i+1)
            nums[i], nums[j] = nums[j], nums[i]
        }
    }
    dfs(0)
    return ans
}

func main() {
    tokens := []string{"宝宝", "肚肚", "打", "雷雷"}
    res := permute(tokens)
    for _, item := range res {
        fmt.Println("♪", item)
    }
}