package leetcode_go

import (
	"math"
)

func minimumRecolors(blocks string, k int) int {
	count := 0
	for i := 0; i < k; i++ {
		if blocks[i] == 'W' {
			count++
		}
	}
	ans := count
	for i := k; i < len(blocks); i++ {
		if blocks[i] == 'W' {
			count++
		}
		if blocks[i-k] == 'W' {
			count--
		}
		ans = min(ans, count)

	}
	return ans
}

// 给你一个长度为 n 下标从 0 开始的字符串 blocks ，
// blocks[i] 要么是 'W' 要么是 'B' ，表示第 i 块的颜色。字符 'W' 和 'B' 分别表示白色和黑色。
//
// 给你一个整数 k ，表示想要 连续 黑色块的数目。
//
// 每一次操作中，你可以选择一个白色块将它 涂成 黑色块。
//
// 请你返回至少出现 一次 连续 k 个黑色块的 最少 操作次数。
func minimumRecolors250321(blocks string, k int) int {
	ans, cnt := math.MaxInt, 0
	for i, x := range blocks {
		if x == 'W' {
			cnt++
		}
		if i < k-1 {
			continue
		}
		ans = min(ans, cnt)
		if blocks[i-k+1] == 'W' {
			cnt--
		}
	}
	return ans
}
