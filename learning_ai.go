package main

func MakeRandomMove(game *GameState) {
	moves := game.PossibleMoves()
	game.AddMove(moves[CheapRandom(len(moves))])
}

func MakeLearnedMove(tree *GameTree, game *GameState) {
	if game == nil {
		panic("Must Pass Valid Game State")
	}

	moves := OptimalMoves(tree, game)
	move := moves[CheapRandom(len(moves))]
	game.AddMove(move)
}

func OptimalMoves(tree *GameTree, game *GameState) []Move {
	tree = tree.FindGame(game)
	possibleMoves := game.PossibleMoves()

	if tree == nil { // Nothing has been learned about this state yet
		return possibleMoves
	}

	moves := WinningMoves(tree) // Win or Tie if possible

	if len(moves) > 0 {
		return moves
	}

	moves = UnexploredMoves(tree, game) // Prefer to explore new possibilities

	if len(moves) > 0 {
		return moves
	}

	moves = GoodMoves(tree, game) // Don't play moves that always lead to a loss

	if len(moves) > 0 {
		return moves
	}

	return possibleMoves
}

func GoodMoves(tree *GameTree, game *GameState) []Move {
	moves := []Move{}
	player := game.NextPlayer()

	for _, child := range tree.Children {
		if child != nil && (child.Results[player] || child.Results[TIE]) { // If player previously won or tied
			moves = append(moves, child.Move)
		}
	}

	return moves
}

func UnexploredMoves(tree *GameTree, game *GameState) []Move {
	moves := []Move{}

	for i, child := range tree.Children {
		if child == nil && !game.SpotMarked(i) { // If the move hasn't been made before and it is a valid move
			moves = append(moves, Move{Spot: i, Player: game.NextPlayer()})
		}
	}
	return moves
}

func WinningMoves(tree *GameTree) []Move {
	moves := []Move{}

	for _, child := range tree.Children {
		if child != nil && !child.HasChildren() { // Assuming only complete games are added, nodes without children are a win or tie
			moves = append(moves, child.Move)
		}
	}
	return moves
}
