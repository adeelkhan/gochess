package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var charMapping map[string]int = map[string]int{"a": 0, "b": 1, "c": 2, "d": 3, "e": 4, "f": 5, "g": 6, "h": 7}

////////// common functions
func GetBoardNotation(rowIndex int, colIndex int) (string, error) {
	if (rowIndex < 0 && rowIndex > 8) || (colIndex < 0 && colIndex > 8) {
		errorStr := fmt.Sprintf("Invalid board indexes (row=%d, col=%d)", rowIndex, colIndex)
		return "", errors.New(errorStr)
	}
	var key string
	for k, value := range charMapping {
		if value == colIndex {
			key = k
			break
		}
	}
	val := strconv.Itoa(8 - rowIndex)
	return key + val, nil
}

func GetIndex(colChar string, rowIndex int) (int, int) {
	var colInd int
	colInd, _ = charMapping[colChar]
	rowInd := 8 - rowIndex
	return rowInd, colInd
}
func isOwn(playerColor, pieceColor string) bool {
	return playerColor == pieceColor
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func haveKing(board *Board, color string, i, j int) bool {
	piece, _ := board.get(i, j)
	return piece != nil && string(piece.getName()[0]) == "+" && color != piece.getColor()
}

///////////////////////

// Piece structures
type Piece interface {
	move(color string, frow, fcol, trow, tcol int) error
	getAllMoves(color string, frow, fcol int) []string
	String() string
	getColor() string
	getName() string
}

type King struct {
	name  string
	color string
	board *Board
}

func (k King) move(color string, frow, fcol, trow, tcol int) error {
	ownPiece := isOwn(k.color, color)
	if ownPiece != true {
		return errors.New("Invalid move: Can't move other player piece.")
	}

	// check validity of pawn moves
	validMoves := k.getValidMoves(color, frow, fcol)
	destPlace, _ := GetBoardNotation(trow, tcol)

	if Contains(validMoves, destPlace) != true {
		return errors.New("Invalid move: Rook cant be moved there")
	}

	// everthing is ok now move
	src, _ := k.board.get(frow, fcol)
	dest, _ := k.board.get(trow, tcol)
	if dest != nil && src.getColor() != dest.getColor() { // capture piece
		fmt.Println("Capturing", dest.getName())
	}
	if k.isCheck(color, trow, tcol) == true {
		fmt.Println("Check!")
	}
	k.board.set(frow, fcol, nil)
	k.board.set(trow, tcol, src)

	fmt.Println(validMoves)
	return nil
}

func (k King) getValidMoves(color string, frow, fcol int) []string {
	validMoves := make([]string, 0)

	i := frow
	j := fcol
	for i+1 < 8 { // down
		piece, _ := k.board.get(i+1, j)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i+1, j)
			validMoves = append(validMoves, move)
		}
	}

	i = frow
	j = fcol
	for i-1 >= 0 { // up
		piece, _ := k.board.get(i-1, j)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i-1, j)
			validMoves = append(validMoves, move)
		}
	}
	i = frow
	j = fcol
	for j+1 < 8 { // right
		piece, _ := k.board.get(i, j+1)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i, j+1)
			validMoves = append(validMoves, move)
		}
	}
	i = frow
	j = fcol
	for j-1 >= 0 { // left
		piece, _ := k.board.get(i, j-1)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i, j-1)
			validMoves = append(validMoves, move)
		}
	}
	i = frow
	j = fcol
	for i+1 < 8 && j+1 < 8 { // right diag down
		piece, _ := k.board.get(i+1, j+1)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i+1, j+1)
			validMoves = append(validMoves, move)
		}
	}
	i = frow
	j = fcol
	for i-1 >= 0 && j+1 < 8 { // right diag up
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

	return validMoves
}
func (k King) getAllMoves(color string, frow, fcol int) []string {
	validMoves := make([]string, 0)

	i := frow
	j := fcol
	for i+1 < 8 { // down
		piece, _ := k.board.get(i+1, j)
		if piece == nil || piece.getColor() == color {
			move, _ := GetBoardNotation(i+1, j)
			validMoves = append(validMoves, move)
		}
	}

	i = frow
	j = fcol
	for i-1 >= 0 { // up
		piece, _ := k.board.get(i-1, j)
		if piece == nil || piece.getColor() == color {
			move, _ := GetBoardNotation(i-1, j)
			validMoves = append(validMoves, move)
		}
	}
	i = frow
	j = fcol
	for j+1 < 8 { // right
		piece, _ := k.board.get(i, j+1)
		if piece == nil || piece.getColor() == color {
			move, _ := GetBoardNotation(i, j+1)
			validMoves = append(validMoves, move)
		}
	}
	i = frow
	j = fcol
	for j-1 >= 0 { // left
		piece, _ := k.board.get(i, j-1)
		if piece == nil || piece.getColor() == color {
			move, _ := GetBoardNotation(i, j-1)
			validMoves = append(validMoves, move)
		}
	}
	i = frow
	j = fcol
	for i+1 < 8 && j+1 < 8 { // right diag down
		piece, _ := k.board.get(i+1, j+1)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i+1, j+1)
			validMoves = append(validMoves, move)
		}
	}
	i = frow
	j = fcol
	for i-1 >= 0 && j+1 < 8 { // right diag up
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
		if piece == nil || piece.getColor() == color {
			move, _ := GetBoardNotation(i+1, j-1)
			validMoves = append(validMoves, move)
		}
	}
	i = frow
	j = fcol
	if i-1 >= 0 && j-1 >= 0 { // left diag up
		piece, _ := k.board.get(i-1, j-1)
		if piece == nil || piece.getColor() == color {
			move, _ := GetBoardNotation(i-1, j-1)
			validMoves = append(validMoves, move)
		}
	}
	return validMoves
}

func (k King) isCheck(color string, row, col int) bool {
	validMoves := make([]string, 0)
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			piece, _ := k.board.get(i, j)
			if piece != nil && piece.getColor() != color {
				pieceMoves := piece.getAllMoves(color, i, j)
				validMoves = append(validMoves, pieceMoves...)
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

func (q Queen) getAllMoves(color string, frow, fcol int) []string {
	validMoves := make([]string, 0)

	i := frow
	j := fcol
	for i+1 < 8 { // down
		piece, _ := q.board.get(i+1, j)
		if piece == nil {
			move, _ := GetBoardNotation(i+1, j)
			validMoves = append(validMoves, move)
			i++
		} else if piece.getColor() == color {
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
		} else if piece.getColor() == color {
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
		} else if piece.getColor() == color {
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
		} else if piece.getColor() == color {
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
		} else if piece.getColor() == color {
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
		} else if piece.getColor() == color {
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
		} else if piece.getColor() == color {
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
		} else if piece.getColor() == color {
			move, _ := GetBoardNotation(i-1, j-1)
			validMoves = append(validMoves, move)
			break
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
		if haveKing(q.board, color, i+1, j) == true {
			return true
		}
		i++
	}
	i = row
	j = col
	for i-1 >= 0 { // up
		if haveKing(q.board, color, i-1, j) == true {
			return true
		}
		i--
	}
	i = row
	j = col
	for j+1 < 8 { // right
		if haveKing(q.board, color, i, j+1) == true {
			return true
		}
		j++
	}
	i = row
	j = col
	for j-1 >= 0 { // left
		if haveKing(q.board, color, i, j-1) == true {
			return true
		}
		j--
	}
	i = row
	j = col
	for i+1 < 8 && j+1 < 8 { // down diagonal right
		if haveKing(q.board, color, i+1, j+1) == true {
			return true
		}
		i++
		j++
	}
	i = row
	j = col
	for i-1 >= 0 && j+1 < 8 { // up diagonal right
		if haveKing(q.board, color, i-1, j+1) == true {
			return true
		}
		i--
		j++
	}
	i = row
	j = col
	for i-1 >= 0 && j+1 < 8 { // up diagonal left
		if haveKing(q.board, color, i-1, j+1) == true {
			return true
		}
		i--
		j++
	}
	i = row
	j = col
	for i+1 < 8 && j-1 >= 0 { // down diagonal left
		if haveKing(q.board, color, i+1, j-1) == true {
			return true
		}
		i++
		j--
	}

	return false
}
func (q Queen) move(color string, frow, fcol, trow, tcol int) error {
	ownPiece := isOwn(q.color, color)
	if ownPiece != true {
		return errors.New("Invalid move: Can't move other player piece.")
	}

	validMoves := q.getValidMoves(color, frow, fcol)
	destPlace, _ := GetBoardNotation(trow, tcol)

	if Contains(validMoves, destPlace) != true {
		return errors.New("Invalid move: Rook cant be moved there")
	}

	src, _ := q.board.get(frow, fcol)
	dest, _ := q.board.get(trow, tcol)
	if dest != nil && src.getColor() != dest.getColor() { // capture piece
		fmt.Println("capturing", dest.getName())
	}
	if q.isCheck(color, trow, tcol) == true {
		fmt.Println("Check!")
	}
	q.board.set(frow, fcol, nil)
	q.board.set(trow, tcol, src)

	fmt.Println(validMoves)
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

type Bishop struct {
	name  string
	color string
	board *Board
}

func (b Bishop) move(color string, frow, fcol, trow, tcol int) error {
	ownPiece := isOwn(b.color, color)
	if ownPiece != true {
		return errors.New("Invalid move: Can't move other player piece.")
	}

	// check validity of pawn moves
	validMoves := b.getValidMoves(color, frow, fcol)
	destPlace, _ := GetBoardNotation(trow, tcol)

	if Contains(validMoves, destPlace) != true {
		return errors.New("Invalid move: Rook cant be moved there")
	}

	// everthing is ok now move
	src, _ := b.board.get(frow, fcol)
	dest, _ := b.board.get(trow, tcol)
	if dest != nil && src.getColor() != dest.getColor() { // capture piece
		fmt.Println("capturing", dest.getName())
	}
	if b.isCheck(color, trow, tcol) == true {
		fmt.Println("Check!")
	}
	b.board.set(frow, fcol, nil)
	b.board.set(trow, tcol, src)

	fmt.Println(validMoves)
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

func (b Bishop) getAllMoves(color string, frow, fcol int) []string {
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
		} else if piece.getColor() == color {
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
		} else if piece.getColor() == color {
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
		} else if piece.getColor() == color {
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
		} else if piece.getColor() == color {
			move, _ := GetBoardNotation(i-1, j-1)
			validMoves = append(validMoves, move)
			break
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
		if haveKing(b.board, color, i+1, j+1) == true {
			return true
		}
		i++
		j++
	}
	i = row
	j = col
	for i-1 >= 0 && j+1 < 8 { // up diagonal right
		if haveKing(b.board, color, i-1, j+1) == true {
			return true
		}
		i--
		j++
	}
	i = row
	j = col
	for i-1 >= 0 && j+1 < 8 { // up diagonal left
		if haveKing(b.board, color, i-1, j+1) == true {
			return true
		}
		i--
		j++
	}
	i = row
	j = col
	for i+1 < 8 && j-1 >= 0 { // down diagonal left
		if haveKing(b.board, color, i+1, j-1) == true {
			return true
		}
		i++
		j--
	}
	return false
}

type Knight struct {
	name  string
	color string
	board *Board
}

func (k Knight) move(color string, frow, fcol, trow, tcol int) error {
	return errors.New("Invalid move")
}

func (k Knight) getValidMoves(color string, frow, fcol int) []string {
	validMoves := make([]string, 0)

	i := frow
	j := fcol
	// <-|
	//   k
	if i-2 >= 0 && j-1 >= 0 {
		piece, _ := k.board.get(i-2, j-1)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i-2, j-1)
			validMoves = append(validMoves, move)
		}
	}
	//   |->
	//   k
	i = frow
	j = fcol
	for i-2 >= 0 && j+1 < 8 {
		piece, _ := k.board.get(i-2, j+1)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i-2, j+1)
			validMoves = append(validMoves, move)
		}
	}
	//   |-k
	//   v
	i = frow
	j = fcol
	for i+2 < 8 && j-1 >= 0 {
		piece, _ := k.board.get(i+2, j-1)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i+2, j-1)
			validMoves = append(validMoves, move)
		}
	}
	//  k-|
	//    v
	i = frow
	j = fcol
	for i+2 >= 0 && j+1 >= 0 {
		piece, _ := k.board.get(i+2, j+1)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i+2, j+1)
			validMoves = append(validMoves, move)
		}
	}
	//   |--|k
	//   v
	i = frow
	j = fcol
	for i+1 >= 0 && j-2 >= 0 {
		piece, _ := k.board.get(i+1, j-2)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i+1, j-2)
			validMoves = append(validMoves, move)
		}
	}
	//   ^
	//   |--k
	i = frow
	j = fcol
	for i-1 >= 0 && j-2 >= 0 {
		piece, _ := k.board.get(i-1, j-2)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i-1, j-2)
			validMoves = append(validMoves, move)
		}
	}
	//   	^
	//   k--|
	i = frow
	j = fcol
	for i-1 >= 0 && j+2 >= 0 {
		piece, _ := k.board.get(i-1, j+2)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i-1, j+2)
			validMoves = append(validMoves, move)
		}
	}

	//  k--|
	// 	   v
	i = frow
	j = fcol
	for i+1 >= 0 && j+2 >= 0 {
		piece, _ := k.board.get(i+1, j+2)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i+1, j+2)
			validMoves = append(validMoves, move)
		}
	}

	return validMoves
}
func (k Knight) getAllMoves(color string, frow, fcol int) []string {
	validMoves := make([]string, 0)

	i := frow
	j := fcol
	// <-|
	//   k
	if i-2 >= 0 && j-1 >= 0 {
		piece, _ := k.board.get(i-2, j-1)
		if piece == nil || piece.getColor() == color {
			move, _ := GetBoardNotation(i-2, j-1)
			validMoves = append(validMoves, move)
		}
	}
	//   |->
	//   k
	i = frow
	j = fcol
	for i-2 >= 0 && j+1 < 8 {
		piece, _ := k.board.get(i-2, j+1)
		if piece == nil || piece.getColor() == color {
			move, _ := GetBoardNotation(i-2, j+1)
			validMoves = append(validMoves, move)
		}
	}
	//   |-k
	//   v
	i = frow
	j = fcol
	for i+2 < 8 && j-1 >= 0 {
		piece, _ := k.board.get(i+2, j-1)
		if piece == nil || piece.getColor() == color {
			move, _ := GetBoardNotation(i+2, j-1)
			validMoves = append(validMoves, move)
		}
	}
	//  k-|
	//    v
	i = frow
	j = fcol
	for i+2 >= 0 && j+1 >= 0 {
		piece, _ := k.board.get(i+2, j+1)
		if piece == nil || piece.getColor() == color {
			move, _ := GetBoardNotation(i+2, j+1)
			validMoves = append(validMoves, move)
		}
	}
	//   |--|k
	//   v
	i = frow
	j = fcol
	for i+1 >= 0 && j-2 >= 0 {
		piece, _ := k.board.get(i+1, j-2)
		if piece == nil || piece.getColor() == color {
			move, _ := GetBoardNotation(i+1, j-2)
			validMoves = append(validMoves, move)
		}
	}
	//   ^
	//   |--k
	i = frow
	j = fcol
	for i-1 >= 0 && j-2 >= 0 {
		piece, _ := k.board.get(i-1, j-2)
		if piece == nil || piece.getColor() == color {
			move, _ := GetBoardNotation(i-1, j-2)
			validMoves = append(validMoves, move)
		}
	}
	//   	^
	//   k--|
	i = frow
	j = fcol
	for i-1 >= 0 && j+2 >= 0 {
		piece, _ := k.board.get(i-1, j+2)
		if piece == nil || piece.getColor() == color {
			move, _ := GetBoardNotation(i-1, j+2)
			validMoves = append(validMoves, move)
		}
	}

	//  k--|
	// 	   v
	i = frow
	j = fcol
	for i+1 >= 0 && j+2 >= 0 {
		piece, _ := k.board.get(i+1, j+2)
		if piece == nil || piece.getColor() == color {
			move, _ := GetBoardNotation(i+1, j+2)
			validMoves = append(validMoves, move)
		}
	}

	return validMoves
}

func (k Knight) isCheck(color string, row, col int) bool {
	i := row
	j := col
	// <-|
	//   k
	if i-2 >= 0 && j-1 >= 0 {
		if haveKing(k.board, color, i-2, j-1) == true {
			return true
		}
	}
	//   |->
	//   k
	i = row
	j = col
	for i-2 >= 0 && j+1 < 8 {
		if haveKing(k.board, color, i-2, j+1) == true {
			return true
		}
	}
	//   |-k
	//   v
	i = row
	j = col
	for i+2 < 8 && j-1 >= 0 {
		if haveKing(k.board, color, i+2, j-1) == true {
			return true
		}
	}
	//  k-|
	//    v
	i = row
	j = col
	for i+2 >= 0 && j+1 >= 0 {
		if haveKing(k.board, color, i+2, j+1) == true {
			return true
		}
	}
	//   |--|k
	//   v
	i = row
	j = col
	for i+1 >= 0 && j-2 >= 0 {
		if haveKing(k.board, color, i+1, j-2) == true {
			return true
		}
	}
	//   ^
	//   |--k
	i = row
	j = col
	for i-1 >= 0 && j-2 >= 0 {
		if haveKing(k.board, color, i-1, j-2) == true {
			return true
		}
	}
	//   	^
	//   k--|
	i = row
	j = col
	for i-1 >= 0 && j+2 >= 0 {
		if haveKing(k.board, color, i-1, j+2) == true {
			return true
		}
	}
	//  k--|
	// 	   v
	i = row
	j = col
	for i+1 >= 0 && j+2 >= 0 {
		if haveKing(k.board, color, i+1, j+2) == true {
			return true
		}
	}
	return false
}
func (k Knight) String() string {
	return fmt.Sprintf(k.name)
}
func (k Knight) getColor() string {
	return k.color
}
func (k Knight) getName() string {
	return k.name
}

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
func (r Rook) getAllMoves(color string, frow, fcol int) []string {
	validMoves := make([]string, 0)

	i := frow
	j := fcol
	for i+1 < 8 { // down
		piece, _ := r.board.get(i+1, j)
		if piece == nil {
			move, _ := GetBoardNotation(i+1, j)
			validMoves = append(validMoves, move)
			i++
		} else if piece.getColor() == color {
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
		} else if piece.getColor() == color {
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
		} else if piece.getColor() == color {
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
		if piece == nil || piece.getColor() == color {
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
		if haveKing(r.board, color, i+1, j) == true {
			return true
		}
		i++
	}
	i = row
	j = col
	for i-1 >= 0 { // up
		if haveKing(r.board, color, i-1, j) == true {
			return true
		}
		i--
	}
	i = row
	j = col
	for j+1 < 8 { // right
		if haveKing(r.board, color, i, j+1) == true {
			return true
		}
		j++
	}
	i = row
	j = col
	for j-1 >= 0 { // left
		if haveKing(r.board, color, i, j-1) == true {
			return true
		}
		j--
	}
	return false
}
func (r Rook) move(color string, frow, fcol, trow, tcol int) error {

	ownPiece := isOwn(r.color, color)
	if ownPiece != true {
		return errors.New("Invalid move: Can't move other player piece.")
	}

	// check validity of pawn moves
	validMoves := r.getValidMoves(color, frow, fcol)
	destPlace, _ := GetBoardNotation(trow, tcol)

	if Contains(validMoves, destPlace) != true {
		return errors.New("Invalid move: Rook cant be moved there")
	}

	// everthing is ok now move
	src, _ := r.board.get(frow, fcol)
	dest, _ := r.board.get(trow, tcol)
	if dest != nil && src.getColor() != dest.getColor() { // capture piece
		fmt.Println("capturing", dest.getName())
	}
	if r.isCheck(color, trow, tcol) == true {
		fmt.Println("Check!")
	}
	r.board.set(frow, fcol, nil)
	r.board.set(trow, tcol, src)

	fmt.Println(validMoves)
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
func (p Pawn) getAllMoves(color string, frow, fcol int) []string {
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
			if piece != nil && piece.getColor() == color {
				move, _ := GetBoardNotation(i+1, j+1)
				validMoves = append(validMoves, move)
			}
		}
		if i+1 >= 0 && j-1 >= 0 { // left diag down
			piece, _ := p.board.get(i+1, j-1)
			if piece != nil && piece.getColor() == color {
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
			if piece != nil && piece.getColor() == color {
				move, _ := GetBoardNotation(i-1, j-1)
				validMoves = append(validMoves, move)
			}
		}
		if i-1 >= 0 && j+1 >= 0 { // right diag up
			piece, _ := p.board.get(i+1, j-1)
			if piece != nil && piece.getColor() == color {
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
			if haveKing(p.board, color, i+1, j+1) == true { // right diag down
				return true
			}
		}
		if i+1 >= 0 && j-1 >= 0 {
			if haveKing(p.board, color, i+1, j-1) == true { // left diag down
				return true
			}
		}
	} else {
		if i-1 < 8 && j-1 < 8 {
			if haveKing(p.board, color, i-1, j-1) == true { // left diag up
				return true
			}
		}
		if i-1 >= 0 && j+1 >= 0 {
			if haveKing(p.board, color, i-1, j+1) == true { // right diag up
				return true
			}
		}
	}
	return false
}
func (p Pawn) move(color string, frow, fcol, trow, tcol int) error {
	ownPiece := isOwn(p.color, color)
	if ownPiece != true {
		return errors.New("Invalid move: Can't move other player piece.")
	}

	// check validity of pawn moves
	validMoves := p.getValidMoves(color, frow, fcol)
	destPlace, _ := GetBoardNotation(trow, tcol)

	if Contains(validMoves, destPlace) != true {
		return errors.New("Invalid move: Pawn cant be moved there")
	}

	// everthing is ok now move
	src, _ := p.board.get(frow, fcol)
	dest, _ := p.board.get(trow, tcol)
	if dest != nil && src.getColor() != dest.getColor() { // capture piece
		fmt.Println("capturing", dest.getName())
	}
	if p.isCheck(color, trow, tcol) == true {
		fmt.Println("Check!")
	}
	p.board.set(frow, fcol, nil)
	p.board.set(trow, tcol, src)

	fmt.Println(validMoves)
	return nil
}
func (p Pawn) String() string {
	return fmt.Sprintf(p.name)
}

// Game structures
type Board struct {
	board [8][8]Box
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
		err = piece.move(p.color, frow, fcol, trow, tcol)
	}
	return err
}

func (b *Board) parseMove(move string) (int, int, int, int, error) {
	playerMove := strings.Split(move, "->")

	fmt.Println(playerMove)
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
	i, j = GetIndex("d", 7)
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

type Box struct {
	piece Piece
}

func (b *Box) setPiece(piece Piece) {
	b.piece = piece
}
func (b *Box) getPiece() Piece {
	return b.piece
}

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

func (g *Game) start() {
	turn := 0
	g.printBoard()
	for {
		if turn == 0 {
			turn = 0
			move := g.player1.getMove(g.Board)
			err := g.Board.move(g.player1, move)

			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			//break

		} else {
			turn = 0
			move := g.player2.getMove(g.Board)
			err := g.Board.move(g.player1, move)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
		}
		g.printBoard()
	}

}

type Player struct {
	name  string
	color string
}

func (p *Player) getMove(board *Board) string {
	fmt.Printf("Player: %s ", p.name)
	var move string
	fmt.Scanf("%s", &move)
	return move
}
func main() {
	// setup player
	p1 := Player{"adeel", "white"}
	p2 := Player{"wahaj", "black"}

	g := Game{}
	g.Initialize(&p1, &p2)
	//g.printBoard()
	g.start()
}
