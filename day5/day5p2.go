package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/Kazalo11/advent-of-code-2024/util"
)

var count2 int = 0

func reOrder(arr []string) []string {
	arr2 := arr
	for i := 0; i < len(arr2)-1; i++ {
		poss := ordering[arr2[i]]
		for j := i + 1; j < len(arr2); j++ {
			curr := arr2[j]
			if !slices.Contains(poss, curr) {
				arr2[i], arr2[j] = arr2[j], arr2[i]
			}
		}

	}
	if isOrdered(arr2) {
		return arr2
	}
	return reOrder(arr2)
}

func part2() {
	file, err := os.Open("day5/day5.txt")
	util.Check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			switchCode = true
			continue
		}
		if !switchCode {
			getOrder(line)
		} else {
			arr := strings.Split(line, ",")

			if !isOrdered(arr) {
				new_arr := reOrder(arr)
				fmt.Println(new_arr)
				count2 += calculateMiddle(new_arr)
			}

		}
	}
	fmt.Printf("Total for part 2: %d \n", count2)

}
