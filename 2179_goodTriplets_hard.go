package leetcode_go

type fenwick []int

// a[i] 增加 val
// 1 <= i <= n
func (f fenwick) update(i int, val int) {
	for ; i < len(f); i += i & -i {
		f[i] += val
	}
}

// 求前缀和 a[1] + ... + a[i]
// 1 <= i <= n
func (f fenwick) pre(i int) int {
	sum := 0
	for ; i > 0; i -= i & -i {
		sum += f[i]
	}
	return sum
}

func newFenwickTree(n int) fenwick {
	return make(fenwick, n+1) // 使用下标 1 到 n
}

func goodTriplets(nums1 []int, nums2 []int) (ans int64) {
	n := len(nums1)
	pos := make([]int, n)
	for i, x := range nums1 {
		pos[x] = i
	}

	tree := newFenwickTree(n)
	for i, y := range nums2[:n-1] {
		y = pos[y]
		less := tree.pre(y)
		ans += int64(less) * int64(n-1-y-(i-less))
		tree.update(y+1, 1)
	}
	return
}
