package gol

import (
	"testing"
)

func TestLiveNeighbourCount00(t *testing.T) {
	gridMap := make([][]bool, 4)

	gridMap[0] = make([]bool, 8)
	gridMap[0] = []bool{true, true, true, true, true, true, true, true}

	gridMap[1] = make([]bool, 8)
	gridMap[1] = []bool{true, true, true, true, true, true, true, true}

	gridMap[2] = make([]bool, 8)
	gridMap[2] = []bool{true, true, true, true, true, true, true, true}

	gridMap[3] = make([]bool, 8)
	gridMap[3] = []bool{true, true, true, true, true, true, true, true}

	actual := getliveNeighbourCount(0, 0, gridMap, 4, 8)

	if actual != 3 {
		t.Errorf("NUmber of neighbours should be 3, not %d", actual)
	}
}

func TestLiveNeighbourCount37(t *testing.T) {
	gridMap := make([][]bool, 4)

	gridMap[0] = make([]bool, 8)
	gridMap[0] = []bool{true, true, true, true, true, true, true, true}

	gridMap[1] = make([]bool, 8)
	gridMap[1] = []bool{true, true, true, true, true, true, true, true}

	gridMap[2] = make([]bool, 8)
	gridMap[2] = []bool{true, true, true, true, true, true, true, true}

	gridMap[3] = make([]bool, 8)
	gridMap[3] = []bool{true, true, true, true, true, true, true, true}

	actual := getliveNeighbourCount(3, 7, gridMap, 4, 8)

	if actual != 3 {
		t.Errorf("NUmber of neighbours should be 3, not %d", actual)
	}
}

func TestLiveNeighbourCount30(t *testing.T) {
	gridMap := make([][]bool, 4)

	gridMap[0] = make([]bool, 8)
	gridMap[0] = []bool{true, true, true, true, true, true, true, true}

	gridMap[1] = make([]bool, 8)
	gridMap[1] = []bool{true, true, true, true, true, true, true, true}

	gridMap[2] = make([]bool, 8)
	gridMap[2] = []bool{true, true, true, true, true, true, true, true}

	gridMap[3] = make([]bool, 8)
	gridMap[3] = []bool{true, true, true, true, true, true, true, true}

	actual := getliveNeighbourCount(3, 0, gridMap, 4, 8)

	if actual != 3 {
		t.Errorf("NUmber of neighbours should be 3, not %d", actual)
	}
}

func TestLiveNeighbourCount07(t *testing.T) {
	gridMap := make([][]bool, 4)

	gridMap[0] = make([]bool, 8)
	gridMap[0] = []bool{true, true, true, true, true, true, true, true}

	gridMap[1] = make([]bool, 8)
	gridMap[1] = []bool{true, true, true, true, true, true, true, true}

	gridMap[2] = make([]bool, 8)
	gridMap[2] = []bool{true, true, true, true, true, true, true, true}

	gridMap[3] = make([]bool, 8)
	gridMap[3] = []bool{true, true, true, true, true, true, true, true}

	actual := getliveNeighbourCount(3, 7, gridMap, 4, 8)

	if actual != 3 {
		t.Errorf("NUmber of neighbours should be 3, not %d", actual)
	}
}

func TestLiveNeighbourCount22(t *testing.T) {
	gridMap := make([][]bool, 4)

	gridMap[0] = make([]bool, 8)
	gridMap[0] = []bool{true, true, true, true, true, true, true, true}

	gridMap[1] = make([]bool, 8)
	gridMap[1] = []bool{true, true, true, true, true, true, true, true}

	gridMap[2] = make([]bool, 8)
	gridMap[2] = []bool{true, true, true, true, true, true, true, true}

	gridMap[3] = make([]bool, 8)
	gridMap[3] = []bool{true, true, true, true, true, true, true, true}

	actual := getliveNeighbourCount(2, 2, gridMap, 4, 8)

	if actual != 8 {
		t.Errorf("NUmber of neighbours should be 8, not %d", actual)
	}
}

func TestGetNextStatusAlive1(t *testing.T) {
	actual := getNextStatus(true, 1)

	if actual != false {
		t.Errorf("Any live cell with fewer than two live neighbours should die.")
	}
}

func TestGetNextStatusAlive4(t *testing.T) {
	actual := getNextStatus(true, 4)

	if actual != false {
		t.Errorf("Any live cell with more than three live neighbours should die.")
	}
}

func TestGetNextStatusAlive2(t *testing.T) {
	actual := getNextStatus(true, 2)

	if actual != true {
		t.Errorf("Any live cell with two or three live neighbours should live on to the next generation.")
	}
}

func TestGetNextStatusAlive3(t *testing.T) {
	actual := getNextStatus(true, 3)

	if actual != true {
		t.Errorf("Any live cell with two or three live neighbours should live on to the next generation.")
	}
}

func TestGetNextStatusDead3(t *testing.T) {
	actual := getNextStatus(false, 3)

	if actual != true {
		t.Errorf("Any dead cell with exactly three live neighbours should become a live cell.")
	}
}

func TestGetNextStatusDead(t *testing.T) {
	actual := getNextStatus(false, 1)

	if actual != false {
		t.Errorf("Any dead cell will remain dead unless it has exactly 3 live neighbours.")
	}
}