package core

import (
	"gochess/src/core"
	"gochess/src/core/position"
	"strings"
)

type Fen string

func (fen Fen) CreatePosition() *position.Position {
	pos := position.NewPosition()
	sq := 0
	for _, ch := range fen {
		if strings.Contains("12345678", string(ch)) {
			sq += int(ch - '0')
		} else if ch == '/' {
			continue
		} else if ch == ' ' {
			break
		} else {
			piece := core.MakePieceFromChar(ch)
			if piece != core.NO_PIECE {
				pos.PutPiece(piece, core.Square(sq))
				sq++
			}
		}
	}
	return pos
}
