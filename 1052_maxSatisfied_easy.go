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

// 有一个书店老板，他的书店开了 n 分钟。每分钟都有一些顾客进入这家商店。给定一个长度为 n 的整数数组 customers ，
// 其中 customers[i] 是在第 i 分钟开始时进入商店的顾客数量，所有这些顾客在第 i 分钟结束后离开。
//
// 在某些分钟内，书店老板会生气。 如果书店老板在第 i 分钟生气，那么 grumpy[i] = 1，否则 grumpy[i] = 0。
//
// 当书店老板生气时，那一分钟的顾客就会不满意，若老板不生气则顾客是满意的。
//
// 书店老板知道一个秘密技巧，能抑制自己的情绪，可以让自己连续 minutes 分钟不生气，但却只能使用一次。
//
// 请你返回 这一天营业下来，最多有多少客户能够感到满意 。
func maxSatisfied250321(customers []int, grumpy []int, minutes int) int {
	// 先计算本来就满意的顾客
	sum := 0
	for i := 0; i < len(customers); i++ {
		if grumpy[i] == 0 {
			sum += customers[i]
		}
	}
	// 计算如果最开始的minutes冷静，会多多少顾客
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
