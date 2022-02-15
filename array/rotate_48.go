package array

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for _, row := range matrix {
		i, j := 0, len(row)-1
		for j > i {
			row[i], row[j] = row[j], row[i]
			j--
			i++
		}
	}
}
