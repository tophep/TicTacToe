package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func ReadInt(reader *bufio.Reader) int {
	input, err := reader.ReadString('\n')

	if err != nil {
		return -1
	}

	num, err := strconv.ParseInt(strings.TrimSpace(input), 10, 32)

	if err != nil {
		return -1
	}

	return int(num)
}

func PrintIntro() {
	fmt.Println()
	fmt.Println("Welcome To Tic-Tac-Toe")
}

func PrintOptions() {
	fmt.Println()
	fmt.Println("Choose 1 To Play With A Friend")
	fmt.Println("Choose 2 To Play A Perfect AI")
	fmt.Println("Choose 3 To Play A Learning AI")
	fmt.Println()
	fmt.Println("Extras:")
	fmt.Println("Choose 4 To Train The Learning AI")
	fmt.Println("Choose 5 To Watch Game Replays")
	fmt.Println("Choose 6 To Print Game Tree")
	fmt.Println()
	fmt.Println("Choose 7 To Exit")
	fmt.Println()
	fmt.Print("Option Number: ")
}

func PrintInvalidOption() {
	fmt.Println()
	fmt.Println("Please Choose A Valid Option")
}

func PrintGameIntro() {
	fmt.Println("Take Turns Choosing A Spot To Place Your Marker")
	fmt.Println("X Goes First, Ø Will Follow, Then Back To X And So On...")
	fmt.Println()
}

func PrintMovePrompt(marker string) {
	fmt.Print(marker + "'s Move: ")
}

func PrintInvalidSpot() {
	fmt.Println("You Can't Mark That Spot")
	fmt.Println()
}

func PrintResult(result int) {
	if result == TIE {
		fmt.Println("The Game Ended In A Tie")
		return
	}

	fmt.Println(PlayerString(result) + " Won The Game!")
}

func PrintTrainingPrompt() {
	fmt.Println("The AI Will Play Itself And Learn Over Time")
	fmt.Println("How Many Times Should It Play?")
	fmt.Print()
	fmt.Print("Games: ")
}

func PrintNewAIState() {
	fmt.Println()
	fmt.Println("The AI Has Never Encountered This State Before!")
	fmt.Println()
}

func PrintNoReplays() {
	fmt.Println()
	fmt.Println("No Replays Available!")
	fmt.Println()
}

func PrintReplayIntro(replayCount int) {
	fmt.Println()
	fmt.Printf("%d Replays Available To Watch\n", replayCount)
	fmt.Println("Which One Do You Want To Watch?")
	fmt.Println()
}

func ReadReplayNumber(reader *bufio.Reader, replayCount int) int {
	number := 0
	for {
		fmt.Print("Replay Number: ")
		number = ReadInt(reader) - 1
		fmt.Println()

		if number < 0 || number >= replayCount {
			fmt.Println("Invalid Replay Number!")
			fmt.Println()
		} else {
			break
		}
	}
	return number
}
