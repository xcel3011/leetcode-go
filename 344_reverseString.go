package leetcode_go

func reverseString(s []byte) {
	lo, hi := 0, len(s)-1
	for hi > lo {
		s[lo], s[hi] = s[hi], s[lo]
		lo++
		hi--
	}
}

// 编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。
// 不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
//
// 示例 1：
// 输入：s = ["h","e","l","l","o"]
// 输出：["o","l","l","e","h"]
//
// 示例 2：
// 输入：s = ["H","a","n","n","a","h"]
// 输出：["h","a","n","n","a","H"]
//
// 提示：
// 1 <= s.length <= 105
// s[i] 都是 ASCII 码表中的可打印字符
func reverseString250326(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
