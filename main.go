package main

import (
	"gochess/src/core/bitboards"
)

func main() {
	// bitboard := boards.A1_H8_DIAG
	// piece := pieces.QUEEN
	// color := bitboards.BLACK
	// whitePiece := pieces.MakePiece(color, piece)
	// fmt.Println(whitePiece.Color())
	// fmt.Println(whitePiece.TypeOf())
	var bitboard bitboards.BitBoard
	// bitboard = bitboards.RANK_1
	// bitboard = bitboards.CENTER
	bitboard.SetBit(bitboards.NewCoords(1, 2))
	// bitboard.SetBit(bitboards.NewCoords(2, 4))
	// bitboard.SetBit(bitboards.NewCoords(3, 2))
	bitboard.PrintBitBoard()
	// fmt.Println(boards.FILE_A)
}
