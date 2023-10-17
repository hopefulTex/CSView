package main

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	terminal "golang.org/x/term"
)

var (
	lineColor lipgloss.AdaptiveColor = lipgloss.AdaptiveColor{Light: "#F000F0", Dark: "#F000F0"}
	evenColor lipgloss.AdaptiveColor = lipgloss.AdaptiveColor{Light: "#202020", Dark: "#202020"}
	oddColor  lipgloss.AdaptiveColor = lipgloss.AdaptiveColor{Light: "#454545", Dark: "#454545"}
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

			switch {
			case row == 0:
				style = HeaderStyle
			case row%2 == 0:
				style = EvenRowStyle.AlignHorizontal(direction)
			default:
				style = OddRowStyle.AlignHorizontal(direction)
			}
			return style
		})

	return t
}

// func Convert(data [][]string, kind string) [][]string {

// }
