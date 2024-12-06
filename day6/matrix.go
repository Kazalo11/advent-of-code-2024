package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/Kazalo11/advent-of-code-2024/util"
)

type Matrix struct {
	data [][]string
}

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

func NewMatrix(filePath string) *Matrix {
	var matrix [][]string

	file, err := os.Open(filePath)
	util.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		matrix = append(matrix, row)
	}

	return &Matrix{data: matrix}
}

func (m *Matrix) FindObstacles() [][2]int {
	obstacles := make([][2]int, 0)
	for rIndex, row := range m.data {
		for cIndex, _ := range row {
			if m.IsObstacle(rIndex, cIndex) {
				obstacles = append(obstacles, [2]int{rIndex, cIndex})
			}
		}
	}
	return obstacles
}

func (m *Matrix) IsObstacle(row, col int) bool {
	i := m.data[row][col]
	return i == "#"
}

func (m *Matrix) FindGuard() [2]int {
	for rowIndex, row := range m.data {
		for colIndex, i := range row {
			if i == "^" {
				return [2]int{rowIndex, colIndex}
			}
		}
	}
	panic("No guard found")
}
