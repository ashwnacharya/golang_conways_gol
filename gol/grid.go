package gol

import (
	"fmt"
	"errors"
	"strconv"
	"strings"
)

// CreateGridMapFromLines creates a GridMap from an array of strings
func CreateGridMapFromLines(lines []string) (gridMap [][]bool, rows int, columns int, e error) {

	var headers = strings.Split(lines[0], " ")

	if len(headers) != 2 {
		return gridMap, rows, columns, errors.New("Invalid Headers")
	}

	rows, rowErr := strconv.Atoi(headers[0])
	if rowErr != nil {
		return gridMap, rows, columns, errors.New("Failed to convert the row header to a valid integer")
	}

	columns, colErr := strconv.Atoi(headers[1])
	if colErr != nil {
		return gridMap, rows, columns, errors.New("Failed to convert column header to a valid integer")
	}

	gridMap = make([][]bool, rows)
	for i := range gridMap {
		gridMap[i] = make([]bool, columns)
	}

	for i := 1; i <= rows; i++ {
		for j := 0; j < columns; j++ {
			if string(lines[i][j]) == "*" {
				gridMap[i-1][j] = true
			}
			if string(lines[i][j]) == "." {
				gridMap[i-1][j] = false
			}
		}
	}
	return gridMap, rows, columns, nil
}


// GetLinesFromGridMap creates test content that represents a gridmap
func GetLinesFromGridMap(gridMap [][]bool, rows int, columns int) ([]string, error) {
	var lines []string

	lines = append(lines, fmt.Sprintf("%d %d", rows, columns))

	for i := 0; i < rows; i++ {

		str := ""
		for j := 0; j < columns; j++ {
			if gridMap[i][j] == true {
				str = str + "*"
			} else {
				str = str + "."
			}
		}

		lines = append(lines, str)
	}

	return lines, nil
}
