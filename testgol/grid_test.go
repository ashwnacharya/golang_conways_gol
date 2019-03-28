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

	gridMap, rows, columns, gridMapErr := gol.CreateGridMapFromLines(lines)

	if gridMapErr != nil {
		t.Error("GetGridMap should not return error")
	}

	if rows != 4 {
		t.Errorf("Rows should be 4, not %d", rows)
	}

	if columns != 8 {
		t.Errorf("Columns should be 8, not %d", columns)
	}

	if gridMap == nil {
		t.Errorf("Grid map should not be nil")
	}

	var expectedGridMap [4][8]bool

	expectedGridMap[0] = [8]bool{false, false, false, false, false, false, false, false}
	expectedGridMap[1] = [8]bool{false, false, false, false, true, false, false, false}
	expectedGridMap[2] = [8]bool{false, false, false, true, true, false, false, false}
	expectedGridMap[3] = [8]bool{false, false, false, false, false, false, false, false}

	for i := 0; i < 4; i++ {
		for j := 0; j < 8; j++ {
			if expectedGridMap[i][j] != gridMap[i][j] {
				t.Errorf("Incorrect flag at position Row: %d Col: %d", i, j)
			}
		}
	}
}


func TestCreateLinesFromGridMap(t *testing.T) {
	var gridMap [4][]bool

	gridMap[0] = []bool{false, false, false, false, false, false, false, false}
	gridMap[1] = []bool{false, false, false, false, true, false, false, false}
	gridMap[2] = []bool{false, false, false, true, true, false, false, false}
	gridMap[3] = []bool{false, false, false, false, false, false, false, false}

	var expectedLines []string

	expectedLines = append(expectedLines, "4 8")
	expectedLines = append(expectedLines, "........")
	expectedLines = append(expectedLines, "....*...")
	expectedLines = append(expectedLines, "...**...")
	expectedLines = append(expectedLines, "........")

	actualLines, err := gol.GetLinesFromGridMap(gridMap[:], 4, 8)

	if err != nil {
		t.Error("Should not return error")
	}

	for i := 1; i <= 4; i ++ {
		if expectedLines[i] != actualLines[i] {
			t.Errorf("Line does not match")
		}
	}
}
