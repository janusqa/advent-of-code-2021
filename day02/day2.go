package day2

import (
	"adventofcode/aocutils"
	"fmt"
	"strings"
)

func Navigate(filename string) int {
	input := aocutils.GetInputFromFile(filename)

	horizontal, depth, aim := 0, 0, 0
	var command []string

	for _, c := range input {
		command = strings.Fields(c)
		value := aocutils.StringToInt(command[1], 0)
		switch command[0] {
		case "forward":
			horizontal += value
			depth += aim * value
		case "up":
			aim -= value
		case "down":
			aim += value
		}
	}

	fmt.Printf("Horizontal: %d, Depth: %d\n", horizontal, depth)

	return horizontal * depth
}
