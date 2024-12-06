package main

import (
	"bufio"
	"fmt"
	"os"
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

func part1() {
	total := getData()
	fmt.Printf("Total for part 1: %d \n", total)
}

func main() {
	part1()
	part2()
}
