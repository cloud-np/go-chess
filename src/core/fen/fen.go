package core

import (
	"gochess/src/core"
	"gochess/src/core/position"
	"strings"
)

type Fen string

func (fen Fen) CreatePosition() *position.Position {
	pos := position.NewPosition()
	rank := 7
	file := 0

	for _, ch := range fen {
		if strings.Contains("12345678", string(ch)) {
			file += int(ch - '0')
		} else if ch == '/' {
			rank--
			file = 0
		} else if ch == ' ' {
			break
		} else {
			piece := core.MakePieceFromChar(ch)
			if piece != core.NO_PIECE {
				square := core.Square(rank*8 + file)
				pos.PutPiece(piece, square)
				file++
			}
		}
	}
	return pos
}

// func (fen Fen) CreatePosition() *position.Position {
// 	pos := position.NewPosition()
// 	sq := 0
// 	for _, ch := range fen {
// 		if ch >= '1' && ch <= '8' {
// 			sq += int(ch - '0')
// 		} else if ch == '/' {
// 			continue
// 		} else if ch == ' ' {
// 			break
// 		} else {
// 			piece := core.MakePieceFromChar(ch)
// 			if piece != core.NO_PIECE {
// 				fmt.Print(core.Square(sq), " ")
// 				fmt.Printf("%c", ch)
// 				pos.PutPiece(piece, core.Square(sq))
// 				sq++
// 			}
// 		}
// 	}
// 	return pos
// }
