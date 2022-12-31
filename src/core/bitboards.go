package core

import (
	"fmt"
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

type Square uint8

const (
	A1 Square = 0
	B1 Square = 1
	C1 Square = 2
	D1 Square = 3
	E1 Square = 4
	F1 Square = 5
	G1 Square = 6
	H1 Square = 7
	A2 Square = 8
	B2 Square = 9
	C2 Square = 10
	D2 Square = 11
	E2 Square = 12
	F2 Square = 13
	G2 Square = 14
	H2 Square = 15
	A3 Square = 16
	B3 Square = 17
	C3 Square = 18
	D3 Square = 19
	E3 Square = 20
	F3 Square = 21
	G3 Square = 22
	H3 Square = 23
	A4 Square = 24
	B4 Square = 25
	C4 Square = 26
	D4 Square = 27
	E4 Square = 28
	F4 Square = 29
	G4 Square = 30
	H4 Square = 31
	A5 Square = 32
	B5 Square = 33
	C5 Square = 34
	D5 Square = 35
	E5 Square = 36
	F5 Square = 37
	G5 Square = 38
	H5 Square = 39
	A6 Square = 40
	B6 Square = 41
	C6 Square = 42
	D6 Square = 43
	E6 Square = 44
	F6 Square = 45
	G6 Square = 46
	H6 Square = 47
	A7 Square = 48
	B7 Square = 49
	C7 Square = 50
	D7 Square = 51
	E7 Square = 52
	F7 Square = 53
	G7 Square = 54
	H7 Square = 55
	A8 Square = 56
	B8 Square = 57
	C8 Square = 58
	D8 Square = 59
	E8 Square = 60
	F8 Square = 61
	G8 Square = 62
	H8 Square = 63
)

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

type FormatingBitBoard interface {
	PrintBitBoard()
}

func (bitboard BitBoard) PrintBB() {

}

func (bitboard BitBoard) PrintPosition() {

	// Flipping was much easier because trying
	// to enumrate the squares in reverse order
	// was a pain. Mainly because of the Enum?
	bitboard = bitboard.FlipVertical()
	piece := ""
	for i := A1; i <= H8; i++ {
		if i%8 == 0 {
			fmt.Println("\n  +-------+-------+-------+-------+-------+-------+-------+-------+")
			fmt.Printf("%d |", 8-i/8)
		}
		// fmt.Print(i, j, Coords{j, i}.ToSquare().String())
		if bitboard.IsBitSet(Square(i)) {
			// if bitboard.IsBitSetByCoords(Square(j + i*8)) {
			piece = string(MakePiece(WHITE, BISHOP).ColoredSymbol())
		} else {
			piece = ""
		}
		fmt.Printf("   %-4s|", piece)

		// fmt.Print("|")
		// fmt.Println()

		// To print extra information needs fixing tho
		// if i%8 == 0 && i != 0 {
		// 	fmt.Print("\n|")
		// 	for k := 0; k < 8; k++ {
		// 		fmt.Printf("   \x1b[30m%-4s\x1b[0m", string(Coords{k, int(i / 8)}.ToSquare().String()))
		// 		fmt.Print("|")
		// 	}
		// }

	}
	fmt.Println("\n  +-------+-------+-------+-------+-------+-------+-------+-------+")
	fmt.Printf("     A        B       C       D       E        F       G       H\n\n")
}

// fmt.Println("\n+-------+-------+-------+-------+-------+-------+-------+-------+")
// piece := ""
// for i := 7; i >= 0; i-- {
// 	fmt.Println("\n+-------+-------+-------+-------+-------+-------+-------+-------+")
// 	for j := 7; j >= 0; j-- {
// 		if j == 7 {
// 			fmt.Print("|")
// 		}
// 		fmt.Print(i, j, Coords{j, i}.ToSquare().String())
// 		if bitboard.IsBitSet(Coords{j, i}.ToSquare()) {
// 			// if bitboard.IsBitSetByCoords(Square(j + i*8)) {
// 			piece = string(MakePiece(WHITE, BISHOP).ColoredSymbol())
// 		} else {
// 			piece = ""
// 		}
// 		fmt.Printf("   %-4s", piece)

// 		fmt.Print("|")
// 	}
// 	fmt.Println()

// 	for k := 0; k < 8; k++ {
// 		if k == 0 {
// 			fmt.Print("|")
// 		}
// 		fmt.Printf("   \x1b[30m%-4s\x1b[0m", string(Coords{k, i}.ToSquare().String()))

// 		fmt.Print("|")
// 	}
// }

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
