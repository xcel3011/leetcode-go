package leetcode_go

import (
	"math"
)

// 给你一个下标从 0 开始、大小为 m x n 的二维矩阵 grid ，请你求解大小同样为 m x n 的答案矩阵 answer 。
//
// 矩阵 answer 中每个单元格 (r, c) 的值可以按下述方式进行计算：
//
// 令 topLeft[r][c] 为矩阵 grid 中单元格 (r, c) 左上角对角线上 不同值 的数量。
// 令 bottomRight[r][c] 为矩阵 grid 中单元格 (r, c) 右下角对角线上 不同值 的数量。
// 然后 answer[r][c] = |topLeft[r][c] - bottomRight[r][c]| 。
//
// 返回矩阵 answer 。
//
// 矩阵对角线 是从最顶行或最左列的某个单元格开始，向右下方向走到矩阵末尾的对角线。
//
// 如果单元格 (r1, c1) 和单元格 (r, c) 属于同一条对角线且 r1 < r ，则单元格 (r1, c1) 属于单元格 (r, c) 的左上对角线。类似地，可以定义右下对角线。
func differenceOfDistinctValues(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 统计左上角值
			cnt1 := map[int]struct{}{}
			r1, c1 := i-1, j-1
			for r1 >= 0 && c1 >= 0 {
				cnt1[grid[r1][c1]] = struct{}{}
				r1--
				c1--
			}

			// 统计右下角值
			cnt2 := map[int]struct{}{}
			r2, c2 := i+1, j+1
			for r2 < m && c2 < n {
				cnt2[grid[r2][c2]] = struct{}{}
				r2++
				c2++
			}
			ans[i] = append(ans[i], int(math.Abs(float64(len(cnt1)-len(cnt2)))))
		}
	}
	return ans
}
