package griddlersolver

import (
	"github.com/yinonavraham/go-griddlersolver/grid"
	"io"
	"strings"
)

const (
	// https://en.wikipedia.org/wiki/List_of_Unicode_characters#Block_Elements
	blockShadeLight  = "\u2591" // ░
	blockShadeMedium = "\u2592" // ▒
	blockShadeDark   = "\u2593" // ▓
	blockFull        = "\u2588" // █

	// https://en.wikipedia.org/wiki/List_of_Unicode_characters#Block_Drawing
	boxHorizontalLine  = "\u2500" // ─
	boxVerticalLine    = "\u2502" // │
	boxCornerUpLeft    = "\u250c" // ┌
	boxCornerUpRight   = "\u2510" // ┐
	boxCornerDownLeft  = "\u2514" // └
	boxCornerDownRight = "\u2518" // ┘
	boxMidConnectLeft  = "\u251c"
	boxMidConnectRight = "\u2524"
	boxMidConnectUp    = "\u252c"
	boxMidConnectDown  = "\u2534"
)

func PrintGrid(w io.Writer, g grid.Grid) error {
	cellSize := 2
	if _, err := w.Write([]byte(boxCornerUpLeft)); err != nil {
		return err
	}
	if _, err := w.Write([]byte(strings.Repeat(boxHorizontalLine, g.Columns()*cellSize))); err != nil {
		return err
	}
	if _, err := w.Write([]byte(boxCornerUpRight + "\n")); err != nil {
		return err
	}
	for r := 0; r < g.Rows(); r++ {
		if _, err := w.Write([]byte(boxVerticalLine)); err != nil {
			return err
		}
		for c := 0; c < g.Columns(); c++ {
			value := g.GetCell(r, c)
			var char string
			switch {
			case value == 0:
				char = " "
			case value < 0.33:
				char = blockShadeLight
			case value < 0.66:
				char = blockShadeMedium
			case value < 1:
				char = blockShadeDark
			case value >= 1:
				char = blockFull
			}
			if _, err := w.Write([]byte(strings.Repeat(char, cellSize))); err != nil {
				return err
			}
		}
		if _, err := w.Write([]byte(boxVerticalLine + "\n")); err != nil {
			return err
		}
	}
	if _, err := w.Write([]byte(boxCornerDownLeft)); err != nil {
		return err
	}
	if _, err := w.Write([]byte(strings.Repeat(boxHorizontalLine, g.Columns()*cellSize))); err != nil {
		return err
	}
	if _, err := w.Write([]byte(boxCornerDownRight + "\n")); err != nil {
		return err
	}
	return nil
}
