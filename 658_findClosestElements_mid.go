package leetcode_go

// 给定一个 排序好 的数组 arr ，两个整数 k 和 x ，从数组中找到最靠近 x（两数之差最小）的 k 个数。返回的结果必须要是按升序排好的。
// 整数 a 比整数 b 更接近 x 需要满足：
// |a - x| < |b - x| 或者
// |a - x| == |b - x| 且 a < b
//
// 示例 1：
// 输入：arr = [1,2,3,4,5], k = 4, x = 3
// 输出：[1,2,3,4]
//
// 示例 2：
// 输入：arr = [1,1,2,3,4,5], k = 4, x = -1
// 输出：[1,1,2,3]
//
// 提示：
// 1 <= k <= arr.length
// 1 <= arr.length <= 104
// arr 按 升序 排列
// -104 <= arr[i], x <= 104
func findClosestElements(arr []int, k int, x int) []int {
	_abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	// 从0到k-1开始滑动
	a, b := 0, k-1
	for i := k; i < len(arr); i++ {
		// 遇到
		if _abs(arr[i]-x) < _abs(arr[a]-x) {
			b = i
			a = b - k + 1
		}
	}
	return arr[a : b+1]
}
