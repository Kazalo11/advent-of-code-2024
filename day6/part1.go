package main

import "fmt"

func part1() int {
	direction := UP
	matrix := NewMatrix("day6/day6.txt")
	guardLocation := matrix.FindGuard()
	obstacles := matrix.FindObstacles()

	path := make([][2]int, 0)
	var obstacle [2]int
	var escape bool

	for {
		switch direction {
		case UP:

			obstacle, escape = NextObstacleUp(guardLocation, obstacles)

			for i := guardLocation[0]; i > obstacle[0]; i-- {
				path = append(path, [2]int{
					i,
					guardLocation[1],
				})
			}
			guardLocation = [2]int{
				obstacle[0] + 1,
				guardLocation[1],
			}

		case DOWN:

			obstacle, escape = NextObstacleDown(guardLocation, obstacles)

			for i := guardLocation[0]; i < obstacle[0]; i++ {
				path = append(path, [2]int{
					i,
					guardLocation[1],
				})
			}
			guardLocation = [2]int{
				obstacle[0] - 1,
				guardLocation[1],
			}

		case LEFT:
			obstacle, escape = NextObstacleLeft(guardLocation, obstacles)

			for i := guardLocation[1]; i > obstacle[1]; i-- {
				path = append(path, [2]int{
					guardLocation[0],
					i,
				})
			}
			guardLocation = [2]int{
				guardLocation[0],
				obstacle[1] + 1,
			}

		case RIGHT:
			obstacle, escape = NextObstacleRight(guardLocation, obstacles)

			for i := guardLocation[1]; i < obstacle[1]; i++ {
				path = append(path, [2]int{
					guardLocation[0],
					i,
				})
			}
			guardLocation = [2]int{
				guardLocation[0],
				obstacle[1] - 1,
			}

		}
		fmt.Printf("New guard location: %v \n", guardLocation)
		fmt.Printf("Path: %v \n", path)

		if escape {
			return GetDistinctItems(path)
		}
		direction = (direction + 1) % 4
	}
}
