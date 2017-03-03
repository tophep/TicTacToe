package main

const (
	X   = 0
	O   = 1
	TIE = 2
)

func OtherPlayer(player int) int {
	return 1 - player
}

type Move struct {
	Spot   int // 0 through 8
	Player int // 0 or 1
}

type GameState struct {
	Sequence []Move
	Board    *GameBoard
}

func NewGameState() *GameState {
	return &GameState{
		Board: NewGameBoard(),
	}
}

func (gs *GameState) Copy() *GameState {
	newSeq := make([]Move, len(gs.Sequence))
	copy(newSeq, gs.Sequence)
	return &GameState{Sequence: newSeq, Board: gs.Board.Copy()}
}

func (gs *GameState) NextPlayer() int {
	return len(gs.Sequence) % 2 // Oscilates between 0 and 1
}

func (gs *GameState) LastPlayer() int {
	return (len(gs.Sequence) + 1) % 2 // Oscilates between 0 and 1
}

func (gs *GameState) LastMarked() int {
	if len(gs.Sequence) == 0 {
		return -1
	}
	return gs.Sequence[len(gs.Sequence)-1].Spot
}

func (gs *GameState) SpotMarked(spot int) bool {
	for _, move := range gs.Sequence {
		if move.Spot == spot {
			return true
		}
	}
	return false
}

func (gs *GameState) ValidSpot(spot int) bool {
	return spot >= 0 && spot <= 8 && !gs.SpotMarked(spot)
}

// Might end up unneeded
// func (gs *GameState) ValidMove(move Move) bool {
// 	return !gs.SpotMarked(move.Spot)
// }

func (gs *GameState) AddMove(move Move) {
	gs.Sequence = append(gs.Sequence, move)
	gs.Board.AddMove(move)
}

func (gs *GameState) PossibleMoves() []Move {
	moves := []Move{}
	player := gs.NextPlayer()

	for spot := 0; spot < 9; spot++ {
		if !gs.SpotMarked(spot) {
			moves = append(moves, Move{Spot: spot, Player: player})
		}
	}
	return moves
}

func (gs *GameState) ChildState(move Move) *GameState {
	child := gs.Copy()
	child.AddMove(move)
	return child
}

func (gs *GameState) ChildStates() []*GameState {
	children := []*GameState{}
	moves := gs.PossibleMoves()

	for _, move := range moves {
		children = append(children, gs.ChildState(move))
	}
	return children
}

func (gs *GameState) IsOver() (bool, int) {
	if gs.Board.IsWon() {
		return true, gs.LastPlayer()
	}

	if len(gs.Sequence) == 9 {
		return true, TIE
	}

	return false, -1
}
