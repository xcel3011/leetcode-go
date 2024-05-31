package leetcode_go

func longestBeautifulSubstring(word string) int {
	ans := 0
	n := len(word)
	for i := 0; i < n; {
		// 必须a开始
		if word[i] != 'a' {
			i++
			continue
		}

		// 符合条件的情况
		start := i
		i++
		count := 1
		for i < n && word[i] >= word[i-1] {
			if word[i] > word[i-1] {
				count++
			}
			i++
		}

		if count == 5 {
			ans = max(ans, i-start)
		}
	}
	return ans
}
