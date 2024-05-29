package leetcode_go

func missingRolls(rolls []int, mean int, n int) []int {
	var sum int
	for i := 0; i < len(rolls); i++ {
		sum += rolls[i]
	}
	diff := (n+len(rolls))*mean - sum
	if diff < n || diff > 6*n {
		return nil
	}

	item := diff / n
	remainder := diff % n
	nn := make([]int, n)
	for i := 0; i < n; i++ {
		if i < remainder {
			nn[i] = item + 1
		} else {
			nn[i] = item
		}
	}
	return nn
}
