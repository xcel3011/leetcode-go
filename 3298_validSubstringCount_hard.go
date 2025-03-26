package leetcode_go

// 给你两个字符串 word1 和 word2 。
// 如果一个字符串 x 重新排列后，word2 是重排字符串的 前缀 ，那么我们称字符串 x 是 合法的 。
// 请你返回 word1 中 合法 子字符串 的数目。
// 注意 ，这个问题中的内存限制比其他题目要 小 ，所以你 必须 实现一个线性复杂度的解法。
//
// 示例 1：
// 输入：word1 = "bcca", word2 = "abc"
// 输出：1
// 解释：
// 唯一合法的子字符串是 "bcca" ，可以重新排列得到 "abcc" ，"abc" 是它的前缀。
//
// 示例 2：
// 输入：word1 = "abcabc", word2 = "abc"
// 输出：10
// 解释：
// 除了长度为 1 和 2 的所有子字符串都是合法的。
//
// 示例 3：
// 输入：word1 = "abcabc", word2 = "aaabc"
// 输出：0
//
// 解释：
// 1 <= word1.length <= 106
// 1 <= word2.length <= 104
// word1 和 word2 都只包含小写英文字母。
func validSubstringCount(word1 string, word2 string) int64 {
	if len(word1) < len(word2) {
		return 0
	}

	cnt := map[byte]int{}
	for i := 0; i < len(word2); i++ {
		cnt[word2[i]]++
	}
	less := 0
	for _, v := range cnt {
		if v > 0 {
			less++
		}
	}
	l, r := 0, 0
	ans := int64(0)
	for r < len(word1) {
		in := word1[r]
		r++
		cnt[in]--
		if cnt[in] == 0 {
			less--
		}
		for less == 0 {
			if cnt[word1[l]] == 0 {
				less++
			}
			cnt[word1[l]]++
			l++
		}
		ans += int64(l)
	}

	return ans
}
