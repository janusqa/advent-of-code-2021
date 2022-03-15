package day16

import (
	"adventofcode/aocutils"
	"fmt"
)

var input []string

func PartI(filename string) {
	input = aocutils.GetInputFromFile(filename)
	fmt.Println(input)

	message_hex := "38006F45291200"
	message_binary := hex_to_binary(message_hex)
	fmt.Println(message_binary)
}

func hex_to_binary(message_hex string) string {
	hex_to_binary_map := map[string]string{
		"0": "0000",
		"1": "0001",
		"2": "0010",
		"3": "0011",
		"4": "0100",
		"5": "0101",
		"6": "0110",
		"7": "0111",
		"8": "1000",
		"9": "1001",
		"A": "1010",
		"B": "1011",
		"C": "1100",
		"D": "1101",
		"E": "1110",
		"F": "1111",
	}

	message_binary := ""

	for _, char := range message_hex {
		message_binary = fmt.Sprintf("%s%s", message_binary, hex_to_binary_map[string(char)])
	}

	return message_binary
}
