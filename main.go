package main

import (
	"conways_game_of_life/gol"
	"io/ioutil"
	"log"
	"strings"
	"os"
)

func main() {
	contentsOfFile, fileReadErr := ioutil.ReadFile("input.txt")

	if fileReadErr != nil {
		log.Fatal(fileReadErr)
	}

	var lines = strings.Split(string(contentsOfFile), "\n")

	var gridMap, rows, columns, gridMapErr = gol.CreateGridMapFromLines(lines)

	if gridMapErr != nil {
		log.Fatal(gridMapErr)
	}

	nextgridMap := gol.GetNextGridMap(gridMap, rows, columns)

	newLines, linesErr := gol.GetLinesFromGridMap(nextgridMap, rows, columns)

	if linesErr != nil {
		log.Fatal("Cannot generate lines from grid map")
	}

	f, err := os.OpenFile("output.txt", os.O_CREATE | os.O_APPEND|os.O_WRONLY, 0644) 
	defer f.Close()
	
	if err != nil {
        log.Fatal(err)
	}
	
	for _, line := range newLines {
		f.WriteString(line)
		f.WriteString("\n") 
	}

}
