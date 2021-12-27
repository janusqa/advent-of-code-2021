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

	for _, element := range readingStr {
		readingInt = append(readingInt, aocutils.StringToInt(element, 10))
	}

	st := stateTable{0, 0, 0, 0, 0, 0, 0, 0, 0}

	for _, fish := range readingInt {
		if st[fish] == 0 {
			st[fish] = count(&readingInt, fish)
		}
	}

	fmt.Println((&st).compute(days))

}

func (st *stateTable) compute(days int) int {

	for day := 0; day < days; day++ {

		// fmt.Println(st)

		var zeroState int
		for i := range *st {
			if i > 0 {
				(*st)[i-1] = (*st)[i]
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

	for _, item := range *readingInt {
		if item == element {
			numCount += 1
		}

	}
	return numCount
}

func (st *stateTable) sum() int {

	sum := 0

	for _, fish := range *st {
		sum += fish
	}
	return sum
}
