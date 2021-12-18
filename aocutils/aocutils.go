package aocutils

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"unicode/utf8"
)

func GetInputFromFile(filename string) []string {

	input := []string{}

	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("Failed to get input from file due to %s", err)
		os.Exit(1)
	}

	defer file.Close()

	// The bufio.NewScanner() function is called in which the
	// object os.File passed as its parameter and this returns a
	// object bufio.Scanner which is further used on the
	// bufio.Scanner.Split() method.
	scanner := bufio.NewScanner(file)

	// The bufio.ScanLines is used as an
	// input to the method bufio.Scanner.Split()
	// and then the scanning forwards to each
	// new line using the bufio.Scanner.Scan()
	// method.
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

func StringToInt(s string, base int) int {
	value, err := strconv.ParseInt(s, base, 64)

	if err != nil {
		log.Fatalf("Error: %s. Unable to convert string \"%s\" to int", err, s)
		value = 0
	}
	return int(value)
}

func StrLen(s string) int {
	return utf8.RuneCountInString(s)
}
