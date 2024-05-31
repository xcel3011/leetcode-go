package leetcode_go

func winnerOfGame(colors string) bool {
	ans := [2]int{}
	n := len(colors)
	for i := 0; i < n; {
		start := i
		i++
		for i < n && colors[i] == colors[i-1] {
			i++
		}
		ans[colors[start]-'A'] += max(0, i-start-2)
	}
	return ans[0] > ans[1]
}
