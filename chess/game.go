package chess

import "fmt"

type Game struct {
	player1 *Player
	player2 *Player
	Board   *Board
}
type Move struct {
	from int
	to   int
}

func (g *Game) Initialize(p1 *Player, p2 *Player) {

	g.player1 = p1
	g.player2 = p2

	g.Board = &Board{}
	g.Board.SetUp()
}

func (g *Game) printBoardHeader() {
	header := fmt.Sprintf("     |  A   |  B   |  C   |  D   |  E   |  F   |  G   |  H\n")
	header += fmt.Sprintf("----------------------------------------------------------\n")
	fmt.Print(header)
}
func (g *Game) printBoardFooter() {
	footer := fmt.Sprintf("----------------------------------------------------------\n")
	footer += fmt.Sprintf("     |  A   |  B   |  C   |  D   |  E   |  F   |  G   |  H\n")
	fmt.Print(footer)
}

func (g *Game) printBoard() {
	g.printBoardHeader()
	board := g.Board
	for i := 0; i < 8; i++ {
		fmt.Printf(" %d ", 8-i)
		for j := 0; j < 8; j++ {
			piece, _ := board.get(i, j)
			switch piece.(type) {
			case Queen, King, Rook, Bishop, Knight, Pawn:
				fmt.Printf("  | %s ", piece)
			default:
				fmt.Printf("  |  %s ", "X")
			}
		}
		fmt.Println()
		fmt.Println()
	}
	g.printBoardFooter()
}

func (g *Game) Start() {
	turn := 0
	g.printBoard()
	for {
		if turn == 0 {
			turn = 0
			// move := g.player1.getMove(g.Board)
			err := g.Board.move(g.player1, "d8->d7")

			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			break

		} else {
			turn = 0
			move := g.player2.getMove(g.Board)
			err := g.Board.move(g.player2, move)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
		}
		g.printBoard()
	}

}
