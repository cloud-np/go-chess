package boards

import "fmt"

type BitBoard uint64

const (
	a1 = iota
	b1
	c1
	d1
	e1
	f1
	g1
	h1
	a2
	b2
	c2
	d2
	e2
	f2
	g2
	h2
	a3
	b3
	c3
	d3
	e3
	f3
	g3
	h3
	a4
	b4
	c4
	d4
	e4
	f4
	g4
	h4
	a5
	b5
	c5
	d5
	e5
	f5
	g5
	h5
	a6
	b6
	c6
	d6
	e6
	f6
	g6
	h6
	a7
	b7
	c7
	d7
	e7
	f7
	g7
	h7
	a8
	b8
	c8
	d8
	e8
	f8
	g8
	h8
)

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

type BitOps interface {
	SetBit(coords Coords)
	GetBit(coords Coords) uint64
}

// Not needed for now
type BitBoardMovement interface {
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
