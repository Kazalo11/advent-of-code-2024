package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Kazalo11/advent-of-code-2024/util"
)

func getData() float64 {
	file, err := os.Open("day7.txt")
	util.Check(err)

	scanner := bufio.NewScanner(file)
	count := 0.0
	for scanner.Scan() {

		line := scanner.Text()
		lines := strings.Split(line, ":")
		end, err := strconv.ParseFloat(lines[0], 64)
		if err != nil {
			panic(fmt.Sprintf("Error converting string to number %v", lines[0]))
		}
		lines_trimmed := strings.TrimSpace(lines[1])
		numbersString := strings.Split(lines_trimmed, " ")
		numbers := make([]float64, 0)
		for _, numString := range numbersString {
			num, err := strconv.ParseFloat(numString, 64)
			if err != nil {
				panic(fmt.Sprintf("Error converting string to number %v", numString))
			}
			numbers = append(numbers, num)
		}
		if checkIfTestValue(end, numbers) {
			count += end
		}

	}
	return count

}
