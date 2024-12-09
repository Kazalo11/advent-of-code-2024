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

func constructBlocks(line string) []rune {
	var ans []rune
	id := 0
	for index, charRune := range line {
		char := string(charRune)
		num, err := strconv.Atoi(char)
		if err != nil {
			panic(fmt.Sprintf("Can't convert %s to num", char))
		}
		if index%2 == 0 {

			for i := 0; i < num; i++ {
				ans = append(ans, rune(id))
			}

			id++
		} else {
			for i := 0; i < num; i++ {
				ans = append(ans, '.')
			}
		}

	}

	fmt.Println(len(ans))
	return ans
}

func calculateCheckSum(block []rune) int {
	ans := 0
	p1 := 0
	p2 := len(block) - 1
	for p1 < p2 {
		char := block[p1]
		if string(char) != "." {

			ans += (int(char) * p1)
		} else {
			char2 := block[p2]
			for string(char2) == "." {
				p2--
				char2 = block[p2]
			}

			ans += (int(char2) * p1)
			p2--

		}
		p1++

	}
	fmt.Println(p2)
	return ans
}

func part1() int {
	line := getData()

	block := constructBlocks(line)

	return calculateCheckSum(block)
}
