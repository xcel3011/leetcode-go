package leetcode_go

func longestEqualSubarray(nums []int, k int) int {
	pos := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		pos[nums[i]] = append(pos[nums[i]], i)
	}
	ans := 0
	for _, vec := range pos {
		j := 0
		for i := 0; i < len(vec); i++ {
			if vec[i]-vec[j]-(i-j) > k {
				j++
			}
			ans = max(ans, i-j+1)
		}
	}
	return ans
}
