package day5

import (
	"adventofcode/aocutils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type screen [][]int
type line_segment []point
type point struct {
	x int
	y int
}

func getHydrothermalReadings(filename string) []line_segment {

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

	reading := []line_segment{}
	for scanner.Scan() {
		reading = append(reading, ls(scanner.Text()))
	}
	return reading
}

// func plot(ls line_segment) {
// 	for _, point := range ls {

// 		fmt.Println(len(display[point.y]))

// 	}
// }

func (s *screen) on(reading []line_segment) {
	max_rows := 0
	max_cols := 0

	for _, ls := range reading {
		for _, point := range ls {
			if point.x > max_cols {
				max_cols = point.x
			}
			if point.y > max_rows {
				max_rows = point.y
			}
		}
	}

	for i := 0; i <= max_rows; i++ {
		*s = append(*s, []int{})
		for j := 0; j <= max_cols; j++ {
			(*s)[i] = append((*s)[i], 0)
		}
	}
}

func (s *screen) plot(reading []line_segment) {
	for _, ls := range reading {
		//if (ls[0].x == ls[len(ls)-1].x) || (ls[0].y == ls[len(ls)-1].y) {
		for _, point := range ls {
			(*s)[point.y][point.x] += 1
		}
		//}
	}
}

func (s *screen) warn() int {
	alert := 0

	for _, row := range *s {
		for _, point := range row {
			if point > 1 {
				alert += 1
			}
		}
	}
	return alert
}

func (s *screen) print() {
	for _, row := range *s {
		fmt.Println(row)
	}
	fmt.Println()
}

func ls(reading string) line_segment {

	line_segment := []point{}
	ls := strings.Split(reading, " -> ")
	ls0 := strings.Split(ls[0], ",")
	ls1 := strings.Split(ls[1], ",")
	p0 := point{x: aocutils.StringToInt(ls0[0], 10), y: aocutils.StringToInt(ls0[1], 10)}
	pn := point{x: aocutils.StringToInt(ls1[0], 10), y: aocutils.StringToInt(ls1[1], 10)}

	var x_min int
	var x_max int
	var y_min int
	var y_max int
	var xy_max int

	if p0.x < pn.x {
		x_min = p0.x
		x_max = pn.x
	} else if p0.x > pn.x {
		x_min = pn.x
		x_max = p0.x
	} else {
		x_min = p0.x
		x_max = p0.x
	}

	if p0.y < pn.y {
		y_min = p0.y
		y_max = pn.y
	} else if p0.y > pn.y {
		y_min = pn.y
		y_max = p0.y
	} else {
		y_min = p0.y
		y_max = p0.y
	}

	x_diff := x_max - x_min
	y_diff := y_max - y_min

	if x_diff < y_diff {
		xy_max = y_diff
	} else if x_diff > y_diff {
		xy_max = x_diff
	} else {
		xy_max = x_diff
	}

	line_segment = append(line_segment, p0)

	x := p0.x
	y := p0.y
	for i := 0; i < (xy_max - 1); i++ {
		if p0.x > pn.x {
			x -= 1
		} else if p0.x < pn.x {
			x += 1
		}

		if p0.y > pn.y {
			y -= 1
		} else if p0.y < pn.y {
			y += 1
		}

		line_segment = append(line_segment, point{x: x, y: y})
	}

	line_segment = append(line_segment, pn)

	return line_segment
}

func HydrothermalVentDetection(filename string) {
	display := screen{}
	reading := getHydrothermalReadings(filename)
	(&display).on(reading)
	(&display).plot(reading)
	//(&display).print()
	fmt.Println((&display).warn())
}
