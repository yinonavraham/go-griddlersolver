package grid

// CellValue type
type CellValue interface {
	Value() float32
}

var _ CellValue = cellValue(0)

type cellValue float32

func (v cellValue) Value() float32 {
	return float32(v)
}

const (
	// ZeroCellValue is the zero value for a cell in the grid
	ZeroCellValue cellValue = -1
	// FullCellValue is the value for a full cell
	FullCellValue cellValue = 1
	// EmptyCellValue is the value for an empty cell
	EmptyCellValue cellValue = 0
)
