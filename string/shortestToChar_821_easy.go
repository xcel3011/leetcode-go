package string

func shortestToChar(s string, c byte) []int {
	min := func(nums ...int) int {
		num := nums[0]
		for i := range nums {
			if nums[i] < num {
				num = nums[i]
			}
		}
		return num
	}

	n := len(s)
	visit := make([]int, n)

	index := -n
	for i := 0; i < n; i++ {
		if s[i] == c {
			index = i
		}
		visit[i] = i - index
	}

	index = 2 * n
	for i := n - 1; i >= 0; i-- {
		if s[i] == c {
			index = i
		}
		visit[i] = min(visit[i], index-i)

	}
	return visit
}
