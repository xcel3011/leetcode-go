package leetcode_go

// 右侧新进入窗口的字母为元音字母，左侧移出窗口的字母也是元音字母，这样一进一出抵消掉了
// 右侧新进入窗口的字母为元音字母，左侧移出窗口的字母非元音字母，此时元音字母个数+1
// 右侧新进入窗口的字母非元音字母，左侧移出窗口的字母为元音字母，此时元音字母个数-1
func maxVowels(s string, k int) int {
	m := map[byte]int{'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1}
	count := 0
	for i := 0; i < k; i++ {
		count += m[s[i]]
	}
	ans := count
	for i := k; i < len(s); i++ {
		count += m[s[i]] - m[s[i-k]]
		ans = max(ans, count)
	}
	return ans
}
