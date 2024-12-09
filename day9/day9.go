package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/Kazalo11/advent-of-code-2024/util"
)

func getData() string {
	file, err := os.Open("day9.txt")
	util.Check(err)
	scanner := bufio.NewScanner(file)
	var line string

	for scanner.Scan() {
		line = scanner.Text()
	}

	return line
}

func constructBlock(line string) []string {
	var ans []string
	id := 0
	for index, charRune := range line {
		char := string(charRune)
		num, err := strconv.Atoi(char)
		if err != nil {
			panic(fmt.Sprintf("Can't convert %s to num", char))
		}
		if index%2 == 0 {

			for i := 0; i < num; i++ {
				ans = append(ans, strconv.Itoa(id))
			}

			id++
		} else {
			for i := 0; i < num; i++ {
				ans = append(ans, ".")
			}
		}

	}

	return ans
}

func calculateCheckSum(block []string) int {
	ans := 0
	p1 := 0
	p2 := len(block) - 1
	for p1 <= p2 {
		char := block[p1]
		if char == "." {
			for block[p2] == "." {
				p2--
			}
			if p1 <= p2 {

				block[p1], block[p2] = block[p2], block[p1]
			}

		}
		// fmt.Println(p1, p2)
		// fmt.Println(block)

		p1++

	}
	for i, val := range block {
		if val != "." {
			num, err := strconv.Atoi(val)
			if err != nil {
				panic(fmt.Sprintf("Can't convert string %s to number", val))
			}
			ans += (num * i)
		}
	}

	return ans
}

func part1() int {
	line := getData()

	block := constructBlock(line)

	return calculateCheckSum(block)
}
