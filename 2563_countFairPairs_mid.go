package leetcode_go

import (
	"slices"
	"sort"
)

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	var ans int64
	for i := 0; i < len(nums); i++ {
		r := sort.SearchInts(nums[:i], upper-nums[i]+1)
		l := sort.SearchInts(nums[:i], nums[i]-lower)
		ans += int64(r - l)
	}
	return ans
}

func countFairPairs250417(nums []int, lower int, upper int) (ans int64) {
	slices.Sort(nums)
	// lower<=nums[i]+nums[j]<=upper
	// lower-nums[i]<=nums[j]<=upper-nums[i]
	// 对于每个nums[i]，找到lower-nums[i]<=nums[j]<=upper-nums[i]的j的个数
	for i, x := range nums {
		// 先找到upper-x的位置
		right, _ := slices.BinarySearch(nums[:i], upper-x+1)
		// 再找到lower-x的位置
		left, _ := slices.BinarySearch(nums[:i], lower-x)
		ans += int64(right - left)
	}
	return
}
