package main

import (
	"fmt"
	"slices"
)

func part2() int {
	matrix := NewMatrix("day6/day6.txt")
	obstacles := matrix.FindObstacles()
	guardLocation := matrix.FindGuard()

	count := 0

	for i := 0; i < 130; i++ {
		for j := 0; j < 130; j++ {
			item := [2]int{i, j}
			fmt.Printf("Checking with obstacle put at: %v \n", item)

			if item != guardLocation && !slices.Contains(obstacles, item) {
				obstacles2 := append([][2]int{}, obstacles...)
				obstacles2 = append(obstacles2, item)
				if checkIfLoop(obstacles2, guardLocation) {
					fmt.Printf("Loop found if obstacle is put at row %d and col %d \n", item[0], item[1])
					count++
				}
			}
		}
	}

	return count
}

func checkIfLoop(obstacles [][2]int, guardLocation [2]int) bool {

	direction := UP

	oldLocations := [][2]int{guardLocation}

	var obstacle [2]int
	var escape bool

	iteration := 0

	for {
		switch direction {
		case UP:

			obstacle, escape = NextObstacleUp(guardLocation, obstacles)

			guardLocation = [2]int{
				obstacle[0] + 1,
				guardLocation[1],
			}

		case DOWN:

			obstacle, escape = NextObstacleDown(guardLocation, obstacles)

			guardLocation = [2]int{
				obstacle[0] - 1,
				guardLocation[1],
			}

		case LEFT:
			obstacle, escape = NextObstacleLeft(guardLocation, obstacles)

			guardLocation = [2]int{
				guardLocation[0],
				obstacle[1] + 1,
			}

		case RIGHT:
			obstacle, escape = NextObstacleRight(guardLocation, obstacles)

			guardLocation = [2]int{
				guardLocation[0],
				obstacle[1] - 1,
			}

		}

		if escape {
			return false
		}

		if checkIfInLoop(oldLocations, guardLocation) {
			for i := 0; i < len(oldLocations)-3; i++ {
				if checkIfRectangle(oldLocations[i : i+4]) {
					return true
				}
			}
			return false
		}

		oldLocations = append(oldLocations, guardLocation)

		direction = (direction + 1) % 4
		iteration++
	}
}

func checkIfInLoop(locations [][2]int, guardLocation [2]int) bool {
	count := 0
	for _, item := range locations {
		if item == guardLocation {
			count++
		}
	}
	return count > 2
}

func checkIfRectangle(oldLocations [][2]int) bool {
	first := oldLocations[0]
	last := oldLocations[3]

	return first[0] == last[0] || first[1] == last[1]

}
