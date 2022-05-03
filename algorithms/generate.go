package algorithms

import (
	"github.com/mbwk/gomaze/maze"
	"github.com/mbwk/gomaze/random"
)

func GenerateMaze(m *maze.Maze, algorithm string, seed int64) int64 {
	rng := random.NewRng(seed)
	switch algorithm {
	case "btree":
		{
			BinaryTreePath(m, rng)
		}
	case "sidewinder":
		{
			Sidewinder(m, rng)
		}
	default:
		{
			PseudoRandomPath(m, rng)
		}
	}
	return rng.Seed
}
