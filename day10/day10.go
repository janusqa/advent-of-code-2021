package day10

import (
	"adventofcode/aocutils"
	"fmt"
	"sort"
)

var input []string

func PartI(filename string) {

	input = aocutils.GetInputFromFile(filename)

	symbolsSyntaxError := map[string]int{")": 3, "]": 57, "}": 1197, ">": 25137}

	syntaxErrorScore := 0

	inputCleaned := input[:0]
	for _, data := range input {
		illegalCharacter := isCorrupted(data)
		if len(illegalCharacter) > 0 {
			syntaxErrorScore += symbolsSyntaxError[illegalCharacter]
		} else {
			inputCleaned = append(inputCleaned, data)
		}
	}

	for i := len(inputCleaned); i < len(input); i++ {
		input[i] = ""
	}

	fmt.Println(syntaxErrorScore)

	scores := []int{}
	for _, data := range inputCleaned {
		scores = append(scores, autoCompleteScore(data))
	}
	sort.Slice(
		scores,
		func(i int, j int) bool {
			return scores[i] < scores[j]
		},
	)

	fmt.Println(scores[int(len(scores)/2)])
}

func isCorrupted(s string) string {
	symbols := map[string]string{")": "(", "]": "[", "}": "{", ">": "<"}

	buffer := []string{}

	for _, c := range s {

		if _, ok := symbols[string(c)]; ok {
			if len(buffer) == 0 {
				return string(c)
			} else if buffer[len(buffer)-1] == symbols[string(c)] {
				buffer = buffer[:len(buffer)-1]
			} else {
				return string(c)
			}
		} else {
			buffer = append(buffer, string(c))
		}
	}

	return ""
}

func autoCompleteScore(s string) int {
	symbols := map[string]string{"(": ")", "[": "]", "{": "}", "<": ">"}
	symbolsAutocomplete := map[string]int{"(": 1, "[": 2, "{": 3, "<": 4}
	buffer := []string{}
	score := 0

	for _, c := range s {
		if (len(buffer) == 0) || string(c) != symbols[buffer[len(buffer)-1]] {
			buffer = append(buffer, string(c))
		} else {
			buffer = buffer[:len(buffer)-1]
		}
	}

	for len(buffer) > 0 {
		score = score*5 + symbolsAutocomplete[buffer[len(buffer)-1]]
		buffer = buffer[:len(buffer)-1]
	}

	return score
}
