package day13

import (
	"adventofcode/aocutils"
	"fmt"
	"regexp"
	"strings"
)

var input []string

func PartI(filename string) {
	input = aocutils.GetInputFromFile(filename)

	// initilize some needed variables
	rePoint, _ := regexp.Compile(`^\d`)
	reFold, _ := regexp.Compile(`([xy])=(\d+)$`)

	type point struct {
		x int
		y int
	}
	type fold struct {
		axis   string
		offset int
	}
	folds := []fold{}
	points := []point{}
	maxX := 0
	maxY := 0

	// parse input into appropriate data structures
	for _, line := range input {
		if rePoint.MatchString(line) {
			xy := strings.Split(line, ",")
			x := aocutils.StringToInt(xy[0], 10)
			y := aocutils.StringToInt(xy[1], 10)
			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
			points = append(points, point{x: x, y: y})
		}
		if reFold.MatchString(line) {
			parts := reFold.FindStringSubmatch(line)
			folds = append(folds, fold{axis: parts[1], offset: aocutils.StringToInt(parts[2], 10)})
		}
	}

	// create paper
	paper := make([]string, (maxX+1)*(maxY+1))

	// fmt.Println(points)
	// fmt.Println(folds)
	// fmt.Println(maxX, maxY)

	// mark points on paper
	for _, point := range points {
		index := getIndex(point.x, point.y, maxX)
		paper[index] = "#"
	}

	// Folding
	for i, fold := range folds {
		if fold.axis == "x" {
			foldX(&paper, fold.offset, &maxX, maxY)
		} else {
			foldY(&paper, fold.offset, maxX, &maxY)
		}
		if i == len(folds)-1 {
			printPaper(paper, maxX)
		}
		fmt.Println(countDots(paper))
		//fmt.Println(maxX, maxY)
	}
}

func getIndex(x int, y int, maxX int) int {
	// translats a point (x,y) in 2d lattice to an index
	// in a 1d lattice
	row := (maxX + 1) * y
	index := row + x
	return index
}

func printPaper(paper []string, maxX int) {
	for i, p := range paper {
		if p == "#" {
			fmt.Print(p)
		} else {
			fmt.Print(".")

		}
		if (i+1)%(maxX+1) == 0 {
			fmt.Println()
		}
	}
}

func countDots(paper []string) int {
	count := 0
	for _, point := range paper {
		if point == "#" {
			count++
		}
	}
	return count
}

func foldY(paper *[]string, foldAtY int, maxX int, maxY *int) {
	for y := foldAtY + 1; y <= *maxY; y++ {
		mirrorY := foldAtY - (y - foldAtY)
		for x := 0; x <= maxX; x++ {
			copyFrom := getIndex(x, y, maxX)
			copyTo := getIndex(x, mirrorY, maxX)
			if (*paper)[copyFrom] == "#" {
				(*paper)[copyTo] = (*paper)[copyFrom]
				(*paper)[copyFrom] = ""
			}
		}
	}
	*maxY = foldAtY - 1
	*paper = (*paper)[:(maxX+1)*(foldAtY-1)+maxX+1]
}

func foldX(paper *[]string, foldAtX int, maxX *int, maxY int) {

	indexCopyTo := foldAtX
	for y := 0; y <= maxY; y++ {
		for x := foldAtX + 1; x <= *maxX; x++ {
			copyFrom := getIndex(x, y, *maxX)
			copyTo := getIndex(foldAtX-(x-foldAtX), y, *maxX)
			if (*paper)[copyFrom] == "#" {
				(*paper)[copyTo] = (*paper)[copyFrom]
				(*paper)[copyFrom] = ""
			}
		}
		// prep array for turncation
		if y > 0 {
			for x := 0; x < foldAtX; x++ {
				oldIndex := getIndex(x, y, *maxX)
				(*paper)[indexCopyTo] = (*paper)[oldIndex]
				(*paper)[oldIndex] = ""
				indexCopyTo++
			}
		}
	}

	*maxX = foldAtX - 1
	*paper = (*paper)[:((*maxX)+1)*(maxY)+(*maxX)+1]
}
