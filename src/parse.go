package main

import (
	"strings"
)

func toCSV(cells [][]string) string {
	var str strings.Builder
	for _, row := range cells {
		for i, cell := range row {
			if strings.ContainsAny(cell, ",") {
				row[i] = "\"" + cell + "\""
			}
		}
		line := strings.Join(row, ",")
		str.WriteString(line)
		str.WriteRune('\n')
	}
	return str.String()
}

func toMD(cells [][]string) string {
	var str strings.Builder
	for i, row := range cells {
		str.WriteRune('|')
		if i == 1 {
			for range row {
				str.WriteString("-|")
			}
			str.WriteString("\n|")
		}

		line := strings.Join(row, "|")
		str.WriteString(line)
		str.WriteString("|\n")
	}
	return str.String()
}

func parseMD(data string) ([][]string, []int, error) {
	var alignment []int // grabs alignment for each column from the neck

	// list of rows
	lines := strings.Split(data, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	var rowOffset int = 0 // the neck isnt visible
	var cells [][]string = make([][]string, len(lines)-1)
	var row []string

	for i, line := range lines {
		// row -> []cells,
		row = strings.Split(line, "|")
		// check alignment
		if i == 1 {
			alignment = getAlignment(row)
			rowOffset = 1
			continue
		}
		// trim padding
		if row[0] == "" {
			row = row[1:]
		}
		if row[len(row)-1] == "" {
			row = row[:len(row)-1]
		}
		cells[i-rowOffset] = row
	}

	return cells, alignment, nil
}

// 0,1 - left; 2 - right; 3 - center
func getAlignment(cells []string) []int {
	var alignment []int = make([]int, len(cells))

	for i := range alignment {
		alignment[i] = 0
		if strings.HasPrefix(cells[i], ":") {
			alignment[i] = 1
		}
		if strings.HasSuffix(cells[i], ":") {
			alignment[i] += 2
		}
	}

	return alignment
}
