package leetcode_go

// dp
// f[i]=f[i-len(word)]+1
func maxRepeating(sequence string, word string) int {
	n, m := len(sequence), len(word)
	if n < m {
		return 0
	}
	f := make([]int, n)
	count := 0
	for i := m - 1; i < n; i++ {
		if sequence[i-m+1:i+1] == word {
			if i == m-1 {
				f[i] = 1
			} else {
				f[i] = f[i-m] + 1
			}
			if f[i] > count {
				count = f[i]
			}
		}
	}
	return count
}
