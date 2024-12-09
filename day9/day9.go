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

func constructBlocks(line string) string {
	var ans string
	id := 0
	for index, charRune := range line {
		char := string(charRune)
		num, err := strconv.Atoi(char)
		if err != nil {
			panic(fmt.Sprintf("Can't convert %s to num", char))
		}
		if index%2 == 0 {

			for i := 0; i < num; i++ {
				ans += fmt.Sprint(id)
			}

			id++
		} else {
			for i := 0; i < num; i++ {
				ans += "."
			}
		}

	}

	fmt.Println(len(ans))
	return ans
}

func calculateCheckSum(block string) int {
	ans := 0
	p1 := 0
	p2 := len(block) - 1
	for p1 <= p2 {
		char := string(block[p1])
		if char != "." {
			num, err := strconv.Atoi(char)
			if err != nil {
				panic(fmt.Sprintf("Can't convert %s to num", char))
			}
			ans += (num * p1)
		} else {
			char2 := string(block[p2])
			for char2 == "." {
				p2--
				char2 = string(block[p2])
			}
			num, err := strconv.Atoi(char2)
			if err != nil {
				panic(fmt.Sprintf("Can't convert %s to num", char2))
			}
			ans += (num * p1)
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
