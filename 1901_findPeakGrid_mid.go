package leetcode_go

import (
	"slices"
)

func findPeakGrid(mat [][]int) []int {
	left, right := -1, len(mat[0])-1
	for left+1 < right {
		mid := left + (right-left)/2
		// 找到mid这一列的最大值
		mx := slices.Max(mat[mid])
		if mx > mat[mid+1][slices.Index(mat[mid], mx)] {
			// mid行的最大值比下一行大，吧right标记为蓝色
			right = mid
		} else {
			// mid行的最大值比下一行小，吧left标记为红色
			left = mid
		}
	}
	return []int{right, slices.Index(mat[right], slices.Max(mat[right]))}
}
