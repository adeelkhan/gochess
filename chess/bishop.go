package chess

import (
	"errors"
	"fmt"
)

type Bishop struct {
	name  string
	color string
	board *Board
}

func (b Bishop) move(color string, frow, fcol, trow, tcol int) error {
	ownPiece := IsOwn(b.color, color)
	if ownPiece != true {
		return errors.New("Invalid move: Can't move other player piece.")
	}

	validMoves := b.getValidMoves(color, frow, fcol)
	destPlace, _ := GetBoardNotation(trow, tcol)

	if Contains(validMoves, destPlace) != true {
		return errors.New("Invalid move: Bishop cant be moved there")
	}

	src, _ := b.board.get(frow, fcol)
	dest, _ := b.board.get(trow, tcol)
	if dest != nil && src.getColor() != dest.getColor() {
		fmt.Println("Capturing", dest.getName())
	}
	if b.isCheck(color, trow, tcol) == true {
		if IsCheckMate(b.board, b.getColor()) == true {
			fmt.Println("Check mate!")
			return nil
		} else {
			fmt.Println("Check!")
		}
	}
	b.board.set(frow, fcol, nil)
	b.board.set(trow, tcol, src)

	return nil
}

func (b Bishop) String() string {
	return fmt.Sprintf(b.name)
}
func (b Bishop) getColor() string {
	return b.color
}
func (b Bishop) getName() string {
	return b.name
}

func (b Bishop) getAllMoves(frow, fcol int) []string {
	validMoves := make([]string, 0)
	bishopColor := b.getColor()

	i := frow
	j := fcol
	for i+1 < 8 && j+1 < 8 { // right diag down
		piece, _ := b.board.get(i+1, j+1)

		if piece == nil || piece.getColor() == bishopColor {
			move, _ := GetBoardNotation(i+1, j+1)
			validMoves = append(validMoves, move)
			i++
			j++
		} else {
			break
		}
	}

	i = frow
	j = fcol
	for i-1 >= 0 && j+1 < 8 { // right diag up
		piece, _ := b.board.get(i-1, j+1)
		if piece == nil || piece.getColor() == bishopColor {
			move, _ := GetBoardNotation(i-1, j+1)
			validMoves = append(validMoves, move)
			i--
			j++
		} else {
			break
		}
	}
	i = frow
	j = fcol
	for i+1 < 8 && j-1 >= 0 { // left diag down
		piece, _ := b.board.get(i+1, j-1)
		if piece == nil || piece.getColor() == bishopColor {
			move, _ := GetBoardNotation(i+1, j-1)
			validMoves = append(validMoves, move)
			i++
			j--
		} else {
			break
		}
	}
	i = frow
	j = fcol
	for i-1 >= 0 && j-1 >= 0 { // left diag up
		piece, _ := b.board.get(i-1, j-1)
		if piece == nil || piece.getColor() == bishopColor {
			move, _ := GetBoardNotation(i-1, j-1)
			validMoves = append(validMoves, move)
			i--
			j--
		} else {
			break
		}
	}
	return validMoves
}
func (b Bishop) getValidMoves(color string, frow, fcol int) []string {
	validMoves := make([]string, 0)

	i := frow
	j := fcol
	for i+1 < 8 && j+1 < 8 { // right diag down
		piece, _ := b.board.get(i+1, j+1)

		if piece == nil {
			move, _ := GetBoardNotation(i+1, j+1)
			validMoves = append(validMoves, move)
			i++
			j++
		} else if piece.getColor() != color {
			move, _ := GetBoardNotation(i+1, j+1)
			validMoves = append(validMoves, move)
			break
		} else {
			break
		}
	}

	i = frow
	j = fcol
	for i-1 >= 0 && j+1 < 8 { // right diag up
		piece, _ := b.board.get(i-1, j+1)
		if piece == nil {
			move, _ := GetBoardNotation(i-1, j+1)
			validMoves = append(validMoves, move)
			i--
			j++
		} else if piece.getColor() != color {
			move, _ := GetBoardNotation(i-1, j+1)
			validMoves = append(validMoves, move)
			break
		} else {
			break
		}
	}
	i = frow
	j = fcol
	for i+1 < 8 && j-1 >= 0 { // left diag down
		piece, _ := b.board.get(i+1, j-1)
		if piece == nil {
			move, _ := GetBoardNotation(i+1, j-1)
			validMoves = append(validMoves, move)
			i++
			j--
		} else if piece.getColor() != color {
			move, _ := GetBoardNotation(i+1, j-1)
			validMoves = append(validMoves, move)
			break
		} else {
			break
		}
	}
	i = frow
	j = fcol
	for i-1 >= 0 && j-1 >= 0 { // left diag up
		piece, _ := b.board.get(i-1, j-1)
		if piece == nil {
			move, _ := GetBoardNotation(i-1, j-1)
			validMoves = append(validMoves, move)
			i--
			j--
		} else if piece.getColor() != color {
			move, _ := GetBoardNotation(i-1, j-1)
			validMoves = append(validMoves, move)
			break
		} else {
			break
		}
	}
	return validMoves
}

func (b Bishop) isCheck(color string, row, col int) bool {
	i := row
	j := col
	for i+1 < 8 && j+1 < 8 { // down diagonal right
		if HaveKing(b.board, color, i+1, j+1) == true {
			return true
		}
		i++
		j++
	}
	i = row
	j = col
	for i-1 >= 0 && j+1 < 8 { // up diagonal right
		if HaveKing(b.board, color, i-1, j+1) == true {
			return true
		}
		i--
		j++
	}
	i = row
	j = col
	for i-1 >= 0 && j+1 < 8 { // up diagonal left
		if HaveKing(b.board, color, i-1, j+1) == true {
			return true
		}
		i--
		j++
	}
	i = row
	j = col
	for i+1 < 8 && j-1 >= 0 { // down diagonal left
		if HaveKing(b.board, color, i+1, j-1) == true {
			return true
		}
		i++
		j--
	}
	return false
}
