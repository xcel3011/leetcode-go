package leetcode_go

func singleNumber(nums []int) int {
	var single int
	for _, num := range nums {
		single ^= num
	}

	return single
}
