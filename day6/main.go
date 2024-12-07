package main

import "fmt"

func main() {
	path := part1()
	fmt.Printf("Answer to part 1: %d \n", len(path))

	// ans2 := part2()
	// fmt.Printf("Answer to part 2: %d \n", ans2)

	ans2 := part2take2(path)
	fmt.Printf("Answer to part 2: %d \n", ans2)
}
