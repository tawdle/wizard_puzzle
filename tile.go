package main

type Tile [8]int

func (t Tile) ToPosition(position, rotation int) int {
	i := (position - 2*rotation) % 8
	j := (t[i] + 2*rotation) % 8
	return j
}

var Tiles []Tile

func init() {
	Tiles = []Tile{
		Tile{1, 0, 7, 5, 6, 3, 4, 2}, // 0
		Tile{7, 3, 6, 1, 5, 4, 2, 0}, // 1
		Tile{4, 7, 6, 5, 0, 3, 2, 1}, // 2
		Tile{6, 3, 7, 1, 5, 4, 0, 2}, // 3
		Tile{7, 4, 6, 5, 1, 3, 2, 0}, // 4
		Tile{3, 4, 6, 0, 1, 7, 2, 5}, // 5
		Tile{5, 7, 4, 6, 2, 0, 3, 1}, // 6
		Tile{7, 3, 4, 1, 2, 6, 5, 0}, // 7
		Tile{7, 5, 4, 6, 2, 1, 3, 0}, // 8
	}
}
