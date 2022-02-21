package array

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

func merge1(nums1 []int, m int, nums2 []int, n int) {
	res := make([]int, 0, m+n)
	i, j := 0, 0
	for {
		if i == m {
			res = append(res, nums2[j:]...)
			break
		}
		if j == m {
			res = append(res, nums1[i:]...)
		}
		if nums1[i] < nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}
	copy(nums1, res)
}
