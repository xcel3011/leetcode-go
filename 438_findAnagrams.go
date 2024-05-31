package leetcode_go

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
