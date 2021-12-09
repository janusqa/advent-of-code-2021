package main

import (
	"adventofcode/day3"
	"fmt"
)

func main() {

	// *** Day 1
	// result := day1.Depths("./day1/input_test.txt")
	// fmt.Printf("%s\n", result)

	// *** Day 2
	// result := day2.Navigate("./day2/input_test.txt")
	// fmt.Printf("%d\n", result)

	// *** Day 3
	// result := day3.GetPowerConsumption("./day3/input_test.txt")
	result := day3.GetLifeSupportRating("./day3/input.txt")
	fmt.Printf("%d\n", result)
}
