package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Play Perfect AI
// Play Learning AI
// Train Learning AI
// Replay Games
// Switch Statement

func main() {
	records := NewGameRecords()
	reader := bufio.NewReader(os.Stdin)

	PrintIntro()

	for {
		PrintOptions()
		option := ReadInt(reader)
		fmt.Println()

		switch option {
		case 1:
			PlayGame(reader, records, "human")
		case 2:
			PlayGame(reader, records, "perfect")
		case 3:
			PlayGame(reader, records, "learning")
		case 4:
			TrainAI(reader, records)
		case 5:
			ReplayGame(reader, records)
		case 6:
			PrintGameTree(records)
		case 7:
			return
		default:
			PrintInvalidOption()
		}
	}
}

func PlayGame(reader *bufio.Reader, records *GameRecords, opponent string) {
	game := NewGameState()
	PrintGameIntro()
	fmt.Println(game.Board)

	for {
		player := game.NextPlayer()
		PrintMovePrompt(PlayerString(player))
		spot := ReadInt(reader) - 1 // Account for 0 indexing
		fmt.Println()

		if game.ValidSpot(spot) {
			game.AddMove(Move{Spot: spot, Player: player})
		} else {
			PrintInvalidSpot()
			continue
		}

		if EvaluateGame(game) {
			break
		}

		if opponent != "human" {
			player := game.NextPlayer()

			if opponent == "learning" {
				if records.Tree.FindGame(game) == nil {
					PrintNewAIState()
				}
				MakeLearnedMove(records.Tree, game)
			} else {
				MakePerfectMove(game)
			}

			PrintMovePrompt(PlayerString(player))
			fmt.Println(game.LastMarked() + 1) // Account for 0 indexing
			fmt.Println()

			if EvaluateGame(game) {
				break
			}
		}
	}
	records.AddGame(game)
}

func TrainAI(reader *bufio.Reader, records *GameRecords) {
	PrintTrainingPrompt()
	count := ReadInt(reader)

	for i := 0; i < count; i++ {
		game := NewGameState()

		for {
			MakePerfectMove(game)

			isOver, _ := game.IsOver()
			if isOver {
				break
			}

			MakeLearnedMove(records.Tree, game)

			isOver, _ = game.IsOver()
			if isOver {
				break
			}
		}

		records.AddGame(game)
	}
}

func ReplayGame(reader *bufio.Reader, records *GameRecords) {
	replayCount := len(records.Replays)
	if replayCount == 0 {
		PrintNoReplays()
		return
	}

	PrintReplayIntro(replayCount)
	number := ReadReplayNumber(reader, replayCount)
	replay := records.Replays[number]
	game := NewGameState()

	fmt.Println()
	fmt.Println(game.Board)
	fmt.Println()

	for _, move := range replay.Sequence {
		time.Sleep(time.Second)
		fmt.Println(PlayerString(move.Player)+"'s Move:", move.Spot+1)
		fmt.Println()
		game.AddMove(move)
		EvaluateGame(game)
	}
}

func PrintGameTree(records *GameRecords) {
	fmt.Println()
	fmt.Println(records.Tree)
	fmt.Println()
}

func EvaluateGame(game *GameState) bool {
	fmt.Println(game.Board)

	isOver, result := game.IsOver()

	if isOver {
		PrintResult(result)
	}
	return isOver
}
