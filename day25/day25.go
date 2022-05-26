package day25

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func PartI(filename string) {
	east := make(map[[2]int]bool)
	south := make(map[[2]int]bool)
	maxr := 0
	maxc := 0
	getInput(filename, east, south, &maxr, &maxc)
	// display(east, south, maxr, maxc)
	steps := 0
	for step(east, south, maxr, maxc) {
		steps++
		//display(east, south, maxr, maxc)
	}
	display(east, south, maxr, maxc)
	fmt.Println("steps: ", steps+1)
}

func step(east map[[2]int]bool, south map[[2]int]bool, maxr int, maxc int) bool {

	moved := false

	for loc := range east {
		nextPos := next(loc, ">", east, south, maxr, maxc)
		if nextPos != loc {
			east[loc] = true
			moved = true
		}
	}
	for loc := range east {
		if east[loc] {
			nextPos := next(loc, ">", east, south, maxr, maxc)
			east[nextPos] = false
			delete(east, loc)
		}
	}
	for loc := range south {
		nextPos := next(loc, "v", east, south, maxr, maxc)
		if nextPos != loc {
			south[loc] = true
			moved = true
		}
	}
	for loc := range south {
		if south[loc] {
			nextPos := next(loc, "v", east, south, maxr, maxc)
			south[nextPos] = false
			delete(south, loc)
		}
	}

	return moved
}

func next(loc [2]int, herd string, east map[[2]int]bool, south map[[2]int]bool, maxr int, maxc int) [2]int {
	row := loc[0]
	col := loc[1]

	if herd == ">" {
		col = (loc[1] + 1) % maxc
	} else if herd == "v" {
		row = (loc[0] + 1) % maxr
	}

	if _, ok := east[[2]int{row, col}]; !ok {
		if _, ok := south[[2]int{row, col}]; !ok {
			return [2]int{row, col}
		}
	}

	return loc
}

func display(east map[[2]int]bool, south map[[2]int]bool, maxr int, maxc int) {
	seafloor := make([][]string, maxr)
	for i := range seafloor {
		seafloor[i] = make([]string, maxc)
	}
	for loc := range east {
		seafloor[loc[0]][loc[1]] = ">"
	}
	for loc := range south {
		seafloor[loc[0]][loc[1]] = "v"
	}
	for _, row := range seafloor {
		for _, col := range row {
			if len(col) == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(col)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func getInput(filename string, east map[[2]int]bool, south map[[2]int]bool, maxr *int, maxc *int) {

	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("Failed to get input from file due to %s", err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	i := 0
	for scanner.Scan() {
		for j, cucumber := range strings.Split(scanner.Text(), "") {
			if cucumber == ">" {
				east[[2]int{i, j}] = false
			} else if cucumber == "v" {
				south[[2]int{i, j}] = false
			}
			if i == 0 {
				*maxc++
			}
		}
		i++
	}
	*maxr = i
}
