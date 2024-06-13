package leetcode_go

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

func merge20240612(nums1 []int, m int, nums2 []int, n int) {
	ans := make([]int, 0, m+n)
	i, j := 0, 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			ans = append(ans, nums1[i])
			i++
		} else if nums1[i] > nums2[j] {
			ans = append(ans, nums2[j])
			j++
		} else {
			ans = append(ans, nums1[i], nums2[j])
			i++
			j++
		}
	}
	if i < m {
		ans = append(ans, nums1[i:]...)
	}
	if j < n {
		ans = append(ans, nums2[j:]...)
	}
	copy(nums1, ans)
}
