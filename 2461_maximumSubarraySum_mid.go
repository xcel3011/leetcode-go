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

// 给你一个整数数组 nums 和一个整数 k 。请你从 nums 中满足下述条件的全部子数组中找出最大子数组和：
//
// 子数组的长度是 k，且
// 子数组中的所有元素 各不相同 。
// 返回满足题面要求的最大子数组和。如果不存在子数组满足这些条件，返回 0 。
//
// 子数组 是数组中一段连续非空的元素序列。
func maximumSubarraySum250321(nums []int, k int) int64 {
	cnt := make(map[int]int)
	var ans, sum int64
	for i, in := range nums {
		cnt[in]++
		sum += int64(in)
		if i < k-1 {
			continue
		}
		if len(cnt) >= k {
			ans = max(ans, sum)
		}
		out := nums[i-k+1]
		cnt[out]--
		sum -= int64(out)
		if cnt[out] == 0 {
			delete(cnt, out)
		}
	}
	return ans
}
