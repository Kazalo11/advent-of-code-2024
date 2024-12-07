package main

import (
	"fmt"
)

func part1() float64 {
	count := 0.0
	data := getData()
	for total, numbers := range data {
		if checkIfTestValue(total, numbers) {
			count += total
		} else {
			fmt.Printf("No combination can be found for %v to equal %f \n", numbers, total)
		}

	}
	return count

}

func checkIfTestValue(total float64, numbers []float64) bool {
	if len(numbers) < 2 {
		return false
	}

	return runningTotal(total, numbers, 0, numbers[0], "+") || runningTotal(total, numbers, 0, numbers[0], "*")

}

func runningTotal(total float64, numbers []float64, index int, curr float64, operator string) bool {
	switch operator {
	case "*":
		curr *= numbers[index]
	case "+":
		curr += numbers[index]

	}
	if curr == total {
		return true
	}
	if index == len(numbers)-1 {
		return false
	}
	return runningTotal(total, numbers, index+1, curr, "*") || runningTotal(total, numbers, index+1, curr, "+")

}

func generateCombinations(n int, combination []string, result *[][]string) {
	if len(combination) == n {
		combinationCopy := append([]string{}, combination...)
		*result = append(*result, combinationCopy)
		return
	}

	generateCombinations(n, append(combination, "+"), result)
	generateCombinations(n, append(combination, "*"), result)
}
