package leetcode_go

func subsetXORSum(nums []int) int {
	var dfs func(val, idx int) int
	dfs = func(val, idx int) int {
		if idx == len(nums) {
			return val
		}
		return dfs(val^nums[idx], idx+1) + dfs(val, idx+1)
	}

	return dfs(0, 0)
}
