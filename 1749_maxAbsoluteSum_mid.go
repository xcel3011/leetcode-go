package leetcode_go

import (
	"math"
)

func maxAbsoluteSum(nums []int) int {
	sums := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		sums[i+1] = sums[i] + nums[i]
	}

	var (
		mx, mn int
		res    int
	)
	for i := 0; i < len(sums); i++ {
		mx = max(mx, sums[i])
		mn = min(mn, sums[i])
		// 找到前缀和数组中的最大值减去最小值，得到一个最大正数
		//（前提是最大值出现在最小值的后面，并且最小值是个负数，否则应该直接取最大值作为答案）
		res = max(res, int(math.Abs(float64(sums[i]-mn))))
		//找到前缀和的最小值减去最大值，得到一个最小负数
		//（前提是最小值出现在最大值的后面，而且最大值是一个正数，否则直接取最小值作为答案
		res = max(res, int(math.Abs(float64(sums[i]-mx))))
	}
	return res
}
