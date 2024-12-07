package main

import (
	"math"
)

func part1() float64 {
	return getData()

}

func checkIfTestValue(total float64, numbers []float64) bool {
	if len(numbers) < 2 {
		return false
	}

	return runningTotal(total, numbers, 1, numbers[0], "+") || runningTotal(total, numbers, 1, numbers[0], "*") || runningTotal(total, numbers, 1, numbers[0], "||")

}

func concatNumbers(n1 float64, n2 float64) float64 {
	digits := int(math.Log10(n2)) + 1
	sol := math.Pow(10, float64(digits))*n1 + n2
	return sol

}

func runningTotal(total float64, numbers []float64, index int, curr float64, operator string) bool {
	switch operator {
	case "*":
		curr *= numbers[index]
	case "+":
		curr += numbers[index]

	case "||":
		curr = concatNumbers(curr, numbers[index])
	}
	if curr == total {
		return true
	}
	if curr > total {
		return false
	}
	if index == len(numbers)-1 {
		return false
	}
	return runningTotal(total, numbers, index+1, curr, "*") || runningTotal(total, numbers, index+1, curr, "+") || runningTotal(total, numbers, index+1, curr, "||")

}
