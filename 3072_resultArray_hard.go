package leetcode_go

import (
	"slices"
	"sort"
)

type fenwick []int

func (f fenwick) add(i int) {
	for ; i < len(f); i += i & -i {
		f[i]++
	}
}

func (f fenwick) pre(i int) int {
	sum := 0
	for ; i > 0; i &= i - 1 {
		sum += f[i]
	}
	return sum
}

func resultArray(nums []int) (ans []int) {
	sorted := slices.Clone(nums)
	slices.Sort(sorted)
	sorted = slices.Compact(sorted)
	m := len(sorted)

	a := nums[:1]
	b := []int{nums[1]}
	t1 := make(fenwick, m+1)
	t2 := make(fenwick, m+1)
	t1.add(sort.SearchInts(sorted, nums[0]) + 1)
	t2.add(sort.SearchInts(sorted, nums[1]) + 1)
	for _, x := range nums[2:] {
		v := sort.SearchInts(sorted, x) + 1
		gc1 := len(a) - t1.pre(v) // greaterCount(a, v)
		gc2 := len(b) - t2.pre(v) // greaterCount(b, v)
		if gc1 > gc2 || gc1 == gc2 && len(a) <= len(b) {
			a = append(a, x)
			t1.add(v)
		} else {
			b = append(b, x)
			t2.add(v)
		}
	}
	return append(a, b...)
}
