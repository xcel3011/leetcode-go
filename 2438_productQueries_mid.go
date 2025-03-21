package leetcode_go

func productQueries(n int, queries [][]int) []int {
	const mod int = 1e9 + 7
	var cnt []int
	for n > 0 {
		lb := n & (-n) // 找到最低位的1
		cnt = append(cnt, lb)
		n ^= lb // 去掉最低位的1
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		mul := 1
		for _, x := range cnt[q[0] : q[1]+1] {
			mul = mul * x % mod // 累乘
		}
		ans[i] = mul
	}
	return ans
}
