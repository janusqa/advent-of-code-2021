package day9

import (
	"adventofcode/aocutils"
	"fmt"
	"sort"
)

var input []string

func PartI(filename string) {

	input = aocutils.GetInputFromFile(filename)

	lowpoints := [][2]int{}
	riskLevelTotal := 0

	for i, row := range input {
		row_rune := []rune(row)
		for j := range row_rune {
			var left, right, up, down int32 = 0, 0, 0, 0

			if j-1 >= 0 {
				left = row_rune[j-1]
			}

			if i-1 >= 0 {
				up = []rune(input[i-1])[j]
			}

			if j+1 < len(row_rune) {
				right = row_rune[j+1]
			}

			if i+1 < len(input) {
				down = []rune(input[i+1])[j]
			}

			isLowpoint := true
			if left > 0 {
				isLowpoint = isLowpoint && row_rune[j] < left
			}

			if up > 0 {
				isLowpoint = isLowpoint && row_rune[j] < up
			}

			if right > 0 {
				isLowpoint = isLowpoint && row_rune[j] < right
			}

			if down > 0 {
				isLowpoint = isLowpoint && row_rune[j] < down
			}

			if isLowpoint {
				lowpoints = append(lowpoints, [2]int{i, j})
				riskLevelTotal += aocutils.StringToInt(string(row_rune[j]), 10) + 1
			}

		}
	}
	//fmt.Println(lowpoints)
	fmt.Println(riskLevelTotal)
	PartII(lowpoints)
}

func PartII(lowpoints [][2]int) {

	basinSizes := []int{}
	basinMultiplex := 1

	for _, lowpoint := range lowpoints {
		memo := make(map[string]bool)
		basinSizes = append(basinSizes, mapBasin(input, lowpoint, memo))
	}

	sort.Slice(
		basinSizes,
		func(i int, j int) bool {
			return basinSizes[j] < basinSizes[i]
		},
	)

	for i := 0; i < 3; i++ {
		basinMultiplex *= basinSizes[i]
	}

	fmt.Println(basinMultiplex)
}

func mapBasin(thermalMap []string, location [2]int, memo map[string]bool) int {

	size := 1

	if location[0] < 0 || location[0] > (len(thermalMap)-1) || location[1] < 0 || location[1] > (len(thermalMap[0])-1) {
		return 0
	}

	memo_key := fmt.Sprintf("%v,%v", location[0], location[1])
	if _, ok := memo[memo_key]; ok {
		return 0
	} else {
		memo[memo_key] = true
	}

	if string(thermalMap[location[0]][location[1]]) == string('9') {
		return 0
	}

	// fmt.Println(string(thermalMap[location[0]][location[1]]), location[0], location[1])

	left := location[1] - 1
	up := location[0] - 1
	right := location[1] + 1
	down := location[0] + 1

	// fmt.Println([2]int{location[0], left}, [2]int{up, location[1]}, [2]int{location[0], right}, [2]int{down, location[1]})
	// return 0

	size +=
		mapBasin(thermalMap, [2]int{location[0], left}, memo) +
			mapBasin(thermalMap, [2]int{up, location[1]}, memo) +
			mapBasin(thermalMap, [2]int{location[0], right}, memo) +
			mapBasin(thermalMap, [2]int{down, location[1]}, memo)

	return size
}
