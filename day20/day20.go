package day20

import (
	"adventofcode/aocutils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type data struct {
	algorithm []string
	image     [][]string
	infinity  string
}

func PartI(filename string) {

	input := *getInput(filename)

	for i := 0; i < 50; i++ {
		(&input).zoom()
		(&input).enhance()
	}
	(&input).display()

}

func getInput(filename string) *data {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to get input from file due to %s", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	input := data{nil, [][]string{}, "."}
	scanner.Scan()
	input.algorithm = strings.Split(scanner.Text(), "")
	scanner.Scan()

	for scanner.Scan() {
		input.image = append(input.image, strings.Split(scanner.Text(), ""))
	}

	return &input
}

func (d *data) display() int {

	numOfPixelsOn := 0

	for i := 0; i < len((*d).image); i++ {
		for j := 0; j < len((*d).image[i]); j++ {
			if (*d).image[i][j] == "#" {
				numOfPixelsOn++
			}
			fmt.Print((*d).image[i][j])
			if j == len((*d).image[i])-1 {
				fmt.Println()
			}
		}
	}
	fmt.Println()
	fmt.Printf("%d pixels turned on.\n", numOfPixelsOn)
	fmt.Println(strings.Repeat("*", len((*d).image[0])))
	fmt.Println()

	return numOfPixelsOn
}

func (d *data) zoom() {

	for i := 0; i < len((*d).image); i++ {
		(*d).image[i] = append([]string{(*d).infinity}, (*d).image[i]...)
		(*d).image[i] = append((*d).image[i], (*d).infinity)
	}
	(*d).image = append([][]string{strings.Split(strings.Repeat((*d).infinity, len((*d).image[0])), "")}, (*d).image...)
	(*d).image = append((*d).image, strings.Split(strings.Repeat((*d).infinity, len((*d).image[0])), ""))

}

func (d *data) processPixel(row int, col int) string {

	binaryRepresentation := []string{}

	translation := map[string]string{
		"#": "1",
		".": "0",
	}

	for row_n := row - 1; row_n < row-1+3; row_n++ {
		for col_n := col - 1; col_n < col-1+3; col_n++ {
			if (row_n < 0 || row_n >= len((*d).image)) || (col_n < 0 || col_n >= len((*d).image[row])) {
				binaryRepresentation = append(binaryRepresentation, translation[(*d).infinity])
			} else {
				binaryRepresentation = append(binaryRepresentation, translation[(*d).image[row_n][col_n]])
			}
		}
	}
	return (*d).algorithm[aocutils.StringToInt(strings.Join(binaryRepresentation, ""), 2)]
}

func (d *data) enhance() {

	translation := map[string]string{
		"#": "1",
		".": "0",
	}

	var enhancedImage [][]string
	for i := 0; i < len((*d).image); i++ {
		enhancedImage = append(enhancedImage, []string{})
		for j := 0; j < len((*d).image[i]); j++ {
			enhancedImage[i] = append(enhancedImage[i], (*d).processPixel(i, j))

		}
	}
	(*d).image = nil
	(*d).image = enhancedImage
	(*d).infinity = (*d).algorithm[aocutils.StringToInt(strings.Repeat(translation[(*d).infinity], 9), 2)]
}
