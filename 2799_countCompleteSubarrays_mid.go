package leetcode_go

// 给你一个由正整数组成的数组 nums 。
// 如果数组中的某个子数组满足下述条件，则称之为完全子数 ：
// 子数组中不同元素的数目等于整个数组不同元素的数目。
// 返回数组中完全子数组的数目。
// 子数组是数组中的一个连续非空序列。
//
// 示例 1：
// 输入：nums = [1,3,1,2,2]
// 输出：4
// 解释：完全子数组有：[1,3,1,2]、[1,3,1,2,2]、[3,1,2] 和 [3,1,2,2] 。
//
// 示例 2：
// 输入：nums = [5,5,5,5]
// 输出：10
// 解释：数组仅由整数 5 组成，所以任意子数组都满足完全子数组的条件。子数组的总数为 10 。
//
// 提示：
// 1 <= nums.length <= 1000
// 1 <= nums[i] <= 2000
func countCompleteSubarrays(nums []int) int {
	// 统计数组中不同元素的数目
	total := map[int]struct{}{}
	for _, num := range nums {
		total[num] = struct{}{}
	}

	l, r := 0, 0
	ans := 0
	cnt := map[int]int{}
	for r < len(nums) {
		in := nums[r]
		r++
		cnt[in]++
		for len(cnt) == len(total) {
			out := nums[l]
			l++
			cnt[out]--
			if cnt[out] == 0 {
				delete(cnt, out)
			}
		}
		ans += l
	}
	return ans
}
