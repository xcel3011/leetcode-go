package string

func minInsertions(s string) int {
	res, need := 0, 0

	for i := range s {
		if s[i] == '(' {
			need += 2
			if need%2 == 1 {
				res++
				need--
			}
		}
		if s[i] == ')' {
			need--
			if need == -1 {
				res++
				need = 1
			}
		}
	}

	return res + need
}
