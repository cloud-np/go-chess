package main

import (
	"gochess/src/core"
	fen "gochess/src/core/fen"
	"gochess/src/core/moves"
)

func main() {
	// bitboard := boards.A1_H8_DIAG
	// piece := pieces.QUEEN
	// color := bitboards.BLACK
	// whitePiece := pieces.MakePiece(color, piece)
	// fmt.Println(whitePiece.Color())
	// fmt.Println(whitePiece.TypeOf())
	// from := core.NewBitBoard()
	// to := core.NewBitBoard()
	// from.SetBit(core.A2)
	// to.SetBit(core.A8)
	// fromTo := from.Xor(to)
	// fromTo.PrintBB()
	fenStr := "rnbqk2r/ppp2ppp/3bpn2/3p4/3P4/3BPN2/PPP2PPP/RNBQK2R w KQkq - 0 1"
	// fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	pos := fen.Fen(fenStr).CreatePosition()
	pos.MovePiece(moves.NewMove(core.A2, core.A4))
	pos.PrintPosition()
	// pos.GetPieces()[core.B_BISHOP].PrintBB()
	// val.PrintBB()
	// fromTo.PrintPosition()
	// piecesBB[pieces.B_BISHOP].UpdateBB
	// piecesBB[pieces.B_BISHOP].PrintBitBoard()
	// bitboard = bitboards.RANK_1
	// bitboard = bitboards.CENTER

	// bitboard.SetBit(bitboards.G4)
	// fmt.Println(bitboard.IsBitSet(bitboards.G4))
	// fmt.Println(bitboard & bitboards.BitBoard(bitboards.G4))
	// bitboard.SetBitByCoords(bitboards.NewCoords(1, 2))
	// bitboard.PrintBitBoard()
	// fmt.Println()
	// bitboard.SetBitByCoords(bitboards.NewCoords(2, 4))
	// bitboard.SetBitByCoords(bitboards.NewCoords(3, 2))
	// fmt.Println(boards.FILE_A)
}
