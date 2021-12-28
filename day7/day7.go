package day7

import (
	"adventofcode/aocutils"
	"fmt"
	"strings"
)

func CrabController(filename string) {
	input := strings.Split(aocutils.GetInputFromFile(filename)[0], ",")

	// Map each position and how many crabs are in that position
	var reading = make(map[int]int)
	var maximum_horizontal_position = 0
	for _, crab := range input {
		position := aocutils.StringToInt(crab, 10)
		reading[position] += 1
		if position > maximum_horizontal_position {
			maximum_horizontal_position = position
		}
	}

	for i := 0; i <= maximum_horizontal_position; i++ {
		if _, ok := reading[i]; !ok {
			reading[i] = 0
		}
	}

	// maximum_at_position := 0
	// for position := range reading {
	// 	if reading[position] > maximum_at_position {
	// 		maximum_at_position = reading[position]
	// 	}
	// }

	fuel_consumption_min := 0
	for position := range reading {
		//if reading[position] == maximum_at_position {
		fuel_consumption := compute_fuel_consumption(&reading, position)
		if (fuel_consumption_min == 0) || (fuel_consumption < fuel_consumption_min) {
			fuel_consumption_min = fuel_consumption
		}
		//}
	}

	// fmt.Println(reading)
	fmt.Println(fuel_consumption_min)
}

func compute_fuel_consumption(reading *map[int]int, position int) int {

	fuel_consumption := 0

	for pos := range *reading {
		if (pos != position) && ((*reading)[pos] > 0) {
			distance := pos - position
			if distance < 0 {
				distance /= -1
			}
			// ***original calculation
			// fuel_consumption += distance * (*reading)[pos]

			// ***New calculation adjusted for new relevation of crab tech.
			// *** We can use a for loop here to calculate the sum of sequence 1+2+3+...+n
			// for i := 1; i <= distance; i++ {
			// 	distance_augumented += i
			// }
			// *** or we can use Gausses formula
			distance_augumented := ((distance + 1) * distance) / 2

			fuel_consumption += distance_augumented * (*reading)[pos]
			// fmt.Printf("Move from %v to %v: %v fuel [%v]\n", pos, position, fuel_consumption, (*reading)[pos])
		}
	}

	return fuel_consumption
}
