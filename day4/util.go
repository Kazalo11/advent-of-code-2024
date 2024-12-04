package main

import (
	"bufio"
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
