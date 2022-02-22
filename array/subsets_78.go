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

func subsets1(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}

	var (
		num int
		n   = len(nums)
	)
	num, nums = nums[n-1], nums[:n-1]

	// 递归算出所有的子集
	res := subsets1(nums)
	size := len(res)
	for i := 0; i < size; i++ {
		res = append(res, append([]int{}, append(res[i], num)...))
	}

	return res
}
