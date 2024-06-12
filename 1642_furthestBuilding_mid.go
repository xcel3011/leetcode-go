package leetcode_go

import (
	"slices"
	"sort"
)

func furthestBuilding(heights []int, bricks int, ladders int) int {
	n := len(heights)
	// 构建辅助数组
	// 将每个height和前一个的差保存
	diff := make([]int, n-1)
	for i, x := range heights[1:] {
		if x > heights[i] {
			diff[i] = x - heights[i]
		}
	}
	// 闭区间[0,n-1]
	// 最坏情况一步也不能走,最好的情况可以走到末尾
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)/2 // 防止溢出
		// 如果梯子大于等于当前步数,说明这几个都可以过
		if ladders >= mid {
			l = mid + 1
			continue
		}

		// 取辅助数组的前mid个排序
		tmp := slices.Clone(diff[:mid])
		sort.Ints(tmp)
		// 消耗转最多的几个用梯子,剩下的用砖
		sum := 0
		for _, h := range tmp[:mid-ladders] {
			sum += h
		}

		if sum <= bricks {
			// 需要总砖sum小于等于需要的bricks说明mid步可以走
			l = mid + 1 //[mid+1,r]
		} else {
			r = mid - 1 //[l,mid-1]
		}
	}

	return r
}
