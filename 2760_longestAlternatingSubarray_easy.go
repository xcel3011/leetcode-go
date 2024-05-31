package leetcode_go

func longestAlternatingSubarray(nums []int, threshold int) int {
	n := len(nums)
	ans := 0
	i := 0
	for i < n {
		if nums[i]%2 != 0 || nums[i] > threshold {
			i++
			continue
		}
		start := i
		i++
		for i < n && nums[i] <= threshold && nums[i]%2 != nums[i-1]%2 {
			i++
		}
		ans = max(ans, i-start)

	}
	return ans
}
