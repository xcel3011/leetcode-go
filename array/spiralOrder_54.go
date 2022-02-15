package array

func spiralOrder(matrix [][]int) []int {
	var (
		res                       []int
		n                         = len(matrix)
		upper, lower, left, right = 0, len(matrix) - 1, 0, len(matrix) - 1
	)

	for len(res) < n*n {
		if upper <= right {
			for i := left; i <= right; i++ {
				res = append(res, matrix[upper][i])
			}
			upper++
		}

		if right >= left {
			for i := upper; i <= lower; i++ {
				res = append(res, matrix[i][right])
			}
			right--
		}

		if lower >= upper {
			for i := right; i >= left; i-- {
				res = append(res, matrix[lower][i])
			}
			lower--
		}

		if left <= right {
			for i := lower; i >= upper; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}

	return res
}
