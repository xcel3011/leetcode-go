package leetcode_go

import (
	"slices"
)

// 你有 k 个 非递减排列 的整数列表。
// 找到一个 最小 区间，使得 k 个列表中的每个列表至少有一个数包含在其中。
// 我们定义如果 b-a < d-c 或者在 b-a == d-c 时 a < c，则区间 [a,b] 比 [c,d] 小。
//
// 示例 1：
// 输入：nums = [[4,10,15,24,26], [0,9,12,20], [5,18,22,30]]
// 输出：[20,24]
// 解释：
// 列表 1：[4, 10, 15, 24, 26]，24 在区间 [20,24] 中。
// 列表 2：[0, 9, 12, 20]，20 在区间 [20,24] 中。
// 列表 3：[5, 18, 22, 30]，22 在区间 [20,24] 中。
//
// 示例 2：
// 输入：nums = [[1,2,3],[1,2,3],[1,2,3]]
// 输出：[1,1]
//
// 提示：
// nums.length == k
// 1 <= k <= 3500
// 1 <= nums[i].length <= 50
// -105 <= nums[i][j] <= 105
// nums[i] 按非递减顺序排列
func smallestRange(nums [][]int) []int {
	type pair struct{ x, i int }
	pairs := []pair{}
	for i, arr := range nums {
		for _, x := range arr {
			pairs = append(pairs, pair{x, i})
		}
	}
	slices.SortFunc(pairs, func(a, b pair) int { return a.x - b.x })

	ansL, ansR := pairs[0].x, pairs[len(pairs)-1].x
	empty := len(nums)
	cnt := make([]int, empty)
	left := 0
	for _, p := range pairs {
		r, i := p.x, p.i
		if cnt[i] == 0 { // 包含 nums[i] 的数字
			empty--
		}
		cnt[i]++
		for empty == 0 { // 每个列表都至少包含一个数
			l, i := pairs[left].x, pairs[left].i
			if r-l < ansR-ansL {
				ansL, ansR = l, r
			}
			cnt[i]--
			if cnt[i] == 0 {
				// 不包含 nums[i] 的数字
				empty++
			}
			left++
		}
	}
	return []int{ansL, ansR}
}
