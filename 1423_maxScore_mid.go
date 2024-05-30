package leetcode_go

// 转换下思路，不直接两端选取的和
func maxScore(cardPoints []int, k int) int {
	// 如果最开始选最右边，那左边的n-k个是不要的
	window, sum := 0, 0
	for i := 0; i < len(cardPoints); i++ {
		sum += cardPoints[i]
		if i < len(cardPoints)-k {
			window += cardPoints[i]
		}
	}

	// 初始化返回
	ans := sum - window

	// 不选取的n-k个向右滑动
	for i := len(cardPoints) - k; i < len(cardPoints); i++ {
		window += cardPoints[i] - cardPoints[i-(len(cardPoints)-k)]
		ans = max(ans, sum-window)
	}
	return ans
}
