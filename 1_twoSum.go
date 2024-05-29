package leetcode_go

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		m[num] = i
	}

	for i, num := range nums {
		other := target - num
		index, ok := m[other]
		if ok && i != index {
			return []int{i, index}
		}
	}

	return nil
}
