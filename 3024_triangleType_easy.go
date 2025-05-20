package leetcode_go

func triangleType(nums []int) string {
	if nums[0]+nums[1] <= nums[2] || nums[0]+nums[2] <= nums[1] || nums[1]+nums[2] <= nums[0] {
		return "none"
	} else if nums[0] == nums[1] && nums[1] == nums[2] {
		return "equilateral"
	} else if nums[0] == nums[1] || nums[1] == nums[2] || nums[0] == nums[2] {
		return "isosceles"
	}
	return "scalene"
}
