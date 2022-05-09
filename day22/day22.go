package day22

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

type Cuboid struct {
	xrange [2]int
	yrange [2]int
	zrange [2]int
}

type Command struct {
	operation string
	cubes     Cuboid
}

func PartI(filename string) {

	commands := []Command{}
	getInput(filename, &commands)

	// Brute Force Part I
	var statusOnPartI = make(map[string]string)
	minXyz := -50
	maxXyz := 50
	for _, command := range commands {
		for x := max(command.cubes.xrange[0], minXyz); x <= min(command.cubes.xrange[1], maxXyz); x++ {
			for y := max(command.cubes.yrange[0], minXyz); y <= min(command.cubes.yrange[1], maxXyz); y++ {
				for z := max(command.cubes.zrange[0], minXyz); z <= min(command.cubes.zrange[1], maxXyz); z++ {
					key := fmt.Sprintf("%d,%d,%d", x, y, z)
					if command.operation == "on" {
						statusOnPartI[key] = command.operation
					} else {
						delete(statusOnPartI, key)
					}
				}
			}
		}
	}
	fmt.Println("Part 1: ", len(statusOnPartI))

	// Inclusion Exclusion Principle, an optimization to brute forcing
	// https://en.wikipedia.org/wiki/Inclusion%E2%80%93exclusion_principle
	// Part II
	coreCommands := []Command{}
	for _, command := range commands {
		activeCommands := []Command{}
		// Only counting cuboids/cubes that are on
		// No need to count cuboids/cubes that are off
		if command.operation == "on" {
			activeCommands = append(activeCommands, command)
		}
		for _, core := range coreCommands {
			if intersection, ok := command.cubes.intersect(core.cubes); ok {
				if core.operation == "on" {
					activeCommands = append(activeCommands, Command{"off", intersection})
				} else {
					activeCommands = append(activeCommands, Command{"on", intersection})
				}
			}
		}
		coreCommands = append(coreCommands, activeCommands...)
	}

	statusOnPartII := 0
	for _, core := range coreCommands {
		if core.operation == "on" {
			statusOnPartII += (core.cubes.xrange[1] - core.cubes.xrange[0] + 1) *
				(core.cubes.yrange[1] - core.cubes.yrange[0] + 1) *
				(core.cubes.zrange[1] - core.cubes.zrange[0] + 1)
		} else {
			statusOnPartII -= (core.cubes.xrange[1] - core.cubes.xrange[0] + 1) *
				(core.cubes.yrange[1] - core.cubes.yrange[0] + 1) *
				(core.cubes.zrange[1] - core.cubes.zrange[0] + 1)
		}
	}
	fmt.Println("Part 2: ", statusOnPartII)

}

func (c Cuboid) intersect(c1 Cuboid) (Cuboid, bool) {

	/***
	if any axis a of cuboid c does not fall between
	the min and max bounds of axis a of cuboid b, then
	cuboids a and b do not intersect.
	***/
	if (c.xrange[0] > c1.xrange[1]) || (c1.xrange[0] > c.xrange[1]) ||
		(c.yrange[0] > c1.yrange[1]) || (c1.yrange[0] > c.yrange[1]) ||
		(c.zrange[0] > c1.zrange[1] || (c1.zrange[0] > c.zrange[1])) {
		return Cuboid{[2]int{0, 0}, [2]int{0, 0}, [2]int{0, 0}}, false
	}

	/***
	if all axes a,b,c of cuboid c falls between
	the min and max bounds of axes a,b,  of cuboid b, then cuboids a and b
	intersect. The intersection is the maximum of the lower bounds of each
	axis and the minimum of the upper bounds of each axis
	***/
	return Cuboid{
		[2]int{max(c.xrange[0], c1.xrange[0]), min(c.xrange[1], c1.xrange[1])},
		[2]int{max(c.yrange[0], c1.yrange[0]), min(c.yrange[1], c1.yrange[1])},
		[2]int{max(c.zrange[0], c1.zrange[0]), min(c.zrange[1], c1.zrange[1])},
	}, true
}

// can only be used in go versision >= 1.18
func min[T Number](a T, b T) T {
	if a < b {
		return a
	}
	return b
}

// can only be used in go versision >= 1.18
func max[T Number](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

func getInput(filename string, commands *[]Command) {

	reCommands, _ := regexp.Compile(`^(on|off) ((?:(?:x|y|z)=(?:-?\d+)\.\.(?:-?\d+),?){3})$`)
	reCubes, _ := regexp.Compile(`(?:(x|y|z)=(-?\d+)\.\.(-?\d+))`)

	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("Failed to get input from file due to %s", err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		commandLine := reCommands.FindStringSubmatch(scanner.Text())
		cubes := reCubes.FindAllStringSubmatch(commandLine[2], -1)
		*commands = append(*commands,
			Command{
				commandLine[1],
				Cuboid{
					[2]int{aocutils.StringToInt(cubes[0][2], 10), aocutils.StringToInt(cubes[0][3], 10)},
					[2]int{aocutils.StringToInt(cubes[1][2], 10), aocutils.StringToInt(cubes[1][3], 10)},
					[2]int{aocutils.StringToInt(cubes[2][2], 10), aocutils.StringToInt(cubes[2][3], 10)},
				},
			},
		)
	}
}
