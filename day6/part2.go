package main

import (
	"fmt"
	"math"
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
			fmt.Printf("Checking if obstacle at %v works \n", item)
			if item != guardLocation && !slices.Contains(obstacles, item) {
				obstacles2 := obstacles
				obstacles2 = append(obstacles2, item)
				if checkIfLoop(obstacles2, guardLocation) {
					fmt.Printf("Loop found if obstacle is put at row %d and col %d \n", i, j)
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
		fmt.Printf("Total old locations: %v \n", oldLocations)
		fmt.Printf("New guard location: %v \n", guardLocation)
		if escape {
			return false
		}

		if slices.Contains(oldLocations, guardLocation) {
			fmt.Printf("Found a possible loop %v \n", oldLocations)
			if len(oldLocations) < 4 {
				return false
			}
			first := oldLocations[len(oldLocations)-4]
			second := oldLocations[len(oldLocations)-3]
			third := oldLocations[len(oldLocations)-2]
			fourth := oldLocations[len(oldLocations)-1]

			return checkIfRectangle(first, second, third, fourth)
		}
		oldLocations = append(oldLocations, guardLocation)

		direction = (direction + 1) % 4
	}
}

func checkIfRectangle(first, second, third, fourth [2]int) bool {
	cx := float64(first[0]+second[0]+third[0]+fourth[0]) / 4.0
	cy := float64(first[1]+second[1]+third[1]+fourth[1]) / 4.0

	dd1 := math.Pow(cx-float64(first[0]), 2) + math.Pow(cy-float64(first[1]), 2)
	dd2 := math.Pow(cx-float64(second[0]), 2) + math.Pow(cy-float64(second[1]), 2)
	dd3 := math.Pow(cx-float64(third[0]), 2) + math.Pow(cy-float64(third[1]), 2)
	dd4 := math.Pow(cx-float64(fourth[0]), 2) + math.Pow(cy-float64(fourth[1]), 2)

	return dd1 == dd2 && dd1 == dd3 && dd1 == dd4

}
