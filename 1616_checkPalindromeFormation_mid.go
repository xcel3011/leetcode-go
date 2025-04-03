package leetcode_go

func checkPalindromeFormation(a string, b string) bool {
	// 检查字符串本身是否为回文串
	checkSelfPalindrome := func(s string) bool {
		left, right := 0, len(s)-1
		for left < right && s[left] == s[right] {
			left++
			right--
		}
		return left >= right
	}

	// 检查两个字符串是否可以拼接成回文串
	checkConcatenation := func(a, b string) bool {
		n := len(a)
		left, right := 0, n-1
		for left < right && a[left] == b[right] {
			left++
			right--
		}
		if left >= right {
			return true
		}
		return checkSelfPalindrome(a[left:right+1]) || checkSelfPalindrome(b[left:right+1])
	}

	return checkConcatenation(a, b) || checkConcatenation(b, a)
}
