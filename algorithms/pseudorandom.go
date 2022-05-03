package algorithms

import (
	"github.com/mbwk/gomaze/maze"
	"github.com/mbwk/gomaze/random"
)

func PseudoRandomPath(m *maze.Maze, rng *random.Rng) {
	direction := maze.Direction(maze.DirectionNorth)
	if rng.IntN(2) == 1 {
		direction = maze.DirectionWest
	}
	m.Tunnel(0, 0, direction)
	for x, y := 0, 0; x < m.Width() && y < m.Height(); {
		headsOrTails := rng.IntN(2)
		if headsOrTails == 1 {
			direction = maze.DirectionEast
		} else {
			direction = maze.DirectionSouth
		}
		m.Tunnel(x, y, direction)
		delta := maze.DIRECTION_MAP[direction]
		x, y = x+delta.X, y+delta.Y
	}
}