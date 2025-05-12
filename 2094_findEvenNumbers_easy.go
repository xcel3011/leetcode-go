package leetcode_go

import (
	"slices"
)

func findEvenNumbers(digits []int) []int {
	cnt := make(map[int]int)
	for _, d := range digits {
		cnt[d]++
	}
	var ans []int
next:
	for i := 100; i < 1000; i += 2 {
		c := make(map[int]int)
		for x := i; x > 0; x /= 10 {
			d := x % 10
			c[d]++
			if c[d] > cnt[d] {
				continue next
			}
		}
		ans = append(ans, i)
	}
	return ans
}

func findEvenNumbers1(digits []int) []int {
	ans := []int{}
	for i := 0; i < len(digits); i++ {
		for j := 0; j < len(digits); j++ {
			for k := 0; k < len(digits); k++ {
				if i == k || j == k || i == j || digits[i] == 0 {
					continue
				}
				num := digits[i]*100 + digits[j]*10 + digits[k]
				if num%2 == 0 {
					ans = append(ans, num)
				}
			}
		}
	}
	slices.Sort(ans)

	return slices.Compact(ans)
}
