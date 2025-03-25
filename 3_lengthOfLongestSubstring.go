package leetcode_go

func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)

	left, right := 0, 0
	res := 0

	for right < len(s) {
		c := s[right]
		right++
		window[c]++

		for window[c] > 1 {
			d := s[left]
			left++
			window[d]--
		}

		diff := right - left
		if res < diff {
			res = diff
		}
	}

	return res
}

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。
//
// 示例 1:
//
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 示例 2:
//
// 输入: s = "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 示例 3:
//
// 输入: s = "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//
//	请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
func lengthOfLongestSubstring250325(s string) int {
	l, r := 0, 0
	window := map[byte]int{}
	ans := 0

	for r < len(s) {
		// 获取右侧窗口字符
		in := s[r]
		// 窗口右移
		r++
		// 右侧进入的字符+1
		window[in]++

		// 判断时候有重复字符
		for window[in] > 1 {
			// 左侧窗口字符
			out := s[l]
			// 窗口左移
			l++
			// 左侧字符-1
			window[out]--
		}

		ans = max(ans, r-l)
	}
	return ans
}
