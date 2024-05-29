package leetcode_go

func reverseString(s []byte) {
	lo, hi := 0, len(s)-1
	for hi > lo {
		s[lo], s[hi] = s[hi], s[lo]
		lo++
		hi--
	}
}
