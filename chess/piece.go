package chess

type Piece interface {
	move(color string, frow, fcol, trow, tcol int) error
	getAllMoves(frow, fcol int) []string
	String() string
	getColor() string
	getName() string
}
