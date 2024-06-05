package leetcode_go

import (
	"math"
	"sort"
)

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	// 把数组二排序
	sort.Ints(arr2)
	ans := 0
	for _, x := range arr1 {
		// 找到比x大且最接近x的值
		j := sort.SearchInts(arr2, x+1)
		dis := math.MaxInt

		// 判断j==len(arr2)的特殊情况
		if j < len(arr2) {
			dis = arr2[j] - x
		}
		i := j - 1

		// 判断i==-1的特殊情况
		if i >= 0 && x-arr2[i] < dis {
			dis = x - arr2[i]
		}

		// 判断是否满足条件
		if dis > d {
			ans++
		}
	}
	return ans
}
