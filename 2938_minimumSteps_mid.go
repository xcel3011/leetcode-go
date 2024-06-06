package leetcode_go

func minimumSteps(s string) int64 {
	var cnt, ans int64
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			cnt++
		} else {
			ans += cnt
		}
	}
	return ans
}
