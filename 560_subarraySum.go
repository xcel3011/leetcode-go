package leetcode_go

func subarraySum(nums []int, k int) int {
	prevSum := make(map[int]int)
	prevSum[0]++

	res, sumI := 0, 0
	for _, num := range nums {
		sumI += num
		sumJ := sumI - k

		if count, ok := prevSum[sumJ]; ok {
			res += count
		}

		prevSum[sumI]++
	}

	return res
}
