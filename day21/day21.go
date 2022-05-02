package day21

import (
	"fmt"
)

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

type Player struct {
	position int
	score    int
}

type DeterministicDice struct {
	face   int
	faces  int
	rolls  int
	tumble int
}

type Game struct {
	currentPlayer int
	maxScore      int
	boardSize     int
}

func PartI(filename string) {
	game := &Game{1, 1000, 10}
	dice := &DeterministicDice{1, 100, 0, 3}
	players := make(map[int]*Player)
	players[1] = &Player{4, 0}
	players[2] = &Player{8, 0}

	for (*players[1]).score < (*game).maxScore && (*players[2]).score < (*game).maxScore {
		players[(*game).currentPlayer].turn(game, dice)
		fmt.Printf(
			"Current Player: %d, dice: %d, times rolled: %d, position: %d, score: %d\n",
			(*game).currentPlayer,
			(*dice).face-1,
			(*dice).rolls,
			(*players[(*game).currentPlayer]).position,
			(*players[(*game).currentPlayer]).score,
		)
		(*game).currentPlayer = next((*game).currentPlayer+1, 2)

	}
	fmt.Printf("\nPart I: %d\n", min((*players[1]).score, (*players[2]).score)*(*dice).rolls)
}

func (p *Player) turn(game *Game, dice *DeterministicDice) {
	moveBy := dice.roll()
	(*dice).rolls += (*dice).tumble
	(*p).position = next((*p).position+moveBy, (*game).boardSize)
	(*p).score += (*p).position
}

func (d *DeterministicDice) roll() int {
	moveBy := 0

	currentDice := (*d).face
	for i := currentDice; i <= (currentDice-1)+(*d).tumble; i++ {
		moveBy += i
		(*d).face = next((*d).face+1, (*d).faces)
	}

	return moveBy
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

// func random(lbound int, ubound int) int {
// 	rand.Seed(time.Now().UnixNano())
// 	return rand.Intn(ubound+1-lbound) + lbound
// }
