package array

import (
	"strconv"
)

func calPoints(operations []string) (ans int) {
	var nums []int
	for _, operation := range operations {
		switch operation {
		case "C":
			nums = nums[:len(nums)-1]
		case "D":
			nums = append(nums, nums[len(nums)-1]*2)
		case "+":
			nums = append(nums, nums[len(nums)-2]+nums[len(nums)-1])
		default:
			num, _ := strconv.Atoi(operation)
			nums = append(nums, num)
		}
	}
	for _, num := range nums {
		ans += num
	}
	return
}
