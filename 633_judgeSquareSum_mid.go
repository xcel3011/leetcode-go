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
