package day8

import (
	"adventofcode/aocutils"
	"fmt"
	"sort"
	"strings"
)

var input []string

func PartI(filename string) {
	input = aocutils.GetInputFromFile(filename)
	unique_digit_count := 0

	for _, entry := range input {
		entry_components := strings.Split(entry, " | ")
		outputs := strings.Split(entry_components[1], " ")

		for _, digit := range outputs {
			if (aocutils.StrLen(digit) == 2) || (aocutils.StrLen(digit) == 3) || (aocutils.StrLen(digit) == 4) || (aocutils.StrLen(digit) == 7) {
				unique_digit_count += 1
			}
		}
	}
	fmt.Printf("Unique: %v\n", unique_digit_count)
}

func PartII() {

	throughput := 0

	for _, entry := range input {
		var signals_decoded_inverse = make(map[int]string)
		var signals_decoded = make(map[string]int)

		entry_components := strings.Split(entry, " | ")
		inputs := strings.Split(entry_components[0], " ")
		outputs := strings.Split(entry_components[1], " ")

		for len(inputs) > 0 {
			input_temp := strings.Split(inputs[0], "")
			sort.Strings(input_temp)
			inputs[0] = strings.Join(input_temp, "")
			if (len(inputs[0]) == 2) || (len(inputs[0]) == 4) || (len(inputs[0]) == 3) || (len(inputs[0]) == 7) {
				if len(inputs[0]) == 2 {
					signals_decoded[inputs[0]] = 1
					signals_decoded_inverse[1] = inputs[0]
				} else if len(inputs[0]) == 4 {
					signals_decoded[inputs[0]] = 4
					signals_decoded_inverse[4] = inputs[0]
				} else if len(inputs[0]) == 3 {
					signals_decoded[inputs[0]] = 7
					signals_decoded_inverse[7] = inputs[0]
				} else if len(inputs[0]) == 7 {
					signals_decoded[inputs[0]] = 8
					signals_decoded_inverse[8] = inputs[0]
				}
				inputs = inputs[1:]
			} else if len(inputs[0]) == 5 {
				processed := false
				value_1, ok_1 := signals_decoded_inverse[1]
				value_6, ok_6 := signals_decoded_inverse[6]
				if ok_1 && !processed {
					if contains(inputs[0], value_1) {
						signals_decoded[inputs[0]] = 3
						signals_decoded_inverse[3] = inputs[0]
						inputs = inputs[1:]
						processed = true
					}
				}
				if ok_6 && !processed {
					if contains(value_6, inputs[0]) {
						signals_decoded[inputs[0]] = 5
						signals_decoded_inverse[5] = inputs[0]
						inputs = inputs[1:]
						processed = true
					}
				}
				if (len(inputs) <= 2) && !processed {
					signals_decoded[inputs[0]] = 2
					signals_decoded_inverse[2] = inputs[0]
					inputs = inputs[1:]
					processed = true
				}
				if !processed {
					temp := inputs[0]
					inputs = inputs[1:]
					inputs = append(inputs, temp)
				}
			} else if len(inputs[0]) == 6 {
				processed := false
				value_1, ok_1 := signals_decoded_inverse[1]
				value_4, ok_4 := signals_decoded_inverse[4]
				value_7, ok_7 := signals_decoded_inverse[7]
				if ok_4 && ok_7 && !processed {
					if contains(inputs[0], value_4) && contains(inputs[0], value_7) {
						signals_decoded[inputs[0]] = 9
						signals_decoded_inverse[9] = inputs[0]
						inputs = inputs[1:]
						processed = true
					}
				}
				if ok_1 && !processed {
					if !contains(inputs[0], value_1) {
						signals_decoded[inputs[0]] = 6
						signals_decoded_inverse[6] = inputs[0]
						inputs = inputs[1:]
						processed = true
					}
				}
				if (len(inputs) <= 2) && !processed {
					signals_decoded[inputs[0]] = 0
					signals_decoded_inverse[0] = inputs[0]
					inputs = inputs[1:]
					processed = true
				}
				if !processed {
					temp := inputs[0]
					inputs = inputs[1:]
					inputs = append(inputs, temp)
				}
			}
		}
		p := 1000
		temp := 0
		for _, digit := range outputs {
			digit_temp := strings.Split(digit, "")
			sort.Strings(digit_temp)
			digit = strings.Join(digit_temp, "")
			temp += signals_decoded[digit] * p
			p /= 10
		}
		throughput += temp
	}
	fmt.Printf("Throughput: %v\n", throughput)
}

func contains(a string, b string) bool {

	result := true

	for _, char := range strings.Split(b, "") {
		if !strings.Contains(a, char) {
			result = false
		}
	}

	return result

}
