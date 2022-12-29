package moves

type Move uint16

const (
	MOVE_NONE Move = 0
	MOVE_NULL Move = 65
)

type MoveType uint16

const (
	NORMAL     MoveType = 0
	PROMOTION  MoveType = 1 << 14
	EN_PASSANT MoveType = 2 << 14
	CASTLING   MoveType = 3 << 14
)

type CastlingRights uint8

const (
	NO_CASTLING CastlingRights = 0
	WHITE_OO    CastlingRights = 1
	WHITE_OOO   CastlingRights = WHITE_OO << 1
	BLACK_OO    CastlingRights = WHITE_OO << 2
	BLACK_OOO   CastlingRights = WHITE_OO << 3

	KING_SIDE      CastlingRights = WHITE_OO | BLACK_OO
	QUEEN_SIDE     CastlingRights = WHITE_OOO | BLACK_OOO
	WHITE_CASTLING CastlingRights = WHITE_OO | WHITE_OOO
	BLACK_CASTLING CastlingRights = BLACK_OO | BLACK_OOO
	ANY_CASTLING   CastlingRights = WHITE_CASTLING | BLACK_CASTLING

	CASTLING_RIGHT_NB CastlingRights = 16
)
