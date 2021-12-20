package day4

import (
	"adventofcode/aocutils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type bingoGame struct {
	draws  []int
	boards []bingoBoard
}

type bingoBoard [][]int

func getBingoGame(filename string, bingo *bingoGame) {

	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("Failed to get input from file due to %s", err)
		os.Exit(1)
	}

	defer file.Close()

	// The bufio.NewScanner() function is called in which the
	// object os.File passed as its parameter and this returns a
	// object bufio.Scanner which is further used on the
	// bufio.Scanner.Split() method.
	scanner := bufio.NewScanner(file)

	// The bufio.ScanLines is used as an
	// input to the method bufio.Scanner.Split()
	// and then the scanning forwards to each
	// new line using the bufio.Scanner.Scan()
	// method.
	scanner.Split(bufio.ScanLines)

	lineNumber := 0
	boardNumber := 0
	rowNum := 0
	line := ""

	for scanner.Scan() {
		line = scanner.Text()
		if lineNumber == 0 && aocutils.StrLen(strings.TrimSpace(line)) > 0 {
			inputS := strings.Split(line, ",")
			for _, chr := range inputS {
				(*bingo).draws = append((*bingo).draws, aocutils.StringToInt(chr, 10))
			}
		} else if aocutils.StrLen(strings.TrimSpace(line)) > 0 {
			inputS := strings.Fields(line)
			row := []int{}
			for _, chr := range inputS {
				row = append(row, aocutils.StringToInt(chr, 10))
			}
			(*bingo).boards[boardNumber] = append((*bingo).boards[boardNumber], row)
			rowNum++
		} else {
			(*bingo).boards = append((*bingo).boards, bingoBoard{})
			rowNum = 0
			if lineNumber > 1 {
				boardNumber++
			}
		}
		lineNumber++
	}
}

func BingoSubsystem(filename string) {
	bingo := bingoGame{}
	wins := []int{}

	getBingoGame(filename, &bingo)
	for _, draw := range bingo.draws {
		for i, board := range bingo.boards {
			if !board.skip(i, wins) {
				board.mark(draw)
				if board.wins() {
					// fmt.Println(board.score(draw))
					wins = append(wins, i)
					if len(wins) == len(bingo.boards) {
						fmt.Println(board.score(draw))
						return
					}
				}
			}
		}
	}
}

func (b bingoBoard) mark(draw int) {
	for _, row := range b {
		for j, col := range row {
			if col == draw {
				row[j] = -1
			}
		}
	}
}

func (b bingoBoard) wins() bool {

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

func (b bingoBoard) score(draw int) int {
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

func (bingoBoard) skip(board int, winningBoards []int) bool {
	for _, winningBoard := range winningBoards {
		if board == winningBoard {
			return true
		}
	}
	return false
}
