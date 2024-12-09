package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
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

func calculateCheckSum2(block []string) int {
	p2 := len(block) - 1

	for p2 > 0 {
		curr := block[p2]
		if curr == "." {
			p2--
			continue
		}
		index := p2 - 1
		for block[index] == curr {
			index--
			if index < 0 {
				return calculateTotal(block)
			}
		}

		length := p2 - index
		numBlock := block[index+1 : index+length+1]
		p1 := 0
		dots := make([]string, length)
		for i := range dots {
			dots[i] = "."
		}
		for p1 < index {
			if reflect.DeepEqual(block[p1:p1+length], dots) {
				tmp := make([]string, length)
				copy(tmp, numBlock)
				copy(numBlock, block[p1:p1+length])
				copy(block[p1:p1+length], tmp)
				break
			}
			p1++
		}
		p2 = index

	}

	return calculateTotal(block)
}

func calculateTotal(block []string) int {
	ans := 0
	for idx, val := range block {
		if val != "." {
			num, err := strconv.Atoi(val)
			if err != nil {
				panic("Can't convert string to number")
			}

			ans += (num * idx)
		}

	}

	return ans
}

func part2() int {
	line := getData()

	block := constructBlock(line)

	return calculateCheckSum2(block)
}
