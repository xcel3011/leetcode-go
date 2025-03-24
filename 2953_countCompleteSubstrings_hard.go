package leetcode_go

// 给你一个字符串 word 和一个整数 k 。
//
// 如果 word 的一个子字符串 s 满足以下条件，我们称它是 完全字符串：
//
// s 中每个字符 恰好 出现 k 次。
// 相邻字符在字母表中的顺序 至多 相差 2 。也就是说，s 中两个相邻字符 c1 和 c2 ，它们在字母表中的位置相差 至多 为 2 。
// 请你返回 word 中 完全 子字符串的数目。
//
// 子字符串 指的是一个字符串中一段连续 非空 的字符序列。

func countCompleteSubstrings(word string, k int) (ans int) {
	f := func(s string, k int) (res int) {
		for m := 1; m <= 26 && k*m <= len(s); m++ {
			cnt := [26]int{}
			check := func() {
				for i := range cnt {
					if cnt[i] > 0 && cnt[i] != k {
						return
					}
				}
				res++
			}
			for right, in := range s {
				cnt[in-'a']++
				if left := right + 1 - k*m; left >= 0 {
					check()
					cnt[s[left]-'a']--
				}
			}
		}
		return
	}

	_abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	for i, n := 0, len(word); i < n; {
		st := i
		for i++; i < n && _abs(int(word[i])-int(word[i-1])) <= 2; i++ {
		}
		ans += f(word[st:i], k)
	}
	return
}
