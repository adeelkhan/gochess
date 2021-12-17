package chess

import (
	"errors"
	"fmt"
)

type Knight struct {
	name  string
	color string
	board *Board
}

func (k Knight) move(color string, frow, fcol, trow, tcol int) error {
	ownPiece := IsOwn(k.color, color)
	if ownPiece != true {
		return errors.New("Invalid move: Can't move other player piece.")
	}

	// check validity of pawn moves
	validMoves := k.getValidMoves(color, frow, fcol)
	destPlace, _ := GetBoardNotation(trow, tcol)

	if Contains(validMoves, destPlace) != true {
		return errors.New("Invalid move: Knight cant be moved there")
	}

	src, _ := k.board.get(frow, fcol)
	dest, _ := k.board.get(trow, tcol)
	if dest != nil && src.getColor() != dest.getColor() {
		fmt.Println("Capturing", dest.getName())
	}
	if k.isCheck(color, trow, tcol) == true {
		if IsCheckMate(k.board, k.getColor()) == true {
			fmt.Println("Check mate!")
			return nil
		} else {
			fmt.Println("Check!")
		}
	}
	k.board.set(frow, fcol, nil)
	k.board.set(trow, tcol, src)

	return nil
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
	if i-2 >= 0 && j+1 < 8 {
		piece, _ := k.board.get(i-2, j+1)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i-2, j+1)
			validMoves = append(validMoves, move)
		}
	}
	//     k
	//   v-|
	i = frow
	j = fcol
	if i+2 < 8 && j-1 >= 0 {
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
	if i+2 < 8 && j+1 < 8 {
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
	if i+1 < 8 && j-2 >= 0 {
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
	if i-1 >= 0 && j-2 >= 0 {
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
	if i-1 >= 0 && j+2 < 8 {
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
	if i+1 < 0 && j+2 < 8 {
		piece, _ := k.board.get(i+1, j+2)
		if piece == nil || piece.getColor() != color {
			move, _ := GetBoardNotation(i+1, j+2)
			validMoves = append(validMoves, move)
		}
	}

	return validMoves
}
func (k Knight) getAllMoves(frow, fcol int) []string {
	validMoves := make([]string, 0)
	knightColor := k.getColor()

	i := frow
	j := fcol
	// <-|
	//   k
	if i-2 >= 0 && j-1 >= 0 {
		piece, _ := k.board.get(i-2, j-1)
		if piece == nil || piece.getColor() == knightColor {
			move, _ := GetBoardNotation(i-2, j-1)
			validMoves = append(validMoves, move)
		}
	}
	//   |->
	//   k
	i = frow
	j = fcol
	if i-2 >= 0 && j+1 < 8 {
		piece, _ := k.board.get(i-2, j+1)
		if piece == nil || piece.getColor() == knightColor {
			move, _ := GetBoardNotation(i-2, j+1)
			validMoves = append(validMoves, move)
		}
	}
	//   |-k
	//   v
	i = frow
	j = fcol
	if i+2 < 8 && j-1 >= 0 {
		piece, _ := k.board.get(i+2, j-1)
		if piece == nil || piece.getColor() == knightColor {
			move, _ := GetBoardNotation(i+2, j-1)
			validMoves = append(validMoves, move)
		}
	}
	//  k-|
	//    v
	i = frow
	j = fcol
	if i+2 < 8 && j+1 < 8 {
		piece, _ := k.board.get(i+2, j+1)
		if piece == nil || piece.getColor() == knightColor {
			move, _ := GetBoardNotation(i+2, j+1)
			validMoves = append(validMoves, move)
		}
	}
	//   |--|k
	//   v
	i = frow
	j = fcol
	if i+1 < 8 && j-2 >= 0 {
		piece, _ := k.board.get(i+1, j-2)
		if piece == nil || piece.getColor() == knightColor {
			move, _ := GetBoardNotation(i+1, j-2)
			validMoves = append(validMoves, move)
		}
	}
	//   ^
	//   |--k
	i = frow
	j = fcol
	if i-1 >= 0 && j-2 >= 0 {
		piece, _ := k.board.get(i-1, j-2)
		if piece == nil || piece.getColor() == knightColor {
			move, _ := GetBoardNotation(i-1, j-2)
			validMoves = append(validMoves, move)
		}
	}
	//   	  ^
	//   k--|
	i = frow
	j = fcol
	if i-1 >= 0 && j+2 < 8 {
		piece, _ := k.board.get(i-1, j+2)
		if piece == nil || piece.getColor() == knightColor {
			move, _ := GetBoardNotation(i-1, j+2)
			validMoves = append(validMoves, move)
		}
	}

	//  k--|
	// 	   v
	i = frow
	j = fcol
	if i+1 < 8 && j+2 < 8 {
		piece, _ := k.board.get(i+1, j+2)
		if piece == nil || piece.getColor() == knightColor {
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
		if HaveKing(k.board, color, i-2, j-1) == true {
			return true
		}
	}
	//   |->
	//   k
	i = row
	j = col
	if i-2 >= 0 && j+1 < 8 {
		if HaveKing(k.board, color, i-2, j+1) == true {
			return true
		}
	}
	//   |-k
	//   v
	i = row
	j = col
	if i+2 < 8 && j-1 >= 0 {
		if HaveKing(k.board, color, i+2, j-1) == true {
			return true
		}
	}
	//  k-|
	//    v
	i = row
	j = col
	if i+2 < 8 && j+1 < 8 {
		if HaveKing(k.board, color, i+2, j+1) == true {
			return true
		}
	}
	//   |--k
	//   v
	i = row
	j = col
	if i+1 < 8 && j-2 >= 0 {
		if HaveKing(k.board, color, i+1, j-2) == true {
			return true
		}
	}
	//   ^
	//   |--k
	i = row
	j = col
	if i-1 >= 0 && j-2 >= 0 {
		if HaveKing(k.board, color, i-1, j-2) == true {
			return true
		}
	}
	//   	  ^
	//   k--|
	i = row
	j = col
	if i-1 >= 0 && j+2 < 8 {
		if HaveKing(k.board, color, i-1, j+2) == true {
			return true
		}
	}
	//  k--|
	// 	   v
	i = row
	j = col
	if i+1 < 8 && j+2 < 8 {
		if HaveKing(k.board, color, i+1, j+2) == true {
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
