package leetcode_go

// https://leetcode.cn/problems/split-array-with-same-average/solutions/1967149/by-lcbin-1dm3/
func splitArraySameAverage(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return false
	}
	s := 0
	for i := range nums {
		s += nums[i]
	}
	for i := range nums {
		nums[i] = nums[i]*n - s
	}

	m := n >> 1
	vis := make(map[int]bool)
	for i := 1; i < 1<<m; i++ {
		t := 0
		for j, v := range nums[:m] {
			if (i >> j & 1) == 1 {
				t += v
			}
		}
		if t == 0 {
			return true
		}
		vis[t] = true
	}
	for i := 1; i < 1<<(n-m); i++ {
		t := 0
		for j, v := range nums[m:] {
			if (i >> j & 1) == 1 {
				t += v
			}
		}
		if t == 0 || (i != (1<<(n-m))-1 && vis[-t]) {
			return true
		}
	}
	return false
}
