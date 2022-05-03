package search

import (
	"math"
	"sort"

	"golang.org/x/exp/slices"

	"github.com/mbwk/gomaze/maze"
)

type Heuristic func(node maze.Coords) int

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func produceHeuristic(m *maze.Maze, start, end maze.Coords) Heuristic {
	return func(node maze.Coords) int {
		hx := (2 * abs(end.X-node.X)) - abs(node.X-start.X)
		hy := (2 * abs(end.Y-node.Y)) - abs(node.Y-start.Y)
		return hx + hy
	}
}

type NodeMap map[maze.Coords]int

func initializeNodeMap(m *maze.Maze) NodeMap {
	nodeMap := make(NodeMap)
	for x := 0; x < m.Width(); x++ {
		for y := 0; y < m.Height(); y++ {
			nodeMap[maze.Coords{x, y}] = math.MaxInt
		}
	}
	return nodeMap
}

func reconstructPath(cameFrom map[maze.Coords]maze.Coords, current maze.Coords) []maze.Coords {
	nodes := make([]maze.Coords, 0)
	for {
		nodes = append(nodes, current)
		back, present := cameFrom[current]
		if !present {
			break
		}
		current = back
	}
	return nodes
}

func distance(current, neighbour maze.Coords) int {
	return 1
}

func AStarSearch(m *maze.Maze, start, end maze.Coords) []maze.Coords {
	openSet := make([]maze.Coords, 0, 1)
	openSet = append(openSet, start)
	heuristic := produceHeuristic(m, start, end)

	cameFrom := make(map[maze.Coords]maze.Coords)

	gScore := initializeNodeMap(m)
	gScore[start] = 0

	fScore := initializeNodeMap(m)
	fScore[start] = heuristic(start)

	for len(openSet) > 0 {
		current := openSet[0]
		if current == end {
			return reconstructPath(cameFrom, current)
		}

		openSet = openSet[1:]
		for _, neighbour := range m.GetAccessibleNeighbours(current.X, current.Y) {
			tentative_gScore := gScore[current] + distance(current, neighbour)
			if tentative_gScore < gScore[neighbour] {
				cameFrom[neighbour] = current
				gScore[neighbour] = tentative_gScore
				fScore[neighbour] = tentative_gScore + heuristic(neighbour)
				if !slices.Contains(openSet, neighbour) {
					openSet = append(openSet, neighbour)
				}
			}
		}
		sort.Slice(openSet, func(i, j int) bool {
			return fScore[openSet[i]] < fScore[openSet[j]]
		})

	}
	return nil
}

func FindPath(m *maze.Maze, start, end maze.Coords) {
	path := AStarSearch(m, start, end)
	for i := 0; i < (len(path) - 1); i++ {
		m.PathBetween(path[i], path[i + 1])
	}
	m.Set(start.X, start.Y, maze.CellStart)
	m.Set(end.X, end.Y, maze.CellEnd)
}
