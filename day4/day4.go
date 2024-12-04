package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Kazalo11/advent-of-code-2024/util"
)

var xmas = [4]string{"X", "M", "A", "S"}
var rowWidth int
var columnHeight int

type Coordinate struct{ row, column int }

var solutionPaths [][]Coordinate
var count int
var directions = map[string][2]int{
	"down-right": {1, 1},
	"down":       {1, 0},
	"down-left":  {1, -1},
	"right":      {0, 1},
	"left":       {0, -1},
	"up-right":   {-1, 1},
	"up":         {-1, 0},
	"up-left":    {-1, -1},
}

func init() {

	xmas = [4]string{"X", "M", "A", "S"}
}

func getData() [][]string {
	var matrix [][]string

	file, err := os.Open("day4/day4.txt")
	util.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		rowWidth = len(line)
		matrix = append(matrix, strings.Split(line, ""))
		count++
	}
	columnHeight = count

	return matrix

}

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

func part1() {

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
func main() {
	part1()
}
