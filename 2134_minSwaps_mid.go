package leetcode_go

import (
	"fmt"
)

// 窗口是所有1的个数
// 窗口内有多少0就要交换几次
func minSwaps(nums []int) int {
	k := 0 // 窗口大小
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			k++
		}
	}

	// 第一个窗口从下标0开始
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	ans := k - sum

	// 窗口向右滑动
	for i := 1; i < len(nums); i++ {
		sum += nums[(i+k-1)%len(nums)] - nums[i-1]
		ans = min(ans, k-sum)
	}
	return ans
}

// 交换 定义为选中一个数组中的两个 互不相同 的位置并交换二者的值。
//
// 环形 数组是一个数组，可以认为 第一个 元素和 最后一个 元素 相邻 。
//
// 给你一个 二进制环形 数组 nums ，返回在 任意位置 将数组中的所有 1 聚集在一起需要的最少交换次数。
func minSwaps250324(nums []int) int {
	// 统计1的个数
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			k++
		}
	}

	// 将环转换成链
	nums = append(nums, nums...)
	ans, sum := 0, 0
	for i, num := range nums {
		fmt.Println("i:", i, "k:", k)
		sum += num
		if i < k-1 {
			continue
		}
		fmt.Println("sum:", sum, "ans:", ans)
		ans = max(ans, sum)
		if i-k+1 < len(nums) {
			sum -= nums[i-k+1]
		}

	}
	return k - ans
}
