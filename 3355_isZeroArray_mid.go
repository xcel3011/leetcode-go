package leetcode_go

func isZeroArray(nums []int, queries [][]int) bool {
	// diff[i] = nums[i] - nums[i-1] （i ≥ 1）
	// nums[i] = diff[0] + diff[1] + ... + diff[i] （前缀和）
	diff := make([]int, len(nums)+1)
	for _, q := range queries {
		// 区间 [l,r] 中的数都加一
		diff[q[0]]++
		diff[q[1]+1]--
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += diff[i]
		if nums[i] > sum {
			return false
		}
	}
	return true
}
