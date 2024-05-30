package leetcode_go

func maximumSubarraySum(nums []int, k int) int64 {
	// 标记窗口内不同元素个数
	window := make(map[int]int)

	// 初始窗口
	var sum, ans int64
	for i := 0; i < k; i++ {
		sum += int64(nums[i])
		window[nums[i]]++
	}

	// 初始返回
	if len(window) >= k {
		ans = sum
	}

	// 窗口向右滑动
	for i := k; i < len(nums); i++ {
		// 右侧滑入
		sum += int64(nums[i])
		window[nums[i]]++

		// 左侧滑出
		sum -= int64(nums[i-k])
		if window[nums[i-k]]--; window[nums[i-k]] == 0 {
			delete(window, nums[i-k])
		}

		if len(window) >= k {
			ans = max(ans, sum)
		}

	}
	return ans
}
