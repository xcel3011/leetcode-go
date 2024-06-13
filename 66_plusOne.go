package leetcode_go

import (
	"slices"
)

func plusOne(digits []int) []int {
	carry := 0
	n := 0
	l := len(digits) - 1
	var res []int
	for ; l >= 0; l-- {
		if l == len(digits)-1 {
			n = digits[l] + 1 + carry
		} else {
			n = digits[l] + carry
		}

		carry = n / 10
		n = n % 10
		res = append([]int{n}, res...)
	}

	if carry != 0 {
		res = append([]int{carry}, res...)
	}
	return res
}

func plusOne20240612(digits []int) []int {
	ans := make([]int, 0, len(digits))
	carry := 0
	for i := len(digits) - 1; i >= 0; i-- {
		n := digits[i] + 1
		if i != len(digits)-1 {
			n = digits[i] + carry
		}

		carry = n / 10
		n = n % 10
		ans = append(ans, n)
	}
	if carry != 0 {
		ans = append(ans, carry)
	}
	slices.Reverse(ans)
	return ans
}
