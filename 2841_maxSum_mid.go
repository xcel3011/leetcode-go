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

// 给你一个整数数组 nums 和两个正整数 m 和 k 。
//
// 请你返回 nums 中长度为 k 的 几乎唯一 子数组的 最大和 ，如果不存在几乎唯一子数组，请你返回 0 。
//
// 如果 nums 的一个子数组有至少 m 个互不相同的元素，我们称它是 几乎唯一 子数组。
//
// 子数组指的是一个数组中一段连续 非空 的元素序列。
func maxSum020321(nums []int, m int, k int) int64 {
	cnt := make(map[int]int)
	var ans, sum int64
	for i, in := range nums {
		cnt[in]++
		sum += int64(in)

		if i < k-1 {
			continue
		}
		if len(cnt) >= m {
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
