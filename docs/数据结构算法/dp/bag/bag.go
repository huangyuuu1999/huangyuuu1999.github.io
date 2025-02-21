package main

func main() {

}
// 0-1背包 回溯法
func zero_one_knapsack(cap int, w, v []int) int {
	var dfs func(i int, leftCap int) int // 考虑前i个物品，容量为laftCap时的最大价值
	dfs = func(i, leftCap int) int {
		if i < 0 {
			return 0
		}
		skip := dfs(i-1, leftCap)
		take := dfs(i-1, leftCap - w[i]) + v[i]
		if leftCap < w[i] { // 背包容量小于物品体积，无法拿i号物品
			return skip
		}
		return max(skip, take)
	}
	return dfs(len(w), cap)
}

// 0-1背包 回溯法
func zero_one_knapsack_backtrace(cap int, w, v []int) int {
	
}