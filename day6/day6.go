package day6

import (
	"adventofcode/aocutils"
	"fmt"
	"strings"
)

type stateTable [9]int

func LanternFishModel(filename string, days int) {
	input := aocutils.GetInputFromFile(filename)
	readingStr := strings.Split(input[0], ",")
	readingInt := []int{}
	st := stateTable{0, 0, 0, 0, 0, 0, 0, 0, 0}

	for _, element := range readingStr {
		readingInt = append(readingInt, aocutils.StringToInt(element, 10))
	}

	for _, state := range readingInt {
		st[state] = count(&readingInt, state)
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

func count(readingInt *[]int, element int) int {

	numCount := 0

	for _, state := range *readingInt {
		if state == element {
			numCount += 1
		}

	}
	return numCount
}

func (st *stateTable) sum() int {

	sum := 0

	for _, fishCount := range *st {
		sum += fishCount
	}
	return sum
}
