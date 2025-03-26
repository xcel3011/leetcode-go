package leetcode_go

// 给你一个字符串 s ，它只包含三种字符 a, b 和 c 。
// 请你返回 a，b 和 c 都 至少 出现过一次的子字符串数目。
//
// 示例 1：
// 输入：s = "abcabc"
// 输出：10
// 解释：包含 a，b 和 c 各至少一次的子字符串为 "abc", "abca", "abcab", "abcabc", "bca", "bcab", "bcabc", "cab", "cabc" 和 "abc" (相同字符串算多次)。
//
// 示例 2：
// 输入：s = "aaacb"
// 输出：3
// 解释：包含 a，b 和 c 各至少一次的子字符串为 "aaacb", "aacb" 和 "acb" 。
//
// 示例 3：
// 输入：s = "abc"
// 输出：1
//
// 提示：
// 3 <= s.length <= 5 x 10^4
// s 只包含字符 a，b 和 c 。
func numberOfSubstrings(s string) int {
	cnt := map[byte]int{}
	ans := 0
	l, r := 0, 0
	for r < len(s) {
		in := s[r]
		r++
		cnt[in]++
		for len(cnt) == 3 {
			out := s[l]
			l++
			cnt[out]--
			if cnt[out] == 0 {
				delete(cnt, out)
			}
		}

		// 右端点固定在 right，左端点在 0,1,2,…,left−1 的所有子串都是合法的，这一共有 left 个，
		ans += l
	}
	return ans
}
