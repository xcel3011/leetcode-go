package leetcode_go

import (
	"maps"
)

func findAnagrams(s string, p string) []int {
	need, window := make(map[byte]int), make(map[byte]int)
	for i := range p {
		need[p[i]]++
	}

	left, right := 0, 0
	valid := 0
	var res []int

	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		if right-left >= len(p) {
			if valid == len(need) {
				res = append(res, left)
			}

			d := s[left]
			left++
			if _, ok := window[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	return res
}

func findAnagrams240531(s string, p string) []int {
	m, n := len(s), len(p)
	if n > m {
		return nil
	}

	// 初始化窗口
	counts, countp := [26]int{}, [26]int{}
	for i := 0; i < n; i++ {
		counts[s[i]-'a']++
		countp[p[i]-'a']++
	}

	// 判断初始窗口是否合法
	var ans []int
	if countp == counts {
		ans = append(ans, 0)
	}

	// 向右滑动窗口
	for i := n; i < m; i++ {
		counts[s[i]-'a']++
		counts[s[i-n]-'a']--
		if countp == counts {
			ans = append(ans, i-n+1)
		}
	}
	return ans
}

// 给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
// 示例 1:
//
// 输入: s = "cbaebabacd", p = "abc"
// 输出: [0,6]
// 解释:
// 起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
// 起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
// 示例 2:
//
// 输入: s = "abab", p = "ab"
// 输出: [0,1,2]
// 解释:
// 起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
// 起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
// 起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。
func findAnagrams250324(s string, p string) []int {
	n, m := len(s), len(p)
	if m > n {
		return nil
	}

	// 初始化窗口
	var ans []int
	cnt1, cnt2 := map[byte]int{}, map[byte]int{}
	for i := range p {
		cnt1[p[i]]++
		cnt2[s[i]]++
	}

	// 判断初始窗口是否合法
	if maps.Equal(cnt1, cnt2) {
		ans = append(ans, 0)
	}

	// 窗口右移
	for i := m; i < n; i++ {
		// 加入新元素
		cnt2[s[i]]++
		// 移除旧元素
		cnt2[s[i-m]]--
		if cnt2[s[i-m]] == 0 {
			delete(cnt2, s[i-m])
		}
		// 判断窗口是否合法
		if maps.Equal(cnt1, cnt2) {
			ans = append(ans, i-m+1)
		}
	}
	return ans
}
