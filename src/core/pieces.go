package core

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

// We have to include the NO_PIECE
const PIECES_NUM = 13
const (
	NO_PIECE Piece = 0
	W_PAWN   Piece = 1
	W_KNIGHT Piece = 2
	W_BISHOP Piece = 3
	W_ROOK   Piece = 4
	W_QUEEN  Piece = 5
	W_KING   Piece = 6
	B_PAWN   Piece = 7
	B_KNIGHT Piece = 8
	B_BISHOP Piece = 9
	B_ROOK   Piece = 10
	B_QUEEN  Piece = 11
	B_KING   Piece = 12
	PIECE_NB Piece = 13
)

func (p Piece) Symbol() rune {
	color := p.Color()
	ptype := p.TypeOf()
	pieceCode := Piece(ptype) + Piece(color)
	switch pieceCode {
	case W_PAWN:
		return '♙'
	case B_PAWN:
		return '♟'
	case W_KNIGHT:
		return '♘'
	case B_KNIGHT:
		return '♞'
	case W_BISHOP:
		return '♗'
	case B_BISHOP:
		return '♝'
	case W_ROOK:
		return '♖'
	case B_ROOK:
		return '♜'
	case W_QUEEN:
		return '♕'
	case B_QUEEN:
		return '♛'
	case W_KING:
		return '♔'
	case B_KING:
		return '♚'
	}
	panic("Invalid piece given: " + string(p))
}

func (p Piece) TypeOf() PieceType {
	return PieceType(p & PIECE_MASK)
}

func (p Piece) Color() Color {
	return Color(p >> 3)
}

// NOTE: It would be ideal to find a way
// without casting. But castling in Go is not slow
// since the value doesn't change only the type.
func MakePiece(c Color, pt PieceType) Piece {
	return Piece((uint8(c) << 3) + uint8(pt))
}
