package leetcode_go

import (
	"slices"
)

// max(a[0]−2,min(a[0]−1,a[1]),a[2])
func maximumLength(s string) int {
	//"abcaba"
	// 按字母分类，找到连续出现次数
	//对啊来说就是[0][]int{1,1}
	groups := [26][]int{}

	// 标记连续出现次数
	count := 0
	for i := 0; i < len(s); i++ {
		count++
		// 下一个不连续或者是最后一个了，将当前count存入group
		if i+1 == len(s) || s[i] != s[i+1] {
			groups[s[i]-'a'] = append(groups[s[i]-'a'], count)
			count = 0
		}
	}

	ans := 0
	for _, group := range groups {
		// 没出现的字符跳过
		if len(group) == 0 {
			continue
		}

		// 连续出现次数从大到小排序
		slices.SortStableFunc(group, func(a, b int) int {
			return b - a
		})

		// 统一补上第二和第三连续出现个数
		group = append(group, 0, 0)

		// 比较
		ans = max(ans, group[0]-2, min(group[0]-1, group[1]), group[2])
	}

	if ans == 0 {
		return -1
	}

	return ans
}
