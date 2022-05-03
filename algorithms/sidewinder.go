package algorithms

import (
	"github.com/mbwk/gomaze/maze"
	"github.com/mbwk/gomaze/random"
)

func Sidewinder(m *maze.Maze, rng *random.Rng) {
	runningSet := make([]maze.Coords, 0)

	for y := 0; y < m.Height(); y++ {
		for x := 0; x < m.Width(); x++ {
			runningSet = append(runningSet, maze.Coords{X: x, Y: y})
			possibleDirections := make([]maze.Direction, 0)
			if x+1 < m.Width() {
				possibleDirections = append(possibleDirections, maze.DirectionEast)
			}
			if y+1 < m.Height() {
				possibleDirections = append(possibleDirections, maze.DirectionSouth)
			}
			possibilities := len(possibleDirections)
			if possibilities > 0 {
				choice := 0
				if possibilities > 1 {
					choice = rng.IntN(possibilities)
				}
				chosenDirection := possibleDirections[choice]
				if chosenDirection == maze.DirectionSouth {
					chosenInRun := random.RandomSelection(rng, runningSet)
					m.Tunnel(chosenInRun.X, chosenInRun.Y, chosenDirection)
					runningSet = make([]maze.Coords, 0)
				} else {
					m.Tunnel(x, y, chosenDirection)
				}
			}
		}
	}
}