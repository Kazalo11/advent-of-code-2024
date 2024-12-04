package main

import "fmt"

var mas = map[[2]string][2]string{
	{"M", "S"}: {"M", "S"},
	{"S", "M"}: {"S", "M"},
	{"S", "S"}: {"M", "M"},
	{"M", "M"}: {"S", "S"},
}

func checkMas(row, column int, matrix [][]string) bool {
	if row == 0 || column == 0 || row == rowWidth-1 || column == columnHeight-1 {
		return false
	}

	rowAbove := [2]string{matrix[row-1][column-1], matrix[row-1][column+1]}
	rowBelow := [2]string{matrix[row+1][column-1], matrix[row+1][column+1]}

	if mas[[2]string(rowAbove)] == [2]string(rowBelow) {
		fmt.Printf("Found X-MAS for position %d, %d \n", row, column)
		return true
	}
	return false

}

func Part2() {
	matrix := getData()
	count := 0
	for rowIndex, row := range matrix {
		for column := range row {
			currentLetter := matrix[rowIndex][column]
			if currentLetter == "A" {
				if checkMas(rowIndex, column, matrix) {
					count++
				}
			}
		}
	}
	fmt.Printf("Answer to part 2: %d \n", count)
}
