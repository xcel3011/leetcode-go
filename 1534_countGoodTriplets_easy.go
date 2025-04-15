package leetcode_go

func countGoodTriplets(arr []int, a int, b int, c int) int {
	_abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	ans := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				if _abs(arr[i]-arr[j]) <= a && _abs(arr[j]-arr[k]) <= b && _abs(arr[i]-arr[k]) <= c {
					ans++
				}
			}
		}
	}
	return ans
}
