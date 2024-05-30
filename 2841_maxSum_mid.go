package leetcode_go

func maxSum(nums []int, m int, k int) int64 {
	// 用map标记窗口内不同的数字
	window := make(map[int]int)

	// 初始化窗口
	var sum, ans int64
	for i := 0; i < k; i++ {
		window[nums[i]]++
		sum += int64(nums[i])
	}
	if len(window) >= m {
		ans = sum
	}

	// 从第一个窗口开始向后移动

	for i := k; i < len(nums); i++ {
		// 右侧新加入窗口
		window[nums[i]]++
		sum += int64(nums[i])

		// 左侧滑出窗口
		window[nums[i-k]]--
		if window[nums[i-k]] == 0 {
			delete(window, nums[i-k])
		}
		sum -= int64(nums[i-k])

		if len(window) >= m {
			ans = max(ans, sum)
		}
	}
	return ans
}
