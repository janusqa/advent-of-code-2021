package day16

import (
	"adventofcode/aocutils"
	"fmt"
	"regexp"
	"strings"
)

var input []string

func PartI(filename string) {
	input = aocutils.GetInputFromFile(filename)

	message_hex := "9C0141080250320F1802104A08"
	message_binary := hex_to_binary(message_hex)
	message_parse(message_binary)
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

func message_parse(message_binary string) {

	re_version_type, _ := regexp.Compile(`^(\d{3})(\d{3})`)
	// re_literal, _ := regexp.Compile(`^(?:\d{6})(?:1(\d{4}))*0(\d{4})`) // repeat a capture group
	re_literal, _ := regexp.Compile(`^\d{6}((?:1\d{4})*)(0\d{4})`) // capture a repeating group. Put extra parens around the repeating capture group
	re_literal_parse, _ := regexp.Compile(`(?:\d(\d{4}))`)
	//re_lengthtype, _ := regexp.Compile(`^\d{6}(\d{1})`)
	re_length_type_0, _ := regexp.Compile(`^\d{6}0(\d{15})`)
	re_length_type_1, _ := regexp.Compile(`^\d{6}1(\d{11})`)
	packet_literal_type := "100"
	packet_version_sum := 0

	// fmt.Println("Original Message: ", message_binary)

	for re_version_type.MatchString(message_binary) {
		version_type := re_version_type.FindStringSubmatch(message_binary)
		packet_version := version_type[1]
		packet_type := version_type[2]
		packet_version_sum += aocutils.StringToInt(packet_version, 2)
		fmt.Println("Version / Type: ", packet_version, " / ", packet_type)

		if packet_type == packet_literal_type {
			if re_literal.MatchString(message_binary) {
				literal := re_literal.FindStringSubmatch(message_binary)
				literal_string := strings.Join(literal[1:], "")
				literal_string = re_literal_parse.ReplaceAllString(literal_string, "$1")
				message_binary = re_literal.ReplaceAllString(message_binary, "")
				fmt.Println("Literal String: ", literal_string)
			}
		} else {
			if re_length_type_0.MatchString(message_binary) {
				message_binary = re_length_type_0.ReplaceAllString(message_binary, "")
			} else if re_length_type_1.MatchString(message_binary) {
				message_binary = re_length_type_1.ReplaceAllString(message_binary, "")
			} else {
				message_binary = ""
			}
		}
		// fmt.Println("Message: ", message_binary)
	}
	fmt.Println("Packet Version Sum: ", packet_version_sum)

}
