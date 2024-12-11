package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Kazalo11/advent-of-code-2024/util"
)

func main() {
	// part1()
	part2()

}

func readData() []int {
	file, err := os.Open("day11.txt")
	util.Check(err)

	scanner := bufio.NewScanner(file)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}

	ans := make([]int, 0)

	for _, item := range strings.Split(line, " ") {
		num, err := strconv.Atoi(item)
		if err != nil {
			panic(fmt.Sprintf("Can't convert %s to number", item))
		}
		ans = append(ans, num)
	}

	return ans

}

func calculateNumberOfDigits(num int) int {
	count := 0
	for num > 0 {
		num /= 10
		count++
	}
	return count
}

func splitNumber(num int) []int {
	length := calculateNumberOfDigits(num)
	num2 := 0
	count := 0
	for i := 0; i < length/2; i++ {
		val := num % 10
		num2 += int(math.Pow10(count)) * val
		num /= 10

		count++

	}
	return []int{num, num2}

}

func blink(arr []int) []int {
	ans := make([]int, 0)

	for _, item := range arr {
		if item == 0 {
			ans = append(ans, 1)
		} else if calculateNumberOfDigits(item)%2 == 0 {
			ans = append(ans, splitNumber(item)...)
		} else {
			ans = append(ans, item*2024)
		}

	}
	return ans
}

func blink2(input map[int]int) map[int]int {
	new_input := make(map[int]int)

	// Extract keys and sort them
	keys := make([]int, 0, len(input))
	for key := range input {
		keys = append(keys, key)
	}
	sort.Ints(keys) // Sort keys in ascending order

	// Process keys in sorted order
	for _, key := range keys {
		val := input[key]
		if key == 0 {
			new_input[1] = val
		} else if calculateNumberOfDigits(key)%2 == 0 {
			digits := splitNumber(key)
			left := digits[0]
			right := digits[1]

			_, exists1 := new_input[left]
			if !exists1 {
				new_input[left] = val
			} else {
				new_input[left] += val
			}

			_, exists2 := new_input[right]
			if !exists2 {
				new_input[right] = val
			} else {
				new_input[right] += val
			}
		} else {
			_, exists := new_input[key*2024]
			if !exists {
				new_input[key*2024] = val
			} else {
				new_input[key*2024] += val
			}
		}
	}
	return new_input
}

func part1() {
	now := time.Now()
	data := readData()

	for i := 0; i < 25; i++ {
		fmt.Println(i)
		data = blink(data)
	}

	fmt.Println(len(data))
	fmt.Println(time.Since(now))

}

func part2() {
	data := readData()

	input := make(map[int]int)

	for _, val := range data {
		input[val] = 1

	}
	for i := 0; i < 75; i++ {
		input = blink2(input)
	}
	count := 0
	for _, val := range input {
		count += val
	}
	fmt.Println(count)

}
