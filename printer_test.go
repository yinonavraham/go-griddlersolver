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
				{0, 0.2, 0.4, 0.6, 0.8, 1},
				{0.1, 0.3, 0.5, 0.7, 0.9, 0},
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
