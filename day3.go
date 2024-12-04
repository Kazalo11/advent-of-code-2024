package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getData() string {
	data, err := os.ReadFile("day3.txt")
	check(err)
	return string(data)
}

func multiplyNumbers(numbers []string) int {
	if len(numbers) > 2 {
		log.Fatalf("More than 2 numbers found: %v", numbers)
	}
	n1, err := strconv.Atoi(numbers[0])
	check(err)
	n2, err := strconv.Atoi(numbers[1])
	check(err)
	return n1 * n2
}

func calculateTotal1(matches []string) int {
	total := 0
	re2 := regexp.MustCompile(`\d+`)
	for _, match := range matches {
		numbers := re2.FindAllString(match, -1)

		total += multiplyNumbers(numbers)

	}
	return total
}

func calculateTotal2(matches []string) int {
	fmt.Printf("matches are: %v \n", matches)
	total := 0
	re2 := regexp.MustCompile(`\d+`)
	doMatch := true
	for _, match := range matches {
		fmt.Printf("doMatch is %v \n", doMatch)
		if match == "do()" {
			doMatch = true
		} else if match == "don't()" {
			doMatch = false
		} else {
			if doMatch {
				numbers := re2.FindAllString(match, -1)
				total += multiplyNumbers(numbers)
			}
		}

	}
	return total

}

func part1() {
	info := getData()

	re := regexp.MustCompile(`mul\(\d+,\d+\)`)

	matches := re.FindAllString(info, -1)

	total := calculateTotal1(matches)

	fmt.Printf("Total for part 1 is: %d \n", total)
}

func part2() {
	info := getData()

	re := regexp.MustCompile(`mul\(\d+,\d+\)|don't\(\)|do\(\)`)

	matches := re.FindAllString(info, -1)

	total := calculateTotal2(matches)
	fmt.Printf("Total for part 2 is: %d \n", total)

}

func main() {
	part1()
	part2()
}
