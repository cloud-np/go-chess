package core

import (
	"fmt"
	"unicode"
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
	ALL_PIECES    PieceType = 0
	PIECE_TYPE_NB PieceType = 8
)

func (pt PieceType) ToString() string {
	if pt == ALL_PIECES {
		return "ALL_PIECES || NO_PIECE_TYPE"
	}
	return [...]string{"NO_PIECE_TYPE", "PAWN", "KNIGHT", "BISHOP", "ROOK", "QUEEN", "KING", "PIECE_MASK", "PIECE_TYPE_NB"}[pt]
}

type Piece uint8

// We have to include the NO_PIECE
const PIECES_NUM = 13
const (
	NO_PIECE = 0
	W_PAWN   = 1
	W_KNIGHT = 2
	W_BISHOP = 3
	W_ROOK   = 4
	W_QUEEN  = 5
	W_KING   = 6
	B_PAWN   = 7
	B_KNIGHT = 8
	B_BISHOP = 9
	B_ROOK   = 10
	B_QUEEN  = 11
	B_KING   = 12
	PIECE_NB = 13
)

func (p Piece) ToString() string {
	return fmt.Sprint(p.Color().ToString(), " ", p.TypeOf().ToString())
}

func (p Piece) Char() rune {
	ptype := p.TypeOf()
	if ptype == NO_PIECE_TYPE {
		return ' '
	}

	color := p.Color()
	if color == WHITE {
		switch ptype {
		case PAWN:
			return 'P'
		case BISHOP:
			return 'B'
		case QUEEN:
			return 'Q'
		case KING:
			return 'K'
		case KNIGHT:
			return 'N'
		case ROOK:
			return 'R'
		}
	} else {
		switch ptype {
		case KNIGHT:
			return 'n'
		case BISHOP:
			return 'b'
		case ROOK:
			return 'r'
		case PAWN:
			return 'p'
		case QUEEN:
			return 'q'
		case KING:
			return 'k'
		}
	}
	panic("Invalid piece given: " + string(p))
}

func (p Piece) ColoredChar() string {
	color := 30
	if p.Color() == WHITE {
		color = 31
	}
	return fmt.Sprintf("\x1b[%dm%-4s\x1b[0m", color, string(p.Char()))
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
// without casting. But casting in Go is not slow
// since the value doesn't change only the type.
func MakePiece(c Color, pt PieceType) Piece {
	return Piece((uint8(c) << 3) + uint8(pt))
}

func MakePieceFromChar(ch rune) Piece {
	var pcolor Color
	var ptype PieceType
	if unicode.IsUpper(ch) {
		pcolor = WHITE
	} else {
		pcolor = BLACK
	}
	chl := unicode.ToLower(ch)
	switch chl {
	case 'p':
		ptype = PAWN
	case 'n':
		ptype = KNIGHT
	case 'b':
		ptype = BISHOP
	case 'r':
		ptype = ROOK
	case 'q':
		ptype = QUEEN
	case 'k':
		ptype = KING
	default:
		panic("Invalid piece given: " + string(ch))
	}
	return MakePiece(pcolor, ptype)
}
