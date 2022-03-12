package day15

import (
	"adventofcode/aocutils"
	"fmt"
	"strings"
)

var input []string

func PartI(filename string) {
	input = aocutils.GetInputFromFile(filename)
	fmt.Println(input)

	cave := [][]int{}

	for row_index, row := range input {
		row_items := strings.Split(row, "")
		cave = append(cave, []int{})
		for _, column := range row_items {
			cave[row_index] = append(cave[row_index], aocutils.StringToInt(column, 10))
		}
	}

}
