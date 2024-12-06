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
	FORWARD Direction = iota
	BACKWARD
	LEFT
	RIGHT
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

func (m *Matrix) GetRow(rowIndex int) []string {
	if rowIndex < 0 || rowIndex >= len(m.data) {
		return nil
	}
	return m.data[rowIndex]
}

func (m *Matrix) GetColumn(columnIndex int) []string {
	column := []string{}
	for _, row := range m.data {
		if columnIndex >= 0 && columnIndex < len(row) {
			column = append(column, row[columnIndex])
		}
	}
	return column
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
}
