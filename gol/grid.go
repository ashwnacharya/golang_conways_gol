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

// GetNextGridMap calculates the next grid map, given the current grid map
func GetNextGridMap(gridMap [][]bool, rows int, columns int) ([][]bool){

	nextGridMap := make([][]bool, rows)

	for i := 0; i < rows; i++ {
		nextGridMap[i] = make([]bool, columns)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			currentCellStatus := gridMap[i][j]
			liveNeighbourCount := getliveNeighbourCount(i, j, gridMap, rows, columns)
			nextGridMap[i][j] = getNextStatus(currentCellStatus, liveNeighbourCount)
		}
	}

	return nextGridMap
}

func getNextStatus(currentCellStatus bool, liveNeighbourCount int) (bool) {
	// Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.
	if currentCellStatus == true && liveNeighbourCount < 2 {
		return false
	}

	// Any live cell with more than three live neighbours dies, as if by overcrowding.
	if currentCellStatus == true && liveNeighbourCount > 3 {
		return false
	}

	// Any live cell with two or three live neighbours lives on to the next generation.
	if currentCellStatus == true && (liveNeighbourCount == 2 || liveNeighbourCount == 3) {
		return true
	}

	// Any dead cell with exactly three live neighbours becomes a live cell.
	if currentCellStatus == false && liveNeighbourCount == 3 {
		return true
	}

	// All other dead cells remain dead
	return false
}

func getliveNeighbourCount(x int, y int, gridMap [][]bool, rows int, columns int) (int) {
	
	liveNeighbourCount := 0

	for i := x-1; i <= x+1; i++ {
		for j := y-1; j <= y+1; j++ {

			if i == x && j == y {
				continue
			}

			if i < 0 || j < 0 || i >= rows || j >= columns {
				continue
			}

			if gridMap[i][j] == true {
				liveNeighbourCount = liveNeighbourCount + 1
			}
		}
	}

	return liveNeighbourCount
}