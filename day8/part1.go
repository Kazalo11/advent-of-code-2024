package main

import (
	"bufio"
	"os"

	"github.com/Kazalo11/advent-of-code-2024/util"
)

var height int
var width int

func getData() map[string][][2]int {
	antennas := make(map[string][][2]int, 0)

	file, err := os.Open("day8.txt")
	util.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		var splits []string
		for _, char := range line {
			splits = append(splits, string(char))
		}
		width = len(splits)
		for col, val := range splits {
			if val != "." {
				_, exists := antennas[val]
				if !exists {
					var init [][2]int
					init = append(init, [2]int{height, col})
					antennas[val] = init
				} else {
					antennas[val] = append(antennas[val], [2]int{height, col})
				}
			}
		}
		height++
	}

	return antennas

}

func part1() int {
	count := 0
	ans := getData()
	antinodes := make([][2]int, 0)

	for _, antennas := range ans {
		antinodes = append(antinodes, getAllAntinodes(antennas)...)
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

func getAllAntinodes(antennas [][2]int) [][2]int {
	ans := make([][2]int, 0)
	for i := 0; i < len(antennas)-1; i++ {
		for j := i + 1; j < len(antennas); j++ {
			ans = append(ans, getAntinodes1(antennas[i], antennas[j])...)
		}
	}
	return ans

}

func getAntinodes1(a1, a2 [2]int) [][2]int {
	an1 := [2]int{0, 0}
	an2 := [2]int{0, 0}

	diff0 := a1[0] - a2[0]

	diff1 := a1[1] - a2[1]

	an1[0] = diff0 + a1[0]
	an1[1] = diff1 + a1[1]

	an2[0] = a2[0] - diff0

	an2[1] = a2[1] - diff1

	return [][2]int{an1, an2}

}
