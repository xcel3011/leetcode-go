package leetcode_go

func getDescentPeriods(prices []int) int64 {
	count := func(n int) int64 {
		return int64(n) * int64(n+1) / 2
	}
	var ans int64
	n := len(prices)
	for i := 0; i < n; {
		start := i
		i++
		for i < n && prices[i-1]-prices[i] == 1 {
			i++
		}
		ans += count(i - start)
	}
	return ans
}
