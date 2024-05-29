package leetcode_go

func moveZeroes(nums []int) {
	index := removeElement(nums, 0)
	for ; index < len(nums); index++ {
		nums[index] = 0
	}
}
