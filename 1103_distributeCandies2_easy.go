package leetcode_go

func distributeCandies(candies int, num_people int) []int {
	ans := make([]int, num_people)
	n := 0
	for candies > 0 {
		ans[n%num_people] += min(candies, n+1)
		candies -= min(candies, n+1)
		n++
	}
	return ans
}
