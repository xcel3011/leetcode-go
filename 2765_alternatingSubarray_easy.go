package leetcode_go

func alternatingSubarray(nums []int) int {
	ans := -1
	for i := 0; i < len(nums)-1; {
		if nums[i+1]-nums[i] != 1 {
			i++
			continue
		}
		start := i
		i += 2
		for i < len(nums) && nums[i] == nums[i-2] {
			i++
		}
		ans = max(ans, i-start)
		i--
	}
	return ans
}
