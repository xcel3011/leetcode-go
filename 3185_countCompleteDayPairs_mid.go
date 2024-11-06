package leetcode_go

func countCompleteDayPairs2(hours []int) int64 {
	cnt := [24]int{}
	var ans int64
	for _, h := range hours {
		ans += int64(cnt[24-h%24])
		cnt[h%24]++
	}

	return ans
}
