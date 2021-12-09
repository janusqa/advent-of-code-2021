package day3

import (
	"adventofcode/aocutils"
	"fmt"
	"strconv"
)

func GetPowerConsumption(filename string) int {
	input := aocutils.GetInputFromFile(filename)
	width := aocutils.StrLen(input[0])
	gamma, epsilon := "", ""

	for i := 0; i < width; i++ {
		data := ""
		for j := 0; j < len(input); j++ {
			data = fmt.Sprintf("%s%s", data, string([]rune(input[j])[i]))
		}

		ham := hammingWeight(data)
		if ham == 1 {
			gamma = fmt.Sprintf("%s%d", gamma, 1)
			epsilon = fmt.Sprintf("%s%d", epsilon, 0)
		} else if ham == 0 {
			gamma = fmt.Sprintf("%s%d", gamma, 0)
			epsilon = fmt.Sprintf("%s%d", epsilon, 1)
		}

	}

	gammaInt := aocutils.StringToInt(gamma, 2)
	epsilonInt := aocutils.StringToInt(epsilon, 2)

	return gammaInt * epsilonInt

}

func GetLifeSupportRating(filename string) int {
	return getOxygenGeneratorRating(filename) * getCO2ScrubberRating(filename)
}

func getOxygenGeneratorRating(filename string) int {
	input := aocutils.GetInputFromFile(filename)

	k := 0
	for len(input) > 1 {
		data := ""
		for j := 0; j < len(input); j++ {
			data = fmt.Sprintf("%s%s", data, string([]rune(input[j])[k]))
		}

		ham := hammingWeight(data)
		if ham == -1 {
			ham = 1
		}

		temp := []string{}
		for i := range input {
			if string([]rune(input[i])[k]) == strconv.Itoa(ham) {
				temp = append(temp, input[i])
			}
		}
		input = temp
		k++
	}

	return aocutils.StringToInt(input[0], 2)
}

func getCO2ScrubberRating(filename string) int {
	input := aocutils.GetInputFromFile(filename)

	k := 0
	for len(input) > 1 {
		data := ""
		for j := 0; j < len(input); j++ {
			data = fmt.Sprintf("%s%s", data, string([]rune(input[j])[k]))
		}

		ham := hammingWeight(data)
		if ham == -1 {
			ham = 0
		} else if ham == 1 {
			ham = 0
		} else {
			ham = 1
		}

		temp := []string{}
		for i := range input {
			if string([]rune(input[i])[k]) == strconv.Itoa(ham) {
				temp = append(temp, input[i])
			}
		}
		input = temp
		k++
	}

	return aocutils.StringToInt(input[0], 2)
}

func hammingWeight(data string) int {

	one := 0
	zero := 0

	for i := 0; i < aocutils.StrLen(data); i++ {
		if (string([]rune(data)[i])) == "1" {
			one++
		} else {
			zero++
		}
	}

	if one > zero {
		return 1
	} else if zero > one {
		return 0
	} else {
		return -1
	}
}
