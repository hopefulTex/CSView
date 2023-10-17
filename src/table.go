package main

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	terminal "golang.org/x/term"
)

var (
	lineColor lipgloss.AdaptiveColor = lipgloss.AdaptiveColor{Light: "#F000F0", Dark: "#A000F0"}
	evenColor lipgloss.AdaptiveColor = lipgloss.AdaptiveColor{Light: "#ECDFEC", Dark: "#202020"}
	oddColor  lipgloss.AdaptiveColor = lipgloss.AdaptiveColor{Light: "", Dark: "#35353F"}
)

var (
	OuterStyle   lipgloss.Style = lipgloss.NewStyle()
	HeaderStyle  lipgloss.Style = lipgloss.NewStyle()
	EvenRowStyle lipgloss.Style = lipgloss.NewStyle().Background(evenColor)
	OddRowStyle  lipgloss.Style = lipgloss.NewStyle().Background(oddColor)
)

func NewTable(cells [][]string, alignment []int) *table.Table {
	if len(cells) < 1 {
		return table.New()
	}

	width, height, err := terminal.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 60
		height = 80
	}
	t := table.New().
		Headers(cells[0]...).
		Rows(cells[1:]...).Height(height).Width(width).
		BorderStyle(
			lipgloss.NewStyle().
				Foreground(lineColor)).
		StyleFunc(func(row, col int) lipgloss.Style {
			var style lipgloss.Style
			var direction lipgloss.Position
			switch alignment[col] {
			case 2:
				direction = lipgloss.Right
			case 3:
				direction = lipgloss.Center
			default:
				direction = lipgloss.Left
			}

			// flipped even and odd because the start is 1 not 0
			switch {
			case row == 0:
				style = HeaderStyle
			case row%2 == 0:
				style = OddRowStyle.AlignHorizontal(direction)
			default:
				style = EvenRowStyle.AlignHorizontal(direction)
			}
			return style
		})

	return t
}

// func Convert(data [][]string, kind string) [][]string {

// }
