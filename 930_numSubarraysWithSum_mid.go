package leetcode_go

func numSubarraysWithSum(nums []int, goal int) int {
	m := map[int]int{}
	var res, sum int
	for i := 0; i < len(nums); i++ {
		m[sum]++
		sum += nums[i]
		res += m[sum-goal]

	}
	return res
}

// 给你一个二元数组 nums ，和一个整数 goal ，请你统计并返回有多少个和为 goal 的 非空 子数组。
// 子数组 是数组的一段连续部分。
//
// 示例 1：
// 输入：nums = [1,0,1,0,1], goal = 2
// 输出：4
// 解释：
// 有 4 个满足题目要求的子数组：[1,0,1]、[1,0,1,0]、[0,1,0,1]、[1,0,1]
//
// 示例 2：
// 输入：nums = [0,0,0,0,0], goal = 0
// 输出：15
//
// 提示：
// 1 <= nums.length <= 3 * 104
// nums[i] 不是 0 就是 1
// 0 <= goal <= nums.length
func numSubarraysWithSum250326(nums []int, goal int) int {
	sum1, sum2 := 0, 0
	l1, l2, r := 0, 0, 0
	ans := 0
	for r < len(nums) {
		in := nums[r]
		sum1 += in
		sum2 += in
		r++
		for l1 < r && sum1 > goal {
			out := nums[l1]
			sum1 -= out
			l1++
		}
		for l2 < r && sum2 >= goal {
			out := nums[l2]
			sum2 -= out
			l2++
		}
		ans += l2 - l1
	}
	return ans
}
