package grid

import "fmt"

// Grid is an interface which represents a grid
type Grid interface {
	// Number of rows in the grid
	Rows() int
	// Number of columns in the grid
	Columns() int
	// Get the value of a given cell at a given row and column
	GetCell(row, column int) CellValue
	// Clone the grid
	Clone() Grid
}

// MutableGrid is an interface which represents a mutable grid - a Grid which can be modified
//
// Usually used with type assertion:
//   g := ...
//   if mg, ok := g.(grid.MutableGrid); ok {
//     mg.SetCell(...)
//   }
type MutableGrid interface {
	Grid
	// Set tne value of a cell in the grid at a given row and column
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

// NewWithValues creates a new grid with the provided values.
// The grid size is determined by the number of rows in the provided values and the max row length is the number of columns.
// If there are rows of different sizes, the undefined cells are initialized with the ZeroCellValue.
func NewWithValues(values [][]CellValue) Grid {
	rows := len(values)
	columns := 0
	if rows > 0 {
		for _, row := range values {
			if len(row) > columns {
				columns = len(row)
			}
		}
	}
	g := New(rows, columns).(*grid)
	for r := 0; r < rows; r++ {
		for c := 0; c < len(values[r]); c++ {
			g.SetCell(r, c, values[r][c])
		}
	}
	return g
}

// New creates a new grid with the provided size (number of rows and columns)
func New(rows, columns int) Grid {
	g := grid{
		rows:    rows,
		columns: columns,
	}
	g.cells = make([][]CellValue, g.rows)
	for r := 0; r < g.rows; r++ {
		g.cells[r] = make([]CellValue, g.columns)
		for c := 0; c < g.columns; c++ {
			g.cells[r][c] = ZeroCellValue
		}
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

func (g *grid) Clone() Grid {
	clone := grid{
		rows:    g.rows,
		columns: g.columns,
	}
	clone.cells = make([][]CellValue, g.rows)
	for r := 0; r < g.rows; r++ {
		clone.cells[r] = make([]CellValue, g.columns)
		for c := 0; c < g.columns; c++ {
			clone.cells[r][c] = g.cells[r][c]
		}
	}
	return &clone
}
