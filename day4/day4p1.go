package main

import (
	"fmt"
)

func nextLetter(index int, row int, column int, matrix [][]string, path []Coordinate, direction string) bool {
	if row < 0 || column < 0 || row >= rowWidth || column >= columnHeight {
		return false
	}

	// fmt.Printf("Current path for %d, %d is %v \n", row, column, path)

	currentLetter := matrix[row][column]

	if currentLetter == "S" && index == len(xmas)-1 {
		coordinate := Coordinate{row: row, column: column}
		path = append(path, coordinate)
		solutionPaths = append(solutionPaths, path)
		fmt.Printf("Found XMAS with path: %v \n", path)
		return true
	}

	expectedLetter := xmas[index]

	if currentLetter != expectedLetter {
		return false
	}

	coordinate := Coordinate{row: row, column: column}
	path = append(path, coordinate)
	dir := directions[direction]
	newRow, newCol := row+dir[0], column+dir[1]

	return nextLetter(index+1, newRow, newCol, matrix, path, direction)

}

func Part1() {

	matrix := getData()
	for rowIndex, row := range matrix {
		for column := range row {
			path := make([]Coordinate, 0)

			start_letter := matrix[rowIndex][column]
			if start_letter == "X" {
				for direction := range directions {

					nextLetter(0, rowIndex, column, matrix, path, direction)
				}
			}

		}
	}
	fmt.Printf("Answer to part 1: %d \n", len(solutionPaths))

}
