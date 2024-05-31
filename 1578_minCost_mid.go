package leetcode_go

func minCost(colors string, neededTime []int) int {
	ans := 0
	n := len(colors)
	for i := 0; i < n; {
		start := i
		i++
		totalTime, maxTime := neededTime[start], neededTime[start]
		for i < n && colors[i] == colors[i-1] {
			maxTime = max(maxTime, neededTime[i])
			totalTime += neededTime[i]
			i++
		}
		ans += totalTime - maxTime
	}
	return ans
}
