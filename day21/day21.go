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
	face  int
	faces int
	rolls int
}

type Game struct {
	currentPlayer int
	maxScore      int
	boardSize     int
}

func PartI(filename string) {
	chances := 3

	deterministicGame(map[int]*Player{1: &Player{4, 0}, 2: &Player{8, 0}}, &Game{1, 1000, 10}, &DeterministicDice{1, 100, 0}, chances)

	wins := quantumGame(map[int]*Player{1: &Player{4, 0}, 2: &Player{8, 0}}, &Game{1, 21, 10}, chances, map[string][2]int{})
	fmt.Println(wins, max(wins[0], wins[1]))

}

func deterministicGame(players map[int]*Player, game *Game, dice *DeterministicDice, chances int) {
	for (*players[1]).score < (*game).maxScore && (*players[2]).score < (*game).maxScore {
		players[(*game).currentPlayer].turn(game, dice, chances)
		// fmt.Printf(
		// 	"Current Player: %d, dice: %d, times rolled: %d, position: %d, score: %d\n",
		// 	(*game).currentPlayer,
		// 	(*dice).face-1,
		// 	(*dice).rolls,
		// 	(*players[(*game).currentPlayer]).position,
		// 	(*players[(*game).currentPlayer]).score,
		// )
		(*game).currentPlayer = next((*game).currentPlayer+1, 2)

	}
	fmt.Printf("\nPart I: %d\n", min((*players[1]).score, (*players[2]).score)*(*dice).rolls)
}

func (p *Player) turn(game *Game, dice *DeterministicDice, chances int) {
	moveBy := dice.roll(chances)
	(*dice).rolls += chances
	(*p).position = next((*p).position+moveBy, (*game).boardSize)
	(*p).score += (*p).position
}

func (d *DeterministicDice) roll(chances int) int {
	moveBy := 0

	currentDice := (*d).face
	for i := currentDice; i <= (currentDice-1)+chances; i++ {
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

// can only be used in go versision >= 1.18
func max[T Number](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

// func random(lbound int, ubound int) int {
// 	rand.Seed(time.Now().UnixNano())
// 	return rand.Intn(ubound+1-lbound) + lbound
// }

func quantumGame(players map[int]*Player, game *Game, chances int, memo map[string][2]int) [2]int {

	key := fmt.Sprintf("%d%d%d%d%d", players[1].position, players[1].score, players[2].position, players[2].score, (*game).currentPlayer)

	if winners, ok := memo[key]; ok {
		return winners
	}

	if (*players[1]).score >= (*game).maxScore {
		return [2]int{1, 0}
	}

	if (*players[2]).score >= (*game).maxScore {
		return [2]int{0, 1}
	}

	wins := [2]int{0, 0}

	for i := 1; i <= chances; i++ {
		for j := 1; j <= chances; j++ {
			for k := 1; k <= chances; k++ {
				moveBy := i + j + k
				newPosition := next((*players[(*game).currentPlayer]).position+moveBy, (*game).boardSize)
				newScore := (*players[(*game).currentPlayer]).score + newPosition
				nextPlayer := next((*game).currentPlayer+1, 2)
				winners := quantumGame(
					map[int]*Player{
						(*game).currentPlayer: &Player{newPosition, newScore},
						nextPlayer:            &Player{(*players[nextPlayer]).position, (*players[nextPlayer]).score},
					},
					&Game{nextPlayer, (*game).maxScore, (*game).boardSize},
					chances,
					memo,
				)
				wins[0] += winners[0]
				wins[1] += winners[1]
			}
		}
	}

	memo[key] = wins
	return wins
}
