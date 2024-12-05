package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func calculateMiddle(line []string) int {
	mid := len(line) / 2
	val, err := strconv.Atoi(line[mid])
	if err != nil {
		fmt.Printf("Line used: %v \n", line)
		fmt.Printf("Could not convert %s to number \n", line[mid])
	}
	return val
}

func getOrder(line string) {
	arr := strings.Split(line, "|")
	if len(arr) > 2 {
		panic("More than 2 values in the array")
	}
	before := arr[0]
	after := arr[1]
	_, exists := ordering[before]
	if !exists {
		ordering[before] = []string{after}
	} else {
		ordering[before] = append(ordering[before], after)
	}

}
func isOrdered(arr []string) bool {
	for i := 0; i < len(arr)-1; i++ {
		poss := ordering[arr[i]]
		for j := i + 1; j < len(arr); j++ {
			curr := arr[j]
			if !slices.Contains(poss, curr) {
				return false
			}

		}
	}
	return true
}
