package main

import (
	"fmt"
	"gochess/src/core"
	fen "gochess/src/core/fen"
	"gochess/src/core/moves"
)

func main() {
	// fenStr := "rbqk2r/ppp2ppp/3bpn2/3p4/3P4/3BPN2/PPP2PPP/RNBQK2R w KQkq - 0 1"
	fenStr := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	pos := fen.Fen(fenStr).CreatePosition()
	fmt.Printf("\n\n%064b", pos.GetByColorBB()[core.WHITE])
	fmt.Print("\n\n ", pos.GetByColorBB()[core.WHITE])
	fmt.Print("\n\n ", pos.GetByColorBB()[core.BLACK])
	fmt.Printf("\n\n%064b\n\n", pos.GetByColorBB()[core.BLACK])
	move := moves.NewMove(core.A2, core.A4)

	// fmt.Printf("\n %c", pos.GetPiece(core.E8).Char())
	// core.BitBoard(move).PrintBB()
	// pos.GetPiece(core.E1)
	// fmt.Print(move.From())
	// fmt.Print(move.To())
	pos.MovePiece(move)
	// pos.GetByColorBB()[core.BLACK].PrintBB()
	// pos.GetByColorBB()[core.WHITE].PrintBB()
	// fmt.Printf(pos.GetPiece(core.E1).ColoredChar())
	pos.PrintPosition(false)
	// position := fen.Fen(fenStr).CreatePosition()
	// position.PrintBoardSquares()

	// move := moves.NewMove(core.A2, core.A4)
	// position.MovePiece(move)
	// position.PrintBoardSquares()

}
