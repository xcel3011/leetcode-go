package leetcode_go

import (
	"math"
)

func maxProfit(prices []int) int {
	minimum := math.MaxInt
	ans := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minimum {
			minimum = prices[i]
		} else {
			ans = max(prices[i]-minimum, ans)
		}
	}
	return ans
}
