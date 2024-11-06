package leetcode_go

func subarraysDivByK(nums []int, k int) int {
	cnt := map[int]int{0: 1}
	var res, sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		mod := (sum%k + k) % k
		res += cnt[mod]
		cnt[mod]++
	}
	return res
}
