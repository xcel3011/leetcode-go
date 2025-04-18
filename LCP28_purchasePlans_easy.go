package leetcode_go

import (
	"slices"
)

// 小力将 N 个零件的报价存于数组 nums。
// 小力预算为 target，假定小力仅购买两个零件，要求购买零件的花费不超过预算，请问他有多少种采购方案。
// 注意：答案需要以 1e9 + 7 (1000000007) 为底取模，如：计算初始结果为：1000000008，请返回 1
//
// 示例 1：
// 输入：nums = [2,5,3,5], target = 6
// 输出：1
// 解释：预算内仅能购买 nums[0] 与 nums[2]。
//
// 示例 2：
// 输入：nums = [2,2,1,9], target = 10
// 输出：4
// 解释：符合预算的采购方案如下：
// nums[0] + nums[1] = 4
// nums[0] + nums[2] = 3
// nums[1] + nums[2] = 3
// nums[2] + nums[3] = 10
//
// 提示：
// 2 <= nums.length <= 10^5
// 1 <= nums[i], target <= 10^5
func purchasePlans(nums []int, target int) int {
	slices.Sort(nums)
	ans := 0
	l, r := 0, len(nums)-1
	for l < r {
		sum := nums[l] + nums[r]
		if sum > target {
			r--
		} else {
			ans += r - l
			l++
		}
	}
	return ans % (1e9 + 7)
}
