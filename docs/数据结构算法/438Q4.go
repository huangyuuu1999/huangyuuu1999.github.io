package main

import (
	"sort"
)

func maxDistance(side int, points [][]int, k int) int {
	n := len(points)
	sort.Slice(points, func(i, j int) bool {
		d1, d2 := points[i][0]+points[i][1], points[j][0]+points[j][1]
		return d1 < d2
	})

	check := func(d int) bool {
		// 选k个，如果k是偶数，左边k个 右边k个
		if k%2 == 0 {
			sp := k / 2
			lp, rp := points[sp-1], points[n-sp]
			dx, dy := rp[0]-lp[0], rp[1]-rp[1]
			if dx < 0 {
				dx = -dx
			}
			if dy < 0 {
				dy = -dy
			}
			return dx+dy >= d
		} else {
			sp := k / 2
			lp, rp := points[sp], points[n-sp]
			dx, dy := rp[0]-lp[0], rp[1]-rp[1]
			if dx < 0 {
				dx = -dx
			}
			if dy < 0 {
				dy = -dy
			}
			ans1 := dx+dy >= d
			lp, rp = points[sp-1], points[n-sp-1]
			dx, dy = rp[0]-lp[0], rp[1]-rp[1]
			if dx < 0 {
				dx = -dx
			}
			if dy < 0 {
				dy = -dy
			}
			ans2 := dx+dy >= d
			return ans1 || ans2
		}
	}
	ans := sort.Search(side+1, func(i int) bool {
		if i == 0 {
			return false
		}
		return !check(i)
	})
	return ans - 1
}

/*
# 二分 可能达到的最大距离
# check怎么写？k距离能否达到？
points.sort()

def check(k) -> bool:  # 检查选k个点，最小距离能否达到k
*/
