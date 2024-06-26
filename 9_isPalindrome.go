package leetcode_go

import "strconv"

func isPalindrome9(x int) bool {
	str := strconv.Itoa(x)
	left, right := 0, len(str)-1

	for right > left {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}

	return true
}
