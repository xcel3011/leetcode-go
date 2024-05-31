package leetcode_go

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
