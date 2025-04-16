package leetcode_go

func longestSemiRepetitiveSubstring(s string) int {
	ans := 1
	left := 0
	repeat := 0
	for right := 1; right < len(s); right++ {
		// 右移指针，统计重复的情况出现了几次
		if s[right] == s[right-1] {
			repeat++
		}

		// 如果重复次数大于1，左指针右移，直到重复次数为1
		if repeat > 1 {
			left++
			for s[left] != s[left-1] {
				left++
			}
			repeat = 1
		}
		// 计算当前子串的长度
		ans = max(ans, right-left+1)
	}
	return ans
}
