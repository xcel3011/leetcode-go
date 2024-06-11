package leetcode_go

import (
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	l, r := 0, len(people)-1
	ans := 0
	for l <= r {
		// 每艘船最多可同时载两人
		// 但条件是这些人的重量之和最多为 limit。
		sum := people[l] + people[r]
		r -= 1
		if sum <= limit {
			l += 1
		}
		ans++
	}
	return ans
}
