package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"

	"github.com/Kazalo11/advent-of-code-2024/util"
)

var (
	numericKeypad = map[[2]int]string{
		{0, 0}: "7",
		{1, 0}: "8",
		{2, 0}: "9",
		{0, 1}: "4",
		{1, 1}: "5",
		{2, 1}: "6",
		{0, 2}: "1",
		{1, 2}: "2",
		{2, 2}: "3",
		{1, 3}: "0",
		{2, 3}: "A",
	}
	directionalKeypad = map[[2]int]string{
		{1, 0}: "^",
		{2, 0}: "A",
		{0, 1}: "<",
		{1, 1}: "v",
		{1, 2}: ">",
	}
	total_path = make(map[string][][2]int)
)

func findPath(route string, keypad map[[2]int]string, index int, coord [2]int, coord_path [][2]int, direction_path []string) {
	fmt.Println(direction_path)
	_, exists := keypad[coord]
	if !exists {
		fmt.Println("Coordinate isn't in keypad, returning")
		return
	}

	if slices.Contains(coord_path, coord) {
		fmt.Println("Coordinate has already been travelled to")
		return
	}
	current_key := keypad[coord]
	ideal_key := string(route[index])

	if current_key == ideal_key {
		fmt.Println("Key found")
		direction_path = append(direction_path, "A")
		if index == len(route)-1 {
			fmt.Println("Found path")

		}
		index++
		fmt.Println(index)
	}

	coord_path = append(coord_path, coord)
	x, y := coord[0], coord[1]

	left := append(direction_path, "<")
	right := append(direction_path, ">")
	down := append(direction_path, "v")
	up := append(direction_path, "^")

	findPath(route, keypad, index, [2]int{x + 1, y}, coord_path, right)
	findPath(route, keypad, index, [2]int{x - 1, y}, coord_path, left)
	findPath(route, keypad, index, [2]int{x, y + 1}, coord_path, down)
	findPath(route, keypad, index, [2]int{x, y - 1}, coord_path, up)

}

func getFullPath(route string) {
	empty_path := make([][2]int, 0)
	empty_string_path := make([]string, 0)
	findPath(route, numericKeypad, 0, [2]int{2, 3}, empty_path, empty_string_path)
	fmt.Println()
}

func part1() {
	file, err := os.Open("day21.txt")
	util.Check(err)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		getFullPath(line)

	}
}

func main() {
	part1()
}
