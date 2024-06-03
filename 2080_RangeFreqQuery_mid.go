package leetcode_go

import (
	"sort"
)

type RangeFreqQuery map[int][]int

func Constructor2080(arr []int) RangeFreqQuery {
	pos := map[int][]int{}
	for i := 0; i < len(arr); i++ {
		pos[arr[i]] = append(pos[arr[i]], i)
	}
	return pos
}

func (pos RangeFreqQuery) Query(left int, right int, value int) int {
	a := pos[value]
	return sort.SearchInts(a, right+1) - sort.SearchInts(a, left)
}
