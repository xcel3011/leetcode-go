package leetcode_go

func minAddToMakeValid(s string) int {
	var res []byte
	of := func(c byte) byte {
		if c == '(' {
			return ' '
		}
		return '('
	}

	for i := range s {
		if len(res) != 0 && res[len(res)-1] == of(s[i]) {
			res = res[:len(res)-1]
		} else {
			res = append(res, s[i])
		}
	}
	return len(res)
}
