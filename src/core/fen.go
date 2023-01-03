package core

import (
	"strings"
)

type Fen string

func (fen Fen) CreatePosition() *Position {
	pos := NewPosition()
	sq := 0
	for _, ch := range fen {
		if strings.Contains("12345678", string(ch)) {
			sq += int(ch - '0')
		} else if ch == '/' {
			continue
		} else if ch == ' ' {
			break
		} else {
			piece := MakePieceFromChar(ch)
			if piece != NO_PIECE {
				// if piece.Color() == WHITE && piece.TypeOf() == KING {
				// 	fmt.Print("White King")
				// }
				pos.PutPiece(piece, Square(sq))
				sq++
			}
		}
	}
	return pos
}
