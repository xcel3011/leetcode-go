package leetcode_go

func maxOperations2(nums []int) int {
	// dfs[i][j]=max(dfs[i+2][j]+1,dfs[i+1][j-1]+1,dfs[i,j-1]+1)
	helper := func(a []int, target int) int {
		n := len(a)

		memo := make([][]int, n)
		for i := 0; i < n; i++ {
			memo[i] = make([]int, n)
			for j := 0; j < n; j++ {
				memo[i][j] = -1 // 标记未遍历过
			}
		}

		var dfs func(int, int) int
		dfs = func(i, j int) int {
			if i >= j {
				return 0
			}

			p := &memo[i][j]
			if *p != -1 { // 之前计算过
				return *p
			}

			res := 0
			if a[i]+a[i+1] == target {
				res = max(res, dfs(i+2, j)+1)
			}
			if a[i]+a[j] == target {
				res = max(res, dfs(i+1, j-1)+1)
			}
			if a[j]+a[j-1] == target {
				res = max(res, dfs(i, j-2)+1)
			}
			*p = res
			return res
		}
		return dfs(0, n-1)
	}
	n := len(nums)
	return max(
		helper(nums[2:], nums[0]+nums[1]),
		helper(nums[1:n-1], nums[0]+nums[n-1]),
		helper(nums[:n-2], nums[n-1]+nums[n-2]),
	)
}
