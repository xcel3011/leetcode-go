package leetcode_go

func lengthAfterTransformations(s string, t int) int {
	//超时
	//var b strings.Builder
	//for i := 0; i < t; i++ {
	//	b.Reset()
	//	for j := 0; j < len(s); j++ {
	//		if s[j] != 'z' {
	//			b.WriteByte(s[j] + 1)
	//		} else {
	//			b.WriteString("ab")
	//		}
	//	}
	//	s = b.String()
	//}
	//return len(s) % (1e9 + 7)
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}
	for i := 0; i < t; i++ {
		next := [26]int{}
		next[0] = cnt[25]
		next[1] = (cnt[0] + cnt[25]) % (1e9 + 7)
		for j := 2; j < 26; j++ {
			next[j] = cnt[j-1]
		}
		cnt = next
	}

	ans := 0
	for i := 0; i < len(cnt); i++ {
		ans = (ans + cnt[i]) % (1e9 + 7)
	}
	return ans
}

func lengthAfterTransformations250515(s string, t int) int {
	// 记录字符串中的字符出现次数
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}

	// 模拟t次变换
	for i := 0; i < t; i++ {
		// 记录下次变化后字符串出现的次数
		next := [26]int{}
		//'z'，则将其替换为字符串 "ab"
		next[0] = cnt[25]
		next[1] = (cnt[0] + cnt[25]) % (1e9 + 7)
		for j := 2; j < 26; j++ {
			next[j] = cnt[j-1]
		}
		cnt = next
	}

	// 计算最终长度
	ans := 0
	for i := 0; i < len(cnt); i++ {
		ans = (ans + cnt[i]) % (1e9 + 7)
	}
	return ans
}
