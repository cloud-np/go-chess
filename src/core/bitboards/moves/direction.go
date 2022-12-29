package moves

type Direction int8

const (
	NORTH Direction = 8
	EAST  Direction = 1
	SOUTH Direction = -NORTH
	WEST  Direction = -EAST

	NORTH_EAST Direction = NORTH + EAST
	SOUTH_EAST Direction = SOUTH + EAST
	SOUTH_WEST Direction = SOUTH + WEST
	NORTH_WEST Direction = NORTH + WEST
)
