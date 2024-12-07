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

		if slices.Contains(oldLocations, guardLocation) {

			fmt.Printf("Guard Location: %v \n", guardLocation)
			if len(oldLocations) < 4 {
				return false
			}

			var found int
			for index, val := range oldLocations {
				if val == guardLocation {
					found = index
				}
			}
			if found < 3 || found > len(oldLocations)-3 {
				return false
			}

			fmt.Printf("Possible loop locations: %v \n", oldLocations[found-3:found+4])

			for i := found; i < found+4; i++ {
				locations := oldLocations[i-3 : i+1]
				if checkIfRectangle(locations) {
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

func checkIfRectangle(oldLocations [][2]int) bool {
	first := oldLocations[0]
	last := oldLocations[3]

	return first[0] == last[0] || first[1] == last[1]

}
