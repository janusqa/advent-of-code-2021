package day15

import (
	"adventofcode/aocutils"
	"fmt"
	"strings"
)

var input []string

func PartI(filename string) {
	input = aocutils.GetInputFromFile(filename)

	cave := [][]int{}

	for row_index, row := range input {
		row_items := strings.Split(row, "")
		cave = append(cave, []int{})
		for _, column := range row_items {
			cave[row_index] = append(cave[row_index], aocutils.StringToInt(column, 10))
		}
	}

	m := len(cave)
	n := len(cave[0])
	minimum_risk := get_minimum_risk(cave, m-1, n-1, map[string]int{})
	fmt.Println(minimum_risk - cave[0][0])

}

func get_minimum_risk(cave [][]int, start_row int, start_column int, memo map[string]int) int {

	key := fmt.Sprintf("%v,%v", start_row, start_column)

	if memoized_risks, ok := memo[key]; ok {
		// fmt.Println("Retrieving stored risk assessment")
		return memoized_risks
	}

	if (start_row == 0) && (start_column == 0) {
		return cave[start_row][start_column]
	}

	if (start_row == -1) || (start_column == -1) {
		return 0
	}

	var minimum_risk int

	risk_1 := get_minimum_risk(cave, start_row-1, start_column, memo)
	risk_2 := get_minimum_risk(cave, start_row, start_column-1, memo)

	if risk_1 == 0 {
		minimum_risk = risk_2
	} else if risk_2 == 0 {
		minimum_risk = risk_1
	} else if risk_1 < risk_2 {
		minimum_risk = risk_1
	} else {
		minimum_risk = risk_2
	}

	minimum_risk += cave[start_row][start_column]

	memo[key] = minimum_risk
	return minimum_risk
}
