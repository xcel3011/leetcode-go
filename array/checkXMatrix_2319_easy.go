package array

func checkXMatrix(grid [][]int) bool {
	for i, row := range grid {
		for j, num := range row {
			if i == j || i+j == len(row)-1 {
				if num == 0 {
					return false
				}
			} else {
				if num != 0 {
					return false
				}
			}
		}
	}
	return true
}
