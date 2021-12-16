package chess

import (
	"errors"
	"fmt"
)

type Rook struct {
	name  string
	color string
	board *Board
}

func (r Rook) getValidMoves(color string, frow, fcol int) []string {
	validMoves := make([]string, 0)

	i := frow
	j := fcol
	for i+1 < 8 { // down
		piece, _ := r.board.get(i+1, j)
		if piece == nil {
			move, _ := GetBoardNotation(i+1, j)
			validMoves = append(validMoves, move)
			i++
		} else if piece.getColor() != color {
			move, _ := GetBoardNotation(i+1, j)
			validMoves = append(validMoves, move)
			break
		} else {
			break
		}
	}

	i = frow
	j = fcol
	for i-1 >= 0 { // up
		piece, _ := r.board.get(i-1, j)
		if piece == nil {
			move, _ := GetBoardNotation(i-1, j)
			validMoves = append(validMoves, move)
			i--
		} else if piece.getColor() != color {
			move, _ := GetBoardNotation(i-1, j)
			validMoves = append(validMoves, move)
			break
		} else {
			break
		}
	}
	i = frow
	j = fcol
	for j+1 < 8 { // right
		piece, _ := r.board.get(i, j+1)
		if piece == nil {
			move, _ := GetBoardNotation(i, j+1)
			validMoves = append(validMoves, move)
			j++
		} else if piece.getColor() != color {
			move, _ := GetBoardNotation(i, j+1)
			validMoves = append(validMoves, move)
			break
		} else {
			break
		}
	}
	i = frow
	j = fcol
	for j-1 >= 0 { // left
		piece, _ := r.board.get(i, j-1)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i, j-1)
			validMoves = append(validMoves, move)
			j--
		} else {
			break
		}
	}
	return validMoves
}
func (r Rook) getAllMoves(frow, fcol int) []string {
	validMoves := make([]string, 0)
	rookColor := r.getColor()

	i := frow
	j := fcol
	for i+1 < 8 { // down
		piece, _ := r.board.get(i+1, j)
		if piece == nil || piece.getColor() == rookColor {
			move, _ := GetBoardNotation(i+1, j)
			validMoves = append(validMoves, move)
			i++
		} else {
			break
		}
	}

	i = frow
	j = fcol
	for i-1 >= 0 { // up
		piece, _ := r.board.get(i-1, j)
		if piece == nil || piece.getColor() == rookColor {
			move, _ := GetBoardNotation(i-1, j)
			validMoves = append(validMoves, move)
			i--
		} else {
			break
		}
	}
	i = frow
	j = fcol
	for j+1 < 8 { // right
		piece, _ := r.board.get(i, j+1)
		if piece == nil || piece.getColor() == rookColor {
			move, _ := GetBoardNotation(i, j+1)
			validMoves = append(validMoves, move)
			j++
		} else {
			break
		}
	}
	i = frow
	j = fcol
	for j-1 >= 0 { // left
		piece, _ := r.board.get(i, j-1)
		if piece == nil || piece.getColor() == rookColor {
			move, _ := GetBoardNotation(i, j-1)
			validMoves = append(validMoves, move)
			j--
		} else {
			break
		}
	}
	return validMoves
}

func (r Rook) isCheck(color string, row, col int) bool {
	i := row
	j := col
	for i+1 < 8 { // down
		if HaveKing(r.board, color, i+1, j) == true {
			return true
		}
		i++
	}
	i = row
	j = col
	for i-1 >= 0 { // up
		if HaveKing(r.board, color, i-1, j) == true {
			return true
		}
		i--
	}
	i = row
	j = col
	for j+1 < 8 { // right
		if HaveKing(r.board, color, i, j+1) == true {
			return true
		}
		j++
	}
	i = row
	j = col
	for j-1 >= 0 { // left
		if HaveKing(r.board, color, i, j-1) == true {
			return true
		}
		j--
	}
	return false
}
func (r Rook) move(color string, frow, fcol, trow, tcol int) error {

	ownPiece := IsOwn(r.color, color)
	if ownPiece != true {
		return errors.New("Invalid move: Can't move other player piece.")
	}

	validMoves := r.getValidMoves(color, frow, fcol)
	destPlace, _ := GetBoardNotation(trow, tcol)

	if Contains(validMoves, destPlace) != true {
		return errors.New("Invalid move: Rook cant be moved there")
	}

	src, _ := r.board.get(frow, fcol)
	dest, _ := r.board.get(trow, tcol)
	if dest != nil && src.getColor() != dest.getColor() {
		fmt.Println("Capturing", dest.getName())
	}
	if r.isCheck(color, trow, tcol) == true {
		if IsCheckMate(r.board, r.getColor()) == true {
			fmt.Println("Check mate!")
		} else {
			fmt.Println("Check!")
		}
		return nil
	}
	r.board.set(frow, fcol, nil)
	r.board.set(trow, tcol, src)

	return nil
}
func (r Rook) String() string {
	return fmt.Sprintf(r.name)
}
func (r Rook) getColor() string {
	return r.color
}

func (r Rook) getName() string {
	return r.name
}
