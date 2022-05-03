package algorithms

import (
	"github.com/mbwk/gomaze/maze"
	"github.com/mbwk/gomaze/random"
)

func BinaryTreePath(m *maze.Maze, rng *random.Rng) {
	for y := 0; y < m.Height(); y++ {
		for x := 0; x < m.Width(); x++ {
			possibleDirections := make([]maze.Direction, 0)
			if x+1 < m.Width() {
				possibleDirections = append(possibleDirections, maze.DirectionEast)
			}
			if y+1 < m.Height() {
				possibleDirections = append(possibleDirections, maze.DirectionSouth)
			}
			possibilities := len(possibleDirections)
			if possibilities > 0 {
				chosen := 0
				if possibilities > 1 {
					chosen = rng.IntN(possibilities + 2)
					if chosen > 1 {
						chosen = 1
					} else if chosen == 1 {
						chosen = 0
					}
				}
				m.Tunnel(x, y, possibleDirections[chosen])
			}
		}
	}
}