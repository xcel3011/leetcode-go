package leetcode_go

// 窗口是所有1的个数
// 窗口内有多少0就要交换几次
func minSwaps(nums []int) int {
	k := 0 // 窗口大小
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			k++
		}
	}

	// 第一个窗口从下标0开始
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	ans := k - sum

	// 窗口向右滑动
	for i := 1; i < len(nums); i++ {
		sum += nums[(i+k-1)%len(nums)] - nums[i-1]
		ans = min(ans, k-sum)
	}
	return ans
}
