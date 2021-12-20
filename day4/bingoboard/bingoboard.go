package bingoboard

type BingoBoard [][]int

func (b BingoBoard) Mark(draw int) {
	for _, row := range b {
		for j, col := range row {
			if col == draw {
				row[j] = -1
			}
		}
	}
}

func (b BingoBoard) Wins() bool {

	rowCheck := 0
	colCheck := 0

	// check each row for a win
	for row := 0; row < len(b); row++ {
		for col := 0; col < len(b[row]); col++ {
			if b[row][col] == -1 {
				rowCheck++
			}
		}
		if rowCheck == len(b[row]) {
			return true
		}
		rowCheck = 0
	}

	// check each col for a win
	row := 0
	for col := 0; col < len(b[row]); col++ {
		for row < len(b) {
			if b[row][col] == -1 {
				colCheck++
			}
			row++
		}
		if colCheck == len(b) {
			return true
		}
		row = 0
		colCheck = 0
	}

	return false
}

func (b BingoBoard) Score(draw int) int {
	unmarkedSum := 0

	for row := 0; row < len(b); row++ {
		for col := 0; col < len(b[row]); col++ {
			if b[row][col] != -1 {
				unmarkedSum += b[row][col]
			}
		}
	}
	return unmarkedSum * draw
}

func (BingoBoard) Skip(board int, winningBoards []int) bool {
	for _, winningBoard := range winningBoards {
		if board == winningBoard {
			return true
		}
	}
	return false
}
