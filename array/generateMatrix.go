package array

func generateMatrix(n int) [][]int {
	var (
		matrix                    = make([][]int, n)
		upper, lower, left, right = 0, n - 1, 0, n - 1
		count                     = 1
	)

	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for count <= n*n {
		if upper <= lower {
			for i := left; i <= right; i++ {
				matrix[upper][i] = count
				count++
			}
			upper++
		}

		if right >= left {
			for i := upper; i <= lower; i++ {
				matrix[i][right] = count
				count++
			}
			right--
		}

		if lower >= upper {
			for i := right; i >= left; i-- {
				matrix[lower][i] = count
				count++
			}
			lower--
		}

		if left <= right {
			for i := lower; i >= upper; i-- {
				matrix[i][left] = count
				count++
			}
			left++
		}
	}

	return matrix
}
