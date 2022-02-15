package string

import "math"

func minWindow(s string, t string) string {
	need := make(map[int]int)
	window := make(map[int]int)
	for i := range t {
		need[int(t[i])]++
	}

	left, right := 0, 0
	valid := 0
	start, l := 0, math.MaxInt

	for right < len(s) {
		c := int(s[right])
		right++

		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left < l {
				start = left
				l = right - left
			}

			d := int(s[left])
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if l == math.MaxInt {
		return ""
	}

	return s[start : start+l]
}
