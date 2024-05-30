package leetcode_go

func findMaxAverage1(nums []int, k int) (ans float64) {
	var sum int
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	ans = float64(sum) / float64(k)

	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		ans = max(ans, float64(sum)/float64(k))
	}
	return
}
