package core

import (
	"fmt"
	"math/bits"
)

type Color uint8

func (c Color) Other() Color {
	return Color(COLOR_NB - c)
}

const (
	WHITE    Color = 1
	BLACK    Color = 2
	COLOR_NB Color = 3
)

func (c Color) ToString() string {
	if c <= 0 || c > COLOR_NB {
		panic("INVALID COLOR: " + fmt.Sprint(c))
	}
	return [...]string{1: "WHITE", 2: "BLACK", 3: "COLOR_NB"}[c]
}

type Square uint8

// Little-Endian Rank-File Mapping
// https://www.chessprogramming.org/Square_Mapping_Considerations
const (
	A1 Square = 56
	B1 Square = 57
	C1 Square = 58
	D1 Square = 59
	E1 Square = 60
	F1 Square = 61
	G1 Square = 62
	H1 Square = 63
	A2 Square = 48
	B2 Square = 49
	C2 Square = 50
	D2 Square = 51
	E2 Square = 52
	F2 Square = 53
	G2 Square = 54
	H2 Square = 55
	A3 Square = 40
	B3 Square = 41
	C3 Square = 42
	D3 Square = 43
	E3 Square = 44
	F3 Square = 45
	G3 Square = 46
	H3 Square = 47
	A4 Square = 32
	B4 Square = 33
	C4 Square = 34
	D4 Square = 35
	E4 Square = 36
	F4 Square = 37
	G4 Square = 38
	H4 Square = 39
	A5 Square = 24
	B5 Square = 25
	C5 Square = 26
	D5 Square = 27
	E5 Square = 28
	F5 Square = 29
	G5 Square = 30
	H5 Square = 31
	A6 Square = 16
	B6 Square = 17
	C6 Square = 18
	D6 Square = 19
	E6 Square = 20
	F6 Square = 21
	G6 Square = 22
	H6 Square = 23
	A7 Square = 8
	B7 Square = 9
	C7 Square = 10
	D7 Square = 11
	E7 Square = 12
	F7 Square = 13
	G7 Square = 14
	H7 Square = 15
	A8 Square = 0
	B8 Square = 1
	C8 Square = 2
	D8 Square = 3
	E8 Square = 4
	F8 Square = 5
	G8 Square = 6
	H8 Square = 7
)

func (s Square) isOk() bool {
	return s >= A1 && s <= H8
}

func (s Square) String() string {
	return fmt.Sprintf("%c%c", 'a'+s%8, '1'+s/8)
}

// Returns whether the square is on the board
func isOk(s Square) bool {
	return s >= A1 && s <= H8
}

type BitBoard uint64

const (
	FILE_A BitBoard = 0x0101010101010101
	FILE_B BitBoard = FILE_A << 1
	FILE_C BitBoard = FILE_A << 2
	FILE_D BitBoard = FILE_A << 3
	FILE_E BitBoard = FILE_A << 4
	FILE_F BitBoard = FILE_A << 5
	FILE_G BitBoard = FILE_A << 6
	FILE_H BitBoard = FILE_A << 7

	RANK_1 BitBoard = 0x00000000000000FF
	RANK_2 BitBoard = RANK_1 << (8 * 1)
	RANK_3 BitBoard = RANK_1 << (8 * 2)
	RANK_4 BitBoard = RANK_1 << (8 * 3)
	RANK_5 BitBoard = RANK_1 << (8 * 4)
	RANK_6 BitBoard = RANK_1 << (8 * 5)
	RANK_7 BitBoard = RANK_1 << (8 * 6)
	RANK_8 BitBoard = RANK_1 << (8 * 7)

	QUEEN_SIDE    BitBoard = FILE_A | FILE_B | FILE_C | FILE_D
	CENTER_FILES  BitBoard = FILE_C | FILE_D | FILE_E | FILE_F
	KING_SIDE     BitBoard = FILE_E | FILE_F | FILE_G | FILE_H
	CENTER        BitBoard = (FILE_D | FILE_E) & (RANK_4 | RANK_5)
	A1_H8_DIAG    BitBoard = 0x8040201008040201
	H1_A8_DIAG    BitBoard = 0x0102040810204080
	LIGHT_SQUARES BitBoard = 0x55AA55AA55AA55AA
	DARK_SQUARES  BitBoard = 0xAA55AA55AA55AA55
	NOT_A_FILE    BitBoard = 0xfefefefefefefefe
	NOT_H_FILE    BitBoard = 0x7f7f7f7f7f7f7f7f
)

type Coords struct {
	x int
	y int
}

func (coords Coords) ToSquare() Square {
	return Square(coords.x + coords.y*8)
}

func NewCoords(x, y int) Coords {
	return Coords{x: x, y: y}
}

func (coords Coords) X() int {
	return coords.x
}

func (coords *Coords) SetX(x int) {
	coords.x = x
}

func (coords Coords) Y() int {
	return coords.y
}

func (coords *Coords) SetY(y int) {
	coords.y = y
}

// We need these because in Go, some bitwise operations are different.
// Specifically, the NOT operator is different and its the same symbol
// if you want to perfom a XOR operation so to avoid confusion
// and improve readability we use these functions.
func (b BitBoard) Union(other BitBoard) BitBoard {
	return b | other
}

func (b BitBoard) Not() BitBoard {
	return ^b
}

func (b BitBoard) RelativeComplement(other BitBoard) BitBoard {
	// other with out b
	return b.Not() & other
}

func (b BitBoard) Implication(other BitBoard) BitBoard {
	// Eveyrthing but b.
	return b.Not() | other
}

func (b BitBoard) Xor(other BitBoard) BitBoard {
	return b ^ other
}

// Not needed for now
type BitBoardShfiting interface {
	File()
	Rank()

	// 	northwest    north   northeast
	//   noWe         nort         noEa
	//           +7    +8    +9
	//               \  |  /
	//   west    -1 <-  0 -> +1    east
	//               /  |  \
	//           -9    -8    -7
	//   soWe         sout         soEa
	//   southwest    south   southeast

	// Verticals
	SouthOne()
	NorthOne()

	// Horizontals
	EastOne()
	WestOne()

	// Diagonals
	NoEaOne()
	SoEaOne()
	NoWeOne()
	SoWeOne()
}

// Verticals do not need check for overflows
func (b *BitBoard) SouthOne() {
	*b >>= 8
}

// Verticals do not need check for overflows
func (b *BitBoard) NorthOne() {
	*b <<= 8
}

func (b *BitBoard) EastOne() {
	*b = (*b & NOT_H_FILE) << 1
}

func (b *BitBoard) NoEaOne() {
	*b = (*b & NOT_H_FILE) << 9
}

func (b *BitBoard) SoEaOne() {
	*b = (*b & NOT_H_FILE) >> 7
}

func (b *BitBoard) WestOne() {
	*b = (*b & NOT_A_FILE) >> 1
}

func (b *BitBoard) SoWeOne() {
	*b = (*b & NOT_H_FILE) >> 9
}

func (b *BitBoard) NoWeOne() {
	*b = (*b & NOT_H_FILE) << 7
}

func (b BitBoard) Count() int {
	return bits.OnesCount64(uint64(b))
}

func (bitboard BitBoard) PrintBB() {
	bitboard = bitboard.FlipVertical()
	// piece := ""
	for i := A1; i <= H8; i++ {
		if i%8 == 0 {
			fmt.Println()
		}
		if bitboard.IsBitSet(Square(i)) {
			fmt.Print(" \x1b[31m1\x1b[0m ")
		} else {
			fmt.Print(" 0 ")
		}
	}
}

func (b BitBoard) FlipVertical() BitBoard {
	k1 := BitBoard(0x00FF00FF00FF00FF)
	k2 := BitBoard(0x0000FFFF0000FFFF)
	x := ((b >> 8) & k1) | ((b & k1) << 8)
	x = ((x >> 16) & k2) | ((x & k2) << 16)
	x = (x >> 32) | (x << 32)
	return x
}

func NewBitBoard() BitBoard {
	return BitBoard(0)
}

func (bitboard BitBoard) GetBit(square Square) BitBoard {
	// Needs the parenthesis for proper order of operations
	return bitboard & (1 << square)
}

func (bitboard BitBoard) IsBitSet(square Square) bool {
	return bitboard.GetBit(square) > 0
}

func (bitboard *BitBoard) SetBit(square Square) {
	*bitboard |= 1 << square
}

func (bitboard BitBoard) GetBitByCoords(coords Coords) BitBoard {
	// Needs the parenthesis for proper order of operations
	return bitboard & (1 << coords.ToSquare())
}

func (bitboard BitBoard) IsBitSetByCoords(coords Coords) bool {
	return bitboard.GetBitByCoords(coords) > 0
}

func (bitboard *BitBoard) SetBitByCoords(coords Coords) {
	*bitboard |= 1 << coords.ToSquare()
}
