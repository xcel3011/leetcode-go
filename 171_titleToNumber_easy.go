package leetcode_go

import (
	"math"
)

func titleToNumber(columnTitle string) int {
	ans := 0
	n := len(columnTitle)
	for i := n - 1; i >= 0; i-- {
		ans += int(math.Pow(26, float64(n-1-i))) * int(columnTitle[i]-'A'+1)
	}

	return ans
}
