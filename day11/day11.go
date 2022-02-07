package day11

import (
	"adventofcode/aocutils"
	"fmt"
)

const ROWCOUNT = 10
const COLUMNCOUNT = ROWCOUNT

var input []string
var octupussy [ROWCOUNT * COLUMNCOUNT]int

func PartI(filename string) {
	const STEPS = 200
	flashes := 0
	queue := []int{}
	in_queue := false

	input = aocutils.GetInputFromFile(filename)
	for i, row := range input {
		for j, octupus := range row {
			octupussy[i*ROWCOUNT+j] = aocutils.StringToInt(string(octupus), 10)
		}
	}

	// PART I for loop
	//for i := 0; i < STEPS; i++ {

	// PART II for loop
	i := 0
	for power_output() != 0 {
		// increment energy level of each octupus
		// add each one that exceeds lvl 9 to queue
		// for flash processing

		for j := 0; j < len(octupussy); j++ {
			process_octupus(j, &queue, in_queue, &flashes)
		}

		// Process octupii that surpassed lvl 9
		for len(queue) > 0 {
			in_queue = true
			octupus := queue[0]
			queue = queue[1:]

			// check left
			if within_bounds(octupus-1, octupus) {
				process_octupus(octupus-1, &queue, in_queue, &flashes)
			}

			// check right
			if within_bounds(octupus+1, octupus) {
				process_octupus(octupus+1, &queue, in_queue, &flashes)
			}

			// check above
			if within_bounds(octupus-COLUMNCOUNT, octupus-COLUMNCOUNT) {
				process_octupus(octupus-COLUMNCOUNT, &queue, in_queue, &flashes)

				// check left
				if within_bounds(octupus-COLUMNCOUNT-1, octupus-COLUMNCOUNT) {
					process_octupus(octupus-COLUMNCOUNT-1, &queue, in_queue, &flashes)
				}

				// check right
				if within_bounds(octupus-COLUMNCOUNT+1, octupus-COLUMNCOUNT) {
					process_octupus(octupus-COLUMNCOUNT+1, &queue, in_queue, &flashes)
				}
			}

			// check below
			if within_bounds(octupus+COLUMNCOUNT, octupus+COLUMNCOUNT) {
				process_octupus(octupus+COLUMNCOUNT, &queue, in_queue, &flashes)

				// check left
				if within_bounds(octupus+COLUMNCOUNT-1, octupus+COLUMNCOUNT) {
					process_octupus(octupus+COLUMNCOUNT-1, &queue, in_queue, &flashes)
				}

				// check right
				if within_bounds(octupus+COLUMNCOUNT+1, octupus+COLUMNCOUNT) {
					process_octupus(octupus+COLUMNCOUNT+1, &queue, in_queue, &flashes)
				}
			}
			in_queue = false
		}

		if power_output() == 0 {
			oprint()
			fmt.Printf("Will synchronize at STEP %v\n\n", i+1)
		}
		i++
	}
	// Part I
	// fmt.Printf("Flashes: %v\n", flashes)
}

func process_octupus(index int, queue *[]int, in_queue bool, flashes *int) {
	if !in_queue || (in_queue && octupussy[index] > 0) {
		octupussy[index]++
	}
	if octupussy[index] > 9 {
		*queue = append(*queue, index)
		octupussy[index] = 0
		(*flashes)++
	}

}

func within_bounds(index int, range_index int) bool {
	lower := int(range_index/ROWCOUNT) * COLUMNCOUNT
	upper := lower + COLUMNCOUNT
	return index >= lower && index < upper && lower >= 0 && upper <= ROWCOUNT*COLUMNCOUNT
}

func oprint() {
	fmt.Println()
	fmt.Println()
	for i := 0; i < len(octupussy); i++ {
		if i%COLUMNCOUNT == 0 && i > 0 {
			fmt.Println()
		}
		fmt.Printf("%v", octupussy[i])

	}
	fmt.Println()
}

func power_output() int {

	result := 0
	for _, po := range octupussy {
		result += po
	}
	return result

}
