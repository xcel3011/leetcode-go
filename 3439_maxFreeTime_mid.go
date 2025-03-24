package leetcode_go

// 给你一个整数 eventTime 表示一个活动的总时长，这个活动开始于 t = 0 ，结束于 t = eventTime 。
//
// 同时给你两个长度为 n 的整数数组 startTime 和 endTime 。它们表示这次活动中 n 个时间 没有重叠 的会议，
// 其中第 i 个会议的时间为 [startTime[i], endTime[i]] 。
//
// 你可以重新安排 至多 k 个会议，安排的规则是将会议时间平移，且保持原来的 会议时长 ，
// 你的目的是移动会议后 最大化 相邻两个会议之间的 最长 连续空余时间。
//
// 移动前后所有会议之间的 相对 顺序需要保持不变，而且会议时间也需要保持互不重叠。
//
// 请你返回重新安排会议以后，可以得到的 最大 空余时间。
//
// 注意，会议 不能 安排到整个活动的时间以外。
func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
	n := len(startTime)
	getTime := func(t int) int {
		if t == 0 {
			return startTime[0]
		} else if t == n {
			return eventTime - endTime[n-1]
		} else {
			return startTime[t] - endTime[t-1]
		}
	}

	ans, sum := 0, 0
	for i := 0; i <= n; i++ {
		sum += getTime(i)
		if i < k {
			continue
		}
		ans = max(ans, sum)
		sum -= getTime(i - k)
	}

	return ans
}
