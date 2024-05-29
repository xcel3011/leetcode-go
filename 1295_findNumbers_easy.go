package leetcode_go

import "strconv"

func findNumbers(nums []int) (count int) {
	for _, num := range nums {
		if len(strconv.Itoa(num))%2 == 0 {
			count++
		}
	}
	return
}
