package arrays

func rotateMatrix(matrix [][]int) {
	if len(matrix) == 0 || len(matrix) == 1 {
		return
	}
	mm := make([][]int, len(matrix))
	for i := range matrix {
		mm[i] = make([]int, len(matrix[i]))
		copy(mm[i], matrix[i])
	}
	for i, row := range matrix {
		y := len(matrix) - 1 - i
		for x := range row {
			matrix[x][y] = mm[i][x]
		}
	}
}
