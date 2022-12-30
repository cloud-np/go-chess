package main

import (
	"gochess/src/core"
)

func main() {
	// bitboard := boards.A1_H8_DIAG
	// piece := pieces.QUEEN
	// color := bitboards.BLACK
	// whitePiece := pieces.MakePiece(color, piece)
	// fmt.Println(whitePiece.Color())
	// fmt.Println(whitePiece.TypeOf())
	// var piecesBB map[]bitboards.BitBoard ma
	// piecesBB[pieces.B_BISHOP] = bitboards.NewBitBoard()
	from := core.NewBitBoard()
	to := core.NewBitBoard()
	from.SetBit(core.A2)
	to.SetBit(core.A8)
	fromTo := from.Xor(to)
	fromTo.PrintBitBoard()
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
