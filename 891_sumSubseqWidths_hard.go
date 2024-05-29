package leetcode_go

import "sort"

// 超出时间限制
func sumSubseqWidths(nums []int) int {
	var subsets [][]int
	subsets = append(subsets, []int{})
	for _, num := range nums {
		subsetsLen := len(subsets)
		for i := 0; i < subsetsLen; i++ {
			subsets = append(subsets, append([]int{}, append(subsets[i], num)...))
		}
	}

	var ans int
	for _, subset := range subsets {
		if len(subset) < 2 {
			continue
		}
		sort.Ints(subset)
		ans += subset[len(subset)-1] - subset[0]
	}
	return ans
}

func sumSubseqWidthsMath(nums []int) int {
	const mod int = 1e9 + 7
	sort.Ints(nums)
	p, n := 1, len(nums)
	ans := 0
	for i, num := range nums {
		ans = (ans + (num-nums[n-i-1])*p + mod) & mod
		p = (p << 1) % mod
	}
	return ans
}
