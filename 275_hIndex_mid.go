package leetcode_go

func hIndex(citations []int) int {
	l, r := 1, len(citations) //[1,n]闭区间
	for l <= r {
		// 循环不变量：
		// left-1 的回答一定为「是」
		// right+1 的回答一定为「否」
		mid := l + (r-l)/2

		// 引用次数最多的 mid 篇论文，引用次数均 >= mid
		if citations[len(citations)-mid] >= mid {
			l = mid + 1 // 询问范围缩小到 [mid+1, right]
		} else {
			r = mid - 1 // 询问范围缩小到 [left, mid-1]
		}
	}
	return r
}
