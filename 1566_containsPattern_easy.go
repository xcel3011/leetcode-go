package leetcode_go

func containsPattern(arr []int, m int, k int) bool {
	for i := 0; i <= len(arr)-m*k; i++ {
		var offset int
		for ; offset < m*k; offset++ {
			if arr[i+offset] != arr[i+offset%m] {
				break
			}
		}
		if offset == m*k {
			return true
		}
	}
	return false
}
