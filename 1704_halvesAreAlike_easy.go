package leetcode_go

func halvesAreAlike(s string) bool {
	n := len(s)
	mid := n / 2
	s1, s2 := s[:mid], s[mid:]
	count1, count2 := 0, 0
	m := map[int32]struct{}{97: {}, 101: {}, 105: {}, 111: {}, 117: {}, 65: {}, 69: {}, 73: {}, 79: {}, 85: {}}
	for _, b := range s1 {
		if _, ok := m[b]; ok {
			count1++
		}
	}
	for _, b := range s2 {
		if _, ok := m[b]; ok {
			count2++
		}
	}
	return count1 == count2
}
