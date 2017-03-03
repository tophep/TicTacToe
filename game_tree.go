package main

import (
	"fmt"
	"strings"
)

type GameTree struct {
	Move     Move
	Results  [3]bool      // Cache the results of children for quick lookup, if Results[0] - X won, if Results[1] - O won, if Results[2] - Tie
	Children [9]*GameTree // Maps possible next moves to the next node
}

func NewGameTree(move Move, result int) *GameTree {
	node := &GameTree{Move: move}
	node.AddResult(result)
	return node
}

func (root *GameTree) AddGame(gs *GameState) {
	isOver, result := gs.IsOver()
	if !isOver {
		panic("Can't Add Incomplete Game To Tree")
	}
	root.AddGameHelper(gs, result, 0)
}

func (root *GameTree) AddGameHelper(gs *GameState, result, depth int) {
	if depth == len(gs.Sequence) {
		return
	}

	child := root.AddChild(gs.Sequence[depth], result)
	child.AddGameHelper(gs, result, depth+1)
}

func (root *GameTree) AddChild(move Move, result int) *GameTree {
	child := root.Children[move.Spot]

	if child == nil {
		child = NewGameTree(move, result)
		root.Children[move.Spot] = child
	} else {
		child.AddResult(result)
	}

	return child
}

func (root *GameTree) AddResult(result int) {
	root.Results[result] = true
}

func (root *GameTree) FindGame(gs *GameState) *GameTree { // Walk the tree using the GameState
	node := root

	for _, move := range gs.Sequence {
		node = node.Children[move.Spot]
		if node == nil {
			return nil
		}
	}
	return node
}

func (root *GameTree) HasChildren() bool {
	for _, child := range root.Children {
		if child != nil {
			return true
		}
	}
	return false
}

func (root *GameTree) String() string {
	return root.StringHelper(0)
}

func (root *GameTree) StringHelper(depth int) string {
	s := "~ Tree Root"
	if depth > 0 {
		s = fmt.Sprintf("%d. %s Marked %d", depth, PlayerString(root.Move.Player), root.Move.Spot)
	}

	isLeaf := !root.HasChildren()

	if isLeaf && root.Results[X] {
		s += " | X Won"
	}

	if isLeaf && root.Results[O] {
		s += " | Ø Won"
	}

	if isLeaf && root.Results[TIE] {
		s += " | Tie"
	}

	for _, child := range root.Children {
		if child != nil {
			s += "\n" + strings.Repeat("  ", depth+1) + child.StringHelper(depth+1)
		}
	}

	return s
}
