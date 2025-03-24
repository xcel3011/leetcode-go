package leetcode_go

// 给你一个字符串 s ，请你返回满足以下条件且出现次数最大的 任意 子串的出现次数：
//
// 子串中不同字母的数目必须小于等于 maxLetters 。
// 子串的长度必须大于等于 minSize 且小于等于 maxSize 。
// 示例 1：
//
// 输入：s = "aababcaab", maxLetters = 2, minSize = 3, maxSize = 4
// 输出：2
// 解释：子串 "aab" 在原字符串中出现了 2 次。
// 它满足所有的要求：2 个不同的字母，长度为 3 （在 minSize 和 maxSize 范围内）。
// 示例 2：
//
// 输入：s = "aaaa", maxLetters = 1, minSize = 3, maxSize = 3
// 输出：2
// 解释：子串 "aaa" 在原字符串中出现了 2 次，且它们有重叠部分。
// 示例 3：
//
// 输入：s = "aabcabcab", maxLetters = 2, minSize = 2, maxSize = 3
// 输出：3
// 示例 4：
//
// 输入：s = "abcde", maxLetters = 2, minSize = 3, maxSize = 3
// 输出：0
func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	m := make(map[string]int)
	for i := minSize; i <= maxSize; i++ {
		chars := make(map[byte]int)
		for j := 0; j < len(s); j++ {
			chars[s[j]]++
			// fmt.Println(chars)
			if j < i-1 {
				continue
			}
			if len(chars) <= maxLetters {
				m[s[j-i+1:j+1]]++
			}
			chars[s[j-i+1]]--
			if chars[s[j-i+1]] == 0 {
				delete(chars, s[j-i+1])
			}
		}
	}

	ans := 0
	// fmt.Println(m)
	for _, count := range m {
		ans = max(ans, count)
	}
	return ans
}
