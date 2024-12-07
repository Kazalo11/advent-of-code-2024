package main

import (
	"math"
)

func part1() int64 {
	return getData()
}

func checkIfTestValue(total int64, numbers []int64) bool {
	if len(numbers) == 1 {
		return numbers[0] == total
	}

	return runningTotal(total, numbers, 1, numbers[0], "+") || runningTotal(total, numbers, 1, numbers[0], "*") || runningTotal(total, numbers, 1, numbers[0], "||")
}

func concatNumbers(n1 int64, n2 int64) int64 {
	digits := 0
	temp := n2
	for temp > 0 {
		temp /= 10
		digits++
	}
	powerOf10 := int64(1)
	for i := 0; i < digits; i++ {
		powerOf10 *= 10
	}
	if n1 > math.MaxInt64/powerOf10 {
		panic("int64 overflow in concatNumbers")
	}
	return n1*powerOf10 + n2
}

func runningTotal(total int64, numbers []int64, index int, curr int64, operator string) bool {
	switch operator {
	case "*":
		curr *= numbers[index]
	case "+":
		curr += numbers[index]
	case "||":
		curr = concatNumbers(curr, numbers[index])
	}
	if curr == total && index == len(numbers)-1 {
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
