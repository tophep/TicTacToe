package main

type GameRecords struct {
	Replays []*GameState
	Tree    *GameTree
}

func (records *GameRecords) AddGame(game *GameState) {
	records.Replays = append(records.Replays, game)
	records.Tree.AddGame(game)
}

func NewGameRecords() *GameRecords {
	return &GameRecords{Tree: &GameTree{}}
}
