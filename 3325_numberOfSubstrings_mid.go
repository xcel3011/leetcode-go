package leetcode_go

// 给你一个字符串 s 和一个整数 k，在 s 的所有子字符串中，
// 请你统计并返回 至少有一个 字符 至少出现 k 次的子字符串总数。
// 子字符串 是字符串中的一个连续、 非空 的字符序列。
//
// 示例 1：
// 输入： s = "abacb", k = 2
// 输出： 4
// 解释：
// 符合条件的子字符串如下：
// "aba"（字符 'a' 出现 2 次）。
// "abac"（字符 'a' 出现 2 次）。
// "abacb"（字符 'a' 出现 2 次）。
// "bacb"（字符 'b' 出现 2 次）。
//
// 示例 2：
// 输入： s = "abcde", k = 1
// 输出： 15
// 解释：
// 所有子字符串都有效，因为每个字符至少出现一次。
//
// 提示：
// 1 <= s.length <= 3000
// 1 <= k <= s.length
// s 仅由小写英文字母组成。
func numberOfSubstrings250325(s string, k int) int {
	cnt := map[byte]int{}
	ans := 0
	l, r := 0, 0
	for r < len(s) {
		in := s[r]
		r++
		cnt[in]++
		for cnt[in] >= k {
			out := s[l]
			l++
			cnt[out]--
		}
		ans += l
	}
	return ans
}
