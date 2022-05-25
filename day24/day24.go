package day24

import (
	"adventofcode/aocutils"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func PartI(filename string) {
	instructions := getInput(filename)

	// *** BEGIN DP method - SLOW
	memory := map[string]int{"w": 0, "x": 0, "y": 0, "z": 0}
	ans := monad(instructions, 0, 17, 0, memory, map[string]int{}, map[string]int{})
	fmt.Println(ans)
	// *** END DP method

	// **** BEGIN STACK METHOD
	// modelNumberMax := ""
	// modelNumberMin := ""
	// for _, tuple := range exec(compile(instructions)) {
	// 	modelNumberMax = fmt.Sprintf("%s%d", modelNumberMax, tuple[1])
	// 	modelNumberMin = fmt.Sprintf("%s%d", modelNumberMin, tuple[0])
	// }
	// fmt.Println("Max model number: ", modelNumberMax)
	// fmt.Println("Min model number: ", modelNumberMin)
	// **** END STACK METHOD
}

// ** BEGIN DP method - Very Slow
func monad(instructions []string, sequenceStart int, sequenceStop int, position int, memory map[string]int, memo map[string]int, aluCache map[string]int) int {
	key := fmt.Sprintf("pos=%d:w=%d:z=%d", position, memory["w"], memory["z"])

	if cached, ok := memo[key]; ok {
		return cached
	}

	if position > 13 {
		if memory["z"] == 0 {
			return memory["w"]
		}
		return 0
	}

	//for w := 9; w >= 1; w-- { // to get largest
	for w := 1; w <= 9; w++ { // to get smallest reverse the direction of this loop
		prevZ := memory["z"]
		alu(instructions, sequenceStart, sequenceStop, memory, w, aluCache)
		partial := monad(instructions, sequenceStart+18, sequenceStop+18, position+1, memory, memo, aluCache)
		if partial > 0 {
			if position == 13 {
				// memo[key] = partial
				return partial
			} else {
				// memo[key] = aocutils.StringToInt(fmt.Sprintf("%d%d", w, partial), 10)
				return aocutils.StringToInt(fmt.Sprintf("%d%d", w, partial), 10)
			}
		}
		memory["z"] = prevZ
		//memo[key(position, memory["w"], memory["z"])] = 0
	}
	memo[key] = 0
	return 0
}

// func key(pos int, w int, z int) string {
// 	return fmt.Sprintf("pos=%d:w=%d:z=%d", pos, w, z)
// }

func alu(instructions []string, sequenceStart int, sequenceStop int, memory map[string]int, w int, aluCache map[string]int) {

	key := fmt.Sprintf("%d:%d:%d:%d", w, sequenceStart, sequenceStop, memory["z"])

	if z, ok := aluCache[key]; ok {
		memory["z"] = z
		return
	}

	instructionSet := map[string]interface{}{"add": add, "mul": mul, "div": div, "mod": mod, "eql": eql}
	reInterpreter, _ := regexp.Compile(`^(inp|add|mul|div|mod|eql) ([wxyz]) ?([wxyz]|-?\d{1,})?$`)

	for sequenceStart <= sequenceStop {
		interpreter := reInterpreter.FindStringSubmatch(instructions[sequenceStart])
		operator := interpreter[1]
		operands := []string{}
		for i := 2; i < len(interpreter); i++ {
			operands = append(operands, interpreter[i])
		}
		switch operator {
		case "inp":
			memory[operands[0]] = w
		default:
			if (int(operands[1][0]) > 118) && (int(operands[1][0]) < 123) {
				memory[operands[0]] = instructionSet[operator].(func(int, int) int)(memory[operands[0]], memory[operands[1]])
			} else {
				memory[operands[0]] = instructionSet[operator].(func(int, int) int)(memory[operands[0]], aocutils.StringToInt(operands[1], 10))
			}
		}
		sequenceStart++
	}
	aluCache[key] = memory["z"]
}

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {

	if b != 0 {
		return a / b
	}
	return 0
}

func mod(a, b int) int {

	if a >= 0 && b > 0 {
		return a % b
	}
	return 0
}

func eql(a, b int) int {
	if a == b {
		return 1
	}
	return 0
}

func getInput(filename string) []string {

	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("Failed to get input from file due to %s", err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	instructions := []string{}
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}
	return instructions
}

// ** END DP method - Very Slow

// **** BEGIN STACK METHOD - Very Fast
// func exec(byteCode [][3]int) [14][2]int {
// 	type stackValue struct {
// 		code    [3]int
// 		pointer int
// 	}
// 	modelNumber := [14][2]int{{9, 1}, {9, 1}, {9, 1}, {9, 1}, {9, 1}, {9, 1}, {9, 1}, {9, 1}, {9, 1}, {9, 1}, {9, 1}, {9, 1}, {9, 1}, {9, 1}}
// 	callStack := []stackValue{}
// 	for index, code := range byteCode {
// 		if code[2] == 26 {
// 			offset := callStack[len(callStack)-1]
// 			callStack = callStack[:len(callStack)-1]
// 			for input := 1; input <= 9; input++ {
// 				check := input + offset.code[1] + code[0]
// 				if check > 0 && check < 10 {
// 					modelNumber[index][0] = min(modelNumber[index][0], check)
// 					modelNumber[index][1] = max(modelNumber[index][1], check)
// 					modelNumber[offset.pointer][0] = min(modelNumber[offset.pointer][0], input)
// 					modelNumber[offset.pointer][1] = max(modelNumber[offset.pointer][1], input)
// 				}
// 			}
// 		} else {
// 			callStack = append(callStack, stackValue{code, index})
// 		}
// 		//fmt.Println(modelNumber)
// 	}
// 	return modelNumber
// }

// func compile(instructions []string) [][3]int {
// 	constOneInstruction := 5
// 	constTwoInstruction := 15
// 	operationInstruction := 4
// 	endBlock := 17
// 	instructionPointer := 0
// 	byteCode := [][3]int{}
// 	constOne := 0
// 	constTwo := 0
// 	operation := 0

// 	reInterpreter, _ := regexp.Compile(`^(inp|add|mul|div|mod|eql) ([wxyz]) ?([wxyz]|-?\d{1,})?$`)

// 	for _, instruction := range instructions {
// 		if instructionPointer == operationInstruction {
// 			operation = aocutils.StringToInt(reInterpreter.FindStringSubmatch(instruction)[3], 10)
// 		}
// 		if instructionPointer == constOneInstruction {
// 			constOne = aocutils.StringToInt(reInterpreter.FindStringSubmatch(instruction)[3], 10)
// 		}
// 		if instructionPointer == constTwoInstruction {
// 			constTwo = aocutils.StringToInt(reInterpreter.FindStringSubmatch(instruction)[3], 10)
// 		}
// 		if instructionPointer == endBlock {
// 			byteCode = append(byteCode, [3]int{constOne, constTwo, operation})
// 			instructionPointer = 0
// 		} else {
// 			instructionPointer++
// 		}
// 	}
// 	return byteCode
// }

// func max[T Number](a T, b T) T {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func min[T Number](a T, b T) T {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
// **** END STACK METHOD - Very Fast
