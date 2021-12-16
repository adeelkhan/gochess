package main

import "github.com/adeelkhan/gochess/chess"

func main() {
	// setup player
	p1 := chess.Player{Name: "adeel", Color: "white", Captured: make([]string, 1)}
	p2 := chess.Player{Name: "wahaj", Color: "black", Captured: make([]string, 1)}

	g := chess.Game{}
	g.Initialize(&p1, &p2)
	g.Start()
}
