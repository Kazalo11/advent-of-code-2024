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
	num_of_operators := len(numbers) - 1
	operators := [][]string{}

	generateCombinations(num_of_operators, []string{}, &operators)
	expected := 1 << num_of_operators
	if len(operators) != expected {
		panic(fmt.Sprintf("Not enough operators. Generated %d, expected %d.\n", len(operators), expected))
	}

	for _, operator := range operators {
		// fmt.Printf("Checking for operators: %v if they can be used with numbers %v to get total %d \n", operator, numbers, total)
		if checkIfPossible(numbers, operator, total) {
			fmt.Printf("Operators %v work for total %f \n", operator, total)
			return true
		}

	}
	return false

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
