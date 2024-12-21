package main

import (
	"fmt"
	"slices"
)

func getPath(route string, keypad map[[2]int]string, index int, coord [2]int, path [][2]int) {
	fmt.Println(path)
	_, exists := keypad[coord]
	if !exists {
		fmt.Println("coordinate is not in map, returning")
		return
	}
	if slices.Contains(path, coord) {
		fmt.Println("coordinate already exists in path, returning")
		return
	}
	curr := keypad[coord]
	ideal := string(route[index])
	if curr == ideal {
		fmt.Printf("Current index is %d \n", index)
		if index == len(route)-1 {
			path = append(path, [2]int{3, 2})
			total_path[route] = path
			return
		}
		index++

	}
	path = append(path, coord)
	x, y := coord[0], coord[1]
	getPath(route, keypad, index, [2]int{x + 1, y}, path)
	getPath(route, keypad, index, [2]int{x - 1, y}, path)
	getPath(route, keypad, index, [2]int{x, y + 1}, path)
	getPath(route, keypad, index, [2]int{x, y - 1}, path)

}

func getFullPath(route string) {
	empty_path := make([][2]int, 0)
	getPath(route, numericKeypad, 0, [2]int{3, 2}, empty_path)
	fmt.Println(total_path)

}

func main() {
	part1()

}
