package leetcode_go

import (
	"sort"
)

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	rec := append(sort.IntSlice{}, nums1...)
	sort.Ints(rec)
	abs := func(n int) int {
		if n >= 0 {
			return n
		}
		return -n
	}
	sum, maxn, n := 0, 0, len(nums2)
	for i := 0; i < n; i++ {
		diff := abs(nums1[i] - nums2[i])
		sum += diff
		j := rec.Search(nums2[i])
		if j < n {
			maxn = max(maxn, diff-(rec[j]-nums2[i]))
		}
		if j > 0 {
			maxn = max(maxn, diff-(nums2[i]-rec[j-1]))
		}
	}
	return (sum - maxn) % (1e9 + 7)
}
