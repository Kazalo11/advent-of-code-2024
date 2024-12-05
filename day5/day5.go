package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/Kazalo11/advent-of-code-2024/util"
)

var ordering map[string][]string = make(map[string][]string)
var switchCode bool = false
var count int = 0

func getData() int {
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
			if isOrdered(arr) {
				count += calculateMiddle(arr)
			}
		}

	}
	return count
}

func getOrder(line string) {
	arr := strings.Split(line, "|")
	if len(arr) > 2 {
		panic("More than 2 values in the array")
	}
	before := arr[0]
	after := arr[1]
	_, exists := ordering[before]
	if !exists {
		ordering[before] = []string{after}
	} else {
		ordering[before] = append(ordering[before], after)
	}

}

func isOrdered(arr []string) bool {
	for i := 0; i < len(arr)-1; i++ {
		poss := ordering[arr[i]]
		for j := i + 1; j < len(arr); j++ {
			curr := arr[j]
			if !slices.Contains(poss, curr) {
				return false
			}

		}
	}
	return true
}

func calculateMiddle(line []string) int {
	mid := len(line) / 2
	val, err := strconv.Atoi(line[mid])
	if err != nil {
		fmt.Printf("Line used: %v \n", line)
		fmt.Printf("Could not convert %s to number \n", line[mid])
	}
	return val
}

func part1() {
	total := getData()
	fmt.Printf("Total for part 1: %d \n", total)
}

func main() {
	part1()
}
