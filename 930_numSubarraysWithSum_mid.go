package leetcode_go

func numSubarraysWithSum(nums []int, goal int) int {
	m := map[int]int{}
	var res, sum int
	for i := 0; i < len(nums); i++ {
		m[sum]++
		sum += nums[i]
		res += m[sum-goal]

	}
	return res
}
