package main

import "fmt"

func part2() int {
	count := 0
	antinodes := make([][2]int, 0)

	for _, antennas := range ans {
		antinodes = append(antinodes, getAllAntinodes2(antennas)...)
	}
	seen := make(map[[2]int]bool)

	unique := make([][2]int, 0)
	for _, val := range antinodes {
		if _, exists := seen[val]; !exists {
			seen[val] = true
			unique = append(unique, val)
		}
	}

	for _, antinode := range unique {
		if antinode[0] > -1 && antinode[0] < width && antinode[1] > -1 && antinode[1] < height {
			count++
		}
	}

	return count

}
func getAllAntinodes2(antennas [][2]int) [][2]int {
	ans := make([][2]int, 0)
	for i := 0; i < len(antennas)-1; i++ {
		for j := i + 1; j < len(antennas); j++ {
			ans = append(ans, getAntinodes2(antennas[i], antennas[j])...)
		}
	}
	return ans

}

func getAntinodes2(a1, a2 [2]int) [][2]int {
	ans := make([][2]int, 0)
	ans = append(ans, a1)
	ans = append(ans, a2)

	diff0 := a1[0] - a2[0]

	diff1 := a1[1] - a2[1]

	an1 := [2]int{0, 0}
	an1[0] = diff0 + a1[0]
	an1[1] = diff1 + a1[1]

	an2 := [2]int{0, 0}

	an2[0] = a2[0] - diff0

	an2[1] = a2[1] - diff1
	fmt.Println("Inputs")

	fmt.Println(a1, a2)

	fmt.Println("First calculation")

	fmt.Println(an1, an2)

	for an1[0] > -1 && an1[0] < width && an1[1] > -1 && an1[1] < height {
		ans = append(ans, an1)
		an1[0] = diff0 + an1[0]
		an1[1] = diff1 + an1[0]
	}

	for an2[0] > -1 && an2[0] < width && an2[1] > -1 && an2[1] < height {
		ans = append(ans, an2)
		an2[0] = an2[0] - diff0
		an2[1] = an2[1] - diff1
	}

	fmt.Println(ans)
	return ans

}
