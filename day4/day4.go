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

func nextLetter(index int, row int, column int, matrix [][]string, path []Coordinate) bool {
	if row < 0 || column < 0 || row >= rowWidth || column >= columnHeight {
		return false
	}

	currentLetter := matrix[row][column]

	if (currentLetter == "X" && index != 0) || (currentLetter != "X" && index == 0) {
		return false
	}

	if currentLetter == "S" && index == len(xmas)-1 {
		coordinate := Coordinate{row: row, column: column}
		path = append(path, coordinate)
		solutionPaths = append(solutionPaths, path)
		fmt.Printf("Found XMAS with path: %v \n", path)
		return true
	}

	expectedLetter := xmas[index]
	// fmt.Printf("Current letter is %s, expected Letter is %s \n", currentLetter, expectedLetter)

	if currentLetter != expectedLetter {
		return false
	}

	coordinate := Coordinate{row: row, column: column}
	path = append(path, coordinate)

	return nextLetter(index+1, row+1, column+1, matrix, path) ||
		nextLetter(index+1, row+1, column, matrix, path) ||
		nextLetter(index+1, row+1, column-1, matrix, path) ||
		nextLetter(index+1, row, column+1, matrix, path) ||
		nextLetter(index+1, row, column-1, matrix, path) ||
		nextLetter(index+1, row-1, column+1, matrix, path) ||
		nextLetter(index+1, row-1, column, matrix, path) ||
		nextLetter(index+1, row-1, column-1, matrix, path)

}

func main() {
	// count := 0
	matrix := getData()
	for rowIndex, row := range matrix {
		for column := range row {
			path := make([]Coordinate, 0)
			// path = append(path, coordinate)
			// if nextLetter(0, rowIndex, column, matrix, path) {
			// 	count++
			// }
			nextLetter(0, rowIndex, column, matrix, path)
		}
	}
	fmt.Print(len(solutionPaths))

}
