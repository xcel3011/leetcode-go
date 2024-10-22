package leetcode_go

func countCompleteDayPairs(hours []int) int {
	var res int
	for i := 0; i < len(hours); i++ {
		for j := i + 1; j < len(hours); j++ {
			if (hours[i]+hours[j])%24 == 0 {
				res++
			}
		}
	}
	return res
}
