package leetcode_go

// 给你一个字符串 s 。一个字符串的 分数 定义为相邻字符 ASCII 码差值绝对值的和。
// 请你返回 s 的 分数 。
//
// 示例 1：
// 输入：s = "hello"
// 输出：13
// 解释：
// s 中字符的 ASCII 码分别为：'h' = 104 ，'e' = 101 ，'l' = 108 ，'o' = 111 。所以 s 的分数为 |104 - 101| + |101 - 108| + |108 - 108| + |108 - 111| = 3 + 7 + 0 + 3 = 13 。
//
// 示例 2：
// 输入：s = "zaz"
// 输出：50
// 解释：
// s 中字符的 ASCII 码分别为：'z' = 122 ，'a' = 97 。所以 s 的分数为 |122 - 97| + |97 - 122| = 25 + 25 = 50 。
//
// 提示：
// 2 <= s.length <= 100
// s 只包含小写英文字母。
func scoreOfString(s string) int {
	_abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	_val := func(x byte) int {
		return int(x)
	}
	ans := 0
	for i := 1; i < len(s); i++ {
		ans += _abs(_val(s[i]) - _val(s[i-1]))
	}
	return ans
}
