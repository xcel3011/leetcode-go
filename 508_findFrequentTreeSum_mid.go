package leetcode_go

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findFrequentTreeSum(root *TreeNode) []int {
	cnt := map[int]int{}
	mx := 0
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		sum := root.Val + dfs(root.Left) + dfs(root.Right)
		cnt[sum]++
		if cnt[sum] > mx {
			mx = cnt[sum]
		}
		return sum
	}
	dfs(root)
	ans := make([]int, 0, len(cnt))
	for num := range cnt {
		if cnt[num] == mx {
			ans = append(ans, num)
		}
	}
	return ans
}
