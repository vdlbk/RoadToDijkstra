package arrays

func isValidSudoku(board [][]byte) bool {
	sudoku := make(map[byte][]int, len(board))
	squares := make([][]byte, len(board))

	for i := 0; i < len(squares); i++ {
		squares[i] = make([]byte, 0, 9)
	}

	currentSquare := 0
	z := 0
	for y, row := range board {
		for x, cell := range row {
			if z != 0 && z%3 == 0 {
				currentSquare = ((currentSquare + 1) % 3) + ((z / 27) * 3)
			}
			if cell != 46 {
				sudoku[cell] = append(sudoku[cell], x, y)
				squares[currentSquare] = append(squares[currentSquare], cell)
			}

			z++
		}
	}
	for _, positions := range sudoku {
		// check row
		rows := make(map[int]bool, 0)
		for i := 0; i < len(positions); i += 2 {
			num := positions[i]
			if _, ok := rows[num]; ok {
				return false
			}

			rows[num] = true
		}
		// check columns
		columns := make(map[int]bool, 0)
		for i := 1; i < len(positions); i += 2 {
			num := positions[i]
			if _, ok := columns[num]; ok {
				return false
			}

			columns[num] = true
		}
	}
	for _, square := range squares {
		rows := make(map[byte]bool, 0)
		for i := 0; i < len(square); i++ {
			num := square[i]
			if _, ok := rows[num]; ok {
				return false
			}

			rows[num] = true
		}
	}
	return true
}
