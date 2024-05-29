package leetcode_go

func findShortestSubArray(nums []int) int {
	type entry struct{ cnt, l, r int }
	m := make(map[int]entry)
	for i, num := range nums {
		if v, ok := m[num]; ok {
			v.cnt++
			v.r = i
			m[num] = v
		} else {
			m[num] = entry{1, i, i}
		}
	}

	min := func(nums ...int) int {
		min := nums[0]
		for _, num := range nums {
			if num < min {
				min = num
			}
		}
		return min
	}

	var ans, max int
	for _, e := range m {
		if e.cnt > max {
			ans, max = e.r-e.l+1, e.cnt
		} else if e.cnt == max {
			ans = min(e.r-e.l+1, ans)
		}
	}
	return ans
}
