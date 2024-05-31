package leetcode_go

func countHomogenous(s string) int {
	count := func(n int) int {
		return (n * (n + 1) / 2) % 1000000007
	}
	ans := 0
	n := len(s)
	for i := 0; i < n; {
		start := i
		i++
		for i < n && s[i] == s[i-1] {
			i++
		}
		ans += count(i - start)
	}
	return ans % 1000000007
}
