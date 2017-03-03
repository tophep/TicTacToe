package main

import "time"

func MakePerfectMove(game *GameState) {
	moves := game.PossibleMoves()

	if len(moves) == 9 { // Heuristic to speed up training, where the perfect AI plays first 0,2,4,6,8
		game.AddMove(moves[CheapRandom(9)])
		return
	}

	player := game.NextPlayer()
	scores := []int{}

	for _, move := range moves {
		childState := game.ChildState(move)
		scores = append(scores, MiniMax(childState, player))
	}

	max := MaxScore(scores)
	bestMoves := []Move{}

	for i, score := range scores {
		if score == max {
			bestMoves = append(bestMoves, moves[i])
		}
	}

	game.AddMove(bestMoves[CheapRandom(len(bestMoves))])
}

func MiniMax(game *GameState, player int) int {
	isOver, result := game.IsOver()

	if isOver && result == TIE {
		return 0
	} else if isOver && result == player {
		return 1
	} else if isOver {
		return -1
	}

	scores := []int{}
	states := game.ChildStates()

	for _, state := range states {
		score := MiniMax(state, player)
		scores = append(scores, score)
	}

	if game.NextPlayer() != player { // Choose the worst move for player / best move for other player
		return MinScore(scores)
	}
	return MaxScore(scores)
}

func MaxScore(scores []int) int {
	max := scores[0]

	for _, score := range scores {
		if score > max {
			max = score
		}
	}
	return max
}

func MinScore(scores []int) int {
	min := scores[0]

	for _, score := range scores {
		if score < min {
			min = score
		}
	}
	return min
}

func CheapRandom(upper int) int {
	return int(time.Now().UnixNano()) % upper
}
