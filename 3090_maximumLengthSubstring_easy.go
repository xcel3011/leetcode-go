package leetcode_go

// 给你一个字符串 s ，请找出满足每个字符最多出现两次的最长子字符串，并返回该子字符串的 最大 长度。
//
// 示例 1：
// 输入： s = "bcbbbcba"
// 输出： 4
// 解释：
// 以下子字符串长度为 4，并且每个字符最多出现两次："bcba"。
//
// 示例 2：
// 输入： s = "aaaa"
// 输出： 2
// 解释：
// 以下子字符串长度为 2，并且每个字符最多出现两次："aa"。
// 提示：
// 2 <= s.length <= 100
// s 仅由小写英文字母组成。
func maximumLengthSubstring(s string) int {
	l, r := 0, 0
	w := map[byte]int{}
	ans := 0
	for r < len(s) {
		in := s[r]
		r++
		w[in]++

		for w[in] > 2 {
			out := s[l]
			l++
			w[out]--
		}

		ans = max(ans, r-l)
	}
	return ans
}
