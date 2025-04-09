package leetcode_go

import (
	"slices"
)

func recoverArray(nums []int) []int {
	// 排序后low[0]肯定是nums[0]
	slices.Sort(nums)

	// 遍历nums[1:]
	for i := 1; i < len(nums); i++ {
		// 跳过重复的nums[i]
		if nums[i] == nums[i-1] {
			continue
		}

		// 从当前i开始，模拟计算k
		d := nums[i] - nums[0]

		// k必须是整数
		if d%2 != 0 {
			continue
		}
		k := d / 2

		// 标记已经找个d
		vis := make([]bool, len(nums))
		vis[i] = true

		// 认为nums[0]和nums[i]还原后是第一个数
		ans := []int{(nums[0] + nums[i]) / 2}

		l, r := 0, i+1
		for r < len(nums) {
			l++
			// 跳过出现在higher中的数
			for vis[l] {
				l++
			}

			// 找 higher
			for r < len(nums) && nums[r]-nums[l] < 2*k {
				r++
			}

			// 不存在满足等式的 higher 值
			if r == len(nums) || nums[r]-nums[l] > 2*k {
				break
			}

			vis[r] = true
			ans = append(ans, (nums[l]+nums[r])/2)
			r++
		}
		if len(ans) == len(nums)/2 {
			return ans
		}
	}
	return nil
}
