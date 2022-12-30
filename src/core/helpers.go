package core

func CoordsToIndex(coords Coords) int {
	return coords.x*8 + coords.y
}
