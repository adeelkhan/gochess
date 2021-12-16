package chess

type Box struct {
	piece Piece
}

func (b *Box) setPiece(piece Piece) {
	b.piece = piece
}
func (b *Box) getPiece() Piece {
	return b.piece
}
