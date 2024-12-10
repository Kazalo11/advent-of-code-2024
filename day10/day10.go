package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Kazalo11/advent-of-code-2024/util"
)

func part1() {
	grid := getData()
	fmt.Println(grid)

	ans := searchGrid(grid)
	fmt.Println(ans)

}

func getData() [][]int {
	file, err := os.Open("day10.txt")
	util.Check(err)
	grid := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		arr := make([]int, 0)
		line := scanner.Text()
		for _, char := range line {
			arr = append(arr, int(char-'0'))
		}
		grid = append(grid, arr)

	}
	return grid
}

func countTrailheads(grid [][]int, row, col int, nextNum int) int {
	if row < 0 || row >= len(grid[0]) || col < 0 || col >= len(grid) {
		return 0
	}
	curr := grid[row][col]
	if curr != nextNum {
		return 0
	}
	if curr == 9 && nextNum == 9 {
		return 1
	}
	return countTrailheads(grid, row+1, col, curr+1) + countTrailheads(grid, row, col+1, curr+1) + countTrailheads(grid, row, col-1, curr+1) + countTrailheads(grid, row-1, col, curr+1)

}

func searchGrid(grid [][]int) int {
	count := 0
	for i, row := range grid {
		for j, item := range row {
			if item == 0 {
				count += countTrailheads(grid, i, j, 0)
			}
		}
	}
	return count
}
