package leetcode_go

func numOfSubarrays(arr []int, k int, threshold int) int {
	sum, count := 0, 0
	for i := 0; i < k; i++ {
		sum += arr[i]
	}
	if sum/k >= threshold {
		count++
	}
	for i := k; i < len(arr); i++ {
		sum += arr[i] - arr[i-k]
		if sum/k >= threshold {
			count++
		}
	}
	return count
}

// 给你一个整数数组 arr 和两个整数 k 和 threshold 。
//
// 请你返回长度为 k 且平均值大于等于 threshold 的子数组数目。
//
// 示例 1：
//
// 输入：arr = [2,2,2,2,5,5,5,8], k = 3, threshold = 4
// 输出：3
// 解释：子数组 [2,5,5],[5,5,5] 和 [5,5,8] 的平均值分别为 4，5 和 6 。其他长度为 3 的子数组的平均值都小于 4 （threshold 的值)。
func numOfSubarrays250321(arr []int, k int, threshold int) int {
	sum, cnt := 0, 0
	for i, num := range arr {
		sum += num
		if i < k-1 {
			continue
		}
		if sum/k >= threshold {
			cnt++
		}
		sum -= arr[i-k+1]
	}
	return cnt
}
