package main

import (
	"gochess/src/core"
	fen "gochess/src/core/fen"
	"gochess/src/core/moves"
)

func main() {
	fenStr := "rnbqk2r/ppp2ppp/3bpn2/3p4/3P4/3BPN2/PPP2PPP/RNBQK2R w KQkq - 0 1"
	// fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	pos := fen.Fen(fenStr).CreatePosition()
	move := moves.NewMove(core.A2, core.A4)
	// fmt.Print(move.From())
	// fmt.Print(move.To())
	pos.MovePiece(move)
	pos.PrintPosition(false)
}
