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

// 几张卡牌 排成一行，每张卡牌都有一个对应的点数。点数由整数数组 cardPoints 给出。
//
// 每次行动，你可以从行的开头或者末尾拿一张卡牌，最终你必须正好拿 k 张卡牌。
//
// 你的点数就是你拿到手中的所有卡牌的点数之和。
//
// 给你一个整数数组 cardPoints 和整数 k，请你返回可以获得的最大点数。
func maxScore250321(cardPoints []int, k int) int {
	// 只能从两边取
	// 中间n-k个是不要的
	// 不要的n-k个滑动
	sum, cnt := 0, 0
	for i, point := range cardPoints {
		sum += point

		// 剩余左边的n-k个，相当于最开始取最右边的k个
		if i < len(cardPoints)-k {
			cnt += point
		}
	}

	ans := sum - cnt
	// n-k个滑动,从左边开始滑动，模拟从两边取
	for i, point := range cardPoints[len(cardPoints)-k:] {
		cnt += point - cardPoints[i]
		ans = max(ans, sum-cnt)
	}
	return ans
}
