package bitboards

import "fmt"

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
	a1 Square = 0
	b1 Square = 1
	c1 Square = 2
	d1 Square = 3
	e1 Square = 4
	f1 Square = 5
	g1 Square = 6
	h1 Square = 7
	a2 Square = 8
	b2 Square = 9
	c2 Square = 10
	d2 Square = 11
	e2 Square = 12
	f2 Square = 13
	g2 Square = 14
	h2 Square = 15
	a3 Square = 16
	b3 Square = 17
	c3 Square = 18
	d3 Square = 19
	e3 Square = 20
	f3 Square = 21
	g3 Square = 22
	h3 Square = 23
	a4 Square = 24
	b4 Square = 25
	c4 Square = 26
	d4 Square = 27
	e4 Square = 28
	f4 Square = 29
	g4 Square = 30
	h4 Square = 31
	a5 Square = 32
	b5 Square = 33
	c5 Square = 34
	d5 Square = 35
	e5 Square = 36
	f5 Square = 37
	g5 Square = 38
	h5 Square = 39
	a6 Square = 40
	b6 Square = 41
	c6 Square = 42
	d6 Square = 43
	e6 Square = 44
	f6 Square = 45
	g6 Square = 46
	h6 Square = 47
	a7 Square = 48
	b7 Square = 49
	c7 Square = 50
	d7 Square = 51
	e7 Square = 52
	f7 Square = 53
	g7 Square = 54
	h7 Square = 55
	a8 Square = 56
	b8 Square = 57
	c8 Square = 58
	d8 Square = 59
	e8 Square = 60
	f8 Square = 61
	g8 Square = 62
	h8 Square = 63
)

// Returns whether the square is on the board
func isOk(s Square) bool {
	return s >= a1 && s <= h8
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
)

type Coords struct {
	x int
	y int
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

func (bitboard BitBoard) Board() BitBoard {
	return bitboard
}

// Not needed for now
type BitBoardShfiting interface {
	File() BitBoard
	Rank() BitBoard

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
	SouthOne() BitBoard
	NorthOne() BitBoard

	// Horizontals
	EastOne() BitBoard
	WestOne() BitBoard

	// Diagonals
	NoEaOne() BitBoard
	SoEaOne() BitBoard
	NoWeOne() BitBoard
	SoWeOne() BitBoard
}

func (bitboard BitBoard) SouthOne() BitBoard {
	return bitboard >> 8
}

func (bitboard BitBoard) NorthOne() BitBoard {
	return bitboard >> 8
}

type FormatingBitBoard interface {
	PrintBitBoard()
}

func (bitboard BitBoard) PrintBitBoard() {
	for i := 0; i < 8; i++ {
		fmt.Println()
		for j := 0; j < 8; j++ {
			square := 0
			if bitboard.GetBit(NewCoords(i, j)) > 0 {
				square = 1
			}
			fmt.Print(square)
		}
	}
}

func (bitboard BitBoard) GetBit(coords Coords) BitBoard {
	// Needs the parenthesis for proper order of operations
	return bitboard & (1 << CoordsToIndex(coords))
}

func (bitboard BitBoard) SetBit(coords Coords) {
	bitboard |= 1 << CoordsToIndex(coords)
}
