package day18

import (
	"adventofcode/aocutils"
	"crypto/sha1"
	"fmt"
	"regexp"
)

// This puzzle can be easily solved, or it can be a pain.  I woke up this morning and chose pain.
// Easy way is probably to parse the stings into threes.
// The solution below works directly with the string so it is very slow.

func PartI(filename string) {
	input := aocutils.GetInputFromFile(filename)

	re_numbers, _ := regexp.Compile(`\d{1,}`)
	re_pair, _ := regexp.Compile(`\[\d{1,},\d{1,}\]`)

	final_snail_sum := ""
	snail_sum_two_largest := 0

	// Calculate final snail sum
	for _, snail_number := range input {
		if final_snail_sum == "" {
			final_snail_sum = snail_number
			continue
		} else {
			final_snail_sum = reduce(add(final_snail_sum, snail_number), re_numbers)
		}
	}

	// Calculate the magnitude of the final snail sum
	magnitude_final_sum := magnitude(final_snail_sum, re_pair, re_numbers)

	// Calculate the largest magnitude possible between any two pairs of Snail Fish Numbers
	for _, snail_number_1 := range input {
		for _, snail_number_2 := range input {
			sum_two := magnitude(reduce(add(snail_number_1, snail_number_2), re_numbers), re_pair, re_numbers)
			sum_two_int := aocutils.StringToInt(sum_two, 10)
			if sum_two_int > snail_sum_two_largest {
				snail_sum_two_largest = sum_two_int
			}
		}
	}

	// Output
	fmt.Println("Final Snail Sum: ", final_snail_sum)
	fmt.Println("Magnitude: ", magnitude_final_sum)
	fmt.Println("Largest Magnitude of any two: ", snail_sum_two_largest)
}

func add(s1 string, s2 string) string {
	return fmt.Sprintf("[%s,%s]", s1, s2)
}

func reduce(s string, re_numbers *regexp.Regexp) string {

	reduced := false
	check_split := 0
	check_explode := true
	for !reduced {
		parenthesis_count := -1
		regular_number_count := 0
		position := 0
		final_snail_sum_sha1_start := fmt.Sprintf("%x", sha1.Sum([]byte(s)))
		for parenthesis_count != 0 {
			numbers := re_numbers.FindAllStringIndex(s, -1)
			if string(s[position]) == "[" {
				if parenthesis_count == -1 {
					parenthesis_count = 0
				}
				parenthesis_count++
			} else if string(s[position]) == "]" {
				parenthesis_count--
			} else if string(s[position]) != "," {
				regular_number := aocutils.StringToInt(s[numbers[regular_number_count][0]:numbers[regular_number_count][1]], 10)
				if parenthesis_count > 4 {
					// Explode
					s = explode(s, regular_number, regular_number_count, &check_split, numbers)
					break
				} else if regular_number > 9 && !check_explode {
					// Split
					s = split(s, regular_number, regular_number_count, &check_split, &check_explode, numbers)
					break
				}
				position = numbers[regular_number_count][1] - 1
				regular_number_count++
			}
			position++
		}

		// file, err := os.OpenFile("./day18/day18.out", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		// if err == nil {
		// 	// file.WriteString(fmt.Sprintf("%s: Splits: %d\n", final_snail_sum, check_split))
		// 	file.WriteString(fmt.Sprintf("%s\n", final_snail_sum))
		// }
		// defer file.Close()

		final_snail_sum_sha1_stop := fmt.Sprintf("%x", sha1.Sum([]byte(s)))
		if final_snail_sum_sha1_start == final_snail_sum_sha1_stop {
			if check_explode {
				check_explode = false
			}
			if !check_explode && check_split < 1 {
				reduced = true
			}
		}
	}

	return s
}

func explode(s string, regular_number, regular_number_count int, check_split *int, numbers [][]int) string {

	left := -1
	right := -1
	if regular_number_count > 0 {
		left_original := aocutils.StringToInt(s[numbers[regular_number_count-1][0]:numbers[regular_number_count-1][1]], 10)
		left = left_original + regular_number
		if left_original < 10 && left > 9 {
			(*check_split)++
		}
	}
	if regular_number_count+2 < len(numbers) {
		right_original := aocutils.StringToInt(s[numbers[regular_number_count+2][0]:numbers[regular_number_count+2][1]], 10)
		right = right_original + aocutils.StringToInt(s[numbers[regular_number_count+1][0]:numbers[regular_number_count+1][1]], 10)
		if right_original < 10 && right > 9 {
			(*check_split)++
		}
	}

	if regular_number > 9 {
		(*check_split)--
	}
	if aocutils.StringToInt(s[numbers[regular_number_count+1][0]:numbers[regular_number_count+1][1]], 10) > 9 {
		(*check_split)--
	}

	if right != -1 {
		s = s[:numbers[regular_number_count+2][0]] + fmt.Sprintf("%d", right) + s[numbers[regular_number_count+2][1]:]
	}
	s = s[:numbers[regular_number_count][0]-1] + fmt.Sprintf("%d", 0) + s[numbers[regular_number_count+1][1]+1:]
	if left != -1 {
		s = s[:numbers[regular_number_count-1][0]] + fmt.Sprintf("%d", left) + s[numbers[regular_number_count-1][1]:]
	}

	return s
}

func split(s string, regular_number, regular_number_count int, check_split *int, check_explode *bool, numbers [][]int) string {
	left := regular_number / 2
	right := left
	if regular_number%2 > 0 {
		right += 1
	}
	if left > 9 {
		(*check_split)++
	}
	if right > 9 {
		(*check_split)++
	}
	split := fmt.Sprintf("[%d,%d]", left, right)
	s = s[:numbers[regular_number_count][0]] + split + s[numbers[regular_number_count][1]:]
	(*check_split)--
	(*check_explode) = true

	return s
}

func magnitude(s string, re_pair *regexp.Regexp, re_numbers *regexp.Regexp) string {

	pair := (*re_pair).FindStringIndex(s)
	for pair != nil {
		pair_elements := (*re_numbers).FindAllStringIndex(s[pair[0]:pair[1]], -1)
		left := aocutils.StringToInt(s[pair[0]:pair[1]][pair_elements[0][0]:pair_elements[0][1]], 10)
		right := aocutils.StringToInt(s[pair[0]:pair[1]][pair_elements[1][0]:pair_elements[1][1]], 10)
		collapse := 3*left + 2*right
		s = s[:pair[0]] + fmt.Sprintf("%d", collapse) + s[pair[1]:]
		pair = (*re_pair).FindStringIndex(s)
	}
	return s
}
