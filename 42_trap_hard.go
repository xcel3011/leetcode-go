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
