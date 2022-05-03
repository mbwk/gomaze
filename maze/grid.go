package maze

import "strings"

type Cell int

func MazeToGrid(coord int) int {
	return (2 * coord) + 1
}

func GridToMaze(coord int) int {
	if coord == 0 {
		return 0
	}
	return (coord - 1) / 2
}

const (
	CellBlank = 0
	CellWall  = 1
	CellPath  = 2
	CellStart = 3
	CellEnd   = 4
)

type Pair2i struct {
	X int
	Y int
}

type Dimensions Pair2i

type Coords Pair2i

type Grid struct {
	D     Dimensions
	Cells []Cell
}

func (g *Grid) Deref(X, Y int) *Cell {
	return &g.Cells[(Y*g.D.X)+X]
}

func (g *Grid) GridSet(X, Y int, c Cell) {
	g.Cells[(Y*g.D.X)+X] = c
}

func (g *Grid) Width() int {
	return g.D.X
}

func (g *Grid) Height() int {
	return g.D.Y
}

func PrintGrid(g *Grid) string {
	lines := make([]string, 0, g.Height())
	line := make([]rune, g.Width())
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			cell := g.Deref(x, y)
			switch *cell {
			case CellBlank:
				{
					line[x] = '⬜'
				}
			case CellWall:
				{
					line[x] = '⬛'
				}
			case CellPath:
				{
					line[x] = '🟩'
				}
			case CellStart:
				{
					line[x] = '🏡'
				}
			case CellEnd:
				{
					line[x] = '🏁'
				}
			default:
				{
					line[x] = '❓'
				}
			}
		}
		lines = append(lines, string(line))
	}
	x := strings.Join(lines, "\n")
	return x
}
