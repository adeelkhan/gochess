package chess

import (
	"errors"
	"fmt"
)

type King struct {
	name  string
	color string
	board *Board
}

func (k King) move(color string, frow, fcol, trow, tcol int) error {
	ownPiece := IsOwn(k.color, color)
	if ownPiece != true {
		return errors.New("Invalid move: Can't move other player piece.")
	}

	validMoves := k.getValidMoves(color, frow, fcol)
	destPlace, _ := GetBoardNotation(trow, tcol)

	if Contains(validMoves, destPlace) != true {
		return errors.New("Invalid move: King cant be moved there")
	}

	src, _ := k.board.get(frow, fcol)
	dest, _ := k.board.get(trow, tcol)
	if dest != nil && src.getColor() != dest.getColor() {
		fmt.Println("Capturing", dest.getName())
	}
	if k.isCheck(color, trow, tcol) == true {
		fmt.Println("Cant move there, there is a check!")
		return nil
	}
	k.board.set(frow, fcol, nil)
	k.board.set(trow, tcol, src)

	k.board.setKingLocation(k.getColor(), destPlace)
	return nil
}

func (k King) getValidMoves(color string, frow, fcol int) []string {
	validMoves := make([]string, 0)

	i := frow
	j := fcol
	if i+1 < 8 { // down
		piece, _ := k.board.get(i+1, j)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i+1, j)
			validMoves = append(validMoves, move)
		}
	}

	i = frow
	j = fcol
	if i-1 >= 0 { // up
		piece, _ := k.board.get(i-1, j)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i-1, j)
			validMoves = append(validMoves, move)
		}
	}
	i = frow
	j = fcol
	if j+1 < 8 { // right
		piece, _ := k.board.get(i, j+1)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i, j+1)
			validMoves = append(validMoves, move)
		}
	}
	i = frow
	j = fcol
	if j-1 >= 0 { // left
		piece, _ := k.board.get(i, j-1)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i, j-1)
			validMoves = append(validMoves, move)
		}
	}
	i = frow
	j = fcol
	if i+1 < 8 && j+1 < 8 { // right diag down
		piece, _ := k.board.get(i+1, j+1)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i+1, j+1)
			validMoves = append(validMoves, move)
		}
	}
	i = frow
	j = fcol
	if i-1 >= 0 && j+1 < 8 { // right diag up
		piece, _ := k.board.get(i-1, j+1)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i-1, j+1)
			validMoves = append(validMoves, move)
		}
	}
	i = frow
	j = fcol
	if i+1 < 8 && j-1 >= 0 { // left diag down
		piece, _ := k.board.get(i+1, j-1)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i+1, j-1)
			validMoves = append(validMoves, move)
		}
	}
	i = frow
	j = fcol
	if i-1 >= 0 && j-1 >= 0 { // left diag up
		piece, _ := k.board.get(i-1, j-1)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i-1, j-1)
			validMoves = append(validMoves, move)
		}
	}
	fmt.Println(validMoves)
	return validMoves
}
func (k King) getAllMoves(frow, fcol int) []string {
	validMoves := make([]string, 0)
	kingColor := k.getColor()

	i := frow
	j := fcol
	if i+1 < 8 { // down
		piece, _ := k.board.get(i+1, j)
		if piece == nil || piece.getColor() == kingColor {
			move, _ := GetBoardNotation(i+1, j)
			validMoves = append(validMoves, move)
		}
	}

	i = frow
	j = fcol
	if i-1 >= 0 { // up
		piece, _ := k.board.get(i-1, j)
		if piece == nil || piece.getColor() == kingColor {
			move, _ := GetBoardNotation(i-1, j)
			validMoves = append(validMoves, move)
		}
	}
	i = frow
	j = fcol
	if j+1 < 8 { // right
		piece, _ := k.board.get(i, j+1)
		if piece == nil || piece.getColor() == kingColor {
			move, _ := GetBoardNotation(i, j+1)
			validMoves = append(validMoves, move)
		}
	}
	i = frow
	j = fcol
	if j-1 >= 0 { // left
		piece, _ := k.board.get(i, j-1)
		if piece == nil || piece.getColor() == kingColor {
			move, _ := GetBoardNotation(i, j-1)
			validMoves = append(validMoves, move)
		}
	}
	i = frow
	j = fcol
	if i+1 < 8 && j+1 < 8 { // right diag down
		piece, _ := k.board.get(i+1, j+1)
		if piece == nil || piece.getColor() == kingColor {
			move, _ := GetBoardNotation(i+1, j+1)
			validMoves = append(validMoves, move)
		}
	}
	i = frow
	j = fcol
	if i-1 >= 0 && j+1 < 8 { // right diag up
		piece, _ := k.board.get(i-1, j+1)
		if piece == nil || piece.getColor() == kingColor {
			move, _ := GetBoardNotation(i-1, j+1)
			validMoves = append(validMoves, move)
		}
	}
	i = frow
	j = fcol
	if i+1 < 8 && j-1 >= 0 { // left diag down
		piece, _ := k.board.get(i+1, j-1)
		if piece == nil || piece.getColor() == kingColor {
			move, _ := GetBoardNotation(i+1, j-1)
			validMoves = append(validMoves, move)
		}
	}
	i = frow
	j = fcol
	if i-1 >= 0 && j-1 >= 0 { // left diag up
		piece, _ := k.board.get(i-1, j-1)
		if piece == nil || piece.getColor() == kingColor {
			move, _ := GetBoardNotation(i-1, j-1)
			validMoves = append(validMoves, move)
		}
	}
	return validMoves
}

func (k King) isCheck(color string, row, col int) bool {
	validMoves := make([]string, 0)
	keys := make(map[string]bool)

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			piece, _ := k.board.get(i, j)
			if piece != nil && piece.getColor() != color {
				fmt.Println(piece.getName())
				pieceMoves := piece.getAllMoves(i, j)
				fmt.Println(piece.getName(), " ", pieceMoves)
				for _, entry := range pieceMoves {
					if _, value := keys[entry]; !value {
						keys[entry] = true
						validMoves = append(validMoves, entry)
					}
				}
			}
		}
	}
	move, _ := GetBoardNotation(row, col)
	for _, value := range validMoves {
		if value == move {
			return true
		}
	}
	return false
}
func (k King) String() string {
	return fmt.Sprintf(k.name)
}

func (k King) getColor() string {
	return k.color
}

func (k King) getName() string {
	return k.name
}
