package leetcode_go

func getLongestSubsequence(words []string, groups []int) []string {
	ans := make([]string, 0, len(groups))
	for i := 0; i < len(groups); i++ {
		if i == 0 || groups[i] != groups[i-1] {
			ans = append(ans, words[i])
		}
	}
	return ans
}
