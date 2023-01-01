package core

import "fmt"

type Position struct {
	// Bitboards
	byTypeBB  [PIECE_TYPE_NB]BitBoard
	byColorBB [COLOR_NB]BitBoard
	// byTypeBB  map[PieceType]*BitBoard
	// byColorBB map[Color]*BitBoard
	// Squares
	board [64]Piece
}

func (p Position) GetByTypeBB() [PIECE_TYPE_NB]BitBoard {
	return p.byTypeBB
}

func (p Position) GetColors() [COLOR_NB]BitBoard {
	return p.byColorBB
}

// func (p Position) GetByTypeBB() map[PieceType]*BitBoard {
// 	return p.byTypeBB
// }

// func (p Position) GetColors() map[Color]*BitBoard {
// 	return p.byColorBB
// }

func (p Position) GetBoard() [64]Piece {
	return p.board
}

func (p *Position) PutPiece(piece Piece, sq Square) {
	// Set based on piece type
	p.byTypeBB[ALL_PIECES].SetBit(sq)
	p.byTypeBB[piece.TypeOf()].SetBit(sq)
	p.byColorBB[piece.Color()].SetBit(sq)
	p.board[sq] = piece
}

// NOTE: XOR seems wrong here
func (p *Position) RemovePiece(sq Square) {
	piece := p.board[sq]
	// Unset the bits based on piece type and colors
	p.byTypeBB[ALL_PIECES].Xor(BitBoard(sq))
	p.byTypeBB[piece.TypeOf()].Xor(BitBoard(sq))
	p.byColorBB[piece.Color()].Xor(BitBoard(sq))
	p.board[sq] = NO_PIECE
}

func (p *Position) MovePiece(from Square, to Square) {
	piece := p.board[from]
	fromTo := BitBoard(from | to)
	// Unset the bits based on piece type and colors
	p.byTypeBB[ALL_PIECES].Xor(fromTo)
	p.byTypeBB[piece.TypeOf()].Xor(fromTo)
	p.byColorBB[piece.Color()].Xor(fromTo)
	p.board[from] = NO_PIECE
	p.board[to] = piece
}

func NewPosition() *Position {
	p := &Position{}
	// p.byTypeBB = make(map[PieceType]*BitBoard)
	// p.byColorBB = make(map[Color]*BitBoard)
	return p
}

func (p Position) PrintPosition() {

	// Flipping was much easier because trying
	// to enumrate the squares in reverse order
	// was a pain. Mainly because of the Enum?
	// p = bitboard.FlipVertical()
	pstr := ""
	for i, piece := range p.board {
		if i%8 == 0 {
			fmt.Println("\n  +-------+-------+-------+-------+-------+-------+-------+-------+")
			fmt.Printf("%d |", 8-i/8)
		}
		// fmt.Print(i, j, Coords{j, i}.ToSquare().String())
		if piece != NO_PIECE {
			// if bitboard.IsBitSetByCoords(Square(j + i*8)) {
			pstr = string(piece.ColoredSymbol())
		} else {
			pstr = ""
		}
		fmt.Printf("   %-4s|", pstr)

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
