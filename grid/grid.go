package grid

import "fmt"

type CellValue float32

const (
	ZeroCellValue CellValue = 0
)

type Grid interface {
	Rows() int
	Columns() int
	GetCell(row, column int) CellValue
}

type MutableGrid interface {
	Grid
	SetCell(row, column int, value CellValue)
}

// Ensure grid satisfies both Grid and MutableGrid interfaces
var _ Grid = &grid{}
var _ MutableGrid = &grid{}

type grid struct {
	rows    int
	columns int
	cells   [][]CellValue
}

func NewWithValues(values [][]CellValue) *grid {
	rows := len(values)
	columns := 0
	if rows > 0 {
		for _, row := range values {
			if len(row) > columns {
				columns = len(row)
			}
		}
	}
	g := New(rows, columns)
	for r := 0; r < rows; r++ {
		for c := 0; c < len(values[r]); c++ {
			g.SetCell(r, c, values[r][c])
		}
	}
	return g
}

func New(rows, columns int) *grid {
	g := grid{
		rows:    rows,
		columns: columns,
	}
	g.cells = make([][]CellValue, g.rows)
	for i := 0; i < g.rows; i++ {
		g.cells[i] = make([]CellValue, g.columns)
	}
	return &g
}

func (g *grid) Rows() int {
	return g.rows
}
func (g *grid) Columns() int {
	return g.columns
}

func (g *grid) SetCell(row, col int, value CellValue) {
	g.assertCoordinates(row, col)
	g.cells[row][col] = value
}

func (g *grid) GetCell(row, col int) CellValue {
	g.assertCoordinates(row, col)
	return g.cells[row][col]
}

func (g *grid) assertCoordinates(row int, col int) {
	if row < 0 || row >= g.rows {
		panic(fmt.Errorf("row out of range: %d [0..%d)", row, g.rows))
	}
	if col < 0 || col >= g.columns {
		panic(fmt.Errorf("column out of range: %d [0..%d)", col, g.columns))
	}
}
