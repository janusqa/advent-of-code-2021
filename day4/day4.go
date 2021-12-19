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

func Zeus(filename string) {
	bingo := bingoGame{}
	getBingoGame(filename, &bingo)
	fmt.Println(bingo)
}
