package leetcode_go

func spiralOrder(matrix [][]int) []int {
	var (
		m, n                      = len(matrix), len(matrix[0])
		res                       = make([]int, m*n)
		count                     int
		upper, lower, left, right = 0, len(matrix) - 1, 0, len(matrix[0]) - 1
	)

	for count < n*m {
		if upper <= right {
			for i := left; i <= right; i++ {
				res[count] = matrix[upper][i]
				count++
			}
			upper++
		}

		if right >= left {
			for i := upper; i <= lower; i++ {
				res[count] = matrix[i][right]
				count++
			}
			right--
		}

		if lower >= upper {
			for i := right; i >= left; i-- {
				res[count] = matrix[lower][i]
				count++
			}
			lower--
		}

		if left <= right {
			for i := lower; i >= upper; i-- {
				res[count] = matrix[i][left]
				count++
			}
			left++
		}
	}

	return res
}
