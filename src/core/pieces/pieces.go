package pieces

import (
	"gochess/src/core/bitboards"
)

type PieceType uint8

const (
	NO_PIECE_TYPE PieceType = 0
	PAWN          PieceType = 1
	KNIGHT        PieceType = 2
	BISHOP        PieceType = 3
	ROOK          PieceType = 4
	QUEEN         PieceType = 5
	KING          PieceType = 6
	PIECE_MASK              = 7
	// ALL_PIECES    PieceType = 0
	PIECE_TYPE_NB PieceType = 8
)

type Piece uint8

const (
	NO_PIECE Piece = 0
	W_PAWN   Piece = 1
	W_KNIGHT Piece = 2
	W_BISHOP Piece = 3
	W_ROOK   Piece = 4
	W_QUEEN  Piece = 5
	W_KING   Piece = 6
	B_PAWN   Piece = 9
	B_KNIGHT Piece = 9
	B_BISHOP Piece = 10
	B_ROOK   Piece = 11
	B_QUEEN  Piece = 12
	B_KING   Piece = 13
	PIECE_NB Piece = 16
)

type PieceFuncs interface {
	type_of() PieceType
}

func (p Piece) type_of() PieceType {
	return PieceType(p & PIECE_MASK)
}

// func (p Piece) Color() bitboards.Color {
// 	return bitboards.Color(p >> 3)
// }

// NOTE: It would be ideal to find a way
// without casting. But castling in Go is not slow
// since the value doesn't change only the type.
func makePiece(c bitboards.Color, pt PieceType) Piece {
	return Piece((uint8(c) << 3) + uint8(pt))
}
