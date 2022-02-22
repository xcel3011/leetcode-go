package array

func subsets(nums []int) [][]int {
	var res [][]int
	res = append(res, []int{})

	for i := 0; i < len(nums); i++ {

		all := len(res)
		for j := 0; j < all; j++ {
			res = append(res, append([]int{}, append(res[j], nums[i])...))
		}
	}

	return res
}
