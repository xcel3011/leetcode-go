package leetcode_go

func maximumCandies(candies []int, k int64) int {
	l, r := 1, 0
	for _, candy := range candies {
		r += candy
	}
	r = int(int64(r) / k)

	// 闭区间[l,r]
	for l <= r {
		mid := l + (r-l)/2
		cnt := 0

		// 判断当前candy能分多少人
		for _, candy := range candies {
			cnt += candy / mid
		}

		if cnt >= int(k) {
			// 如果分的数量大于k
			// 让left向右移动,给每个人分的糖果更多些
			l = mid + 1
		} else {
			// 如果小于k
			// 说明分的糖果过多了需要right向做移动,给每个人分的糖果少些
			r = mid - 1
		}

	}
	return r
}
