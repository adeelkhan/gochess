package chess

import "fmt"

type Player struct {
	Name     string
	Color    string
	Captured []string
}

func (p *Player) getMove(board *Board) string {
	fmt.Printf("Player: %s (%s):", p.Name, p.Color)
	var move string
	fmt.Scanf("%s", &move)
	return move
}
