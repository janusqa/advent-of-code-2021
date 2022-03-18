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

	// message_hex := input[0]
	message_hex_arr := []string{
		// "C200B40A82",
		// "04005AC33890",
		// "880086C3E88112",
		// "CE00C43D881120",
		// "D8005AC2A8F0",
		// "F600BC2D8F",
		// "9C005AC2F8F0",
		// "9C0141080250320F1802104A08",
		input[0],
	}
	for _, message_hex := range message_hex_arr {
		message_binary := hex_to_binary(message_hex)
		/*** DEBUG SART ************************************************/
		// fmt.Println("Original Message: ", message_binary)
		/*** DEBUG STOP ************************************************/
		fmt.Println("Packet Version Sum: ", message_version(message_binary))
		fmt.Println("Packet Value: ", message_parse(&message_binary, 0))
	}
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

func message_version(message_binary string) int {

	re_version_type, _ := regexp.Compile(`^(\d{3})(\d{3})`)
	// re_literal_type, _ := regexp.Compile(`^(?:\d{6})(?:1(\d{4}))*0(\d{4})`) // repeat a capture group
	re_literal_type, _ := regexp.Compile(`^\d{3}100((?:1\d{4})*)(0\d{4})`) // capture a repeating group. Put extra parens around the repeating capture group
	// re_literal, _ := regexp.Compile(`(?:\d(\d{4}))`)
	//re_lengthtype, _ := regexp.Compile(`^\d{6}(\d{1})`)
	re_operator_type_0, _ := regexp.Compile(`^\d{3}(?:000|001|010|011|101|110|111)0(\d{15})`)
	re_operator_type_1, _ := regexp.Compile(`^\d{3}(?:000|001|010|011|101|110|111)1(\d{11})`)
	// packet_literal_type := "100"
	packet_version_sum := 0
	// packets := &customStack{stack: make([]string, 0)}

	// fmt.Println("Original Message: ", message_binary)

	for re_version_type.MatchString(message_binary) {
		version_type := re_version_type.FindStringSubmatch(message_binary)
		packet_version := version_type[1]
		// packet_type := version_type[2]
		packet_version_sum += aocutils.StringToInt(packet_version, 2)
		// fmt.Println("Version / Type: ", packet_version, " / ", packet_type)
		// if packet_type != packet_literal_type {
		// 	packets.Push(packet_type)
		// }

		if re_literal_type.MatchString(message_binary) {
			// literal := re_literal_type.FindStringSubmatch(message_binary)
			// literal_string := strings.Join(literal[1:], "")
			// literal_string = re_literal.ReplaceAllString(literal_string, "$1")
			message_binary = re_literal_type.ReplaceAllString(message_binary, "")
			// fmt.Println("Literal String: ", literal_string)
			// packets.Push(literal_string)
		} else {
			if re_operator_type_0.MatchString(message_binary) {
				message_binary = re_operator_type_0.ReplaceAllString(message_binary, "")
			} else if re_operator_type_1.MatchString(message_binary) {
				message_binary = re_operator_type_1.ReplaceAllString(message_binary, "")
			} else {
				message_binary = ""
			}
		}
		// fmt.Println("Message: ", message_binary)
	}
	// fmt.Println("Packet Version Sum: ", packet_version_sum)
	// fmt.Println(subpacket_stack.stack)
	return packet_version_sum
}

func message_parse(message_binary *string, indent int) int {

	re_version_type, _ := regexp.Compile(`^(\d{3})(\d{3})`)
	re_literal_type, _ := regexp.Compile(`^\d{3}100((?:1\d{4})*)(0\d{4})`)
	re_literal, _ := regexp.Compile(`(?:\d(\d{4}))`)
	re_operator_type_0, _ := regexp.Compile(`^\d{3}(?:000|001|010|011|101|110|111)0(\d{15})`)
	re_operator_type_1, _ := regexp.Compile(`^\d{3}(?:000|001|010|011|101|110|111)1(\d{11})`)

	packet_version_type := re_version_type.FindStringSubmatch(*message_binary)
	// packet_version := version_type[1]
	packet_type := packet_version_type[2]
	// packet_literal_type := "100"

	packet_translation := -1

	/*** DEBUG START ************************************************/
	// the_packet := print_packet(*message_binary)
	// fmt.Println(the_packet)
	indent_string := ""
	for i := 0; i < indent; i++ {
		indent_string = fmt.Sprintf("%s%s", indent_string, ">")
	}
	if len(indent_string) > 0 {
		indent_string = fmt.Sprintf("%s%s", indent_string, " ")
	}
	operators := map[string]string{
		"000": "sum",
		"001": "mul",
		"010": "min",
		"011": "max",
		"101": "gt",
		"110": "lt",
		"111": "eq",
	}
	if _, ok := operators[packet_type]; ok {
		fmt.Printf("%soperator: %s\n", indent_string, operators[packet_type])
	}
	/*** DEBUG STOP ************************************************/

	if re_literal_type.MatchString(*message_binary) {
		literal := re_literal_type.FindStringSubmatch(*message_binary)
		literal_string := strings.Join(literal[1:], "")
		literal_string = re_literal.ReplaceAllString(literal_string, "$1")
		*message_binary = re_literal_type.ReplaceAllString(*message_binary, "")
		/*** DEBUG START ************************************************/
		fmt.Printf("%sliteral: %d\n", indent_string, aocutils.StringToInt(literal_string, 2))
		/*** DEBUG STOP ************************************************/
		return aocutils.StringToInt(literal_string, 2)
	}

	if re_operator_type_0.MatchString(*message_binary) {
		size_in_bits_bin := re_operator_type_0.FindStringSubmatch(*message_binary)
		size_in_bits_dec := aocutils.StringToInt(size_in_bits_bin[1], 2)
		*message_binary = re_operator_type_0.ReplaceAllString(*message_binary, "")
		size_in_bits_after_parse := len(*message_binary) - size_in_bits_dec
		for len(*message_binary) > size_in_bits_after_parse {
			packet_translation = compute(packet_translation, message_parse(message_binary, indent+2), packet_type)
		}
	} else if re_operator_type_1.MatchString(*message_binary) {
		size_in_count_bin := re_operator_type_1.FindStringSubmatch(*message_binary)
		size_in_count_dec := aocutils.StringToInt(size_in_count_bin[1], 2)
		*message_binary = re_operator_type_1.ReplaceAllString(*message_binary, "")
		for packet_count := 0; packet_count < size_in_count_dec; packet_count++ {
			packet_translation = compute(packet_translation, message_parse(message_binary, indent+2), packet_type)
		}
	}

	/*** DEBUG START ************************************************/
	fmt.Printf("%send subpackets (%s)\n", indent_string, operators[packet_type])
	/*** DEBUG STOP ************************************************/

	return packet_translation
}

func compute(packet_translation int, packet_value int, packet_type string) int {

	operators := map[string]string{
		"000": "sum",
		"001": "mul",
		"010": "min",
		"011": "max",
		"101": "gt",
		"110": "lt",
		"111": "eq",
	}

	if value, ok := operators[packet_type]; ok {
		if packet_translation == -1 {
			packet_translation = packet_value
		} else {
			if value == "sum" {
				packet_translation += packet_value
			} else if value == "mul" {
				packet_translation *= packet_value
			} else if value == "min" {
				if packet_value < packet_translation {
					packet_translation = packet_value
				}
			} else if value == "max" {
				if packet_value > packet_translation {
					packet_translation = packet_value
				}
			} else if value == "gt" {
				if packet_translation > packet_value {
					packet_translation = 1
				} else {
					packet_translation = 0
				}
			} else if value == "lt" {
				if packet_translation < packet_value {
					packet_translation = 1
				} else {
					packet_translation = 0
				}
			} else if value == "eq" {
				if packet_translation == packet_value {
					packet_translation = 1
				} else {
					packet_translation = 0
				}
			}
		}
	}
	return packet_translation
}

// func print_packet(message_binary string) string {
// 	re_literal_type, _ := regexp.Compile(`^\d{3}100(?:(?:1\d{4})*)(?:0\d{4})`)
// 	re_operator_type_0, _ := regexp.Compile(`^\d{3}(?:000|001|010|011|101|110|111)0(\d{15})`)
// 	re_operator_type_1, _ := regexp.Compile(`^\d{3}(?:000|001|010|011|101|110|111)1(\d{11})`)

// 	if re_literal_type.MatchString(message_binary) {
// 		return re_literal_type.FindStringSubmatch(message_binary)[0]
// 	} else if re_operator_type_0.MatchString(message_binary) {
// 		return re_operator_type_0.FindStringSubmatch(message_binary)[0]
// 	} else if re_operator_type_1.MatchString(message_binary) {
// 		return re_operator_type_1.FindStringSubmatch(message_binary)[0]
// 	}
// 	return ""
// }
