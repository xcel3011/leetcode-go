package leetcode_go

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) (ans []string) {
	n := len(nums)
	for i := 0; i < n; {
		start := i
		i++
		for i < n && nums[i]-nums[i-1] == 1 {
			i++
		}
		if nums[start] == nums[i-1] {
			ans = append(ans, strconv.Itoa(nums[start]))
		} else {
			ans = append(ans, fmt.Sprintf("%d->%d", nums[start], nums[i]-1))
		}

	}
	return
}
