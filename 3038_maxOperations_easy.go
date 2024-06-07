package leetcode_go

import (
	"fmt"
)

func maxOperations(nums []int) int {
	ans := 0
	sum := 0
	for i := 0; i < len(nums); i += 2 {
		if i+1 >= len(nums) {
			break
		}
		fmt.Println(nums[i], nums[i+1], sum, ans)
		if i == 0 {
			ans++
			sum = nums[0] + nums[1]
			continue
		}
		if nums[i]+nums[i+1] > sum {
			ans++
		}
	}
	return ans
}
