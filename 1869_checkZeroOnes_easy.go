package leetcode_go

func checkZeroOnes(s string) bool {
	n := len(s)
	ans := [2]int{}
	i := 0
	for i < n {
		start := i
		i++
		for i < n && s[i] == s[i-1] {
			i++
		}

		ans[s[start]-'0'] = max(ans[s[start]-'0'], i-start)
	}
	return ans[1]-ans[0] > 0
}
