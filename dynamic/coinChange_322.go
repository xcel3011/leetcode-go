package dynamic

import "math"

func coinChangeRecursion(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	res := math.MaxInt
	for _, coin := range coins {
		sub := coinChangeRecursion(coins, amount-coin)
		// 子问题无解则跳过
		if sub == -1 {
			continue
		}
		// 在子问题中选择最优解，然后加一
		if sub+1 < res {
			res = sub + 1
		}
	}

	if res == math.MaxInt {
		return -1
	}
	return res
}

func coinChange(coins []int, amount int) int {
	// dp 数组的定义：当目标金额为 i 时，至少需要 dp[i] 枚硬币凑出。
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = amount + 1
	}

	// bad case
	dp[0] = 0

	// 外层 for 循环在遍历所有状态的所有取值
	for i := 0; i <= amount; i++ {
		// 内层 for 循环在求所有选择的最小值
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			if 1+dp[i-coin] < dp[i] {
				dp[i] = 1 + dp[i-coin]
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
