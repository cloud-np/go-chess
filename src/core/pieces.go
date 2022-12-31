package core

import "fmt"

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

func (p Piece) Char() string {
	color := p.Color()
	ptype := p.TypeOf()
	if color == WHITE {
		switch ptype {
		case PAWN:
			return "P"
		case BISHOP:
			return "B"
		case QUEEN:
			return "Q"
		case KING:
			return "K"
		case KNIGHT:
			return "Kn"
		case ROOK:
			return "R"
		}
	} else {
		switch ptype {
		case KNIGHT:
			return "kn"
		case BISHOP:
			return "b"
		case ROOK:
			return "r"
		case PAWN:
			return "p"
		case QUEEN:
			return "q"
		case KING:
			return "k"
		}
	}
	panic("Invalid piece given: " + string(p))
}

func (p Piece) ColoredChar() string {
	color := 30
	if p.Color() == WHITE {
		color = 31
	}
	return fmt.Sprintf("\x1b[%dm%-4s\x1b[0m", color, p.Char())
}

func (p Piece) ColoredSymbol() string {
	color := 30
	if p.Color() == WHITE {
		color = 31
	}
	return fmt.Sprintf("\x1b[%dm%-4s\x1b[0m", color, string(p.Symbol()))
}

func (p Piece) Symbol() rune {
	color := p.Color()
	ptype := p.TypeOf()
	if color == WHITE {
		switch ptype {
		case PAWN:
			return '♙'
		case BISHOP:
			return '♗'
		case QUEEN:
			return '♕'
		case KING:
			return '♔'
		case KNIGHT:
			return '♘'
		case ROOK:
			return '♖'
		}
	} else {
		switch ptype {
		case KNIGHT:
			return '♞'
		case BISHOP:
			return '♝'
		case ROOK:
			return '♜'
		case PAWN:
			return '♟'
		case QUEEN:
			return '♛'
		case KING:
			return '♚'
		}
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
