package string

func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}
	var res []byte
	leftOf := func(c byte) byte {
		if c == '}' {
			return '{'
		} else if c == ']' {
			return '['
		}
		return '('
	}
	for i := range s {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			res = append(res, s[i])
		} else {
			if len(res) != 0 && leftOf(s[i]) == res[len(res)-1] {
				res = res[:len(res)-1]
			} else {
				return false
			}
		}
	}

	return len(res) == 0
}
