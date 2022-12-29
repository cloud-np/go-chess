package main

import (
	"fmt"
	"gochess/src/core/bitboards"
	"gochess/src/core/pieces"
)

func main() {
	// bitboard := boards.A1_H8_DIAG
	// var bitboard bitboards.BitBoard
	piece := pieces.QUEEN
	color := bitboards.BLACK
	whitePiece := pieces.MakePiece(color, piece)
	fmt.Println(whitePiece.Color())
	fmt.Println(whitePiece.TypeOf())
	// bitboard = bitboards.RANK_1
	// bitboard = bitboards.CENTER
	// bitboard.SetBit(boards.NewCoords(1, 2))
	// bitboard.SetBit(boards.NewCoords(2, 4))
	// bitboard.SetBit(boards.NewCoords(3, 2))
	// bitboard.PrintBitBoard()
	// fmt.Println(boards.FILE_A)
}
