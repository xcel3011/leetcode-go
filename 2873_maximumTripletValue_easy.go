package leetcode_go

// 给你一个下标从 0 开始的整数数组 nums 。
// 请你从所有满足 i < j < k 的下标三元组 (i, j, k) 中，找出并返回下标三元组的最大值。
// 如果所有满足条件的三元组的值都是负数，则返回 0 。
// 下标三元组 (i, j, k) 的值等于 (nums[i] - nums[j]) * nums[k] 。
//
// 示例 1：
// 输入：nums = [12,6,1,2,7]
// 输出：77
// 解释：下标三元组 (0, 2, 4) 的值是 (nums[0] - nums[2]) * nums[4] = 77 。
// 可以证明不存在值大于 77 的有序下标三元组。
//
// 示例 2：
// 输入：nums = [1,10,3,4,19]
// 输出：133
// 解释：下标三元组 (1, 2, 4) 的值是 (nums[1] - nums[2]) * nums[4] = 133 。
// 可以证明不存在值大于 133 的有序下标三元组。
//
// 示例 3：
// 输入：nums = [1,2,3]
// 输出：0
// 解释：唯一的下标三元组 (0, 1, 2) 的值是一个负数，(nums[0] - nums[1]) * nums[2] = -3 。因此，答案是 0 。
//
// 提示：
// 3 <= nums.length <= 100
// 1 <= nums[i] <= 106
func maximumTripletValue(nums []int) int64 {
	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if (nums[i]-nums[j])*nums[k] > ans {
					ans = (nums[i] - nums[j]) * nums[k]
				}
			}
		}
	}
	return int64(ans)
}
