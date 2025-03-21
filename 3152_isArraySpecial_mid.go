package leetcode_go

func isArraySpecial(nums []int, queries [][]int) []bool {
	cnt := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		cnt[i] = cnt[i-1]
		if nums[i]%2 == nums[i-1]%2 {
			cnt[i]++
		}
	}
	res := make([]bool, len(queries))
	for i := 0; i < len(queries); i++ {
		res[i] = cnt[queries[i][0]] == cnt[queries[i][1]]
	}
	return res
}
