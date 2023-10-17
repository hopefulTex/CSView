package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func Write(path string, kind string, cells [][]string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	var data string

	switch kind {
	case "md":
		data = toMD(cells)
	case "csv":
		data = toCSV(cells)
	}

	_, err = f.WriteString(data)
	return err
}

func Open(path string, kind string) ([][]string, []int, error) {
	var alignment []int
	var data [][]string

	switch kind {
	case "csv":
		f, err := os.Open(path)
		if err != nil {
			return nil, nil, err
		}
		r := csv.NewReader(f)
		data, err = r.ReadAll()
		f.Close()
		if err != nil {
			return nil, nil, err
		}
		alignment = make([]int, len(data[0]))
		for i := range alignment {
			alignment[i] = 0
		}
	case "md":
		bytes, err := os.ReadFile(path)
		if err != nil {
			return nil, nil, err
		}

		data, alignment, err = parseMD(string(bytes))
		if err != nil {
			return nil, nil, err
		}
	default:
		return nil, nil, fmt.Errorf("incorrect file type. Want csv or md")
	}

	return data, alignment, nil
}

func parseMD(data string) ([][]string, []int, error) {
	var content [][]string
	var alignment []int
	lines := strings.Split(data, "\n")
	var cells [][]string = make([][]string, len(lines))
	for i, line := range lines {
		cells[i] = strings.Split(line, "|")
		if cells[i][0] == "" {
			cells[i] = cells[i][1:]
		}
		if cells[i][len(cells[i])-1] == "" {
			cells[i] = cells[i][:len(cells[i])-1]
		}
	}

	for i, row := range cells {
		if i == 1 {
			alignment = getAlignment(row)
		}
	}

	return content, alignment, nil
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

func toCSV(cells [][]string) string {
	var str strings.Builder
	for _, row := range cells {
		for i, cell := range row {
			if strings.ContainsAny(cell, ",") {
				row[i] = "\"" + cell + "\""
			}
		}
		line := strings.Join(row, ", ")
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
		}

		line := strings.Join(row, "|")
		str.WriteString(line)
		str.WriteString("|\n")
	}
	return str.String()
}
