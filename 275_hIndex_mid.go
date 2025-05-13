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

func hIndex250417(citations []int) int {
	left, right := 1, len(citations)+1 //[1,n+1)左闭右开
	for left < right {
		// 至少有mid篇论文的引用次数>=mid
		mid := left + (right-left)/2

		// 判断第n-mid篇论文的引用次数是否>=mid
		if citations[len(citations)-mid] >= mid {
			left = mid + 1 //[mid+1,right)
		} else {
			right = mid //[left,mid)
		}
	}
	return left - 1
}
