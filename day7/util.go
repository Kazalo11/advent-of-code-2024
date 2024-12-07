package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getData() int64 {
	// Open the file
	file, err := os.Open("day7.txt")
	if err != nil {
		panic(fmt.Sprintf("Error opening file: %v", err))
	}
	defer file.Close()

	// Initialize total count
	var count int64 = 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Read and parse each line
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			panic(fmt.Sprintf("Malformed line: %v", line))
		}

		// Parse the target value
		end, err := strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 64)
		if err != nil {
			panic(fmt.Sprintf("Error parsing target number: %v", parts[0]))
		}

		// Parse the number sequence
		numberStrings := strings.Fields(parts[1])
		numbers := make([]int64, 0, len(numberStrings))
		for _, numStr := range numberStrings {
			num, err := strconv.ParseInt(numStr, 10, 64)
			if err != nil {
				panic(fmt.Sprintf("Error parsing number in sequence: %v", numStr))
			}
			numbers = append(numbers, num)
		}

		// Check if the equation can be satisfied
		if checkIfTestValue(end, numbers) {
			count += end
		}
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading file: %v", err))
	}

	return count
}
