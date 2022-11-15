package array

// 正序
func arrayChange(nums []int, operations [][]int) []int {
	idx := make(map[int]int, len(nums))
	for i, num := range nums {
		idx[num] = i
	}

	for _, operation := range operations {
		x, y := operation[0], operation[1]
		i := idx[x]
		nums[i] = y
		delete(idx, x)
		idx[y] = i
	}
	return nums
}
