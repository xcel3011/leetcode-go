package leetcode_go

func maxPower(s string) int {
	n := len(s)
	ans, i := 0, 0
	for i < n {
		start := i
		i++
		for i < n && s[i] == s[start] {
			i++
		}
		ans = max(ans, i-start)
	}
	return ans
}
