package chess

import (
	"errors"
	"strconv"
	"strings"
)

type Board struct {
	board     [8][8]Box
	whiteKing string
	blackKing string
}

func (b *Board) setKingLocation(color string, kingPosition string) {
	if color == "white" {
		b.whiteKing = kingPosition
	} else {
		b.blackKing = kingPosition
	}

}
func (b *Board) getKingLocation(color string) string {
	if color == "white" {
		return b.whiteKing
	} else {
		return b.blackKing
	}
}

func (b *Board) set(i, j int, p Piece) {
	b.board[i][j].setPiece(p)
}
func (b *Board) get(i, j int) (Piece, error) {
	piece := b.board[i][j].getPiece()
	return piece, nil
}

func (b *Board) move(p *Player, move string) error {
	frow, fcol, trow, tcol, err := b.parseMove(move)
	if err != nil {
		return err
	}

	err = nil
	piece, err := b.get(frow, fcol)
	if err == nil && piece != nil {
		err = piece.move(p.Color, frow, fcol, trow, tcol)
	}
	return err
}

func (b *Board) parseMove(move string) (int, int, int, int, error) {
	playerMove := strings.Split(move, "->")

	colChar := string(playerMove[0][0])
	rowIndex := string(playerMove[0][1])

	row, err := strconv.Atoi(rowIndex)
	if err != nil {
		return 0, 0, 0, 0, errors.New("Move is not parseable")
	}
	frow, fcol := GetIndex(colChar, row)

	colChar = string(playerMove[1][0])
	rowIndex = string(playerMove[1][1])

	row, err = strconv.Atoi(rowIndex)
	if err != nil {
		return 0, 0, 0, 0, errors.New("Move is not parseable")
	}
	trow, tcol := GetIndex(colChar, row)
	return frow, fcol, trow, tcol, nil
}

func (b *Board) SetUp() {
	// setup board
	// setting up white
	i, j := GetIndex("a", 8)
	b.set(i, j, Rook{"Rw", "white", b})
	i, j = GetIndex("b", 8)
	b.set(i, j, Knight{"Kw", "white", b})
	i, j = GetIndex("c", 8)
	b.set(i, j, Bishop{"Bw", "white", b})
	i, j = GetIndex("d", 8)
	b.set(i, j, Queen{"Qw", "white", b})
	i, j = GetIndex("e", 8)
	b.set(i, j, King{"+w", "white", b})
	b.setKingLocation("white", "e8")
	i, j = GetIndex("f", 8)
	b.set(i, j, Knight{"Kw", "white", b})
	i, j = GetIndex("g", 8)
	b.set(i, j, Bishop{"Bw", "white", b})
	i, j = GetIndex("h", 8)
	b.set(i, j, Rook{"Rw", "white", b})

	// setting up pawn
	i, j = GetIndex("a", 7)
	b.set(i, j, Pawn{"Pw", "white", true, b})
	i, j = GetIndex("b", 7)
	b.set(i, j, Pawn{"Pw", "white", true, b})
	i, j = GetIndex("c", 7)
	b.set(i, j, Pawn{"Pw", "white", true, b})
	i, j = GetIndex("d", 6)
	b.set(i, j, Pawn{"Pw", "white", true, b})
	i, j = GetIndex("e", 7)
	b.set(i, j, Pawn{"Pw", "white", true, b})
	i, j = GetIndex("f", 7)
	b.set(i, j, Pawn{"Pw", "white", true, b})
	i, j = GetIndex("g", 7)
	b.set(i, j, Pawn{"Pw", "white", true, b})
	i, j = GetIndex("h", 7)
	b.set(i, j, Pawn{"Pw", "white", true, b})

	// setting up black
	i, j = GetIndex("a", 1)
	b.set(i, j, Rook{"Rb", "black", b})
	i, j = GetIndex("b", 1)
	b.set(i, j, Knight{"Kb", "black", b})
	i, j = GetIndex("c", 1)
	b.set(i, j, Bishop{"Bb", "black", b})
	i, j = GetIndex("d", 5)
	b.set(i, j, King{"+b", "black", b})
	b.setKingLocation("black", "d5")
	i, j = GetIndex("e", 1)
	b.set(i, j, Queen{"Qb", "black", b})
	i, j = GetIndex("f", 1)
	b.set(i, j, Knight{"Kb", "black", b})
	i, j = GetIndex("g", 1)
	b.set(i, j, Bishop{"Bb", "black", b})
	i, j = GetIndex("h", 1)
	b.set(i, j, Rook{"Rb", "black", b})

	// setting up pawn
	i, j = GetIndex("a", 2)
	b.set(i, j, Pawn{"Pb", "black", true, b})
	i, j = GetIndex("b", 2)
	b.set(i, j, Pawn{"Pb", "black", true, b})
	i, j = GetIndex("c", 2)
	b.set(i, j, Pawn{"Pb", "black", true, b})
	i, j = GetIndex("d", 2)
	b.set(i, j, Pawn{"Pb", "black", true, b})
	i, j = GetIndex("e", 2)
	b.set(i, j, Pawn{"Pb", "black", true, b})
	i, j = GetIndex("f", 2)
	b.set(i, j, Pawn{"Pb", "black", true, b})
	i, j = GetIndex("g", 2)
	b.set(i, j, Pawn{"Pb", "black", true, b})
	i, j = GetIndex("h", 2)
	b.set(i, j, Pawn{"Pb", "black", true, b})
}
