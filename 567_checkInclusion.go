package leetcode_go

import (
	"fmt"
	"maps"
)

func checkInclusion(s1 string, s2 string) bool {
	need, window := make(map[byte]int), make(map[byte]int)
	for i := range s1 {
		need[s1[i]]++
	}

	left, right := 0, 0
	valid := 0

	for right < len(s2) {
		c := s2[right]
		right++

		if _, ok := need[c]; ok {
			window[c]++
			if need[c] == window[c] {
				valid++
			}
		}

		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}

			d := s2[left]
			left++

			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	return false
}

func checkInclusion240531(s1 string, s2 string) bool {
	// 肯定不存在
	m, n := len(s1), len(s2)
	if n < m {
		return false
	}

	// 初始化窗口
	cs1, cs2 := [26]int{}, [26]int{}
	for i := 0; i < m; i++ {
		cs1[s1[i]-'a']++
		cs2[s2[i]-'a']++
	}

	// 判断初始化的情况是否相等
	if cs1 == cs2 {
		return true
	}

	// 窗口向右移动
	for i := m; i < n; i++ {
		cs2[s2[i]-'a']++
		cs2[s2[i-m]-'a']--
		if cs1 == cs2 {
			return true
		}
	}
	return false
}

// 给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的 排列。如果是，返回 true ；否则，返回 false 。
//
// 换句话说，s1 的排列之一是 s2 的 子串 。
// 示例 1：
//
// 输入：s1 = "ab" s2 = "eidbaooo"
// 输出：true
// 解释：s2 包含 s1 的排列之一 ("ba").
// 示例 2：
//
// 输入：s1= "ab" s2 = "eidboaoo"
// 输出：false
func checkInclusion250324(s1 string, s2 string) bool {
	// 如果s2长度小于s1，肯定不存在
	m, n := len(s1), len(s2)
	fmt.Println(m, n)
	if n < m {
		return false
	}

	// 初始化窗口
	cnt1, cnt2 := map[byte]int{}, map[byte]int{}
	for i := range s1 {
		// 窗口右移
		cnt1[s1[i]]++
		cnt2[s2[i]]++
	}

	// 达到窗口的条件，判断是否相等
	if maps.Equal(cnt1, cnt2) {
		return true
	}

	for i := m; i < n; i++ {
		// 窗口右移
		cnt2[s2[i]]++
		cnt2[s2[i-m]]--
		if cnt2[s2[i-m]] == 0 {
			delete(cnt2, s2[i-m])
		}
		fmt.Println("cnt1:", cnt1)
		fmt.Println("cnt2:", cnt2)
		if maps.Equal(cnt1, cnt2) {
			return true
		}
	}
	return false
}
