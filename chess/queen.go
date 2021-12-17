package chess

import (
	"errors"
	"fmt"
)

type Queen struct {
	name  string
	color string
	board *Board
}

func (q Queen) getValidMoves(color string, frow, fcol int) []string {
	validMoves := make([]string, 0)

	i := frow
	j := fcol
	for i+1 < 8 { // down
		piece, _ := q.board.get(i+1, j)
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
		piece, _ := q.board.get(i-1, j)
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
		piece, _ := q.board.get(i, j+1)
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
		piece, _ := q.board.get(i, j-1)
		if piece == nil {
			move, _ := GetBoardNotation(i, j-1)
			validMoves = append(validMoves, move)
			j--
		} else if piece.getColor() != color {
			move, _ := GetBoardNotation(i, j-1)
			validMoves = append(validMoves, move)
			break
		} else {
			break
		}
	}
	i = frow
	j = fcol
	for i+1 < 8 && j+1 < 8 { // right diag down
		piece, _ := q.board.get(i+1, j+1)

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
		piece, _ := q.board.get(i-1, j+1)
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
		piece, _ := q.board.get(i+1, j-1)
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
		piece, _ := q.board.get(i-1, j-1)
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

func (q Queen) getAllMoves(frow, fcol int) []string {
	validMoves := make([]string, 0)
	queenColor := q.getColor()

	i := frow
	j := fcol
	for i+1 < 8 { // down
		piece, _ := q.board.get(i+1, j)
		if piece == nil || piece.getColor() == queenColor {
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
		piece, _ := q.board.get(i-1, j)
		if piece == nil || piece.getColor() == queenColor {
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
		piece, _ := q.board.get(i, j+1)
		if piece == nil || piece.getColor() == queenColor {
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
		piece, _ := q.board.get(i, j-1)
		if piece == nil || piece.getColor() == queenColor {
			move, _ := GetBoardNotation(i, j-1)
			validMoves = append(validMoves, move)
			j--
		} else {
			break
		}
	}
	i = frow
	j = fcol
	for i+1 < 8 && j+1 < 8 { // right diag down
		piece, _ := q.board.get(i+1, j+1)

		if piece == nil || piece.getColor() == queenColor {
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
		piece, _ := q.board.get(i-1, j+1)
		if piece == nil || piece.getColor() == queenColor {
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
		piece, _ := q.board.get(i+1, j-1)
		if piece == nil || piece.getColor() == queenColor {
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
		piece, _ := q.board.get(i-1, j-1)
		if piece == nil || piece.getColor() == queenColor {
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

func (q Queen) isCheck(color string, row, col int) bool {
	i := row
	j := col
	for i+1 < 8 { // down
		if HaveKing(q.board, color, i+1, j) == true {
			return true
		}
		piece, _ := q.board.get(i+1, j)
		if piece == nil { // keep searching
			i++
		} else {
			break
		}
	}
	i = row
	j = col
	for i-1 >= 0 { // up
		if HaveKing(q.board, color, i-1, j) == true {
			return true
		}
		piece, _ := q.board.get(i-1, j)
		if piece == nil { // keep searching
			i--
		} else {
			break
		}
	}
	i = row
	j = col
	for j+1 < 8 { // right
		if HaveKing(q.board, color, i, j+1) == true {
			return true
		}
		piece, _ := q.board.get(i, j+1)
		if piece == nil { // keep searching
			j++
		} else {
			break
		}
	}
	i = row
	j = col
	for j-1 >= 0 { // left
		if HaveKing(q.board, color, i, j-1) == true {
			return true
		}
		piece, _ := q.board.get(i, j-1)
		if piece == nil { // keep searching
			j--
		} else {
			break
		}
	}
	i = row
	j = col
	for i+1 < 8 && j+1 < 8 { // down diagonal right
		if HaveKing(q.board, color, i+1, j+1) == true {
			return true
		}
		piece, _ := q.board.get(i+1, j+1)
		if piece == nil { // keep searching
			i++
			j++
		} else {
			break
		}
	}
	i = row
	j = col
	for i-1 >= 0 && j+1 < 8 { // up diagonal right
		if HaveKing(q.board, color, i-1, j+1) == true {
			return true
		}
		piece, _ := q.board.get(i-1, j+1)
		if piece == nil { // keep searching
			i--
			j++
		} else {
			break
		}
	}
	i = row
	j = col
	for i-1 >= 0 && j+1 < 8 { // up diagonal left
		if HaveKing(q.board, color, i-1, j-1) == true {
			return true
		}
		piece, _ := q.board.get(i-1, j-1)
		if piece == nil { // keep searching
			i--
			j--
		} else {
			break
		}
	}
	i = row
	j = col
	for i+1 < 8 && j-1 >= 0 { // down diagonal left
		if HaveKing(q.board, color, i+1, j-1) == true {
			return true
		}
		piece, _ := q.board.get(i+1, j-1)
		if piece == nil { // keep searching
			i++
			j--
		} else {
			break
		}
	}

	return false
}
func (q Queen) move(color string, frow, fcol, trow, tcol int) error {
	ownPiece := IsOwn(q.color, color)
	if ownPiece != true {
		return errors.New("Invalid move: Can't move other player piece.")
	}

	validMoves := q.getValidMoves(color, frow, fcol)
	destPlace, _ := GetBoardNotation(trow, tcol)

	if Contains(validMoves, destPlace) != true {
		return errors.New("Invalid move: Queen cant be moved there")
	}

	src, _ := q.board.get(frow, fcol)
	dest, _ := q.board.get(trow, tcol)
	if dest != nil && src.getColor() != dest.getColor() {
		fmt.Println("Capturing", dest.getName())
	}
	if q.isCheck(color, trow, tcol) == true {
		if IsCheckMate(q.board, q.getColor()) == true {
			fmt.Println("Check mate!")
			return nil
		} else {
			fmt.Println("Check!")
		}
	}
	q.board.set(frow, fcol, nil)
	q.board.set(trow, tcol, src)

	return nil
}
func (q Queen) String() string {
	return fmt.Sprintf(q.name)
}
func (q Queen) getColor() string {
	return q.color
}
func (q Queen) getName() string {
	return q.name
}
