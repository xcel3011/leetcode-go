package leetcode_go

// 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
//
// 示例 1：
// 输入：nums = [-4,-1,0,3,10]
// 输出：[0,1,9,16,100]
// 解释：平方后，数组变为 [16,1,0,9,100]
// 排序后，数组变为 [0,1,9,16,100]
//
// 示例 2：
// 输入：nums = [-7,-3,2,3,11]
// 输出：[4,9,9,49,121]
//
// 提示：
// 1 <= nums.length <= 104
// -104 <= nums[i] <= 104
// nums 已按 非递减顺序 排序
//
// 进阶：
// 请你设计时间复杂度为 O(n) 的算法解决本问题
func sortedSquares(nums []int) []int {
	//for i := 0; i < len(nums); i++ {
	//	nums[i] *= nums[i]
	//}
	//slices.Sort(nums)
	//return nums
	//平方后两边往中间递减
	n := len(nums)
	ans := make([]int, n)
	l, r := 0, n-1
	for i := n - 1; i >= 0; i-- {
		v := nums[l] * nums[l]
		w := nums[r] * nums[r]
		if v > w {
			ans[i] = v
			l++
		} else {
			ans[i] = w
			r--
		}
	}
	return ans
}
