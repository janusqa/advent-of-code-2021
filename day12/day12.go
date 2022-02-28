package day12

import (
	"adventofcode/aocutils"
	"fmt"
	"regexp"
	"strings"
)

var input []string

func PartI(filename string) {
	input = aocutils.GetInputFromFile(filename)
	paths := [][]string{}
	for _, path := range input {
		temp_path := strings.Split(path, "-")
		//reUpper, _ := regexp.Compile(`^[A-Z]+$`)
		reStartEnd, _ := regexp.Compile(`^(start|end)$`)

		if temp_path[1] == "start" {
			paths = append(paths, reverse(temp_path))
		} else if temp_path[0] == "end" {
			paths = append(paths, reverse(temp_path))
		} else if !reStartEnd.MatchString(temp_path[1]) && !reStartEnd.MatchString(temp_path[0]) {
			paths = append(paths, temp_path)
			paths = append(paths, reverse(temp_path))
		} else {
			paths = append(paths, temp_path)
		}
	}

	routesPotential := findPaths("start", paths, map[string]int{})
	routes := []string{}
	for _, route := range routesPotential {
		if route[len(route)-1] == "end" {
			routes = append(routes, strings.Join(route, ","))
		}
	}
	fmt.Println("paths: ", paths)
	// fmt.Println("Potential routes: ", routesPotential, len(routesPotential))
	fmt.Println("routes: ", routes, len(routes))
}

func reverse(s []string) []string {

	r := make([]string, len(s))
	copy(r, s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return r
}

func findPaths(root string, branches [][]string, visited map[string]int) [][]string {
	if root == "end" {
		return [][]string{}
	}
	result := [][]string{}
	reUpper, _ := regexp.Compile(`^[A-Z]+$`)
	reLower, _ := regexp.Compile(`^[a-z]+$`)
	if reLower.MatchString(root) {
		visited[root] += 1
	}
	for _, branch := range branches {
		if root == branch[0] && ((visited[branch[1]] < 2 && reLower.MatchString(branch[1])) || reUpper.MatchString(branch[1])) {
			small_cave_quota_reached := false
			if reLower.MatchString(branch[1]) {
				for node, count := range visited {
					if (count > 1 && node != branch[1]) && visited[branch[1]] > 0 {
						small_cave_quota_reached = true
						break
					}
				}
			}
			if (!small_cave_quota_reached) || reUpper.MatchString(branch[1]) {
				pathWays := findPaths(branch[1], branches, visited)
				rootWays := [][]string{}
				if len(pathWays) > 0 {
					for _, pathWay := range pathWays {
						if root == "start" {
							rootWays = append(rootWays, append([]string{root, branch[1]}, pathWay...))
						} else {
							rootWays = append(rootWays, append([]string{branch[1]}, pathWay...))
						}
					}
				} else {
					if root == "start" {
						rootWays = append(rootWays, []string{root, branch[1]})
					} else {
						rootWays = append(rootWays, []string{branch[1]})
					}
				}
				result = append(result, rootWays...)
			}
		}
	}
	if val, ok := visited[root]; ok {
		visited[root] = val - 1
	}

	return result
}
