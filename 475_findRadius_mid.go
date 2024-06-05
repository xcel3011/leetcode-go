package leetcode_go

import (
	"math"
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	sort.Ints(heaters)
	ans := 0
	for _, house := range houses {
		j := sort.SearchInts(heaters, house+1)
		dis := math.MaxInt
		if j < len(heaters) {
			dis = heaters[j] - house
		}
		i := j - 1
		if i >= 0 && house-heaters[i] < dis {
			dis = house - heaters[i]
		}
		if dis > ans {
			ans = dis
		}
	}
	return ans
}
