package day21

import (
	"fmt"
)

type player struct {
	position int
	score    int
}

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func PartI(filename string) {
	dice := 1
	currentPlayer := 1
	maxScore := 1000
	boardSize := 10
	diceSize := 100
	chances := 3
	countRolls := 0
	players := make(map[int]*player)
	players[1] = &player{10, 0}
	players[2] = &player{9, 0}

	for (*players[1]).score < maxScore && (*players[2]).score < maxScore {
		players[currentPlayer].turn(boardSize, &dice, &countRolls, diceSize, chances)
		fmt.Printf(
			"Current Player: %d, dice: %d, times rolled: %d, position: %d, score: %d\n",
			currentPlayer,
			dice,
			countRolls,
			(*players[currentPlayer]).position,
			(*players[currentPlayer]).score,
		)
		currentPlayer = next(currentPlayer+1, 2)

	}
	fmt.Printf("\nPart I: %d\n", min((*players[1]).score, (*players[2]).score)*(countRolls))
}

func (p *player) turn(boardSize int, dice *int, countRolls *int, diceSize int, chances int) {
	moveBy := 0
	currentDice := *dice
	for i := currentDice; i <= (currentDice-1)+chances; i++ {
		(*dice) = next((*dice)+1, diceSize)
		(*countRolls)++
		moveBy += i
	}
	(*p).position = next((*p).position+moveBy, boardSize)
	(*p).score += (*p).position
}

func next(item int, wrapAt int) int {
	// wraps a sequence at a max position
	// back around to the beginning
	return ((item - 1) % wrapAt) + 1
}

// can only be used in go versision >= 1.18
func min[T Number](a T, b T) T {
	if a < b {
		return a
	}
	return b
}
