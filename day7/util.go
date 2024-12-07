package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Kazalo11/advent-of-code-2024/util"
)

func getData() map[float64][]float64 {
	file, err := os.Open("day7.txt")
	util.Check(err)

	equation := make(map[float64][]float64, 0)

	scanner := bufio.NewScanner(file)
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

		equation[end] = numbers

	}
	return equation

}

func checkIfPossible(numbers []float64, operators []string, desiredTotal float64) bool {
	total := numbers[0]
	if len(numbers) < 2 {
		return false
	}

	for i, operator := range operators {
		number := numbers[i+1]
		switch operator {
		case "+":
			total += number
		case "*":
			total *= number

		}
		if total > desiredTotal {
			return false
		}

	}
	return total == desiredTotal

}
