package maze

type Maze struct {
	D Dimensions
	G Grid
}

type Direction uint64

const (
	DirectionNone  = 0
	DirectionNorth = 1
	DirectionEast  = 2
	DirectionSouth = 3
	DirectionWest  = 4
)

var DIRECTION_MAP map[Direction]Dimensions

func init() {
	DIRECTION_MAP = make(map[Direction]Dimensions)
	DIRECTION_MAP[DirectionNorth] = Dimensions{X: 0, Y: -1}
	DIRECTION_MAP[DirectionEast] = Dimensions{X: 1, Y: 0}
	DIRECTION_MAP[DirectionSouth] = Dimensions{X: 0, Y: 1}
	DIRECTION_MAP[DirectionWest] = Dimensions{X: -1, Y: 0}
}

func ApplyMapDelta(x, y int, direction Direction) Coords {
	delta := DIRECTION_MAP[direction]
	xDir, yDir := delta.X, delta.Y
	return Coords{x + xDir, y + yDir}
}

func ApplyGridDelta(x, y int, direction Direction) (int, int) {
	delta := DIRECTION_MAP[direction]
	xDir, yDir := delta.X, delta.Y
	return MazeToGrid(x)+xDir, MazeToGrid(y)+yDir
}

func (m *Maze) Tunnel(x, y int, direction Direction) {
	newX, newY := ApplyGridDelta(x, y, direction)
	m.G.GridSet(newX, newY, CellBlank)
}

func (m *Maze) PathTo(x, y int, direction Direction) {
	newX, newY := ApplyGridDelta(x, y, direction)
	m.Set(x, y, CellPath)
	m.G.GridSet(newX, newY, CellPath)	
}

func (m *Maze) PathBetween(first, second Coords) {
	m2g := MazeToGrid
	g1x, g1y, g2x, g2y := m2g(first.X), m2g(first.Y), m2g(second.X), m2g(second.Y)
	betweenX := g2x - ((g2x - g1x) / 2)
	betweenY := g2y - ((g2y - g1y) / 2)
	m.G.GridSet(betweenX, betweenY, CellPath)
	m.G.GridSet(g2x, g2y, CellPath)
}

func (m *Maze) Set(x, y int, c Cell) {
	m.G.GridSet(MazeToGrid(x), MazeToGrid(y), c)
}

func (m *Maze) At(x, y int) *Cell {
	return m.G.Deref(MazeToGrid(x), MazeToGrid(y))
}

func (m *Maze) Width() int {
	return m.D.X
}

func (m *Maze) Height() int {
	return m.D.Y
}

func (m *Maze) GetPossibleDirections(x, y int) []Direction {
	possibilities := make([]Direction, 0)
	for _, d := range []Direction{DirectionNorth, DirectionEast, DirectionSouth, DirectionWest} {
		newX, newY := ApplyGridDelta(x, y, d)
		gridCell := m.G.Deref(newX, newY)
		if *gridCell == CellBlank {
			possibilities = append(possibilities, d)
		}
	}
	return possibilities
}

func (m *Maze) GetAccessibleNeighbours(x, y int) []Coords {
	neighbours := make([]Coords, 0)
	for _, direction := range m.GetPossibleDirections(x, y) {
		neighbours = append(neighbours, ApplyMapDelta(x, y, direction))
	}
	return neighbours
}

func NewMaze(d Dimensions) *Maze {
	gridDimensions := Dimensions{X: MazeToGrid(d.X), Y: MazeToGrid(d.Y)}
	m := Maze{
		D: d,
		G: Grid{
			D:     gridDimensions,
			Cells: make([]Cell, gridDimensions.X*gridDimensions.Y),
		},
	}
	return &m
}

func (m *Maze) InitializeGrid() {
	for i := range m.G.Cells {
		m.G.Cells[i] = CellWall
	}
	for x := 0; x < m.Width(); x++ {
		for y := 0; y < m.Height(); y++ {
			m.Set(x, y, CellBlank)
		}
	}
}
