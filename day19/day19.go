package day19

import (
	"adventofcode/aocutils"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type location struct {
	x int
	y int
	z int
}

type scannerList [][]location

var scanners scannerList = scannerList{}

func PartI(filename string) {
	getInput(filename)

	known := MakeSet()
	for _, location := range scanners[0] {
		known.Add(location)
	}
	scanners.remove(0)
	scanner_locations := []location{}
	scanner_locations = append(scanner_locations, location{0, 0, 0})

	for len(scanners) > 0 {
		found_scanner, found_orientation, offset := overlap(known, scanners)
		if found_orientation != nil {
			for _, location := range found_orientation {
				known.Add(location)
			}
			scanners.remove(found_scanner)
			scanner_locations = append(scanner_locations, offset)
		}
		fmt.Println("Part 1: ", known.Size(), len(scanners), offset)
	}

	max_manhattan := -1
	for _, location_1 := range scanner_locations {
		for _, location_2 := range scanner_locations {
			if location_1 != location_2 {
				if max_manhattan == -1 {
					max_manhattan = location_1.manhattan(location_2)
				} else {
					max_manhattan = max(location_1.manhattan(location_2), max_manhattan)
				}
			}
		}
	}
	fmt.Println("Part 2: ", max_manhattan)
}

// remove element from slice
func (s *scannerList) remove(index int) {
	// Remove element at index from slice
	copy((*s)[index:], (*s)[index+1:])
	(*s)[len((*s))-1] = nil
	(*s) = (*s)[:len((*s))-1]
}

func (l1 location) manhattan(l2 location) int {
	return abs(l1.x-l2.x) + abs(l1.y-l2.y) + abs(l1.z-l2.z)
}

func abs(i int) int {
	if i < 0 {
		return i * (-1)
	}
	return i
}

func max(m1 int, m2 int) int {
	if m1 > m2 {
		return m1
	}
	return m2
}

func overlap(known *Set, scanners scannerList) (int, []location, location) {
	for scannerId, scanner := range scanners {
		orientations := getScannerData(scanner)
		for _, orientation := range orientations {
			distances := make(map[location]int)
			for p1 := range (*known).m {
				for _, p2 := range orientation {
					distances[location{p1.x - p2.x, p1.y - p2.y, p1.z - p2.z}] += 1
				}
			}
			for offset, count := range distances {
				if count >= 12 {
					for index := range orientation {
						orientation[index].x += offset.x
						orientation[index].y += offset.y
						orientation[index].z += offset.z
					}
					return scannerId, orientation, offset
				}
			}
		}
	}
	return -1, nil, location{}
}

func getScannerData(scanner []location) scannerList {

	rawData := scannerList{}

	for _, location := range scanner {
		rawData = append(rawData, location.getOrientations())
	}

	return sortByOrientation(rawData)
}

func (l location) getOrientations() []location {
	x := l.x
	y := l.y
	z := l.z

	return []location{
		{x, y, z},
		{x, -z, y},
		{x, -y, -z},
		{x, z, -y},
		{-x, y, -z},
		{-x, -z, -y},
		{-x, -y, z},
		{-x, z, y},
		{y, x, -z},
		{y, -z, -x},
		{y, -x, z},
		{y, z, x},
		{-y, x, z},
		{-y, -z, x},
		{-y, -x, -z},
		{-y, z, -x},
		{z, y, -x},
		{z, x, y},
		{z, -y, x},
		{z, -x, -y},
		{-z, y, x},
		{-z, -x, y},
		{-z, -y, -x},
		{-z, x, -y},
	}
}

func sortByOrientation(rawData scannerList) scannerList {

	sortedRotations := scannerList{}

	for i := 0; i < len(rawData[0]); i++ {
		sortedRotations = append(sortedRotations, []location{})
		for _, pointRotations := range rawData {
			sortedRotations[i] = append(sortedRotations[i], pointRotations[i])
		}
	}

	return sortedRotations
}

func getInput(filename string) {

	re_scanner, _ := regexp.Compile(`^-{3} scanner (\d{1,}) -{3}$`)
	re_beacon, _ := regexp.Compile(`^(-?\d{1,}),(-?\d{1,}),(-?\d{1,})$`)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to get input from file due to %s", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	currentScanner := -1
	for scanner.Scan() {
		input := scanner.Text()

		if re_scanner.MatchString(input) {
			currentScanner = aocutils.StringToInt(re_scanner.FindStringSubmatch(input)[1], 10)
			scanners = append(scanners, []location{})
		} else if re_beacon.MatchString(input) {
			coordinates := re_beacon.FindStringSubmatch(input)
			scanners[currentScanner] = append(scanners[currentScanner],
				location{
					aocutils.StringToInt(coordinates[1], 10),
					aocutils.StringToInt(coordinates[2], 10),
					aocutils.StringToInt(coordinates[3], 10),
				},
			)
		}
	}
}
