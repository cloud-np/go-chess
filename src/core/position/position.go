package position

import (
	"fmt"
	"gochess/src/core"
	"gochess/src/core/moves"
	"math/bits"
)

type Position struct {
	// Bitboards
	byTypeBB  [core.PIECE_TYPE_NB]core.BitBoard
	byColorBB [core.COLOR_NB]core.BitBoard
	// byTypeBB  map[PieceType]*BitBoard
	// bycore.ColorBB map[core.Color]*BitBoard
	// Squares
	board      [64]core.Piece
	sideToMove core.Color
}

func (p *Position) GetByTypeBB() [core.PIECE_TYPE_NB]core.BitBoard {
	return p.byTypeBB
}

func (p *Position) GetByColorBB() [core.COLOR_NB]core.BitBoard {
	return p.byColorBB
}

func (p *Position) GetSideToMove() core.Color {
	return p.sideToMove
}

// Not needed yet
// func (p *Position) Count() uint {
// }
func (p *Position) GetPiecesSquares(color core.Color, pieceType core.PieceType) core.BitBoard {
	return p.byColorBB[color] & p.byTypeBB[pieceType]
}

// This has to implement LSB and we have to have only one bit set.
// Meaning we only have one PieceType in that BitBoard. Ideally we
// should make a function to count the number of bits in a BitBoard.
func (p *Position) GetSquare(color core.Color, pieceType core.PieceType) core.Square {
	pbb := p.GetPiecesSquares(color, pieceType)
	// core.Assert(pbb.Count() == 1, fmt.Sprintf("More than one piece of type %s and color %s", pieceType.ToString(), color.ToString()))
	// Why do I have to Flip it? Not sure
	return core.Square(uint8(bits.TrailingZeros64(uint64(pbb.FlipVertical()))))
}

// func (p *Position) GetSquareByPiece(piece Piece) Square {
// 	return p.board[sq]
// }

// func (p *Position) GetKingSquare(color core.Color) Square {
// 	return p.byTypeBB[KING]
// }

func (p *Position) GetPiece(sq core.Square) core.Piece {
	return p.board[sq]
}

func (p *Position) GetBoard() [64]core.Piece {
	return p.board
}

func (p *Position) PutPiece(piece core.Piece, sq core.Square) {
	// fmt.Print(string(piece.Char()))
	// fmt.Print(piece.Color() == core.WHITE)
	// fmt.Print(p.GetSquare(core.BLACK, core.ROOK))
	// Set based on piece type
	p.byTypeBB[core.ALL_PIECES].SetBit(sq)
	p.byTypeBB[piece.TypeOf()].SetBit(sq)
	p.byColorBB[piece.Color()].SetBit(sq)
	p.board[sq] = piece
}

// NOTE: XOR seems wrong here
func (p *Position) RemovePiece(sq core.Square) {
	piece := p.board[sq]
	// Unset the bits based on piece type and colors
	p.byTypeBB[core.ALL_PIECES].Xor(core.BitBoard(sq))
	p.byTypeBB[piece.TypeOf()].Xor(core.BitBoard(sq))
	p.byColorBB[piece.Color()].Xor(core.BitBoard(sq))
	p.board[sq] = core.NO_PIECE
}

func (p *Position) MovePiece(move moves.Move) {
	from := move.From()
	to := move.To()
	piece := p.board[from]
	fromTo := core.BitBoard(from | to)
	// Unset the bits based on piece type and colors
	p.byTypeBB[core.ALL_PIECES].Xor(fromTo)
	p.byTypeBB[core.PAWN].Xor(fromTo)
	p.byColorBB[piece.Color()].Xor(fromTo)
	p.board[from] = core.NO_PIECE
	p.board[to] = piece
}

func NewPosition() *Position {
	p := &Position{sideToMove: core.WHITE}
	// p.byTypeBB = make(map[PieceType]*BitBoard)
	// p.bycore.ColorBB = make(map[core.Color]*BitBoard)
	return p
}

func (p Position) IsMoveLegal(move moves.Move) bool {
	return true
}

func (p Position) PrintPosition(fancy bool) {
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
		if piece != core.NO_PIECE {
			if fancy {
				pstr = piece.ColoredSymbol()
			} else {
				pstr = piece.ColoredChar()
			}
		} else {
			pstr = ""
		}
		fmt.Printf("   %-4s|", pstr)

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
