package main

import (
	"fmt"
	"math"
)

func part1() int {
	count := 0
	data := getData()
	for total, numbers := range data {
		fmt.Printf("Checking for total %d if numbers %v work \n", total, numbers)
		if checkIfTestValue(total, numbers) {
			count += total
		}

	}
	return count

}

func checkIfTestValue(total int, numbers []int) bool {
	num_of_operators := len(numbers) - 1
	operators := [][]string{}

	generateCombinations(num_of_operators, []string{}, &operators)
	if len(operators) != int(math.Pow(2, (float64(num_of_operators)))) {
		panic(fmt.Sprintf("Not enough operators, len of operators is %d when it should be %d \n", len(operators), int(math.Pow(2, (float64(num_of_operators))))))
	}

	for _, operator := range operators {
		if checkIfPossible(numbers, operator, total) {
			fmt.Printf("Operators %v work for total %d \n", operator, total)
			return true
		}

	}
	return false

}

func generateCombinations(n int, combination []string, result *[][]string) {
	if len(combination) == n {
		combinationCopy := make([]string, len(combination))
		copy(combinationCopy, combination)
		*result = append(*result, combinationCopy)
		return
	}

	generateCombinations(n, append(combination, "+"), result)
	generateCombinations(n, append(combination, "*"), result)
}
