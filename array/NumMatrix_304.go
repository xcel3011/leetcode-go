package array

type NumMatrix struct {
	PrevSum [][]int
}

func NewNumMatrix(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	if m == 0 || n == 0 {
		return NumMatrix{}
	}

	prevSum := make([][]int, m+1)
	prevSum[0] = make([]int, n+1)
	for i := 0; i < m; i++ {
		prevSum[i+1] = make([]int, n+1)
		for j := 0; j < n; j++ {
			prevSum[i+1][j+1] = matrix[i][j] + prevSum[i][j+1] + prevSum[i+1][j] - prevSum[i][j]
		}
	}

	return NumMatrix{PrevSum: prevSum}
}

func (matrix *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return matrix.PrevSum[row2+1][col2+1] - matrix.PrevSum[row1][col2+1] - matrix.PrevSum[row2+1][col1] + matrix.PrevSum[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
