package main

import (
	"gochess/src/core/bitboards"
)

func main() {
	// bitboard := boards.A1_H8_DIAG
	var bitboard bitboards.BitBoard
	bitboard = bitboards.RANK_1
	bitboard = bitboards.CENTER
	// bitboard.SetBit(boards.NewCoords(1, 2))
	// bitboard.SetBit(boards.NewCoords(2, 4))
	// bitboard.SetBit(boards.NewCoords(3, 2))
	bitboard.PrintBitBoard()
	// fmt.Println(boards.FILE_A)
}
