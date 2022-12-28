package boards

func CoordsToIndex(coords Coords) int {
	return coords.x*8 + coords.y
}
