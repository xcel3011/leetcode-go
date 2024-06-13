package leetcode_go

import (
	"strconv"
)

func addBinary(a string, b string) string {
	ans := ""
	m, n := len(a), len(b)
	i, j := m-1, n-1
	carry := 0
	for i >= 0 || j >= 0 {
		sum := 0
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		sum += carry
		ans = strconv.Itoa(sum%2) + ans
		carry = sum / 2
	}
	if carry > 0 {
		ans = strconv.Itoa(carry) + ans
	}
	return ans
}
