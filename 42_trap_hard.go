package leetcode_go

func trap(height []int) int {
	ans := 0
	l, r := 0, len(height)-1
	lm, rm := 0, 0
	for l < r {
		lm = max(lm, height[l])
		rm = max(rm, height[r])
		if lm < rm {
			ans += lm - height[l]
			l++
		} else {
			ans += rm - height[r]
			r--
		}
	}
	return ans
}

// 时间复杂的度：O(n)，其中 n 是数组 height 的长度。
// 空间复杂度：O(n)，其中 n 是数组 height 的长度。空间复杂度主要取决于两个数组 preMax 和 sufMax 的长度。
func trap250415(height []int) int {
	n := len(height)
	preMax := make([]int, n)
	preMax[0] = height[0]
	for i := 1; i < n; i++ {
		preMax[i] = max(preMax[i-1], height[i])
	}
	sufMax := make([]int, n)
	sufMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		sufMax[i] = max(sufMax[i+1], height[i])
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans += min(preMax[i], sufMax[i]) - height[i]
	}
	return ans
}
