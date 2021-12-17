package chess

import (
	"errors"
	"fmt"
)

type Pawn struct {
	name      string
	color     string
	firstMove bool
	board     *Board
}

func (p Pawn) getColor() string {
	return p.color
}
func (p Pawn) getName() string {
	return p.name
}

func (p Pawn) getValidMoves(color string, frow, fcol int) []string {
	validMoves := make([]string, 0)

	if color == "white" {
		i := frow
		j := fcol
		if i+1 < 8 { // down
			piece, _ := p.board.get(i+1, j)
			if piece == nil {
				move, _ := GetBoardNotation(i+1, j)
				validMoves = append(validMoves, move)
				if p.firstMove == true { // down
					piece, _ = p.board.get(i+2, j)
					if piece == nil {
						move, _ := GetBoardNotation(i+2, j)
						validMoves = append(validMoves, move)
					}
				}
			}
		}
		if i+1 < 8 && j+1 < 8 { // right diag down
			piece, _ := p.board.get(i+1, j+1)
			if piece != nil && piece.getColor() != color {
				move, _ := GetBoardNotation(i+1, j+1)
				validMoves = append(validMoves, move)
			}
		}
		if i+1 >= 0 && j-1 >= 0 { // left diag down
			piece, _ := p.board.get(i+1, j-1)
			if piece != nil && piece.getColor() != color {
				move, _ := GetBoardNotation(i+1, j-1)
				validMoves = append(validMoves, move)
			}
		}

	} else {
		i := frow
		j := fcol
		if i > 0 {
			piece, _ := p.board.get(i-1, j)
			if piece == nil {
				move, _ := GetBoardNotation(i-1, j)
				validMoves = append(validMoves, move)
				if p.firstMove == true {
					piece, _ = p.board.get(i-2, j)
					if piece == nil {
						move, _ := GetBoardNotation(i-2, j)
						validMoves = append(validMoves, move)
						p.firstMove = false
					}
				}
			}
		}
		if i-1 < 8 && j-1 < 8 { // left diag up
			piece, _ := p.board.get(i-1, j-1)
			if piece != nil && piece.getColor() != color {
				move, _ := GetBoardNotation(i-1, j-1)
				validMoves = append(validMoves, move)
			}
		}
		if i-1 >= 0 && j+1 >= 0 { // right diag up
			piece, _ := p.board.get(i+1, j-1)
			if piece != nil && piece.getColor() != color {
				move, _ := GetBoardNotation(i-1, j+1)
				validMoves = append(validMoves, move)
			}
		}
	}
	return validMoves
}
func (p Pawn) getAllMoves(frow, fcol int) []string {
	validMoves := make([]string, 0)

	pawnColor := p.getColor()
	if pawnColor == "white" {
		i := frow
		j := fcol
		if i+1 < 8 { // down
			piece, _ := p.board.get(i+1, j)
			if piece == nil {
				move, _ := GetBoardNotation(i+1, j)
				validMoves = append(validMoves, move)
				if p.firstMove == true { // down
					piece, _ = p.board.get(i+2, j)
					if piece == nil {
						move, _ := GetBoardNotation(i+2, j)
						validMoves = append(validMoves, move)
					}
				}
			}
		}
		if i+1 < 8 && j+1 < 8 { // right diag down
			piece, _ := p.board.get(i+1, j+1)
			if piece != nil && piece.getColor() == pawnColor {
				move, _ := GetBoardNotation(i+1, j+1)
				validMoves = append(validMoves, move)
			}
		}
		if i+1 >= 0 && j-1 >= 0 { // left diag down
			piece, _ := p.board.get(i+1, j-1)
			if piece != nil && piece.getColor() == pawnColor {
				move, _ := GetBoardNotation(i+1, j-1)
				validMoves = append(validMoves, move)
			}
		}

	} else {
		i := frow
		j := fcol
		if i > 0 {
			piece, _ := p.board.get(i-1, j)
			if piece == nil {
				move, _ := GetBoardNotation(i-1, j)
				validMoves = append(validMoves, move)
				if p.firstMove == true {
					piece, _ = p.board.get(i-2, j)
					if piece == nil {
						move, _ := GetBoardNotation(i-2, j)
						validMoves = append(validMoves, move)
						p.firstMove = false
					}
				}
			}
		}
		if i-1 < 8 && j-1 < 8 { // left diag up
			piece, _ := p.board.get(i-1, j-1)
			if piece != nil && piece.getColor() == pawnColor {
				move, _ := GetBoardNotation(i-1, j-1)
				validMoves = append(validMoves, move)
			}
		}
		if i-1 >= 0 && j+1 >= 0 { // right diag up
			piece, _ := p.board.get(i+1, j-1)
			if piece != nil && piece.getColor() == pawnColor {
				move, _ := GetBoardNotation(i-1, j+1)
				validMoves = append(validMoves, move)
			}
		}
	}
	return validMoves
}

func (p Pawn) isCheck(color string, i, j int) bool {
	if color == "white" {
		if i+1 < 8 && j+1 < 8 {
			if HaveKing(p.board, color, i+1, j+1) == true { // right diag down
				return true
			}
		}
		if i+1 >= 0 && j-1 >= 0 {
			if HaveKing(p.board, color, i+1, j-1) == true { // left diag down
				return true
			}
		}
	} else {
		if i-1 < 8 && j-1 < 8 {
			if HaveKing(p.board, color, i-1, j-1) == true { // left diag up
				return true
			}
		}
		if i-1 >= 0 && j+1 >= 0 {
			if HaveKing(p.board, color, i-1, j+1) == true { // right diag up
				return true
			}
		}
	}
	return false
}
func (p Pawn) move(color string, frow, fcol, trow, tcol int) error {
	ownPiece := IsOwn(p.color, color)
	if ownPiece != true {
		return errors.New("Invalid move: Can't move other player piece.")
	}

	// check validity of pawn moves
	validMoves := p.getValidMoves(color, frow, fcol)
	destPlace, _ := GetBoardNotation(trow, tcol)

	if Contains(validMoves, destPlace) != true {
		return errors.New("Invalid move: Pawn cant be moved there")
	}

	src, _ := p.board.get(frow, fcol)
	dest, _ := p.board.get(trow, tcol)
	if dest != nil && src.getColor() != dest.getColor() {
		fmt.Println("Capturing", dest.getName())
	}
	if p.isCheck(color, trow, tcol) == true {
		if IsCheckMate(p.board, p.getColor()) == true {
			fmt.Println("Check mate!")
			return nil
		} else {
			fmt.Println("Check!")
		}
	}
	p.board.set(frow, fcol, nil)
	p.board.set(trow, tcol, src)

	return nil
}
func (p Pawn) String() string {
	return fmt.Sprintf(p.name)
}
