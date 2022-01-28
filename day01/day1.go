package day1

import (
	"adventofcode/aocutils"
	"fmt"
)

func Depths(filename string) string {
	input := aocutils.GetInputFromFile(filename)

	count_larger_measurements_1, count_larger_measurements_2 := 0, 0
	turtoise := 0
	hare := turtoise + 2
	three_measurement_window := 0

	for index := 1; index < len(input); index++ {
		if aocutils.StringToInt(input[index], 0) > aocutils.StringToInt(input[index-1], 0) {
			count_larger_measurements_1 += 1
		}
	}

	if hare < len(input) {
		three_measurement_window = aocutils.StringToInt(input[0], 0) + aocutils.StringToInt(input[1], 0) + aocutils.StringToInt(input[2], 0)
	}

	for hare < len(input) {
		temp := three_measurement_window - aocutils.StringToInt(input[turtoise], 0)
		turtoise += 1
		hare += 1
		if hare < len(input) {
			temp += aocutils.StringToInt(input[hare], 0)
			if temp > three_measurement_window {
				count_larger_measurements_2 += 1
			}
			three_measurement_window = temp
		}
	}

	return fmt.Sprintf("%d %d", count_larger_measurements_1, count_larger_measurements_2)
}
