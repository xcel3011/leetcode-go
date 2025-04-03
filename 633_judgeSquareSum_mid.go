package leetcode_go

import (
	"math"
)

func judgeSquareSum(c int) bool {
	//for a := 0; a*a < c; a++ {
	//	b := math.Sqrt(float64(c - a*a))
	//	if b == math.Floor(b) {
	//		return true
	//	}
	//}
	//return false
	l := 0
	r := int(math.Sqrt(float64(c)))
	for l <= r {
		sum := l*l + r*r
		if sum == c {
			return true
		} else if sum > c {
			r--
		} else {
			l++
		}
	}
	return false
}

// 给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a2 + b2 = c 。
//
// 示例 1：
// 输入：c = 5
// 输出：true
// 解释：1 * 1 + 2 * 2 = 5
//
// 示例 2：
// 输入：c = 3
// 输出：false
//
// 提示：
// 0 <= c <= 231 - 1
func judgeSquareSum250327(c int) bool {
	// 右边界从C开始会超时
	l, r := 0, int(math.Sqrt(float64(c)))
	for l <= r {
		sum := l*l + r*r
		// fmt.Println(l, r, sum)
		if sum == c {
			return true
		} else if sum > c {
			r--
		} else {
			l++
		}
	}
	return false
}
