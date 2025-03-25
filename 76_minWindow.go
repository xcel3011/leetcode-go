package leetcode_go

import "math"

func minWindow(s string, t string) string {
	need := make(map[byte]int)
	window := make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}

	left, right := 0, 0
	valid := 0
	start, l := 0, math.MaxInt

	for right < len(s) {
		c := s[right]
		right++

		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left < l {
				start = left
				l = right - left
			}

			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if l == math.MaxInt {
		return ""
	}

	return s[start : start+l]
}

// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。
// 如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
// 注意：
// 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
// 如果 s 中存在这样的子串，我们保证它是唯一的答案。
//
// 示例 1：
// 输入：s = "ADOBECODEBANC", t = "ABC"
// 输出："BANC"
// 解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
//
// 示例 2：
// 输入：s = "a", t = "a"
// 输出："a"
// 解释：整个字符串 s 是最小覆盖子串。
//
// 示例 3:
// 输入: s = "a", t = "aa"
// 输出: ""
// 解释: t 中两个字符 'a' 均应包含在 s 的子串中，
// 因此没有符合条件的子字符串，返回空字符串。
//
// 提示：
// m == s.length
// n == t.length
// 1 <= m, n <= 105
// s 和 t 由英文字母组成
func minWindow250325(s string, t string) string {
	need := make(map[byte]int)
	window := make(map[byte]int)
	for i := range t {
		need[t[i]]++
	}
	l, r := 0, 0
	valid := 0
	start, end := 0, math.MaxInt
	for r < len(s) {
		in := s[r]
		r++
		_, ok := need[in]
		if ok {
			window[in]++
			if window[in] == need[in] {
				valid++
			}
		}
		for valid == len(need) {
			if r-l < end-start {
				start = l
				end = r
			}

			out := s[l]
			l++
			_, ok = need[out]
			if ok {
				if window[out] == need[out] {
					valid--
				}
				window[out]--
			}
		}
	}
	if end == math.MaxInt {
		return ""
	}
	return s[start:end]
}
