package main

import "fmt"

func p2take2(path [][2]int) int {
	matrix := NewMatrix("day6.txt")
	obstacles := matrix.FindObstacles()
	guardLocation := matrix.FindGuard()
	count := 0

	for _, item := range path {
		if item != guardLocation {
			obstacles2 := append([][2]int{}, obstacles...)
			obstacles2 = append(obstacles2, item)
			if checkIfLoop(obstacles2, guardLocation) {
				fmt.Printf("Loop found if obstacle is put at row %d and col %d \n", item[0], item[1])
				count++
			}
		}
	}
	return count

}
