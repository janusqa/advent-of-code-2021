package day6

import (
	"adventofcode/aocutils"
	"fmt"
	"strings"
)

type stateTable [9]int

func LanternFishModel(filename string, days int) {
	input := aocutils.GetInputFromFile(filename)
	reading := strings.Split(input[0], ",")
	st := stateTable{0, 0, 0, 0, 0, 0, 0, 0, 0}

	for _, state := range reading {
		st[aocutils.StringToInt(state, 10)] = count(&reading, state)
	}

	fmt.Println((&st).compute(days))

}

func (st *stateTable) compute(days int) int {

	for day := 0; day < days; day++ {

		// fmt.Println(st)

		var zeroState int
		for state := range *st {
			if state > 0 {
				(*st)[state-1] = (*st)[state]
			} else {
				zeroState = (*st)[0]
			}
		}
		(*st)[6] += zeroState
		(*st)[8] = zeroState
	}

	return st.sum()
}

func (st *stateTable) sum() int {

	sum := 0

	for _, fishCount := range *st {
		sum += fishCount
	}
	return sum
}

func count(reading *[]string, element string) int {

	numCount := 0

	for _, state := range *reading {
		if state == element {
			numCount += 1
		}

	}
	return numCount
}
