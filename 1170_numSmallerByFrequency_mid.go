package leetcode_go

func numSmallerByFrequency(queries []string, words []string) []int {
	f := func(s string) int {
		ch := 'z'
		count := 0
		for _, c := range s {
			if c < ch {
				ch = c
				count = 1
			} else {
				count++
			}
		}
		return count
	}
	count := [12]int{}
	for i := 0; i < len(words); i++ {
		count[f(words[i])]++
	}
	for i := 9; i >= 1; i-- {
		count[i] += count[i+1]
	}
	ans := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		ans[i] = count[f(queries[i])+1]
	}
	return ans
}
