package leetcode_go

func getAverages(nums []int, k int) []int {
	ans := make([]int, len(nums))
	for i := 0; i < len(ans); i++ {
		ans[i] = -1
	}
	if len(nums) <= 2*k {
		return ans
	}

	sum := 0
	// 初始化窗口
	for i := 0; i <= 2*k; i++ {
		sum += nums[i]
	}
	ans[k] = sum / (2*k + 1)

	// 从第二位开始算
	for i := k + 1; i < len(nums)-k; i++ {
		sum += nums[i+k] - nums[i-k-1]
		ans[i] = sum / (2*k + 1)
	}
	return ans
}
