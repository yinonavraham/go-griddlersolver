package grid_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yinonavraham/go-griddlersolver/grid"
	"testing"
)

func TestGrid(t *testing.T) {
	type cell struct {
		r int
		c int
		v grid.CellValue
	}
	tests := []struct {
		name  string
		rows  int
		cols  int
		cells []cell
	}{
		{
			name: "empty",
			rows: 0,
			cols: 0,
		},
		{
			name: "sparse",
			rows: 4,
			cols: 3,
			cells: []cell{
				{r: 0, c: 0, v: 0.2},
				{r: 2, c: 1, v: 0.7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var g grid.Grid
			g = grid.New(tt.rows, tt.cols)
			assert.Equal(t, tt.rows, g.Rows())
			assert.Equal(t, tt.cols, g.Columns())
			for r := 0; r < g.Rows(); r++ {
				for c := 0; c < g.Columns(); c++ {
					v := g.GetCell(r, c)
					assert.Equal(t, grid.ZeroCellValue, v, fmt.Sprintf("unexpected initial cell value at (%d,%d)", r, c))
				}
			}

			assert.Panics(t, func() { g.GetCell(tt.rows+1, 0) }, "GetCell expected to panic on row index out of bounds")
			assert.Panics(t, func() { g.GetCell(-1, 0) }, "GetCell expected to panic on negative row index")
			assert.Panics(t, func() { g.GetCell(0, tt.cols+1) }, "GetCell expected to panic on column index out of bounds")
			assert.Panics(t, func() { g.GetCell(0, -1) }, "GetCell expected to panic on negative column index")

			mg := g.(grid.MutableGrid)
			v := grid.CellValue(0)
			assert.Panics(t, func() { mg.SetCell(tt.rows+1, 0, v) }, "SetCell expected to panic on row index out of bounds")
			assert.Panics(t, func() { mg.SetCell(-1, 0, v) }, "SetCell expected to panic on negative row index")
			assert.Panics(t, func() { mg.SetCell(0, tt.cols+1, v) }, "SetCell expected to panic on column index out of bounds")
			assert.Panics(t, func() { mg.SetCell(0, -1, v) }, "SetCell expected to panic on negative column index")

			for _, cell := range tt.cells {
				mg.SetCell(cell.r, cell.c, cell.v)
				assert.Equal(t, cell.v, g.GetCell(cell.r, cell.c), "unexpected cell value after update")
			}
		})
	}
}

func TestNewWithValues(t *testing.T) {
	tests := []struct {
		name   string
		values [][]grid.CellValue
	}{
		{
			name:   "empty",
			values: [][]grid.CellValue{},
		},
		{
			name: "single cell",
			values: [][]grid.CellValue{
				{0.5},
			},
		},
		{
			name: "single row",
			values: [][]grid.CellValue{
				{0.1, 0.5, 0.7},
			},
		},
		{
			name: "multiple rows, equal lengths",
			values: [][]grid.CellValue{
				{0.1, 0.5, 0.7},
				{0.2, 0.4, 0.8},
				{0.4, 0.9, 0.9},
				{0.3, 0.6, 0.3},
			},
		},
		{
			name: "multiple rows, different lengths",
			values: [][]grid.CellValue{
				{0.1, 0.5},
				{0.2, 0.4, 0.8},
				{},
				{0.3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := grid.NewWithValues(tt.values)
			assert.Equal(t, len(tt.values), g.Rows(), "number of rows not as expected")
			var maxCols int
			for r := 0; r < g.Rows(); r++ {
				if len(tt.values[r]) > maxCols {
					maxCols = len(tt.values[r])
				}
				for c := 0; c < g.Columns(); c++ {
					expected := grid.ZeroCellValue
					if c < len(tt.values[r]) {
						expected = tt.values[r][c]
					}
					v := g.GetCell(r, c)
					assert.Equal(t, expected, v, fmt.Sprintf("unexpected cell value as (%d,%d)", r, c))
				}
			}
			assert.Equal(t, maxCols, g.Columns(), "number of columns not as expected")
		})
	}
}
