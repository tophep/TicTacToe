package main

import "strconv"

type GameBoard [9]string

func NewGameBoard() *GameBoard {
	gb := &GameBoard{}
	gb.Clear()
	return gb
}

func (gb *GameBoard) Copy() *GameBoard {
	newBoard := &GameBoard{}
	for spot, marker := range gb {
		newBoard[spot] = marker
	}
	return newBoard
}

func (gb *GameBoard) Clear() {
	for i, _ := range gb {
		gb[i] = strconv.Itoa(i + 1)
	}
}

func (gb *GameBoard) AddMove(move Move) {
	gb[move.Spot] = PlayerString(move.Player)
}

func PlayerString(player int) string {
	if player == X {
		return "X"
	} else if player == O {
		return "Ø"
	} else if player == TIE {
		return "TIE"
	} else {
		panic("Invalid Player Code")
	}
}

func (gb *GameBoard) IsWon() bool {
	// Check the rows for a win
	if gb[0] == gb[1] && gb[1] == gb[2] {
		return true
	}
	if gb[3] == gb[4] && gb[4] == gb[5] {
		return true
	}
	if gb[6] == gb[7] && gb[7] == gb[8] {
		return true
	}

	// Check the columns for a win
	if gb[0] == gb[3] && gb[3] == gb[6] {
		return true
	}
	if gb[1] == gb[4] && gb[4] == gb[7] {
		return true
	}
	if gb[2] == gb[5] && gb[5] == gb[8] {
		return true
	}

	// Check the diagonals for a win
	if gb[0] == gb[4] && gb[4] == gb[8] {
		return true
	}
	if gb[2] == gb[4] && gb[4] == gb[6] {
		return true
	}

	return false
}

func (gb *GameBoard) String() string {
	s := " " + gb[0] + "  |  " + gb[1] + "  |  " + gb[2] + "\n"
	s += "--- + --- + ---\n"
	s += " " + gb[3] + "  |  " + gb[4] + "  |  " + gb[5] + "\n"
	s += "--- + --- + ---\n"
	s += " " + gb[6] + "  |  " + gb[7] + "  |  " + gb[8] + "\n"
	return s
}
