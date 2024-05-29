package leetcode_go

func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)

	left, right := 0, 0
	res := 0

	for right < len(s) {
		c := s[right]
		right++
		window[c]++

		for window[c] > 1 {
			d := s[left]
			left++
			window[d]--
		}

		diff := right - left
		if res < diff {
			res = diff
		}
	}

	return res
}
