package core

import (
	"fmt"
	"strings"
)

type Fen string

func (fen Fen) CreatePosition() *Position {
	pos := NewPosition()
	// sq.SetFen(fen)
	sq := 0
	for _, ch := range fen {
		if strings.Contains("12345678", string(ch)) {
			sq += int(ch-'0') * 1
			continue
		} else if ch == '/' {
			sq += 2 * -8
			continue
		} else if ch == ' ' {
			break
		}
		piece := MakePieceFromChar(ch)
		if piece != NO_PIECE {
			sq++
			fmt.Println(piece, sq)
			pos.PutPiece(piece, Square(sq))
		}
		// pieceBB := pos.GetPieceBB(piece)
		// pieceBB.SetBit(i)
		// pos.SetPieceBB(piece, pieceBB)
	}
	return pos
}
