package day14

import (
	"adventofcode/aocutils"
	"fmt"
	"strings"
)

var input []string

func PartI(filename string) {
	input = aocutils.GetInputFromFile(filename)

	polymer_template := input[0]
	insertion_elements := make(map[string]string)
	adjacent_pairs := make(map[string]int)
	element_count := make(map[string]int)
	steps := 40

	for _, insertion_pair := range input[2:] {
		insertion_pair_map := strings.Split(insertion_pair, " -> ")
		insertion_elements[insertion_pair_map[0]] = insertion_pair_map[1]
	}

	for i := range polymer_template {
		element_count[string(polymer_template[i])] += 1
		if i+1 < len(polymer_template) {
			key := fmt.Sprintf("%s%s", string(polymer_template[i]), string(polymer_template[i+1]))
			adjacent_pairs[key] += 1
		}
	}

	for i := 0; i < steps; i++ {
		pass := make(map[string]int)
		for key, _ := range adjacent_pairs {
			element_count[insertion_elements[key]] += adjacent_pairs[key]
			pair_array := strings.Split(key, "")
			key_left := fmt.Sprintf("%s%s", pair_array[0], insertion_elements[key])
			key_right := fmt.Sprintf("%s%s", insertion_elements[key], pair_array[1])
			pass[key_left] += adjacent_pairs[key]
			pass[key_right] += adjacent_pairs[key]
			delete(adjacent_pairs, key)
		}
		for key, _ := range pass {
			adjacent_pairs[key] += pass[key]
			delete(pass, key)
		}
	}

	min := 0
	max := 0
	for _, value := range element_count {
		if min == 0 {
			min = value
		} else if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	fmt.Println(polymer_template)
	fmt.Println(len(polymer_template))
	fmt.Println(element_count)
	fmt.Printf("Max: %v - Min: %v = %v\n", max, min, max-min)

	// fmt.Println(adjacent_pairs)
	// fmt.Println(element_count)
	// fmt.Println()

	// **** Old implementation. Too slow ^^^
	// for i := 0; i < steps; i++ {
	// 	new_polymer_template := []string{}
	// 	for _, char := range polymer_template {
	// 		if len(new_polymer_template) > 0 {
	// 			insertion_pairs_key := fmt.Sprintf("%s%s", new_polymer_template[len(new_polymer_template)-1], string(char))
	// 			element, ok := insertion_pairs[insertion_pairs_key]
	// 			if ok {
	// 				new_polymer_template = append(new_polymer_template, element)
	// 				if i == steps-1 {
	// 					element_count[element] += 1
	// 				}
	// 			}
	// 		}
	// 		new_polymer_template = append(new_polymer_template, string(char))
	// 		if i == steps-1 {
	// 			element_count[string(char)] += 1
	// 		}
	// 	}
	// 	polymer_template = strings.Join(new_polymer_template, "")
	// }

}
