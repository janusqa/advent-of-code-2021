package day23

import (
	"container/heap"
	"fmt"
	"math"
)

const AMBER = 1
const BRONZE = 2
const COPPER = 3
const DESERT = 4
const ROOM_SIZE = 4
const POPULATION = ROOM_SIZE * 4

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}
type Room [ROOM_SIZE]int
type Location [2]int
type State []interface{}
type Node struct {
	state     State
	cost      int
	amphipods [POPULATION]Amphipod
}
type Amphipod struct {
	class    int
	location Location
}

func PartI(filename string) {

	fmt.Println("Least Energy Consumed: ", aStarSearch())

}

func aStarSearch() int {

	// initial := Node{
	// 	State{0, 0, Room{BRONZE, AMBER}, 0, Room{COPPER, DESERT}, 0, Room{BRONZE, COPPER}, 0, Room{DESERT, AMBER}, 0, 0},
	// 	0,
	// 	[POPULATION]Amphipod{
	// 		{AMBER, Location{2, 1}},
	// 		{AMBER, Location{8, 1}},
	// 		{BRONZE, Location{2, 0}},
	// 		{BRONZE, Location{6, 0}},
	// 		{COPPER, Location{4, 0}},
	// 		{COPPER, Location{6, 1}},
	// 		{DESERT, Location{4, 1}},
	// 		{DESERT, Location{8, 0}},
	// 	},
	// }

	// initial := Node{
	// 	State{0, 0, Room{COPPER, BRONZE}, 0, Room{BRONZE, DESERT}, 0, Room{DESERT, AMBER}, 0, Room{AMBER, COPPER}, 0, 0},
	// 	0,
	// 	[POPULATION]Amphipod{
	// 		{AMBER, Location{6, 1}},
	// 		{AMBER, Location{8, 0}},
	// 		{BRONZE, Location{2, 1}},
	// 		{BRONZE, Location{4, 0}},
	// 		{COPPER, Location{2, 0}},
	// 		{COPPER, Location{8, 1}},
	// 		{DESERT, Location{4, 1}},
	// 		{DESERT, Location{6, 0}},
	// 	},
	// }

	// final := Node{
	// 	State{0, 0, Room{AMBER, AMBER}, 0, Room{BRONZE, BRONZE}, 0, Room{COPPER, COPPER}, 0, Room{DESERT, DESERT}, 0, 0},
	// 	0,
	// 	[POPULATION]Amphipod{
	// 		{AMBER, Location{2, 0}},
	// 		{AMBER, Location{2, 1}},
	// 		{BRONZE, Location{4, 0}},
	// 		{BRONZE, Location{4, 1}},
	// 		{COPPER, Location{6, 0}},
	// 		{COPPER, Location{6, 1}},
	// 		{DESERT, Location{8, 0}},
	// 		{DESERT, Location{8, 0}},
	// 	},
	// }

	// initial := Node{
	// 	State{
	// 		0,
	// 		0,
	// 		Room{BRONZE, DESERT, DESERT, AMBER},
	// 		0,
	// 		Room{COPPER, COPPER, BRONZE, DESERT},
	// 		0,
	// 		Room{BRONZE, BRONZE, AMBER, COPPER},
	// 		0,
	// 		Room{DESERT, AMBER, COPPER, AMBER},
	// 		0,
	// 		0,
	// 	},
	// 	0,
	// 	[POPULATION]Amphipod{
	// 		{AMBER, Location{2, 3}},
	// 		{AMBER, Location{6, 2}},
	// 		{AMBER, Location{8, 1}},
	// 		{AMBER, Location{8, 3}},
	// 		{BRONZE, Location{2, 0}},
	// 		{BRONZE, Location{4, 2}},
	// 		{BRONZE, Location{6, 0}},
	// 		{BRONZE, Location{6, 1}},
	// 		{COPPER, Location{4, 0}},
	// 		{COPPER, Location{4, 1}},
	// 		{COPPER, Location{6, 3}},
	// 		{COPPER, Location{8, 2}},
	// 		{DESERT, Location{2, 1}},
	// 		{DESERT, Location{2, 2}},
	// 		{DESERT, Location{4, 3}},
	// 		{DESERT, Location{8, 0}},
	// 	},
	// }

	initial := Node{
		State{
			0,
			0,
			Room{COPPER, DESERT, DESERT, BRONZE},
			0,
			Room{BRONZE, COPPER, BRONZE, DESERT},
			0,
			Room{DESERT, BRONZE, AMBER, AMBER},
			0,
			Room{AMBER, AMBER, COPPER, COPPER},
			0,
			0,
		},
		0,
		[POPULATION]Amphipod{
			{AMBER, Location{6, 2}},
			{AMBER, Location{6, 3}},
			{AMBER, Location{8, 0}},
			{AMBER, Location{8, 1}},
			{BRONZE, Location{2, 3}},
			{BRONZE, Location{4, 0}},
			{BRONZE, Location{4, 2}},
			{BRONZE, Location{6, 1}},
			{COPPER, Location{2, 0}},
			{COPPER, Location{4, 1}},
			{COPPER, Location{8, 2}},
			{COPPER, Location{8, 3}},
			{DESERT, Location{2, 1}},
			{DESERT, Location{2, 2}},
			{DESERT, Location{4, 3}},
			{DESERT, Location{6, 0}},
		},
	}

	final := Node{
		State{
			0,
			0,
			Room{AMBER, AMBER, AMBER, AMBER},
			0,
			Room{BRONZE, BRONZE, BRONZE, BRONZE},
			0,
			Room{COPPER, COPPER, COPPER, COPPER},
			0,
			Room{DESERT, DESERT, DESERT, DESERT},
			0,
			0,
		},
		0,
		[POPULATION]Amphipod{
			{AMBER, Location{2, 0}},
			{AMBER, Location{2, 1}},
			{AMBER, Location{2, 2}},
			{AMBER, Location{2, 3}},
			{BRONZE, Location{4, 0}},
			{BRONZE, Location{4, 1}},
			{BRONZE, Location{4, 2}},
			{BRONZE, Location{4, 3}},
			{COPPER, Location{6, 0}},
			{COPPER, Location{6, 1}},
			{COPPER, Location{6, 2}},
			{COPPER, Location{6, 3}},
			{DESERT, Location{8, 0}},
			{DESERT, Location{8, 1}},
			{DESERT, Location{8, 2}},
			{DESERT, Location{8, 3}},
		},
	}

	visited := make(map[string]int)
	destination := (&final).serialize()

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	item := pq_item(&initial)
	item.priority = 0 // (&initial).h()
	heap.Push(&pq, item)
	for pq.Len() > 0 {
		currentNode := heap.Pop(&pq).(*Item)
		nodeKey := (*currentNode).value.serialize()
		// fmt.Println(nodeKey, (*currentNode).priority)
		if _, ok := visited[nodeKey]; !ok {
			visited[nodeKey] = (*currentNode).priority //(*(*currentNode).value).cost

			if nodeKey == destination {
				break
			}

			for _, amphipod := range (*currentNode).value.amphipods {
				for indexHall, valueHall := range (*currentNode).value.state {
					if (Location{indexHall, 0}.inHall()) {
						to := Location{indexHall, 0}
						if (*currentNode).value.valid(amphipod.location, to) {
							cost := (*currentNode).value.g(&amphipod, amphipod.location, to)
							newNode := (*currentNode).value.child(amphipod.location, to)
							item := pq_item(newNode)
							(*newNode).cost = cost
							priority := (*newNode).cost + newNode.h()
							if priority < (*item).priority {
								(*item).priority = priority
							}
							heap.Push(&pq, item)
						}
					} else {
						for indexRoom := range valueHall.(Room) {
							to := Location{indexHall, indexRoom}
							if (*currentNode).value.valid(amphipod.location, to) {
								cost := (*currentNode).value.g(&amphipod, amphipod.location, to)
								newNode := (*currentNode).value.child(amphipod.location, to)
								item := pq_item(newNode)
								(*newNode).cost = cost
								priority := (*newNode).cost + newNode.h()
								if priority < (*item).priority {
									(*item).priority = priority
								}
								heap.Push(&pq, item)
							}
						}
					}
				}
			}
		}
	}
	return visited[destination]
}

func (n *Node) serialize() string {
	key := ""
	for index, value := range (*n).state {
		if (Location{index, 0}.inHall()) {
			key = fmt.Sprintf("%s%d", key, value.(int))
		} else {
			for _, value := range value.(Room) {
				key = fmt.Sprintf("%s%d", key, value)
			}
		}
	}
	return key
}

func (n *Node) child(from Location, to Location) *Node {

	newNode := Node{State{}, 0, [POPULATION]Amphipod{}}

	// deep copy  current node to new node
	newNode.state = append(newNode.state, (*n).state...)

	for index, amphipod := range (*n).amphipods {
		newLocation := amphipod.location
		if (amphipod.location[0] == from[0]) && (amphipod.location[1] == from[1]) {
			newLocation[0] = to[0]
			newLocation[1] = to[1]

			// update the halls and rooms as needed
			if to.inHall() {
				newNode.state[to[0]] = newNode.state[from[0]].(Room)[from[1]]
				room := (newNode.state[from[0]]).(Room)
				room[from[1]] = 0
				newNode.state[from[0]] = room
			} else {
				toRoom := (newNode.state[to[0]]).(Room)
				if from.inHall() {
					toRoom[to[1]] = newNode.state[from[0]].(int)
					newNode.state[to[0]] = toRoom
					newNode.state[from[0]] = 0
				} else {
					fromRoom := (newNode.state[from[0]]).(Room)
					toRoom[to[1]] = fromRoom[from[1]]
					newNode.state[to[0]] = toRoom
					fromRoom[from[1]] = 0
					newNode.state[from[0]] = fromRoom
				}
			}
		}
		// update list of amphipod locations as needed
		newNode.amphipods[index] = Amphipod{amphipod.class, newLocation}
	}
	return &newNode
}

func (n *Node) g(amphipod *Amphipod, from Location, to Location) int {
	// Calculates cost from the start/initial node/state to another node/state
	steps := 0

	steps += abs(from[0] - to[0])
	if from.inRoom() {
		steps += abs(0-from[1]) + 1
	}
	if to.inRoom() {
		steps += abs(0-to[1]) + 1
	}

	return (steps * unitEnergy((*amphipod).class)) + (*n).cost
}

func (n *Node) h() int {
	// Returns a heuristic of Amphipods not in position
	// if totally out of place the weight is 2
	// if in place in the first Room but elment in second Room is not in place the weight is 1
	// if in place (in proper burrow in room 1 or in place in proper burrow in room 0 while room 1 occupant is also in place then weight 0)

	count := 0
	for _, amphipod := range (*n).amphipods {
		if amphipod.class*2 != amphipod.location[0] {
			count += (n.g(&amphipod, amphipod.location, Location{amphipod.class * 2, 0}) - (*n).cost)
			// if !n.valid(amphipod.location, Location{amphipod.class * 2, ROOM_SIZE - 1}) {
			// 	count += 1000
			// }
		}
	}

	return count
}

func unitEnergy(amphipodType int) int {
	energy := 0
	switch amphipodType {
	case AMBER:
		energy = 1
	case BRONZE:
		energy = 10
	case COPPER:
		energy = 100
	case DESERT:
		energy = 1000
	}
	return energy
}

func (n *Node) valid(from Location, to Location) bool {

	// A non move, that is trying to move from where you are to where you are
	// or trying to move within a burrow is not really a valid move
	if from[0] == to[0] {
		return false
	}

	// Once moved into hall. Cannot move again unless moving to a room
	if from.inHall() && to.inHall() {
		return false
	}

	// When moving to a location that location must be vacant (in this case hall location)
	if to.inHall() {
		if (*n).state[to[0]].(int) != 0 {
			return false
		}
	}

	// cannot move past other amphipods in hall
	left := min(from[0], to[0])
	right := max(from[0], to[0])
	for i := left; i <= right; i++ {
		if (i != from[0]) && (Location{i, 0}.inHall()) && ((*n).state[i].(int) != 0) {
			return false
		}
	}

	// when moving from a room
	// 1. cannot move past an occupied room
	if from.inRoom() {
		for i := 0; i < from[1]; i++ {
			if (*n).state[from[0]].(Room)[i] != 0 {
				return false
			}
		}
	}

	// When moving to a location that location must be vacant (in this case room location)
	if to.inRoom() {
		if (*n).state[to[0]].(Room)[to[1]] != 0 {
			return false
		}
	}

	// In order to move into a room
	// It must be the correct burrow for that occupant
	if to.inRoom() {
		amphipodType := 0
		if from.inHall() {
			amphipodType = (*n).state[from[0]].(int)
		} else {
			amphipodType = (*n).state[from[0]].(Room)[from[1]]
		}
		if to[0]/2 != amphipodType {
			return false
		}
	}

	// In order to move into a room
	// Any occupants in the burrow must be of like type
	if to.inRoom() {
		amphipodType := 0
		if from.inHall() {
			amphipodType = (*n).state[from[0]].(int)
		} else {
			amphipodType = (*n).state[from[0]].(Room)[from[1]]
		}
		for i := 0; i < ROOM_SIZE; i++ {
			if ((*n).state[to[0]].(Room)[i] != amphipodType) && ((*n).state[to[0]].(Room)[i] != 0) {
				return false
			}
		}
	}

	// if moving into a burrow, you must move to deepest room possible in burrow
	if to.inRoom() {
		for i := to[1] + 1; i < ROOM_SIZE; i++ {
			if (*n).state[to[0]].(Room)[i] == 0 {
				return false
			}
		}
	}

	return true
}

func (l Location) inHall() bool {
	// returns true if in the hallway
	if (l[0]%2 != 0) || (l[0] < 2) || (l[0] > 8) {
		return true
	}
	return false
}

func (l Location) inRoom() bool {
	// returns true if in a room
	if !l.inHall() && (l[1] >= 0) && (l[1] < ROOM_SIZE) {
		return true
	}
	return false
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

// can only be used in go versision >= 1.18
func abs[T Number](a T) T {
	if a > 0 {
		return a
	}
	return -a
}

func pq_item(node *Node) *Item {
	return &Item{
		value:    node,
		priority: math.MaxInt,
	}
}
