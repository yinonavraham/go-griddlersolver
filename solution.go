package griddlersolver

import "github.com/yinonavraham/go-griddlersolver/grid"

// Solution is an iterator over states - from the initial (empty) state to the last solved state
type Solution interface {
	// Next returns the next state in the solution and an indicator whether there are more states
	Next() (state grid.Grid, hasNext bool)
}

type solution struct {
	problem Problem
}

func (s *solution) Next() (state grid.Grid, hasNext bool) {
	return nil, false
}

func (s *solution) updateProbabilities(g grid.Grid, rowCount int, colCount int, rowPossibleSolutions []possibleSolutions, colPossibleSolutions []possibleSolutions) (hasNext bool) {
	changed := false
	for r := 0; r < rowCount; r++ {
		for c := 0; c < colCount; c++ {
			value := g.GetCell(r, c)
			if value == grid.FullCellValue || value == grid.EmptyCellValue {
				continue
			}
			changed = true
			// - rows
			trueCount := 0
			falseCount := 0
			for i := 0; i < len(rowPossibleSolutions[r]); i++ {
				if rowPossibleSolutions[r][i][c] == true {
					trueCount++
				} else {
					falseCount++
				}
			}
			if trueCount == 0 {
				g.(grid.MutableGrid).SetCell(r, c, grid.EmptyCellValue)
				continue
			}
			if falseCount == 0 {
				g.(grid.MutableGrid).SetCell(r, c, grid.FullCellValue)
				continue
			}
			intermediateValue := intermediateCellValue{}
			intermediateValue.fullPossibilities = trueCount
			intermediateValue.emptyPossibilities = falseCount
			g.(grid.MutableGrid).SetCell(r, c, intermediateValue)

			// - columns
			trueCount = 0
			falseCount = 0
			for i := 0; i < len(colPossibleSolutions[c]); i++ {
				if colPossibleSolutions[c][i][r] == true {
					trueCount++
				} else {
					falseCount++
				}
			}
			if trueCount == 0 {
				intermediateValue.fullPossibilities = 0
			}
			intermediateValue.fullPossibilities += trueCount
			if falseCount == 0 {
				intermediateValue.emptyPossibilities = 0
			}
			intermediateValue.emptyPossibilities += falseCount
			g.(grid.MutableGrid).SetCell(r, c, intermediateValue)
		}
	}
	return changed
}

func (s *solution) removeFalseSolutions(g grid.Grid, rowCount int, colCount int, rowPossibleSolutions []possibleSolutions, colPossibleSolutions []possibleSolutions) (hasNext bool) {
	changed := false
	for r := 0; r < rowCount; r++ {
		for c := 0; c < colCount; c++ {
			value := g.GetCell(r, c)
			if value == grid.ZeroCellValue {
				continue
			}
			expected := false
			if v, ok := value.(intermediateCellValue); ok {
				if v.emptyPossibilities != 0 && v.fullPossibilities != 0 {
					continue
				}
				expected = v.fullPossibilities != 0
			} else {
				expected = value == grid.FullCellValue
			}
			// - rows
			var toRemove []int
			for i := 0; i < len(rowPossibleSolutions[r]); i++ {
				if rowPossibleSolutions[r][i][c] != expected {
					toRemove = append(toRemove, i)
				}
			}
			for i := len(toRemove)-1; i >= 0; i-- {
				rowPossibleSolutions[r] = removeAt(rowPossibleSolutions[r], i)
			}
			// - columns
			toRemove = []int{}
			for i := 0; i < len(colPossibleSolutions[c]); i++ {
				if colPossibleSolutions[c][i][r] != expected {
					toRemove = append(toRemove, i)
				}
			}
			for i := len(toRemove)-1; i >= 0; i-- {
				colPossibleSolutions[c] = removeAt(colPossibleSolutions[c], i)
			}

			changed = true
		}
	}
	return changed
}

type possibleSolutions []possibleSolution
type possibleSolution []bool

func (s possibleSolution) Clone() possibleSolution {
	clone := make(possibleSolution, len(s))
	for _, v := range s {
		clone = append(clone, v)
	}
	return clone
}

func removeAt(s possibleSolutions, index int) possibleSolutions {
	s1 := make(possibleSolutions, len(s)-1)
	for i := 0; i < len(s); i++ {
		if i == index {
			continue
		}
		s1 = append(s1, s[i])
	}
	return s1
}

func calcLinePossibleSolutions(def Definition, size int) possibleSolutions {
	solutions := possibleSolutions{}
	var solution = make(possibleSolution, size)
	for i := 0; i < size; i++ {
		solution[i] = false
	}
	calcRecursive(&solutions, solution, def, size, 0, 0)
	return solutions
}

func calcTotalLength(def Definition, part int) int {
	total := 0
	for i := part; i < len(def); i++ {
		if i > part {
			total += 1 // Add space
		}
		total += def[i]
	}
	return total
}

func calcRecursive(solutions *possibleSolutions, solution possibleSolution, def Definition, size int, location int, part int) bool {
	// If no more parts to lay - add the solution and return true
	if part >= len(def) {
		*solutions = append(*solutions, solution.Clone())
		return true
	}
	// If the remaining parts cannot fit in the current location and size - return false
	var defTotalLength = calcTotalLength(def, part)
	if defTotalLength > (size - location) {
		return false
	}
	// For each location from current location to size
	for i := location; i < (size - defTotalLength + 1); i++ {
		// - Place the current part and continue recursively to the next part
		for j := location; j < size; j++ {
			solution[j] = false
		}
		for j := i; j < i+def[part]; j++ {
			solution[j] = true
		}
		result := calcRecursive(solutions, solution, def, size, i+def[part]+1, part+1)
		// - If the return value is false - break
		if result == false {
			break
		}
	}
	return true
}

type intermediateCellValue struct {
	fullPossibilities  int
	emptyPossibilities int
}

func (v intermediateCellValue) Value() float32 {
	return float32(v.fullPossibilities) / float32(v.fullPossibilities+v.emptyPossibilities)
}
