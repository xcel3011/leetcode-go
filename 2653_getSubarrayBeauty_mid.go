package leetcode_go

// 假设我们有一个数列或集合，其中的元素是按某种顺序排列的（通常是从小到大）。"第 x 小的数"就是这个数列或集合中按顺序排在第 x 个位置的数。
//
// 例如，给定数列 {3, 1, 4, 1, 5, 9}，按从小到大的顺序排列后得到 {1, 1, 3, 4, 5, 9}。在这个排序后的数列中，第 3 小的数是 3。
func getSubarrayBeauty(nums []int, k, x int) []int {
	const bias = 50
	cnt := [bias*2 + 1]int{}
	for _, num := range nums[:k-1] { // 先往窗口内添加 k-1 个数
		cnt[num+bias]++
	}
	ans := make([]int, len(nums)-k+1)
	for i, num := range nums[k-1:] {
		cnt[num+bias]++ // 进入窗口（保证窗口有恰好 k 个数）
		left := x
		for j, c := range cnt[:bias] { // 暴力枚举负数范围 [-50,-1]
			left -= c
			if left <= 0 { // 找到美丽值
				ans[i] = j - bias
				break
			}
		}
		cnt[nums[i]+bias]-- // 离开窗口
	}
	return ans
}
