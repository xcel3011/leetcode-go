package leetcode_go

func minSum(nums1 []int, nums2 []int) int64 {
	var (
		sum1, sum2   int64
		zero1, zero2 bool
	)
	for _, num := range nums1 {
		sum1 += int64(num)
		// 先把0替换成1
		if num == 0 {
			sum1++
			zero1 = true
		}
	}
	for _, num := range nums2 {
		sum2 += int64(num)
		// 先把0替换成1
		if num == 0 {
			sum2++
			zero2 = true
		}
	}

	// 如果sum1<sum2且zero1为false，那么sum1不能再加了
	if sum1 < sum2 && !zero1 {
		return -1
	} else if sum1 > sum2 && !zero2 {
		return -1
	}
	return max(sum1, sum2)
}
