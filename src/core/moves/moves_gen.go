package moves

import "gochess/src/core"

type GenType uint8

const (
	QUIET        GenType = 0
	CAPTURE      GenType = 1
	QUIET_CHECKS GenType = 2
	EVASIONS     GenType = 3
	NON_EVASIONS GenType = 4
	LEGAL        GenType = 5
)

func NewMove(from core.Square, to core.Square) Move {
	return Move((uint32(from) << 6) + uint32(to))
}

func (m Move) From() core.Square {
	return core.Square(m >> 6)
}

func (m Move) To() core.Square {
	return core.Square(m & 0x3F)
}

// func generatePawnMoves(pos *core.Position, moves *[]Move, genType GenType) *[]Move {

// }

// func generateMoves(pos *core.Position, moves *[]Move, genType GenType) *[]Move {
// 	cToPlay := pos.GetSideToMove()
// 	checks := genType == QUIET_CHECKS
// 	ksq := pos.GetKingSquare(cToPlay)

// 	// Generate all pseudo-legal moves
// 	if (genType == QUIET) || (genType == CAPTURE) || (genType == QUIET_CHECKS) {
// 		// Generate all pawn moves
// 		// Generate all knight moves
// 		// Generate all bishop moves
// 		// Generate all rook moves
// 		// Generate all queen moves
// 		// Generate all king moves
// 	}

// }

// struct Move {

// }
