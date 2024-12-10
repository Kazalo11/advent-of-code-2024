package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Kazalo11/advent-of-code-2024/util"
)

func part1() {
	grid := getData()

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

var trailheads = make(map[[2]int][][2]int)

func countingTrailheads(grid [][]int, row, col, nextNum int, startPosition [2]int) {
	if row < 0 || row >= len(grid[0]) || col < 0 || col >= len(grid) {
		return
	}

	curr := grid[row][col]
	if curr != nextNum {
		return
	}

	if curr == 9 && nextNum == 9 {
		trailheads[startPosition] = append(trailheads[startPosition], [2]int{row, col})
		// positions := trailheads[startPosition]
		// if !slices.Contains(positions, [2]int{row, col}) {
		// 	trailheads[startPosition] = append(trailheads[startPosition], [2]int{row, col})

		// }

		//Add in for part 1

		return
	}
	countingTrailheads(grid, row+1, col, nextNum+1, startPosition)
	countingTrailheads(grid, row-1, col, nextNum+1, startPosition)
	countingTrailheads(grid, row, col-1, nextNum+1, startPosition)
	countingTrailheads(grid, row, col+1, nextNum+1, startPosition)
}

func searchGrid(grid [][]int) int {
	count := 0
	for i, row := range grid {
		for j, item := range row {
			if item == 0 {
				trailheads[[2]int{i, j}] = make([][2]int, 0)
				countingTrailheads(grid, i, j, 0, [2]int{i, j})
			}
		}
	}
	for _, item := range trailheads {
		count += len(item)
	}
	return count
}
