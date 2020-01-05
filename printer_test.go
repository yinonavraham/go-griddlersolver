package griddlersolver_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/yinonavraham/go-griddlersolver"
	"github.com/yinonavraham/go-griddlersolver/grid"
	"strings"
	"testing"
)

func TestPrintGrid(t *testing.T) {
	tests := []struct {
		name   string
		values [][]grid.CellValue
		want   string
	}{
		{
			name: "empty",
			want: `
┌┐
└┘
`,
		},
		{
			name: "with values",
			values: [][]grid.CellValue{
				{cellValue(0), cellValue(0.2), cellValue(0.4), cellValue(0.6), cellValue(0.8), cellValue(1)},
				{cellValue(0.1), cellValue(0.3), cellValue(0.5), cellValue(0.7), cellValue(0.9), cellValue(0)},
			},
			want: `
┌────────────┐
│  ░░▒▒▒▒▓▓██│
│░░░░▒▒▓▓▓▓  │
└────────────┘
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := grid.NewWithValues(tt.values)
			w := strings.Builder{}
			w.WriteString("\n")
			_ = griddlersolver.PrintGrid(&w, g)
			assert.Equal(t, tt.want, w.String())
		})
	}
}

type cellValue float32

func (v cellValue) Value() float32 {
	return float32(v)
}
