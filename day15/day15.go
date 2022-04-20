package day15

import (
	"adventofcode/aocutils"
	"container/heap"
	"fmt"
	"math"
	"strings"
)

var input []string

func PartI(filename string) {
	input = aocutils.GetInputFromFile(filename)

	cave := [][]int{}

	// Initilize the cave map
	for row_index, row := range input {
		row_items := strings.Split(row, "")
		cave = append(cave, []int{})
		for _, column := range row_items {
			cave[row_index] = append(cave[row_index], aocutils.StringToInt(column, 10))
		}
	}

	scale := 5
	// minimum_risk_1 := get_minimum_risk(cave, 0, 0, scale, map[string]int{})
	minimum_risk_2 := get_minimum_risk(cave, 0, 0, scale)
	fmt.Println(minimum_risk_2)

}

// Dijkstra's Algorithmn
func get_minimum_risk(cave [][]int, row int, column int, scale int) int {

	num_rows := len(cave)
	num_columns := len(cave[0])
	visited := make(map[string]int)
	children := map[string][2]int{
		"left":  {0, -1},
		"right": {0, 1},
		"up":    {-1, 0},
		"down":  {1, 0},
	}
	destination := fmt.Sprintf("%d,%d", num_rows*scale-1, num_columns*scale-1)

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	item := pq_item(row, column)
	item.priority = 0
	heap.Push(&pq, item)
	// pq.update(item, item.value, 0)

	for pq.Len() > 0 {
		var item *Item
		var risk_weight int

		current_node := heap.Pop(&pq).(*Item)

		if _, ok := visited[current_node.value]; !ok {
			visited[current_node.value] = current_node.priority
			current_node_id := strings.Split(current_node.value, ",")
			current_node_row := aocutils.StringToInt(current_node_id[0], 10)
			current_node_column := aocutils.StringToInt(current_node_id[1], 10)
			if current_node.value == destination {
				break
			}

			for _, child := range children {
				row = current_node_row + child[0]
				column = current_node_column + child[1]
				if row < num_rows*scale && column < num_columns*scale && row >= 0 && column >= 0 {
					item = pq_item(row, column)
					risk_weight = get_risk_level(cave, row, column) + current_node.priority
					if risk_weight < item.priority {
						item.priority = risk_weight
					}
					heap.Push(&pq, item)
				}
			}

		}
	}
	return visited[destination]
}

func pq_item(row int, column int) *Item {
	return &Item{
		value:    fmt.Sprintf("%d,%d", row, column),
		priority: math.MaxInt16,
	}
}

func get_risk_level(cave [][]int, row int, column int) int {
	num_rows := len(cave)
	num_columns := len(cave[0])
	maximum_risk_level := 9
	tile_row := row / num_rows
	tile_column := column / num_columns
	start_row_translated := row - (tile_row * num_rows)
	start_column_translated := column - (tile_column * num_columns)
	risk_level := cave[start_row_translated][start_column_translated] + tile_row + tile_column
	if risk_level > maximum_risk_level {
		risk_level %= maximum_risk_level
		if risk_level == 0 {
			risk_level = maximum_risk_level
		}
	}
	return risk_level
}

// calculate the mimimum risk of all paths
// *** Failure ***
// func get_minimum_risk(cave [][]int, start_row int, start_column int, scale int, memo map[string]int) int {

// 	key := fmt.Sprintf("%v,%v", start_row, start_column)
// 	num_rows := len(cave)
// 	num_columns := len(cave[0])
// 	children := map[string][2]int{
// 		"left":  {0, -1},
// 		"right": {0, 1},
// 		"up":    {-1, 0},
// 		"down":  {1, 0},
// 	}

// 	if memoized_risk_levels, ok := memo[key]; ok {
// 		return memoized_risk_levels
// 	}

// 	if (start_row == num_rows*scale-1) && (start_column == num_columns*scale-1) {
// 		return risk_level(cave, start_row, start_column)
// 	}

// 	if (start_row == num_rows*scale) || (start_column == num_columns*scale) {
// 		return math.MaxInt16
// 	}

// 	minimum_risk := math.MaxInt16

// 	for _, child := range children {
// 		cumalative_risk_levels := get_minimum_risk(cave, start_row+child[0], start_column+child[1], scale, memo)
// 		if cumalative_risk_levels < minimum_risk {
// 			minimum_risk = cumalative_risk_levels
// 		}
// 	}

// 	minimum_risk += risk_level(cave, start_row, start_column)

// 	memo[key] = minimum_risk
// 	return minimum_risk
// }
