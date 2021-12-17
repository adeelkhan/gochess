package chess

import (
	"errors"
	"fmt"
	"strconv"
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
func IsOwn(playerColor, pieceColor string) bool {
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

func HaveKing(board *Board, color string, i, j int) bool {
	piece, _ := board.get(i, j)
	return piece != nil &&
		string(piece.getName()[0]) == "+" && color != piece.getColor()
}
func GetKingValidMoves(board *Board, color string, frow, fcol int) []string {
	validMoves := make([]string, 0)

	i := frow
	j := fcol
	if i+1 < 8 { // down
		piece, _ := board.get(i+1, j)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i+1, j)
			validMoves = append(validMoves, move)
		}
	}

	i = frow
	j = fcol
	if i-1 >= 0 { // up
		piece, _ := board.get(i-1, j)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i-1, j)
			validMoves = append(validMoves, move)
		}
	}
	i = frow
	j = fcol
	if j+1 < 8 { // right
		piece, _ := board.get(i, j+1)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i, j+1)
			validMoves = append(validMoves, move)
		}
	}
	i = frow
	j = fcol
	if j-1 >= 0 { // left
		piece, _ := board.get(i, j-1)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i, j-1)
			validMoves = append(validMoves, move)
		}
	}
	i = frow
	j = fcol
	if i+1 < 8 && j+1 < 8 { // right diag down
		piece, _ := board.get(i+1, j+1)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i+1, j+1)
			validMoves = append(validMoves, move)
		}
	}
	i = frow
	j = fcol
	if i-1 >= 0 && j+1 < 8 { // right diag up
		piece, _ := board.get(i-1, j+1)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i-1, j+1)
			validMoves = append(validMoves, move)
		}
	}
	i = frow
	j = fcol
	if i+1 < 8 && j-1 >= 0 { // left diag down
		piece, _ := board.get(i+1, j-1)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i+1, j-1)
			validMoves = append(validMoves, move)
		}
	}
	i = frow
	j = fcol
	if i-1 >= 0 && j-1 >= 0 { // left diag up
		piece, _ := board.get(i-1, j-1)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i-1, j-1)
			validMoves = append(validMoves, move)
		}
	}

	return validMoves

}

func IsCheckMate(board *Board, color string) bool {
	var kingLoc string
	if color == "white" {
		kingLoc = board.getKingLocation("black")
	} else {
		kingLoc = board.getKingLocation("white")
	}
	rowInd, _ := strconv.Atoi(string(kingLoc[1]))
	row, col := GetIndex(string(kingLoc[0]), rowInd)
	kingPossibleMoves := GetKingValidMoves(board, color, row, col)

	possibleMoves := make([]string, 0)
	keys := make(map[string]bool)

	// find all possible moves for colored player
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			piece, _ := board.get(i, j)
			if piece != nil && piece.getColor() == color {
				fmt.Println(piece.getName())
				pieceMoves := piece.getAllMoves(i, j)
				fmt.Println(piece.getName(), " ", pieceMoves)
				for _, entry := range pieceMoves {
					if _, value := keys[entry]; !value {
						keys[entry] = true
						possibleMoves = append(possibleMoves, entry)
					}
				}
			}
		}
	}

	totalMoves := len(kingPossibleMoves)
	i := 0
	for _, kingMove := range kingPossibleMoves {
		for _, opponentMove := range possibleMoves {
			if kingMove == opponentMove {
				i++
			}
		}
	}
	if i == totalMoves { // if all possible king moves are subset of opponent moves
		return true
	}
	return false
}
