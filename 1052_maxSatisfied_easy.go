package leetcode_go

func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	// 不连续冷静的时候最多顾客
	sum := 0
	for i := 0; i < len(customers); i++ {
		if grumpy[i] == 0 {
			sum += customers[i]
		}
	}

	// 如果最开始的minutes冷静，会多多少顾客
	for i := 0; i < minutes; i++ {
		if grumpy[i] == 1 {
			sum += customers[i]
		}
	}

	// 窗口向后滑动找到最大的
	ans := sum
	for i := minutes; i < len(customers); i++ {
		if grumpy[i] == 1 {
			sum += customers[i]
		}
		if grumpy[i-minutes] == 1 {
			sum -= customers[i-minutes]
		}
		ans = max(ans, sum)
	}
	return ans
}
