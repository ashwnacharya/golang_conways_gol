package testgol

import (
	"conways_game_of_life/gol"
	"testing"
)

func TestCreateGridMapFromLines(t *testing.T) {
	var lines []string

	lines = append(lines, "4 8")
	lines = append(lines, "........")
	lines = append(lines, "....*...")
	lines = append(lines, "...**...")
	lines = append(lines, "........")

	actualGridMap, rows, columns, gridMapErr := gol.CreateGridMapFromLines(lines)

	if gridMapErr != nil {
		t.Error("GetGridMap should not return error")
	}

	if rows != 4 {
		t.Errorf("Rows should be 4, not %d", rows)
	}

	if columns != 8 {
		t.Errorf("Columns should be 8, not %d", columns)
	}

	if actualGridMap == nil {
		t.Errorf("Grid map should not be nil")
	}

	expectedGridMap := make([][]bool, 4)

	expectedGridMap[0] = make([]bool, 8)
	expectedGridMap[0] = []bool{false, false, false, false, false, false, false, false}

	expectedGridMap[1] = make([]bool, 8)
	expectedGridMap[1] = []bool{false, false, false, false, true, false, false, false}
	
	expectedGridMap[2] = make([]bool, 8)
	expectedGridMap[2] = []bool{false, false, false, true, true, false, false, false}
	
	expectedGridMap[3] = make([]bool, 8)
	expectedGridMap[3] = []bool{false, false, false, false, false, false, false, false}

	for i := 0; i < 4; i++ {
		for j := 0; j < 8; j++ {
			if expectedGridMap[i][j] != actualGridMap[i][j] {
				t.Errorf("Incorrect flag at position Row: %d Col: %d", i, j)
			}
		}
	}
}

func TestCreateLinesFromGridMap(t *testing.T) {
	gridMap := make([][]bool, 4)

	gridMap[0] = make([]bool, 8)
	gridMap[0] = []bool{false, false, false, false, false, false, false, false}

	gridMap[1] = make([]bool, 8)
	gridMap[1] = []bool{false, false, false, false, true, false, false, false}
	
	gridMap[2] = make([]bool, 8)
	gridMap[2] = []bool{false, false, false, true, true, false, false, false}
	
	gridMap[3] = make([]bool, 8)
	gridMap[3] = []bool{false, false, false, false, false, false, false, false}

	var expectedLines []string

	expectedLines = append(expectedLines, "4 8")
	expectedLines = append(expectedLines, "........")
	expectedLines = append(expectedLines, "....*...")
	expectedLines = append(expectedLines, "...**...")
	expectedLines = append(expectedLines, "........")

	actualLines, err := gol.GetLinesFromGridMap(gridMap, 4, 8)

	if err != nil {
		t.Error("Should not return error")
	}

	for i := 1; i <= 4; i ++ {
		if expectedLines[i] != actualLines[i] {
			t.Errorf("Line does not match")
		}
	}
}

func TestGetNextGridMap(t *testing.T) {
	gridMap := make([][]bool, 4)

	gridMap[0] = make([]bool, 8)
	gridMap[0] = []bool{false, false, false, false, false, false, false, false}

	gridMap[1] = make([]bool, 8)
	gridMap[1] = []bool{false, false, false, false, true, false, false, false}
	
	gridMap[2] = make([]bool, 8)
	gridMap[2] = []bool{false, false, false, true, true, false, false, false}
	
	gridMap[3] = make([]bool, 8)
	gridMap[3] = []bool{false, false, false, false, false, false, false, false}

	actualNextGridMap := gol.GetNextGridMap(gridMap, 4, 8)


	expectedGridMap := make([][]bool, 4)

	expectedGridMap[0] = make([]bool, 8)
	expectedGridMap[0] = []bool{false, false, false, false, false, false, false, false}

	expectedGridMap[1] = make([]bool, 8)
	expectedGridMap[1] = []bool{false, false, false, true, true, false, false, false}
	
	expectedGridMap[2] = make([]bool, 8)
	expectedGridMap[2] = []bool{false, false, false, true, true, false, false, false}
	
	expectedGridMap[3] = make([]bool, 8)
	expectedGridMap[3] = []bool{false, false, false, false, false, false, false, false}

	for i := 0; i < 4; i++ {
		for j := 0; j < 8; j++ {
			if expectedGridMap[i][j] != actualNextGridMap[i][j] {
				t.Errorf("Incorrect flag at position Row: %d Col: %d", i, j)
			}
		}
	}
}