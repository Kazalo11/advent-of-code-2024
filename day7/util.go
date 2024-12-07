package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Kazalo11/advent-of-code-2024/util"
)

func getData() map[int][]int {
	file, err := os.Open("day7.txt")
	util.Check(err)

	equation := make(map[int][]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines := strings.Split(line, ":")
		end, err := strconv.Atoi(lines[0])
		if err != nil {
			panic(fmt.Sprintf("Error converting string to number %v", lines[0]))
		}
		lines_trimmed := strings.TrimSpace(lines[1])
		numbersString := strings.Split(lines_trimmed, " ")
		numbers := make([]int, 0)
		for _, numString := range numbersString {
			num, err := strconv.Atoi(numString)
			if err != nil {
				panic(fmt.Sprintf("Error converting string to number %v", numString))
			}
			numbers = append(numbers, num)
		}

		equation[end] = numbers

	}
	return equation

}

func checkIfPossible(numbers []int, operators []string, desiredTotal int) bool {
	total := numbers[0]
	if len(numbers) < 2 {
		return false
	}
	i := 1

	for _, operator := range operators {
		number := numbers[i]
		switch operator {
		case "+":

			total += number
		case "*":
			total *= number

		}
		i++
		if total > desiredTotal {
			return false
		}

	}
	return total == desiredTotal

}
