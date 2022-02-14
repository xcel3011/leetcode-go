package array

func moveZeroes(nums []int) {
	index := removeElement(nums, 0)
	for ; index < len(nums); index++ {
		nums[index] = 0
	}
}
