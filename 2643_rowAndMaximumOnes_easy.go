package leetcode_go

func rowAndMaximumOnes(mat [][]int) []int {
	ans := []int{-1, -1}
	for i := range mat {
		cnt := 0
		for _, v := range mat[i] {
			if v == 1 {
				cnt++
			}
		}
		if cnt > ans[1] {
			ans[0] = i
			ans[1] = cnt
		}
	}
	return ans
}
