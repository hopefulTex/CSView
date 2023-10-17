package main

import (
	"encoding/csv"
	"fmt"
	"os"
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
